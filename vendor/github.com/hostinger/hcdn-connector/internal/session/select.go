package session

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/hostinger/hcdn-connector/pkg/config"
)

// Candidate is a probed PoP as seen by the selector.
type Candidate struct {
	Code    string
	Region  string
	Allowed bool
	Pinned  bool
	RTT     time.Duration // 0 with Err set = unreachable
	Err     string
}

func (c Candidate) reachable() bool { return c.Err == "" }

// Current describes a PoP that has a supervisor now.
type Current struct {
	Code  string
	Since time.Time
	Ready bool // at least one ready session
	Down  bool // no ready session for popDownAfter outside a pause
}

// Policy is the selection part of the config.
type Policy struct {
	HomePops   int
	MinGainPct int
	MinGainMs  int
	Holddown   time.Duration
	Mode       string
	ManualPops []string
}

// Select picks the home PoPs. It returns the
// ordered selection and a human-readable reason for `status`.
func Select(cands []Candidate, current []Current, pol Policy, now time.Time) ([]string, string) {
	byCode := map[string]Candidate{}
	for _, c := range cands {
		byCode[c.Code] = c
	}
	if pol.Mode == config.ModeManual {
		var sel, skipped []string
		for _, code := range pol.ManualPops {
			if len(sel) >= pol.HomePops {
				break
			}
			c, ok := byCode[code]
			switch {
			case !ok:
				skipped = append(skipped, code+" (not in pop list)")
			case !c.Allowed:
				skipped = append(skipped, code+" (not allowed)")
			default:
				sel = append(sel, code)
			}
		}
		reason := "manual: " + strings.Join(sel, ",")
		if len(skipped) > 0 {
			reason += "; skipped " + strings.Join(skipped, ", ")
		}
		return sel, reason
	}

	curBy := map[string]Current{}
	for _, c := range current {
		curBy[c.Code] = c
	}
	var reasons []string
	selected := []string{}
	regions := map[string]bool{}
	add := func(code, why string) {
		selected = append(selected, code)
		regions[byCode[code].Region] = true
		reasons = append(reasons, code+": "+why)
	}
	has := func(code string) bool {
		for _, s := range selected {
			if s == code {
				return true
			}
		}
		return false
	}

	// 1. Pinned PoPs first (fastest first, unreachable last).
	var pinned []Candidate
	for _, c := range cands {
		if c.Pinned && c.Allowed {
			pinned = append(pinned, c)
		}
	}
	sortByRTT(pinned)
	for _, c := range pinned {
		if len(selected) >= pol.HomePops {
			break
		}
		add(c.Code, "pinned")
	}

	// 2. Keep current PoPs that are still allowed and reachable (or carrying
	//    sessions). A down PoP is only a last resort (step 3b).
	var kept, down []Current
	for _, cur := range current {
		c, ok := byCode[cur.Code]
		if !ok || !c.Allowed {
			reasons = append(reasons, cur.Code+": dropped (no longer allowed)")
			continue
		}
		if !c.reachable() && !cur.Ready {
			reasons = append(reasons, cur.Code+": dropped (unreachable: "+c.Err+")")
			continue
		}
		if cur.Down {
			down = append(down, cur)
			continue
		}
		kept = append(kept, cur)
	}
	isDown := func(code string) bool {
		for _, d := range down {
			if d.Code == code {
				return true
			}
		}
		return false
	}
	sort.Slice(kept, func(i, j int) bool { return byCode[kept[i].Code].RTT < byCode[kept[j].Code].RTT })
	for _, cur := range kept {
		if len(selected) >= pol.HomePops {
			break
		}
		if !has(cur.Code) {
			add(cur.Code, "current")
		}
	}

	// 3. Fill from the fastest reachable candidates, preferring new regions.
	var pool []Candidate
	for _, c := range cands {
		if c.Allowed && c.reachable() && !has(c.Code) && !isDown(c.Code) {
			pool = append(pool, c)
		}
	}
	sortByRTT(pool)
	for len(selected) < pol.HomePops && len(pool) > 0 {
		idx := 0
		for i, c := range pool {
			if !regions[c.Region] {
				idx = i
				break
			}
		}
		c := pool[idx]
		pool = append(pool[:idx], pool[idx+1:]...)
		why := fmt.Sprintf("fastest (%.1f ms)", ms(c.RTT))
		if _, wasCur := curBy[c.Code]; wasCur {
			why = "current"
		}
		add(c.Code, why)
	}

	// 3b. A down PoP beats an empty slot; apply drops it once a replacement is ready.
	for _, cur := range down {
		switch {
		case has(cur.Code): // pinned
		case len(selected) < pol.HomePops:
			add(cur.Code, "current (no ready session; no alternative)")
		default:
			reasons = append(reasons, cur.Code+": dropped (no ready session)")
		}
	}

	// 4. One switch per round: replace the slowest non-pinned current PoP past
	//    its holddown when a candidate beats it by both thresholds.
	if len(pool) > 0 {
		worstIdx := -1
		for i, code := range selected {
			cur, isCur := curBy[code]
			c := byCode[code]
			if !isCur || c.Pinned || !c.reachable() {
				continue
			}
			if now.Sub(cur.Since) < pol.Holddown {
				continue
			}
			if worstIdx < 0 || c.RTT > byCode[selected[worstIdx]].RTT {
				worstIdx = i
			}
		}
		best, ok := Candidate{}, false
		if worstIdx >= 0 {
			best, ok = fastestKeepingRegions(pool, selected, worstIdx, byCode)
		}
		if ok {
			cur := byCode[selected[worstIdx]]
			gain := cur.RTT - best.RTT
			pct := 0.0
			if cur.RTT > 0 {
				pct = float64(gain) / float64(cur.RTT) * 100
			}
			// gain > 0: with both thresholds at 0, a tie must not swap PoPs every reprobe.
			if gain > 0 && gain >= time.Duration(pol.MinGainMs)*time.Millisecond && pct >= float64(pol.MinGainPct) {
				reasons = append(reasons, fmt.Sprintf("switch %s -> %s (%.1f ms -> %.1f ms, %.0f%% / %.1f ms gain)",
					cur.Code, best.Code, ms(cur.RTT), ms(best.RTT), pct, ms(gain)))
				selected[worstIdx] = best.Code
			}
		}
	}
	return selected, strings.Join(reasons, "; ")
}

// fastestKeepingRegions returns the fastest pool candidate that, replacing
// selected[idx], leaves at least as many distinct regions selected.
func fastestKeepingRegions(pool []Candidate, selected []string, idx int, byCode map[string]Candidate) (Candidate, bool) {
	before, rest := map[string]bool{}, map[string]bool{}
	for i, code := range selected {
		r := byCode[code].Region
		before[r] = true
		if i != idx {
			rest[r] = true
		}
	}
	for _, c := range pool {
		after := len(rest)
		if !rest[c.Region] {
			after++
		}
		if after >= len(before) {
			return c, true
		}
	}
	return Candidate{}, false
}

func sortByRTT(cs []Candidate) {
	sort.SliceStable(cs, func(i, j int) bool {
		if cs[i].reachable() != cs[j].reachable() {
			return cs[i].reachable()
		}
		return cs[i].RTT < cs[j].RTT
	})
}

func ms(d time.Duration) float64 { return float64(d) / float64(time.Millisecond) }
