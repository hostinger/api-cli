package api

import (
	"context"
	"fmt"
	"github.com/hostinger/api-cli/client"
	"github.com/hostinger/api-cli/utils"
	"github.com/oapi-codegen/oapi-codegen/v2/pkg/securityprovider"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

// Request returns the API client for a command run. Without a static token it
// signs in through OAuth, opening the browser when needed.
func Request() client.ClientWithResponsesInterface {
	return newClient(true)
}

// CompletionRequest returns an API client for shell completion: it uses the
// static token or an already stored OAuth sign-in, and never starts an
// interactive login.
func CompletionRequest() client.ClientWithResponsesInterface {
	return newClient(false)
}

func newClient(interactive bool) client.ClientWithResponsesInterface {
	apiURL, opts := clientOptions(interactive, interactive && viper.GetBool("dry_run"))
	c, err := client.NewClientWithResponses(apiURL, opts...)
	if err != nil {
		log.Fatal(err)
	}
	return c
}

// clientOptions are shared by the typed client and Raw. dryRun prints each
// request instead of sending it.
func clientOptions(interactive, dryRun bool) (string, []client.ClientOption) {
	apiURL := viper.GetString("api_url")
	if apiURL == "" {
		apiURL = "https://developers.hostinger.com"
	}

	opts := []client.ClientOption{
		client.WithRequestEditorFn(addUserAgent),
	}

	var doer client.HttpRequestDoer = &http.Client{}
	if viper.GetBool("verbose") {
		doer = &verboseDoer{base: doer}
	}

	if dryRun {
		// Print the request instead of sending it; the request itself needs no
		// credentials (lookups through Raw may, see Raw).
		opts = append(opts, client.WithHTTPClient(&dryRunDoer{}))
	} else if staticToken := viper.GetString("api_token"); staticToken != "" {
		// Static token wins: use it unchanged. OAuth and 401 re-auth are skipped
		// because there is no refresh source for a user-managed token.
		bearer, err := securityprovider.NewSecurityProviderBearerToken(staticToken)
		if err != nil {
			panic(err)
		}
		opts = append(opts, client.WithRequestEditorFn(bearer.Intercept), client.WithHTTPClient(doer))
	} else {
		// No static token: obtain the bearer via the OAuth service, and retry
		// once on a 401 after forcing a fresh token.
		oauth, err := NewOAuthService()
		if err != nil {
			log.Fatal(err)
		}
		opts = append(opts,
			client.WithRequestEditorFn(oauthBearer(oauth, interactive)),
			client.WithHTTPClient(&oauthRetryDoer{oauth: oauth, base: doer, interactive: interactive}),
		)
	}

	return apiURL, opts
}

// oauthBearer resolves a token (cached -> refresh -> interactive login) and
// attaches it as the Authorization header before each request is sent. When
// interactive is false the login step is skipped.
func oauthBearer(oauth *OAuthService, interactive bool) client.RequestEditorFn {
	return func(ctx context.Context, req *http.Request) error {
		resolve := oauth.Token
		if !interactive {
			resolve = oauth.TokenNoLogin
		}
		token, err := resolve(ctx)
		if err != nil {
			return err
		}
		req.Header.Set("Authorization", "Bearer "+token)
		return nil
	}
}

// oauthRetryDoer performs API requests and, on a 401, forces a fresh token via
// Reauthenticate and retries the request exactly once (spec §8). A second 401
// is returned to the caller unchanged — it never loops. Non-interactive
// clients return the first 401 as is.
type oauthRetryDoer struct {
	oauth       *OAuthService
	base        client.HttpRequestDoer
	interactive bool
}

func (d *oauthRetryDoer) Do(req *http.Request) (*http.Response, error) {
	resp, err := d.base.Do(req)
	if err != nil || resp.StatusCode != http.StatusUnauthorized || !d.interactive {
		return resp, err
	}

	// Rebuild the request with a rewound body. If the body can't be replayed,
	// return the original 401 rather than sending a corrupt retry.
	retry := req.Clone(req.Context())
	if req.Body != nil && req.Body != http.NoBody {
		if req.GetBody == nil {
			return resp, nil
		}
		body, gerr := req.GetBody()
		if gerr != nil {
			return resp, nil
		}
		retry.Body = body
	}

	token, rerr := d.oauth.Reauthenticate(req.Context())
	if rerr != nil {
		resp.Body.Close()
		return nil, rerr
	}

	resp.Body.Close() // discard the 401 body; we are retrying
	retry.Header.Set("Authorization", "Bearer "+token)
	return d.base.Do(retry)
}

func addUserAgent(ctx context.Context, req *http.Request) error {
	req.Header.Set("User-Agent", fmt.Sprintf("hostinger-cli/%s", utils.CLIVersion.String(false)))
	return nil
}
