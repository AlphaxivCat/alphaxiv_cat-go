// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package alphaxivcat

import (
	"context"
	"net/http"
	"slices"

	"github.com/stainless-sdks/alphaxiv_cat-go/internal/apijson"
	"github.com/stainless-sdks/alphaxiv_cat-go/internal/requestconfig"
	"github.com/stainless-sdks/alphaxiv_cat-go/option"
	"github.com/stainless-sdks/alphaxiv_cat-go/packages/respjson"
)

// AssistantService contains methods and other services that help with interacting
// with the alphaxiv_cat API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAssistantService] method instead.
type AssistantService struct {
	options []option.RequestOption
	V2      AssistantV2Service
}

// NewAssistantService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAssistantService(opts ...option.RequestOption) (r AssistantService) {
	r = AssistantService{}
	r.options = opts
	r.V2 = NewAssistantV2Service(opts...)
	return
}

// Upload a file for use with the assistant (max 30MB)
//
// Source file: `api-server/src/controllers/v1/assistant/upload-file.controller.ts`
func (r *AssistantService) UploadFile(ctx context.Context, opts ...option.RequestOption) (res *AssistantUploadFileResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/assistant/upload-file"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

type AssistantUploadFileResponse struct {
	Data AssistantUploadFileResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AssistantUploadFileResponse) RawJSON() string { return r.JSON.raw }
func (r *AssistantUploadFileResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AssistantUploadFileResponseData struct {
	ID          string  `json:"id" api:"required"`
	ContentType string  `json:"content_type" api:"required"`
	Filename    string  `json:"filename" api:"required"`
	Size        float64 `json:"size" api:"required"`
	URL         string  `json:"url" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ContentType respjson.Field
		Filename    respjson.Field
		Size        respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AssistantUploadFileResponseData) RawJSON() string { return r.JSON.raw }
func (r *AssistantUploadFileResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
