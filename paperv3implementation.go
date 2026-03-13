// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package alphaxivcat

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/AlphaxivCat/alphaxiv_cat-go/internal/apijson"
	"github.com/AlphaxivCat/alphaxiv_cat-go/internal/apiquery"
	"github.com/AlphaxivCat/alphaxiv_cat-go/internal/requestconfig"
	"github.com/AlphaxivCat/alphaxiv_cat-go/option"
	"github.com/AlphaxivCat/alphaxiv_cat-go/packages/param"
	"github.com/AlphaxivCat/alphaxiv_cat-go/packages/respjson"
)

// PaperV3ImplementationService contains methods and other services that help with
// interacting with the alphaxiv_cat API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPaperV3ImplementationService] method instead.
type PaperV3ImplementationService struct {
	options []option.RequestOption
}

// NewPaperV3ImplementationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewPaperV3ImplementationService(opts ...option.RequestOption) (r PaperV3ImplementationService) {
	r = PaperV3ImplementationService{}
	r.options = opts
	return
}

// Add an implementation (AlphaXiv, Marimo, Author, or Other)
//
// Source file:
// `api-server/src/controllers/papers/v3/add-implementation.controller.ts`
func (r *PaperV3ImplementationService) New(ctx context.Context, paperGroupID string, body PaperV3ImplementationNewParams, opts ...option.RequestOption) (res *PaperV3ImplementationNewResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if paperGroupID == "" {
		err = errors.New("missing required paperGroupId parameter")
		return nil, err
	}
	path := fmt.Sprintf("papers/v3/%s/implementations", paperGroupID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Get all implementations for a paper (AlphaXiv, Marimo, Author, and Other)
//
// Source file:
// `api-server/src/controllers/papers/v3/get-implementations.controller.ts`
func (r *PaperV3ImplementationService) List(ctx context.Context, paperGroupID string, opts ...option.RequestOption) (res *PaperV3ImplementationListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if paperGroupID == "" {
		err = errors.New("missing required paperGroupId parameter")
		return nil, err
	}
	path := fmt.Sprintf("papers/v3/%s/implementations", paperGroupID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Delete an implementation (AlphaXiv, Marimo, Author, or Other)
//
// Source file:
// `api-server/src/controllers/papers/v3/delete-implementation.controller.ts`
func (r *PaperV3ImplementationService) Delete(ctx context.Context, implementationID string, params PaperV3ImplementationDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if params.PaperGroupID == "" {
		err = errors.New("missing required paperGroupId parameter")
		return err
	}
	if implementationID == "" {
		err = errors.New("missing required implementationId parameter")
		return err
	}
	path := fmt.Sprintf("papers/v3/%s/implementations/%s", params.PaperGroupID, implementationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, nil, opts...)
	return err
}

type PaperV3ImplementationNewResponse struct {
	ID      string `json:"id" api:"required"`
	Message string `json:"message" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3ImplementationNewResponse) RawJSON() string { return r.JSON.raw }
func (r *PaperV3ImplementationNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3ImplementationListResponse struct {
	AlphaXivImplementations []PaperV3ImplementationListResponseAlphaXivImplementation `json:"alphaXivImplementations" api:"required"`
	PaperResources          []PaperV3ImplementationListResponsePaperResource          `json:"paperResources" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AlphaXivImplementations respjson.Field
		PaperResources          respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3ImplementationListResponse) RawJSON() string { return r.JSON.raw }
func (r *PaperV3ImplementationListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3ImplementationListResponseAlphaXivImplementation struct {
	ID string `json:"id" api:"required"`
	// Any of "github", "marimo".
	Type string `json:"type" api:"required"`
	URL  string `json:"url" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Type        respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3ImplementationListResponseAlphaXivImplementation) RawJSON() string { return r.JSON.raw }
func (r *PaperV3ImplementationListResponseAlphaXivImplementation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3ImplementationListResponsePaperResource struct {
	ID          string `json:"id" api:"required"`
	Description string `json:"description" api:"required"`
	Language    string `json:"language" api:"required"`
	// Any of "author", "other".
	Source string  `json:"source" api:"required"`
	Stars  float64 `json:"stars" api:"required"`
	// Any of "github", "marimo".
	Type string `json:"type" api:"required"`
	URL  string `json:"url" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Description respjson.Field
		Language    respjson.Field
		Source      respjson.Field
		Stars       respjson.Field
		Type        respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3ImplementationListResponsePaperResource) RawJSON() string { return r.JSON.raw }
func (r *PaperV3ImplementationListResponsePaperResource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3ImplementationNewParams struct {
	// Any of "alphaxiv", "marimo", "author".
	ImplementationType PaperV3ImplementationNewParamsImplementationType `json:"implementationType,omitzero" api:"required"`
	URL                string                                           `json:"url" api:"required" format:"uri"`
	paramObj
}

func (r PaperV3ImplementationNewParams) MarshalJSON() (data []byte, err error) {
	type shadow PaperV3ImplementationNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PaperV3ImplementationNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3ImplementationNewParamsImplementationType string

const (
	PaperV3ImplementationNewParamsImplementationTypeAlphaxiv PaperV3ImplementationNewParamsImplementationType = "alphaxiv"
	PaperV3ImplementationNewParamsImplementationTypeMarimo   PaperV3ImplementationNewParamsImplementationType = "marimo"
	PaperV3ImplementationNewParamsImplementationTypeAuthor   PaperV3ImplementationNewParamsImplementationType = "author"
)

type PaperV3ImplementationDeleteParams struct {
	PaperGroupID string `path:"paperGroupId" api:"required" format:"uuid" json:"-"`
	// Any of "alphaxiv", "resource".
	Type PaperV3ImplementationDeleteParamsType `query:"type,omitzero" api:"required" json:"-"`
	paramObj
}

// URLQuery serializes [PaperV3ImplementationDeleteParams]'s query parameters as
// `url.Values`.
func (r PaperV3ImplementationDeleteParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PaperV3ImplementationDeleteParamsType string

const (
	PaperV3ImplementationDeleteParamsTypeAlphaxiv PaperV3ImplementationDeleteParamsType = "alphaxiv"
	PaperV3ImplementationDeleteParamsTypeResource PaperV3ImplementationDeleteParamsType = "resource"
)
