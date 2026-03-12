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

// PaperV3OverviewService contains methods and other services that help with
// interacting with the alphaxiv_cat API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPaperV3OverviewService] method instead.
type PaperV3OverviewService struct {
	options []option.RequestOption
}

// NewPaperV3OverviewService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPaperV3OverviewService(opts ...option.RequestOption) (r PaperV3OverviewService) {
	r = PaperV3OverviewService{}
	r.options = opts
	return
}

// Retrieve paper version overview for given language. Does not create it if it
// doesn't exist
//
// Source file: `api-server/src/controllers/papers/v3/get-overviews.controller.ts`
func (r *PaperV3OverviewService) Get(ctx context.Context, language PaperV3OverviewGetParamsLanguage, query PaperV3OverviewGetParams, opts ...option.RequestOption) (res *PaperV3OverviewGetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if query.PaperVersion == "" {
		err = errors.New("missing required paperVersion parameter")
		return nil, err
	}
	path := fmt.Sprintf("papers/v3/%s/overview/%v", query.PaperVersion, language)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Retrieve paper version status for overview generation and translations
//
// Source file:
// `api-server/src/controllers/papers/v3/get-overviews-status.controller.ts`
func (r *PaperV3OverviewService) GetStatus(ctx context.Context, paperVersion string, opts ...option.RequestOption) (res *PaperV3OverviewGetStatusResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if paperVersion == "" {
		err = errors.New("missing required paperVersion parameter")
		return nil, err
	}
	path := fmt.Sprintf("papers/v3/%s/overview/status", paperVersion)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type PaperV3OverviewGetResponse struct {
	Abstract              string                                          `json:"abstract" api:"required"`
	Citations             []PaperV3OverviewGetResponseCitation            `json:"citations" api:"required"`
	IntermediateReport    string                                          `json:"intermediateReport" api:"required"`
	Overview              string                                          `json:"overview" api:"required"`
	Summary               PaperV3OverviewGetResponseSummary               `json:"summary" api:"required"`
	Title                 string                                          `json:"title" api:"required"`
	AITooltips            PaperV3OverviewGetResponseAITooltips            `json:"aiTooltips"`
	OverviewSectionTitles PaperV3OverviewGetResponseOverviewSectionTitles `json:"overviewSectionTitles"`
	SummarySectionTitles  PaperV3OverviewGetResponseSummarySectionTitles  `json:"summarySectionTitles"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Abstract              respjson.Field
		Citations             respjson.Field
		IntermediateReport    respjson.Field
		Overview              respjson.Field
		Summary               respjson.Field
		Title                 respjson.Field
		AITooltips            respjson.Field
		OverviewSectionTitles respjson.Field
		SummarySectionTitles  respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3OverviewGetResponse) RawJSON() string { return r.JSON.raw }
func (r *PaperV3OverviewGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3OverviewGetResponseCitation struct {
	AlphaxivLink  string `json:"alphaxivLink" api:"required"`
	FullCitation  string `json:"fullCitation" api:"required"`
	Justification string `json:"justification" api:"required"`
	Title         string `json:"title" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AlphaxivLink  respjson.Field
		FullCitation  respjson.Field
		Justification respjson.Field
		Title         respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3OverviewGetResponseCitation) RawJSON() string { return r.JSON.raw }
func (r *PaperV3OverviewGetResponseCitation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3OverviewGetResponseSummary struct {
	KeyInsights     []string `json:"keyInsights" api:"required"`
	OriginalProblem []string `json:"originalProblem" api:"required"`
	Results         []string `json:"results" api:"required"`
	Solution        []string `json:"solution" api:"required"`
	Summary         string   `json:"summary" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		KeyInsights     respjson.Field
		OriginalProblem respjson.Field
		Results         respjson.Field
		Solution        respjson.Field
		Summary         respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3OverviewGetResponseSummary) RawJSON() string { return r.JSON.raw }
func (r *PaperV3OverviewGetResponseSummary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3OverviewGetResponseAITooltips struct {
	AIGeneratedContent string `json:"aiGeneratedContent" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AIGeneratedContent respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3OverviewGetResponseAITooltips) RawJSON() string { return r.JSON.raw }
func (r *PaperV3OverviewGetResponseAITooltips) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3OverviewGetResponseOverviewSectionTitles struct {
	RelevantCitations string `json:"relevantCitations" api:"required"`
	TableOfContents   string `json:"tableOfContents" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		RelevantCitations respjson.Field
		TableOfContents   respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3OverviewGetResponseOverviewSectionTitles) RawJSON() string { return r.JSON.raw }
func (r *PaperV3OverviewGetResponseOverviewSectionTitles) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3OverviewGetResponseSummarySectionTitles struct {
	Abstract  string `json:"abstract" api:"required"`
	Method    string `json:"method" api:"required"`
	Problem   string `json:"problem" api:"required"`
	Results   string `json:"results" api:"required"`
	Summary   string `json:"summary" api:"required"`
	Takeaways string `json:"takeaways" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Abstract    respjson.Field
		Method      respjson.Field
		Problem     respjson.Field
		Results     respjson.Field
		Summary     respjson.Field
		Takeaways   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3OverviewGetResponseSummarySectionTitles) RawJSON() string { return r.JSON.raw }
func (r *PaperV3OverviewGetResponseSummarySectionTitles) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3OverviewGetStatusResponse struct {
	// Any of "queued", "extracting", "generating", "done", "errored".
	State        PaperV3OverviewGetStatusResponseState        `json:"state" api:"required"`
	Translations PaperV3OverviewGetStatusResponseTranslations `json:"translations" api:"required"`
	// Milliseconds since UNIX epoch
	UpdatedAt float64 `json:"updatedAt" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		State        respjson.Field
		Translations respjson.Field
		UpdatedAt    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3OverviewGetStatusResponse) RawJSON() string { return r.JSON.raw }
func (r *PaperV3OverviewGetStatusResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3OverviewGetStatusResponseState string

const (
	PaperV3OverviewGetStatusResponseStateQueued     PaperV3OverviewGetStatusResponseState = "queued"
	PaperV3OverviewGetStatusResponseStateExtracting PaperV3OverviewGetStatusResponseState = "extracting"
	PaperV3OverviewGetStatusResponseStateGenerating PaperV3OverviewGetStatusResponseState = "generating"
	PaperV3OverviewGetStatusResponseStateDone       PaperV3OverviewGetStatusResponseState = "done"
	PaperV3OverviewGetStatusResponseStateErrored    PaperV3OverviewGetStatusResponseState = "errored"
)

type PaperV3OverviewGetStatusResponseTranslations struct {
	ID PaperV3OverviewGetStatusResponseTranslationsID `json:"id"`
	Am PaperV3OverviewGetStatusResponseTranslationsAm `json:"am"`
	Ar PaperV3OverviewGetStatusResponseTranslationsAr `json:"ar"`
	Az PaperV3OverviewGetStatusResponseTranslationsAz `json:"az"`
	Bg PaperV3OverviewGetStatusResponseTranslationsBg `json:"bg"`
	Bn PaperV3OverviewGetStatusResponseTranslationsBn `json:"bn"`
	Ca PaperV3OverviewGetStatusResponseTranslationsCa `json:"ca"`
	Cs PaperV3OverviewGetStatusResponseTranslationsCs `json:"cs"`
	Da PaperV3OverviewGetStatusResponseTranslationsDa `json:"da"`
	De PaperV3OverviewGetStatusResponseTranslationsDe `json:"de"`
	El PaperV3OverviewGetStatusResponseTranslationsEl `json:"el"`
	En PaperV3OverviewGetStatusResponseTranslationsEn `json:"en"`
	Es PaperV3OverviewGetStatusResponseTranslationsEs `json:"es"`
	Et PaperV3OverviewGetStatusResponseTranslationsEt `json:"et"`
	Fa PaperV3OverviewGetStatusResponseTranslationsFa `json:"fa"`
	Fi PaperV3OverviewGetStatusResponseTranslationsFi `json:"fi"`
	Fr PaperV3OverviewGetStatusResponseTranslationsFr `json:"fr"`
	Gu PaperV3OverviewGetStatusResponseTranslationsGu `json:"gu"`
	Ha PaperV3OverviewGetStatusResponseTranslationsHa `json:"ha"`
	He PaperV3OverviewGetStatusResponseTranslationsHe `json:"he"`
	Hi PaperV3OverviewGetStatusResponseTranslationsHi `json:"hi"`
	Hr PaperV3OverviewGetStatusResponseTranslationsHr `json:"hr"`
	Hu PaperV3OverviewGetStatusResponseTranslationsHu `json:"hu"`
	It PaperV3OverviewGetStatusResponseTranslationsIt `json:"it"`
	Ja PaperV3OverviewGetStatusResponseTranslationsJa `json:"ja"`
	Ka PaperV3OverviewGetStatusResponseTranslationsKa `json:"ka"`
	Kn PaperV3OverviewGetStatusResponseTranslationsKn `json:"kn"`
	Ko PaperV3OverviewGetStatusResponseTranslationsKo `json:"ko"`
	Lt PaperV3OverviewGetStatusResponseTranslationsLt `json:"lt"`
	Lv PaperV3OverviewGetStatusResponseTranslationsLv `json:"lv"`
	Ml PaperV3OverviewGetStatusResponseTranslationsMl `json:"ml"`
	Mr PaperV3OverviewGetStatusResponseTranslationsMr `json:"mr"`
	Ms PaperV3OverviewGetStatusResponseTranslationsMs `json:"ms"`
	My PaperV3OverviewGetStatusResponseTranslationsMy `json:"my"`
	Ne PaperV3OverviewGetStatusResponseTranslationsNe `json:"ne"`
	Nl PaperV3OverviewGetStatusResponseTranslationsNl `json:"nl"`
	No PaperV3OverviewGetStatusResponseTranslationsNo `json:"no"`
	Pa PaperV3OverviewGetStatusResponseTranslationsPa `json:"pa"`
	Pl PaperV3OverviewGetStatusResponseTranslationsPl `json:"pl"`
	Pt PaperV3OverviewGetStatusResponseTranslationsPt `json:"pt"`
	Ro PaperV3OverviewGetStatusResponseTranslationsRo `json:"ro"`
	Ru PaperV3OverviewGetStatusResponseTranslationsRu `json:"ru"`
	Si PaperV3OverviewGetStatusResponseTranslationsSi `json:"si"`
	Sk PaperV3OverviewGetStatusResponseTranslationsSk `json:"sk"`
	Sl PaperV3OverviewGetStatusResponseTranslationsSl `json:"sl"`
	Sr PaperV3OverviewGetStatusResponseTranslationsSr `json:"sr"`
	Sv PaperV3OverviewGetStatusResponseTranslationsSv `json:"sv"`
	Sw PaperV3OverviewGetStatusResponseTranslationsSw `json:"sw"`
	Ta PaperV3OverviewGetStatusResponseTranslationsTa `json:"ta"`
	Te PaperV3OverviewGetStatusResponseTranslationsTe `json:"te"`
	Th PaperV3OverviewGetStatusResponseTranslationsTh `json:"th"`
	Tl PaperV3OverviewGetStatusResponseTranslationsTl `json:"tl"`
	Tr PaperV3OverviewGetStatusResponseTranslationsTr `json:"tr"`
	Uk PaperV3OverviewGetStatusResponseTranslationsUk `json:"uk"`
	Ur PaperV3OverviewGetStatusResponseTranslationsUr `json:"ur"`
	Uz PaperV3OverviewGetStatusResponseTranslationsUz `json:"uz"`
	Vi PaperV3OverviewGetStatusResponseTranslationsVi `json:"vi"`
	Yo PaperV3OverviewGetStatusResponseTranslationsYo `json:"yo"`
	Zh PaperV3OverviewGetStatusResponseTranslationsZh `json:"zh"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Am          respjson.Field
		Ar          respjson.Field
		Az          respjson.Field
		Bg          respjson.Field
		Bn          respjson.Field
		Ca          respjson.Field
		Cs          respjson.Field
		Da          respjson.Field
		De          respjson.Field
		El          respjson.Field
		En          respjson.Field
		Es          respjson.Field
		Et          respjson.Field
		Fa          respjson.Field
		Fi          respjson.Field
		Fr          respjson.Field
		Gu          respjson.Field
		Ha          respjson.Field
		He          respjson.Field
		Hi          respjson.Field
		Hr          respjson.Field
		Hu          respjson.Field
		It          respjson.Field
		Ja          respjson.Field
		Ka          respjson.Field
		Kn          respjson.Field
		Ko          respjson.Field
		Lt          respjson.Field
		Lv          respjson.Field
		Ml          respjson.Field
		Mr          respjson.Field
		Ms          respjson.Field
		My          respjson.Field
		Ne          respjson.Field
		Nl          respjson.Field
		No          respjson.Field
		Pa          respjson.Field
		Pl          respjson.Field
		Pt          respjson.Field
		Ro          respjson.Field
		Ru          respjson.Field
		Si          respjson.Field
		Sk          respjson.Field
		Sl          respjson.Field
		Sr          respjson.Field
		Sv          respjson.Field
		Sw          respjson.Field
		Ta          respjson.Field
		Te          respjson.Field
		Th          respjson.Field
		Tl          respjson.Field
		Tr          respjson.Field
		Uk          respjson.Field
		Ur          respjson.Field
		Uz          respjson.Field
		Vi          respjson.Field
		Yo          respjson.Field
		Zh          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3OverviewGetStatusResponseTranslations) RawJSON() string { return r.JSON.raw }
func (r *PaperV3OverviewGetStatusResponseTranslations) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3OverviewGetStatusResponseTranslationsID struct {
	Error       string  `json:"error" api:"required"`
	RequestedAt float64 `json:"requestedAt" api:"required"`
	// Any of "queued", "done", "errored".
	State     string  `json:"state" api:"required"`
	UpdatedAt float64 `json:"updatedAt" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Error       respjson.Field
		RequestedAt respjson.Field
		State       respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3OverviewGetStatusResponseTranslationsID) RawJSON() string { return r.JSON.raw }
func (r *PaperV3OverviewGetStatusResponseTranslationsID) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3OverviewGetStatusResponseTranslationsAm struct {
	Error       string  `json:"error" api:"required"`
	RequestedAt float64 `json:"requestedAt" api:"required"`
	// Any of "queued", "done", "errored".
	State     string  `json:"state" api:"required"`
	UpdatedAt float64 `json:"updatedAt" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Error       respjson.Field
		RequestedAt respjson.Field
		State       respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3OverviewGetStatusResponseTranslationsAm) RawJSON() string { return r.JSON.raw }
func (r *PaperV3OverviewGetStatusResponseTranslationsAm) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3OverviewGetStatusResponseTranslationsAr struct {
	Error       string  `json:"error" api:"required"`
	RequestedAt float64 `json:"requestedAt" api:"required"`
	// Any of "queued", "done", "errored".
	State     string  `json:"state" api:"required"`
	UpdatedAt float64 `json:"updatedAt" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Error       respjson.Field
		RequestedAt respjson.Field
		State       respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3OverviewGetStatusResponseTranslationsAr) RawJSON() string { return r.JSON.raw }
func (r *PaperV3OverviewGetStatusResponseTranslationsAr) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3OverviewGetStatusResponseTranslationsAz struct {
	Error       string  `json:"error" api:"required"`
	RequestedAt float64 `json:"requestedAt" api:"required"`
	// Any of "queued", "done", "errored".
	State     string  `json:"state" api:"required"`
	UpdatedAt float64 `json:"updatedAt" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Error       respjson.Field
		RequestedAt respjson.Field
		State       respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3OverviewGetStatusResponseTranslationsAz) RawJSON() string { return r.JSON.raw }
func (r *PaperV3OverviewGetStatusResponseTranslationsAz) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3OverviewGetStatusResponseTranslationsBg struct {
	Error       string  `json:"error" api:"required"`
	RequestedAt float64 `json:"requestedAt" api:"required"`
	// Any of "queued", "done", "errored".
	State     string  `json:"state" api:"required"`
	UpdatedAt float64 `json:"updatedAt" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Error       respjson.Field
		RequestedAt respjson.Field
		State       respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3OverviewGetStatusResponseTranslationsBg) RawJSON() string { return r.JSON.raw }
func (r *PaperV3OverviewGetStatusResponseTranslationsBg) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3OverviewGetStatusResponseTranslationsBn struct {
	Error       string  `json:"error" api:"required"`
	RequestedAt float64 `json:"requestedAt" api:"required"`
	// Any of "queued", "done", "errored".
	State     string  `json:"state" api:"required"`
	UpdatedAt float64 `json:"updatedAt" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Error       respjson.Field
		RequestedAt respjson.Field
		State       respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3OverviewGetStatusResponseTranslationsBn) RawJSON() string { return r.JSON.raw }
func (r *PaperV3OverviewGetStatusResponseTranslationsBn) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3OverviewGetStatusResponseTranslationsCa struct {
	Error       string  `json:"error" api:"required"`
	RequestedAt float64 `json:"requestedAt" api:"required"`
	// Any of "queued", "done", "errored".
	State     string  `json:"state" api:"required"`
	UpdatedAt float64 `json:"updatedAt" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Error       respjson.Field
		RequestedAt respjson.Field
		State       respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3OverviewGetStatusResponseTranslationsCa) RawJSON() string { return r.JSON.raw }
func (r *PaperV3OverviewGetStatusResponseTranslationsCa) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3OverviewGetStatusResponseTranslationsCs struct {
	Error       string  `json:"error" api:"required"`
	RequestedAt float64 `json:"requestedAt" api:"required"`
	// Any of "queued", "done", "errored".
	State     string  `json:"state" api:"required"`
	UpdatedAt float64 `json:"updatedAt" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Error       respjson.Field
		RequestedAt respjson.Field
		State       respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3OverviewGetStatusResponseTranslationsCs) RawJSON() string { return r.JSON.raw }
func (r *PaperV3OverviewGetStatusResponseTranslationsCs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3OverviewGetStatusResponseTranslationsDa struct {
	Error       string  `json:"error" api:"required"`
	RequestedAt float64 `json:"requestedAt" api:"required"`
	// Any of "queued", "done", "errored".
	State     string  `json:"state" api:"required"`
	UpdatedAt float64 `json:"updatedAt" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Error       respjson.Field
		RequestedAt respjson.Field
		State       respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3OverviewGetStatusResponseTranslationsDa) RawJSON() string { return r.JSON.raw }
func (r *PaperV3OverviewGetStatusResponseTranslationsDa) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3OverviewGetStatusResponseTranslationsDe struct {
	Error       string  `json:"error" api:"required"`
	RequestedAt float64 `json:"requestedAt" api:"required"`
	// Any of "queued", "done", "errored".
	State     string  `json:"state" api:"required"`
	UpdatedAt float64 `json:"updatedAt" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Error       respjson.Field
		RequestedAt respjson.Field
		State       respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3OverviewGetStatusResponseTranslationsDe) RawJSON() string { return r.JSON.raw }
func (r *PaperV3OverviewGetStatusResponseTranslationsDe) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3OverviewGetStatusResponseTranslationsEl struct {
	Error       string  `json:"error" api:"required"`
	RequestedAt float64 `json:"requestedAt" api:"required"`
	// Any of "queued", "done", "errored".
	State     string  `json:"state" api:"required"`
	UpdatedAt float64 `json:"updatedAt" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Error       respjson.Field
		RequestedAt respjson.Field
		State       respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3OverviewGetStatusResponseTranslationsEl) RawJSON() string { return r.JSON.raw }
func (r *PaperV3OverviewGetStatusResponseTranslationsEl) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3OverviewGetStatusResponseTranslationsEn struct {
	Error       string  `json:"error" api:"required"`
	RequestedAt float64 `json:"requestedAt" api:"required"`
	// Any of "queued", "done", "errored".
	State     string  `json:"state" api:"required"`
	UpdatedAt float64 `json:"updatedAt" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Error       respjson.Field
		RequestedAt respjson.Field
		State       respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3OverviewGetStatusResponseTranslationsEn) RawJSON() string { return r.JSON.raw }
func (r *PaperV3OverviewGetStatusResponseTranslationsEn) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3OverviewGetStatusResponseTranslationsEs struct {
	Error       string  `json:"error" api:"required"`
	RequestedAt float64 `json:"requestedAt" api:"required"`
	// Any of "queued", "done", "errored".
	State     string  `json:"state" api:"required"`
	UpdatedAt float64 `json:"updatedAt" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Error       respjson.Field
		RequestedAt respjson.Field
		State       respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3OverviewGetStatusResponseTranslationsEs) RawJSON() string { return r.JSON.raw }
func (r *PaperV3OverviewGetStatusResponseTranslationsEs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3OverviewGetStatusResponseTranslationsEt struct {
	Error       string  `json:"error" api:"required"`
	RequestedAt float64 `json:"requestedAt" api:"required"`
	// Any of "queued", "done", "errored".
	State     string  `json:"state" api:"required"`
	UpdatedAt float64 `json:"updatedAt" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Error       respjson.Field
		RequestedAt respjson.Field
		State       respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3OverviewGetStatusResponseTranslationsEt) RawJSON() string { return r.JSON.raw }
func (r *PaperV3OverviewGetStatusResponseTranslationsEt) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3OverviewGetStatusResponseTranslationsFa struct {
	Error       string  `json:"error" api:"required"`
	RequestedAt float64 `json:"requestedAt" api:"required"`
	// Any of "queued", "done", "errored".
	State     string  `json:"state" api:"required"`
	UpdatedAt float64 `json:"updatedAt" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Error       respjson.Field
		RequestedAt respjson.Field
		State       respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3OverviewGetStatusResponseTranslationsFa) RawJSON() string { return r.JSON.raw }
func (r *PaperV3OverviewGetStatusResponseTranslationsFa) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3OverviewGetStatusResponseTranslationsFi struct {
	Error       string  `json:"error" api:"required"`
	RequestedAt float64 `json:"requestedAt" api:"required"`
	// Any of "queued", "done", "errored".
	State     string  `json:"state" api:"required"`
	UpdatedAt float64 `json:"updatedAt" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Error       respjson.Field
		RequestedAt respjson.Field
		State       respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3OverviewGetStatusResponseTranslationsFi) RawJSON() string { return r.JSON.raw }
func (r *PaperV3OverviewGetStatusResponseTranslationsFi) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3OverviewGetStatusResponseTranslationsFr struct {
	Error       string  `json:"error" api:"required"`
	RequestedAt float64 `json:"requestedAt" api:"required"`
	// Any of "queued", "done", "errored".
	State     string  `json:"state" api:"required"`
	UpdatedAt float64 `json:"updatedAt" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Error       respjson.Field
		RequestedAt respjson.Field
		State       respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3OverviewGetStatusResponseTranslationsFr) RawJSON() string { return r.JSON.raw }
func (r *PaperV3OverviewGetStatusResponseTranslationsFr) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3OverviewGetStatusResponseTranslationsGu struct {
	Error       string  `json:"error" api:"required"`
	RequestedAt float64 `json:"requestedAt" api:"required"`
	// Any of "queued", "done", "errored".
	State     string  `json:"state" api:"required"`
	UpdatedAt float64 `json:"updatedAt" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Error       respjson.Field
		RequestedAt respjson.Field
		State       respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3OverviewGetStatusResponseTranslationsGu) RawJSON() string { return r.JSON.raw }
func (r *PaperV3OverviewGetStatusResponseTranslationsGu) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3OverviewGetStatusResponseTranslationsHa struct {
	Error       string  `json:"error" api:"required"`
	RequestedAt float64 `json:"requestedAt" api:"required"`
	// Any of "queued", "done", "errored".
	State     string  `json:"state" api:"required"`
	UpdatedAt float64 `json:"updatedAt" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Error       respjson.Field
		RequestedAt respjson.Field
		State       respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3OverviewGetStatusResponseTranslationsHa) RawJSON() string { return r.JSON.raw }
func (r *PaperV3OverviewGetStatusResponseTranslationsHa) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3OverviewGetStatusResponseTranslationsHe struct {
	Error       string  `json:"error" api:"required"`
	RequestedAt float64 `json:"requestedAt" api:"required"`
	// Any of "queued", "done", "errored".
	State     string  `json:"state" api:"required"`
	UpdatedAt float64 `json:"updatedAt" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Error       respjson.Field
		RequestedAt respjson.Field
		State       respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3OverviewGetStatusResponseTranslationsHe) RawJSON() string { return r.JSON.raw }
func (r *PaperV3OverviewGetStatusResponseTranslationsHe) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3OverviewGetStatusResponseTranslationsHi struct {
	Error       string  `json:"error" api:"required"`
	RequestedAt float64 `json:"requestedAt" api:"required"`
	// Any of "queued", "done", "errored".
	State     string  `json:"state" api:"required"`
	UpdatedAt float64 `json:"updatedAt" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Error       respjson.Field
		RequestedAt respjson.Field
		State       respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3OverviewGetStatusResponseTranslationsHi) RawJSON() string { return r.JSON.raw }
func (r *PaperV3OverviewGetStatusResponseTranslationsHi) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3OverviewGetStatusResponseTranslationsHr struct {
	Error       string  `json:"error" api:"required"`
	RequestedAt float64 `json:"requestedAt" api:"required"`
	// Any of "queued", "done", "errored".
	State     string  `json:"state" api:"required"`
	UpdatedAt float64 `json:"updatedAt" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Error       respjson.Field
		RequestedAt respjson.Field
		State       respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3OverviewGetStatusResponseTranslationsHr) RawJSON() string { return r.JSON.raw }
func (r *PaperV3OverviewGetStatusResponseTranslationsHr) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3OverviewGetStatusResponseTranslationsHu struct {
	Error       string  `json:"error" api:"required"`
	RequestedAt float64 `json:"requestedAt" api:"required"`
	// Any of "queued", "done", "errored".
	State     string  `json:"state" api:"required"`
	UpdatedAt float64 `json:"updatedAt" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Error       respjson.Field
		RequestedAt respjson.Field
		State       respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3OverviewGetStatusResponseTranslationsHu) RawJSON() string { return r.JSON.raw }
func (r *PaperV3OverviewGetStatusResponseTranslationsHu) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3OverviewGetStatusResponseTranslationsIt struct {
	Error       string  `json:"error" api:"required"`
	RequestedAt float64 `json:"requestedAt" api:"required"`
	// Any of "queued", "done", "errored".
	State     string  `json:"state" api:"required"`
	UpdatedAt float64 `json:"updatedAt" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Error       respjson.Field
		RequestedAt respjson.Field
		State       respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3OverviewGetStatusResponseTranslationsIt) RawJSON() string { return r.JSON.raw }
func (r *PaperV3OverviewGetStatusResponseTranslationsIt) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3OverviewGetStatusResponseTranslationsJa struct {
	Error       string  `json:"error" api:"required"`
	RequestedAt float64 `json:"requestedAt" api:"required"`
	// Any of "queued", "done", "errored".
	State     string  `json:"state" api:"required"`
	UpdatedAt float64 `json:"updatedAt" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Error       respjson.Field
		RequestedAt respjson.Field
		State       respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3OverviewGetStatusResponseTranslationsJa) RawJSON() string { return r.JSON.raw }
func (r *PaperV3OverviewGetStatusResponseTranslationsJa) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3OverviewGetStatusResponseTranslationsKa struct {
	Error       string  `json:"error" api:"required"`
	RequestedAt float64 `json:"requestedAt" api:"required"`
	// Any of "queued", "done", "errored".
	State     string  `json:"state" api:"required"`
	UpdatedAt float64 `json:"updatedAt" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Error       respjson.Field
		RequestedAt respjson.Field
		State       respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3OverviewGetStatusResponseTranslationsKa) RawJSON() string { return r.JSON.raw }
func (r *PaperV3OverviewGetStatusResponseTranslationsKa) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3OverviewGetStatusResponseTranslationsKn struct {
	Error       string  `json:"error" api:"required"`
	RequestedAt float64 `json:"requestedAt" api:"required"`
	// Any of "queued", "done", "errored".
	State     string  `json:"state" api:"required"`
	UpdatedAt float64 `json:"updatedAt" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Error       respjson.Field
		RequestedAt respjson.Field
		State       respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3OverviewGetStatusResponseTranslationsKn) RawJSON() string { return r.JSON.raw }
func (r *PaperV3OverviewGetStatusResponseTranslationsKn) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3OverviewGetStatusResponseTranslationsKo struct {
	Error       string  `json:"error" api:"required"`
	RequestedAt float64 `json:"requestedAt" api:"required"`
	// Any of "queued", "done", "errored".
	State     string  `json:"state" api:"required"`
	UpdatedAt float64 `json:"updatedAt" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Error       respjson.Field
		RequestedAt respjson.Field
		State       respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3OverviewGetStatusResponseTranslationsKo) RawJSON() string { return r.JSON.raw }
func (r *PaperV3OverviewGetStatusResponseTranslationsKo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3OverviewGetStatusResponseTranslationsLt struct {
	Error       string  `json:"error" api:"required"`
	RequestedAt float64 `json:"requestedAt" api:"required"`
	// Any of "queued", "done", "errored".
	State     string  `json:"state" api:"required"`
	UpdatedAt float64 `json:"updatedAt" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Error       respjson.Field
		RequestedAt respjson.Field
		State       respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3OverviewGetStatusResponseTranslationsLt) RawJSON() string { return r.JSON.raw }
func (r *PaperV3OverviewGetStatusResponseTranslationsLt) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3OverviewGetStatusResponseTranslationsLv struct {
	Error       string  `json:"error" api:"required"`
	RequestedAt float64 `json:"requestedAt" api:"required"`
	// Any of "queued", "done", "errored".
	State     string  `json:"state" api:"required"`
	UpdatedAt float64 `json:"updatedAt" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Error       respjson.Field
		RequestedAt respjson.Field
		State       respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3OverviewGetStatusResponseTranslationsLv) RawJSON() string { return r.JSON.raw }
func (r *PaperV3OverviewGetStatusResponseTranslationsLv) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3OverviewGetStatusResponseTranslationsMl struct {
	Error       string  `json:"error" api:"required"`
	RequestedAt float64 `json:"requestedAt" api:"required"`
	// Any of "queued", "done", "errored".
	State     string  `json:"state" api:"required"`
	UpdatedAt float64 `json:"updatedAt" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Error       respjson.Field
		RequestedAt respjson.Field
		State       respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3OverviewGetStatusResponseTranslationsMl) RawJSON() string { return r.JSON.raw }
func (r *PaperV3OverviewGetStatusResponseTranslationsMl) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3OverviewGetStatusResponseTranslationsMr struct {
	Error       string  `json:"error" api:"required"`
	RequestedAt float64 `json:"requestedAt" api:"required"`
	// Any of "queued", "done", "errored".
	State     string  `json:"state" api:"required"`
	UpdatedAt float64 `json:"updatedAt" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Error       respjson.Field
		RequestedAt respjson.Field
		State       respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3OverviewGetStatusResponseTranslationsMr) RawJSON() string { return r.JSON.raw }
func (r *PaperV3OverviewGetStatusResponseTranslationsMr) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3OverviewGetStatusResponseTranslationsMs struct {
	Error       string  `json:"error" api:"required"`
	RequestedAt float64 `json:"requestedAt" api:"required"`
	// Any of "queued", "done", "errored".
	State     string  `json:"state" api:"required"`
	UpdatedAt float64 `json:"updatedAt" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Error       respjson.Field
		RequestedAt respjson.Field
		State       respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3OverviewGetStatusResponseTranslationsMs) RawJSON() string { return r.JSON.raw }
func (r *PaperV3OverviewGetStatusResponseTranslationsMs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3OverviewGetStatusResponseTranslationsMy struct {
	Error       string  `json:"error" api:"required"`
	RequestedAt float64 `json:"requestedAt" api:"required"`
	// Any of "queued", "done", "errored".
	State     string  `json:"state" api:"required"`
	UpdatedAt float64 `json:"updatedAt" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Error       respjson.Field
		RequestedAt respjson.Field
		State       respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3OverviewGetStatusResponseTranslationsMy) RawJSON() string { return r.JSON.raw }
func (r *PaperV3OverviewGetStatusResponseTranslationsMy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3OverviewGetStatusResponseTranslationsNe struct {
	Error       string  `json:"error" api:"required"`
	RequestedAt float64 `json:"requestedAt" api:"required"`
	// Any of "queued", "done", "errored".
	State     string  `json:"state" api:"required"`
	UpdatedAt float64 `json:"updatedAt" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Error       respjson.Field
		RequestedAt respjson.Field
		State       respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3OverviewGetStatusResponseTranslationsNe) RawJSON() string { return r.JSON.raw }
func (r *PaperV3OverviewGetStatusResponseTranslationsNe) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3OverviewGetStatusResponseTranslationsNl struct {
	Error       string  `json:"error" api:"required"`
	RequestedAt float64 `json:"requestedAt" api:"required"`
	// Any of "queued", "done", "errored".
	State     string  `json:"state" api:"required"`
	UpdatedAt float64 `json:"updatedAt" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Error       respjson.Field
		RequestedAt respjson.Field
		State       respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3OverviewGetStatusResponseTranslationsNl) RawJSON() string { return r.JSON.raw }
func (r *PaperV3OverviewGetStatusResponseTranslationsNl) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3OverviewGetStatusResponseTranslationsNo struct {
	Error       string  `json:"error" api:"required"`
	RequestedAt float64 `json:"requestedAt" api:"required"`
	// Any of "queued", "done", "errored".
	State     string  `json:"state" api:"required"`
	UpdatedAt float64 `json:"updatedAt" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Error       respjson.Field
		RequestedAt respjson.Field
		State       respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3OverviewGetStatusResponseTranslationsNo) RawJSON() string { return r.JSON.raw }
func (r *PaperV3OverviewGetStatusResponseTranslationsNo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3OverviewGetStatusResponseTranslationsPa struct {
	Error       string  `json:"error" api:"required"`
	RequestedAt float64 `json:"requestedAt" api:"required"`
	// Any of "queued", "done", "errored".
	State     string  `json:"state" api:"required"`
	UpdatedAt float64 `json:"updatedAt" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Error       respjson.Field
		RequestedAt respjson.Field
		State       respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3OverviewGetStatusResponseTranslationsPa) RawJSON() string { return r.JSON.raw }
func (r *PaperV3OverviewGetStatusResponseTranslationsPa) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3OverviewGetStatusResponseTranslationsPl struct {
	Error       string  `json:"error" api:"required"`
	RequestedAt float64 `json:"requestedAt" api:"required"`
	// Any of "queued", "done", "errored".
	State     string  `json:"state" api:"required"`
	UpdatedAt float64 `json:"updatedAt" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Error       respjson.Field
		RequestedAt respjson.Field
		State       respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3OverviewGetStatusResponseTranslationsPl) RawJSON() string { return r.JSON.raw }
func (r *PaperV3OverviewGetStatusResponseTranslationsPl) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3OverviewGetStatusResponseTranslationsPt struct {
	Error       string  `json:"error" api:"required"`
	RequestedAt float64 `json:"requestedAt" api:"required"`
	// Any of "queued", "done", "errored".
	State     string  `json:"state" api:"required"`
	UpdatedAt float64 `json:"updatedAt" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Error       respjson.Field
		RequestedAt respjson.Field
		State       respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3OverviewGetStatusResponseTranslationsPt) RawJSON() string { return r.JSON.raw }
func (r *PaperV3OverviewGetStatusResponseTranslationsPt) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3OverviewGetStatusResponseTranslationsRo struct {
	Error       string  `json:"error" api:"required"`
	RequestedAt float64 `json:"requestedAt" api:"required"`
	// Any of "queued", "done", "errored".
	State     string  `json:"state" api:"required"`
	UpdatedAt float64 `json:"updatedAt" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Error       respjson.Field
		RequestedAt respjson.Field
		State       respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3OverviewGetStatusResponseTranslationsRo) RawJSON() string { return r.JSON.raw }
func (r *PaperV3OverviewGetStatusResponseTranslationsRo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3OverviewGetStatusResponseTranslationsRu struct {
	Error       string  `json:"error" api:"required"`
	RequestedAt float64 `json:"requestedAt" api:"required"`
	// Any of "queued", "done", "errored".
	State     string  `json:"state" api:"required"`
	UpdatedAt float64 `json:"updatedAt" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Error       respjson.Field
		RequestedAt respjson.Field
		State       respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3OverviewGetStatusResponseTranslationsRu) RawJSON() string { return r.JSON.raw }
func (r *PaperV3OverviewGetStatusResponseTranslationsRu) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3OverviewGetStatusResponseTranslationsSi struct {
	Error       string  `json:"error" api:"required"`
	RequestedAt float64 `json:"requestedAt" api:"required"`
	// Any of "queued", "done", "errored".
	State     string  `json:"state" api:"required"`
	UpdatedAt float64 `json:"updatedAt" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Error       respjson.Field
		RequestedAt respjson.Field
		State       respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3OverviewGetStatusResponseTranslationsSi) RawJSON() string { return r.JSON.raw }
func (r *PaperV3OverviewGetStatusResponseTranslationsSi) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3OverviewGetStatusResponseTranslationsSk struct {
	Error       string  `json:"error" api:"required"`
	RequestedAt float64 `json:"requestedAt" api:"required"`
	// Any of "queued", "done", "errored".
	State     string  `json:"state" api:"required"`
	UpdatedAt float64 `json:"updatedAt" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Error       respjson.Field
		RequestedAt respjson.Field
		State       respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3OverviewGetStatusResponseTranslationsSk) RawJSON() string { return r.JSON.raw }
func (r *PaperV3OverviewGetStatusResponseTranslationsSk) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3OverviewGetStatusResponseTranslationsSl struct {
	Error       string  `json:"error" api:"required"`
	RequestedAt float64 `json:"requestedAt" api:"required"`
	// Any of "queued", "done", "errored".
	State     string  `json:"state" api:"required"`
	UpdatedAt float64 `json:"updatedAt" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Error       respjson.Field
		RequestedAt respjson.Field
		State       respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3OverviewGetStatusResponseTranslationsSl) RawJSON() string { return r.JSON.raw }
func (r *PaperV3OverviewGetStatusResponseTranslationsSl) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3OverviewGetStatusResponseTranslationsSr struct {
	Error       string  `json:"error" api:"required"`
	RequestedAt float64 `json:"requestedAt" api:"required"`
	// Any of "queued", "done", "errored".
	State     string  `json:"state" api:"required"`
	UpdatedAt float64 `json:"updatedAt" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Error       respjson.Field
		RequestedAt respjson.Field
		State       respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3OverviewGetStatusResponseTranslationsSr) RawJSON() string { return r.JSON.raw }
func (r *PaperV3OverviewGetStatusResponseTranslationsSr) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3OverviewGetStatusResponseTranslationsSv struct {
	Error       string  `json:"error" api:"required"`
	RequestedAt float64 `json:"requestedAt" api:"required"`
	// Any of "queued", "done", "errored".
	State     string  `json:"state" api:"required"`
	UpdatedAt float64 `json:"updatedAt" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Error       respjson.Field
		RequestedAt respjson.Field
		State       respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3OverviewGetStatusResponseTranslationsSv) RawJSON() string { return r.JSON.raw }
func (r *PaperV3OverviewGetStatusResponseTranslationsSv) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3OverviewGetStatusResponseTranslationsSw struct {
	Error       string  `json:"error" api:"required"`
	RequestedAt float64 `json:"requestedAt" api:"required"`
	// Any of "queued", "done", "errored".
	State     string  `json:"state" api:"required"`
	UpdatedAt float64 `json:"updatedAt" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Error       respjson.Field
		RequestedAt respjson.Field
		State       respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3OverviewGetStatusResponseTranslationsSw) RawJSON() string { return r.JSON.raw }
func (r *PaperV3OverviewGetStatusResponseTranslationsSw) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3OverviewGetStatusResponseTranslationsTa struct {
	Error       string  `json:"error" api:"required"`
	RequestedAt float64 `json:"requestedAt" api:"required"`
	// Any of "queued", "done", "errored".
	State     string  `json:"state" api:"required"`
	UpdatedAt float64 `json:"updatedAt" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Error       respjson.Field
		RequestedAt respjson.Field
		State       respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3OverviewGetStatusResponseTranslationsTa) RawJSON() string { return r.JSON.raw }
func (r *PaperV3OverviewGetStatusResponseTranslationsTa) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3OverviewGetStatusResponseTranslationsTe struct {
	Error       string  `json:"error" api:"required"`
	RequestedAt float64 `json:"requestedAt" api:"required"`
	// Any of "queued", "done", "errored".
	State     string  `json:"state" api:"required"`
	UpdatedAt float64 `json:"updatedAt" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Error       respjson.Field
		RequestedAt respjson.Field
		State       respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3OverviewGetStatusResponseTranslationsTe) RawJSON() string { return r.JSON.raw }
func (r *PaperV3OverviewGetStatusResponseTranslationsTe) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3OverviewGetStatusResponseTranslationsTh struct {
	Error       string  `json:"error" api:"required"`
	RequestedAt float64 `json:"requestedAt" api:"required"`
	// Any of "queued", "done", "errored".
	State     string  `json:"state" api:"required"`
	UpdatedAt float64 `json:"updatedAt" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Error       respjson.Field
		RequestedAt respjson.Field
		State       respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3OverviewGetStatusResponseTranslationsTh) RawJSON() string { return r.JSON.raw }
func (r *PaperV3OverviewGetStatusResponseTranslationsTh) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3OverviewGetStatusResponseTranslationsTl struct {
	Error       string  `json:"error" api:"required"`
	RequestedAt float64 `json:"requestedAt" api:"required"`
	// Any of "queued", "done", "errored".
	State     string  `json:"state" api:"required"`
	UpdatedAt float64 `json:"updatedAt" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Error       respjson.Field
		RequestedAt respjson.Field
		State       respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3OverviewGetStatusResponseTranslationsTl) RawJSON() string { return r.JSON.raw }
func (r *PaperV3OverviewGetStatusResponseTranslationsTl) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3OverviewGetStatusResponseTranslationsTr struct {
	Error       string  `json:"error" api:"required"`
	RequestedAt float64 `json:"requestedAt" api:"required"`
	// Any of "queued", "done", "errored".
	State     string  `json:"state" api:"required"`
	UpdatedAt float64 `json:"updatedAt" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Error       respjson.Field
		RequestedAt respjson.Field
		State       respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3OverviewGetStatusResponseTranslationsTr) RawJSON() string { return r.JSON.raw }
func (r *PaperV3OverviewGetStatusResponseTranslationsTr) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3OverviewGetStatusResponseTranslationsUk struct {
	Error       string  `json:"error" api:"required"`
	RequestedAt float64 `json:"requestedAt" api:"required"`
	// Any of "queued", "done", "errored".
	State     string  `json:"state" api:"required"`
	UpdatedAt float64 `json:"updatedAt" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Error       respjson.Field
		RequestedAt respjson.Field
		State       respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3OverviewGetStatusResponseTranslationsUk) RawJSON() string { return r.JSON.raw }
func (r *PaperV3OverviewGetStatusResponseTranslationsUk) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3OverviewGetStatusResponseTranslationsUr struct {
	Error       string  `json:"error" api:"required"`
	RequestedAt float64 `json:"requestedAt" api:"required"`
	// Any of "queued", "done", "errored".
	State     string  `json:"state" api:"required"`
	UpdatedAt float64 `json:"updatedAt" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Error       respjson.Field
		RequestedAt respjson.Field
		State       respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3OverviewGetStatusResponseTranslationsUr) RawJSON() string { return r.JSON.raw }
func (r *PaperV3OverviewGetStatusResponseTranslationsUr) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3OverviewGetStatusResponseTranslationsUz struct {
	Error       string  `json:"error" api:"required"`
	RequestedAt float64 `json:"requestedAt" api:"required"`
	// Any of "queued", "done", "errored".
	State     string  `json:"state" api:"required"`
	UpdatedAt float64 `json:"updatedAt" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Error       respjson.Field
		RequestedAt respjson.Field
		State       respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3OverviewGetStatusResponseTranslationsUz) RawJSON() string { return r.JSON.raw }
func (r *PaperV3OverviewGetStatusResponseTranslationsUz) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3OverviewGetStatusResponseTranslationsVi struct {
	Error       string  `json:"error" api:"required"`
	RequestedAt float64 `json:"requestedAt" api:"required"`
	// Any of "queued", "done", "errored".
	State     string  `json:"state" api:"required"`
	UpdatedAt float64 `json:"updatedAt" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Error       respjson.Field
		RequestedAt respjson.Field
		State       respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3OverviewGetStatusResponseTranslationsVi) RawJSON() string { return r.JSON.raw }
func (r *PaperV3OverviewGetStatusResponseTranslationsVi) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3OverviewGetStatusResponseTranslationsYo struct {
	Error       string  `json:"error" api:"required"`
	RequestedAt float64 `json:"requestedAt" api:"required"`
	// Any of "queued", "done", "errored".
	State     string  `json:"state" api:"required"`
	UpdatedAt float64 `json:"updatedAt" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Error       respjson.Field
		RequestedAt respjson.Field
		State       respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3OverviewGetStatusResponseTranslationsYo) RawJSON() string { return r.JSON.raw }
func (r *PaperV3OverviewGetStatusResponseTranslationsYo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3OverviewGetStatusResponseTranslationsZh struct {
	Error       string  `json:"error" api:"required"`
	RequestedAt float64 `json:"requestedAt" api:"required"`
	// Any of "queued", "done", "errored".
	State     string  `json:"state" api:"required"`
	UpdatedAt float64 `json:"updatedAt" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Error       respjson.Field
		RequestedAt respjson.Field
		State       respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3OverviewGetStatusResponseTranslationsZh) RawJSON() string { return r.JSON.raw }
func (r *PaperV3OverviewGetStatusResponseTranslationsZh) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3OverviewGetParams struct {
	PaperVersion string `path:"paperVersion" api:"required" format:"uuid" json:"-"`
	paramObj
}

type PaperV3OverviewGetParamsLanguage string

const (
	PaperV3OverviewGetParamsLanguageAm PaperV3OverviewGetParamsLanguage = "am"
	PaperV3OverviewGetParamsLanguageAr PaperV3OverviewGetParamsLanguage = "ar"
	PaperV3OverviewGetParamsLanguageAz PaperV3OverviewGetParamsLanguage = "az"
	PaperV3OverviewGetParamsLanguageBg PaperV3OverviewGetParamsLanguage = "bg"
	PaperV3OverviewGetParamsLanguageBn PaperV3OverviewGetParamsLanguage = "bn"
	PaperV3OverviewGetParamsLanguageCa PaperV3OverviewGetParamsLanguage = "ca"
	PaperV3OverviewGetParamsLanguageCs PaperV3OverviewGetParamsLanguage = "cs"
	PaperV3OverviewGetParamsLanguageDa PaperV3OverviewGetParamsLanguage = "da"
	PaperV3OverviewGetParamsLanguageDe PaperV3OverviewGetParamsLanguage = "de"
	PaperV3OverviewGetParamsLanguageEl PaperV3OverviewGetParamsLanguage = "el"
	PaperV3OverviewGetParamsLanguageEn PaperV3OverviewGetParamsLanguage = "en"
	PaperV3OverviewGetParamsLanguageEs PaperV3OverviewGetParamsLanguage = "es"
	PaperV3OverviewGetParamsLanguageEt PaperV3OverviewGetParamsLanguage = "et"
	PaperV3OverviewGetParamsLanguageFa PaperV3OverviewGetParamsLanguage = "fa"
	PaperV3OverviewGetParamsLanguageFi PaperV3OverviewGetParamsLanguage = "fi"
	PaperV3OverviewGetParamsLanguageFr PaperV3OverviewGetParamsLanguage = "fr"
	PaperV3OverviewGetParamsLanguageGu PaperV3OverviewGetParamsLanguage = "gu"
	PaperV3OverviewGetParamsLanguageHa PaperV3OverviewGetParamsLanguage = "ha"
	PaperV3OverviewGetParamsLanguageHe PaperV3OverviewGetParamsLanguage = "he"
	PaperV3OverviewGetParamsLanguageHi PaperV3OverviewGetParamsLanguage = "hi"
	PaperV3OverviewGetParamsLanguageHr PaperV3OverviewGetParamsLanguage = "hr"
	PaperV3OverviewGetParamsLanguageHu PaperV3OverviewGetParamsLanguage = "hu"
	PaperV3OverviewGetParamsLanguageID PaperV3OverviewGetParamsLanguage = "id"
	PaperV3OverviewGetParamsLanguageIt PaperV3OverviewGetParamsLanguage = "it"
	PaperV3OverviewGetParamsLanguageJa PaperV3OverviewGetParamsLanguage = "ja"
	PaperV3OverviewGetParamsLanguageKa PaperV3OverviewGetParamsLanguage = "ka"
	PaperV3OverviewGetParamsLanguageKn PaperV3OverviewGetParamsLanguage = "kn"
	PaperV3OverviewGetParamsLanguageKo PaperV3OverviewGetParamsLanguage = "ko"
	PaperV3OverviewGetParamsLanguageLt PaperV3OverviewGetParamsLanguage = "lt"
	PaperV3OverviewGetParamsLanguageLv PaperV3OverviewGetParamsLanguage = "lv"
	PaperV3OverviewGetParamsLanguageMl PaperV3OverviewGetParamsLanguage = "ml"
	PaperV3OverviewGetParamsLanguageMr PaperV3OverviewGetParamsLanguage = "mr"
	PaperV3OverviewGetParamsLanguageMs PaperV3OverviewGetParamsLanguage = "ms"
	PaperV3OverviewGetParamsLanguageMy PaperV3OverviewGetParamsLanguage = "my"
	PaperV3OverviewGetParamsLanguageNe PaperV3OverviewGetParamsLanguage = "ne"
	PaperV3OverviewGetParamsLanguageNl PaperV3OverviewGetParamsLanguage = "nl"
	PaperV3OverviewGetParamsLanguageNo PaperV3OverviewGetParamsLanguage = "no"
	PaperV3OverviewGetParamsLanguagePa PaperV3OverviewGetParamsLanguage = "pa"
	PaperV3OverviewGetParamsLanguagePl PaperV3OverviewGetParamsLanguage = "pl"
	PaperV3OverviewGetParamsLanguagePt PaperV3OverviewGetParamsLanguage = "pt"
	PaperV3OverviewGetParamsLanguageRo PaperV3OverviewGetParamsLanguage = "ro"
	PaperV3OverviewGetParamsLanguageRu PaperV3OverviewGetParamsLanguage = "ru"
	PaperV3OverviewGetParamsLanguageSi PaperV3OverviewGetParamsLanguage = "si"
	PaperV3OverviewGetParamsLanguageSk PaperV3OverviewGetParamsLanguage = "sk"
	PaperV3OverviewGetParamsLanguageSl PaperV3OverviewGetParamsLanguage = "sl"
	PaperV3OverviewGetParamsLanguageSr PaperV3OverviewGetParamsLanguage = "sr"
	PaperV3OverviewGetParamsLanguageSv PaperV3OverviewGetParamsLanguage = "sv"
	PaperV3OverviewGetParamsLanguageSw PaperV3OverviewGetParamsLanguage = "sw"
	PaperV3OverviewGetParamsLanguageTa PaperV3OverviewGetParamsLanguage = "ta"
	PaperV3OverviewGetParamsLanguageTe PaperV3OverviewGetParamsLanguage = "te"
	PaperV3OverviewGetParamsLanguageTh PaperV3OverviewGetParamsLanguage = "th"
	PaperV3OverviewGetParamsLanguageTl PaperV3OverviewGetParamsLanguage = "tl"
	PaperV3OverviewGetParamsLanguageTr PaperV3OverviewGetParamsLanguage = "tr"
	PaperV3OverviewGetParamsLanguageUk PaperV3OverviewGetParamsLanguage = "uk"
	PaperV3OverviewGetParamsLanguageUr PaperV3OverviewGetParamsLanguage = "ur"
	PaperV3OverviewGetParamsLanguageUz PaperV3OverviewGetParamsLanguage = "uz"
	PaperV3OverviewGetParamsLanguageVi PaperV3OverviewGetParamsLanguage = "vi"
	PaperV3OverviewGetParamsLanguageYo PaperV3OverviewGetParamsLanguage = "yo"
	PaperV3OverviewGetParamsLanguageZh PaperV3OverviewGetParamsLanguage = "zh"
)
