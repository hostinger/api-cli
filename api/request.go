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
	bearerToken, err := securityprovider.NewSecurityProviderBearerToken(viper.GetString("api_token"))
	if err != nil {
		panic(err)
	}

	var apiUrl = viper.GetString("api_url")
	if apiUrl == "" {
		apiUrl = "https://developers.hostinger.com"
	}

	c, err := client.NewClientWithResponses(
		apiUrl,
		client.WithRequestEditorFn(bearerToken.Intercept),
		client.WithRequestEditorFn(addUserAgent),
	)
	if err != nil {
		log.Fatal(err)
	}

	return c
}

func addUserAgent(ctx context.Context, req *http.Request) error {
	req.Header.Set("User-Agent", fmt.Sprintf("hostinger-cli/%s", utils.CLIVersion.String(false)))
	return nil
}
