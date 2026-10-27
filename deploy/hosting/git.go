package hosting

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"sort"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

type Source string

const (
	SourceUpload Source = "upload" // pack the directory and upload it
	SourceGit    Source = "git"    // the pushed branch, through a connected Git account
)

func ParseSource(s string) (Source, error) {
	switch Source(s) {
	case SourceUpload, SourceGit:
		return Source(s), nil
	}
	return "", fmt.Errorf("unknown source %q (one of: upload, git)", s)
}

type Repo struct {
	Host       string // github.com, gitlab.com, ...
	Owner      string
	Name       string
	Branch     string
	installUID string
}

func (r *Repo) FullName() string { return r.Owner + "/" + r.Name }

// Installation is a Git account connected in hPanel.
type Installation struct {
	UUID         string `json:"uuid"`
	Provider     string `json:"provider"`
	AccountLogin string `json:"account_login"`
	Status       string `json:"status"`
}

type Repository struct {
	Name          string `json:"name"`
	FullName      string `json:"full_name"`
	Owner         string `json:"owner"`
	DefaultBranch string `json:"default_branch"`
}

// GitSettings are a static website's auto-deploy settings.
type GitSettings struct {
	InstallationUUID string `json:"installation_uuid"`
	Owner            string `json:"owner"`
	Repository       string `json:"repository"`
	Branch           string `json:"branch"`
	IsEnabled        bool   `json:"is_enabled"`
}

// openOptions let a linked worktree see the main repository's remotes.
var openOptions = &git.PlainOpenOptions{EnableDotGitCommonDir: true}

// DetectRepo reads the origin remote and branch; nil, nil when dir is not a
// repository root with an origin.
func DetectRepo(dir string) (*Repo, error) {
	repo, err := git.PlainOpenWithOptions(dir, openOptions)
	if errors.Is(err, git.ErrRepositoryNotExists) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	remote, err := repo.Remote("origin")
	if err != nil {
		return nil, nil
	}
	urls := remote.Config().URLs
	if len(urls) == 0 {
		return nil, nil
	}
	host, owner, name, ok := parseRemote(urls[0])
	if !ok {
		return nil, nil
	}
	head, err := repo.Head()
	if err != nil {
		return nil, fmt.Errorf("reading HEAD: %w", err)
	}
	if !head.Name().IsBranch() {
		return nil, errors.New("the repository is in detached HEAD state; check out a branch to deploy from Git")
	}
	return &Repo{Host: host, Owner: owner, Name: name, Branch: head.Name().Short()}, nil
}

// unpushedWork describes what the pushed branch is missing, or "" when nothing
// (or unknown). Files this command writes are not counted.
func unpushedWork(dir, branch string) string {
	repo, err := git.PlainOpenWithOptions(dir, openOptions)
	if err != nil {
		return ""
	}
	var problems []string
	head, err := repo.Head()
	if err == nil {
		remoteRef, rerr := repo.Reference(plumbing.NewRemoteReferenceName("origin", branch), true)
		switch {
		case errors.Is(rerr, plumbing.ErrReferenceNotFound):
			problems = append(problems, "branch "+branch+" has never been pushed")
		case rerr == nil && remoteRef.Hash() != head.Hash() && !contains(repo, remoteRef.Hash(), head.Hash()):
			problems = append(problems, "commits not pushed to origin/"+branch)
		}
	}
	if wt, err := repo.Worktree(); err == nil {
		if status, err := wt.Status(); err == nil {
			for path, st := range status {
				if st.Worktree == git.Unmodified && st.Staging == git.Unmodified {
					continue
				}
				if path == ".gitignore" || path == IgnoreFile || strings.HasPrefix(path, ".hostinger/") {
					continue
				}
				problems = append(problems, "uncommitted changes")
				break
			}
		}
	}
	return strings.Join(problems, " and ")
}

// contains reports whether commit tip includes commit c, so a checkout that
// is only behind origin is not told to push.
func contains(repo *git.Repository, tip, c plumbing.Hash) bool {
	tc, err := repo.CommitObject(tip)
	if err != nil {
		return false
	}
	cc, err := repo.CommitObject(c)
	if err != nil {
		return false
	}
	ok, err := cc.IsAncestor(tc)
	return err == nil && ok
}

func parseRemote(raw string) (host, owner, name string, ok bool) {
	raw = strings.TrimSpace(raw)
	var path string
	if strings.HasPrefix(raw, "git@") || (strings.Contains(raw, ":") && !strings.Contains(raw, "://")) {
		// git@github.com:owner/name.git, or github.com:owner/name.git without a user
		at := strings.Index(raw, "@") // -1 when there is no user
		colon := strings.Index(raw, ":")
		if colon < at+1 {
			return "", "", "", false
		}
		host = raw[at+1 : colon]
		path = raw[colon+1:]
	} else {
		u, err := url.Parse(raw)
		if err != nil || u.Host == "" {
			return "", "", "", false
		}
		host = u.Hostname()
		path = u.Path
	}
	path = strings.TrimSuffix(strings.Trim(path, "/"), ".git")
	parts := strings.Split(path, "/")
	if len(parts) < 2 {
		return "", "", "", false
	}
	// GitLab groups can nest: owner is everything before the last segment.
	return strings.ToLower(host), strings.Join(parts[:len(parts)-1], "/"), parts[len(parts)-1], true
}

// findInstallation returns nil, nil when no connected Git account can read the repo.
func findInstallation(ctx context.Context, a API, r *Repo) (*Installation, error) {
	installations, err := a.ListGitInstallations(ctx)
	if err != nil {
		return nil, err
	}
	// The same owner/name can exist on GitHub and GitLab, so the remote's
	// provider is searched first. It is a preference, not a filter: an SSH
	// alias (github-work) or ssh.github.com does not spell it out.
	onHost := func(inst Installation) bool {
		return inst.Provider != "" && strings.Contains(r.Host, strings.ToLower(inst.Provider))
	}
	sort.SliceStable(installations, func(i, j int) bool { return onHost(installations[i]) && !onHost(installations[j]) })
	for i := range installations {
		inst := &installations[i]
		if inst.Status != "" && inst.Status != "active" {
			continue
		}
		repos, err := a.ListInstallationRepositories(ctx, inst.UUID)
		if err != nil {
			return nil, err
		}
		for _, repo := range repos {
			if strings.EqualFold(repo.FullName, r.FullName()) {
				r.Owner, r.Name = repo.Owner, repo.Name
				r.installUID = inst.UUID
				return inst, nil
			}
		}
	}
	return nil, nil
}

// HPanelIntegrations is where Git accounts are connected (Account integrations).
const HPanelIntegrations = "https://hpanel.hostinger.com/profile/personal-information"

const ConnectHint = "Connect GitHub or GitLab once in hPanel under Account integrations (" + HPanelIntegrations + "), give it access to the repository, then run this again"

// connectGit guides linking a Git account in hPanel, re-checking on request
// until the repo is visible; nil, nil when the person stops.
func connectGit(ctx context.Context, a API, ui UI, repo *Repo) (*Installation, error) {
	ui.Info("To deploy from Git, link a Git account to your hPanel account. On the Account integrations page:")
	ui.Info("  - connect GitHub or GitLab and give it access to all repositories, or at least to " + repo.FullName() + ";")
	ui.Info("  - an account that is already linked gets more repositories through Manage in GitHub.")
	open, err := ui.Confirm("Open "+HPanelIntegrations+" in your browser now?", true)
	if err != nil {
		return nil, err
	}
	if open && OpenBrowser(HPanelIntegrations) == nil {
		ui.Info("Opened hPanel. Come back here once the account is linked.")
	} else {
		ui.Info("Open " + HPanelIntegrations + " yourself (Account information, Account integrations), then come back here.")
	}
	for {
		ok, err := ui.Confirm("Done linking? Check again for a linked Git account that can read "+repo.FullName(), true)
		if err != nil {
			return nil, err
		}
		if !ok {
			return nil, nil
		}
		inst, err := findInstallation(ctx, a, repo)
		if err != nil {
			return nil, err
		}
		if inst != nil {
			return inst, nil
		}
		ui.Info("Still no Git account linked in your hPanel account can read " + repo.FullName() + ". Check that the linked account was given access to it.")
	}
}

func gitSourceLabel(r *Repo) string {
	return fmt.Sprintf("Deploy from Git (%s, branch %s; redeploys on every push)", r.FullName(), r.Branch)
}

// chooseSource picks upload or Git: --source, else remembered, else asked
// when a connected Git account can read the checkout. Without a terminal,
// or in a dry run, it is upload.
func chooseSource(ctx context.Context, a API, ui UI, opts Options, src string, isDir bool, remembered Source) (Source, *Repo, error) {
	if !isDir {
		if opts.Source == SourceGit {
			return "", nil, errors.New("an archive cannot be deployed from Git; pass a directory")
		}
		return SourceUpload, nil, nil
	}
	want := opts.Source
	if want == "" {
		want = remembered
	}
	if want == SourceUpload {
		return SourceUpload, nil, nil
	}

	// A detached HEAD (most CI checkouts) only matters when Git was asked for.
	repo, err := DetectRepo(src)
	if err != nil {
		if want == SourceGit {
			return "", nil, err
		}
		return SourceUpload, nil, nil
	}
	if repo == nil {
		if want == SourceGit {
			return "", nil, errors.New(src + " is not the root of a Git repository with an origin remote; deploy with --source upload")
		}
		return SourceUpload, nil, nil
	}
	if want == "" && (!ui.Interactive() || opts.DryRun) {
		return SourceUpload, nil, nil
	}
	inst, err := findInstallation(ctx, a, repo)
	if err != nil {
		if want == SourceGit {
			return "", nil, err
		}
		// Git was only on offer: a rate limit or an expired installation must not block an upload.
		ui.Info("Could not check the Git accounts linked in hPanel (" + err.Error() + "); uploading this folder")
		return SourceUpload, nil, nil
	}
	if inst == nil {
		// Reached with --source git, or in a terminal with no source asked for.
		if !ui.Interactive() || opts.DryRun {
			return "", nil, errors.New("no Git account linked in your hPanel account can read " + repo.FullName() + ". " + ConnectHint)
		}
		if want != SourceGit {
			idx, err := ui.Select("This is a Git checkout of "+repo.FullName()+", but no Git account linked in your hPanel account can read it. How do you want to deploy?",
				[]string{"Upload this folder", "Connect GitHub or GitLab in hPanel first, then deploy from Git (redeploys on every push)"}, 0)
			if err != nil {
				return "", nil, err
			}
			if idx == 0 {
				return SourceUpload, nil, nil
			}
		}
		inst, err = connectGit(ctx, a, ui, repo)
		if err != nil {
			return "", nil, err
		}
		if inst != nil {
			return SourceGit, repo, nil
		}
		if want == SourceGit {
			return "", nil, errors.New("no Git account linked in your hPanel account can read " + repo.FullName() + ". " + ConnectHint)
		}
		ui.Info("Uploading the folder instead. " + ConnectHint)
		return SourceUpload, nil, nil
	}
	if want == SourceGit {
		return SourceGit, repo, nil
	}
	idx, err := ui.Select("How do you want to deploy?", []string{gitSourceLabel(repo), "Upload this folder"}, 0)
	if err != nil {
		return "", nil, err
	}
	if idx == 0 {
		return SourceGit, repo, nil
	}
	return SourceUpload, nil, nil
}
