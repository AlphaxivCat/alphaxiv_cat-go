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
	"github.com/stainless-sdks/alphaxiv_cat-go/internal/apiquery"
	"github.com/stainless-sdks/alphaxiv_cat-go/internal/requestconfig"
	"github.com/stainless-sdks/alphaxiv_cat-go/option"
	"github.com/stainless-sdks/alphaxiv_cat-go/packages/param"
	"github.com/stainless-sdks/alphaxiv_cat-go/packages/respjson"
)

// PaperService contains methods and other services that help with interacting with
// the alphaxiv_cat API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPaperService] method instead.
type PaperService struct {
	options                 []option.RequestOption
	Versions                PaperVersionService
	Metadata                PaperMetadataService
	Ingest                  PaperIngestService
	Private                 PaperPrivateService
	KickoffDailyGitHubStars PaperKickoffDailyGitHubStarService
	V3                      PaperV3Service
	V2                      PaperV2Service
}

// NewPaperService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewPaperService(opts ...option.RequestOption) (r PaperService) {
	r = PaperService{}
	r.options = opts
	r.Versions = NewPaperVersionService(opts...)
	r.Metadata = NewPaperMetadataService(opts...)
	r.Ingest = NewPaperIngestService(opts...)
	r.Private = NewPaperPrivateService(opts...)
	r.KickoffDailyGitHubStars = NewPaperKickoffDailyGitHubStarService(opts...)
	r.V3 = NewPaperV3Service(opts...)
	r.V2 = NewPaperV2Service(opts...)
	return
}

// Add a new author to a paper
//
// Source file: `api-server/src/controllers/v2/papers/add-new-author.controller.ts`
func (r *PaperService) AddAuthor(ctx context.Context, paperID string, body PaperAddAuthorParams, opts ...option.RequestOption) (res *PaperAddAuthorResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if paperID == "" {
		err = errors.New("missing required paperId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/papers/%s/add-author", url.PathEscape(paperID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Set paper vote count (admin only)
//
// Source file:
// `api-server/src/controllers/v2/papers/admin-vote-paper.controller.ts`
func (r *PaperService) AdminVote(ctx context.Context, paperID string, body PaperAdminVoteParams, opts ...option.RequestOption) (res *PaperAdminVoteResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if paperID == "" {
		err = errors.New("missing required paperId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/papers/%s/admin-vote", url.PathEscape(paperID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Legacy route for v1 browser extensions to track abstract page clicks
//
// Source file:
// `api-server/src/controllers/v1/papers/crxabstractclick.controller.ts`
func (r *PaperService) CrxAbstractClick(ctx context.Context, ref string, query PaperCrxAbstractClickParams, opts ...option.RequestOption) (res *PaperCrxAbstractClickResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if query.Pid == "" {
		err = errors.New("missing required pid parameter")
		return nil, err
	}
	if ref == "" {
		err = errors.New("missing required ref parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/papers/crxabstractclick/%s/%s", url.PathEscape(query.Pid), url.PathEscape(ref))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Legacy route for v1 browser extensions to track abstract page hits
//
// Source file: `api-server/src/controllers/v1/papers/crxabstracthit.controller.ts`
func (r *PaperService) CrxAbstractHit(ctx context.Context, pid string, opts ...option.RequestOption) (res *PaperCrxAbstractHitResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if pid == "" {
		err = errors.New("missing required pid parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/papers/crxabstracthit/%s", url.PathEscape(pid))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Legacy route for v1 browser extensions to track PDF page clicks
//
// Source file: `api-server/src/controllers/v1/papers/crxpdfclick.controller.ts`
func (r *PaperService) CrxPdfClick(ctx context.Context, ref string, query PaperCrxPdfClickParams, opts ...option.RequestOption) (res *PaperCrxPdfClickResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if query.Pid == "" {
		err = errors.New("missing required pid parameter")
		return nil, err
	}
	if ref == "" {
		err = errors.New("missing required ref parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/papers/crxpdfclick/%s/%s", url.PathEscape(query.Pid), url.PathEscape(ref))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Legacy route for v1 browser extensions to track PDF page hits
//
// Source file: `api-server/src/controllers/v1/papers/crxpdfhit.controller.ts`
func (r *PaperService) CrxPdfHit(ctx context.Context, pid string, opts ...option.RequestOption) (res *PaperCrxPdfHitResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if pid == "" {
		err = errors.New("missing required pid parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/papers/crxpdfhit/%s", url.PathEscape(pid))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Send email to individual author about paper comments or trending
//
// Source file:
// `api-server/src/controllers/v2/papers/email-individual-author.controller.ts`
func (r *PaperService) EmailAuthor(ctx context.Context, paperID string, body PaperEmailAuthorParams, opts ...option.RequestOption) (res *PaperEmailAuthorResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if paperID == "" {
		err = errors.New("missing required paperId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/papers/%s/email-author", url.PathEscape(paperID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Legacy route for v1 browser extensions to get paper information
//
// Source file:
// `api-server/src/controllers/v1/papers/getcrxpaperinfo.controller.ts`
func (r *PaperService) GetCrxPaperInfo(ctx context.Context, pid string, opts ...option.RequestOption) (res *PaperGetCrxPaperInfoResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if pid == "" {
		err = errors.New("missing required pid parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/papers/getcrxpaperinfo/%s", url.PathEscape(pid))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Legacy route for getting paper information from arXiv abstract pages
//
// Source file: `api-server/src/controllers/v1/papers/getpaperinfo.controller.ts`
func (r *PaperService) GetPaperInfo(ctx context.Context, pid string, opts ...option.RequestOption) (res *PaperGetPaperInfoResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if pid == "" {
		err = errors.New("missing required pid parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/papers/getpaperinfo/%s", url.PathEscape(pid))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Kickoff background job to generate abstract embeddings for paper versions
//
// Source file:
// `api-server/src/controllers/v2/papers/kickoff-paper-version-abstract-embed.controller.ts`
func (r *PaperService) KickoffAbstractEmbed(ctx context.Context, opts ...option.RequestOption) (res *PaperKickoffAbstractEmbedResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v2/papers/kickoff-paper-version-abstract-embed"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Kickoff background job to generate AI overviews for papers
//
// Source file:
// `api-server/src/controllers/v2/papers/kickoff-paper-ai.controller.ts`
func (r *PaperService) KickoffAI(ctx context.Context, opts ...option.RequestOption) (res *PaperKickoffAIResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v2/papers/kickoff-paper-ai"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Kickoff background job to generate bibtex for papers
//
// Source file:
// `api-server/src/controllers/v2/papers/kickoff-paper-bibtex.controller.ts`
func (r *PaperService) KickoffBibtex(ctx context.Context, opts ...option.RequestOption) (res *PaperKickoffBibtexResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v2/papers/kickoff-paper-bibtex"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Kickoff background job to link papers with GitHub repositories
//
// Source file: `api-server/src/controllers/v2/papers/kickoff-github.controller.ts`
func (r *PaperService) KickoffGitHub(ctx context.Context, opts ...option.RequestOption) (res *PaperKickoffGitHubResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v2/papers/kickoff-github"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Kickoff background job to categorize papers
//
// Source file:
// `api-server/src/controllers/v2/papers/kickoff-paper-categorization.controller.ts`
func (r *PaperService) KickoffPaperCategorization(ctx context.Context, all string, opts ...option.RequestOption) (res *PaperKickoffPaperCategorizationResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if all == "" {
		err = errors.New("missing required all parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/papers/kickoff-paper-categorization/%s", url.PathEscape(all))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Kickoff background job to ingest recent papers from arXiv
//
// Source file:
// `api-server/src/controllers/v2/papers/kickoff-recent-papers.controller.ts`
func (r *PaperService) KickoffRecentPapers(ctx context.Context, opts ...option.RequestOption) (res *PaperKickoffRecentPapersResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v2/papers/kickoff-recent-papers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Track paper view event for analytics
//
// Source file:
// `api-server/src/controllers/v2/papers/mark-paper-view.controller.ts`
func (r *PaperService) MarkViewed(ctx context.Context, upid string, opts ...option.RequestOption) (res *PaperMarkViewedResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if upid == "" {
		err = errors.New("missing required upid parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/papers/%s/view", url.PathEscape(upid))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Process abstract embedding for a paper version
//
// Source file:
// `api-server/src/controllers/v2/papers/process-paper-version-abstract-embed.controller.ts`
//
// Deprecated: deprecated
func (r *PaperService) ProcessAbstractEmbed(ctx context.Context, body PaperProcessAbstractEmbedParams, opts ...option.RequestOption) (res *PaperProcessAbstractEmbedResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v2/papers/process-paper-version-abstract-embed"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Process various metadata for a paper (thumbnail, github, bibtex, etc.)
//
// Source file:
// `api-server/src/controllers/v2/papers/process-metadata.controller.ts`
func (r *PaperService) ProcessMetadata(ctx context.Context, body PaperProcessMetadataParams, opts ...option.RequestOption) (res *PaperProcessMetadataResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v2/papers/process-metadata"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Request AI overview generation for the latest paper version
//
// Source file:
// `api-server/src/controllers/v2/papers/request-ai-latest.controller.ts`
func (r *PaperService) RequestAILatest(ctx context.Context, upid string, body PaperRequestAILatestParams, opts ...option.RequestOption) (res *PaperRequestAILatestResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if upid == "" {
		err = errors.New("missing required upid parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/papers/%s/request-ai", url.PathEscape(upid))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Request AI overview translation for the latest paper version
//
// Source file:
// `api-server/src/controllers/v2/papers/request-ai-translation-latest.controller.ts`
func (r *PaperService) RequestAITranslationLatest(ctx context.Context, language PaperRequestAITranslationLatestParamsLanguage, body PaperRequestAITranslationLatestParams, opts ...option.RequestOption) (res *PaperRequestAITranslationLatestResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if body.Upid == "" {
		err = errors.New("missing required upid parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/papers/%s/request-ai-translation/%v", url.PathEscape(body.Upid), language)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Set GitHub repository for a paper
//
// Source file: `api-server/src/controllers/v2/papers/set-github.controller.ts`
func (r *PaperService) SetGitHubRepository(ctx context.Context, paperID string, body PaperSetGitHubRepositoryParams, opts ...option.RequestOption) (res *PaperSetGitHubRepositoryResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if paperID == "" {
		err = errors.New("missing required paperId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/papers/%s/github", url.PathEscape(paperID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Toggle paper follow status (add to Want to read folder or remove from all
// folders)
//
// Source file: `api-server/src/controllers/v2/papers/follow-paper.controller.ts`
func (r *PaperService) ToggleFollow(ctx context.Context, paperID string, opts ...option.RequestOption) (res *PaperToggleFollowResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if paperID == "" {
		err = errors.New("missing required paperId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/papers/%s/follow", url.PathEscape(paperID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Translate AI overview to specified language
//
// Source file:
// `api-server/src/controllers/v2/papers/translate-ai-overview.controller.ts`
//
// Deprecated: deprecated
func (r *PaperService) TranslateAIOverview(ctx context.Context, language PaperTranslateAIOverviewParamsLanguage, body PaperTranslateAIOverviewParams, opts ...option.RequestOption) (res *PaperTranslateAIOverviewResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if body.PaperVersionID == "" {
		err = errors.New("missing required paperVersionId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/papers/translate-ai-overview/%s/%v", body.PaperVersionID, language)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Remove authorship claim from a paper
//
// Source file: `api-server/src/controllers/v2/papers/unclaim-paper.controller.ts`
func (r *PaperService) Unclaim(ctx context.Context, paperID string, opts ...option.RequestOption) (res *PaperUnclaimResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if paperID == "" {
		err = errors.New("missing required paperId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/papers/%s/unclaim", url.PathEscape(paperID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Toggle vote for a paper
//
// Source file: `api-server/src/controllers/v2/papers/vote-paper.controller.ts`
func (r *PaperService) Vote(ctx context.Context, paperID string, opts ...option.RequestOption) (res *PaperVoteResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if paperID == "" {
		err = errors.New("missing required paperId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/papers/%s/vote", url.PathEscape(paperID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

type PaperAddAuthorResponse struct {
	// Any of 0, 1, 2.
	Accept float64                    `json:"accept" api:"required"`
	Data   PaperAddAuthorResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Accept      respjson.Field
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperAddAuthorResponse) RawJSON() string { return r.JSON.raw }
func (r *PaperAddAuthorResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperAddAuthorResponseData struct {
	Warning string `json:"warning" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Warning     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperAddAuthorResponseData) RawJSON() string { return r.JSON.raw }
func (r *PaperAddAuthorResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperAdminVoteResponse struct {
	Data PaperAdminVoteResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperAdminVoteResponse) RawJSON() string { return r.JSON.raw }
func (r *PaperAdminVoteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperAdminVoteResponseData struct {
	Success bool `json:"success" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperAdminVoteResponseData) RawJSON() string { return r.JSON.raw }
func (r *PaperAdminVoteResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperCrxAbstractClickResponse struct {
	Message float64 `json:"message" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperCrxAbstractClickResponse) RawJSON() string { return r.JSON.raw }
func (r *PaperCrxAbstractClickResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperCrxAbstractHitResponse struct {
	Message float64 `json:"message" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperCrxAbstractHitResponse) RawJSON() string { return r.JSON.raw }
func (r *PaperCrxAbstractHitResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperCrxPdfClickResponse struct {
	Message float64 `json:"message" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperCrxPdfClickResponse) RawJSON() string { return r.JSON.raw }
func (r *PaperCrxPdfClickResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperCrxPdfHitResponse struct {
	Message float64 `json:"message" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperCrxPdfHitResponse) RawJSON() string { return r.JSON.raw }
func (r *PaperCrxPdfHitResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperEmailAuthorResponse struct {
	Status float64 `json:"status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperEmailAuthorResponse) RawJSON() string { return r.JSON.raw }
func (r *PaperEmailAuthorResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperGetCrxPaperInfoResponse struct {
	HasClaimedAuthorship bool    `json:"hasClaimedAuthorship" api:"required"`
	NumQuestions         float64 `json:"numQuestions" api:"required"`
	ReturnVersion        float64 `json:"returnVersion" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		HasClaimedAuthorship respjson.Field
		NumQuestions         respjson.Field
		ReturnVersion        respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperGetCrxPaperInfoResponse) RawJSON() string { return r.JSON.raw }
func (r *PaperGetCrxPaperInfoResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperGetPaperInfoResponse struct {
	HasClaimedAuthorship bool    `json:"hasClaimedAuthorship" api:"required"`
	NumQuestions         float64 `json:"numQuestions" api:"required"`
	ReturnVersion        float64 `json:"returnVersion" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		HasClaimedAuthorship respjson.Field
		NumQuestions         respjson.Field
		ReturnVersion        respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperGetPaperInfoResponse) RawJSON() string { return r.JSON.raw }
func (r *PaperGetPaperInfoResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperKickoffAbstractEmbedResponse struct {
	Data PaperKickoffAbstractEmbedResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperKickoffAbstractEmbedResponse) RawJSON() string { return r.JSON.raw }
func (r *PaperKickoffAbstractEmbedResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperKickoffAbstractEmbedResponseData struct {
	Message string `json:"message" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperKickoffAbstractEmbedResponseData) RawJSON() string { return r.JSON.raw }
func (r *PaperKickoffAbstractEmbedResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperKickoffAIResponse struct {
	Data PaperKickoffAIResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperKickoffAIResponse) RawJSON() string { return r.JSON.raw }
func (r *PaperKickoffAIResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperKickoffAIResponseData struct {
	Message string `json:"message" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperKickoffAIResponseData) RawJSON() string { return r.JSON.raw }
func (r *PaperKickoffAIResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperKickoffBibtexResponse struct {
	Data PaperKickoffBibtexResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperKickoffBibtexResponse) RawJSON() string { return r.JSON.raw }
func (r *PaperKickoffBibtexResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperKickoffBibtexResponseData struct {
	Message string `json:"message" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperKickoffBibtexResponseData) RawJSON() string { return r.JSON.raw }
func (r *PaperKickoffBibtexResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperKickoffGitHubResponse struct {
	Data PaperKickoffGitHubResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperKickoffGitHubResponse) RawJSON() string { return r.JSON.raw }
func (r *PaperKickoffGitHubResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperKickoffGitHubResponseData struct {
	Message string `json:"message" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperKickoffGitHubResponseData) RawJSON() string { return r.JSON.raw }
func (r *PaperKickoffGitHubResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperKickoffPaperCategorizationResponse struct {
	Data PaperKickoffPaperCategorizationResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperKickoffPaperCategorizationResponse) RawJSON() string { return r.JSON.raw }
func (r *PaperKickoffPaperCategorizationResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperKickoffPaperCategorizationResponseData struct {
	Message string `json:"message" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperKickoffPaperCategorizationResponseData) RawJSON() string { return r.JSON.raw }
func (r *PaperKickoffPaperCategorizationResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperKickoffRecentPapersResponse struct {
	Data PaperKickoffRecentPapersResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperKickoffRecentPapersResponse) RawJSON() string { return r.JSON.raw }
func (r *PaperKickoffRecentPapersResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperKickoffRecentPapersResponseData struct {
	Message string `json:"message" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperKickoffRecentPapersResponseData) RawJSON() string { return r.JSON.raw }
func (r *PaperKickoffRecentPapersResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperMarkViewedResponse = any

type PaperProcessAbstractEmbedResponse struct {
	Data PaperProcessAbstractEmbedResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperProcessAbstractEmbedResponse) RawJSON() string { return r.JSON.raw }
func (r *PaperProcessAbstractEmbedResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperProcessAbstractEmbedResponseData struct {
	Status string `json:"status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperProcessAbstractEmbedResponseData) RawJSON() string { return r.JSON.raw }
func (r *PaperProcessAbstractEmbedResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperProcessMetadataResponse struct {
	Data  PaperProcessMetadataResponseData `json:"data"`
	Error string                           `json:"error"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Error       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperProcessMetadataResponse) RawJSON() string { return r.JSON.raw }
func (r *PaperProcessMetadataResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperProcessMetadataResponseData struct {
	Status string `json:"status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperProcessMetadataResponseData) RawJSON() string { return r.JSON.raw }
func (r *PaperProcessMetadataResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperRequestAILatestResponse struct {
	Data string `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperRequestAILatestResponse) RawJSON() string { return r.JSON.raw }
func (r *PaperRequestAILatestResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperRequestAITranslationLatestResponse struct {
	Data string `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperRequestAITranslationLatestResponse) RawJSON() string { return r.JSON.raw }
func (r *PaperRequestAITranslationLatestResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperSetGitHubRepositoryResponse struct {
	Data PaperSetGitHubRepositoryResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperSetGitHubRepositoryResponse) RawJSON() string { return r.JSON.raw }
func (r *PaperSetGitHubRepositoryResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperSetGitHubRepositoryResponseData struct {
	Message string `json:"message" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperSetGitHubRepositoryResponseData) RawJSON() string { return r.JSON.raw }
func (r *PaperSetGitHubRepositoryResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperToggleFollowResponse struct {
	Data PaperToggleFollowResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperToggleFollowResponse) RawJSON() string { return r.JSON.raw }
func (r *PaperToggleFollowResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperToggleFollowResponseData struct {
	PaperGroupID string `json:"paperGroupId" api:"required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PaperGroupID respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperToggleFollowResponseData) RawJSON() string { return r.JSON.raw }
func (r *PaperToggleFollowResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperTranslateAIOverviewResponse struct {
	Data PaperTranslateAIOverviewResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperTranslateAIOverviewResponse) RawJSON() string { return r.JSON.raw }
func (r *PaperTranslateAIOverviewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperTranslateAIOverviewResponseData struct {
	Message string `json:"message" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperTranslateAIOverviewResponseData) RawJSON() string { return r.JSON.raw }
func (r *PaperTranslateAIOverviewResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperUnclaimResponse struct {
	Data PaperUnclaimResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperUnclaimResponse) RawJSON() string { return r.JSON.raw }
func (r *PaperUnclaimResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperUnclaimResponseData struct {
	Success bool `json:"success" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperUnclaimResponseData) RawJSON() string { return r.JSON.raw }
func (r *PaperUnclaimResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperVoteResponse struct {
	Data PaperVoteResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperVoteResponse) RawJSON() string { return r.JSON.raw }
func (r *PaperVoteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperVoteResponseData struct {
	PaperGroup PaperVoteResponseDataPaperGroup `json:"paperGroup" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PaperGroup  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperVoteResponseData) RawJSON() string { return r.JSON.raw }
func (r *PaperVoteResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperVoteResponseDataPaperGroup struct {
	ID string `json:"_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperVoteResponseDataPaperGroup) RawJSON() string { return r.JSON.raw }
func (r *PaperVoteResponseDataPaperGroup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperAddAuthorParams struct {
	AuthorEmail string          `json:"authorEmail" api:"required" format:"email"`
	ShouldEmail param.Opt[bool] `json:"shouldEmail,omitzero"`
	paramObj
}

func (r PaperAddAuthorParams) MarshalJSON() (data []byte, err error) {
	type shadow PaperAddAuthorParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PaperAddAuthorParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperAdminVoteParams struct {
	Entry float64 `json:"entry" api:"required"`
	paramObj
}

func (r PaperAdminVoteParams) MarshalJSON() (data []byte, err error) {
	type shadow PaperAdminVoteParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PaperAdminVoteParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperCrxAbstractClickParams struct {
	// Paper ID
	Pid string `path:"pid" api:"required" json:"-"`
	paramObj
}

type PaperCrxPdfClickParams struct {
	// Paper ID
	Pid string `path:"pid" api:"required" json:"-"`
	paramObj
}

type PaperEmailAuthorParams struct {
	AuthorIndividualEmail string `json:"authorIndividualEmail" api:"required" format:"email"`
	EmailAuthorName       string `json:"emailAuthorName" api:"required"`
	PaperName             string `json:"paperName" api:"required"`
	// Any of "comment", "trending".
	Type                 PaperEmailAuthorParamsType `json:"type,omitzero" api:"required"`
	IgnoreDuplicateError param.Opt[bool]            `json:"ignoreDuplicateError,omitzero"`
	paramObj
}

func (r PaperEmailAuthorParams) MarshalJSON() (data []byte, err error) {
	type shadow PaperEmailAuthorParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PaperEmailAuthorParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperEmailAuthorParamsType string

const (
	PaperEmailAuthorParamsTypeComment  PaperEmailAuthorParamsType = "comment"
	PaperEmailAuthorParamsTypeTrending PaperEmailAuthorParamsType = "trending"
)

type PaperProcessAbstractEmbedParams struct {
	PaperVersionID string `json:"paperVersionId" api:"required" format:"uuid"`
	paramObj
}

func (r PaperProcessAbstractEmbedParams) MarshalJSON() (data []byte, err error) {
	type shadow PaperProcessAbstractEmbedParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PaperProcessAbstractEmbedParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperProcessMetadataParams struct {
	Metadata         PaperProcessMetadataParamsMetadata `json:"metadata,omitzero" api:"required"`
	UniversalPaperID string                             `json:"universalPaperId" api:"required"`
	paramObj
}

func (r PaperProcessMetadataParams) MarshalJSON() (data []byte, err error) {
	type shadow PaperProcessMetadataParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PaperProcessMetadataParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperProcessMetadataParamsMetadata struct {
	Bibtex           param.Opt[bool] `json:"bibtex,omitzero"`
	CustomCategories param.Opt[bool] `json:"custom_categories,omitzero"`
	Embedding        param.Opt[bool] `json:"embedding,omitzero"`
	GitHub           param.Opt[bool] `json:"github,omitzero"`
	Organizations    param.Opt[bool] `json:"organizations,omitzero"`
	Overview         param.Opt[bool] `json:"overview,omitzero"`
	Thumbnail        param.Opt[bool] `json:"thumbnail,omitzero"`
	paramObj
}

func (r PaperProcessMetadataParamsMetadata) MarshalJSON() (data []byte, err error) {
	type shadow PaperProcessMetadataParamsMetadata
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PaperProcessMetadataParamsMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperRequestAILatestParams struct {
	// Any of "am", "ar", "az", "bg", "bn", "ca", "cs", "da", "de", "el", "en", "es",
	// "et", "fa", "fi", "fr", "gu", "ha", "he", "hi", "hr", "hu", "id", "it", "ja",
	// "ka", "kn", "ko", "lt", "lv", "ml", "mr", "ms", "my", "ne", "nl", "no", "pa",
	// "pl", "pt", "ro", "ru", "si", "sk", "sl", "sr", "sv", "sw", "ta", "te", "th",
	// "tl", "tr", "uk", "ur", "uz", "vi", "yo", "zh".
	PreferredLanguage PaperRequestAILatestParamsPreferredLanguage `query:"preferredLanguage,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PaperRequestAILatestParams]'s query parameters as
// `url.Values`.
func (r PaperRequestAILatestParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PaperRequestAILatestParamsPreferredLanguage string

const (
	PaperRequestAILatestParamsPreferredLanguageAm PaperRequestAILatestParamsPreferredLanguage = "am"
	PaperRequestAILatestParamsPreferredLanguageAr PaperRequestAILatestParamsPreferredLanguage = "ar"
	PaperRequestAILatestParamsPreferredLanguageAz PaperRequestAILatestParamsPreferredLanguage = "az"
	PaperRequestAILatestParamsPreferredLanguageBg PaperRequestAILatestParamsPreferredLanguage = "bg"
	PaperRequestAILatestParamsPreferredLanguageBn PaperRequestAILatestParamsPreferredLanguage = "bn"
	PaperRequestAILatestParamsPreferredLanguageCa PaperRequestAILatestParamsPreferredLanguage = "ca"
	PaperRequestAILatestParamsPreferredLanguageCs PaperRequestAILatestParamsPreferredLanguage = "cs"
	PaperRequestAILatestParamsPreferredLanguageDa PaperRequestAILatestParamsPreferredLanguage = "da"
	PaperRequestAILatestParamsPreferredLanguageDe PaperRequestAILatestParamsPreferredLanguage = "de"
	PaperRequestAILatestParamsPreferredLanguageEl PaperRequestAILatestParamsPreferredLanguage = "el"
	PaperRequestAILatestParamsPreferredLanguageEn PaperRequestAILatestParamsPreferredLanguage = "en"
	PaperRequestAILatestParamsPreferredLanguageEs PaperRequestAILatestParamsPreferredLanguage = "es"
	PaperRequestAILatestParamsPreferredLanguageEt PaperRequestAILatestParamsPreferredLanguage = "et"
	PaperRequestAILatestParamsPreferredLanguageFa PaperRequestAILatestParamsPreferredLanguage = "fa"
	PaperRequestAILatestParamsPreferredLanguageFi PaperRequestAILatestParamsPreferredLanguage = "fi"
	PaperRequestAILatestParamsPreferredLanguageFr PaperRequestAILatestParamsPreferredLanguage = "fr"
	PaperRequestAILatestParamsPreferredLanguageGu PaperRequestAILatestParamsPreferredLanguage = "gu"
	PaperRequestAILatestParamsPreferredLanguageHa PaperRequestAILatestParamsPreferredLanguage = "ha"
	PaperRequestAILatestParamsPreferredLanguageHe PaperRequestAILatestParamsPreferredLanguage = "he"
	PaperRequestAILatestParamsPreferredLanguageHi PaperRequestAILatestParamsPreferredLanguage = "hi"
	PaperRequestAILatestParamsPreferredLanguageHr PaperRequestAILatestParamsPreferredLanguage = "hr"
	PaperRequestAILatestParamsPreferredLanguageHu PaperRequestAILatestParamsPreferredLanguage = "hu"
	PaperRequestAILatestParamsPreferredLanguageID PaperRequestAILatestParamsPreferredLanguage = "id"
	PaperRequestAILatestParamsPreferredLanguageIt PaperRequestAILatestParamsPreferredLanguage = "it"
	PaperRequestAILatestParamsPreferredLanguageJa PaperRequestAILatestParamsPreferredLanguage = "ja"
	PaperRequestAILatestParamsPreferredLanguageKa PaperRequestAILatestParamsPreferredLanguage = "ka"
	PaperRequestAILatestParamsPreferredLanguageKn PaperRequestAILatestParamsPreferredLanguage = "kn"
	PaperRequestAILatestParamsPreferredLanguageKo PaperRequestAILatestParamsPreferredLanguage = "ko"
	PaperRequestAILatestParamsPreferredLanguageLt PaperRequestAILatestParamsPreferredLanguage = "lt"
	PaperRequestAILatestParamsPreferredLanguageLv PaperRequestAILatestParamsPreferredLanguage = "lv"
	PaperRequestAILatestParamsPreferredLanguageMl PaperRequestAILatestParamsPreferredLanguage = "ml"
	PaperRequestAILatestParamsPreferredLanguageMr PaperRequestAILatestParamsPreferredLanguage = "mr"
	PaperRequestAILatestParamsPreferredLanguageMs PaperRequestAILatestParamsPreferredLanguage = "ms"
	PaperRequestAILatestParamsPreferredLanguageMy PaperRequestAILatestParamsPreferredLanguage = "my"
	PaperRequestAILatestParamsPreferredLanguageNe PaperRequestAILatestParamsPreferredLanguage = "ne"
	PaperRequestAILatestParamsPreferredLanguageNl PaperRequestAILatestParamsPreferredLanguage = "nl"
	PaperRequestAILatestParamsPreferredLanguageNo PaperRequestAILatestParamsPreferredLanguage = "no"
	PaperRequestAILatestParamsPreferredLanguagePa PaperRequestAILatestParamsPreferredLanguage = "pa"
	PaperRequestAILatestParamsPreferredLanguagePl PaperRequestAILatestParamsPreferredLanguage = "pl"
	PaperRequestAILatestParamsPreferredLanguagePt PaperRequestAILatestParamsPreferredLanguage = "pt"
	PaperRequestAILatestParamsPreferredLanguageRo PaperRequestAILatestParamsPreferredLanguage = "ro"
	PaperRequestAILatestParamsPreferredLanguageRu PaperRequestAILatestParamsPreferredLanguage = "ru"
	PaperRequestAILatestParamsPreferredLanguageSi PaperRequestAILatestParamsPreferredLanguage = "si"
	PaperRequestAILatestParamsPreferredLanguageSk PaperRequestAILatestParamsPreferredLanguage = "sk"
	PaperRequestAILatestParamsPreferredLanguageSl PaperRequestAILatestParamsPreferredLanguage = "sl"
	PaperRequestAILatestParamsPreferredLanguageSr PaperRequestAILatestParamsPreferredLanguage = "sr"
	PaperRequestAILatestParamsPreferredLanguageSv PaperRequestAILatestParamsPreferredLanguage = "sv"
	PaperRequestAILatestParamsPreferredLanguageSw PaperRequestAILatestParamsPreferredLanguage = "sw"
	PaperRequestAILatestParamsPreferredLanguageTa PaperRequestAILatestParamsPreferredLanguage = "ta"
	PaperRequestAILatestParamsPreferredLanguageTe PaperRequestAILatestParamsPreferredLanguage = "te"
	PaperRequestAILatestParamsPreferredLanguageTh PaperRequestAILatestParamsPreferredLanguage = "th"
	PaperRequestAILatestParamsPreferredLanguageTl PaperRequestAILatestParamsPreferredLanguage = "tl"
	PaperRequestAILatestParamsPreferredLanguageTr PaperRequestAILatestParamsPreferredLanguage = "tr"
	PaperRequestAILatestParamsPreferredLanguageUk PaperRequestAILatestParamsPreferredLanguage = "uk"
	PaperRequestAILatestParamsPreferredLanguageUr PaperRequestAILatestParamsPreferredLanguage = "ur"
	PaperRequestAILatestParamsPreferredLanguageUz PaperRequestAILatestParamsPreferredLanguage = "uz"
	PaperRequestAILatestParamsPreferredLanguageVi PaperRequestAILatestParamsPreferredLanguage = "vi"
	PaperRequestAILatestParamsPreferredLanguageYo PaperRequestAILatestParamsPreferredLanguage = "yo"
	PaperRequestAILatestParamsPreferredLanguageZh PaperRequestAILatestParamsPreferredLanguage = "zh"
)

type PaperRequestAITranslationLatestParams struct {
	Upid string `path:"upid" api:"required" json:"-"`
	paramObj
}

type PaperRequestAITranslationLatestParamsLanguage string

const (
	PaperRequestAITranslationLatestParamsLanguageAm PaperRequestAITranslationLatestParamsLanguage = "am"
	PaperRequestAITranslationLatestParamsLanguageAr PaperRequestAITranslationLatestParamsLanguage = "ar"
	PaperRequestAITranslationLatestParamsLanguageAz PaperRequestAITranslationLatestParamsLanguage = "az"
	PaperRequestAITranslationLatestParamsLanguageBg PaperRequestAITranslationLatestParamsLanguage = "bg"
	PaperRequestAITranslationLatestParamsLanguageBn PaperRequestAITranslationLatestParamsLanguage = "bn"
	PaperRequestAITranslationLatestParamsLanguageCa PaperRequestAITranslationLatestParamsLanguage = "ca"
	PaperRequestAITranslationLatestParamsLanguageCs PaperRequestAITranslationLatestParamsLanguage = "cs"
	PaperRequestAITranslationLatestParamsLanguageDa PaperRequestAITranslationLatestParamsLanguage = "da"
	PaperRequestAITranslationLatestParamsLanguageDe PaperRequestAITranslationLatestParamsLanguage = "de"
	PaperRequestAITranslationLatestParamsLanguageEl PaperRequestAITranslationLatestParamsLanguage = "el"
	PaperRequestAITranslationLatestParamsLanguageEs PaperRequestAITranslationLatestParamsLanguage = "es"
	PaperRequestAITranslationLatestParamsLanguageEt PaperRequestAITranslationLatestParamsLanguage = "et"
	PaperRequestAITranslationLatestParamsLanguageFa PaperRequestAITranslationLatestParamsLanguage = "fa"
	PaperRequestAITranslationLatestParamsLanguageFi PaperRequestAITranslationLatestParamsLanguage = "fi"
	PaperRequestAITranslationLatestParamsLanguageFr PaperRequestAITranslationLatestParamsLanguage = "fr"
	PaperRequestAITranslationLatestParamsLanguageGu PaperRequestAITranslationLatestParamsLanguage = "gu"
	PaperRequestAITranslationLatestParamsLanguageHa PaperRequestAITranslationLatestParamsLanguage = "ha"
	PaperRequestAITranslationLatestParamsLanguageHe PaperRequestAITranslationLatestParamsLanguage = "he"
	PaperRequestAITranslationLatestParamsLanguageHi PaperRequestAITranslationLatestParamsLanguage = "hi"
	PaperRequestAITranslationLatestParamsLanguageHr PaperRequestAITranslationLatestParamsLanguage = "hr"
	PaperRequestAITranslationLatestParamsLanguageHu PaperRequestAITranslationLatestParamsLanguage = "hu"
	PaperRequestAITranslationLatestParamsLanguageID PaperRequestAITranslationLatestParamsLanguage = "id"
	PaperRequestAITranslationLatestParamsLanguageIt PaperRequestAITranslationLatestParamsLanguage = "it"
	PaperRequestAITranslationLatestParamsLanguageJa PaperRequestAITranslationLatestParamsLanguage = "ja"
	PaperRequestAITranslationLatestParamsLanguageKa PaperRequestAITranslationLatestParamsLanguage = "ka"
	PaperRequestAITranslationLatestParamsLanguageKn PaperRequestAITranslationLatestParamsLanguage = "kn"
	PaperRequestAITranslationLatestParamsLanguageKo PaperRequestAITranslationLatestParamsLanguage = "ko"
	PaperRequestAITranslationLatestParamsLanguageLt PaperRequestAITranslationLatestParamsLanguage = "lt"
	PaperRequestAITranslationLatestParamsLanguageLv PaperRequestAITranslationLatestParamsLanguage = "lv"
	PaperRequestAITranslationLatestParamsLanguageMl PaperRequestAITranslationLatestParamsLanguage = "ml"
	PaperRequestAITranslationLatestParamsLanguageMr PaperRequestAITranslationLatestParamsLanguage = "mr"
	PaperRequestAITranslationLatestParamsLanguageMs PaperRequestAITranslationLatestParamsLanguage = "ms"
	PaperRequestAITranslationLatestParamsLanguageMy PaperRequestAITranslationLatestParamsLanguage = "my"
	PaperRequestAITranslationLatestParamsLanguageNe PaperRequestAITranslationLatestParamsLanguage = "ne"
	PaperRequestAITranslationLatestParamsLanguageNl PaperRequestAITranslationLatestParamsLanguage = "nl"
	PaperRequestAITranslationLatestParamsLanguageNo PaperRequestAITranslationLatestParamsLanguage = "no"
	PaperRequestAITranslationLatestParamsLanguagePa PaperRequestAITranslationLatestParamsLanguage = "pa"
	PaperRequestAITranslationLatestParamsLanguagePl PaperRequestAITranslationLatestParamsLanguage = "pl"
	PaperRequestAITranslationLatestParamsLanguagePt PaperRequestAITranslationLatestParamsLanguage = "pt"
	PaperRequestAITranslationLatestParamsLanguageRo PaperRequestAITranslationLatestParamsLanguage = "ro"
	PaperRequestAITranslationLatestParamsLanguageRu PaperRequestAITranslationLatestParamsLanguage = "ru"
	PaperRequestAITranslationLatestParamsLanguageSi PaperRequestAITranslationLatestParamsLanguage = "si"
	PaperRequestAITranslationLatestParamsLanguageSk PaperRequestAITranslationLatestParamsLanguage = "sk"
	PaperRequestAITranslationLatestParamsLanguageSl PaperRequestAITranslationLatestParamsLanguage = "sl"
	PaperRequestAITranslationLatestParamsLanguageSr PaperRequestAITranslationLatestParamsLanguage = "sr"
	PaperRequestAITranslationLatestParamsLanguageSv PaperRequestAITranslationLatestParamsLanguage = "sv"
	PaperRequestAITranslationLatestParamsLanguageSw PaperRequestAITranslationLatestParamsLanguage = "sw"
	PaperRequestAITranslationLatestParamsLanguageTa PaperRequestAITranslationLatestParamsLanguage = "ta"
	PaperRequestAITranslationLatestParamsLanguageTe PaperRequestAITranslationLatestParamsLanguage = "te"
	PaperRequestAITranslationLatestParamsLanguageTh PaperRequestAITranslationLatestParamsLanguage = "th"
	PaperRequestAITranslationLatestParamsLanguageTl PaperRequestAITranslationLatestParamsLanguage = "tl"
	PaperRequestAITranslationLatestParamsLanguageTr PaperRequestAITranslationLatestParamsLanguage = "tr"
	PaperRequestAITranslationLatestParamsLanguageUk PaperRequestAITranslationLatestParamsLanguage = "uk"
	PaperRequestAITranslationLatestParamsLanguageUr PaperRequestAITranslationLatestParamsLanguage = "ur"
	PaperRequestAITranslationLatestParamsLanguageUz PaperRequestAITranslationLatestParamsLanguage = "uz"
	PaperRequestAITranslationLatestParamsLanguageVi PaperRequestAITranslationLatestParamsLanguage = "vi"
	PaperRequestAITranslationLatestParamsLanguageYo PaperRequestAITranslationLatestParamsLanguage = "yo"
	PaperRequestAITranslationLatestParamsLanguageZh PaperRequestAITranslationLatestParamsLanguage = "zh"
)

type PaperSetGitHubRepositoryParams struct {
	GitHub string `json:"github" api:"required" format:"uri"`
	paramObj
}

func (r PaperSetGitHubRepositoryParams) MarshalJSON() (data []byte, err error) {
	type shadow PaperSetGitHubRepositoryParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PaperSetGitHubRepositoryParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperTranslateAIOverviewParams struct {
	PaperVersionID string `path:"paperVersionId" api:"required" format:"uuid" json:"-"`
	paramObj
}

type PaperTranslateAIOverviewParamsLanguage string

const (
	PaperTranslateAIOverviewParamsLanguageAm PaperTranslateAIOverviewParamsLanguage = "am"
	PaperTranslateAIOverviewParamsLanguageAr PaperTranslateAIOverviewParamsLanguage = "ar"
	PaperTranslateAIOverviewParamsLanguageAz PaperTranslateAIOverviewParamsLanguage = "az"
	PaperTranslateAIOverviewParamsLanguageBg PaperTranslateAIOverviewParamsLanguage = "bg"
	PaperTranslateAIOverviewParamsLanguageBn PaperTranslateAIOverviewParamsLanguage = "bn"
	PaperTranslateAIOverviewParamsLanguageCa PaperTranslateAIOverviewParamsLanguage = "ca"
	PaperTranslateAIOverviewParamsLanguageCs PaperTranslateAIOverviewParamsLanguage = "cs"
	PaperTranslateAIOverviewParamsLanguageDa PaperTranslateAIOverviewParamsLanguage = "da"
	PaperTranslateAIOverviewParamsLanguageDe PaperTranslateAIOverviewParamsLanguage = "de"
	PaperTranslateAIOverviewParamsLanguageEl PaperTranslateAIOverviewParamsLanguage = "el"
	PaperTranslateAIOverviewParamsLanguageEs PaperTranslateAIOverviewParamsLanguage = "es"
	PaperTranslateAIOverviewParamsLanguageEt PaperTranslateAIOverviewParamsLanguage = "et"
	PaperTranslateAIOverviewParamsLanguageFa PaperTranslateAIOverviewParamsLanguage = "fa"
	PaperTranslateAIOverviewParamsLanguageFi PaperTranslateAIOverviewParamsLanguage = "fi"
	PaperTranslateAIOverviewParamsLanguageFr PaperTranslateAIOverviewParamsLanguage = "fr"
	PaperTranslateAIOverviewParamsLanguageGu PaperTranslateAIOverviewParamsLanguage = "gu"
	PaperTranslateAIOverviewParamsLanguageHa PaperTranslateAIOverviewParamsLanguage = "ha"
	PaperTranslateAIOverviewParamsLanguageHe PaperTranslateAIOverviewParamsLanguage = "he"
	PaperTranslateAIOverviewParamsLanguageHi PaperTranslateAIOverviewParamsLanguage = "hi"
	PaperTranslateAIOverviewParamsLanguageHr PaperTranslateAIOverviewParamsLanguage = "hr"
	PaperTranslateAIOverviewParamsLanguageHu PaperTranslateAIOverviewParamsLanguage = "hu"
	PaperTranslateAIOverviewParamsLanguageID PaperTranslateAIOverviewParamsLanguage = "id"
	PaperTranslateAIOverviewParamsLanguageIt PaperTranslateAIOverviewParamsLanguage = "it"
	PaperTranslateAIOverviewParamsLanguageJa PaperTranslateAIOverviewParamsLanguage = "ja"
	PaperTranslateAIOverviewParamsLanguageKa PaperTranslateAIOverviewParamsLanguage = "ka"
	PaperTranslateAIOverviewParamsLanguageKn PaperTranslateAIOverviewParamsLanguage = "kn"
	PaperTranslateAIOverviewParamsLanguageKo PaperTranslateAIOverviewParamsLanguage = "ko"
	PaperTranslateAIOverviewParamsLanguageLt PaperTranslateAIOverviewParamsLanguage = "lt"
	PaperTranslateAIOverviewParamsLanguageLv PaperTranslateAIOverviewParamsLanguage = "lv"
	PaperTranslateAIOverviewParamsLanguageMl PaperTranslateAIOverviewParamsLanguage = "ml"
	PaperTranslateAIOverviewParamsLanguageMr PaperTranslateAIOverviewParamsLanguage = "mr"
	PaperTranslateAIOverviewParamsLanguageMs PaperTranslateAIOverviewParamsLanguage = "ms"
	PaperTranslateAIOverviewParamsLanguageMy PaperTranslateAIOverviewParamsLanguage = "my"
	PaperTranslateAIOverviewParamsLanguageNe PaperTranslateAIOverviewParamsLanguage = "ne"
	PaperTranslateAIOverviewParamsLanguageNl PaperTranslateAIOverviewParamsLanguage = "nl"
	PaperTranslateAIOverviewParamsLanguageNo PaperTranslateAIOverviewParamsLanguage = "no"
	PaperTranslateAIOverviewParamsLanguagePa PaperTranslateAIOverviewParamsLanguage = "pa"
	PaperTranslateAIOverviewParamsLanguagePl PaperTranslateAIOverviewParamsLanguage = "pl"
	PaperTranslateAIOverviewParamsLanguagePt PaperTranslateAIOverviewParamsLanguage = "pt"
	PaperTranslateAIOverviewParamsLanguageRo PaperTranslateAIOverviewParamsLanguage = "ro"
	PaperTranslateAIOverviewParamsLanguageRu PaperTranslateAIOverviewParamsLanguage = "ru"
	PaperTranslateAIOverviewParamsLanguageSi PaperTranslateAIOverviewParamsLanguage = "si"
	PaperTranslateAIOverviewParamsLanguageSk PaperTranslateAIOverviewParamsLanguage = "sk"
	PaperTranslateAIOverviewParamsLanguageSl PaperTranslateAIOverviewParamsLanguage = "sl"
	PaperTranslateAIOverviewParamsLanguageSr PaperTranslateAIOverviewParamsLanguage = "sr"
	PaperTranslateAIOverviewParamsLanguageSv PaperTranslateAIOverviewParamsLanguage = "sv"
	PaperTranslateAIOverviewParamsLanguageSw PaperTranslateAIOverviewParamsLanguage = "sw"
	PaperTranslateAIOverviewParamsLanguageTa PaperTranslateAIOverviewParamsLanguage = "ta"
	PaperTranslateAIOverviewParamsLanguageTe PaperTranslateAIOverviewParamsLanguage = "te"
	PaperTranslateAIOverviewParamsLanguageTh PaperTranslateAIOverviewParamsLanguage = "th"
	PaperTranslateAIOverviewParamsLanguageTl PaperTranslateAIOverviewParamsLanguage = "tl"
	PaperTranslateAIOverviewParamsLanguageTr PaperTranslateAIOverviewParamsLanguage = "tr"
	PaperTranslateAIOverviewParamsLanguageUk PaperTranslateAIOverviewParamsLanguage = "uk"
	PaperTranslateAIOverviewParamsLanguageUr PaperTranslateAIOverviewParamsLanguage = "ur"
	PaperTranslateAIOverviewParamsLanguageUz PaperTranslateAIOverviewParamsLanguage = "uz"
	PaperTranslateAIOverviewParamsLanguageVi PaperTranslateAIOverviewParamsLanguage = "vi"
	PaperTranslateAIOverviewParamsLanguageYo PaperTranslateAIOverviewParamsLanguage = "yo"
	PaperTranslateAIOverviewParamsLanguageZh PaperTranslateAIOverviewParamsLanguage = "zh"
)
