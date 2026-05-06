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
	"github.com/AlphaxivCat/alphaxiv_cat-go/internal/requestconfig"
	"github.com/AlphaxivCat/alphaxiv_cat-go/option"
	"github.com/AlphaxivCat/alphaxiv_cat-go/packages/respjson"
)

// PaperV3LegacyService contains methods and other services that help with
// interacting with the alphaxiv_cat API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPaperV3LegacyService] method instead.
type PaperV3LegacyService struct {
	options []option.RequestOption
}

// NewPaperV3LegacyService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPaperV3LegacyService(opts ...option.RequestOption) (r PaperV3LegacyService) {
	r = PaperV3LegacyService{}
	r.options = opts
	return
}

// Retrieve paper version metadata and comments. Fetches from ArXiv if needed. This
// runs the old id resolution implementation, old fetching service, and old JSON
// format. Do not write new code with this please.
//
// Source file:
// `api-server/src/controllers/papers/v3/legacy/get-v2-paper-for-display.controller.ts`
func (r *PaperV3LegacyService) Get(ctx context.Context, unresolved string, opts ...option.RequestOption) (res *PaperV3LegacyGetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if unresolved == "" {
		err = errors.New("missing required unresolved parameter")
		return nil, err
	}
	path := fmt.Sprintf("papers/v3/legacy/%s", url.PathEscape(unresolved))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Retrieve paper group comments. Does not distinguish "paper group does not exist"
// from "there are no comments"
//
// Source file:
// `api-server/src/controllers/papers/v3/legacy/get-v2-comments.controller.ts`
func (r *PaperV3LegacyService) GetComments(ctx context.Context, group string, opts ...option.RequestOption) (res *[]PaperV3LegacyGetCommentsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if group == "" {
		err = errors.New("missing required group parameter")
		return nil, err
	}
	path := fmt.Sprintf("papers/v3/legacy/%s/comments", group)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type PaperV3LegacyGetResponse struct {
	Comments []PaperV3LegacyGetResponseComment `json:"comments" api:"required"`
	Paper    any                               `json:"paper"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Comments    respjson.Field
		Paper       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3LegacyGetResponse) RawJSON() string { return r.JSON.raw }
func (r *PaperV3LegacyGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3LegacyGetResponseComment struct {
	ID              string                                       `json:"id" api:"required" format:"uuid"`
	Annotation      PaperV3LegacyGetResponseCommentAnnotation    `json:"annotation" api:"required"`
	Author          PaperV3LegacyGetResponseCommentAuthor        `json:"author" api:"required"`
	AuthorResponded bool                                         `json:"authorResponded" api:"required"`
	Body            string                                       `json:"body" api:"required"`
	Date            string                                       `json:"date" api:"required"`
	Endorsements    []PaperV3LegacyGetResponseCommentEndorsement `json:"endorsements" api:"required"`
	HasDownvoted    bool                                         `json:"hasDownvoted" api:"required"`
	HasFlagged      bool                                         `json:"hasFlagged" api:"required"`
	HasUpvoted      bool                                         `json:"hasUpvoted" api:"required"`
	IsAuthor        bool                                         `json:"isAuthor" api:"required"`
	PaperGroupID    string                                       `json:"paperGroupId" api:"required" format:"uuid"`
	PaperTitle      string                                       `json:"paperTitle" api:"required"`
	PaperVersionID  string                                       `json:"paperVersionId" api:"required" format:"uuid"`
	Responses       []PaperV3LegacyGetResponseCommentResponse    `json:"responses" api:"required"`
	Tag             string                                       `json:"tag" api:"required"`
	Title           string                                       `json:"title" api:"required"`
	UniversalID     string                                       `json:"universalId" api:"required"`
	Upvotes         float64                                      `json:"upvotes" api:"required"`
	WasEdited       bool                                         `json:"wasEdited" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		Annotation      respjson.Field
		Author          respjson.Field
		AuthorResponded respjson.Field
		Body            respjson.Field
		Date            respjson.Field
		Endorsements    respjson.Field
		HasDownvoted    respjson.Field
		HasFlagged      respjson.Field
		HasUpvoted      respjson.Field
		IsAuthor        respjson.Field
		PaperGroupID    respjson.Field
		PaperTitle      respjson.Field
		PaperVersionID  respjson.Field
		Responses       respjson.Field
		Tag             respjson.Field
		Title           respjson.Field
		UniversalID     respjson.Field
		Upvotes         respjson.Field
		WasEdited       respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3LegacyGetResponseComment) RawJSON() string { return r.JSON.raw }
func (r *PaperV3LegacyGetResponseComment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3LegacyGetResponseCommentAnnotation struct {
	AnchorPosition PaperV3LegacyGetResponseCommentAnnotationAnchorPosition  `json:"anchorPosition" api:"required"`
	FocusPosition  PaperV3LegacyGetResponseCommentAnnotationFocusPosition   `json:"focusPosition" api:"required"`
	HighlightRects []PaperV3LegacyGetResponseCommentAnnotationHighlightRect `json:"highlightRects" api:"required"`
	SelectedText   string                                                   `json:"selectedText" api:"required"`
	// Any of "highlight".
	Type           string `json:"type" api:"required"`
	HighlightColor string `json:"highlightColor" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AnchorPosition respjson.Field
		FocusPosition  respjson.Field
		HighlightRects respjson.Field
		SelectedText   respjson.Field
		Type           respjson.Field
		HighlightColor respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3LegacyGetResponseCommentAnnotation) RawJSON() string { return r.JSON.raw }
func (r *PaperV3LegacyGetResponseCommentAnnotation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3LegacyGetResponseCommentAnnotationAnchorPosition struct {
	Offset    float64 `json:"offset" api:"required"`
	PageIndex float64 `json:"pageIndex" api:"required"`
	SpanIndex float64 `json:"spanIndex" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Offset      respjson.Field
		PageIndex   respjson.Field
		SpanIndex   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3LegacyGetResponseCommentAnnotationAnchorPosition) RawJSON() string { return r.JSON.raw }
func (r *PaperV3LegacyGetResponseCommentAnnotationAnchorPosition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3LegacyGetResponseCommentAnnotationFocusPosition struct {
	Offset    float64 `json:"offset" api:"required"`
	PageIndex float64 `json:"pageIndex" api:"required"`
	SpanIndex float64 `json:"spanIndex" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Offset      respjson.Field
		PageIndex   respjson.Field
		SpanIndex   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3LegacyGetResponseCommentAnnotationFocusPosition) RawJSON() string { return r.JSON.raw }
func (r *PaperV3LegacyGetResponseCommentAnnotationFocusPosition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3LegacyGetResponseCommentAnnotationHighlightRect struct {
	PageIndex float64                                                      `json:"pageIndex" api:"required"`
	Rects     []PaperV3LegacyGetResponseCommentAnnotationHighlightRectRect `json:"rects" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PageIndex   respjson.Field
		Rects       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3LegacyGetResponseCommentAnnotationHighlightRect) RawJSON() string { return r.JSON.raw }
func (r *PaperV3LegacyGetResponseCommentAnnotationHighlightRect) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3LegacyGetResponseCommentAnnotationHighlightRectRect struct {
	X1 float64 `json:"x1" api:"required"`
	X2 float64 `json:"x2" api:"required"`
	Y1 float64 `json:"y1" api:"required"`
	Y2 float64 `json:"y2" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		X1          respjson.Field
		X2          respjson.Field
		Y1          respjson.Field
		Y2          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3LegacyGetResponseCommentAnnotationHighlightRectRect) RawJSON() string {
	return r.JSON.raw
}
func (r *PaperV3LegacyGetResponseCommentAnnotationHighlightRectRect) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3LegacyGetResponseCommentAuthor struct {
	ID               string                                        `json:"id" api:"required" format:"uuid"`
	Avatar           []PaperV3LegacyGetResponseCommentAuthorAvatar `json:"avatar" api:"required"`
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
func (r PaperV3LegacyGetResponseCommentAuthor) RawJSON() string { return r.JSON.raw }
func (r *PaperV3LegacyGetResponseCommentAuthor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3LegacyGetResponseCommentAuthorAvatar struct {
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
func (r PaperV3LegacyGetResponseCommentAuthorAvatar) RawJSON() string { return r.JSON.raw }
func (r *PaperV3LegacyGetResponseCommentAuthorAvatar) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3LegacyGetResponseCommentEndorsement struct {
	ID   string `json:"id" api:"required" format:"uuid"`
	Name string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3LegacyGetResponseCommentEndorsement) RawJSON() string { return r.JSON.raw }
func (r *PaperV3LegacyGetResponseCommentEndorsement) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3LegacyGetResponseCommentResponse struct {
	ID              string                                               `json:"id" api:"required" format:"uuid"`
	Annotation      PaperV3LegacyGetResponseCommentResponseAnnotation    `json:"annotation" api:"required"`
	Author          PaperV3LegacyGetResponseCommentResponseAuthor        `json:"author" api:"required"`
	AuthorResponded bool                                                 `json:"authorResponded" api:"required"`
	Body            string                                               `json:"body" api:"required"`
	Date            string                                               `json:"date" api:"required"`
	Endorsements    []PaperV3LegacyGetResponseCommentResponseEndorsement `json:"endorsements" api:"required"`
	HasDownvoted    bool                                                 `json:"hasDownvoted" api:"required"`
	HasFlagged      bool                                                 `json:"hasFlagged" api:"required"`
	HasUpvoted      bool                                                 `json:"hasUpvoted" api:"required"`
	IsAuthor        bool                                                 `json:"isAuthor" api:"required"`
	PaperGroupID    string                                               `json:"paperGroupId" api:"required" format:"uuid"`
	PaperTitle      string                                               `json:"paperTitle" api:"required"`
	PaperVersionID  string                                               `json:"paperVersionId" api:"required" format:"uuid"`
	Tag             string                                               `json:"tag" api:"required"`
	Title           string                                               `json:"title" api:"required"`
	UniversalID     string                                               `json:"universalId" api:"required"`
	Upvotes         float64                                              `json:"upvotes" api:"required"`
	WasEdited       bool                                                 `json:"wasEdited" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		Annotation      respjson.Field
		Author          respjson.Field
		AuthorResponded respjson.Field
		Body            respjson.Field
		Date            respjson.Field
		Endorsements    respjson.Field
		HasDownvoted    respjson.Field
		HasFlagged      respjson.Field
		HasUpvoted      respjson.Field
		IsAuthor        respjson.Field
		PaperGroupID    respjson.Field
		PaperTitle      respjson.Field
		PaperVersionID  respjson.Field
		Tag             respjson.Field
		Title           respjson.Field
		UniversalID     respjson.Field
		Upvotes         respjson.Field
		WasEdited       respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3LegacyGetResponseCommentResponse) RawJSON() string { return r.JSON.raw }
func (r *PaperV3LegacyGetResponseCommentResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3LegacyGetResponseCommentResponseAnnotation struct {
	AnchorPosition PaperV3LegacyGetResponseCommentResponseAnnotationAnchorPosition  `json:"anchorPosition" api:"required"`
	FocusPosition  PaperV3LegacyGetResponseCommentResponseAnnotationFocusPosition   `json:"focusPosition" api:"required"`
	HighlightRects []PaperV3LegacyGetResponseCommentResponseAnnotationHighlightRect `json:"highlightRects" api:"required"`
	SelectedText   string                                                           `json:"selectedText" api:"required"`
	// Any of "highlight".
	Type           string `json:"type" api:"required"`
	HighlightColor string `json:"highlightColor" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AnchorPosition respjson.Field
		FocusPosition  respjson.Field
		HighlightRects respjson.Field
		SelectedText   respjson.Field
		Type           respjson.Field
		HighlightColor respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3LegacyGetResponseCommentResponseAnnotation) RawJSON() string { return r.JSON.raw }
func (r *PaperV3LegacyGetResponseCommentResponseAnnotation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3LegacyGetResponseCommentResponseAnnotationAnchorPosition struct {
	Offset    float64 `json:"offset" api:"required"`
	PageIndex float64 `json:"pageIndex" api:"required"`
	SpanIndex float64 `json:"spanIndex" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Offset      respjson.Field
		PageIndex   respjson.Field
		SpanIndex   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3LegacyGetResponseCommentResponseAnnotationAnchorPosition) RawJSON() string {
	return r.JSON.raw
}
func (r *PaperV3LegacyGetResponseCommentResponseAnnotationAnchorPosition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3LegacyGetResponseCommentResponseAnnotationFocusPosition struct {
	Offset    float64 `json:"offset" api:"required"`
	PageIndex float64 `json:"pageIndex" api:"required"`
	SpanIndex float64 `json:"spanIndex" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Offset      respjson.Field
		PageIndex   respjson.Field
		SpanIndex   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3LegacyGetResponseCommentResponseAnnotationFocusPosition) RawJSON() string {
	return r.JSON.raw
}
func (r *PaperV3LegacyGetResponseCommentResponseAnnotationFocusPosition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3LegacyGetResponseCommentResponseAnnotationHighlightRect struct {
	PageIndex float64                                                              `json:"pageIndex" api:"required"`
	Rects     []PaperV3LegacyGetResponseCommentResponseAnnotationHighlightRectRect `json:"rects" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PageIndex   respjson.Field
		Rects       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3LegacyGetResponseCommentResponseAnnotationHighlightRect) RawJSON() string {
	return r.JSON.raw
}
func (r *PaperV3LegacyGetResponseCommentResponseAnnotationHighlightRect) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3LegacyGetResponseCommentResponseAnnotationHighlightRectRect struct {
	X1 float64 `json:"x1" api:"required"`
	X2 float64 `json:"x2" api:"required"`
	Y1 float64 `json:"y1" api:"required"`
	Y2 float64 `json:"y2" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		X1          respjson.Field
		X2          respjson.Field
		Y1          respjson.Field
		Y2          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3LegacyGetResponseCommentResponseAnnotationHighlightRectRect) RawJSON() string {
	return r.JSON.raw
}
func (r *PaperV3LegacyGetResponseCommentResponseAnnotationHighlightRectRect) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3LegacyGetResponseCommentResponseAuthor struct {
	ID               string                                                `json:"id" api:"required" format:"uuid"`
	Avatar           []PaperV3LegacyGetResponseCommentResponseAuthorAvatar `json:"avatar" api:"required"`
	BlueskyUsername  string                                                `json:"blueskyUsername" api:"required"`
	GitHubUsername   string                                                `json:"githubUsername" api:"required"`
	GoogleScholarID  string                                                `json:"googleScholarId" api:"required"`
	Institution      string                                                `json:"institution" api:"required"`
	LinkedinUsername string                                                `json:"linkedinUsername" api:"required"`
	OrcidID          string                                                `json:"orcidId" api:"required"`
	PublicEmail      string                                                `json:"publicEmail" api:"required"`
	RealName         string                                                `json:"realName" api:"required"`
	Reputation       float64                                               `json:"reputation" api:"required"`
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
func (r PaperV3LegacyGetResponseCommentResponseAuthor) RawJSON() string { return r.JSON.raw }
func (r *PaperV3LegacyGetResponseCommentResponseAuthor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3LegacyGetResponseCommentResponseAuthorAvatar struct {
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
func (r PaperV3LegacyGetResponseCommentResponseAuthorAvatar) RawJSON() string { return r.JSON.raw }
func (r *PaperV3LegacyGetResponseCommentResponseAuthorAvatar) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3LegacyGetResponseCommentResponseEndorsement struct {
	ID   string `json:"id" api:"required" format:"uuid"`
	Name string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3LegacyGetResponseCommentResponseEndorsement) RawJSON() string { return r.JSON.raw }
func (r *PaperV3LegacyGetResponseCommentResponseEndorsement) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3LegacyGetCommentsResponse struct {
	ID              string                                        `json:"id" api:"required" format:"uuid"`
	Annotation      PaperV3LegacyGetCommentsResponseAnnotation    `json:"annotation" api:"required"`
	Author          PaperV3LegacyGetCommentsResponseAuthor        `json:"author" api:"required"`
	AuthorResponded bool                                          `json:"authorResponded" api:"required"`
	Body            string                                        `json:"body" api:"required"`
	Date            string                                        `json:"date" api:"required"`
	Endorsements    []PaperV3LegacyGetCommentsResponseEndorsement `json:"endorsements" api:"required"`
	HasDownvoted    bool                                          `json:"hasDownvoted" api:"required"`
	HasFlagged      bool                                          `json:"hasFlagged" api:"required"`
	HasUpvoted      bool                                          `json:"hasUpvoted" api:"required"`
	IsAuthor        bool                                          `json:"isAuthor" api:"required"`
	PaperGroupID    string                                        `json:"paperGroupId" api:"required" format:"uuid"`
	PaperTitle      string                                        `json:"paperTitle" api:"required"`
	PaperVersionID  string                                        `json:"paperVersionId" api:"required" format:"uuid"`
	Responses       []PaperV3LegacyGetCommentsResponseResponse    `json:"responses" api:"required"`
	Tag             string                                        `json:"tag" api:"required"`
	Title           string                                        `json:"title" api:"required"`
	UniversalID     string                                        `json:"universalId" api:"required"`
	Upvotes         float64                                       `json:"upvotes" api:"required"`
	WasEdited       bool                                          `json:"wasEdited" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		Annotation      respjson.Field
		Author          respjson.Field
		AuthorResponded respjson.Field
		Body            respjson.Field
		Date            respjson.Field
		Endorsements    respjson.Field
		HasDownvoted    respjson.Field
		HasFlagged      respjson.Field
		HasUpvoted      respjson.Field
		IsAuthor        respjson.Field
		PaperGroupID    respjson.Field
		PaperTitle      respjson.Field
		PaperVersionID  respjson.Field
		Responses       respjson.Field
		Tag             respjson.Field
		Title           respjson.Field
		UniversalID     respjson.Field
		Upvotes         respjson.Field
		WasEdited       respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3LegacyGetCommentsResponse) RawJSON() string { return r.JSON.raw }
func (r *PaperV3LegacyGetCommentsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3LegacyGetCommentsResponseAnnotation struct {
	AnchorPosition PaperV3LegacyGetCommentsResponseAnnotationAnchorPosition  `json:"anchorPosition" api:"required"`
	FocusPosition  PaperV3LegacyGetCommentsResponseAnnotationFocusPosition   `json:"focusPosition" api:"required"`
	HighlightRects []PaperV3LegacyGetCommentsResponseAnnotationHighlightRect `json:"highlightRects" api:"required"`
	SelectedText   string                                                    `json:"selectedText" api:"required"`
	// Any of "highlight".
	Type           string `json:"type" api:"required"`
	HighlightColor string `json:"highlightColor" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AnchorPosition respjson.Field
		FocusPosition  respjson.Field
		HighlightRects respjson.Field
		SelectedText   respjson.Field
		Type           respjson.Field
		HighlightColor respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3LegacyGetCommentsResponseAnnotation) RawJSON() string { return r.JSON.raw }
func (r *PaperV3LegacyGetCommentsResponseAnnotation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3LegacyGetCommentsResponseAnnotationAnchorPosition struct {
	Offset    float64 `json:"offset" api:"required"`
	PageIndex float64 `json:"pageIndex" api:"required"`
	SpanIndex float64 `json:"spanIndex" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Offset      respjson.Field
		PageIndex   respjson.Field
		SpanIndex   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3LegacyGetCommentsResponseAnnotationAnchorPosition) RawJSON() string { return r.JSON.raw }
func (r *PaperV3LegacyGetCommentsResponseAnnotationAnchorPosition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3LegacyGetCommentsResponseAnnotationFocusPosition struct {
	Offset    float64 `json:"offset" api:"required"`
	PageIndex float64 `json:"pageIndex" api:"required"`
	SpanIndex float64 `json:"spanIndex" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Offset      respjson.Field
		PageIndex   respjson.Field
		SpanIndex   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3LegacyGetCommentsResponseAnnotationFocusPosition) RawJSON() string { return r.JSON.raw }
func (r *PaperV3LegacyGetCommentsResponseAnnotationFocusPosition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3LegacyGetCommentsResponseAnnotationHighlightRect struct {
	PageIndex float64                                                       `json:"pageIndex" api:"required"`
	Rects     []PaperV3LegacyGetCommentsResponseAnnotationHighlightRectRect `json:"rects" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PageIndex   respjson.Field
		Rects       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3LegacyGetCommentsResponseAnnotationHighlightRect) RawJSON() string { return r.JSON.raw }
func (r *PaperV3LegacyGetCommentsResponseAnnotationHighlightRect) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3LegacyGetCommentsResponseAnnotationHighlightRectRect struct {
	X1 float64 `json:"x1" api:"required"`
	X2 float64 `json:"x2" api:"required"`
	Y1 float64 `json:"y1" api:"required"`
	Y2 float64 `json:"y2" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		X1          respjson.Field
		X2          respjson.Field
		Y1          respjson.Field
		Y2          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3LegacyGetCommentsResponseAnnotationHighlightRectRect) RawJSON() string {
	return r.JSON.raw
}
func (r *PaperV3LegacyGetCommentsResponseAnnotationHighlightRectRect) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3LegacyGetCommentsResponseAuthor struct {
	ID               string                                         `json:"id" api:"required" format:"uuid"`
	Avatar           []PaperV3LegacyGetCommentsResponseAuthorAvatar `json:"avatar" api:"required"`
	BlueskyUsername  string                                         `json:"blueskyUsername" api:"required"`
	GitHubUsername   string                                         `json:"githubUsername" api:"required"`
	GoogleScholarID  string                                         `json:"googleScholarId" api:"required"`
	Institution      string                                         `json:"institution" api:"required"`
	LinkedinUsername string                                         `json:"linkedinUsername" api:"required"`
	OrcidID          string                                         `json:"orcidId" api:"required"`
	PublicEmail      string                                         `json:"publicEmail" api:"required"`
	RealName         string                                         `json:"realName" api:"required"`
	Reputation       float64                                        `json:"reputation" api:"required"`
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
func (r PaperV3LegacyGetCommentsResponseAuthor) RawJSON() string { return r.JSON.raw }
func (r *PaperV3LegacyGetCommentsResponseAuthor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3LegacyGetCommentsResponseAuthorAvatar struct {
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
func (r PaperV3LegacyGetCommentsResponseAuthorAvatar) RawJSON() string { return r.JSON.raw }
func (r *PaperV3LegacyGetCommentsResponseAuthorAvatar) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3LegacyGetCommentsResponseEndorsement struct {
	ID   string `json:"id" api:"required" format:"uuid"`
	Name string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3LegacyGetCommentsResponseEndorsement) RawJSON() string { return r.JSON.raw }
func (r *PaperV3LegacyGetCommentsResponseEndorsement) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3LegacyGetCommentsResponseResponse struct {
	ID              string                                                `json:"id" api:"required" format:"uuid"`
	Annotation      PaperV3LegacyGetCommentsResponseResponseAnnotation    `json:"annotation" api:"required"`
	Author          PaperV3LegacyGetCommentsResponseResponseAuthor        `json:"author" api:"required"`
	AuthorResponded bool                                                  `json:"authorResponded" api:"required"`
	Body            string                                                `json:"body" api:"required"`
	Date            string                                                `json:"date" api:"required"`
	Endorsements    []PaperV3LegacyGetCommentsResponseResponseEndorsement `json:"endorsements" api:"required"`
	HasDownvoted    bool                                                  `json:"hasDownvoted" api:"required"`
	HasFlagged      bool                                                  `json:"hasFlagged" api:"required"`
	HasUpvoted      bool                                                  `json:"hasUpvoted" api:"required"`
	IsAuthor        bool                                                  `json:"isAuthor" api:"required"`
	PaperGroupID    string                                                `json:"paperGroupId" api:"required" format:"uuid"`
	PaperTitle      string                                                `json:"paperTitle" api:"required"`
	PaperVersionID  string                                                `json:"paperVersionId" api:"required" format:"uuid"`
	Tag             string                                                `json:"tag" api:"required"`
	Title           string                                                `json:"title" api:"required"`
	UniversalID     string                                                `json:"universalId" api:"required"`
	Upvotes         float64                                               `json:"upvotes" api:"required"`
	WasEdited       bool                                                  `json:"wasEdited" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		Annotation      respjson.Field
		Author          respjson.Field
		AuthorResponded respjson.Field
		Body            respjson.Field
		Date            respjson.Field
		Endorsements    respjson.Field
		HasDownvoted    respjson.Field
		HasFlagged      respjson.Field
		HasUpvoted      respjson.Field
		IsAuthor        respjson.Field
		PaperGroupID    respjson.Field
		PaperTitle      respjson.Field
		PaperVersionID  respjson.Field
		Tag             respjson.Field
		Title           respjson.Field
		UniversalID     respjson.Field
		Upvotes         respjson.Field
		WasEdited       respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3LegacyGetCommentsResponseResponse) RawJSON() string { return r.JSON.raw }
func (r *PaperV3LegacyGetCommentsResponseResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3LegacyGetCommentsResponseResponseAnnotation struct {
	AnchorPosition PaperV3LegacyGetCommentsResponseResponseAnnotationAnchorPosition  `json:"anchorPosition" api:"required"`
	FocusPosition  PaperV3LegacyGetCommentsResponseResponseAnnotationFocusPosition   `json:"focusPosition" api:"required"`
	HighlightRects []PaperV3LegacyGetCommentsResponseResponseAnnotationHighlightRect `json:"highlightRects" api:"required"`
	SelectedText   string                                                            `json:"selectedText" api:"required"`
	// Any of "highlight".
	Type           string `json:"type" api:"required"`
	HighlightColor string `json:"highlightColor" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AnchorPosition respjson.Field
		FocusPosition  respjson.Field
		HighlightRects respjson.Field
		SelectedText   respjson.Field
		Type           respjson.Field
		HighlightColor respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3LegacyGetCommentsResponseResponseAnnotation) RawJSON() string { return r.JSON.raw }
func (r *PaperV3LegacyGetCommentsResponseResponseAnnotation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3LegacyGetCommentsResponseResponseAnnotationAnchorPosition struct {
	Offset    float64 `json:"offset" api:"required"`
	PageIndex float64 `json:"pageIndex" api:"required"`
	SpanIndex float64 `json:"spanIndex" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Offset      respjson.Field
		PageIndex   respjson.Field
		SpanIndex   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3LegacyGetCommentsResponseResponseAnnotationAnchorPosition) RawJSON() string {
	return r.JSON.raw
}
func (r *PaperV3LegacyGetCommentsResponseResponseAnnotationAnchorPosition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3LegacyGetCommentsResponseResponseAnnotationFocusPosition struct {
	Offset    float64 `json:"offset" api:"required"`
	PageIndex float64 `json:"pageIndex" api:"required"`
	SpanIndex float64 `json:"spanIndex" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Offset      respjson.Field
		PageIndex   respjson.Field
		SpanIndex   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3LegacyGetCommentsResponseResponseAnnotationFocusPosition) RawJSON() string {
	return r.JSON.raw
}
func (r *PaperV3LegacyGetCommentsResponseResponseAnnotationFocusPosition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3LegacyGetCommentsResponseResponseAnnotationHighlightRect struct {
	PageIndex float64                                                               `json:"pageIndex" api:"required"`
	Rects     []PaperV3LegacyGetCommentsResponseResponseAnnotationHighlightRectRect `json:"rects" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PageIndex   respjson.Field
		Rects       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3LegacyGetCommentsResponseResponseAnnotationHighlightRect) RawJSON() string {
	return r.JSON.raw
}
func (r *PaperV3LegacyGetCommentsResponseResponseAnnotationHighlightRect) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3LegacyGetCommentsResponseResponseAnnotationHighlightRectRect struct {
	X1 float64 `json:"x1" api:"required"`
	X2 float64 `json:"x2" api:"required"`
	Y1 float64 `json:"y1" api:"required"`
	Y2 float64 `json:"y2" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		X1          respjson.Field
		X2          respjson.Field
		Y1          respjson.Field
		Y2          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3LegacyGetCommentsResponseResponseAnnotationHighlightRectRect) RawJSON() string {
	return r.JSON.raw
}
func (r *PaperV3LegacyGetCommentsResponseResponseAnnotationHighlightRectRect) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3LegacyGetCommentsResponseResponseAuthor struct {
	ID               string                                                 `json:"id" api:"required" format:"uuid"`
	Avatar           []PaperV3LegacyGetCommentsResponseResponseAuthorAvatar `json:"avatar" api:"required"`
	BlueskyUsername  string                                                 `json:"blueskyUsername" api:"required"`
	GitHubUsername   string                                                 `json:"githubUsername" api:"required"`
	GoogleScholarID  string                                                 `json:"googleScholarId" api:"required"`
	Institution      string                                                 `json:"institution" api:"required"`
	LinkedinUsername string                                                 `json:"linkedinUsername" api:"required"`
	OrcidID          string                                                 `json:"orcidId" api:"required"`
	PublicEmail      string                                                 `json:"publicEmail" api:"required"`
	RealName         string                                                 `json:"realName" api:"required"`
	Reputation       float64                                                `json:"reputation" api:"required"`
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
func (r PaperV3LegacyGetCommentsResponseResponseAuthor) RawJSON() string { return r.JSON.raw }
func (r *PaperV3LegacyGetCommentsResponseResponseAuthor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3LegacyGetCommentsResponseResponseAuthorAvatar struct {
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
func (r PaperV3LegacyGetCommentsResponseResponseAuthorAvatar) RawJSON() string { return r.JSON.raw }
func (r *PaperV3LegacyGetCommentsResponseResponseAuthorAvatar) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3LegacyGetCommentsResponseResponseEndorsement struct {
	ID   string `json:"id" api:"required" format:"uuid"`
	Name string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3LegacyGetCommentsResponseResponseEndorsement) RawJSON() string { return r.JSON.raw }
func (r *PaperV3LegacyGetCommentsResponseResponseEndorsement) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
