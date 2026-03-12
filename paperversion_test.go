// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package alphaxivcat_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/alphaxiv_cat-go"
	"github.com/stainless-sdks/alphaxiv_cat-go/internal/testutil"
	"github.com/stainless-sdks/alphaxiv_cat-go/option"
)

func TestPaperVersionRequestAIOverviewWithOptionalParams(t *testing.T) {
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
	_, err := client.Papers.Versions.RequestAIOverview(
		context.TODO(),
		"x",
		alphaxivcat.PaperVersionRequestAIOverviewParams{
			Upid:              "x",
			PreferredLanguage: alphaxivcat.PaperVersionRequestAIOverviewParamsPreferredLanguageAm,
		},
	)
	if err != nil {
		var apierr *alphaxivcat.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPaperVersionRequestAITranslation(t *testing.T) {
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
	_, err := client.Papers.Versions.RequestAITranslation(
		context.TODO(),
		alphaxivcat.PaperVersionRequestAITranslationParamsLanguageAm,
		alphaxivcat.PaperVersionRequestAITranslationParams{
			Upid:         "x",
			VersionOrder: "x",
		},
	)
	if err != nil {
		var apierr *alphaxivcat.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
