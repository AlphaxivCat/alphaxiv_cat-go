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

func TestCommentEdit(t *testing.T) {
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
	err := client.Comments.Edit(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		alphaxivcat.CommentEditParams{
			AnchorPosition: alphaxivcat.CommentEditParamsAnchorPosition{
				Offset:    0,
				PageIndex: 0,
				SpanIndex: 0,
			},
			Body: "body",
			FocusPosition: alphaxivcat.CommentEditParamsFocusPosition{
				Offset:    0,
				PageIndex: 0,
				SpanIndex: 0,
			},
			HighlightColor: alphaxivcat.String("#26fBC5dF"),
			HighlightRects: []alphaxivcat.CommentEditParamsHighlightRect{{
				PageIndex: 0,
				Rects: []alphaxivcat.CommentEditParamsHighlightRectRect{{
					X1: 0,
					X2: 0,
					Y1: 0,
					Y2: 0,
				}},
			}},
			SelectedText: alphaxivcat.String("selectedText"),
			Title:        alphaxivcat.String("title"),
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
