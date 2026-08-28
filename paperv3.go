// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package alphaxivcat

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/AlphaxivCat/alphaxiv_cat-go/internal/apijson"
	"github.com/AlphaxivCat/alphaxiv_cat-go/internal/apiquery"
	shimjson "github.com/AlphaxivCat/alphaxiv_cat-go/internal/encoding/json"
	"github.com/AlphaxivCat/alphaxiv_cat-go/internal/requestconfig"
	"github.com/AlphaxivCat/alphaxiv_cat-go/option"
	"github.com/AlphaxivCat/alphaxiv_cat-go/packages/param"
	"github.com/AlphaxivCat/alphaxiv_cat-go/packages/respjson"
)

// PaperV3Service contains methods and other services that help with interacting
// with the alphaxiv_cat API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPaperV3Service] method instead.
type PaperV3Service struct {
	options         []option.RequestOption
	Legacy          PaperV3LegacyService
	Overview        PaperV3OverviewService
	Implementations PaperV3ImplementationService
	XMentions       PaperV3XMentionService
}

// NewPaperV3Service generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewPaperV3Service(opts ...option.RequestOption) (r PaperV3Service) {
	r = PaperV3Service{}
	r.options = opts
	r.Legacy = NewPaperV3LegacyService(opts...)
	r.Overview = NewPaperV3OverviewService(opts...)
	r.Implementations = NewPaperV3ImplementationService(opts...)
	r.XMentions = NewPaperV3XMentionService(opts...)
	return
}

// Retrieve paper version metadata. Fetches from ArXiv if needed.
//
// Source file:
// `api-server/file:/app/api-server/src/controllers/papers/v3/get-paper.controller.ts`
func (r *PaperV3Service) Get(ctx context.Context, unresolved string, opts ...option.RequestOption) (res *PaperV3GetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if unresolved == "" {
		err = errors.New("missing required unresolved parameter")
		return nil, err
	}
	path := fmt.Sprintf("papers/v3/%s", url.PathEscape(unresolved))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Create a public comment or private note on a paper.
//
// Source file:
// `api-server/file:/app/api-server/src/controllers/papers/v3/post-comment.controller.ts`
func (r *PaperV3Service) Comment(ctx context.Context, version string, body PaperV3CommentParams, opts ...option.RequestOption) (res *PaperV3CommentResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if version == "" {
		err = errors.New("missing required version parameter")
		return nil, err
	}
	path := fmt.Sprintf("papers/v3/%s/comment", version)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Remove votes from many papers at once
//
// Source file:
// `api-server/file:/app/api-server/src/controllers/papers/v3/remove-vote-batch.controller.ts`
func (r *PaperV3Service) DeleteVotes(ctx context.Context, body PaperV3DeleteVotesParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "papers/v3/votes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, nil, opts...)
	return err
}

// Create or update an implementation for a paper group
//
// Source file:
// `api-server/file:/app/api-server/src/controllers/papers/v3/create-or-update-implementation.controller.ts`
func (r *PaperV3Service) Implementation(ctx context.Context, paperGroupID string, body PaperV3ImplementationParams, opts ...option.RequestOption) (res *PaperV3ImplementationResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if paperGroupID == "" {
		err = errors.New("missing required paperGroupId parameter")
		return nil, err
	}
	path := fmt.Sprintf("papers/v3/%s/implementation", paperGroupID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Kickoff paper countries processing for hot papers
//
// Source file:
// `api-server/file:/app/api-server/src/controllers/papers/v3/kickoff-paper-countries.controller.ts`
func (r *PaperV3Service) KickoffPaperCountries(ctx context.Context, body PaperV3KickoffPaperCountriesParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "papers/v3/kickoff-paper-countries"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return err
}

// Kickoff paper podcasts on Uptash for a subset of paper groups
//
// Source file:
// `api-server/file:/app/api-server/src/controllers/papers/v3/kickoff-paper-podcasts.controller.ts`
func (r *PaperV3Service) KickoffPaperPodcasts(ctx context.Context, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "papers/v3/kickoff-paper-podcasts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return err
}

// Set your like status on a paper group
//
// Source file:
// `api-server/file:/app/api-server/src/controllers/papers/v3/like-paper.controller.ts`
func (r *PaperV3Service) Like(ctx context.Context, group string, body PaperV3LikeParams, opts ...option.RequestOption) (res *PaperV3LikeResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if group == "" {
		err = errors.New("missing required group parameter")
		return nil, err
	}
	path := fmt.Sprintf("papers/v3/%s/like", group)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Generates a podcast for a paper group
//
// Source file:
// `api-server/file:/app/api-server/src/controllers/papers/v3/generate-paper-podcast.controller.ts`
//
// Deprecated: deprecated
func (r *PaperV3Service) Podcast(ctx context.Context, paperGroupID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if paperGroupID == "" {
		err = errors.New("missing required paperGroupId parameter")
		return err
	}
	path := fmt.Sprintf("papers/v3/%s/podcast", paperGroupID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return err
}

// Generates AI overviews for a paper version
//
// Source file:
// `api-server/file:/app/api-server/src/controllers/papers/v3/process-ai.controller.ts`
//
// Deprecated: deprecated
func (r *PaperV3Service) ProcessAI(ctx context.Context, paperVersionID string, body PaperV3ProcessAIParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if paperVersionID == "" {
		err = errors.New("missing required paperVersionId parameter")
		return err
	}
	path := fmt.Sprintf("papers/v3/%s/process-ai", paperVersionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return err
}

// Processes and generates country metadata for papers based on institutional
// affiliations
//
// Source file:
// `api-server/file:/app/api-server/src/controllers/papers/v3/process-countries.controller.ts`
//
// Deprecated: deprecated
func (r *PaperV3Service) ProcessCountries(ctx context.Context, body PaperV3ProcessCountriesParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "papers/v3/process-countries"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return err
}

// Toggle your implementation request status on a paper group
//
// Source file:
// `api-server/file:/app/api-server/src/controllers/papers/v3/request-paper-implementation.controller.ts`
func (r *PaperV3Service) RequestImplementation(ctx context.Context, group string, body PaperV3RequestImplementationParams, opts ...option.RequestOption) (res *PaperV3RequestImplementationResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if group == "" {
		err = errors.New("missing required group parameter")
		return nil, err
	}
	path := fmt.Sprintf("papers/v3/%s/request-implementation", group)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Request podcast generation for a paper group
//
// Source file:
// `api-server/file:/app/api-server/src/controllers/papers/v3/request-podcast.controller.ts`
func (r *PaperV3Service) RequestPodcast(ctx context.Context, paperGroupID string, opts ...option.RequestOption) (res *PaperV3RequestPodcastResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if paperGroupID == "" {
		err = errors.New("missing required paperGroupId parameter")
		return nil, err
	}
	path := fmt.Sprintf("papers/v3/%s/request-podcast", paperGroupID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Get all paper universal IDs sorted by most recent publication date
//
// Source file:
// `api-server/file:/app/api-server/src/controllers/papers/v3/get-all-papers.controller.ts`
func (r *PaperV3Service) GetAll(ctx context.Context, query PaperV3GetAllParams, opts ...option.RequestOption) (res *PaperV3GetAllResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "papers/v3/all"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Get an initial batch of diverse papers on the given topics for recommendations
//
// Source file:
// `api-server/file:/app/api-server/src/controllers/papers/v3/diverse-papers.controller.ts`
func (r *PaperV3Service) GetDiversePapers(ctx context.Context, query PaperV3GetDiversePapersParams, opts ...option.RequestOption) (res *[]PaperV3GetDiversePapersResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "papers/v3/diverse-papers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Get an optionally filtered list of papers for the main feed
//
// Source file:
// `api-server/file:/app/api-server/src/controllers/papers/v3/feed.controller.ts`
func (r *PaperV3Service) GetFeed(ctx context.Context, query PaperV3GetFeedParams, opts ...option.RequestOption) (res *PaperV3GetFeedResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "papers/v3/feed"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Get list of figure URLs for a paper
//
// Source file:
// `api-server/file:/app/api-server/src/controllers/papers/v3/get-paper-figures.controller.ts`
func (r *PaperV3Service) GetFigures(ctx context.Context, paperGroupID string, opts ...option.RequestOption) (res *PaperV3GetFiguresResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if paperGroupID == "" {
		err = errors.New("missing required paperGroupId parameter")
		return nil, err
	}
	path := fmt.Sprintf("papers/v3/%s/figures", paperGroupID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Get the full extracted text of a paper, page by page
//
// Source file:
// `api-server/file:/app/api-server/src/controllers/papers/v3/get-full-text.controller.ts`
func (r *PaperV3Service) GetFullText(ctx context.Context, paperVersion string, opts ...option.RequestOption) (res *PaperV3GetFullTextResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if paperVersion == "" {
		err = errors.New("missing required paperVersion parameter")
		return nil, err
	}
	path := fmt.Sprintf("papers/v3/%s/full-text", paperVersion)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Retrieve metrics for a paper (comments count, upvotes, views)
//
// Source file:
// `api-server/file:/app/api-server/src/controllers/papers/v3/get-metrics.controller.ts`
func (r *PaperV3Service) GetMetrics(ctx context.Context, unresolved string, opts ...option.RequestOption) (res *PaperV3GetMetricsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if unresolved == "" {
		err = errors.New("missing required unresolved parameter")
		return nil, err
	}
	path := fmt.Sprintf("papers/v3/%s/metrics", url.PathEscape(unresolved))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Retrieve paper data for paper preview cards
//
// Source file:
// `api-server/file:/app/api-server/src/controllers/papers/v3/get-paper-preview.controller.ts`
func (r *PaperV3Service) GetPreview(ctx context.Context, id string, opts ...option.RequestOption) (res *PaperV3GetPreviewResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("papers/v3/%s/preview", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Get papers semantically similar to the selected one
//
// Source file:
// `api-server/file:/app/api-server/src/controllers/papers/v3/get-similar-papers.controller.ts`
func (r *PaperV3Service) GetSimilarPapers(ctx context.Context, id string, query PaperV3GetSimilarPapersParams, opts ...option.RequestOption) (res *[]PaperV3GetSimilarPapersResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("papers/v3/%s/similar-papers", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Get some papers on the provided topics that are unrelated to the provided papers
//
// Source file:
// `api-server/file:/app/api-server/src/controllers/papers/v3/unrelated-papers.controller.ts`
func (r *PaperV3Service) GetUnrelated(ctx context.Context, query PaperV3GetUnrelatedParams, opts ...option.RequestOption) (res *[]PaperV3GetUnrelatedResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "papers/v3/unrelated"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Track paper view event for analytics
//
// Source file:
// `api-server/file:/app/api-server/src/controllers/papers/v3/mark-paper-view.controller.ts`
func (r *PaperV3Service) View(ctx context.Context, group string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if group == "" {
		err = errors.New("missing required group parameter")
		return err
	}
	path := fmt.Sprintf("papers/v3/%s/view", group)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return err
}

type PaperV3GetResponse struct {
	Abstract             string                       `json:"abstract" api:"required"`
	CitationBibtex       string                       `json:"citationBibtex" api:"required"`
	CitationsCount       float64                      `json:"citationsCount" api:"required"`
	FirstPublicationDate float64                      `json:"firstPublicationDate" api:"required"`
	GoogleCitationID     string                       `json:"googleCitationId" api:"required"`
	GroupID              string                       `json:"groupId" api:"required" format:"uuid"`
	License              string                       `json:"license" api:"required"`
	PdfOnly              bool                         `json:"pdfOnly" api:"required"`
	PublicationDate      float64                      `json:"publicationDate" api:"required"`
	Resources            []PaperV3GetResponseResource `json:"resources" api:"required"`
	SourceName           string                       `json:"sourceName" api:"required"`
	SourceURL            string                       `json:"sourceUrl" api:"required"`
	Title                string                       `json:"title" api:"required"`
	// Any of "public", "private".
	Type         PaperV3GetResponseType      `json:"type" api:"required"`
	UniversalID  string                      `json:"universalId" api:"required"`
	Uploader     string                      `json:"uploader" api:"required" format:"uuid"`
	VersionID    string                      `json:"versionId" api:"required" format:"uuid"`
	VersionLabel string                      `json:"versionLabel" api:"required"`
	VersionOrder float64                     `json:"versionOrder" api:"required"`
	Versions     []PaperV3GetResponseVersion `json:"versions" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Abstract             respjson.Field
		CitationBibtex       respjson.Field
		CitationsCount       respjson.Field
		FirstPublicationDate respjson.Field
		GoogleCitationID     respjson.Field
		GroupID              respjson.Field
		License              respjson.Field
		PdfOnly              respjson.Field
		PublicationDate      respjson.Field
		Resources            respjson.Field
		SourceName           respjson.Field
		SourceURL            respjson.Field
		Title                respjson.Field
		Type                 respjson.Field
		UniversalID          respjson.Field
		Uploader             respjson.Field
		VersionID            respjson.Field
		VersionLabel         respjson.Field
		VersionOrder         respjson.Field
		Versions             respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3GetResponse) RawJSON() string { return r.JSON.raw }
func (r *PaperV3GetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3GetResponseResource struct {
	ID          string  `json:"id" api:"required" format:"uuid"`
	Description string  `json:"description" api:"required"`
	Language    string  `json:"language" api:"required"`
	Stars       float64 `json:"stars" api:"required"`
	// Any of "github".
	Type string `json:"type" api:"required"`
	URL  string `json:"url" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Description respjson.Field
		Language    respjson.Field
		Stars       respjson.Field
		Type        respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3GetResponseResource) RawJSON() string { return r.JSON.raw }
func (r *PaperV3GetResponseResource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3GetResponseType string

const (
	PaperV3GetResponseTypePublic  PaperV3GetResponseType = "public"
	PaperV3GetResponseTypePrivate PaperV3GetResponseType = "private"
)

type PaperV3GetResponseVersion struct {
	ID    string  `json:"id" api:"required" format:"uuid"`
	Label string  `json:"label" api:"required"`
	Order float64 `json:"order" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Label       respjson.Field
		Order       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3GetResponseVersion) RawJSON() string { return r.JSON.raw }
func (r *PaperV3GetResponseVersion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3CommentResponse struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3CommentResponse) RawJSON() string { return r.JSON.raw }
func (r *PaperV3CommentResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3ImplementationResponse struct {
	Implementation PaperV3ImplementationResponseImplementation `json:"implementation" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Implementation respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3ImplementationResponse) RawJSON() string { return r.JSON.raw }
func (r *PaperV3ImplementationResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3ImplementationResponseImplementation struct {
	ID           string `json:"id" api:"required" format:"uuid"`
	PaperGroupID string `json:"paperGroupId" api:"required" format:"uuid"`
	// Any of "github", "marimo".
	Type string `json:"type" api:"required"`
	URL  string `json:"url" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		PaperGroupID respjson.Field
		Type         respjson.Field
		URL          respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3ImplementationResponseImplementation) RawJSON() string { return r.JSON.raw }
func (r *PaperV3ImplementationResponseImplementation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3LikeResponse struct {
	Liked bool `json:"liked" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Liked       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3LikeResponse) RawJSON() string { return r.JSON.raw }
func (r *PaperV3LikeResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3RequestImplementationResponse struct {
	Requested bool `json:"requested" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Requested   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3RequestImplementationResponse) RawJSON() string { return r.JSON.raw }
func (r *PaperV3RequestImplementationResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3RequestPodcastResponse struct {
	Message string `json:"message" api:"required"`
	// Any of "queued", "generating", "done", "errored".
	State PaperV3RequestPodcastResponseState `json:"state" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		State       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3RequestPodcastResponse) RawJSON() string { return r.JSON.raw }
func (r *PaperV3RequestPodcastResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3RequestPodcastResponseState string

const (
	PaperV3RequestPodcastResponseStateQueued     PaperV3RequestPodcastResponseState = "queued"
	PaperV3RequestPodcastResponseStateGenerating PaperV3RequestPodcastResponseState = "generating"
	PaperV3RequestPodcastResponseStateDone       PaperV3RequestPodcastResponseState = "done"
	PaperV3RequestPodcastResponseStateErrored    PaperV3RequestPodcastResponseState = "errored"
)

type PaperV3GetAllResponse struct {
	Limit        float64  `json:"limit" api:"required"`
	Skip         float64  `json:"skip" api:"required"`
	UniversalIDs []string `json:"universalIds" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit        respjson.Field
		Skip         respjson.Field
		UniversalIDs respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3GetAllResponse) RawJSON() string { return r.JSON.raw }
func (r *PaperV3GetAllResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3GetDiversePapersResponse struct {
	ID         string                                      `json:"id" api:"required" format:"uuid"`
	Abstract   string                                      `json:"abstract" api:"required"`
	AuthorInfo []PaperV3GetDiversePapersResponseAuthorInfo `json:"author_info" api:"required"`
	Authors    []string                                    `json:"authors" api:"required"`
	// A versioned paper ID (e.g. 1706.03762v1)
	CanonicalID          string                                            `json:"canonical_id" api:"required"`
	ExternalBlog         PaperV3GetDiversePapersResponseExternalBlog       `json:"external_blog" api:"required"`
	FirstPublicationDate string                                            `json:"first_publication_date" api:"required"`
	FullAuthors          []PaperV3GetDiversePapersResponseFullAuthor       `json:"full_authors" api:"required"`
	FullAuthorsV2        []PaperV3GetDiversePapersResponseFullAuthorsV2    `json:"full_authors_v2" api:"required"`
	GitHubStars          float64                                           `json:"github_stars" api:"required"`
	GitHubURL            string                                            `json:"github_url" api:"required"`
	HasRunReport         bool                                              `json:"has_run_report" api:"required"`
	ImageURL             string                                            `json:"image_url" api:"required"`
	Metrics              PaperV3GetDiversePapersResponseMetrics            `json:"metrics" api:"required"`
	OrganizationInfo     []PaperV3GetDiversePapersResponseOrganizationInfo `json:"organization_info" api:"required"`
	PaperGroupID         string                                            `json:"paper_group_id" api:"required" format:"uuid"`
	PaperSummary         PaperV3GetDiversePapersResponsePaperSummary       `json:"paper_summary" api:"required"`
	PdfOnly              bool                                              `json:"pdf_only" api:"required"`
	PublicationDate      string                                            `json:"publication_date" api:"required"`
	Title                string                                            `json:"title" api:"required"`
	Topics               []string                                          `json:"topics" api:"required"`
	// A versionless universal paper ID (e.g. 1706.03762)
	UniversalPaperID      string                                               `json:"universal_paper_id" api:"required"`
	UpdatedAt             string                                               `json:"updated_at" api:"required"`
	VersionID             string                                               `json:"version_id" api:"required" format:"uuid"`
	CardPreviewBlobID     string                                               `json:"card_preview_blob_id" api:"nullable" format:"uuid"`
	NarrationAudioURL     string                                               `json:"narration_audio_url" api:"nullable"`
	RecommendationContext PaperV3GetDiversePapersResponseRecommendationContext `json:"recommendation_context"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                    respjson.Field
		Abstract              respjson.Field
		AuthorInfo            respjson.Field
		Authors               respjson.Field
		CanonicalID           respjson.Field
		ExternalBlog          respjson.Field
		FirstPublicationDate  respjson.Field
		FullAuthors           respjson.Field
		FullAuthorsV2         respjson.Field
		GitHubStars           respjson.Field
		GitHubURL             respjson.Field
		HasRunReport          respjson.Field
		ImageURL              respjson.Field
		Metrics               respjson.Field
		OrganizationInfo      respjson.Field
		PaperGroupID          respjson.Field
		PaperSummary          respjson.Field
		PdfOnly               respjson.Field
		PublicationDate       respjson.Field
		Title                 respjson.Field
		Topics                respjson.Field
		UniversalPaperID      respjson.Field
		UpdatedAt             respjson.Field
		VersionID             respjson.Field
		CardPreviewBlobID     respjson.Field
		NarrationAudioURL     respjson.Field
		RecommendationContext respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3GetDiversePapersResponse) RawJSON() string { return r.JSON.raw }
func (r *PaperV3GetDiversePapersResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3GetDiversePapersResponseAuthorInfo struct {
	ID               string                                            `json:"id" api:"required" format:"uuid"`
	Avatar           []PaperV3GetDiversePapersResponseAuthorInfoAvatar `json:"avatar" api:"required"`
	BlueskyUsername  string                                            `json:"blueskyUsername" api:"required"`
	GitHubUsername   string                                            `json:"githubUsername" api:"required"`
	GoogleScholarID  string                                            `json:"googleScholarId" api:"required"`
	Institution      string                                            `json:"institution" api:"required"`
	LinkedinUsername string                                            `json:"linkedinUsername" api:"required"`
	OrcidID          string                                            `json:"orcidId" api:"required"`
	PublicEmail      string                                            `json:"publicEmail" api:"required"`
	RealName         string                                            `json:"realName" api:"required"`
	Reputation       float64                                           `json:"reputation" api:"required"`
	ResearcherSlug   string                                            `json:"researcherSlug" api:"required"`
	// Any of "user", "reviewer", "admin", "bot".
	Role             string  `json:"role" api:"required"`
	Username         string  `json:"username" api:"required"`
	Verified         bool    `json:"verified" api:"required"`
	WeeklyReputation float64 `json:"weeklyReputation" api:"required"`
	XUsername        string  `json:"xUsername" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		Avatar           respjson.Field
		BlueskyUsername  respjson.Field
		GitHubUsername   respjson.Field
		GoogleScholarID  respjson.Field
		Institution      respjson.Field
		LinkedinUsername respjson.Field
		OrcidID          respjson.Field
		PublicEmail      respjson.Field
		RealName         respjson.Field
		Reputation       respjson.Field
		ResearcherSlug   respjson.Field
		Role             respjson.Field
		Username         respjson.Field
		Verified         respjson.Field
		WeeklyReputation respjson.Field
		XUsername        respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3GetDiversePapersResponseAuthorInfo) RawJSON() string { return r.JSON.raw }
func (r *PaperV3GetDiversePapersResponseAuthorInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3GetDiversePapersResponseAuthorInfoAvatar struct {
	// Any of "full_size", "thumbnail".
	Type string `json:"type" api:"required"`
	URL  string `json:"url" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3GetDiversePapersResponseAuthorInfoAvatar) RawJSON() string { return r.JSON.raw }
func (r *PaperV3GetDiversePapersResponseAuthorInfoAvatar) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3GetDiversePapersResponseExternalBlog struct {
	BodyBlobID  string `json:"body_blob_id" api:"required" format:"uuid"`
	CoverBlobID string `json:"cover_blob_id" api:"required" format:"uuid"`
	SourceName  string `json:"source_name" api:"required"`
	SourceURL   string `json:"source_url" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BodyBlobID  respjson.Field
		CoverBlobID respjson.Field
		SourceName  respjson.Field
		SourceURL   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3GetDiversePapersResponseExternalBlog) RawJSON() string { return r.JSON.raw }
func (r *PaperV3GetDiversePapersResponseExternalBlog) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3GetDiversePapersResponseFullAuthor struct {
	ID             string `json:"id" api:"required" format:"uuid"`
	FullName       string `json:"full_name" api:"required"`
	UserID         string `json:"user_id" api:"required" format:"uuid"`
	Username       string `json:"username" api:"required"`
	ResearcherSlug string `json:"researcher_slug" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		FullName       respjson.Field
		UserID         respjson.Field
		Username       respjson.Field
		ResearcherSlug respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3GetDiversePapersResponseFullAuthor) RawJSON() string { return r.JSON.raw }
func (r *PaperV3GetDiversePapersResponseFullAuthor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3GetDiversePapersResponseFullAuthorsV2 struct {
	FullName   string                                                 `json:"full_name" api:"required"`
	Researcher PaperV3GetDiversePapersResponseFullAuthorsV2Researcher `json:"researcher" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FullName    respjson.Field
		Researcher  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3GetDiversePapersResponseFullAuthorsV2) RawJSON() string { return r.JSON.raw }
func (r *PaperV3GetDiversePapersResponseFullAuthorsV2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3GetDiversePapersResponseFullAuthorsV2Researcher struct {
	Affiliation   string                                                            `json:"affiliation" api:"required"`
	Bio           string                                                            `json:"bio" api:"required"`
	Citations     float64                                                           `json:"citations" api:"required"`
	Headline      string                                                            `json:"headline" api:"required"`
	HIndex        float64                                                           `json:"hIndex" api:"required"`
	LinkedUser    PaperV3GetDiversePapersResponseFullAuthorsV2ResearcherLinkedUser  `json:"linkedUser" api:"required"`
	Links         PaperV3GetDiversePapersResponseFullAuthorsV2ResearcherLinks       `json:"links" api:"required"`
	Name          string                                                            `json:"name" api:"required"`
	PhotoURL      string                                                            `json:"photoUrl" api:"required"`
	ResearchAreas []string                                                          `json:"researchAreas" api:"required"`
	Slug          string                                                            `json:"slug" api:"required"`
	Reason        PaperV3GetDiversePapersResponseFullAuthorsV2ResearcherReasonUnion `json:"reason"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Affiliation   respjson.Field
		Bio           respjson.Field
		Citations     respjson.Field
		Headline      respjson.Field
		HIndex        respjson.Field
		LinkedUser    respjson.Field
		Links         respjson.Field
		Name          respjson.Field
		PhotoURL      respjson.Field
		ResearchAreas respjson.Field
		Slug          respjson.Field
		Reason        respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3GetDiversePapersResponseFullAuthorsV2Researcher) RawJSON() string { return r.JSON.raw }
func (r *PaperV3GetDiversePapersResponseFullAuthorsV2Researcher) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3GetDiversePapersResponseFullAuthorsV2ResearcherLinkedUser struct {
	Name     string `json:"name" api:"required"`
	Username string `json:"username" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Username    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3GetDiversePapersResponseFullAuthorsV2ResearcherLinkedUser) RawJSON() string {
	return r.JSON.raw
}
func (r *PaperV3GetDiversePapersResponseFullAuthorsV2ResearcherLinkedUser) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3GetDiversePapersResponseFullAuthorsV2ResearcherLinks struct {
	Bluesky      string `json:"bluesky" api:"required"`
	Cv           string `json:"cv" api:"required"`
	Dblp         string `json:"dblp" api:"required"`
	Email        string `json:"email" api:"required"`
	GitHub       string `json:"github" api:"required"`
	Huggingface  string `json:"huggingface" api:"required"`
	Linkedin     string `json:"linkedin" api:"required"`
	Openreview   string `json:"openreview" api:"required"`
	Orcid        string `json:"orcid" api:"required"`
	PersonalSite string `json:"personalSite" api:"required"`
	Scholar      string `json:"scholar" api:"required"`
	Twitter      string `json:"twitter" api:"required"`
	Wikipedia    string `json:"wikipedia" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Bluesky      respjson.Field
		Cv           respjson.Field
		Dblp         respjson.Field
		Email        respjson.Field
		GitHub       respjson.Field
		Huggingface  respjson.Field
		Linkedin     respjson.Field
		Openreview   respjson.Field
		Orcid        respjson.Field
		PersonalSite respjson.Field
		Scholar      respjson.Field
		Twitter      respjson.Field
		Wikipedia    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3GetDiversePapersResponseFullAuthorsV2ResearcherLinks) RawJSON() string {
	return r.JSON.raw
}
func (r *PaperV3GetDiversePapersResponseFullAuthorsV2ResearcherLinks) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PaperV3GetDiversePapersResponseFullAuthorsV2ResearcherReasonUnion contains all
// possible properties and values from
// [PaperV3GetDiversePapersResponseFullAuthorsV2ResearcherReasonObject],
// [PaperV3GetDiversePapersResponseFullAuthorsV2ResearcherReasonKind],
// [PaperV3GetDiversePapersResponseFullAuthorsV2ResearcherReasonObject2].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type PaperV3GetDiversePapersResponseFullAuthorsV2ResearcherReasonUnion struct {
	Kind string `json:"kind"`
	// This field is from variant
	// [PaperV3GetDiversePapersResponseFullAuthorsV2ResearcherReasonObject].
	PaperTitle string `json:"paperTitle"`
	// This field is from variant
	// [PaperV3GetDiversePapersResponseFullAuthorsV2ResearcherReasonObject2].
	Count float64 `json:"count"`
	// This field is from variant
	// [PaperV3GetDiversePapersResponseFullAuthorsV2ResearcherReasonObject2].
	Followed PaperV3GetDiversePapersResponseFullAuthorsV2ResearcherReasonObject2Followed `json:"followed"`
	JSON     struct {
		Kind       respjson.Field
		PaperTitle respjson.Field
		Count      respjson.Field
		Followed   respjson.Field
		raw        string
	} `json:"-"`
}

func (u PaperV3GetDiversePapersResponseFullAuthorsV2ResearcherReasonUnion) AsPaperV3GetDiversePapersResponseFullAuthorsV2ResearcherReasonObject() (v PaperV3GetDiversePapersResponseFullAuthorsV2ResearcherReasonObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PaperV3GetDiversePapersResponseFullAuthorsV2ResearcherReasonUnion) AsPaperV3GetDiversePapersResponseFullAuthorsV2ResearcherReasonKind() (v PaperV3GetDiversePapersResponseFullAuthorsV2ResearcherReasonKind) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PaperV3GetDiversePapersResponseFullAuthorsV2ResearcherReasonUnion) AsPaperV3GetDiversePapersResponseFullAuthorsV2ResearcherReasonObject2() (v PaperV3GetDiversePapersResponseFullAuthorsV2ResearcherReasonObject2) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PaperV3GetDiversePapersResponseFullAuthorsV2ResearcherReasonUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *PaperV3GetDiversePapersResponseFullAuthorsV2ResearcherReasonUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3GetDiversePapersResponseFullAuthorsV2ResearcherReasonObject struct {
	// Any of "interest".
	Kind       string `json:"kind" api:"required"`
	PaperTitle string `json:"paperTitle"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Kind        respjson.Field
		PaperTitle  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3GetDiversePapersResponseFullAuthorsV2ResearcherReasonObject) RawJSON() string {
	return r.JSON.raw
}
func (r *PaperV3GetDiversePapersResponseFullAuthorsV2ResearcherReasonObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3GetDiversePapersResponseFullAuthorsV2ResearcherReasonKind struct {
	// Any of "read".
	Kind string `json:"kind" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Kind        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3GetDiversePapersResponseFullAuthorsV2ResearcherReasonKind) RawJSON() string {
	return r.JSON.raw
}
func (r *PaperV3GetDiversePapersResponseFullAuthorsV2ResearcherReasonKind) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3GetDiversePapersResponseFullAuthorsV2ResearcherReasonObject2 struct {
	Count float64 `json:"count" api:"required"`
	// Any of "coauthor".
	Kind     string                                                                      `json:"kind" api:"required"`
	Followed PaperV3GetDiversePapersResponseFullAuthorsV2ResearcherReasonObject2Followed `json:"followed"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Kind        respjson.Field
		Followed    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3GetDiversePapersResponseFullAuthorsV2ResearcherReasonObject2) RawJSON() string {
	return r.JSON.raw
}
func (r *PaperV3GetDiversePapersResponseFullAuthorsV2ResearcherReasonObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3GetDiversePapersResponseFullAuthorsV2ResearcherReasonObject2Followed struct {
	Name string `json:"name" api:"required"`
	Slug string `json:"slug" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Slug        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3GetDiversePapersResponseFullAuthorsV2ResearcherReasonObject2Followed) RawJSON() string {
	return r.JSON.raw
}
func (r *PaperV3GetDiversePapersResponseFullAuthorsV2ResearcherReasonObject2Followed) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3GetDiversePapersResponseMetrics struct {
	PublicTotalVotes float64                                           `json:"public_total_votes" api:"required"`
	TotalVotes       float64                                           `json:"total_votes" api:"required"`
	VisitsCount      PaperV3GetDiversePapersResponseMetricsVisitsCount `json:"visits_count" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PublicTotalVotes respjson.Field
		TotalVotes       respjson.Field
		VisitsCount      respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3GetDiversePapersResponseMetrics) RawJSON() string { return r.JSON.raw }
func (r *PaperV3GetDiversePapersResponseMetrics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3GetDiversePapersResponseMetricsVisitsCount struct {
	All       float64 `json:"all" api:"required"`
	Last7Days float64 `json:"last_7_days" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		All         respjson.Field
		Last7Days   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3GetDiversePapersResponseMetricsVisitsCount) RawJSON() string { return r.JSON.raw }
func (r *PaperV3GetDiversePapersResponseMetricsVisitsCount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3GetDiversePapersResponseOrganizationInfo struct {
	Image string `json:"image" api:"required"`
	Name  string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Image       respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3GetDiversePapersResponseOrganizationInfo) RawJSON() string { return r.JSON.raw }
func (r *PaperV3GetDiversePapersResponseOrganizationInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3GetDiversePapersResponsePaperSummary struct {
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
func (r PaperV3GetDiversePapersResponsePaperSummary) RawJSON() string { return r.JSON.raw }
func (r *PaperV3GetDiversePapersResponsePaperSummary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3GetDiversePapersResponseRecommendationContext struct {
	FollowedAuthors []PaperV3GetDiversePapersResponseRecommendationContextFollowedAuthor `json:"followed_authors"`
	FollowedLikers  []PaperV3GetDiversePapersResponseRecommendationContextFollowedLiker  `json:"followed_likers"`
	Hot             bool                                                                 `json:"hot"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FollowedAuthors respjson.Field
		FollowedLikers  respjson.Field
		Hot             respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3GetDiversePapersResponseRecommendationContext) RawJSON() string { return r.JSON.raw }
func (r *PaperV3GetDiversePapersResponseRecommendationContext) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3GetDiversePapersResponseRecommendationContextFollowedAuthor struct {
	Name string `json:"name" api:"required"`
	Slug string `json:"slug"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Slug        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3GetDiversePapersResponseRecommendationContextFollowedAuthor) RawJSON() string {
	return r.JSON.raw
}
func (r *PaperV3GetDiversePapersResponseRecommendationContextFollowedAuthor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3GetDiversePapersResponseRecommendationContextFollowedLiker struct {
	ID               string                                                                    `json:"id" api:"required" format:"uuid"`
	Avatar           []PaperV3GetDiversePapersResponseRecommendationContextFollowedLikerAvatar `json:"avatar" api:"required"`
	GoogleScholarID  string                                                                    `json:"googleScholarId" api:"required"`
	Institution      string                                                                    `json:"institution" api:"required"`
	RealName         string                                                                    `json:"realName" api:"required"`
	Reputation       float64                                                                   `json:"reputation" api:"required"`
	ResearcherSlug   string                                                                    `json:"researcherSlug" api:"required"`
	Username         string                                                                    `json:"username" api:"required"`
	WeeklyReputation float64                                                                   `json:"weeklyReputation" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		Avatar           respjson.Field
		GoogleScholarID  respjson.Field
		Institution      respjson.Field
		RealName         respjson.Field
		Reputation       respjson.Field
		ResearcherSlug   respjson.Field
		Username         respjson.Field
		WeeklyReputation respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3GetDiversePapersResponseRecommendationContextFollowedLiker) RawJSON() string {
	return r.JSON.raw
}
func (r *PaperV3GetDiversePapersResponseRecommendationContextFollowedLiker) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3GetDiversePapersResponseRecommendationContextFollowedLikerAvatar struct {
	// Any of "full_size", "thumbnail".
	Type string `json:"type" api:"required"`
	URL  string `json:"url" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3GetDiversePapersResponseRecommendationContextFollowedLikerAvatar) RawJSON() string {
	return r.JSON.raw
}
func (r *PaperV3GetDiversePapersResponseRecommendationContextFollowedLikerAvatar) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3GetFeedResponse struct {
	Page          float64                       `json:"page" api:"required"`
	Papers        []PaperV3GetFeedResponsePaper `json:"papers" api:"required"`
	FeedCursor    string                        `json:"feedCursor"`
	FeedRefreshed bool                          `json:"feedRefreshed"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Page          respjson.Field
		Papers        respjson.Field
		FeedCursor    respjson.Field
		FeedRefreshed respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3GetFeedResponse) RawJSON() string { return r.JSON.raw }
func (r *PaperV3GetFeedResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3GetFeedResponsePaper struct {
	ID         string                                  `json:"id" api:"required" format:"uuid"`
	Abstract   string                                  `json:"abstract" api:"required"`
	AuthorInfo []PaperV3GetFeedResponsePaperAuthorInfo `json:"author_info" api:"required"`
	Authors    []string                                `json:"authors" api:"required"`
	// A versioned paper ID (e.g. 1706.03762v1)
	CanonicalID          string                                        `json:"canonical_id" api:"required"`
	ExternalBlog         PaperV3GetFeedResponsePaperExternalBlog       `json:"external_blog" api:"required"`
	FirstPublicationDate string                                        `json:"first_publication_date" api:"required"`
	FullAuthors          []PaperV3GetFeedResponsePaperFullAuthor       `json:"full_authors" api:"required"`
	FullAuthorsV2        []PaperV3GetFeedResponsePaperFullAuthorsV2    `json:"full_authors_v2" api:"required"`
	GitHubStars          float64                                       `json:"github_stars" api:"required"`
	GitHubURL            string                                        `json:"github_url" api:"required"`
	HasRunReport         bool                                          `json:"has_run_report" api:"required"`
	ImageURL             string                                        `json:"image_url" api:"required"`
	Metrics              PaperV3GetFeedResponsePaperMetrics            `json:"metrics" api:"required"`
	OrganizationInfo     []PaperV3GetFeedResponsePaperOrganizationInfo `json:"organization_info" api:"required"`
	PaperGroupID         string                                        `json:"paper_group_id" api:"required" format:"uuid"`
	PaperSummary         PaperV3GetFeedResponsePaperPaperSummary       `json:"paper_summary" api:"required"`
	PdfOnly              bool                                          `json:"pdf_only" api:"required"`
	PublicationDate      string                                        `json:"publication_date" api:"required"`
	Title                string                                        `json:"title" api:"required"`
	Topics               []string                                      `json:"topics" api:"required"`
	// A versionless universal paper ID (e.g. 1706.03762)
	UniversalPaperID      string                                           `json:"universal_paper_id" api:"required"`
	UpdatedAt             string                                           `json:"updated_at" api:"required"`
	VersionID             string                                           `json:"version_id" api:"required" format:"uuid"`
	CardPreviewBlobID     string                                           `json:"card_preview_blob_id" api:"nullable" format:"uuid"`
	NarrationAudioURL     string                                           `json:"narration_audio_url" api:"nullable"`
	RecommendationContext PaperV3GetFeedResponsePaperRecommendationContext `json:"recommendation_context"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                    respjson.Field
		Abstract              respjson.Field
		AuthorInfo            respjson.Field
		Authors               respjson.Field
		CanonicalID           respjson.Field
		ExternalBlog          respjson.Field
		FirstPublicationDate  respjson.Field
		FullAuthors           respjson.Field
		FullAuthorsV2         respjson.Field
		GitHubStars           respjson.Field
		GitHubURL             respjson.Field
		HasRunReport          respjson.Field
		ImageURL              respjson.Field
		Metrics               respjson.Field
		OrganizationInfo      respjson.Field
		PaperGroupID          respjson.Field
		PaperSummary          respjson.Field
		PdfOnly               respjson.Field
		PublicationDate       respjson.Field
		Title                 respjson.Field
		Topics                respjson.Field
		UniversalPaperID      respjson.Field
		UpdatedAt             respjson.Field
		VersionID             respjson.Field
		CardPreviewBlobID     respjson.Field
		NarrationAudioURL     respjson.Field
		RecommendationContext respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3GetFeedResponsePaper) RawJSON() string { return r.JSON.raw }
func (r *PaperV3GetFeedResponsePaper) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3GetFeedResponsePaperAuthorInfo struct {
	ID               string                                        `json:"id" api:"required" format:"uuid"`
	Avatar           []PaperV3GetFeedResponsePaperAuthorInfoAvatar `json:"avatar" api:"required"`
	BlueskyUsername  string                                        `json:"blueskyUsername" api:"required"`
	GitHubUsername   string                                        `json:"githubUsername" api:"required"`
	GoogleScholarID  string                                        `json:"googleScholarId" api:"required"`
	Institution      string                                        `json:"institution" api:"required"`
	LinkedinUsername string                                        `json:"linkedinUsername" api:"required"`
	OrcidID          string                                        `json:"orcidId" api:"required"`
	PublicEmail      string                                        `json:"publicEmail" api:"required"`
	RealName         string                                        `json:"realName" api:"required"`
	Reputation       float64                                       `json:"reputation" api:"required"`
	ResearcherSlug   string                                        `json:"researcherSlug" api:"required"`
	// Any of "user", "reviewer", "admin", "bot".
	Role             string  `json:"role" api:"required"`
	Username         string  `json:"username" api:"required"`
	Verified         bool    `json:"verified" api:"required"`
	WeeklyReputation float64 `json:"weeklyReputation" api:"required"`
	XUsername        string  `json:"xUsername" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		Avatar           respjson.Field
		BlueskyUsername  respjson.Field
		GitHubUsername   respjson.Field
		GoogleScholarID  respjson.Field
		Institution      respjson.Field
		LinkedinUsername respjson.Field
		OrcidID          respjson.Field
		PublicEmail      respjson.Field
		RealName         respjson.Field
		Reputation       respjson.Field
		ResearcherSlug   respjson.Field
		Role             respjson.Field
		Username         respjson.Field
		Verified         respjson.Field
		WeeklyReputation respjson.Field
		XUsername        respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3GetFeedResponsePaperAuthorInfo) RawJSON() string { return r.JSON.raw }
func (r *PaperV3GetFeedResponsePaperAuthorInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3GetFeedResponsePaperAuthorInfoAvatar struct {
	// Any of "full_size", "thumbnail".
	Type string `json:"type" api:"required"`
	URL  string `json:"url" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3GetFeedResponsePaperAuthorInfoAvatar) RawJSON() string { return r.JSON.raw }
func (r *PaperV3GetFeedResponsePaperAuthorInfoAvatar) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3GetFeedResponsePaperExternalBlog struct {
	BodyBlobID  string `json:"body_blob_id" api:"required" format:"uuid"`
	CoverBlobID string `json:"cover_blob_id" api:"required" format:"uuid"`
	SourceName  string `json:"source_name" api:"required"`
	SourceURL   string `json:"source_url" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BodyBlobID  respjson.Field
		CoverBlobID respjson.Field
		SourceName  respjson.Field
		SourceURL   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3GetFeedResponsePaperExternalBlog) RawJSON() string { return r.JSON.raw }
func (r *PaperV3GetFeedResponsePaperExternalBlog) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3GetFeedResponsePaperFullAuthor struct {
	ID             string `json:"id" api:"required" format:"uuid"`
	FullName       string `json:"full_name" api:"required"`
	UserID         string `json:"user_id" api:"required" format:"uuid"`
	Username       string `json:"username" api:"required"`
	ResearcherSlug string `json:"researcher_slug" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		FullName       respjson.Field
		UserID         respjson.Field
		Username       respjson.Field
		ResearcherSlug respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3GetFeedResponsePaperFullAuthor) RawJSON() string { return r.JSON.raw }
func (r *PaperV3GetFeedResponsePaperFullAuthor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3GetFeedResponsePaperFullAuthorsV2 struct {
	FullName   string                                             `json:"full_name" api:"required"`
	Researcher PaperV3GetFeedResponsePaperFullAuthorsV2Researcher `json:"researcher" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FullName    respjson.Field
		Researcher  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3GetFeedResponsePaperFullAuthorsV2) RawJSON() string { return r.JSON.raw }
func (r *PaperV3GetFeedResponsePaperFullAuthorsV2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3GetFeedResponsePaperFullAuthorsV2Researcher struct {
	Affiliation   string                                                        `json:"affiliation" api:"required"`
	Bio           string                                                        `json:"bio" api:"required"`
	Citations     float64                                                       `json:"citations" api:"required"`
	Headline      string                                                        `json:"headline" api:"required"`
	HIndex        float64                                                       `json:"hIndex" api:"required"`
	LinkedUser    PaperV3GetFeedResponsePaperFullAuthorsV2ResearcherLinkedUser  `json:"linkedUser" api:"required"`
	Links         PaperV3GetFeedResponsePaperFullAuthorsV2ResearcherLinks       `json:"links" api:"required"`
	Name          string                                                        `json:"name" api:"required"`
	PhotoURL      string                                                        `json:"photoUrl" api:"required"`
	ResearchAreas []string                                                      `json:"researchAreas" api:"required"`
	Slug          string                                                        `json:"slug" api:"required"`
	Reason        PaperV3GetFeedResponsePaperFullAuthorsV2ResearcherReasonUnion `json:"reason"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Affiliation   respjson.Field
		Bio           respjson.Field
		Citations     respjson.Field
		Headline      respjson.Field
		HIndex        respjson.Field
		LinkedUser    respjson.Field
		Links         respjson.Field
		Name          respjson.Field
		PhotoURL      respjson.Field
		ResearchAreas respjson.Field
		Slug          respjson.Field
		Reason        respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3GetFeedResponsePaperFullAuthorsV2Researcher) RawJSON() string { return r.JSON.raw }
func (r *PaperV3GetFeedResponsePaperFullAuthorsV2Researcher) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3GetFeedResponsePaperFullAuthorsV2ResearcherLinkedUser struct {
	Name     string `json:"name" api:"required"`
	Username string `json:"username" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Username    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3GetFeedResponsePaperFullAuthorsV2ResearcherLinkedUser) RawJSON() string {
	return r.JSON.raw
}
func (r *PaperV3GetFeedResponsePaperFullAuthorsV2ResearcherLinkedUser) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3GetFeedResponsePaperFullAuthorsV2ResearcherLinks struct {
	Bluesky      string `json:"bluesky" api:"required"`
	Cv           string `json:"cv" api:"required"`
	Dblp         string `json:"dblp" api:"required"`
	Email        string `json:"email" api:"required"`
	GitHub       string `json:"github" api:"required"`
	Huggingface  string `json:"huggingface" api:"required"`
	Linkedin     string `json:"linkedin" api:"required"`
	Openreview   string `json:"openreview" api:"required"`
	Orcid        string `json:"orcid" api:"required"`
	PersonalSite string `json:"personalSite" api:"required"`
	Scholar      string `json:"scholar" api:"required"`
	Twitter      string `json:"twitter" api:"required"`
	Wikipedia    string `json:"wikipedia" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Bluesky      respjson.Field
		Cv           respjson.Field
		Dblp         respjson.Field
		Email        respjson.Field
		GitHub       respjson.Field
		Huggingface  respjson.Field
		Linkedin     respjson.Field
		Openreview   respjson.Field
		Orcid        respjson.Field
		PersonalSite respjson.Field
		Scholar      respjson.Field
		Twitter      respjson.Field
		Wikipedia    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3GetFeedResponsePaperFullAuthorsV2ResearcherLinks) RawJSON() string { return r.JSON.raw }
func (r *PaperV3GetFeedResponsePaperFullAuthorsV2ResearcherLinks) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PaperV3GetFeedResponsePaperFullAuthorsV2ResearcherReasonUnion contains all
// possible properties and values from
// [PaperV3GetFeedResponsePaperFullAuthorsV2ResearcherReasonObject],
// [PaperV3GetFeedResponsePaperFullAuthorsV2ResearcherReasonKind],
// [PaperV3GetFeedResponsePaperFullAuthorsV2ResearcherReasonObject2].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type PaperV3GetFeedResponsePaperFullAuthorsV2ResearcherReasonUnion struct {
	Kind string `json:"kind"`
	// This field is from variant
	// [PaperV3GetFeedResponsePaperFullAuthorsV2ResearcherReasonObject].
	PaperTitle string `json:"paperTitle"`
	// This field is from variant
	// [PaperV3GetFeedResponsePaperFullAuthorsV2ResearcherReasonObject2].
	Count float64 `json:"count"`
	// This field is from variant
	// [PaperV3GetFeedResponsePaperFullAuthorsV2ResearcherReasonObject2].
	Followed PaperV3GetFeedResponsePaperFullAuthorsV2ResearcherReasonObject2Followed `json:"followed"`
	JSON     struct {
		Kind       respjson.Field
		PaperTitle respjson.Field
		Count      respjson.Field
		Followed   respjson.Field
		raw        string
	} `json:"-"`
}

func (u PaperV3GetFeedResponsePaperFullAuthorsV2ResearcherReasonUnion) AsPaperV3GetFeedResponsePaperFullAuthorsV2ResearcherReasonObject() (v PaperV3GetFeedResponsePaperFullAuthorsV2ResearcherReasonObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PaperV3GetFeedResponsePaperFullAuthorsV2ResearcherReasonUnion) AsPaperV3GetFeedResponsePaperFullAuthorsV2ResearcherReasonKind() (v PaperV3GetFeedResponsePaperFullAuthorsV2ResearcherReasonKind) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PaperV3GetFeedResponsePaperFullAuthorsV2ResearcherReasonUnion) AsPaperV3GetFeedResponsePaperFullAuthorsV2ResearcherReasonObject2() (v PaperV3GetFeedResponsePaperFullAuthorsV2ResearcherReasonObject2) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PaperV3GetFeedResponsePaperFullAuthorsV2ResearcherReasonUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *PaperV3GetFeedResponsePaperFullAuthorsV2ResearcherReasonUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3GetFeedResponsePaperFullAuthorsV2ResearcherReasonObject struct {
	// Any of "interest".
	Kind       string `json:"kind" api:"required"`
	PaperTitle string `json:"paperTitle"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Kind        respjson.Field
		PaperTitle  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3GetFeedResponsePaperFullAuthorsV2ResearcherReasonObject) RawJSON() string {
	return r.JSON.raw
}
func (r *PaperV3GetFeedResponsePaperFullAuthorsV2ResearcherReasonObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3GetFeedResponsePaperFullAuthorsV2ResearcherReasonKind struct {
	// Any of "read".
	Kind string `json:"kind" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Kind        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3GetFeedResponsePaperFullAuthorsV2ResearcherReasonKind) RawJSON() string {
	return r.JSON.raw
}
func (r *PaperV3GetFeedResponsePaperFullAuthorsV2ResearcherReasonKind) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3GetFeedResponsePaperFullAuthorsV2ResearcherReasonObject2 struct {
	Count float64 `json:"count" api:"required"`
	// Any of "coauthor".
	Kind     string                                                                  `json:"kind" api:"required"`
	Followed PaperV3GetFeedResponsePaperFullAuthorsV2ResearcherReasonObject2Followed `json:"followed"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Kind        respjson.Field
		Followed    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3GetFeedResponsePaperFullAuthorsV2ResearcherReasonObject2) RawJSON() string {
	return r.JSON.raw
}
func (r *PaperV3GetFeedResponsePaperFullAuthorsV2ResearcherReasonObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3GetFeedResponsePaperFullAuthorsV2ResearcherReasonObject2Followed struct {
	Name string `json:"name" api:"required"`
	Slug string `json:"slug" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Slug        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3GetFeedResponsePaperFullAuthorsV2ResearcherReasonObject2Followed) RawJSON() string {
	return r.JSON.raw
}
func (r *PaperV3GetFeedResponsePaperFullAuthorsV2ResearcherReasonObject2Followed) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3GetFeedResponsePaperMetrics struct {
	PublicTotalVotes float64                                       `json:"public_total_votes" api:"required"`
	TotalVotes       float64                                       `json:"total_votes" api:"required"`
	VisitsCount      PaperV3GetFeedResponsePaperMetricsVisitsCount `json:"visits_count" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PublicTotalVotes respjson.Field
		TotalVotes       respjson.Field
		VisitsCount      respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3GetFeedResponsePaperMetrics) RawJSON() string { return r.JSON.raw }
func (r *PaperV3GetFeedResponsePaperMetrics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3GetFeedResponsePaperMetricsVisitsCount struct {
	All       float64 `json:"all" api:"required"`
	Last7Days float64 `json:"last_7_days" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		All         respjson.Field
		Last7Days   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3GetFeedResponsePaperMetricsVisitsCount) RawJSON() string { return r.JSON.raw }
func (r *PaperV3GetFeedResponsePaperMetricsVisitsCount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3GetFeedResponsePaperOrganizationInfo struct {
	Image string `json:"image" api:"required"`
	Name  string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Image       respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3GetFeedResponsePaperOrganizationInfo) RawJSON() string { return r.JSON.raw }
func (r *PaperV3GetFeedResponsePaperOrganizationInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3GetFeedResponsePaperPaperSummary struct {
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
func (r PaperV3GetFeedResponsePaperPaperSummary) RawJSON() string { return r.JSON.raw }
func (r *PaperV3GetFeedResponsePaperPaperSummary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3GetFeedResponsePaperRecommendationContext struct {
	FollowedAuthors []PaperV3GetFeedResponsePaperRecommendationContextFollowedAuthor `json:"followed_authors"`
	FollowedLikers  []PaperV3GetFeedResponsePaperRecommendationContextFollowedLiker  `json:"followed_likers"`
	Hot             bool                                                             `json:"hot"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FollowedAuthors respjson.Field
		FollowedLikers  respjson.Field
		Hot             respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3GetFeedResponsePaperRecommendationContext) RawJSON() string { return r.JSON.raw }
func (r *PaperV3GetFeedResponsePaperRecommendationContext) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3GetFeedResponsePaperRecommendationContextFollowedAuthor struct {
	Name string `json:"name" api:"required"`
	Slug string `json:"slug"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Slug        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3GetFeedResponsePaperRecommendationContextFollowedAuthor) RawJSON() string {
	return r.JSON.raw
}
func (r *PaperV3GetFeedResponsePaperRecommendationContextFollowedAuthor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3GetFeedResponsePaperRecommendationContextFollowedLiker struct {
	ID               string                                                                `json:"id" api:"required" format:"uuid"`
	Avatar           []PaperV3GetFeedResponsePaperRecommendationContextFollowedLikerAvatar `json:"avatar" api:"required"`
	GoogleScholarID  string                                                                `json:"googleScholarId" api:"required"`
	Institution      string                                                                `json:"institution" api:"required"`
	RealName         string                                                                `json:"realName" api:"required"`
	Reputation       float64                                                               `json:"reputation" api:"required"`
	ResearcherSlug   string                                                                `json:"researcherSlug" api:"required"`
	Username         string                                                                `json:"username" api:"required"`
	WeeklyReputation float64                                                               `json:"weeklyReputation" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		Avatar           respjson.Field
		GoogleScholarID  respjson.Field
		Institution      respjson.Field
		RealName         respjson.Field
		Reputation       respjson.Field
		ResearcherSlug   respjson.Field
		Username         respjson.Field
		WeeklyReputation respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3GetFeedResponsePaperRecommendationContextFollowedLiker) RawJSON() string {
	return r.JSON.raw
}
func (r *PaperV3GetFeedResponsePaperRecommendationContextFollowedLiker) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3GetFeedResponsePaperRecommendationContextFollowedLikerAvatar struct {
	// Any of "full_size", "thumbnail".
	Type string `json:"type" api:"required"`
	URL  string `json:"url" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3GetFeedResponsePaperRecommendationContextFollowedLikerAvatar) RawJSON() string {
	return r.JSON.raw
}
func (r *PaperV3GetFeedResponsePaperRecommendationContextFollowedLikerAvatar) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3GetFiguresResponse struct {
	Figures []string `json:"figures" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Figures     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3GetFiguresResponse) RawJSON() string { return r.JSON.raw }
func (r *PaperV3GetFiguresResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3GetFullTextResponse struct {
	Pages []PaperV3GetFullTextResponsePage `json:"pages" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Pages       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3GetFullTextResponse) RawJSON() string { return r.JSON.raw }
func (r *PaperV3GetFullTextResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3GetFullTextResponsePage struct {
	PageNumber float64 `json:"pageNumber" api:"required"`
	Text       string  `json:"text" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PageNumber  respjson.Field
		Text        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3GetFullTextResponsePage) RawJSON() string { return r.JSON.raw }
func (r *PaperV3GetFullTextResponsePage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3GetMetricsResponse struct {
	CommentsCount    float64 `json:"commentsCount" api:"required"`
	PublicTotalVotes float64 `json:"publicTotalVotes" api:"required"`
	VisitsAll        float64 `json:"visitsAll" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CommentsCount    respjson.Field
		PublicTotalVotes respjson.Field
		VisitsAll        respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3GetMetricsResponse) RawJSON() string { return r.JSON.raw }
func (r *PaperV3GetMetricsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3GetPreviewResponse struct {
	ID         string                                `json:"id" api:"required" format:"uuid"`
	Abstract   string                                `json:"abstract" api:"required"`
	AuthorInfo []PaperV3GetPreviewResponseAuthorInfo `json:"author_info" api:"required"`
	Authors    []string                              `json:"authors" api:"required"`
	// A versioned paper ID (e.g. 1706.03762v1)
	CanonicalID          string                                      `json:"canonical_id" api:"required"`
	ExternalBlog         PaperV3GetPreviewResponseExternalBlog       `json:"external_blog" api:"required"`
	FirstPublicationDate string                                      `json:"first_publication_date" api:"required"`
	FullAuthors          []PaperV3GetPreviewResponseFullAuthor       `json:"full_authors" api:"required"`
	FullAuthorsV2        []PaperV3GetPreviewResponseFullAuthorsV2    `json:"full_authors_v2" api:"required"`
	GitHubStars          float64                                     `json:"github_stars" api:"required"`
	GitHubURL            string                                      `json:"github_url" api:"required"`
	HasRunReport         bool                                        `json:"has_run_report" api:"required"`
	ImageURL             string                                      `json:"image_url" api:"required"`
	Metrics              PaperV3GetPreviewResponseMetrics            `json:"metrics" api:"required"`
	OrganizationInfo     []PaperV3GetPreviewResponseOrganizationInfo `json:"organization_info" api:"required"`
	PaperGroupID         string                                      `json:"paper_group_id" api:"required" format:"uuid"`
	PaperSummary         PaperV3GetPreviewResponsePaperSummary       `json:"paper_summary" api:"required"`
	PdfOnly              bool                                        `json:"pdf_only" api:"required"`
	PublicationDate      string                                      `json:"publication_date" api:"required"`
	Title                string                                      `json:"title" api:"required"`
	Topics               []string                                    `json:"topics" api:"required"`
	// A versionless universal paper ID (e.g. 1706.03762)
	UniversalPaperID      string                                         `json:"universal_paper_id" api:"required"`
	UpdatedAt             string                                         `json:"updated_at" api:"required"`
	VersionID             string                                         `json:"version_id" api:"required" format:"uuid"`
	CardPreviewBlobID     string                                         `json:"card_preview_blob_id" api:"nullable" format:"uuid"`
	NarrationAudioURL     string                                         `json:"narration_audio_url" api:"nullable"`
	RecommendationContext PaperV3GetPreviewResponseRecommendationContext `json:"recommendation_context"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                    respjson.Field
		Abstract              respjson.Field
		AuthorInfo            respjson.Field
		Authors               respjson.Field
		CanonicalID           respjson.Field
		ExternalBlog          respjson.Field
		FirstPublicationDate  respjson.Field
		FullAuthors           respjson.Field
		FullAuthorsV2         respjson.Field
		GitHubStars           respjson.Field
		GitHubURL             respjson.Field
		HasRunReport          respjson.Field
		ImageURL              respjson.Field
		Metrics               respjson.Field
		OrganizationInfo      respjson.Field
		PaperGroupID          respjson.Field
		PaperSummary          respjson.Field
		PdfOnly               respjson.Field
		PublicationDate       respjson.Field
		Title                 respjson.Field
		Topics                respjson.Field
		UniversalPaperID      respjson.Field
		UpdatedAt             respjson.Field
		VersionID             respjson.Field
		CardPreviewBlobID     respjson.Field
		NarrationAudioURL     respjson.Field
		RecommendationContext respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3GetPreviewResponse) RawJSON() string { return r.JSON.raw }
func (r *PaperV3GetPreviewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3GetPreviewResponseAuthorInfo struct {
	ID               string                                      `json:"id" api:"required" format:"uuid"`
	Avatar           []PaperV3GetPreviewResponseAuthorInfoAvatar `json:"avatar" api:"required"`
	BlueskyUsername  string                                      `json:"blueskyUsername" api:"required"`
	GitHubUsername   string                                      `json:"githubUsername" api:"required"`
	GoogleScholarID  string                                      `json:"googleScholarId" api:"required"`
	Institution      string                                      `json:"institution" api:"required"`
	LinkedinUsername string                                      `json:"linkedinUsername" api:"required"`
	OrcidID          string                                      `json:"orcidId" api:"required"`
	PublicEmail      string                                      `json:"publicEmail" api:"required"`
	RealName         string                                      `json:"realName" api:"required"`
	Reputation       float64                                     `json:"reputation" api:"required"`
	ResearcherSlug   string                                      `json:"researcherSlug" api:"required"`
	// Any of "user", "reviewer", "admin", "bot".
	Role             string  `json:"role" api:"required"`
	Username         string  `json:"username" api:"required"`
	Verified         bool    `json:"verified" api:"required"`
	WeeklyReputation float64 `json:"weeklyReputation" api:"required"`
	XUsername        string  `json:"xUsername" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		Avatar           respjson.Field
		BlueskyUsername  respjson.Field
		GitHubUsername   respjson.Field
		GoogleScholarID  respjson.Field
		Institution      respjson.Field
		LinkedinUsername respjson.Field
		OrcidID          respjson.Field
		PublicEmail      respjson.Field
		RealName         respjson.Field
		Reputation       respjson.Field
		ResearcherSlug   respjson.Field
		Role             respjson.Field
		Username         respjson.Field
		Verified         respjson.Field
		WeeklyReputation respjson.Field
		XUsername        respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3GetPreviewResponseAuthorInfo) RawJSON() string { return r.JSON.raw }
func (r *PaperV3GetPreviewResponseAuthorInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3GetPreviewResponseAuthorInfoAvatar struct {
	// Any of "full_size", "thumbnail".
	Type string `json:"type" api:"required"`
	URL  string `json:"url" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3GetPreviewResponseAuthorInfoAvatar) RawJSON() string { return r.JSON.raw }
func (r *PaperV3GetPreviewResponseAuthorInfoAvatar) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3GetPreviewResponseExternalBlog struct {
	BodyBlobID  string `json:"body_blob_id" api:"required" format:"uuid"`
	CoverBlobID string `json:"cover_blob_id" api:"required" format:"uuid"`
	SourceName  string `json:"source_name" api:"required"`
	SourceURL   string `json:"source_url" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BodyBlobID  respjson.Field
		CoverBlobID respjson.Field
		SourceName  respjson.Field
		SourceURL   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3GetPreviewResponseExternalBlog) RawJSON() string { return r.JSON.raw }
func (r *PaperV3GetPreviewResponseExternalBlog) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3GetPreviewResponseFullAuthor struct {
	ID             string `json:"id" api:"required" format:"uuid"`
	FullName       string `json:"full_name" api:"required"`
	UserID         string `json:"user_id" api:"required" format:"uuid"`
	Username       string `json:"username" api:"required"`
	ResearcherSlug string `json:"researcher_slug" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		FullName       respjson.Field
		UserID         respjson.Field
		Username       respjson.Field
		ResearcherSlug respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3GetPreviewResponseFullAuthor) RawJSON() string { return r.JSON.raw }
func (r *PaperV3GetPreviewResponseFullAuthor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3GetPreviewResponseFullAuthorsV2 struct {
	FullName   string                                           `json:"full_name" api:"required"`
	Researcher PaperV3GetPreviewResponseFullAuthorsV2Researcher `json:"researcher" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FullName    respjson.Field
		Researcher  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3GetPreviewResponseFullAuthorsV2) RawJSON() string { return r.JSON.raw }
func (r *PaperV3GetPreviewResponseFullAuthorsV2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3GetPreviewResponseFullAuthorsV2Researcher struct {
	Affiliation   string                                                      `json:"affiliation" api:"required"`
	Bio           string                                                      `json:"bio" api:"required"`
	Citations     float64                                                     `json:"citations" api:"required"`
	Headline      string                                                      `json:"headline" api:"required"`
	HIndex        float64                                                     `json:"hIndex" api:"required"`
	LinkedUser    PaperV3GetPreviewResponseFullAuthorsV2ResearcherLinkedUser  `json:"linkedUser" api:"required"`
	Links         PaperV3GetPreviewResponseFullAuthorsV2ResearcherLinks       `json:"links" api:"required"`
	Name          string                                                      `json:"name" api:"required"`
	PhotoURL      string                                                      `json:"photoUrl" api:"required"`
	ResearchAreas []string                                                    `json:"researchAreas" api:"required"`
	Slug          string                                                      `json:"slug" api:"required"`
	Reason        PaperV3GetPreviewResponseFullAuthorsV2ResearcherReasonUnion `json:"reason"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Affiliation   respjson.Field
		Bio           respjson.Field
		Citations     respjson.Field
		Headline      respjson.Field
		HIndex        respjson.Field
		LinkedUser    respjson.Field
		Links         respjson.Field
		Name          respjson.Field
		PhotoURL      respjson.Field
		ResearchAreas respjson.Field
		Slug          respjson.Field
		Reason        respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3GetPreviewResponseFullAuthorsV2Researcher) RawJSON() string { return r.JSON.raw }
func (r *PaperV3GetPreviewResponseFullAuthorsV2Researcher) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3GetPreviewResponseFullAuthorsV2ResearcherLinkedUser struct {
	Name     string `json:"name" api:"required"`
	Username string `json:"username" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Username    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3GetPreviewResponseFullAuthorsV2ResearcherLinkedUser) RawJSON() string {
	return r.JSON.raw
}
func (r *PaperV3GetPreviewResponseFullAuthorsV2ResearcherLinkedUser) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3GetPreviewResponseFullAuthorsV2ResearcherLinks struct {
	Bluesky      string `json:"bluesky" api:"required"`
	Cv           string `json:"cv" api:"required"`
	Dblp         string `json:"dblp" api:"required"`
	Email        string `json:"email" api:"required"`
	GitHub       string `json:"github" api:"required"`
	Huggingface  string `json:"huggingface" api:"required"`
	Linkedin     string `json:"linkedin" api:"required"`
	Openreview   string `json:"openreview" api:"required"`
	Orcid        string `json:"orcid" api:"required"`
	PersonalSite string `json:"personalSite" api:"required"`
	Scholar      string `json:"scholar" api:"required"`
	Twitter      string `json:"twitter" api:"required"`
	Wikipedia    string `json:"wikipedia" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Bluesky      respjson.Field
		Cv           respjson.Field
		Dblp         respjson.Field
		Email        respjson.Field
		GitHub       respjson.Field
		Huggingface  respjson.Field
		Linkedin     respjson.Field
		Openreview   respjson.Field
		Orcid        respjson.Field
		PersonalSite respjson.Field
		Scholar      respjson.Field
		Twitter      respjson.Field
		Wikipedia    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3GetPreviewResponseFullAuthorsV2ResearcherLinks) RawJSON() string { return r.JSON.raw }
func (r *PaperV3GetPreviewResponseFullAuthorsV2ResearcherLinks) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PaperV3GetPreviewResponseFullAuthorsV2ResearcherReasonUnion contains all
// possible properties and values from
// [PaperV3GetPreviewResponseFullAuthorsV2ResearcherReasonObject],
// [PaperV3GetPreviewResponseFullAuthorsV2ResearcherReasonKind],
// [PaperV3GetPreviewResponseFullAuthorsV2ResearcherReasonObject2].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type PaperV3GetPreviewResponseFullAuthorsV2ResearcherReasonUnion struct {
	Kind string `json:"kind"`
	// This field is from variant
	// [PaperV3GetPreviewResponseFullAuthorsV2ResearcherReasonObject].
	PaperTitle string `json:"paperTitle"`
	// This field is from variant
	// [PaperV3GetPreviewResponseFullAuthorsV2ResearcherReasonObject2].
	Count float64 `json:"count"`
	// This field is from variant
	// [PaperV3GetPreviewResponseFullAuthorsV2ResearcherReasonObject2].
	Followed PaperV3GetPreviewResponseFullAuthorsV2ResearcherReasonObject2Followed `json:"followed"`
	JSON     struct {
		Kind       respjson.Field
		PaperTitle respjson.Field
		Count      respjson.Field
		Followed   respjson.Field
		raw        string
	} `json:"-"`
}

func (u PaperV3GetPreviewResponseFullAuthorsV2ResearcherReasonUnion) AsPaperV3GetPreviewResponseFullAuthorsV2ResearcherReasonObject() (v PaperV3GetPreviewResponseFullAuthorsV2ResearcherReasonObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PaperV3GetPreviewResponseFullAuthorsV2ResearcherReasonUnion) AsPaperV3GetPreviewResponseFullAuthorsV2ResearcherReasonKind() (v PaperV3GetPreviewResponseFullAuthorsV2ResearcherReasonKind) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PaperV3GetPreviewResponseFullAuthorsV2ResearcherReasonUnion) AsPaperV3GetPreviewResponseFullAuthorsV2ResearcherReasonObject2() (v PaperV3GetPreviewResponseFullAuthorsV2ResearcherReasonObject2) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PaperV3GetPreviewResponseFullAuthorsV2ResearcherReasonUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *PaperV3GetPreviewResponseFullAuthorsV2ResearcherReasonUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3GetPreviewResponseFullAuthorsV2ResearcherReasonObject struct {
	// Any of "interest".
	Kind       string `json:"kind" api:"required"`
	PaperTitle string `json:"paperTitle"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Kind        respjson.Field
		PaperTitle  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3GetPreviewResponseFullAuthorsV2ResearcherReasonObject) RawJSON() string {
	return r.JSON.raw
}
func (r *PaperV3GetPreviewResponseFullAuthorsV2ResearcherReasonObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3GetPreviewResponseFullAuthorsV2ResearcherReasonKind struct {
	// Any of "read".
	Kind string `json:"kind" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Kind        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3GetPreviewResponseFullAuthorsV2ResearcherReasonKind) RawJSON() string {
	return r.JSON.raw
}
func (r *PaperV3GetPreviewResponseFullAuthorsV2ResearcherReasonKind) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3GetPreviewResponseFullAuthorsV2ResearcherReasonObject2 struct {
	Count float64 `json:"count" api:"required"`
	// Any of "coauthor".
	Kind     string                                                                `json:"kind" api:"required"`
	Followed PaperV3GetPreviewResponseFullAuthorsV2ResearcherReasonObject2Followed `json:"followed"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Kind        respjson.Field
		Followed    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3GetPreviewResponseFullAuthorsV2ResearcherReasonObject2) RawJSON() string {
	return r.JSON.raw
}
func (r *PaperV3GetPreviewResponseFullAuthorsV2ResearcherReasonObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3GetPreviewResponseFullAuthorsV2ResearcherReasonObject2Followed struct {
	Name string `json:"name" api:"required"`
	Slug string `json:"slug" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Slug        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3GetPreviewResponseFullAuthorsV2ResearcherReasonObject2Followed) RawJSON() string {
	return r.JSON.raw
}
func (r *PaperV3GetPreviewResponseFullAuthorsV2ResearcherReasonObject2Followed) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3GetPreviewResponseMetrics struct {
	PublicTotalVotes float64                                     `json:"public_total_votes" api:"required"`
	TotalVotes       float64                                     `json:"total_votes" api:"required"`
	VisitsCount      PaperV3GetPreviewResponseMetricsVisitsCount `json:"visits_count" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PublicTotalVotes respjson.Field
		TotalVotes       respjson.Field
		VisitsCount      respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3GetPreviewResponseMetrics) RawJSON() string { return r.JSON.raw }
func (r *PaperV3GetPreviewResponseMetrics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3GetPreviewResponseMetricsVisitsCount struct {
	All       float64 `json:"all" api:"required"`
	Last7Days float64 `json:"last_7_days" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		All         respjson.Field
		Last7Days   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3GetPreviewResponseMetricsVisitsCount) RawJSON() string { return r.JSON.raw }
func (r *PaperV3GetPreviewResponseMetricsVisitsCount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3GetPreviewResponseOrganizationInfo struct {
	Image string `json:"image" api:"required"`
	Name  string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Image       respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3GetPreviewResponseOrganizationInfo) RawJSON() string { return r.JSON.raw }
func (r *PaperV3GetPreviewResponseOrganizationInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3GetPreviewResponsePaperSummary struct {
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
func (r PaperV3GetPreviewResponsePaperSummary) RawJSON() string { return r.JSON.raw }
func (r *PaperV3GetPreviewResponsePaperSummary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3GetPreviewResponseRecommendationContext struct {
	FollowedAuthors []PaperV3GetPreviewResponseRecommendationContextFollowedAuthor `json:"followed_authors"`
	FollowedLikers  []PaperV3GetPreviewResponseRecommendationContextFollowedLiker  `json:"followed_likers"`
	Hot             bool                                                           `json:"hot"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FollowedAuthors respjson.Field
		FollowedLikers  respjson.Field
		Hot             respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3GetPreviewResponseRecommendationContext) RawJSON() string { return r.JSON.raw }
func (r *PaperV3GetPreviewResponseRecommendationContext) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3GetPreviewResponseRecommendationContextFollowedAuthor struct {
	Name string `json:"name" api:"required"`
	Slug string `json:"slug"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Slug        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3GetPreviewResponseRecommendationContextFollowedAuthor) RawJSON() string {
	return r.JSON.raw
}
func (r *PaperV3GetPreviewResponseRecommendationContextFollowedAuthor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3GetPreviewResponseRecommendationContextFollowedLiker struct {
	ID               string                                                              `json:"id" api:"required" format:"uuid"`
	Avatar           []PaperV3GetPreviewResponseRecommendationContextFollowedLikerAvatar `json:"avatar" api:"required"`
	GoogleScholarID  string                                                              `json:"googleScholarId" api:"required"`
	Institution      string                                                              `json:"institution" api:"required"`
	RealName         string                                                              `json:"realName" api:"required"`
	Reputation       float64                                                             `json:"reputation" api:"required"`
	ResearcherSlug   string                                                              `json:"researcherSlug" api:"required"`
	Username         string                                                              `json:"username" api:"required"`
	WeeklyReputation float64                                                             `json:"weeklyReputation" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		Avatar           respjson.Field
		GoogleScholarID  respjson.Field
		Institution      respjson.Field
		RealName         respjson.Field
		Reputation       respjson.Field
		ResearcherSlug   respjson.Field
		Username         respjson.Field
		WeeklyReputation respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3GetPreviewResponseRecommendationContextFollowedLiker) RawJSON() string {
	return r.JSON.raw
}
func (r *PaperV3GetPreviewResponseRecommendationContextFollowedLiker) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3GetPreviewResponseRecommendationContextFollowedLikerAvatar struct {
	// Any of "full_size", "thumbnail".
	Type string `json:"type" api:"required"`
	URL  string `json:"url" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3GetPreviewResponseRecommendationContextFollowedLikerAvatar) RawJSON() string {
	return r.JSON.raw
}
func (r *PaperV3GetPreviewResponseRecommendationContextFollowedLikerAvatar) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3GetSimilarPapersResponse struct {
	ID         string                                      `json:"id" api:"required" format:"uuid"`
	Abstract   string                                      `json:"abstract" api:"required"`
	AuthorInfo []PaperV3GetSimilarPapersResponseAuthorInfo `json:"author_info" api:"required"`
	Authors    []string                                    `json:"authors" api:"required"`
	// A versioned paper ID (e.g. 1706.03762v1)
	CanonicalID          string                                            `json:"canonical_id" api:"required"`
	ExternalBlog         PaperV3GetSimilarPapersResponseExternalBlog       `json:"external_blog" api:"required"`
	FirstPublicationDate string                                            `json:"first_publication_date" api:"required"`
	FullAuthors          []PaperV3GetSimilarPapersResponseFullAuthor       `json:"full_authors" api:"required"`
	FullAuthorsV2        []PaperV3GetSimilarPapersResponseFullAuthorsV2    `json:"full_authors_v2" api:"required"`
	GitHubStars          float64                                           `json:"github_stars" api:"required"`
	GitHubURL            string                                            `json:"github_url" api:"required"`
	HasRunReport         bool                                              `json:"has_run_report" api:"required"`
	ImageURL             string                                            `json:"image_url" api:"required"`
	Metrics              PaperV3GetSimilarPapersResponseMetrics            `json:"metrics" api:"required"`
	OrganizationInfo     []PaperV3GetSimilarPapersResponseOrganizationInfo `json:"organization_info" api:"required"`
	PaperGroupID         string                                            `json:"paper_group_id" api:"required" format:"uuid"`
	PaperSummary         PaperV3GetSimilarPapersResponsePaperSummary       `json:"paper_summary" api:"required"`
	PdfOnly              bool                                              `json:"pdf_only" api:"required"`
	PublicationDate      string                                            `json:"publication_date" api:"required"`
	Title                string                                            `json:"title" api:"required"`
	Topics               []string                                          `json:"topics" api:"required"`
	// A versionless universal paper ID (e.g. 1706.03762)
	UniversalPaperID      string                                               `json:"universal_paper_id" api:"required"`
	UpdatedAt             string                                               `json:"updated_at" api:"required"`
	VersionID             string                                               `json:"version_id" api:"required" format:"uuid"`
	CardPreviewBlobID     string                                               `json:"card_preview_blob_id" api:"nullable" format:"uuid"`
	NarrationAudioURL     string                                               `json:"narration_audio_url" api:"nullable"`
	RecommendationContext PaperV3GetSimilarPapersResponseRecommendationContext `json:"recommendation_context"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                    respjson.Field
		Abstract              respjson.Field
		AuthorInfo            respjson.Field
		Authors               respjson.Field
		CanonicalID           respjson.Field
		ExternalBlog          respjson.Field
		FirstPublicationDate  respjson.Field
		FullAuthors           respjson.Field
		FullAuthorsV2         respjson.Field
		GitHubStars           respjson.Field
		GitHubURL             respjson.Field
		HasRunReport          respjson.Field
		ImageURL              respjson.Field
		Metrics               respjson.Field
		OrganizationInfo      respjson.Field
		PaperGroupID          respjson.Field
		PaperSummary          respjson.Field
		PdfOnly               respjson.Field
		PublicationDate       respjson.Field
		Title                 respjson.Field
		Topics                respjson.Field
		UniversalPaperID      respjson.Field
		UpdatedAt             respjson.Field
		VersionID             respjson.Field
		CardPreviewBlobID     respjson.Field
		NarrationAudioURL     respjson.Field
		RecommendationContext respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3GetSimilarPapersResponse) RawJSON() string { return r.JSON.raw }
func (r *PaperV3GetSimilarPapersResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3GetSimilarPapersResponseAuthorInfo struct {
	ID               string                                            `json:"id" api:"required" format:"uuid"`
	Avatar           []PaperV3GetSimilarPapersResponseAuthorInfoAvatar `json:"avatar" api:"required"`
	BlueskyUsername  string                                            `json:"blueskyUsername" api:"required"`
	GitHubUsername   string                                            `json:"githubUsername" api:"required"`
	GoogleScholarID  string                                            `json:"googleScholarId" api:"required"`
	Institution      string                                            `json:"institution" api:"required"`
	LinkedinUsername string                                            `json:"linkedinUsername" api:"required"`
	OrcidID          string                                            `json:"orcidId" api:"required"`
	PublicEmail      string                                            `json:"publicEmail" api:"required"`
	RealName         string                                            `json:"realName" api:"required"`
	Reputation       float64                                           `json:"reputation" api:"required"`
	ResearcherSlug   string                                            `json:"researcherSlug" api:"required"`
	// Any of "user", "reviewer", "admin", "bot".
	Role             string  `json:"role" api:"required"`
	Username         string  `json:"username" api:"required"`
	Verified         bool    `json:"verified" api:"required"`
	WeeklyReputation float64 `json:"weeklyReputation" api:"required"`
	XUsername        string  `json:"xUsername" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		Avatar           respjson.Field
		BlueskyUsername  respjson.Field
		GitHubUsername   respjson.Field
		GoogleScholarID  respjson.Field
		Institution      respjson.Field
		LinkedinUsername respjson.Field
		OrcidID          respjson.Field
		PublicEmail      respjson.Field
		RealName         respjson.Field
		Reputation       respjson.Field
		ResearcherSlug   respjson.Field
		Role             respjson.Field
		Username         respjson.Field
		Verified         respjson.Field
		WeeklyReputation respjson.Field
		XUsername        respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3GetSimilarPapersResponseAuthorInfo) RawJSON() string { return r.JSON.raw }
func (r *PaperV3GetSimilarPapersResponseAuthorInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3GetSimilarPapersResponseAuthorInfoAvatar struct {
	// Any of "full_size", "thumbnail".
	Type string `json:"type" api:"required"`
	URL  string `json:"url" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3GetSimilarPapersResponseAuthorInfoAvatar) RawJSON() string { return r.JSON.raw }
func (r *PaperV3GetSimilarPapersResponseAuthorInfoAvatar) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3GetSimilarPapersResponseExternalBlog struct {
	BodyBlobID  string `json:"body_blob_id" api:"required" format:"uuid"`
	CoverBlobID string `json:"cover_blob_id" api:"required" format:"uuid"`
	SourceName  string `json:"source_name" api:"required"`
	SourceURL   string `json:"source_url" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BodyBlobID  respjson.Field
		CoverBlobID respjson.Field
		SourceName  respjson.Field
		SourceURL   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3GetSimilarPapersResponseExternalBlog) RawJSON() string { return r.JSON.raw }
func (r *PaperV3GetSimilarPapersResponseExternalBlog) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3GetSimilarPapersResponseFullAuthor struct {
	ID             string `json:"id" api:"required" format:"uuid"`
	FullName       string `json:"full_name" api:"required"`
	UserID         string `json:"user_id" api:"required" format:"uuid"`
	Username       string `json:"username" api:"required"`
	ResearcherSlug string `json:"researcher_slug" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		FullName       respjson.Field
		UserID         respjson.Field
		Username       respjson.Field
		ResearcherSlug respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3GetSimilarPapersResponseFullAuthor) RawJSON() string { return r.JSON.raw }
func (r *PaperV3GetSimilarPapersResponseFullAuthor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3GetSimilarPapersResponseFullAuthorsV2 struct {
	FullName   string                                                 `json:"full_name" api:"required"`
	Researcher PaperV3GetSimilarPapersResponseFullAuthorsV2Researcher `json:"researcher" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FullName    respjson.Field
		Researcher  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3GetSimilarPapersResponseFullAuthorsV2) RawJSON() string { return r.JSON.raw }
func (r *PaperV3GetSimilarPapersResponseFullAuthorsV2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3GetSimilarPapersResponseFullAuthorsV2Researcher struct {
	Affiliation   string                                                            `json:"affiliation" api:"required"`
	Bio           string                                                            `json:"bio" api:"required"`
	Citations     float64                                                           `json:"citations" api:"required"`
	Headline      string                                                            `json:"headline" api:"required"`
	HIndex        float64                                                           `json:"hIndex" api:"required"`
	LinkedUser    PaperV3GetSimilarPapersResponseFullAuthorsV2ResearcherLinkedUser  `json:"linkedUser" api:"required"`
	Links         PaperV3GetSimilarPapersResponseFullAuthorsV2ResearcherLinks       `json:"links" api:"required"`
	Name          string                                                            `json:"name" api:"required"`
	PhotoURL      string                                                            `json:"photoUrl" api:"required"`
	ResearchAreas []string                                                          `json:"researchAreas" api:"required"`
	Slug          string                                                            `json:"slug" api:"required"`
	Reason        PaperV3GetSimilarPapersResponseFullAuthorsV2ResearcherReasonUnion `json:"reason"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Affiliation   respjson.Field
		Bio           respjson.Field
		Citations     respjson.Field
		Headline      respjson.Field
		HIndex        respjson.Field
		LinkedUser    respjson.Field
		Links         respjson.Field
		Name          respjson.Field
		PhotoURL      respjson.Field
		ResearchAreas respjson.Field
		Slug          respjson.Field
		Reason        respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3GetSimilarPapersResponseFullAuthorsV2Researcher) RawJSON() string { return r.JSON.raw }
func (r *PaperV3GetSimilarPapersResponseFullAuthorsV2Researcher) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3GetSimilarPapersResponseFullAuthorsV2ResearcherLinkedUser struct {
	Name     string `json:"name" api:"required"`
	Username string `json:"username" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Username    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3GetSimilarPapersResponseFullAuthorsV2ResearcherLinkedUser) RawJSON() string {
	return r.JSON.raw
}
func (r *PaperV3GetSimilarPapersResponseFullAuthorsV2ResearcherLinkedUser) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3GetSimilarPapersResponseFullAuthorsV2ResearcherLinks struct {
	Bluesky      string `json:"bluesky" api:"required"`
	Cv           string `json:"cv" api:"required"`
	Dblp         string `json:"dblp" api:"required"`
	Email        string `json:"email" api:"required"`
	GitHub       string `json:"github" api:"required"`
	Huggingface  string `json:"huggingface" api:"required"`
	Linkedin     string `json:"linkedin" api:"required"`
	Openreview   string `json:"openreview" api:"required"`
	Orcid        string `json:"orcid" api:"required"`
	PersonalSite string `json:"personalSite" api:"required"`
	Scholar      string `json:"scholar" api:"required"`
	Twitter      string `json:"twitter" api:"required"`
	Wikipedia    string `json:"wikipedia" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Bluesky      respjson.Field
		Cv           respjson.Field
		Dblp         respjson.Field
		Email        respjson.Field
		GitHub       respjson.Field
		Huggingface  respjson.Field
		Linkedin     respjson.Field
		Openreview   respjson.Field
		Orcid        respjson.Field
		PersonalSite respjson.Field
		Scholar      respjson.Field
		Twitter      respjson.Field
		Wikipedia    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3GetSimilarPapersResponseFullAuthorsV2ResearcherLinks) RawJSON() string {
	return r.JSON.raw
}
func (r *PaperV3GetSimilarPapersResponseFullAuthorsV2ResearcherLinks) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PaperV3GetSimilarPapersResponseFullAuthorsV2ResearcherReasonUnion contains all
// possible properties and values from
// [PaperV3GetSimilarPapersResponseFullAuthorsV2ResearcherReasonObject],
// [PaperV3GetSimilarPapersResponseFullAuthorsV2ResearcherReasonKind],
// [PaperV3GetSimilarPapersResponseFullAuthorsV2ResearcherReasonObject2].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type PaperV3GetSimilarPapersResponseFullAuthorsV2ResearcherReasonUnion struct {
	Kind string `json:"kind"`
	// This field is from variant
	// [PaperV3GetSimilarPapersResponseFullAuthorsV2ResearcherReasonObject].
	PaperTitle string `json:"paperTitle"`
	// This field is from variant
	// [PaperV3GetSimilarPapersResponseFullAuthorsV2ResearcherReasonObject2].
	Count float64 `json:"count"`
	// This field is from variant
	// [PaperV3GetSimilarPapersResponseFullAuthorsV2ResearcherReasonObject2].
	Followed PaperV3GetSimilarPapersResponseFullAuthorsV2ResearcherReasonObject2Followed `json:"followed"`
	JSON     struct {
		Kind       respjson.Field
		PaperTitle respjson.Field
		Count      respjson.Field
		Followed   respjson.Field
		raw        string
	} `json:"-"`
}

func (u PaperV3GetSimilarPapersResponseFullAuthorsV2ResearcherReasonUnion) AsPaperV3GetSimilarPapersResponseFullAuthorsV2ResearcherReasonObject() (v PaperV3GetSimilarPapersResponseFullAuthorsV2ResearcherReasonObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PaperV3GetSimilarPapersResponseFullAuthorsV2ResearcherReasonUnion) AsPaperV3GetSimilarPapersResponseFullAuthorsV2ResearcherReasonKind() (v PaperV3GetSimilarPapersResponseFullAuthorsV2ResearcherReasonKind) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PaperV3GetSimilarPapersResponseFullAuthorsV2ResearcherReasonUnion) AsPaperV3GetSimilarPapersResponseFullAuthorsV2ResearcherReasonObject2() (v PaperV3GetSimilarPapersResponseFullAuthorsV2ResearcherReasonObject2) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PaperV3GetSimilarPapersResponseFullAuthorsV2ResearcherReasonUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *PaperV3GetSimilarPapersResponseFullAuthorsV2ResearcherReasonUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3GetSimilarPapersResponseFullAuthorsV2ResearcherReasonObject struct {
	// Any of "interest".
	Kind       string `json:"kind" api:"required"`
	PaperTitle string `json:"paperTitle"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Kind        respjson.Field
		PaperTitle  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3GetSimilarPapersResponseFullAuthorsV2ResearcherReasonObject) RawJSON() string {
	return r.JSON.raw
}
func (r *PaperV3GetSimilarPapersResponseFullAuthorsV2ResearcherReasonObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3GetSimilarPapersResponseFullAuthorsV2ResearcherReasonKind struct {
	// Any of "read".
	Kind string `json:"kind" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Kind        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3GetSimilarPapersResponseFullAuthorsV2ResearcherReasonKind) RawJSON() string {
	return r.JSON.raw
}
func (r *PaperV3GetSimilarPapersResponseFullAuthorsV2ResearcherReasonKind) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3GetSimilarPapersResponseFullAuthorsV2ResearcherReasonObject2 struct {
	Count float64 `json:"count" api:"required"`
	// Any of "coauthor".
	Kind     string                                                                      `json:"kind" api:"required"`
	Followed PaperV3GetSimilarPapersResponseFullAuthorsV2ResearcherReasonObject2Followed `json:"followed"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Kind        respjson.Field
		Followed    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3GetSimilarPapersResponseFullAuthorsV2ResearcherReasonObject2) RawJSON() string {
	return r.JSON.raw
}
func (r *PaperV3GetSimilarPapersResponseFullAuthorsV2ResearcherReasonObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3GetSimilarPapersResponseFullAuthorsV2ResearcherReasonObject2Followed struct {
	Name string `json:"name" api:"required"`
	Slug string `json:"slug" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Slug        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3GetSimilarPapersResponseFullAuthorsV2ResearcherReasonObject2Followed) RawJSON() string {
	return r.JSON.raw
}
func (r *PaperV3GetSimilarPapersResponseFullAuthorsV2ResearcherReasonObject2Followed) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3GetSimilarPapersResponseMetrics struct {
	PublicTotalVotes float64                                           `json:"public_total_votes" api:"required"`
	TotalVotes       float64                                           `json:"total_votes" api:"required"`
	VisitsCount      PaperV3GetSimilarPapersResponseMetricsVisitsCount `json:"visits_count" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PublicTotalVotes respjson.Field
		TotalVotes       respjson.Field
		VisitsCount      respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3GetSimilarPapersResponseMetrics) RawJSON() string { return r.JSON.raw }
func (r *PaperV3GetSimilarPapersResponseMetrics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3GetSimilarPapersResponseMetricsVisitsCount struct {
	All       float64 `json:"all" api:"required"`
	Last7Days float64 `json:"last_7_days" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		All         respjson.Field
		Last7Days   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3GetSimilarPapersResponseMetricsVisitsCount) RawJSON() string { return r.JSON.raw }
func (r *PaperV3GetSimilarPapersResponseMetricsVisitsCount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3GetSimilarPapersResponseOrganizationInfo struct {
	Image string `json:"image" api:"required"`
	Name  string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Image       respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3GetSimilarPapersResponseOrganizationInfo) RawJSON() string { return r.JSON.raw }
func (r *PaperV3GetSimilarPapersResponseOrganizationInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3GetSimilarPapersResponsePaperSummary struct {
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
func (r PaperV3GetSimilarPapersResponsePaperSummary) RawJSON() string { return r.JSON.raw }
func (r *PaperV3GetSimilarPapersResponsePaperSummary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3GetSimilarPapersResponseRecommendationContext struct {
	FollowedAuthors []PaperV3GetSimilarPapersResponseRecommendationContextFollowedAuthor `json:"followed_authors"`
	FollowedLikers  []PaperV3GetSimilarPapersResponseRecommendationContextFollowedLiker  `json:"followed_likers"`
	Hot             bool                                                                 `json:"hot"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FollowedAuthors respjson.Field
		FollowedLikers  respjson.Field
		Hot             respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3GetSimilarPapersResponseRecommendationContext) RawJSON() string { return r.JSON.raw }
func (r *PaperV3GetSimilarPapersResponseRecommendationContext) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3GetSimilarPapersResponseRecommendationContextFollowedAuthor struct {
	Name string `json:"name" api:"required"`
	Slug string `json:"slug"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Slug        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3GetSimilarPapersResponseRecommendationContextFollowedAuthor) RawJSON() string {
	return r.JSON.raw
}
func (r *PaperV3GetSimilarPapersResponseRecommendationContextFollowedAuthor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3GetSimilarPapersResponseRecommendationContextFollowedLiker struct {
	ID               string                                                                    `json:"id" api:"required" format:"uuid"`
	Avatar           []PaperV3GetSimilarPapersResponseRecommendationContextFollowedLikerAvatar `json:"avatar" api:"required"`
	GoogleScholarID  string                                                                    `json:"googleScholarId" api:"required"`
	Institution      string                                                                    `json:"institution" api:"required"`
	RealName         string                                                                    `json:"realName" api:"required"`
	Reputation       float64                                                                   `json:"reputation" api:"required"`
	ResearcherSlug   string                                                                    `json:"researcherSlug" api:"required"`
	Username         string                                                                    `json:"username" api:"required"`
	WeeklyReputation float64                                                                   `json:"weeklyReputation" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		Avatar           respjson.Field
		GoogleScholarID  respjson.Field
		Institution      respjson.Field
		RealName         respjson.Field
		Reputation       respjson.Field
		ResearcherSlug   respjson.Field
		Username         respjson.Field
		WeeklyReputation respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3GetSimilarPapersResponseRecommendationContextFollowedLiker) RawJSON() string {
	return r.JSON.raw
}
func (r *PaperV3GetSimilarPapersResponseRecommendationContextFollowedLiker) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3GetSimilarPapersResponseRecommendationContextFollowedLikerAvatar struct {
	// Any of "full_size", "thumbnail".
	Type string `json:"type" api:"required"`
	URL  string `json:"url" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3GetSimilarPapersResponseRecommendationContextFollowedLikerAvatar) RawJSON() string {
	return r.JSON.raw
}
func (r *PaperV3GetSimilarPapersResponseRecommendationContextFollowedLikerAvatar) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3GetUnrelatedResponse struct {
	ID         string                                  `json:"id" api:"required" format:"uuid"`
	Abstract   string                                  `json:"abstract" api:"required"`
	AuthorInfo []PaperV3GetUnrelatedResponseAuthorInfo `json:"author_info" api:"required"`
	Authors    []string                                `json:"authors" api:"required"`
	// A versioned paper ID (e.g. 1706.03762v1)
	CanonicalID          string                                        `json:"canonical_id" api:"required"`
	ExternalBlog         PaperV3GetUnrelatedResponseExternalBlog       `json:"external_blog" api:"required"`
	FirstPublicationDate string                                        `json:"first_publication_date" api:"required"`
	FullAuthors          []PaperV3GetUnrelatedResponseFullAuthor       `json:"full_authors" api:"required"`
	FullAuthorsV2        []PaperV3GetUnrelatedResponseFullAuthorsV2    `json:"full_authors_v2" api:"required"`
	GitHubStars          float64                                       `json:"github_stars" api:"required"`
	GitHubURL            string                                        `json:"github_url" api:"required"`
	HasRunReport         bool                                          `json:"has_run_report" api:"required"`
	ImageURL             string                                        `json:"image_url" api:"required"`
	Metrics              PaperV3GetUnrelatedResponseMetrics            `json:"metrics" api:"required"`
	OrganizationInfo     []PaperV3GetUnrelatedResponseOrganizationInfo `json:"organization_info" api:"required"`
	PaperGroupID         string                                        `json:"paper_group_id" api:"required" format:"uuid"`
	PaperSummary         PaperV3GetUnrelatedResponsePaperSummary       `json:"paper_summary" api:"required"`
	PdfOnly              bool                                          `json:"pdf_only" api:"required"`
	PublicationDate      string                                        `json:"publication_date" api:"required"`
	Title                string                                        `json:"title" api:"required"`
	Topics               []string                                      `json:"topics" api:"required"`
	// A versionless universal paper ID (e.g. 1706.03762)
	UniversalPaperID      string                                           `json:"universal_paper_id" api:"required"`
	UpdatedAt             string                                           `json:"updated_at" api:"required"`
	VersionID             string                                           `json:"version_id" api:"required" format:"uuid"`
	CardPreviewBlobID     string                                           `json:"card_preview_blob_id" api:"nullable" format:"uuid"`
	NarrationAudioURL     string                                           `json:"narration_audio_url" api:"nullable"`
	RecommendationContext PaperV3GetUnrelatedResponseRecommendationContext `json:"recommendation_context"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                    respjson.Field
		Abstract              respjson.Field
		AuthorInfo            respjson.Field
		Authors               respjson.Field
		CanonicalID           respjson.Field
		ExternalBlog          respjson.Field
		FirstPublicationDate  respjson.Field
		FullAuthors           respjson.Field
		FullAuthorsV2         respjson.Field
		GitHubStars           respjson.Field
		GitHubURL             respjson.Field
		HasRunReport          respjson.Field
		ImageURL              respjson.Field
		Metrics               respjson.Field
		OrganizationInfo      respjson.Field
		PaperGroupID          respjson.Field
		PaperSummary          respjson.Field
		PdfOnly               respjson.Field
		PublicationDate       respjson.Field
		Title                 respjson.Field
		Topics                respjson.Field
		UniversalPaperID      respjson.Field
		UpdatedAt             respjson.Field
		VersionID             respjson.Field
		CardPreviewBlobID     respjson.Field
		NarrationAudioURL     respjson.Field
		RecommendationContext respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3GetUnrelatedResponse) RawJSON() string { return r.JSON.raw }
func (r *PaperV3GetUnrelatedResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3GetUnrelatedResponseAuthorInfo struct {
	ID               string                                        `json:"id" api:"required" format:"uuid"`
	Avatar           []PaperV3GetUnrelatedResponseAuthorInfoAvatar `json:"avatar" api:"required"`
	BlueskyUsername  string                                        `json:"blueskyUsername" api:"required"`
	GitHubUsername   string                                        `json:"githubUsername" api:"required"`
	GoogleScholarID  string                                        `json:"googleScholarId" api:"required"`
	Institution      string                                        `json:"institution" api:"required"`
	LinkedinUsername string                                        `json:"linkedinUsername" api:"required"`
	OrcidID          string                                        `json:"orcidId" api:"required"`
	PublicEmail      string                                        `json:"publicEmail" api:"required"`
	RealName         string                                        `json:"realName" api:"required"`
	Reputation       float64                                       `json:"reputation" api:"required"`
	ResearcherSlug   string                                        `json:"researcherSlug" api:"required"`
	// Any of "user", "reviewer", "admin", "bot".
	Role             string  `json:"role" api:"required"`
	Username         string  `json:"username" api:"required"`
	Verified         bool    `json:"verified" api:"required"`
	WeeklyReputation float64 `json:"weeklyReputation" api:"required"`
	XUsername        string  `json:"xUsername" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		Avatar           respjson.Field
		BlueskyUsername  respjson.Field
		GitHubUsername   respjson.Field
		GoogleScholarID  respjson.Field
		Institution      respjson.Field
		LinkedinUsername respjson.Field
		OrcidID          respjson.Field
		PublicEmail      respjson.Field
		RealName         respjson.Field
		Reputation       respjson.Field
		ResearcherSlug   respjson.Field
		Role             respjson.Field
		Username         respjson.Field
		Verified         respjson.Field
		WeeklyReputation respjson.Field
		XUsername        respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3GetUnrelatedResponseAuthorInfo) RawJSON() string { return r.JSON.raw }
func (r *PaperV3GetUnrelatedResponseAuthorInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3GetUnrelatedResponseAuthorInfoAvatar struct {
	// Any of "full_size", "thumbnail".
	Type string `json:"type" api:"required"`
	URL  string `json:"url" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3GetUnrelatedResponseAuthorInfoAvatar) RawJSON() string { return r.JSON.raw }
func (r *PaperV3GetUnrelatedResponseAuthorInfoAvatar) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3GetUnrelatedResponseExternalBlog struct {
	BodyBlobID  string `json:"body_blob_id" api:"required" format:"uuid"`
	CoverBlobID string `json:"cover_blob_id" api:"required" format:"uuid"`
	SourceName  string `json:"source_name" api:"required"`
	SourceURL   string `json:"source_url" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BodyBlobID  respjson.Field
		CoverBlobID respjson.Field
		SourceName  respjson.Field
		SourceURL   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3GetUnrelatedResponseExternalBlog) RawJSON() string { return r.JSON.raw }
func (r *PaperV3GetUnrelatedResponseExternalBlog) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3GetUnrelatedResponseFullAuthor struct {
	ID             string `json:"id" api:"required" format:"uuid"`
	FullName       string `json:"full_name" api:"required"`
	UserID         string `json:"user_id" api:"required" format:"uuid"`
	Username       string `json:"username" api:"required"`
	ResearcherSlug string `json:"researcher_slug" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		FullName       respjson.Field
		UserID         respjson.Field
		Username       respjson.Field
		ResearcherSlug respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3GetUnrelatedResponseFullAuthor) RawJSON() string { return r.JSON.raw }
func (r *PaperV3GetUnrelatedResponseFullAuthor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3GetUnrelatedResponseFullAuthorsV2 struct {
	FullName   string                                             `json:"full_name" api:"required"`
	Researcher PaperV3GetUnrelatedResponseFullAuthorsV2Researcher `json:"researcher" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FullName    respjson.Field
		Researcher  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3GetUnrelatedResponseFullAuthorsV2) RawJSON() string { return r.JSON.raw }
func (r *PaperV3GetUnrelatedResponseFullAuthorsV2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3GetUnrelatedResponseFullAuthorsV2Researcher struct {
	Affiliation   string                                                        `json:"affiliation" api:"required"`
	Bio           string                                                        `json:"bio" api:"required"`
	Citations     float64                                                       `json:"citations" api:"required"`
	Headline      string                                                        `json:"headline" api:"required"`
	HIndex        float64                                                       `json:"hIndex" api:"required"`
	LinkedUser    PaperV3GetUnrelatedResponseFullAuthorsV2ResearcherLinkedUser  `json:"linkedUser" api:"required"`
	Links         PaperV3GetUnrelatedResponseFullAuthorsV2ResearcherLinks       `json:"links" api:"required"`
	Name          string                                                        `json:"name" api:"required"`
	PhotoURL      string                                                        `json:"photoUrl" api:"required"`
	ResearchAreas []string                                                      `json:"researchAreas" api:"required"`
	Slug          string                                                        `json:"slug" api:"required"`
	Reason        PaperV3GetUnrelatedResponseFullAuthorsV2ResearcherReasonUnion `json:"reason"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Affiliation   respjson.Field
		Bio           respjson.Field
		Citations     respjson.Field
		Headline      respjson.Field
		HIndex        respjson.Field
		LinkedUser    respjson.Field
		Links         respjson.Field
		Name          respjson.Field
		PhotoURL      respjson.Field
		ResearchAreas respjson.Field
		Slug          respjson.Field
		Reason        respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3GetUnrelatedResponseFullAuthorsV2Researcher) RawJSON() string { return r.JSON.raw }
func (r *PaperV3GetUnrelatedResponseFullAuthorsV2Researcher) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3GetUnrelatedResponseFullAuthorsV2ResearcherLinkedUser struct {
	Name     string `json:"name" api:"required"`
	Username string `json:"username" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Username    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3GetUnrelatedResponseFullAuthorsV2ResearcherLinkedUser) RawJSON() string {
	return r.JSON.raw
}
func (r *PaperV3GetUnrelatedResponseFullAuthorsV2ResearcherLinkedUser) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3GetUnrelatedResponseFullAuthorsV2ResearcherLinks struct {
	Bluesky      string `json:"bluesky" api:"required"`
	Cv           string `json:"cv" api:"required"`
	Dblp         string `json:"dblp" api:"required"`
	Email        string `json:"email" api:"required"`
	GitHub       string `json:"github" api:"required"`
	Huggingface  string `json:"huggingface" api:"required"`
	Linkedin     string `json:"linkedin" api:"required"`
	Openreview   string `json:"openreview" api:"required"`
	Orcid        string `json:"orcid" api:"required"`
	PersonalSite string `json:"personalSite" api:"required"`
	Scholar      string `json:"scholar" api:"required"`
	Twitter      string `json:"twitter" api:"required"`
	Wikipedia    string `json:"wikipedia" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Bluesky      respjson.Field
		Cv           respjson.Field
		Dblp         respjson.Field
		Email        respjson.Field
		GitHub       respjson.Field
		Huggingface  respjson.Field
		Linkedin     respjson.Field
		Openreview   respjson.Field
		Orcid        respjson.Field
		PersonalSite respjson.Field
		Scholar      respjson.Field
		Twitter      respjson.Field
		Wikipedia    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3GetUnrelatedResponseFullAuthorsV2ResearcherLinks) RawJSON() string { return r.JSON.raw }
func (r *PaperV3GetUnrelatedResponseFullAuthorsV2ResearcherLinks) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PaperV3GetUnrelatedResponseFullAuthorsV2ResearcherReasonUnion contains all
// possible properties and values from
// [PaperV3GetUnrelatedResponseFullAuthorsV2ResearcherReasonObject],
// [PaperV3GetUnrelatedResponseFullAuthorsV2ResearcherReasonKind],
// [PaperV3GetUnrelatedResponseFullAuthorsV2ResearcherReasonObject2].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type PaperV3GetUnrelatedResponseFullAuthorsV2ResearcherReasonUnion struct {
	Kind string `json:"kind"`
	// This field is from variant
	// [PaperV3GetUnrelatedResponseFullAuthorsV2ResearcherReasonObject].
	PaperTitle string `json:"paperTitle"`
	// This field is from variant
	// [PaperV3GetUnrelatedResponseFullAuthorsV2ResearcherReasonObject2].
	Count float64 `json:"count"`
	// This field is from variant
	// [PaperV3GetUnrelatedResponseFullAuthorsV2ResearcherReasonObject2].
	Followed PaperV3GetUnrelatedResponseFullAuthorsV2ResearcherReasonObject2Followed `json:"followed"`
	JSON     struct {
		Kind       respjson.Field
		PaperTitle respjson.Field
		Count      respjson.Field
		Followed   respjson.Field
		raw        string
	} `json:"-"`
}

func (u PaperV3GetUnrelatedResponseFullAuthorsV2ResearcherReasonUnion) AsPaperV3GetUnrelatedResponseFullAuthorsV2ResearcherReasonObject() (v PaperV3GetUnrelatedResponseFullAuthorsV2ResearcherReasonObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PaperV3GetUnrelatedResponseFullAuthorsV2ResearcherReasonUnion) AsPaperV3GetUnrelatedResponseFullAuthorsV2ResearcherReasonKind() (v PaperV3GetUnrelatedResponseFullAuthorsV2ResearcherReasonKind) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PaperV3GetUnrelatedResponseFullAuthorsV2ResearcherReasonUnion) AsPaperV3GetUnrelatedResponseFullAuthorsV2ResearcherReasonObject2() (v PaperV3GetUnrelatedResponseFullAuthorsV2ResearcherReasonObject2) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PaperV3GetUnrelatedResponseFullAuthorsV2ResearcherReasonUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *PaperV3GetUnrelatedResponseFullAuthorsV2ResearcherReasonUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3GetUnrelatedResponseFullAuthorsV2ResearcherReasonObject struct {
	// Any of "interest".
	Kind       string `json:"kind" api:"required"`
	PaperTitle string `json:"paperTitle"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Kind        respjson.Field
		PaperTitle  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3GetUnrelatedResponseFullAuthorsV2ResearcherReasonObject) RawJSON() string {
	return r.JSON.raw
}
func (r *PaperV3GetUnrelatedResponseFullAuthorsV2ResearcherReasonObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3GetUnrelatedResponseFullAuthorsV2ResearcherReasonKind struct {
	// Any of "read".
	Kind string `json:"kind" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Kind        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3GetUnrelatedResponseFullAuthorsV2ResearcherReasonKind) RawJSON() string {
	return r.JSON.raw
}
func (r *PaperV3GetUnrelatedResponseFullAuthorsV2ResearcherReasonKind) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3GetUnrelatedResponseFullAuthorsV2ResearcherReasonObject2 struct {
	Count float64 `json:"count" api:"required"`
	// Any of "coauthor".
	Kind     string                                                                  `json:"kind" api:"required"`
	Followed PaperV3GetUnrelatedResponseFullAuthorsV2ResearcherReasonObject2Followed `json:"followed"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Kind        respjson.Field
		Followed    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3GetUnrelatedResponseFullAuthorsV2ResearcherReasonObject2) RawJSON() string {
	return r.JSON.raw
}
func (r *PaperV3GetUnrelatedResponseFullAuthorsV2ResearcherReasonObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3GetUnrelatedResponseFullAuthorsV2ResearcherReasonObject2Followed struct {
	Name string `json:"name" api:"required"`
	Slug string `json:"slug" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Slug        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3GetUnrelatedResponseFullAuthorsV2ResearcherReasonObject2Followed) RawJSON() string {
	return r.JSON.raw
}
func (r *PaperV3GetUnrelatedResponseFullAuthorsV2ResearcherReasonObject2Followed) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3GetUnrelatedResponseMetrics struct {
	PublicTotalVotes float64                                       `json:"public_total_votes" api:"required"`
	TotalVotes       float64                                       `json:"total_votes" api:"required"`
	VisitsCount      PaperV3GetUnrelatedResponseMetricsVisitsCount `json:"visits_count" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PublicTotalVotes respjson.Field
		TotalVotes       respjson.Field
		VisitsCount      respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3GetUnrelatedResponseMetrics) RawJSON() string { return r.JSON.raw }
func (r *PaperV3GetUnrelatedResponseMetrics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3GetUnrelatedResponseMetricsVisitsCount struct {
	All       float64 `json:"all" api:"required"`
	Last7Days float64 `json:"last_7_days" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		All         respjson.Field
		Last7Days   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3GetUnrelatedResponseMetricsVisitsCount) RawJSON() string { return r.JSON.raw }
func (r *PaperV3GetUnrelatedResponseMetricsVisitsCount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3GetUnrelatedResponseOrganizationInfo struct {
	Image string `json:"image" api:"required"`
	Name  string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Image       respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3GetUnrelatedResponseOrganizationInfo) RawJSON() string { return r.JSON.raw }
func (r *PaperV3GetUnrelatedResponseOrganizationInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3GetUnrelatedResponsePaperSummary struct {
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
func (r PaperV3GetUnrelatedResponsePaperSummary) RawJSON() string { return r.JSON.raw }
func (r *PaperV3GetUnrelatedResponsePaperSummary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3GetUnrelatedResponseRecommendationContext struct {
	FollowedAuthors []PaperV3GetUnrelatedResponseRecommendationContextFollowedAuthor `json:"followed_authors"`
	FollowedLikers  []PaperV3GetUnrelatedResponseRecommendationContextFollowedLiker  `json:"followed_likers"`
	Hot             bool                                                             `json:"hot"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FollowedAuthors respjson.Field
		FollowedLikers  respjson.Field
		Hot             respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3GetUnrelatedResponseRecommendationContext) RawJSON() string { return r.JSON.raw }
func (r *PaperV3GetUnrelatedResponseRecommendationContext) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3GetUnrelatedResponseRecommendationContextFollowedAuthor struct {
	Name string `json:"name" api:"required"`
	Slug string `json:"slug"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Slug        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3GetUnrelatedResponseRecommendationContextFollowedAuthor) RawJSON() string {
	return r.JSON.raw
}
func (r *PaperV3GetUnrelatedResponseRecommendationContextFollowedAuthor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3GetUnrelatedResponseRecommendationContextFollowedLiker struct {
	ID               string                                                                `json:"id" api:"required" format:"uuid"`
	Avatar           []PaperV3GetUnrelatedResponseRecommendationContextFollowedLikerAvatar `json:"avatar" api:"required"`
	GoogleScholarID  string                                                                `json:"googleScholarId" api:"required"`
	Institution      string                                                                `json:"institution" api:"required"`
	RealName         string                                                                `json:"realName" api:"required"`
	Reputation       float64                                                               `json:"reputation" api:"required"`
	ResearcherSlug   string                                                                `json:"researcherSlug" api:"required"`
	Username         string                                                                `json:"username" api:"required"`
	WeeklyReputation float64                                                               `json:"weeklyReputation" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		Avatar           respjson.Field
		GoogleScholarID  respjson.Field
		Institution      respjson.Field
		RealName         respjson.Field
		Reputation       respjson.Field
		ResearcherSlug   respjson.Field
		Username         respjson.Field
		WeeklyReputation respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3GetUnrelatedResponseRecommendationContextFollowedLiker) RawJSON() string {
	return r.JSON.raw
}
func (r *PaperV3GetUnrelatedResponseRecommendationContextFollowedLiker) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3GetUnrelatedResponseRecommendationContextFollowedLikerAvatar struct {
	// Any of "full_size", "thumbnail".
	Type string `json:"type" api:"required"`
	URL  string `json:"url" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3GetUnrelatedResponseRecommendationContextFollowedLikerAvatar) RawJSON() string {
	return r.JSON.raw
}
func (r *PaperV3GetUnrelatedResponseRecommendationContextFollowedLikerAvatar) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3CommentParams struct {
	// Any of "anonymous", "general", "personal", "research", "resources".
	Tag    PaperV3CommentParamsTag `json:"tag,omitzero" api:"required"`
	Body   param.Opt[string]       `json:"body,omitzero"`
	Parent param.Opt[string]       `json:"parent,omitzero" format:"uuid"`
	Title  param.Opt[string]       `json:"title,omitzero"`
	paramObj
}

func (r PaperV3CommentParams) MarshalJSON() (data []byte, err error) {
	type shadow PaperV3CommentParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PaperV3CommentParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3CommentParamsTag string

const (
	PaperV3CommentParamsTagAnonymous PaperV3CommentParamsTag = "anonymous"
	PaperV3CommentParamsTagGeneral   PaperV3CommentParamsTag = "general"
	PaperV3CommentParamsTagPersonal  PaperV3CommentParamsTag = "personal"
	PaperV3CommentParamsTagResearch  PaperV3CommentParamsTag = "research"
	PaperV3CommentParamsTagResources PaperV3CommentParamsTag = "resources"
)

type PaperV3DeleteVotesParams struct {
	Body []string
	paramObj
}

func (r PaperV3DeleteVotesParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *PaperV3DeleteVotesParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3ImplementationParams struct {
	URL string `json:"url" api:"required"`
	paramObj
}

func (r PaperV3ImplementationParams) MarshalJSON() (data []byte, err error) {
	type shadow PaperV3ImplementationParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PaperV3ImplementationParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3KickoffPaperCountriesParams struct {
	// Number of papers to process in each batch
	Batch param.Opt[float64] `json:"batch,omitzero"`
	// Maximum number of papers to process
	MaxPapers param.Opt[float64] `json:"maxPapers,omitzero"`
	// Only process papers at least this many months old
	Months param.Opt[float64] `json:"months,omitzero"`
	paramObj
}

func (r PaperV3KickoffPaperCountriesParams) MarshalJSON() (data []byte, err error) {
	type shadow PaperV3KickoffPaperCountriesParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PaperV3KickoffPaperCountriesParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3LikeParams struct {
	// Any of "true", "false".
	Liked PaperV3LikeParamsLiked `query:"liked,omitzero" api:"required" json:"-"`
	paramObj
}

// URLQuery serializes [PaperV3LikeParams]'s query parameters as `url.Values`.
func (r PaperV3LikeParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PaperV3LikeParamsLiked string

const (
	PaperV3LikeParamsLikedTrue  PaperV3LikeParamsLiked = "true"
	PaperV3LikeParamsLikedFalse PaperV3LikeParamsLiked = "false"
)

type PaperV3ProcessAIParams struct {
	// Any of "am", "ar", "az", "bg", "bn", "ca", "cs", "da", "de", "el", "en", "es",
	// "et", "fa", "fi", "fr", "gu", "ha", "he", "hi", "hr", "hu", "id", "it", "ja",
	// "ka", "kn", "ko", "lt", "lv", "ml", "mr", "ms", "my", "ne", "nl", "no", "pa",
	// "pl", "pt", "ro", "ru", "si", "sk", "sl", "sr", "sv", "sw", "ta", "te", "th",
	// "tl", "tr", "uk", "ur", "uz", "vi", "yo", "zh".
	PreferredLanguage PaperV3ProcessAIParamsPreferredLanguage `query:"preferredLanguage,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PaperV3ProcessAIParams]'s query parameters as `url.Values`.
func (r PaperV3ProcessAIParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PaperV3ProcessAIParamsPreferredLanguage string

const (
	PaperV3ProcessAIParamsPreferredLanguageAm PaperV3ProcessAIParamsPreferredLanguage = "am"
	PaperV3ProcessAIParamsPreferredLanguageAr PaperV3ProcessAIParamsPreferredLanguage = "ar"
	PaperV3ProcessAIParamsPreferredLanguageAz PaperV3ProcessAIParamsPreferredLanguage = "az"
	PaperV3ProcessAIParamsPreferredLanguageBg PaperV3ProcessAIParamsPreferredLanguage = "bg"
	PaperV3ProcessAIParamsPreferredLanguageBn PaperV3ProcessAIParamsPreferredLanguage = "bn"
	PaperV3ProcessAIParamsPreferredLanguageCa PaperV3ProcessAIParamsPreferredLanguage = "ca"
	PaperV3ProcessAIParamsPreferredLanguageCs PaperV3ProcessAIParamsPreferredLanguage = "cs"
	PaperV3ProcessAIParamsPreferredLanguageDa PaperV3ProcessAIParamsPreferredLanguage = "da"
	PaperV3ProcessAIParamsPreferredLanguageDe PaperV3ProcessAIParamsPreferredLanguage = "de"
	PaperV3ProcessAIParamsPreferredLanguageEl PaperV3ProcessAIParamsPreferredLanguage = "el"
	PaperV3ProcessAIParamsPreferredLanguageEn PaperV3ProcessAIParamsPreferredLanguage = "en"
	PaperV3ProcessAIParamsPreferredLanguageEs PaperV3ProcessAIParamsPreferredLanguage = "es"
	PaperV3ProcessAIParamsPreferredLanguageEt PaperV3ProcessAIParamsPreferredLanguage = "et"
	PaperV3ProcessAIParamsPreferredLanguageFa PaperV3ProcessAIParamsPreferredLanguage = "fa"
	PaperV3ProcessAIParamsPreferredLanguageFi PaperV3ProcessAIParamsPreferredLanguage = "fi"
	PaperV3ProcessAIParamsPreferredLanguageFr PaperV3ProcessAIParamsPreferredLanguage = "fr"
	PaperV3ProcessAIParamsPreferredLanguageGu PaperV3ProcessAIParamsPreferredLanguage = "gu"
	PaperV3ProcessAIParamsPreferredLanguageHa PaperV3ProcessAIParamsPreferredLanguage = "ha"
	PaperV3ProcessAIParamsPreferredLanguageHe PaperV3ProcessAIParamsPreferredLanguage = "he"
	PaperV3ProcessAIParamsPreferredLanguageHi PaperV3ProcessAIParamsPreferredLanguage = "hi"
	PaperV3ProcessAIParamsPreferredLanguageHr PaperV3ProcessAIParamsPreferredLanguage = "hr"
	PaperV3ProcessAIParamsPreferredLanguageHu PaperV3ProcessAIParamsPreferredLanguage = "hu"
	PaperV3ProcessAIParamsPreferredLanguageID PaperV3ProcessAIParamsPreferredLanguage = "id"
	PaperV3ProcessAIParamsPreferredLanguageIt PaperV3ProcessAIParamsPreferredLanguage = "it"
	PaperV3ProcessAIParamsPreferredLanguageJa PaperV3ProcessAIParamsPreferredLanguage = "ja"
	PaperV3ProcessAIParamsPreferredLanguageKa PaperV3ProcessAIParamsPreferredLanguage = "ka"
	PaperV3ProcessAIParamsPreferredLanguageKn PaperV3ProcessAIParamsPreferredLanguage = "kn"
	PaperV3ProcessAIParamsPreferredLanguageKo PaperV3ProcessAIParamsPreferredLanguage = "ko"
	PaperV3ProcessAIParamsPreferredLanguageLt PaperV3ProcessAIParamsPreferredLanguage = "lt"
	PaperV3ProcessAIParamsPreferredLanguageLv PaperV3ProcessAIParamsPreferredLanguage = "lv"
	PaperV3ProcessAIParamsPreferredLanguageMl PaperV3ProcessAIParamsPreferredLanguage = "ml"
	PaperV3ProcessAIParamsPreferredLanguageMr PaperV3ProcessAIParamsPreferredLanguage = "mr"
	PaperV3ProcessAIParamsPreferredLanguageMs PaperV3ProcessAIParamsPreferredLanguage = "ms"
	PaperV3ProcessAIParamsPreferredLanguageMy PaperV3ProcessAIParamsPreferredLanguage = "my"
	PaperV3ProcessAIParamsPreferredLanguageNe PaperV3ProcessAIParamsPreferredLanguage = "ne"
	PaperV3ProcessAIParamsPreferredLanguageNl PaperV3ProcessAIParamsPreferredLanguage = "nl"
	PaperV3ProcessAIParamsPreferredLanguageNo PaperV3ProcessAIParamsPreferredLanguage = "no"
	PaperV3ProcessAIParamsPreferredLanguagePa PaperV3ProcessAIParamsPreferredLanguage = "pa"
	PaperV3ProcessAIParamsPreferredLanguagePl PaperV3ProcessAIParamsPreferredLanguage = "pl"
	PaperV3ProcessAIParamsPreferredLanguagePt PaperV3ProcessAIParamsPreferredLanguage = "pt"
	PaperV3ProcessAIParamsPreferredLanguageRo PaperV3ProcessAIParamsPreferredLanguage = "ro"
	PaperV3ProcessAIParamsPreferredLanguageRu PaperV3ProcessAIParamsPreferredLanguage = "ru"
	PaperV3ProcessAIParamsPreferredLanguageSi PaperV3ProcessAIParamsPreferredLanguage = "si"
	PaperV3ProcessAIParamsPreferredLanguageSk PaperV3ProcessAIParamsPreferredLanguage = "sk"
	PaperV3ProcessAIParamsPreferredLanguageSl PaperV3ProcessAIParamsPreferredLanguage = "sl"
	PaperV3ProcessAIParamsPreferredLanguageSr PaperV3ProcessAIParamsPreferredLanguage = "sr"
	PaperV3ProcessAIParamsPreferredLanguageSv PaperV3ProcessAIParamsPreferredLanguage = "sv"
	PaperV3ProcessAIParamsPreferredLanguageSw PaperV3ProcessAIParamsPreferredLanguage = "sw"
	PaperV3ProcessAIParamsPreferredLanguageTa PaperV3ProcessAIParamsPreferredLanguage = "ta"
	PaperV3ProcessAIParamsPreferredLanguageTe PaperV3ProcessAIParamsPreferredLanguage = "te"
	PaperV3ProcessAIParamsPreferredLanguageTh PaperV3ProcessAIParamsPreferredLanguage = "th"
	PaperV3ProcessAIParamsPreferredLanguageTl PaperV3ProcessAIParamsPreferredLanguage = "tl"
	PaperV3ProcessAIParamsPreferredLanguageTr PaperV3ProcessAIParamsPreferredLanguage = "tr"
	PaperV3ProcessAIParamsPreferredLanguageUk PaperV3ProcessAIParamsPreferredLanguage = "uk"
	PaperV3ProcessAIParamsPreferredLanguageUr PaperV3ProcessAIParamsPreferredLanguage = "ur"
	PaperV3ProcessAIParamsPreferredLanguageUz PaperV3ProcessAIParamsPreferredLanguage = "uz"
	PaperV3ProcessAIParamsPreferredLanguageVi PaperV3ProcessAIParamsPreferredLanguage = "vi"
	PaperV3ProcessAIParamsPreferredLanguageYo PaperV3ProcessAIParamsPreferredLanguage = "yo"
	PaperV3ProcessAIParamsPreferredLanguageZh PaperV3ProcessAIParamsPreferredLanguage = "zh"
)

type PaperV3ProcessCountriesParams struct {
	// Array of universal paper IDs (versionless)
	UniversalPaperIDs []string `json:"universalPaperIds,omitzero" api:"required"`
	paramObj
}

func (r PaperV3ProcessCountriesParams) MarshalJSON() (data []byte, err error) {
	type shadow PaperV3ProcessCountriesParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PaperV3ProcessCountriesParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3RequestImplementationParams struct {
	PaperTitle       string            `json:"paperTitle" api:"required"`
	UniversalPaperID string            `json:"universalPaperId" api:"required"`
	AdditionalInfo   param.Opt[string] `json:"additionalInfo,omitzero"`
	paramObj
}

func (r PaperV3RequestImplementationParams) MarshalJSON() (data []byte, err error) {
	type shadow PaperV3RequestImplementationParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PaperV3RequestImplementationParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3GetAllParams struct {
	Limit param.Opt[string] `query:"limit,omitzero" json:"-"`
	Skip  param.Opt[string] `query:"skip,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PaperV3GetAllParams]'s query parameters as `url.Values`.
func (r PaperV3GetAllParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PaperV3GetDiversePapersParams struct {
	Topics string `query:"topics" api:"required" json:"-"`
	paramObj
}

// URLQuery serializes [PaperV3GetDiversePapersParams]'s query parameters as
// `url.Values`.
func (r PaperV3GetDiversePapersParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PaperV3GetFeedParams struct {
	// Any of "3 Days", "7 Days", "30 Days", "90 Days", "All time".
	Interval PaperV3GetFeedParamsInterval `query:"interval,omitzero" api:"required" json:"-"`
	PageNum  string                       `query:"pageNum" api:"required" json:"-"`
	PageSize string                       `query:"pageSize" api:"required" json:"-"`
	// Any of "Hot", "Comments", "Views", "Likes", "GitHub", "Recommended", "ForYou",
	// "Recent".
	Sort                 PaperV3GetFeedParamsSort `query:"sort,omitzero" api:"required" json:"-"`
	FeedCursor           param.Opt[string]        `query:"feedCursor,omitzero" json:"-"`
	IncludeExternalBlogs param.Opt[string]        `query:"includeExternalBlogs,omitzero" json:"-"`
	Runnable             param.Opt[string]        `query:"runnable,omitzero" json:"-"`
	Topics               param.Opt[string]        `query:"topics,omitzero" json:"-"`
	// A versionless universal paper ID (e.g. 1706.03762)
	UniversalID param.Opt[string] `query:"universalId,omitzero" json:"-"`
	// Any of "GitHub".
	Source PaperV3GetFeedParamsSource `query:"source,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PaperV3GetFeedParams]'s query parameters as `url.Values`.
func (r PaperV3GetFeedParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PaperV3GetFeedParamsInterval string

const (
	PaperV3GetFeedParamsInterval3Days   PaperV3GetFeedParamsInterval = "3 Days"
	PaperV3GetFeedParamsInterval7Days   PaperV3GetFeedParamsInterval = "7 Days"
	PaperV3GetFeedParamsInterval30Days  PaperV3GetFeedParamsInterval = "30 Days"
	PaperV3GetFeedParamsInterval90Days  PaperV3GetFeedParamsInterval = "90 Days"
	PaperV3GetFeedParamsIntervalAllTime PaperV3GetFeedParamsInterval = "All time"
)

type PaperV3GetFeedParamsSort string

const (
	PaperV3GetFeedParamsSortHot         PaperV3GetFeedParamsSort = "Hot"
	PaperV3GetFeedParamsSortComments    PaperV3GetFeedParamsSort = "Comments"
	PaperV3GetFeedParamsSortViews       PaperV3GetFeedParamsSort = "Views"
	PaperV3GetFeedParamsSortLikes       PaperV3GetFeedParamsSort = "Likes"
	PaperV3GetFeedParamsSortGitHub      PaperV3GetFeedParamsSort = "GitHub"
	PaperV3GetFeedParamsSortRecommended PaperV3GetFeedParamsSort = "Recommended"
	PaperV3GetFeedParamsSortForYou      PaperV3GetFeedParamsSort = "ForYou"
	PaperV3GetFeedParamsSortRecent      PaperV3GetFeedParamsSort = "Recent"
)

type PaperV3GetFeedParamsSource string

const (
	PaperV3GetFeedParamsSourceGitHub PaperV3GetFeedParamsSource = "GitHub"
)

type PaperV3GetSimilarPapersParams struct {
	Exclude param.Opt[string] `query:"exclude,omitzero" json:"-"`
	Limit   param.Opt[string] `query:"limit,omitzero" json:"-"`
	// Any of "false", "true".
	ExcludeLikes PaperV3GetSimilarPapersParamsExcludeLikes `query:"excludeLikes,omitzero" json:"-"`
	// Any of "3 Days", "7 Days", "30 Days", "90 Days", "All time".
	Interval PaperV3GetSimilarPapersParamsInterval `query:"interval,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PaperV3GetSimilarPapersParams]'s query parameters as
// `url.Values`.
func (r PaperV3GetSimilarPapersParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PaperV3GetSimilarPapersParamsExcludeLikes string

const (
	PaperV3GetSimilarPapersParamsExcludeLikesFalse PaperV3GetSimilarPapersParamsExcludeLikes = "false"
	PaperV3GetSimilarPapersParamsExcludeLikesTrue  PaperV3GetSimilarPapersParamsExcludeLikes = "true"
)

type PaperV3GetSimilarPapersParamsInterval string

const (
	PaperV3GetSimilarPapersParamsInterval3Days   PaperV3GetSimilarPapersParamsInterval = "3 Days"
	PaperV3GetSimilarPapersParamsInterval7Days   PaperV3GetSimilarPapersParamsInterval = "7 Days"
	PaperV3GetSimilarPapersParamsInterval30Days  PaperV3GetSimilarPapersParamsInterval = "30 Days"
	PaperV3GetSimilarPapersParamsInterval90Days  PaperV3GetSimilarPapersParamsInterval = "90 Days"
	PaperV3GetSimilarPapersParamsIntervalAllTime PaperV3GetSimilarPapersParamsInterval = "All time"
)

type PaperV3GetUnrelatedParams struct {
	Limit  string `query:"limit" api:"required" json:"-"`
	Papers string `query:"papers" api:"required" json:"-"`
	Topics string `query:"topics" api:"required" json:"-"`
	paramObj
}

// URLQuery serializes [PaperV3GetUnrelatedParams]'s query parameters as
// `url.Values`.
func (r PaperV3GetUnrelatedParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
