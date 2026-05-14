// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package alphaxivcat_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/AlphaxivCat/alphaxiv_cat-go"
	"github.com/AlphaxivCat/alphaxiv_cat-go/internal/testutil"
	"github.com/AlphaxivCat/alphaxiv_cat-go/option"
)

func TestPaperMetadataGetLatestMetadataWithOptionalParams(t *testing.T) {
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
	_, err := client.Papers.Metadata.GetLatestMetadata(
		context.TODO(),
		"upid",
		alphaxivcat.PaperMetadataGetLatestMetadataParams{
			PreventTracking: alphaxivcat.String("preventTracking"),
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

func TestPaperMetadataGetVersionMetadataWithOptionalParams(t *testing.T) {
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
	_, err := client.Papers.Metadata.GetVersionMetadata(
		context.TODO(),
		"469",
		alphaxivcat.PaperMetadataGetVersionMetadataParams{
			Upid:            "upid",
			PreventTracking: alphaxivcat.String("preventTracking"),
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
