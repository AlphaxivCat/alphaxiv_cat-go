// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package alphaxivcat

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/stainless-sdks/alphaxiv_cat-go/internal/apijson"
	"github.com/stainless-sdks/alphaxiv_cat-go/internal/requestconfig"
	"github.com/stainless-sdks/alphaxiv_cat-go/option"
	"github.com/stainless-sdks/alphaxiv_cat-go/packages/respjson"
)

// ArxivV1LabService contains methods and other services that help with interacting
// with the alphaxiv_cat API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewArxivV1LabService] method instead.
type ArxivV1LabService struct {
	options []option.RequestOption
}

// NewArxivV1LabService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewArxivV1LabService(opts ...option.RequestOption) (r ArxivV1LabService) {
	r = ArxivV1LabService{}
	r.options = opts
	return
}

// Gets data necessary to render arXiv labs.
//
// Source file:
// `api-server/src/controllers/arxiv/v1/get-arxiv-labs-data.controller.ts`
func (r *ArxivV1LabService) Get(ctx context.Context, unresolved string, opts ...option.RequestOption) (res *ArxivV1LabGetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if unresolved == "" {
		err = errors.New("missing required unresolved parameter")
		return nil, err
	}
	path := fmt.Sprintf("arxiv/v1/%s/labs", url.PathEscape(unresolved))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type ArxivV1LabGetResponse struct {
	Summary string `json:"summary" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Summary     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ArxivV1LabGetResponse) RawJSON() string { return r.JSON.raw }
func (r *ArxivV1LabGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
