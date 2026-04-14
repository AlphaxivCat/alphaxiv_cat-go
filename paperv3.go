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
// Source file: `api-server/src/controllers/papers/v3/get-paper.controller.ts`
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
// Source file: `api-server/src/controllers/papers/v3/post-comment.controller.ts`
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
// `api-server/src/controllers/papers/v3/remove-vote-batch.controller.ts`
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
// `api-server/src/controllers/papers/v3/create-or-update-implementation.controller.ts`
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
// `api-server/src/controllers/papers/v3/kickoff-paper-countries.controller.ts`
func (r *PaperV3Service) KickoffPaperCountries(ctx context.Context, body PaperV3KickoffPaperCountriesParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "papers/v3/kickoff-paper-countries"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return err
}

// Kickoff paper full text processing for recent papers
//
// Source file:
// `api-server/src/controllers/papers/v3/kickoff-paper-full-text.controller.ts`
func (r *PaperV3Service) KickoffPaperFullText(ctx context.Context, body PaperV3KickoffPaperFullTextParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "papers/v3/kickoff-paper-full-text"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return err
}

// Kickoff paper podcasts on Uptash for a subset of paper groups
//
// Source file:
// `api-server/src/controllers/papers/v3/kickoff-paper-podcasts.controller.ts`
func (r *PaperV3Service) KickoffPaperPodcasts(ctx context.Context, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "papers/v3/kickoff-paper-podcasts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return err
}

// Kickoff background job to generate thumbnails for trending papers
//
// Source file:
// `api-server/src/controllers/papers/v3/kickoff-thumbnails-trending-papers.controller.ts`
func (r *PaperV3Service) KickoffThumbnailsTrendingPapers(ctx context.Context, opts ...option.RequestOption) (res *PaperV3KickoffThumbnailsTrendingPapersResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "papers/v3/kickoff-thumbnails-trending-papers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Kickoff X mentions sync for hot papers. Uses x-mentions-sync-queue with
// parallelism=1 and built-in delays.
//
// Source file:
// `api-server/src/controllers/papers/v3/kickoff-x-mentions-sync.controller.ts`
func (r *PaperV3Service) KickoffXMentionsSync(ctx context.Context, body PaperV3KickoffXMentionsSyncParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "papers/v3/kickoff-x-mentions-sync"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return err
}

// Toggle your like status on a paper group
//
// Source file: `api-server/src/controllers/papers/v3/like-paper.controller.ts`
func (r *PaperV3Service) Like(ctx context.Context, group string, opts ...option.RequestOption) (res *PaperV3LikeResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if group == "" {
		err = errors.New("missing required group parameter")
		return nil, err
	}
	path := fmt.Sprintf("papers/v3/%s/like", group)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Generates a podcast for a paper group
//
// Source file:
// `api-server/src/controllers/papers/v3/generate-paper-podcast.controller.ts`
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
// Source file: `api-server/src/controllers/papers/v3/process-ai.controller.ts`
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
// `api-server/src/controllers/papers/v3/process-countries.controller.ts`
//
// Deprecated: deprecated
func (r *PaperV3Service) ProcessCountries(ctx context.Context, body PaperV3ProcessCountriesParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "papers/v3/process-countries"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return err
}

// Processes and extracts full text from paper PDFs for indexing and search
//
// Source file:
// `api-server/src/controllers/papers/v3/process-full-text.controller.ts`
//
// Deprecated: deprecated
func (r *PaperV3Service) ProcessFullText(ctx context.Context, body PaperV3ProcessFullTextParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "papers/v3/process-full-text"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return err
}

// Clear 'is_last_X_days' flags from paper embeddings that have become too old
//
// Source file:
// `api-server/src/controllers/papers/v3/prune-embeddings-by-date.controller.ts`
//
// Deprecated: deprecated
func (r *PaperV3Service) PruneEmbeddingsByDate(ctx context.Context, opts ...option.RequestOption) (res *PaperV3PruneEmbeddingsByDateResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "papers/v3/prune-embeddings-by-date"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Toggle your implementation request status on a paper group
//
// Source file:
// `api-server/src/controllers/papers/v3/request-paper-implementation.controller.ts`
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
// `api-server/src/controllers/papers/v3/request-podcast.controller.ts`
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
// Source file: `api-server/src/controllers/papers/v3/get-all-papers.controller.ts`
func (r *PaperV3Service) GetAll(ctx context.Context, query PaperV3GetAllParams, opts ...option.RequestOption) (res *PaperV3GetAllResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "papers/v3/all"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Get an initial batch of diverse papers on the given topics for recommendations
//
// Source file: `api-server/src/controllers/papers/v3/diverse-papers.controller.ts`
func (r *PaperV3Service) GetDiversePapers(ctx context.Context, query PaperV3GetDiversePapersParams, opts ...option.RequestOption) (res *[]PaperV3GetDiversePapersResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "papers/v3/diverse-papers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Get an optionally filtered list of papers for the main feed
//
// Source file: `api-server/src/controllers/papers/v3/feed.controller.ts`
func (r *PaperV3Service) GetFeed(ctx context.Context, query PaperV3GetFeedParams, opts ...option.RequestOption) (res *PaperV3GetFeedResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "papers/v3/feed"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Get list of figure URLs for a paper
//
// Source file:
// `api-server/src/controllers/papers/v3/get-paper-figures.controller.ts`
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
// Source file: `api-server/src/controllers/papers/v3/get-full-text.controller.ts`
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

// Retrieve geographical trends and analytics data for papers
//
// Source file: `api-server/src/controllers/papers/v3/get-geo-trends.controller.ts`
func (r *PaperV3Service) GetGeoTrends(ctx context.Context, query PaperV3GetGeoTrendsParams, opts ...option.RequestOption) (res *PaperV3GetGeoTrendsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "papers/v3/geo-trends"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Retrieve metrics for a paper (comments count, upvotes, views)
//
// Source file: `api-server/src/controllers/papers/v3/get-metrics.controller.ts`
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

// Retrieve top papers by country with optional country filter
//
// Source file:
// `api-server/src/controllers/papers/v3/get-papers-by-country.controller.ts`
func (r *PaperV3Service) GetPapersByCountry(ctx context.Context, query PaperV3GetPapersByCountryParams, opts ...option.RequestOption) (res *[]PaperV3GetPapersByCountryResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "papers/v3/papers-by-country"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Retrieve paper data for paper preview cards
//
// Source file:
// `api-server/src/controllers/papers/v3/get-paper-preview.controller.ts`
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
// `api-server/src/controllers/papers/v3/get-similar-papers.controller.ts`
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
// `api-server/src/controllers/papers/v3/unrelated-papers.controller.ts`
func (r *PaperV3Service) GetUnrelated(ctx context.Context, query PaperV3GetUnrelatedParams, opts ...option.RequestOption) (res *[]PaperV3GetUnrelatedResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "papers/v3/unrelated"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Track paper view event for analytics
//
// Source file:
// `api-server/src/controllers/papers/v3/mark-paper-view.controller.ts`
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
	PublicationDate      float64                      `json:"publicationDate" api:"required"`
	Resources            []PaperV3GetResponseResource `json:"resources" api:"required"`
	SourceName           string                       `json:"sourceName" api:"required"`
	SourceURL            string                       `json:"sourceUrl" api:"required"`
	Title                string                       `json:"title" api:"required"`
	// Any of "public", "private", "community".
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
	PaperV3GetResponseTypePublic    PaperV3GetResponseType = "public"
	PaperV3GetResponseTypePrivate   PaperV3GetResponseType = "private"
	PaperV3GetResponseTypeCommunity PaperV3GetResponseType = "community"
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

type PaperV3KickoffThumbnailsTrendingPapersResponse struct {
	Data PaperV3KickoffThumbnailsTrendingPapersResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3KickoffThumbnailsTrendingPapersResponse) RawJSON() string { return r.JSON.raw }
func (r *PaperV3KickoffThumbnailsTrendingPapersResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3KickoffThumbnailsTrendingPapersResponseData struct {
	Message string `json:"message" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3KickoffThumbnailsTrendingPapersResponseData) RawJSON() string { return r.JSON.raw }
func (r *PaperV3KickoffThumbnailsTrendingPapersResponseData) UnmarshalJSON(data []byte) error {
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

type PaperV3PruneEmbeddingsByDateResponse struct {
	RowsUpdated []float64 `json:"rowsUpdated" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		RowsUpdated respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3PruneEmbeddingsByDateResponse) RawJSON() string { return r.JSON.raw }
func (r *PaperV3PruneEmbeddingsByDateResponse) UnmarshalJSON(data []byte) error {
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
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3RequestPodcastResponse) RawJSON() string { return r.JSON.raw }
func (r *PaperV3RequestPodcastResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

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
	FirstPublicationDate string                                            `json:"first_publication_date" api:"required"`
	FullAuthors          []PaperV3GetDiversePapersResponseFullAuthor       `json:"full_authors" api:"required"`
	GitHubStars          float64                                           `json:"github_stars" api:"required"`
	GitHubURL            string                                            `json:"github_url" api:"required"`
	ImageURL             string                                            `json:"image_url" api:"required"`
	Metrics              PaperV3GetDiversePapersResponseMetrics            `json:"metrics" api:"required"`
	OrganizationInfo     []PaperV3GetDiversePapersResponseOrganizationInfo `json:"organization_info" api:"required"`
	PaperGroupID         string                                            `json:"paper_group_id" api:"required" format:"uuid"`
	PaperSummary         PaperV3GetDiversePapersResponsePaperSummary       `json:"paper_summary" api:"required"`
	PublicationDate      string                                            `json:"publication_date" api:"required"`
	Title                string                                            `json:"title" api:"required"`
	Topics               []string                                          `json:"topics" api:"required"`
	// A versionless universal paper ID (e.g. 1706.03762)
	UniversalPaperID string `json:"universal_paper_id" api:"required"`
	UpdatedAt        string `json:"updated_at" api:"required"`
	VersionID        string `json:"version_id" api:"required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                   respjson.Field
		Abstract             respjson.Field
		AuthorInfo           respjson.Field
		Authors              respjson.Field
		CanonicalID          respjson.Field
		FirstPublicationDate respjson.Field
		FullAuthors          respjson.Field
		GitHubStars          respjson.Field
		GitHubURL            respjson.Field
		ImageURL             respjson.Field
		Metrics              respjson.Field
		OrganizationInfo     respjson.Field
		PaperGroupID         respjson.Field
		PaperSummary         respjson.Field
		PublicationDate      respjson.Field
		Title                respjson.Field
		Topics               respjson.Field
		UniversalPaperID     respjson.Field
		UpdatedAt            respjson.Field
		VersionID            respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
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

type PaperV3GetDiversePapersResponseFullAuthor struct {
	ID       string `json:"id" api:"required" format:"uuid"`
	FullName string `json:"full_name" api:"required"`
	UserID   string `json:"user_id" api:"required" format:"uuid"`
	Username string `json:"username" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		FullName    respjson.Field
		UserID      respjson.Field
		Username    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3GetDiversePapersResponseFullAuthor) RawJSON() string { return r.JSON.raw }
func (r *PaperV3GetDiversePapersResponseFullAuthor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3GetDiversePapersResponseMetrics struct {
	PublicTotalVotes float64                                           `json:"public_total_votes" api:"required"`
	TotalVotes       float64                                           `json:"total_votes" api:"required"`
	VisitsCount      PaperV3GetDiversePapersResponseMetricsVisitsCount `json:"visits_count" api:"required"`
	XLikes           float64                                           `json:"x_likes" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PublicTotalVotes respjson.Field
		TotalVotes       respjson.Field
		VisitsCount      respjson.Field
		XLikes           respjson.Field
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

type PaperV3GetFeedResponse struct {
	Page   float64                       `json:"page" api:"required"`
	Papers []PaperV3GetFeedResponsePaper `json:"papers" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Page        respjson.Field
		Papers      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
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
	FirstPublicationDate string                                        `json:"first_publication_date" api:"required"`
	FullAuthors          []PaperV3GetFeedResponsePaperFullAuthor       `json:"full_authors" api:"required"`
	GitHubStars          float64                                       `json:"github_stars" api:"required"`
	GitHubURL            string                                        `json:"github_url" api:"required"`
	ImageURL             string                                        `json:"image_url" api:"required"`
	Metrics              PaperV3GetFeedResponsePaperMetrics            `json:"metrics" api:"required"`
	OrganizationInfo     []PaperV3GetFeedResponsePaperOrganizationInfo `json:"organization_info" api:"required"`
	PaperGroupID         string                                        `json:"paper_group_id" api:"required" format:"uuid"`
	PaperSummary         PaperV3GetFeedResponsePaperPaperSummary       `json:"paper_summary" api:"required"`
	PublicationDate      string                                        `json:"publication_date" api:"required"`
	Title                string                                        `json:"title" api:"required"`
	Topics               []string                                      `json:"topics" api:"required"`
	// A versionless universal paper ID (e.g. 1706.03762)
	UniversalPaperID string `json:"universal_paper_id" api:"required"`
	UpdatedAt        string `json:"updated_at" api:"required"`
	VersionID        string `json:"version_id" api:"required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                   respjson.Field
		Abstract             respjson.Field
		AuthorInfo           respjson.Field
		Authors              respjson.Field
		CanonicalID          respjson.Field
		FirstPublicationDate respjson.Field
		FullAuthors          respjson.Field
		GitHubStars          respjson.Field
		GitHubURL            respjson.Field
		ImageURL             respjson.Field
		Metrics              respjson.Field
		OrganizationInfo     respjson.Field
		PaperGroupID         respjson.Field
		PaperSummary         respjson.Field
		PublicationDate      respjson.Field
		Title                respjson.Field
		Topics               respjson.Field
		UniversalPaperID     respjson.Field
		UpdatedAt            respjson.Field
		VersionID            respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
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

type PaperV3GetFeedResponsePaperFullAuthor struct {
	ID       string `json:"id" api:"required" format:"uuid"`
	FullName string `json:"full_name" api:"required"`
	UserID   string `json:"user_id" api:"required" format:"uuid"`
	Username string `json:"username" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		FullName    respjson.Field
		UserID      respjson.Field
		Username    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3GetFeedResponsePaperFullAuthor) RawJSON() string { return r.JSON.raw }
func (r *PaperV3GetFeedResponsePaperFullAuthor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3GetFeedResponsePaperMetrics struct {
	PublicTotalVotes float64                                       `json:"public_total_votes" api:"required"`
	TotalVotes       float64                                       `json:"total_votes" api:"required"`
	VisitsCount      PaperV3GetFeedResponsePaperMetricsVisitsCount `json:"visits_count" api:"required"`
	XLikes           float64                                       `json:"x_likes" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PublicTotalVotes respjson.Field
		TotalVotes       respjson.Field
		VisitsCount      respjson.Field
		XLikes           respjson.Field
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

type PaperV3GetGeoTrendsResponse struct {
	FrequentCollaborations   []PaperV3GetGeoTrendsResponseFrequentCollaboration    `json:"frequentCollaborations" api:"required"`
	PaperCountGraph          map[string][]float64                                  `json:"paperCountGraph" api:"required"`
	PaperStarsGraph          map[string][]float64                                  `json:"paperStarsGraph" api:"required"`
	TopCountriesByPaperCount []PaperV3GetGeoTrendsResponseTopCountriesByPaperCount `json:"topCountriesByPaperCount" api:"required"`
	TopCountriesByStars      []PaperV3GetGeoTrendsResponseTopCountriesByStar       `json:"topCountriesByStars" api:"required"`
	TopGitHubRepos           []PaperV3GetGeoTrendsResponseTopGitHubRepo            `json:"topGitHubRepos" api:"required"`
	TotalPapersCount         float64                                               `json:"totalPapersCount" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FrequentCollaborations   respjson.Field
		PaperCountGraph          respjson.Field
		PaperStarsGraph          respjson.Field
		TopCountriesByPaperCount respjson.Field
		TopCountriesByStars      respjson.Field
		TopGitHubRepos           respjson.Field
		TotalPapersCount         respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3GetGeoTrendsResponse) RawJSON() string { return r.JSON.raw }
func (r *PaperV3GetGeoTrendsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3GetGeoTrendsResponseFrequentCollaboration struct {
	CollaborationCount float64  `json:"collaborationCount" api:"required"`
	Countries          []string `json:"countries" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CollaborationCount respjson.Field
		Countries          respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3GetGeoTrendsResponseFrequentCollaboration) RawJSON() string { return r.JSON.raw }
func (r *PaperV3GetGeoTrendsResponseFrequentCollaboration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3GetGeoTrendsResponseTopCountriesByPaperCount struct {
	Count   float64 `json:"count" api:"required"`
	Country string  `json:"country" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Country     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3GetGeoTrendsResponseTopCountriesByPaperCount) RawJSON() string { return r.JSON.raw }
func (r *PaperV3GetGeoTrendsResponseTopCountriesByPaperCount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3GetGeoTrendsResponseTopCountriesByStar struct {
	Country    string  `json:"country" api:"required"`
	TotalStars float64 `json:"totalStars" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Country     respjson.Field
		TotalStars  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3GetGeoTrendsResponseTopCountriesByStar) RawJSON() string { return r.JSON.raw }
func (r *PaperV3GetGeoTrendsResponseTopCountriesByStar) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3GetGeoTrendsResponseTopGitHubRepo struct {
	Countries            []string `json:"countries" api:"required"`
	Description          string   `json:"description" api:"required"`
	FirstPublicationDate string   `json:"firstPublicationDate" api:"required"`
	Language             string   `json:"language" api:"required"`
	PaperGroupID         string   `json:"paperGroupId" api:"required"`
	Stars                float64  `json:"stars" api:"required"`
	StarsDaily           float64  `json:"starsDaily" api:"required"`
	Title                string   `json:"title" api:"required"`
	UniversalID          string   `json:"universalId" api:"required"`
	URL                  string   `json:"url" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Countries            respjson.Field
		Description          respjson.Field
		FirstPublicationDate respjson.Field
		Language             respjson.Field
		PaperGroupID         respjson.Field
		Stars                respjson.Field
		StarsDaily           respjson.Field
		Title                respjson.Field
		UniversalID          respjson.Field
		URL                  respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3GetGeoTrendsResponseTopGitHubRepo) RawJSON() string { return r.JSON.raw }
func (r *PaperV3GetGeoTrendsResponseTopGitHubRepo) UnmarshalJSON(data []byte) error {
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

type PaperV3GetPapersByCountryResponse struct {
	Abstract             string   `json:"abstract" api:"required"`
	CitationsCount       float64  `json:"citationsCount" api:"required"`
	CommentsCount        float64  `json:"commentsCount" api:"required"`
	Countries            []string `json:"countries" api:"required"`
	FirstPublicationDate string   `json:"firstPublicationDate" api:"required"`
	PaperGroupID         string   `json:"paperGroupId" api:"required"`
	PublicationDate      string   `json:"publicationDate" api:"required"`
	PublicTotalVotes     float64  `json:"publicTotalVotes" api:"required"`
	Title                string   `json:"title" api:"required"`
	TotalVotes           float64  `json:"totalVotes" api:"required"`
	UniversalID          string   `json:"universalId" api:"required"`
	VisitsAll            float64  `json:"visitsAll" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Abstract             respjson.Field
		CitationsCount       respjson.Field
		CommentsCount        respjson.Field
		Countries            respjson.Field
		FirstPublicationDate respjson.Field
		PaperGroupID         respjson.Field
		PublicationDate      respjson.Field
		PublicTotalVotes     respjson.Field
		Title                respjson.Field
		TotalVotes           respjson.Field
		UniversalID          respjson.Field
		VisitsAll            respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3GetPapersByCountryResponse) RawJSON() string { return r.JSON.raw }
func (r *PaperV3GetPapersByCountryResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3GetPreviewResponse struct {
	ID         string                                `json:"id" api:"required" format:"uuid"`
	Abstract   string                                `json:"abstract" api:"required"`
	AuthorInfo []PaperV3GetPreviewResponseAuthorInfo `json:"author_info" api:"required"`
	Authors    []string                              `json:"authors" api:"required"`
	// A versioned paper ID (e.g. 1706.03762v1)
	CanonicalID          string                                      `json:"canonical_id" api:"required"`
	FirstPublicationDate string                                      `json:"first_publication_date" api:"required"`
	FullAuthors          []PaperV3GetPreviewResponseFullAuthor       `json:"full_authors" api:"required"`
	GitHubStars          float64                                     `json:"github_stars" api:"required"`
	GitHubURL            string                                      `json:"github_url" api:"required"`
	ImageURL             string                                      `json:"image_url" api:"required"`
	Metrics              PaperV3GetPreviewResponseMetrics            `json:"metrics" api:"required"`
	OrganizationInfo     []PaperV3GetPreviewResponseOrganizationInfo `json:"organization_info" api:"required"`
	PaperGroupID         string                                      `json:"paper_group_id" api:"required" format:"uuid"`
	PaperSummary         PaperV3GetPreviewResponsePaperSummary       `json:"paper_summary" api:"required"`
	PublicationDate      string                                      `json:"publication_date" api:"required"`
	Title                string                                      `json:"title" api:"required"`
	Topics               []string                                    `json:"topics" api:"required"`
	// A versionless universal paper ID (e.g. 1706.03762)
	UniversalPaperID string `json:"universal_paper_id" api:"required"`
	UpdatedAt        string `json:"updated_at" api:"required"`
	VersionID        string `json:"version_id" api:"required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                   respjson.Field
		Abstract             respjson.Field
		AuthorInfo           respjson.Field
		Authors              respjson.Field
		CanonicalID          respjson.Field
		FirstPublicationDate respjson.Field
		FullAuthors          respjson.Field
		GitHubStars          respjson.Field
		GitHubURL            respjson.Field
		ImageURL             respjson.Field
		Metrics              respjson.Field
		OrganizationInfo     respjson.Field
		PaperGroupID         respjson.Field
		PaperSummary         respjson.Field
		PublicationDate      respjson.Field
		Title                respjson.Field
		Topics               respjson.Field
		UniversalPaperID     respjson.Field
		UpdatedAt            respjson.Field
		VersionID            respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
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

type PaperV3GetPreviewResponseFullAuthor struct {
	ID       string `json:"id" api:"required" format:"uuid"`
	FullName string `json:"full_name" api:"required"`
	UserID   string `json:"user_id" api:"required" format:"uuid"`
	Username string `json:"username" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		FullName    respjson.Field
		UserID      respjson.Field
		Username    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3GetPreviewResponseFullAuthor) RawJSON() string { return r.JSON.raw }
func (r *PaperV3GetPreviewResponseFullAuthor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3GetPreviewResponseMetrics struct {
	PublicTotalVotes float64                                     `json:"public_total_votes" api:"required"`
	TotalVotes       float64                                     `json:"total_votes" api:"required"`
	VisitsCount      PaperV3GetPreviewResponseMetricsVisitsCount `json:"visits_count" api:"required"`
	XLikes           float64                                     `json:"x_likes" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PublicTotalVotes respjson.Field
		TotalVotes       respjson.Field
		VisitsCount      respjson.Field
		XLikes           respjson.Field
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

type PaperV3GetSimilarPapersResponse struct {
	ID         string                                      `json:"id" api:"required" format:"uuid"`
	Abstract   string                                      `json:"abstract" api:"required"`
	AuthorInfo []PaperV3GetSimilarPapersResponseAuthorInfo `json:"author_info" api:"required"`
	Authors    []string                                    `json:"authors" api:"required"`
	// A versioned paper ID (e.g. 1706.03762v1)
	CanonicalID          string                                            `json:"canonical_id" api:"required"`
	FirstPublicationDate string                                            `json:"first_publication_date" api:"required"`
	FullAuthors          []PaperV3GetSimilarPapersResponseFullAuthor       `json:"full_authors" api:"required"`
	GitHubStars          float64                                           `json:"github_stars" api:"required"`
	GitHubURL            string                                            `json:"github_url" api:"required"`
	ImageURL             string                                            `json:"image_url" api:"required"`
	Metrics              PaperV3GetSimilarPapersResponseMetrics            `json:"metrics" api:"required"`
	OrganizationInfo     []PaperV3GetSimilarPapersResponseOrganizationInfo `json:"organization_info" api:"required"`
	PaperGroupID         string                                            `json:"paper_group_id" api:"required" format:"uuid"`
	PaperSummary         PaperV3GetSimilarPapersResponsePaperSummary       `json:"paper_summary" api:"required"`
	PublicationDate      string                                            `json:"publication_date" api:"required"`
	Title                string                                            `json:"title" api:"required"`
	Topics               []string                                          `json:"topics" api:"required"`
	// A versionless universal paper ID (e.g. 1706.03762)
	UniversalPaperID string `json:"universal_paper_id" api:"required"`
	UpdatedAt        string `json:"updated_at" api:"required"`
	VersionID        string `json:"version_id" api:"required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                   respjson.Field
		Abstract             respjson.Field
		AuthorInfo           respjson.Field
		Authors              respjson.Field
		CanonicalID          respjson.Field
		FirstPublicationDate respjson.Field
		FullAuthors          respjson.Field
		GitHubStars          respjson.Field
		GitHubURL            respjson.Field
		ImageURL             respjson.Field
		Metrics              respjson.Field
		OrganizationInfo     respjson.Field
		PaperGroupID         respjson.Field
		PaperSummary         respjson.Field
		PublicationDate      respjson.Field
		Title                respjson.Field
		Topics               respjson.Field
		UniversalPaperID     respjson.Field
		UpdatedAt            respjson.Field
		VersionID            respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
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

type PaperV3GetSimilarPapersResponseFullAuthor struct {
	ID       string `json:"id" api:"required" format:"uuid"`
	FullName string `json:"full_name" api:"required"`
	UserID   string `json:"user_id" api:"required" format:"uuid"`
	Username string `json:"username" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		FullName    respjson.Field
		UserID      respjson.Field
		Username    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3GetSimilarPapersResponseFullAuthor) RawJSON() string { return r.JSON.raw }
func (r *PaperV3GetSimilarPapersResponseFullAuthor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3GetSimilarPapersResponseMetrics struct {
	PublicTotalVotes float64                                           `json:"public_total_votes" api:"required"`
	TotalVotes       float64                                           `json:"total_votes" api:"required"`
	VisitsCount      PaperV3GetSimilarPapersResponseMetricsVisitsCount `json:"visits_count" api:"required"`
	XLikes           float64                                           `json:"x_likes" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PublicTotalVotes respjson.Field
		TotalVotes       respjson.Field
		VisitsCount      respjson.Field
		XLikes           respjson.Field
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

type PaperV3GetUnrelatedResponse struct {
	ID         string                                  `json:"id" api:"required" format:"uuid"`
	Abstract   string                                  `json:"abstract" api:"required"`
	AuthorInfo []PaperV3GetUnrelatedResponseAuthorInfo `json:"author_info" api:"required"`
	Authors    []string                                `json:"authors" api:"required"`
	// A versioned paper ID (e.g. 1706.03762v1)
	CanonicalID          string                                        `json:"canonical_id" api:"required"`
	FirstPublicationDate string                                        `json:"first_publication_date" api:"required"`
	FullAuthors          []PaperV3GetUnrelatedResponseFullAuthor       `json:"full_authors" api:"required"`
	GitHubStars          float64                                       `json:"github_stars" api:"required"`
	GitHubURL            string                                        `json:"github_url" api:"required"`
	ImageURL             string                                        `json:"image_url" api:"required"`
	Metrics              PaperV3GetUnrelatedResponseMetrics            `json:"metrics" api:"required"`
	OrganizationInfo     []PaperV3GetUnrelatedResponseOrganizationInfo `json:"organization_info" api:"required"`
	PaperGroupID         string                                        `json:"paper_group_id" api:"required" format:"uuid"`
	PaperSummary         PaperV3GetUnrelatedResponsePaperSummary       `json:"paper_summary" api:"required"`
	PublicationDate      string                                        `json:"publication_date" api:"required"`
	Title                string                                        `json:"title" api:"required"`
	Topics               []string                                      `json:"topics" api:"required"`
	// A versionless universal paper ID (e.g. 1706.03762)
	UniversalPaperID string `json:"universal_paper_id" api:"required"`
	UpdatedAt        string `json:"updated_at" api:"required"`
	VersionID        string `json:"version_id" api:"required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                   respjson.Field
		Abstract             respjson.Field
		AuthorInfo           respjson.Field
		Authors              respjson.Field
		CanonicalID          respjson.Field
		FirstPublicationDate respjson.Field
		FullAuthors          respjson.Field
		GitHubStars          respjson.Field
		GitHubURL            respjson.Field
		ImageURL             respjson.Field
		Metrics              respjson.Field
		OrganizationInfo     respjson.Field
		PaperGroupID         respjson.Field
		PaperSummary         respjson.Field
		PublicationDate      respjson.Field
		Title                respjson.Field
		Topics               respjson.Field
		UniversalPaperID     respjson.Field
		UpdatedAt            respjson.Field
		VersionID            respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
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

type PaperV3GetUnrelatedResponseFullAuthor struct {
	ID       string `json:"id" api:"required" format:"uuid"`
	FullName string `json:"full_name" api:"required"`
	UserID   string `json:"user_id" api:"required" format:"uuid"`
	Username string `json:"username" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		FullName    respjson.Field
		UserID      respjson.Field
		Username    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3GetUnrelatedResponseFullAuthor) RawJSON() string { return r.JSON.raw }
func (r *PaperV3GetUnrelatedResponseFullAuthor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3GetUnrelatedResponseMetrics struct {
	PublicTotalVotes float64                                       `json:"public_total_votes" api:"required"`
	TotalVotes       float64                                       `json:"total_votes" api:"required"`
	VisitsCount      PaperV3GetUnrelatedResponseMetricsVisitsCount `json:"visits_count" api:"required"`
	XLikes           float64                                       `json:"x_likes" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PublicTotalVotes respjson.Field
		TotalVotes       respjson.Field
		VisitsCount      respjson.Field
		XLikes           respjson.Field
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

type PaperV3KickoffPaperFullTextParams struct {
	// Maximum number of paper versions to process
	MaxPapers param.Opt[float64] `json:"maxPapers,omitzero"`
	paramObj
}

func (r PaperV3KickoffPaperFullTextParams) MarshalJSON() (data []byte, err error) {
	type shadow PaperV3KickoffPaperFullTextParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PaperV3KickoffPaperFullTextParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3KickoffXMentionsSyncParams struct {
	// If true, only logs papers without queuing
	DryRun param.Opt[bool] `json:"dryRun,omitzero"`
	// Number of hot papers to sync (default: 500)
	Limit param.Opt[int64] `json:"limit,omitzero"`
	paramObj
}

func (r PaperV3KickoffXMentionsSyncParams) MarshalJSON() (data []byte, err error) {
	type shadow PaperV3KickoffXMentionsSyncParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PaperV3KickoffXMentionsSyncParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

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

type PaperV3ProcessFullTextParams struct {
	// Paper version ID to process for full text extraction
	PaperVersionID string `json:"paperVersionId" api:"required"`
	paramObj
}

func (r PaperV3ProcessFullTextParams) MarshalJSON() (data []byte, err error) {
	type shadow PaperV3ProcessFullTextParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PaperV3ProcessFullTextParams) UnmarshalJSON(data []byte) error {
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
	// Any of "Hot", "Comments", "Views", "Likes", "GitHub", "Twitter (X)",
	// "Recommended".
	Sort          PaperV3GetFeedParamsSort `query:"sort,omitzero" api:"required" json:"-"`
	Organizations param.Opt[string]        `query:"organizations,omitzero" json:"-"`
	Topics        param.Opt[string]        `query:"topics,omitzero" json:"-"`
	// A versionless universal paper ID (e.g. 1706.03762)
	UniversalID param.Opt[string] `query:"universalId,omitzero" json:"-"`
	// Any of "true", "false".
	ExcludeSeenBriefs PaperV3GetFeedParamsExcludeSeenBriefs `query:"excludeSeenBriefs,omitzero" json:"-"`
	// Any of "true", "false".
	RequireSummary PaperV3GetFeedParamsRequireSummary `query:"requireSummary,omitzero" json:"-"`
	// Any of "GitHub", "Twitter (X)".
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
	PaperV3GetFeedParamsSortTwitterX    PaperV3GetFeedParamsSort = "Twitter (X)"
	PaperV3GetFeedParamsSortRecommended PaperV3GetFeedParamsSort = "Recommended"
)

type PaperV3GetFeedParamsExcludeSeenBriefs string

const (
	PaperV3GetFeedParamsExcludeSeenBriefsTrue  PaperV3GetFeedParamsExcludeSeenBriefs = "true"
	PaperV3GetFeedParamsExcludeSeenBriefsFalse PaperV3GetFeedParamsExcludeSeenBriefs = "false"
)

type PaperV3GetFeedParamsRequireSummary string

const (
	PaperV3GetFeedParamsRequireSummaryTrue  PaperV3GetFeedParamsRequireSummary = "true"
	PaperV3GetFeedParamsRequireSummaryFalse PaperV3GetFeedParamsRequireSummary = "false"
)

type PaperV3GetFeedParamsSource string

const (
	PaperV3GetFeedParamsSourceGitHub   PaperV3GetFeedParamsSource = "GitHub"
	PaperV3GetFeedParamsSourceTwitterX PaperV3GetFeedParamsSource = "Twitter (X)"
)

type PaperV3GetGeoTrendsParams struct {
	CollaborationLimit param.Opt[string] `query:"collaborationLimit,omitzero" json:"-"`
	PaperLimit         param.Opt[string] `query:"paperLimit,omitzero" json:"-"`
	PastMonths         param.Opt[string] `query:"pastMonths,omitzero" json:"-"`
	RepoLimit          param.Opt[string] `query:"repoLimit,omitzero" json:"-"`
	TopCountries       param.Opt[string] `query:"topCountries,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PaperV3GetGeoTrendsParams]'s query parameters as
// `url.Values`.
func (r PaperV3GetGeoTrendsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PaperV3GetPapersByCountryParams struct {
	Country param.Opt[string] `query:"country,omitzero" json:"-"`
	Limit   param.Opt[string] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PaperV3GetPapersByCountryParams]'s query parameters as
// `url.Values`.
func (r PaperV3GetPapersByCountryParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

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
