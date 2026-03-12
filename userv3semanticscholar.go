// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package alphaxivcat

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/AlphaxivCat/alphaxiv_cat-go/internal/apijson"
	"github.com/AlphaxivCat/alphaxiv_cat-go/internal/requestconfig"
	"github.com/AlphaxivCat/alphaxiv_cat-go/option"
	"github.com/AlphaxivCat/alphaxiv_cat-go/packages/respjson"
)

// UserV3SemanticScholarService contains methods and other services that help with
// interacting with the alphaxiv_cat API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserV3SemanticScholarService] method instead.
type UserV3SemanticScholarService struct {
	options []option.RequestOption
}

// NewUserV3SemanticScholarService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewUserV3SemanticScholarService(opts ...option.RequestOption) (r UserV3SemanticScholarService) {
	r = UserV3SemanticScholarService{}
	r.options = opts
	return
}

// Link a user's account to a Semantic Scholar profile based on claimed papers
//
// Source file:
// `api-server/src/controllers/users/v3/link-semantic-scholar.controller.ts`
func (r *UserV3SemanticScholarService) Link(ctx context.Context, id string, opts ...option.RequestOption) (res *UserV3SemanticScholarLinkResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("users/v3/%s/semantic-scholar/link", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Refresh Semantic Scholar data for a linked user
//
// Source file:
// `api-server/src/controllers/users/v3/scrape-semantic-scholar.controller.ts`
func (r *UserV3SemanticScholarService) Scrape(ctx context.Context, id string, opts ...option.RequestOption) (res *UserV3SemanticScholarScrapeResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("users/v3/%s/semantic-scholar/scrape", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

type UserV3SemanticScholarLinkResponse struct {
	Message string `json:"message" api:"required"`
	Data    any    `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserV3SemanticScholarLinkResponse) RawJSON() string { return r.JSON.raw }
func (r *UserV3SemanticScholarLinkResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserV3SemanticScholarScrapeResponse struct {
	Message string `json:"message" api:"required"`
	Data    any    `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserV3SemanticScholarScrapeResponse) RawJSON() string { return r.JSON.raw }
func (r *UserV3SemanticScholarScrapeResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
