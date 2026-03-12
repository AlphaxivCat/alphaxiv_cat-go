// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package alphaxivcat_test

import (
	"context"
	"os"
	"testing"

	"github.com/stainless-sdks/alphaxiv_cat-go"
	"github.com/stainless-sdks/alphaxiv_cat-go/internal/testutil"
	"github.com/stainless-sdks/alphaxiv_cat-go/option"
)

func TestUsage(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := alphaxivcat.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	response, err := client.Papers.Versions.RequestAITranslation(
		context.TODO(),
		alphaxivcat.PaperVersionRequestAITranslationParamsLanguageAm,
		alphaxivcat.PaperVersionRequestAITranslationParams{
			Upid:         "REPLACE_ME",
			VersionOrder: "REPLACE_ME",
		},
	)
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	t.Logf("%+v\n", response.Data)
}
