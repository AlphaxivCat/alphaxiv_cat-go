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
	"github.com/stainless-sdks/alphaxiv_cat-go/packages/param"
	"github.com/stainless-sdks/alphaxiv_cat-go/packages/respjson"
)

// PaperPrivateService contains methods and other services that help with
// interacting with the alphaxiv_cat API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPaperPrivateService] method instead.
type PaperPrivateService struct {
	options []option.RequestOption
}

// NewPaperPrivateService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPaperPrivateService(opts ...option.RequestOption) (r PaperPrivateService) {
	r = PaperPrivateService{}
	r.options = opts
	return
}

// Upload a private PDF paper
//
// Source file:
// `api-server/src/controllers/v2/papers/upload-private-paper.controller.ts`
func (r *PaperPrivateService) New(ctx context.Context, body PaperPrivateNewParams, opts ...option.RequestOption) (res *PaperPrivateNewResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v2/papers/private"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Update metadata for a private paper
//
// Source file:
// `api-server/src/controllers/v2/papers/update-private-paper.controller.ts`
func (r *PaperPrivateService) UpdateMetadata(ctx context.Context, paperID string, body PaperPrivateUpdateMetadataParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if paperID == "" {
		err = errors.New("missing required paperId parameter")
		return err
	}
	path := fmt.Sprintf("v2/papers/private/%s/metadata", url.PathEscape(paperID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, nil, opts...)
	return err
}

type PaperPrivateNewResponse struct {
	Data PaperPrivateNewResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperPrivateNewResponse) RawJSON() string { return r.JSON.raw }
func (r *PaperPrivateNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperPrivateNewResponseData struct {
	ID          string  `json:"id" api:"required"`
	ContentType string  `json:"content_type" api:"required"`
	Filename    string  `json:"filename" api:"required"`
	Size        float64 `json:"size" api:"required"`
	URL         string  `json:"url" api:"required"`
	PaperData   any     `json:"paper_data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ContentType respjson.Field
		Filename    respjson.Field
		Size        respjson.Field
		URL         respjson.Field
		PaperData   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperPrivateNewResponseData) RawJSON() string { return r.JSON.raw }
func (r *PaperPrivateNewResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperPrivateNewParams struct {
	ContentType string `json:"contentType" api:"required"`
	File        string `json:"file" api:"required"`
	Filename    string `json:"filename" api:"required"`
	paramObj
}

func (r PaperPrivateNewParams) MarshalJSON() (data []byte, err error) {
	type shadow PaperPrivateNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PaperPrivateNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperPrivateUpdateMetadataParams struct {
	Bibtex          param.Opt[string]                        `json:"bibtex,omitzero" api:"required"`
	Abstract        string                                   `json:"abstract" api:"required"`
	Authors         []PaperPrivateUpdateMetadataParamsAuthor `json:"authors,omitzero" api:"required"`
	Categories      []string                                 `json:"categories,omitzero" api:"required"`
	PublicationDate float64                                  `json:"publicationDate" api:"required"`
	Title           string                                   `json:"title" api:"required"`
	paramObj
}

func (r PaperPrivateUpdateMetadataParams) MarshalJSON() (data []byte, err error) {
	type shadow PaperPrivateUpdateMetadataParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PaperPrivateUpdateMetadataParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Name is required.
type PaperPrivateUpdateMetadataParamsAuthor struct {
	Name string `json:"name" api:"required"`
	paramObj
}

func (r PaperPrivateUpdateMetadataParamsAuthor) MarshalJSON() (data []byte, err error) {
	type shadow PaperPrivateUpdateMetadataParamsAuthor
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PaperPrivateUpdateMetadataParamsAuthor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
