package hosting

import (
	"context"
	"errors"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/site"
)

type Order = site.HostingOrder

type Datacenter struct {
	Title string `json:"title"`
	Code  string `json:"code"`
}

type Setup struct {
	Domain   string `json:"domain"`
	Username string `json:"username"`
	Status   string `json:"status"`
}

// NodeSettings are a Node.js site's build settings. Nil pointers are sent as
// null: an app run from source has no build script or output directory.
type NodeSettings struct {
	NodeVersion     int     `json:"node_version"`
	AppType         string  `json:"app_type"`
	RootDirectory   string  `json:"root_directory"`
	OutputDirectory *string `json:"output_directory"`
	BuildScript     *string `json:"build_script"`
	EntryFile       string  `json:"entry_file,omitempty"`
	PackageManager  string  `json:"package_manager,omitempty"`
}

type Build struct {
	UUID  string `json:"uuid"`
	State string `json:"state"`
}

// API is what Run needs from the Hostinger API; tests supply a fake.
type API interface {
	ListOrders(ctx context.Context) ([]Order, error)
	ListWebsites(ctx context.Context, domain string) ([]site.Website, error)
	// StartSetup creates a website; websiteType "" is an empty one.
	StartSetup(ctx context.Context, orderID int, domain, datacenter, websiteType string) (Setup, error)
	// ListDatacenters lists the best match first.
	ListDatacenters(ctx context.Context, orderID int) ([]Datacenter, error)
	ListSetups(ctx context.Context, domain string) ([]Setup, error)
	UploadCredentials(ctx context.Context, username, domain string) (Credentials, error)
	DeployArchive(ctx context.Context, username, domain, archivePath string) error
	BuildSettingsFromArchive(ctx context.Context, username, domain, archivePath string) (NodeSettings, error)
	BuildSettingsFromRepository(ctx context.Context, username, domain string, r *Repo) (NodeSettings, error)
	StartBuild(ctx context.Context, username, domain string, settings NodeSettings, archivePath string) (Build, error)
	BuildDetails(ctx context.Context, username, domain, uuid string) (Build, error)
	BuildLogs(ctx context.Context, username, domain, uuid string) (string, error)
	BuildAnalysis(ctx context.Context, username, domain, uuid string) (analysis, solution string, err error)
	ListGitInstallations(ctx context.Context) ([]Installation, error)
	ListInstallationRepositories(ctx context.Context, installationUUID string) ([]Repository, error)
	UpdateGitAutoDeploy(ctx context.Context, username, domain string, settings GitSettings) error
	BuildSettings(ctx context.Context, username, domain string) (*NodeSettings, error)
	UpdateBuildSettings(ctx context.Context, username, domain string, settings NodeSettings) error
	StartGitBuild(ctx context.Context, username, domain string, settings NodeSettings, r *Repo) (Build, error)
	OwnsDomain(ctx context.Context, domain string) (bool, error)
	HasContent(ctx context.Context, username, domain string) (bool, error)
}

// The file server can lag behind a finished setup; tests lower the delay.
var (
	CredentialRetries    = 5
	CredentialRetryDelay = 2 * time.Second
)

func NewAPI(c *api.RawClient) API {
	return &client{c: c}
}

type client struct {
	c *api.RawClient
}

func sitePath(username, domain, rest string) string {
	return "/api/hosting/v1/accounts/" + url.PathEscape(username) + "/websites/" + url.PathEscape(domain) + rest
}

func (a *client) ListOrders(ctx context.Context) ([]Order, error) {
	return site.HostingOrders(ctx, a.c)
}

func (a *client) ListWebsites(ctx context.Context, domain string) ([]site.Website, error) {
	return site.ListWebsitesWith(ctx, a.c, domain)
}

func (a *client) StartSetup(ctx context.Context, orderID int, domain, datacenter, websiteType string) (Setup, error) {
	body := map[string]any{}
	if websiteType != "" {
		body["type"] = websiteType
	}
	if domain != "" {
		body["domain"] = domain
	}
	if datacenter != "" {
		body["datacenter_code"] = datacenter
	}
	var s Setup
	err := a.c.JSON(ctx, http.MethodPost, "/api/hosting/v1/orders/"+strconv.Itoa(orderID)+"/onboardings", nil, body, &s)
	return s, err
}

func (a *client) ListDatacenters(ctx context.Context, orderID int) ([]Datacenter, error) {
	var out []Datacenter
	err := a.c.JSON(ctx, http.MethodGet, "/api/hosting/v1/datacenters", url.Values{"order_id": {strconv.Itoa(orderID)}}, nil, &out)
	return out, err
}

func (a *client) ListSetups(ctx context.Context, domain string) ([]Setup, error) {
	var setups []Setup
	err := a.c.JSON(ctx, http.MethodGet, "/api/hosting/v1/onboardings", url.Values{"domain": {domain}}, nil, &setups)
	return setups, err
}

// UploadCredentials retries briefly: for a few seconds after setup completes
// the answer can be an error or lack a URL.
func (a *client) UploadCredentials(ctx context.Context, username, domain string) (Credentials, error) {
	var creds Credentials
	var err error
	for attempt := 0; ; attempt++ {
		creds = Credentials{}
		err = a.c.JSON(ctx, http.MethodPost, "/api/hosting/v1/files/upload-urls", nil, map[string]string{"username": username, "domain": domain}, &creds)
		if err == nil && creds.URL != "" {
			return creds, nil
		}
		// A rejected token or request will not get better by waiting.
		var apiErr *api.Error
		if errors.As(err, &apiErr) && apiErr.Status < 500 && apiErr.Status != 404 && apiErr.Status != 409 {
			return creds, err
		}
		if attempt >= CredentialRetries {
			break
		}
		select {
		case <-ctx.Done():
			return creds, ctx.Err()
		case <-time.After(CredentialRetryDelay):
		}
	}
	if err == nil {
		err = errors.New("the file server gave no upload URL; run the command again in a moment")
	}
	return creds, err
}

var platformFiles = map[string]bool{".htaccess": true, "default.php": true, "index.php": false}

func (a *client) HasContent(ctx context.Context, username, domain string) (bool, error) {
	var resp struct {
		Items []struct {
			Name string `json:"name"`
		} `json:"items"`
	}
	q := url.Values{"max_depth": {"1"}, "max_items": {"20"}}
	if err := a.c.JSON(ctx, http.MethodGet, "/api/hosting/v1/accounts/"+url.PathEscape(username)+"/domains/"+url.PathEscape(domain)+"/files", q, nil, &resp); err != nil {
		var apiErr *api.Error
		if errors.As(err, &apiErr) && apiErr.Status == http.StatusNotFound {
			return false, nil // a new site has no document root to list yet
		}
		return false, err
	}
	for _, item := range resp.Items {
		if !platformFiles[item.Name] {
			return true, nil
		}
	}
	return false, nil
}

// OwnsDomain also matches a parent domain in the portfolio.
func (a *client) OwnsDomain(ctx context.Context, domain string) (bool, error) {
	var portfolio []struct {
		Domain string `json:"domain"`
	}
	if err := a.c.JSON(ctx, http.MethodGet, "/api/domains/v1/portfolio", nil, nil, &portfolio); err != nil {
		return false, err
	}
	domain = strings.ToLower(domain)
	for _, d := range portfolio {
		owned := strings.ToLower(d.Domain)
		if domain == owned || strings.HasSuffix(domain, "."+owned) {
			return true, nil
		}
	}
	return false, nil
}

func (a *client) DeployArchive(ctx context.Context, username, domain, archivePath string) error {
	return a.c.JSON(ctx, http.MethodPost, sitePath(username, domain, "/deploy"), nil, map[string]string{"archive_path": archivePath}, nil)
}

func (a *client) BuildSettingsFromArchive(ctx context.Context, username, domain, archivePath string) (NodeSettings, error) {
	var s NodeSettings
	err := a.c.JSON(ctx, http.MethodGet, sitePath(username, domain, "/nodejs/builds/settings/from-archive"), url.Values{"archive_path": {archivePath}}, nil, &s)
	return s, err
}

func (a *client) BuildSettingsFromRepository(ctx context.Context, username, domain string, r *Repo) (NodeSettings, error) {
	var s NodeSettings
	q := url.Values{"owner": {r.Owner}, "repository": {r.Name}, "branch": {r.Branch}, "installation_uuid": {r.installUID}}
	err := a.c.JSON(ctx, http.MethodGet, sitePath(username, domain, "/nodejs/builds/settings/from-repository"), q, nil, &s)
	return s, err
}

func (a *client) StartBuild(ctx context.Context, username, domain string, settings NodeSettings, archivePath string) (Build, error) {
	body := buildBody(settings)
	body["source_type"] = "archive"
	body["source_options"] = map[string]string{"archive_path": archivePath}
	var b Build
	err := a.c.JSON(ctx, http.MethodPost, sitePath(username, domain, "/nodejs/builds"), nil, body, &b)
	return b, err
}

func (a *client) BuildDetails(ctx context.Context, username, domain, uuid string) (Build, error) {
	var b Build
	err := a.c.JSON(ctx, http.MethodGet, sitePath(username, domain, "/nodejs/builds/"+url.PathEscape(uuid)), nil, nil, &b)
	return b, err
}

func (a *client) BuildLogs(ctx context.Context, username, domain, uuid string) (string, error) {
	var resp struct {
		Logs string `json:"logs"`
	}
	err := a.c.JSON(ctx, http.MethodGet, sitePath(username, domain, "/nodejs/builds/"+url.PathEscape(uuid)+"/logs"), nil, nil, &resp)
	return resp.Logs, err
}

func (a *client) BuildAnalysis(ctx context.Context, username, domain, uuid string) (string, string, error) {
	var resp struct {
		Analysis string `json:"analysis"`
		Solution string `json:"solution"`
	}
	err := a.c.JSON(ctx, http.MethodGet, sitePath(username, domain, "/nodejs/builds/"+url.PathEscape(uuid)+"/analysis"), nil, nil, &resp)
	return resp.Analysis, resp.Solution, err
}

func (a *client) ListGitInstallations(ctx context.Context) ([]Installation, error) {
	var out []Installation
	err := a.c.JSON(ctx, http.MethodGet, "/api/hosting/v1/git/installations", url.Values{"status": {"active"}}, nil, &out)
	return out, err
}

func (a *client) ListInstallationRepositories(ctx context.Context, installationUUID string) ([]Repository, error) {
	var out []Repository
	err := a.c.JSON(ctx, http.MethodGet, "/api/hosting/v1/git/installations/"+url.PathEscape(installationUUID)+"/repositories", nil, nil, &out)
	return out, err
}

func (a *client) UpdateGitAutoDeploy(ctx context.Context, username, domain string, s GitSettings) error {
	return a.c.JSON(ctx, http.MethodPut, sitePath(username, domain, "/git/auto-deployments/settings"), nil, s, nil)
}

func (a *client) BuildSettings(ctx context.Context, username, domain string) (*NodeSettings, error) {
	var s NodeSettings
	err := a.c.JSON(ctx, http.MethodGet, sitePath(username, domain, "/nodejs/builds/settings"), nil, nil, &s)
	var apiErr *api.Error
	if errors.As(err, &apiErr) && apiErr.Status == http.StatusNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &s, nil
}

func (a *client) UpdateBuildSettings(ctx context.Context, username, domain string, s NodeSettings) error {
	return a.c.JSON(ctx, http.MethodPut, sitePath(username, domain, "/nodejs/builds/settings"), nil, buildBody(s), nil)
}

func (a *client) StartGitBuild(ctx context.Context, username, domain string, s NodeSettings, r *Repo) (Build, error) {
	body := buildBody(s)
	body["source_type"] = "git"
	body["source_options"] = map[string]string{"owner": r.Owner, "repository": r.Name, "branch": r.Branch, "installation_uuid": r.installUID}
	var b Build
	err := a.c.JSON(ctx, http.MethodPost, sitePath(username, domain, "/nodejs/builds"), nil, body, &b)
	return b, err
}

func buildBody(settings NodeSettings) map[string]any {
	body := map[string]any{
		"node_version":     settings.NodeVersion,
		"app_type":         settings.AppType,
		"root_directory":   settings.RootDirectory,
		"output_directory": settings.OutputDirectory,
		"build_script":     settings.BuildScript,
	}
	if settings.EntryFile != "" {
		body["entry_file"] = settings.EntryFile
	}
	if settings.PackageManager != "" {
		body["package_manager"] = settings.PackageManager
	}
	return body
}
