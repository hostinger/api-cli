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

func Request() client.ClientWithResponsesInterface {
	apiURL := viper.GetString("api_url")
	if apiURL == "" {
		apiURL = "https://developers.hostinger.com"
	}

	opts := []client.ClientOption{
		client.WithRequestEditorFn(addUserAgent),
	}

	if staticToken := viper.GetString("api_token"); staticToken != "" {
		// Static token wins: use it unchanged. OAuth and 401 re-auth are skipped
		// because there is no refresh source for a user-managed token.
		bearer, err := securityprovider.NewSecurityProviderBearerToken(staticToken)
		if err != nil {
			panic(err)
		}
		opts = append(opts, client.WithRequestEditorFn(bearer.Intercept))
	} else {
		// No static token: obtain the bearer via the OAuth service, and retry
		// once on a 401 after forcing a fresh token.
		oauth, err := NewOAuthService()
		if err != nil {
			log.Fatal(err)
		}
		opts = append(opts,
			client.WithRequestEditorFn(oauthBearer(oauth)),
			client.WithHTTPClient(&oauthRetryDoer{oauth: oauth, base: &http.Client{}}),
		)
	}

	c, err := client.NewClientWithResponses(apiURL, opts...)
	if err != nil {
		log.Fatal(err)
	}

	return c
}

// oauthBearer resolves a token (cached -> refresh -> interactive login) and
// attaches it as the Authorization header before each request is sent.
func oauthBearer(oauth *OAuthService) client.RequestEditorFn {
	return func(ctx context.Context, req *http.Request) error {
		token, err := oauth.Token(ctx)
		if err != nil {
			return err
		}
		req.Header.Set("Authorization", "Bearer "+token)
		return nil
	}
}

// oauthRetryDoer performs API requests and, on a 401, forces a fresh token via
// Reauthenticate and retries the request exactly once (spec §8). A second 401
// is returned to the caller unchanged — it never loops.
type oauthRetryDoer struct {
	oauth *OAuthService
	base  client.HttpRequestDoer
}

func (d *oauthRetryDoer) Do(req *http.Request) (*http.Response, error) {
	resp, err := d.base.Do(req)
	if err != nil || resp.StatusCode != http.StatusUnauthorized {
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
