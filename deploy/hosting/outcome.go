package hosting

// Outcome is what `hostinger up` prints.
type Outcome struct {
	URL         string `json:"url"`
	Target      string `json:"target"`
	ProjectFile string `json:"projectFile,omitempty"`
	// NewLink is true when this deploy linked the directory.
	NewLink bool `json:"newLink,omitempty"`
	DryRun  bool `json:"dryRun,omitempty"`
	// Note replaces the "live" line when the deploy is not finished yet.
	Note string `json:"note,omitempty"`
	// Reachable is false when URL cannot answer yet (no DNS); ReachableNote says what to do.
	Reachable     bool   `json:"reachable"`
	ReachableNote string `json:"reachableNote,omitempty"`
	Data          any    `json:"data,omitempty"`
}
