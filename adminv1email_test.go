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

func TestAdminV1EmailSendMonthlyDigestWithOptionalParams(t *testing.T) {
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
	err := client.Admin.V1.Emails.SendMonthlyDigest(context.TODO(), alphaxivcat.AdminV1EmailSendMonthlyDigestParams{
		Role: alphaxivcat.AdminV1EmailSendMonthlyDigestParamsRoleAdmin,
	})
	if err != nil {
		var apierr *alphaxivcat.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAdminV1EmailSendWeeklyDigestWithOptionalParams(t *testing.T) {
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
	err := client.Admin.V1.Emails.SendWeeklyDigest(context.TODO(), alphaxivcat.AdminV1EmailSendWeeklyDigestParams{
		Events: []alphaxivcat.AdminV1EmailSendWeeklyDigestParamsEvent{{
			Date:         "date",
			Description:  "description",
			Link:         "link",
			Title:        "title",
			EndTimeRaw:   alphaxivcat.String("endTimeRaw"),
			StartTimeRaw: alphaxivcat.String("startTimeRaw"),
		}},
		IntroText: alphaxivcat.String("introText"),
		Role:      alphaxivcat.AdminV1EmailSendWeeklyDigestParamsRoleAdmin,
		Subject:   alphaxivcat.String("subject"),
	})
	if err != nil {
		var apierr *alphaxivcat.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
