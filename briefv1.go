// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package alphaxivcat

import (
	"context"
	"net/http"
	"slices"

	"github.com/AlphaxivCat/alphaxiv_cat-go/internal/apijson"
	"github.com/AlphaxivCat/alphaxiv_cat-go/internal/requestconfig"
	"github.com/AlphaxivCat/alphaxiv_cat-go/option"
	"github.com/AlphaxivCat/alphaxiv_cat-go/packages/param"
	"github.com/AlphaxivCat/alphaxiv_cat-go/packages/respjson"
)

// BriefV1Service contains methods and other services that help with interacting
// with the alphaxiv_cat API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBriefV1Service] method instead.
type BriefV1Service struct {
	options []option.RequestOption
	Seen    BriefV1SeenService
}

// NewBriefV1Service generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewBriefV1Service(opts ...option.RequestOption) (r BriefV1Service) {
	r = BriefV1Service{}
	r.options = opts
	r.Seen = NewBriefV1SeenService(opts...)
	return
}

// Generates speech audio from brief summary text using ElevenLabs. Returns the
// CloudFront URL. Client should try the CDN URL first and only call this if it
// 404s.
//
// Source file:
// `api-server/src/controllers/briefs/v1/generate-speech.controller.ts`
func (r *BriefV1Service) GenerateSpeech(ctx context.Context, body BriefV1GenerateSpeechParams, opts ...option.RequestOption) (res *BriefV1GenerateSpeechResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "briefs/v1/speech"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type BriefV1GenerateSpeechResponse struct {
	URL string `json:"url" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BriefV1GenerateSpeechResponse) RawJSON() string { return r.JSON.raw }
func (r *BriefV1GenerateSpeechResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BriefV1GenerateSpeechParams struct {
	// Paper group ID for caching
	PaperGroupID string `json:"paperGroupId" api:"required" format:"uuid"`
	// Text to convert to speech
	Text string `json:"text" api:"required"`
	paramObj
}

func (r BriefV1GenerateSpeechParams) MarshalJSON() (data []byte, err error) {
	type shadow BriefV1GenerateSpeechParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BriefV1GenerateSpeechParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
