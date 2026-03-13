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

// AnalyticsPaperViewCountService contains methods and other services that help
// with interacting with the alphaxiv_cat API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAnalyticsPaperViewCountService] method instead.
type AnalyticsPaperViewCountService struct {
	options []option.RequestOption
}

// NewAnalyticsPaperViewCountService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAnalyticsPaperViewCountService(opts ...option.RequestOption) (r AnalyticsPaperViewCountService) {
	r = AnalyticsPaperViewCountService{}
	r.options = opts
	return
}

// Track a paper view event for analytics
//
// Source file:
// `api-server/src/controllers/v1/analytics/ingest-paper-view-count-event.controller.ts`
//
// Deprecated: deprecated
func (r *AnalyticsPaperViewCountService) IngestEvent(ctx context.Context, body AnalyticsPaperViewCountIngestEventParams, opts ...option.RequestOption) (res *AnalyticsPaperViewCountIngestEventResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/analytics/paper-view-count/ingest"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Kicks off a background job to aggregate paper view counts
//
// Source file:
// `api-server/src/controllers/v1/analytics/kickoff-paper-view-count-aggregation-job.controller.ts`
func (r *AnalyticsPaperViewCountService) KickoffJob(ctx context.Context, body AnalyticsPaperViewCountKickoffJobParams, opts ...option.RequestOption) (res *AnalyticsPaperViewCountKickoffJobResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/analytics/paper-view-count/kickoff-job"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Process view count aggregation for a specific paper
//
// Source file:
// `api-server/src/controllers/v1/analytics/process-paper-view-count-aggregation-job.controller.ts`
func (r *AnalyticsPaperViewCountService) ProcessJob(ctx context.Context, body AnalyticsPaperViewCountProcessJobParams, opts ...option.RequestOption) (res *AnalyticsPaperViewCountProcessJobResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/analytics/paper-view-count/process-job"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type AnalyticsPaperViewCountIngestEventResponse struct {
	Data AnalyticsPaperViewCountIngestEventResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AnalyticsPaperViewCountIngestEventResponse) RawJSON() string { return r.JSON.raw }
func (r *AnalyticsPaperViewCountIngestEventResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AnalyticsPaperViewCountIngestEventResponseData struct {
	IngestedPaper string `json:"ingestedPaper" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IngestedPaper respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AnalyticsPaperViewCountIngestEventResponseData) RawJSON() string { return r.JSON.raw }
func (r *AnalyticsPaperViewCountIngestEventResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AnalyticsPaperViewCountKickoffJobResponse struct {
	Data AnalyticsPaperViewCountKickoffJobResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AnalyticsPaperViewCountKickoffJobResponse) RawJSON() string { return r.JSON.raw }
func (r *AnalyticsPaperViewCountKickoffJobResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AnalyticsPaperViewCountKickoffJobResponseData struct {
	Message string `json:"message" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AnalyticsPaperViewCountKickoffJobResponseData) RawJSON() string { return r.JSON.raw }
func (r *AnalyticsPaperViewCountKickoffJobResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AnalyticsPaperViewCountProcessJobResponse struct {
	Data AnalyticsPaperViewCountProcessJobResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AnalyticsPaperViewCountProcessJobResponse) RawJSON() string { return r.JSON.raw }
func (r *AnalyticsPaperViewCountProcessJobResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AnalyticsPaperViewCountProcessJobResponseData struct {
	ProcessedPaper string `json:"processedPaper" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ProcessedPaper respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AnalyticsPaperViewCountProcessJobResponseData) RawJSON() string { return r.JSON.raw }
func (r *AnalyticsPaperViewCountProcessJobResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AnalyticsPaperViewCountIngestEventParams struct {
	// Paper ID to track view for
	PaperID string `json:"paperId" api:"required"`
	// Optional user ID who viewed the paper
	UserID param.Opt[string] `json:"userId,omitzero"`
	// Optional timestamp for the view event
	CreatedAt param.Opt[string] `json:"createdAt,omitzero"`
	paramObj
}

func (r AnalyticsPaperViewCountIngestEventParams) MarshalJSON() (data []byte, err error) {
	type shadow AnalyticsPaperViewCountIngestEventParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AnalyticsPaperViewCountIngestEventParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AnalyticsPaperViewCountKickoffJobParams struct {
	// Time period filter: 'all' or number of days
	Type string `json:"type" api:"required"`
	// Optional filter for likes
	Like param.Opt[string] `json:"like,omitzero"`
	paramObj
}

func (r AnalyticsPaperViewCountKickoffJobParams) MarshalJSON() (data []byte, err error) {
	type shadow AnalyticsPaperViewCountKickoffJobParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AnalyticsPaperViewCountKickoffJobParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AnalyticsPaperViewCountProcessJobParams struct {
	// Paper ID to process view counts for
	PaperID string `json:"paperId" api:"required"`
	// Publication date for age decay calculation
	PublicationDate string `json:"publicationDate" api:"required"`
	// Whether to add noise to votes
	Like param.Opt[bool] `json:"like,omitzero"`
	paramObj
}

func (r AnalyticsPaperViewCountProcessJobParams) MarshalJSON() (data []byte, err error) {
	type shadow AnalyticsPaperViewCountProcessJobParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AnalyticsPaperViewCountProcessJobParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
