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
	"github.com/AlphaxivCat/alphaxiv_cat-go/packages/respjson"
)

// PaperVersionService contains methods and other services that help with
// interacting with the alphaxiv_cat API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPaperVersionService] method instead.
type PaperVersionService struct {
	options []option.RequestOption
}

// NewPaperVersionService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPaperVersionService(opts ...option.RequestOption) (r PaperVersionService) {
	r = PaperVersionService{}
	r.options = opts
	return
}

// Request AI overview generation for a paper version
//
// Source file: `api-server/src/controllers/v2/papers/request-ai.controller.ts`
func (r *PaperVersionService) RequestAIOverview(ctx context.Context, versionOrder string, params PaperVersionRequestAIOverviewParams, opts ...option.RequestOption) (res *PaperVersionRequestAIOverviewResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if params.Upid == "" {
		err = errors.New("missing required upid parameter")
		return nil, err
	}
	if versionOrder == "" {
		err = errors.New("missing required versionOrder parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/papers/%s/versions/%s/request-ai", url.PathEscape(params.Upid), url.PathEscape(versionOrder))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Request AI overview translation for a paper version
//
// Source file:
// `api-server/src/controllers/v2/papers/request-ai-translation.controller.ts`
func (r *PaperVersionService) RequestAITranslation(ctx context.Context, language PaperVersionRequestAITranslationParamsLanguage, body PaperVersionRequestAITranslationParams, opts ...option.RequestOption) (res *PaperVersionRequestAITranslationResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if body.Upid == "" {
		err = errors.New("missing required upid parameter")
		return nil, err
	}
	if body.VersionOrder == "" {
		err = errors.New("missing required versionOrder parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/papers/%s/versions/%s/request-ai-translation/%v", url.PathEscape(body.Upid), url.PathEscape(body.VersionOrder), language)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

type PaperVersionRequestAIOverviewResponse struct {
	Data string `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperVersionRequestAIOverviewResponse) RawJSON() string { return r.JSON.raw }
func (r *PaperVersionRequestAIOverviewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperVersionRequestAITranslationResponse struct {
	Data string `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperVersionRequestAITranslationResponse) RawJSON() string { return r.JSON.raw }
func (r *PaperVersionRequestAITranslationResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperVersionRequestAIOverviewParams struct {
	Upid string `path:"upid" api:"required" json:"-"`
	// Any of "am", "ar", "az", "bg", "bn", "ca", "cs", "da", "de", "el", "en", "es",
	// "et", "fa", "fi", "fr", "gu", "ha", "he", "hi", "hr", "hu", "id", "it", "ja",
	// "ka", "kn", "ko", "lt", "lv", "ml", "mr", "ms", "my", "ne", "nl", "no", "pa",
	// "pl", "pt", "ro", "ru", "si", "sk", "sl", "sr", "sv", "sw", "ta", "te", "th",
	// "tl", "tr", "uk", "ur", "uz", "vi", "yo", "zh".
	PreferredLanguage PaperVersionRequestAIOverviewParamsPreferredLanguage `query:"preferredLanguage,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PaperVersionRequestAIOverviewParams]'s query parameters as
// `url.Values`.
func (r PaperVersionRequestAIOverviewParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PaperVersionRequestAIOverviewParamsPreferredLanguage string

const (
	PaperVersionRequestAIOverviewParamsPreferredLanguageAm PaperVersionRequestAIOverviewParamsPreferredLanguage = "am"
	PaperVersionRequestAIOverviewParamsPreferredLanguageAr PaperVersionRequestAIOverviewParamsPreferredLanguage = "ar"
	PaperVersionRequestAIOverviewParamsPreferredLanguageAz PaperVersionRequestAIOverviewParamsPreferredLanguage = "az"
	PaperVersionRequestAIOverviewParamsPreferredLanguageBg PaperVersionRequestAIOverviewParamsPreferredLanguage = "bg"
	PaperVersionRequestAIOverviewParamsPreferredLanguageBn PaperVersionRequestAIOverviewParamsPreferredLanguage = "bn"
	PaperVersionRequestAIOverviewParamsPreferredLanguageCa PaperVersionRequestAIOverviewParamsPreferredLanguage = "ca"
	PaperVersionRequestAIOverviewParamsPreferredLanguageCs PaperVersionRequestAIOverviewParamsPreferredLanguage = "cs"
	PaperVersionRequestAIOverviewParamsPreferredLanguageDa PaperVersionRequestAIOverviewParamsPreferredLanguage = "da"
	PaperVersionRequestAIOverviewParamsPreferredLanguageDe PaperVersionRequestAIOverviewParamsPreferredLanguage = "de"
	PaperVersionRequestAIOverviewParamsPreferredLanguageEl PaperVersionRequestAIOverviewParamsPreferredLanguage = "el"
	PaperVersionRequestAIOverviewParamsPreferredLanguageEn PaperVersionRequestAIOverviewParamsPreferredLanguage = "en"
	PaperVersionRequestAIOverviewParamsPreferredLanguageEs PaperVersionRequestAIOverviewParamsPreferredLanguage = "es"
	PaperVersionRequestAIOverviewParamsPreferredLanguageEt PaperVersionRequestAIOverviewParamsPreferredLanguage = "et"
	PaperVersionRequestAIOverviewParamsPreferredLanguageFa PaperVersionRequestAIOverviewParamsPreferredLanguage = "fa"
	PaperVersionRequestAIOverviewParamsPreferredLanguageFi PaperVersionRequestAIOverviewParamsPreferredLanguage = "fi"
	PaperVersionRequestAIOverviewParamsPreferredLanguageFr PaperVersionRequestAIOverviewParamsPreferredLanguage = "fr"
	PaperVersionRequestAIOverviewParamsPreferredLanguageGu PaperVersionRequestAIOverviewParamsPreferredLanguage = "gu"
	PaperVersionRequestAIOverviewParamsPreferredLanguageHa PaperVersionRequestAIOverviewParamsPreferredLanguage = "ha"
	PaperVersionRequestAIOverviewParamsPreferredLanguageHe PaperVersionRequestAIOverviewParamsPreferredLanguage = "he"
	PaperVersionRequestAIOverviewParamsPreferredLanguageHi PaperVersionRequestAIOverviewParamsPreferredLanguage = "hi"
	PaperVersionRequestAIOverviewParamsPreferredLanguageHr PaperVersionRequestAIOverviewParamsPreferredLanguage = "hr"
	PaperVersionRequestAIOverviewParamsPreferredLanguageHu PaperVersionRequestAIOverviewParamsPreferredLanguage = "hu"
	PaperVersionRequestAIOverviewParamsPreferredLanguageID PaperVersionRequestAIOverviewParamsPreferredLanguage = "id"
	PaperVersionRequestAIOverviewParamsPreferredLanguageIt PaperVersionRequestAIOverviewParamsPreferredLanguage = "it"
	PaperVersionRequestAIOverviewParamsPreferredLanguageJa PaperVersionRequestAIOverviewParamsPreferredLanguage = "ja"
	PaperVersionRequestAIOverviewParamsPreferredLanguageKa PaperVersionRequestAIOverviewParamsPreferredLanguage = "ka"
	PaperVersionRequestAIOverviewParamsPreferredLanguageKn PaperVersionRequestAIOverviewParamsPreferredLanguage = "kn"
	PaperVersionRequestAIOverviewParamsPreferredLanguageKo PaperVersionRequestAIOverviewParamsPreferredLanguage = "ko"
	PaperVersionRequestAIOverviewParamsPreferredLanguageLt PaperVersionRequestAIOverviewParamsPreferredLanguage = "lt"
	PaperVersionRequestAIOverviewParamsPreferredLanguageLv PaperVersionRequestAIOverviewParamsPreferredLanguage = "lv"
	PaperVersionRequestAIOverviewParamsPreferredLanguageMl PaperVersionRequestAIOverviewParamsPreferredLanguage = "ml"
	PaperVersionRequestAIOverviewParamsPreferredLanguageMr PaperVersionRequestAIOverviewParamsPreferredLanguage = "mr"
	PaperVersionRequestAIOverviewParamsPreferredLanguageMs PaperVersionRequestAIOverviewParamsPreferredLanguage = "ms"
	PaperVersionRequestAIOverviewParamsPreferredLanguageMy PaperVersionRequestAIOverviewParamsPreferredLanguage = "my"
	PaperVersionRequestAIOverviewParamsPreferredLanguageNe PaperVersionRequestAIOverviewParamsPreferredLanguage = "ne"
	PaperVersionRequestAIOverviewParamsPreferredLanguageNl PaperVersionRequestAIOverviewParamsPreferredLanguage = "nl"
	PaperVersionRequestAIOverviewParamsPreferredLanguageNo PaperVersionRequestAIOverviewParamsPreferredLanguage = "no"
	PaperVersionRequestAIOverviewParamsPreferredLanguagePa PaperVersionRequestAIOverviewParamsPreferredLanguage = "pa"
	PaperVersionRequestAIOverviewParamsPreferredLanguagePl PaperVersionRequestAIOverviewParamsPreferredLanguage = "pl"
	PaperVersionRequestAIOverviewParamsPreferredLanguagePt PaperVersionRequestAIOverviewParamsPreferredLanguage = "pt"
	PaperVersionRequestAIOverviewParamsPreferredLanguageRo PaperVersionRequestAIOverviewParamsPreferredLanguage = "ro"
	PaperVersionRequestAIOverviewParamsPreferredLanguageRu PaperVersionRequestAIOverviewParamsPreferredLanguage = "ru"
	PaperVersionRequestAIOverviewParamsPreferredLanguageSi PaperVersionRequestAIOverviewParamsPreferredLanguage = "si"
	PaperVersionRequestAIOverviewParamsPreferredLanguageSk PaperVersionRequestAIOverviewParamsPreferredLanguage = "sk"
	PaperVersionRequestAIOverviewParamsPreferredLanguageSl PaperVersionRequestAIOverviewParamsPreferredLanguage = "sl"
	PaperVersionRequestAIOverviewParamsPreferredLanguageSr PaperVersionRequestAIOverviewParamsPreferredLanguage = "sr"
	PaperVersionRequestAIOverviewParamsPreferredLanguageSv PaperVersionRequestAIOverviewParamsPreferredLanguage = "sv"
	PaperVersionRequestAIOverviewParamsPreferredLanguageSw PaperVersionRequestAIOverviewParamsPreferredLanguage = "sw"
	PaperVersionRequestAIOverviewParamsPreferredLanguageTa PaperVersionRequestAIOverviewParamsPreferredLanguage = "ta"
	PaperVersionRequestAIOverviewParamsPreferredLanguageTe PaperVersionRequestAIOverviewParamsPreferredLanguage = "te"
	PaperVersionRequestAIOverviewParamsPreferredLanguageTh PaperVersionRequestAIOverviewParamsPreferredLanguage = "th"
	PaperVersionRequestAIOverviewParamsPreferredLanguageTl PaperVersionRequestAIOverviewParamsPreferredLanguage = "tl"
	PaperVersionRequestAIOverviewParamsPreferredLanguageTr PaperVersionRequestAIOverviewParamsPreferredLanguage = "tr"
	PaperVersionRequestAIOverviewParamsPreferredLanguageUk PaperVersionRequestAIOverviewParamsPreferredLanguage = "uk"
	PaperVersionRequestAIOverviewParamsPreferredLanguageUr PaperVersionRequestAIOverviewParamsPreferredLanguage = "ur"
	PaperVersionRequestAIOverviewParamsPreferredLanguageUz PaperVersionRequestAIOverviewParamsPreferredLanguage = "uz"
	PaperVersionRequestAIOverviewParamsPreferredLanguageVi PaperVersionRequestAIOverviewParamsPreferredLanguage = "vi"
	PaperVersionRequestAIOverviewParamsPreferredLanguageYo PaperVersionRequestAIOverviewParamsPreferredLanguage = "yo"
	PaperVersionRequestAIOverviewParamsPreferredLanguageZh PaperVersionRequestAIOverviewParamsPreferredLanguage = "zh"
)

type PaperVersionRequestAITranslationParams struct {
	Upid         string `path:"upid" api:"required" json:"-"`
	VersionOrder string `path:"versionOrder" api:"required" json:"-"`
	paramObj
}

type PaperVersionRequestAITranslationParamsLanguage string

const (
	PaperVersionRequestAITranslationParamsLanguageAm PaperVersionRequestAITranslationParamsLanguage = "am"
	PaperVersionRequestAITranslationParamsLanguageAr PaperVersionRequestAITranslationParamsLanguage = "ar"
	PaperVersionRequestAITranslationParamsLanguageAz PaperVersionRequestAITranslationParamsLanguage = "az"
	PaperVersionRequestAITranslationParamsLanguageBg PaperVersionRequestAITranslationParamsLanguage = "bg"
	PaperVersionRequestAITranslationParamsLanguageBn PaperVersionRequestAITranslationParamsLanguage = "bn"
	PaperVersionRequestAITranslationParamsLanguageCa PaperVersionRequestAITranslationParamsLanguage = "ca"
	PaperVersionRequestAITranslationParamsLanguageCs PaperVersionRequestAITranslationParamsLanguage = "cs"
	PaperVersionRequestAITranslationParamsLanguageDa PaperVersionRequestAITranslationParamsLanguage = "da"
	PaperVersionRequestAITranslationParamsLanguageDe PaperVersionRequestAITranslationParamsLanguage = "de"
	PaperVersionRequestAITranslationParamsLanguageEl PaperVersionRequestAITranslationParamsLanguage = "el"
	PaperVersionRequestAITranslationParamsLanguageEs PaperVersionRequestAITranslationParamsLanguage = "es"
	PaperVersionRequestAITranslationParamsLanguageEt PaperVersionRequestAITranslationParamsLanguage = "et"
	PaperVersionRequestAITranslationParamsLanguageFa PaperVersionRequestAITranslationParamsLanguage = "fa"
	PaperVersionRequestAITranslationParamsLanguageFi PaperVersionRequestAITranslationParamsLanguage = "fi"
	PaperVersionRequestAITranslationParamsLanguageFr PaperVersionRequestAITranslationParamsLanguage = "fr"
	PaperVersionRequestAITranslationParamsLanguageGu PaperVersionRequestAITranslationParamsLanguage = "gu"
	PaperVersionRequestAITranslationParamsLanguageHa PaperVersionRequestAITranslationParamsLanguage = "ha"
	PaperVersionRequestAITranslationParamsLanguageHe PaperVersionRequestAITranslationParamsLanguage = "he"
	PaperVersionRequestAITranslationParamsLanguageHi PaperVersionRequestAITranslationParamsLanguage = "hi"
	PaperVersionRequestAITranslationParamsLanguageHr PaperVersionRequestAITranslationParamsLanguage = "hr"
	PaperVersionRequestAITranslationParamsLanguageHu PaperVersionRequestAITranslationParamsLanguage = "hu"
	PaperVersionRequestAITranslationParamsLanguageID PaperVersionRequestAITranslationParamsLanguage = "id"
	PaperVersionRequestAITranslationParamsLanguageIt PaperVersionRequestAITranslationParamsLanguage = "it"
	PaperVersionRequestAITranslationParamsLanguageJa PaperVersionRequestAITranslationParamsLanguage = "ja"
	PaperVersionRequestAITranslationParamsLanguageKa PaperVersionRequestAITranslationParamsLanguage = "ka"
	PaperVersionRequestAITranslationParamsLanguageKn PaperVersionRequestAITranslationParamsLanguage = "kn"
	PaperVersionRequestAITranslationParamsLanguageKo PaperVersionRequestAITranslationParamsLanguage = "ko"
	PaperVersionRequestAITranslationParamsLanguageLt PaperVersionRequestAITranslationParamsLanguage = "lt"
	PaperVersionRequestAITranslationParamsLanguageLv PaperVersionRequestAITranslationParamsLanguage = "lv"
	PaperVersionRequestAITranslationParamsLanguageMl PaperVersionRequestAITranslationParamsLanguage = "ml"
	PaperVersionRequestAITranslationParamsLanguageMr PaperVersionRequestAITranslationParamsLanguage = "mr"
	PaperVersionRequestAITranslationParamsLanguageMs PaperVersionRequestAITranslationParamsLanguage = "ms"
	PaperVersionRequestAITranslationParamsLanguageMy PaperVersionRequestAITranslationParamsLanguage = "my"
	PaperVersionRequestAITranslationParamsLanguageNe PaperVersionRequestAITranslationParamsLanguage = "ne"
	PaperVersionRequestAITranslationParamsLanguageNl PaperVersionRequestAITranslationParamsLanguage = "nl"
	PaperVersionRequestAITranslationParamsLanguageNo PaperVersionRequestAITranslationParamsLanguage = "no"
	PaperVersionRequestAITranslationParamsLanguagePa PaperVersionRequestAITranslationParamsLanguage = "pa"
	PaperVersionRequestAITranslationParamsLanguagePl PaperVersionRequestAITranslationParamsLanguage = "pl"
	PaperVersionRequestAITranslationParamsLanguagePt PaperVersionRequestAITranslationParamsLanguage = "pt"
	PaperVersionRequestAITranslationParamsLanguageRo PaperVersionRequestAITranslationParamsLanguage = "ro"
	PaperVersionRequestAITranslationParamsLanguageRu PaperVersionRequestAITranslationParamsLanguage = "ru"
	PaperVersionRequestAITranslationParamsLanguageSi PaperVersionRequestAITranslationParamsLanguage = "si"
	PaperVersionRequestAITranslationParamsLanguageSk PaperVersionRequestAITranslationParamsLanguage = "sk"
	PaperVersionRequestAITranslationParamsLanguageSl PaperVersionRequestAITranslationParamsLanguage = "sl"
	PaperVersionRequestAITranslationParamsLanguageSr PaperVersionRequestAITranslationParamsLanguage = "sr"
	PaperVersionRequestAITranslationParamsLanguageSv PaperVersionRequestAITranslationParamsLanguage = "sv"
	PaperVersionRequestAITranslationParamsLanguageSw PaperVersionRequestAITranslationParamsLanguage = "sw"
	PaperVersionRequestAITranslationParamsLanguageTa PaperVersionRequestAITranslationParamsLanguage = "ta"
	PaperVersionRequestAITranslationParamsLanguageTe PaperVersionRequestAITranslationParamsLanguage = "te"
	PaperVersionRequestAITranslationParamsLanguageTh PaperVersionRequestAITranslationParamsLanguage = "th"
	PaperVersionRequestAITranslationParamsLanguageTl PaperVersionRequestAITranslationParamsLanguage = "tl"
	PaperVersionRequestAITranslationParamsLanguageTr PaperVersionRequestAITranslationParamsLanguage = "tr"
	PaperVersionRequestAITranslationParamsLanguageUk PaperVersionRequestAITranslationParamsLanguage = "uk"
	PaperVersionRequestAITranslationParamsLanguageUr PaperVersionRequestAITranslationParamsLanguage = "ur"
	PaperVersionRequestAITranslationParamsLanguageUz PaperVersionRequestAITranslationParamsLanguage = "uz"
	PaperVersionRequestAITranslationParamsLanguageVi PaperVersionRequestAITranslationParamsLanguage = "vi"
	PaperVersionRequestAITranslationParamsLanguageYo PaperVersionRequestAITranslationParamsLanguage = "yo"
	PaperVersionRequestAITranslationParamsLanguageZh PaperVersionRequestAITranslationParamsLanguage = "zh"
)
