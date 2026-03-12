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

func TestPaperV2CommentWithOptionalParams(t *testing.T) {
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
	_, err := client.Papers.V2.Comment(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		alphaxivcat.PaperV2CommentParams{
			AnchorPosition: alphaxivcat.PaperV2CommentParamsAnchorPosition{
				Offset:    0,
				PageIndex: 0,
				SpanIndex: 0,
			},
			Body: "body",
			FocusPosition: alphaxivcat.PaperV2CommentParamsFocusPosition{
				Offset:    0,
				PageIndex: 0,
				SpanIndex: 0,
			},
			HighlightRects: []alphaxivcat.PaperV2CommentParamsHighlightRect{{
				PageIndex: 0,
				Rects: []alphaxivcat.PaperV2CommentParamsHighlightRectRect{{
					X1: 0,
					X2: 0,
					Y1: 0,
					Y2: 0,
				}},
			}},
			ParentCommentID: alphaxivcat.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			SelectedText:    alphaxivcat.String("selectedText"),
			Tag:             alphaxivcat.PaperV2CommentParamsTagAnonymous,
			Title:           alphaxivcat.String("title"),
			HighlightColor:  alphaxivcat.String("#26fBC5dF"),
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
