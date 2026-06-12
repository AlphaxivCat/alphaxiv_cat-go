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
// `api-server/file:/app/api-server/src/controllers/papers/v3/legacy/get-v2-paper-for-display.controller.ts`
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
// `api-server/file:/app/api-server/src/controllers/papers/v3/legacy/get-v2-comments.controller.ts`
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
	Author          PaperV3LegacyGetResponseCommentAuthorUnion   `json:"author" api:"required"`
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

// PaperV3LegacyGetResponseCommentAuthorUnion contains all possible properties and
// values from [PaperV3LegacyGetResponseCommentAuthorObject],
// [PaperV3LegacyGetResponseCommentAuthorObject2].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type PaperV3LegacyGetResponseCommentAuthorUnion struct {
	// This field is a union of [string], [any]
	ID PaperV3LegacyGetResponseCommentAuthorUnionID `json:"id"`
	// This field is a union of [[]PaperV3LegacyGetResponseCommentAuthorObjectAvatar],
	// [[]PaperV3LegacyGetResponseCommentAuthorObject2Avatar]
	Avatar           PaperV3LegacyGetResponseCommentAuthorUnionAvatar `json:"avatar"`
	BlueskyUsername  string                                           `json:"blueskyUsername"`
	GitHubUsername   string                                           `json:"githubUsername"`
	GoogleScholarID  string                                           `json:"googleScholarId"`
	Institution      string                                           `json:"institution"`
	LinkedinUsername string                                           `json:"linkedinUsername"`
	OrcidID          string                                           `json:"orcidId"`
	PublicEmail      string                                           `json:"publicEmail"`
	RealName         string                                           `json:"realName"`
	Reputation       float64                                          `json:"reputation"`
	Role             string                                           `json:"role"`
	Username         string                                           `json:"username"`
	Verified         bool                                             `json:"verified"`
	WeeklyReputation float64                                          `json:"weeklyReputation"`
	XUsername        string                                           `json:"xUsername"`
	JSON             struct {
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
		raw              string
	} `json:"-"`
}

func (u PaperV3LegacyGetResponseCommentAuthorUnion) AsPaperV3LegacyGetResponseCommentAuthorObject() (v PaperV3LegacyGetResponseCommentAuthorObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PaperV3LegacyGetResponseCommentAuthorUnion) AsPaperV3LegacyGetResponseCommentAuthorObject2() (v PaperV3LegacyGetResponseCommentAuthorObject2) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PaperV3LegacyGetResponseCommentAuthorUnion) RawJSON() string { return u.JSON.raw }

func (r *PaperV3LegacyGetResponseCommentAuthorUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PaperV3LegacyGetResponseCommentAuthorUnionID is an implicit subunion of
// [PaperV3LegacyGetResponseCommentAuthorUnion].
// PaperV3LegacyGetResponseCommentAuthorUnionID provides convenient access to the
// sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [PaperV3LegacyGetResponseCommentAuthorUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfPaperV3LegacyGetResponseCommentAuthorObject2ID]
type PaperV3LegacyGetResponseCommentAuthorUnionID struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfPaperV3LegacyGetResponseCommentAuthorObject2ID any `json:",inline"`
	JSON                                             struct {
		OfString                                         respjson.Field
		OfPaperV3LegacyGetResponseCommentAuthorObject2ID respjson.Field
		raw                                              string
	} `json:"-"`
}

func (r *PaperV3LegacyGetResponseCommentAuthorUnionID) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PaperV3LegacyGetResponseCommentAuthorUnionAvatar is an implicit subunion of
// [PaperV3LegacyGetResponseCommentAuthorUnion].
// PaperV3LegacyGetResponseCommentAuthorUnionAvatar provides convenient access to
// the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [PaperV3LegacyGetResponseCommentAuthorUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfPaperV3LegacyGetResponseCommentAuthorObjectAvatarArray
// OfPaperV3LegacyGetResponseCommentAuthorObject2AvatarArray]
type PaperV3LegacyGetResponseCommentAuthorUnionAvatar struct {
	// This field will be present if the value is a
	// [[]PaperV3LegacyGetResponseCommentAuthorObjectAvatar] instead of an object.
	OfPaperV3LegacyGetResponseCommentAuthorObjectAvatarArray []PaperV3LegacyGetResponseCommentAuthorObjectAvatar `json:",inline"`
	// This field will be present if the value is a
	// [[]PaperV3LegacyGetResponseCommentAuthorObject2Avatar] instead of an object.
	OfPaperV3LegacyGetResponseCommentAuthorObject2AvatarArray []PaperV3LegacyGetResponseCommentAuthorObject2Avatar `json:",inline"`
	JSON                                                      struct {
		OfPaperV3LegacyGetResponseCommentAuthorObjectAvatarArray  respjson.Field
		OfPaperV3LegacyGetResponseCommentAuthorObject2AvatarArray respjson.Field
		raw                                                       string
	} `json:"-"`
}

func (r *PaperV3LegacyGetResponseCommentAuthorUnionAvatar) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3LegacyGetResponseCommentAuthorObject struct {
	ID               string                                              `json:"id" api:"required" format:"uuid"`
	Avatar           []PaperV3LegacyGetResponseCommentAuthorObjectAvatar `json:"avatar" api:"required"`
	BlueskyUsername  string                                              `json:"blueskyUsername" api:"required"`
	GitHubUsername   string                                              `json:"githubUsername" api:"required"`
	GoogleScholarID  string                                              `json:"googleScholarId" api:"required"`
	Institution      string                                              `json:"institution" api:"required"`
	LinkedinUsername string                                              `json:"linkedinUsername" api:"required"`
	OrcidID          string                                              `json:"orcidId" api:"required"`
	PublicEmail      string                                              `json:"publicEmail" api:"required"`
	RealName         string                                              `json:"realName" api:"required"`
	Reputation       float64                                             `json:"reputation" api:"required"`
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
func (r PaperV3LegacyGetResponseCommentAuthorObject) RawJSON() string { return r.JSON.raw }
func (r *PaperV3LegacyGetResponseCommentAuthorObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3LegacyGetResponseCommentAuthorObjectAvatar struct {
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
func (r PaperV3LegacyGetResponseCommentAuthorObjectAvatar) RawJSON() string { return r.JSON.raw }
func (r *PaperV3LegacyGetResponseCommentAuthorObjectAvatar) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3LegacyGetResponseCommentAuthorObject2 struct {
	ID               any                                                  `json:"id" api:"required"`
	Avatar           []PaperV3LegacyGetResponseCommentAuthorObject2Avatar `json:"avatar" api:"required"`
	BlueskyUsername  string                                               `json:"blueskyUsername" api:"required"`
	GitHubUsername   string                                               `json:"githubUsername" api:"required"`
	GoogleScholarID  string                                               `json:"googleScholarId" api:"required"`
	Institution      string                                               `json:"institution" api:"required"`
	LinkedinUsername string                                               `json:"linkedinUsername" api:"required"`
	OrcidID          string                                               `json:"orcidId" api:"required"`
	PublicEmail      string                                               `json:"publicEmail" api:"required"`
	RealName         string                                               `json:"realName" api:"required"`
	Reputation       float64                                              `json:"reputation" api:"required"`
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
func (r PaperV3LegacyGetResponseCommentAuthorObject2) RawJSON() string { return r.JSON.raw }
func (r *PaperV3LegacyGetResponseCommentAuthorObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3LegacyGetResponseCommentAuthorObject2Avatar struct {
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
func (r PaperV3LegacyGetResponseCommentAuthorObject2Avatar) RawJSON() string { return r.JSON.raw }
func (r *PaperV3LegacyGetResponseCommentAuthorObject2Avatar) UnmarshalJSON(data []byte) error {
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
	Author          PaperV3LegacyGetResponseCommentResponseAuthorUnion   `json:"author" api:"required"`
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

// PaperV3LegacyGetResponseCommentResponseAuthorUnion contains all possible
// properties and values from
// [PaperV3LegacyGetResponseCommentResponseAuthorObject],
// [PaperV3LegacyGetResponseCommentResponseAuthorObject2].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type PaperV3LegacyGetResponseCommentResponseAuthorUnion struct {
	// This field is a union of [string], [any]
	ID PaperV3LegacyGetResponseCommentResponseAuthorUnionID `json:"id"`
	// This field is a union of
	// [[]PaperV3LegacyGetResponseCommentResponseAuthorObjectAvatar],
	// [[]PaperV3LegacyGetResponseCommentResponseAuthorObject2Avatar]
	Avatar           PaperV3LegacyGetResponseCommentResponseAuthorUnionAvatar `json:"avatar"`
	BlueskyUsername  string                                                   `json:"blueskyUsername"`
	GitHubUsername   string                                                   `json:"githubUsername"`
	GoogleScholarID  string                                                   `json:"googleScholarId"`
	Institution      string                                                   `json:"institution"`
	LinkedinUsername string                                                   `json:"linkedinUsername"`
	OrcidID          string                                                   `json:"orcidId"`
	PublicEmail      string                                                   `json:"publicEmail"`
	RealName         string                                                   `json:"realName"`
	Reputation       float64                                                  `json:"reputation"`
	Role             string                                                   `json:"role"`
	Username         string                                                   `json:"username"`
	Verified         bool                                                     `json:"verified"`
	WeeklyReputation float64                                                  `json:"weeklyReputation"`
	XUsername        string                                                   `json:"xUsername"`
	JSON             struct {
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
		raw              string
	} `json:"-"`
}

func (u PaperV3LegacyGetResponseCommentResponseAuthorUnion) AsPaperV3LegacyGetResponseCommentResponseAuthorObject() (v PaperV3LegacyGetResponseCommentResponseAuthorObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PaperV3LegacyGetResponseCommentResponseAuthorUnion) AsPaperV3LegacyGetResponseCommentResponseAuthorObject2() (v PaperV3LegacyGetResponseCommentResponseAuthorObject2) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PaperV3LegacyGetResponseCommentResponseAuthorUnion) RawJSON() string { return u.JSON.raw }

func (r *PaperV3LegacyGetResponseCommentResponseAuthorUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PaperV3LegacyGetResponseCommentResponseAuthorUnionID is an implicit subunion of
// [PaperV3LegacyGetResponseCommentResponseAuthorUnion].
// PaperV3LegacyGetResponseCommentResponseAuthorUnionID provides convenient access
// to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [PaperV3LegacyGetResponseCommentResponseAuthorUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString
// OfPaperV3LegacyGetResponseCommentResponseAuthorObject2ID]
type PaperV3LegacyGetResponseCommentResponseAuthorUnionID struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfPaperV3LegacyGetResponseCommentResponseAuthorObject2ID any `json:",inline"`
	JSON                                                     struct {
		OfString                                                 respjson.Field
		OfPaperV3LegacyGetResponseCommentResponseAuthorObject2ID respjson.Field
		raw                                                      string
	} `json:"-"`
}

func (r *PaperV3LegacyGetResponseCommentResponseAuthorUnionID) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PaperV3LegacyGetResponseCommentResponseAuthorUnionAvatar is an implicit subunion
// of [PaperV3LegacyGetResponseCommentResponseAuthorUnion].
// PaperV3LegacyGetResponseCommentResponseAuthorUnionAvatar provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [PaperV3LegacyGetResponseCommentResponseAuthorUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfPaperV3LegacyGetResponseCommentResponseAuthorObjectAvatarArray
// OfPaperV3LegacyGetResponseCommentResponseAuthorObject2AvatarArray]
type PaperV3LegacyGetResponseCommentResponseAuthorUnionAvatar struct {
	// This field will be present if the value is a
	// [[]PaperV3LegacyGetResponseCommentResponseAuthorObjectAvatar] instead of an
	// object.
	OfPaperV3LegacyGetResponseCommentResponseAuthorObjectAvatarArray []PaperV3LegacyGetResponseCommentResponseAuthorObjectAvatar `json:",inline"`
	// This field will be present if the value is a
	// [[]PaperV3LegacyGetResponseCommentResponseAuthorObject2Avatar] instead of an
	// object.
	OfPaperV3LegacyGetResponseCommentResponseAuthorObject2AvatarArray []PaperV3LegacyGetResponseCommentResponseAuthorObject2Avatar `json:",inline"`
	JSON                                                              struct {
		OfPaperV3LegacyGetResponseCommentResponseAuthorObjectAvatarArray  respjson.Field
		OfPaperV3LegacyGetResponseCommentResponseAuthorObject2AvatarArray respjson.Field
		raw                                                               string
	} `json:"-"`
}

func (r *PaperV3LegacyGetResponseCommentResponseAuthorUnionAvatar) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3LegacyGetResponseCommentResponseAuthorObject struct {
	ID               string                                                      `json:"id" api:"required" format:"uuid"`
	Avatar           []PaperV3LegacyGetResponseCommentResponseAuthorObjectAvatar `json:"avatar" api:"required"`
	BlueskyUsername  string                                                      `json:"blueskyUsername" api:"required"`
	GitHubUsername   string                                                      `json:"githubUsername" api:"required"`
	GoogleScholarID  string                                                      `json:"googleScholarId" api:"required"`
	Institution      string                                                      `json:"institution" api:"required"`
	LinkedinUsername string                                                      `json:"linkedinUsername" api:"required"`
	OrcidID          string                                                      `json:"orcidId" api:"required"`
	PublicEmail      string                                                      `json:"publicEmail" api:"required"`
	RealName         string                                                      `json:"realName" api:"required"`
	Reputation       float64                                                     `json:"reputation" api:"required"`
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
func (r PaperV3LegacyGetResponseCommentResponseAuthorObject) RawJSON() string { return r.JSON.raw }
func (r *PaperV3LegacyGetResponseCommentResponseAuthorObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3LegacyGetResponseCommentResponseAuthorObjectAvatar struct {
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
func (r PaperV3LegacyGetResponseCommentResponseAuthorObjectAvatar) RawJSON() string {
	return r.JSON.raw
}
func (r *PaperV3LegacyGetResponseCommentResponseAuthorObjectAvatar) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3LegacyGetResponseCommentResponseAuthorObject2 struct {
	ID               any                                                          `json:"id" api:"required"`
	Avatar           []PaperV3LegacyGetResponseCommentResponseAuthorObject2Avatar `json:"avatar" api:"required"`
	BlueskyUsername  string                                                       `json:"blueskyUsername" api:"required"`
	GitHubUsername   string                                                       `json:"githubUsername" api:"required"`
	GoogleScholarID  string                                                       `json:"googleScholarId" api:"required"`
	Institution      string                                                       `json:"institution" api:"required"`
	LinkedinUsername string                                                       `json:"linkedinUsername" api:"required"`
	OrcidID          string                                                       `json:"orcidId" api:"required"`
	PublicEmail      string                                                       `json:"publicEmail" api:"required"`
	RealName         string                                                       `json:"realName" api:"required"`
	Reputation       float64                                                      `json:"reputation" api:"required"`
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
func (r PaperV3LegacyGetResponseCommentResponseAuthorObject2) RawJSON() string { return r.JSON.raw }
func (r *PaperV3LegacyGetResponseCommentResponseAuthorObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3LegacyGetResponseCommentResponseAuthorObject2Avatar struct {
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
func (r PaperV3LegacyGetResponseCommentResponseAuthorObject2Avatar) RawJSON() string {
	return r.JSON.raw
}
func (r *PaperV3LegacyGetResponseCommentResponseAuthorObject2Avatar) UnmarshalJSON(data []byte) error {
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
	Author          PaperV3LegacyGetCommentsResponseAuthorUnion   `json:"author" api:"required"`
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

// PaperV3LegacyGetCommentsResponseAuthorUnion contains all possible properties and
// values from [PaperV3LegacyGetCommentsResponseAuthorObject],
// [PaperV3LegacyGetCommentsResponseAuthorObject2].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type PaperV3LegacyGetCommentsResponseAuthorUnion struct {
	// This field is a union of [string], [any]
	ID PaperV3LegacyGetCommentsResponseAuthorUnionID `json:"id"`
	// This field is a union of [[]PaperV3LegacyGetCommentsResponseAuthorObjectAvatar],
	// [[]PaperV3LegacyGetCommentsResponseAuthorObject2Avatar]
	Avatar           PaperV3LegacyGetCommentsResponseAuthorUnionAvatar `json:"avatar"`
	BlueskyUsername  string                                            `json:"blueskyUsername"`
	GitHubUsername   string                                            `json:"githubUsername"`
	GoogleScholarID  string                                            `json:"googleScholarId"`
	Institution      string                                            `json:"institution"`
	LinkedinUsername string                                            `json:"linkedinUsername"`
	OrcidID          string                                            `json:"orcidId"`
	PublicEmail      string                                            `json:"publicEmail"`
	RealName         string                                            `json:"realName"`
	Reputation       float64                                           `json:"reputation"`
	Role             string                                            `json:"role"`
	Username         string                                            `json:"username"`
	Verified         bool                                              `json:"verified"`
	WeeklyReputation float64                                           `json:"weeklyReputation"`
	XUsername        string                                            `json:"xUsername"`
	JSON             struct {
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
		raw              string
	} `json:"-"`
}

func (u PaperV3LegacyGetCommentsResponseAuthorUnion) AsPaperV3LegacyGetCommentsResponseAuthorObject() (v PaperV3LegacyGetCommentsResponseAuthorObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PaperV3LegacyGetCommentsResponseAuthorUnion) AsPaperV3LegacyGetCommentsResponseAuthorObject2() (v PaperV3LegacyGetCommentsResponseAuthorObject2) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PaperV3LegacyGetCommentsResponseAuthorUnion) RawJSON() string { return u.JSON.raw }

func (r *PaperV3LegacyGetCommentsResponseAuthorUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PaperV3LegacyGetCommentsResponseAuthorUnionID is an implicit subunion of
// [PaperV3LegacyGetCommentsResponseAuthorUnion].
// PaperV3LegacyGetCommentsResponseAuthorUnionID provides convenient access to the
// sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [PaperV3LegacyGetCommentsResponseAuthorUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfPaperV3LegacyGetCommentsResponseAuthorObject2ID]
type PaperV3LegacyGetCommentsResponseAuthorUnionID struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfPaperV3LegacyGetCommentsResponseAuthorObject2ID any `json:",inline"`
	JSON                                              struct {
		OfString                                          respjson.Field
		OfPaperV3LegacyGetCommentsResponseAuthorObject2ID respjson.Field
		raw                                               string
	} `json:"-"`
}

func (r *PaperV3LegacyGetCommentsResponseAuthorUnionID) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PaperV3LegacyGetCommentsResponseAuthorUnionAvatar is an implicit subunion of
// [PaperV3LegacyGetCommentsResponseAuthorUnion].
// PaperV3LegacyGetCommentsResponseAuthorUnionAvatar provides convenient access to
// the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [PaperV3LegacyGetCommentsResponseAuthorUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfPaperV3LegacyGetCommentsResponseAuthorObjectAvatarArray
// OfPaperV3LegacyGetCommentsResponseAuthorObject2AvatarArray]
type PaperV3LegacyGetCommentsResponseAuthorUnionAvatar struct {
	// This field will be present if the value is a
	// [[]PaperV3LegacyGetCommentsResponseAuthorObjectAvatar] instead of an object.
	OfPaperV3LegacyGetCommentsResponseAuthorObjectAvatarArray []PaperV3LegacyGetCommentsResponseAuthorObjectAvatar `json:",inline"`
	// This field will be present if the value is a
	// [[]PaperV3LegacyGetCommentsResponseAuthorObject2Avatar] instead of an object.
	OfPaperV3LegacyGetCommentsResponseAuthorObject2AvatarArray []PaperV3LegacyGetCommentsResponseAuthorObject2Avatar `json:",inline"`
	JSON                                                       struct {
		OfPaperV3LegacyGetCommentsResponseAuthorObjectAvatarArray  respjson.Field
		OfPaperV3LegacyGetCommentsResponseAuthorObject2AvatarArray respjson.Field
		raw                                                        string
	} `json:"-"`
}

func (r *PaperV3LegacyGetCommentsResponseAuthorUnionAvatar) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3LegacyGetCommentsResponseAuthorObject struct {
	ID               string                                               `json:"id" api:"required" format:"uuid"`
	Avatar           []PaperV3LegacyGetCommentsResponseAuthorObjectAvatar `json:"avatar" api:"required"`
	BlueskyUsername  string                                               `json:"blueskyUsername" api:"required"`
	GitHubUsername   string                                               `json:"githubUsername" api:"required"`
	GoogleScholarID  string                                               `json:"googleScholarId" api:"required"`
	Institution      string                                               `json:"institution" api:"required"`
	LinkedinUsername string                                               `json:"linkedinUsername" api:"required"`
	OrcidID          string                                               `json:"orcidId" api:"required"`
	PublicEmail      string                                               `json:"publicEmail" api:"required"`
	RealName         string                                               `json:"realName" api:"required"`
	Reputation       float64                                              `json:"reputation" api:"required"`
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
func (r PaperV3LegacyGetCommentsResponseAuthorObject) RawJSON() string { return r.JSON.raw }
func (r *PaperV3LegacyGetCommentsResponseAuthorObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3LegacyGetCommentsResponseAuthorObjectAvatar struct {
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
func (r PaperV3LegacyGetCommentsResponseAuthorObjectAvatar) RawJSON() string { return r.JSON.raw }
func (r *PaperV3LegacyGetCommentsResponseAuthorObjectAvatar) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3LegacyGetCommentsResponseAuthorObject2 struct {
	ID               any                                                   `json:"id" api:"required"`
	Avatar           []PaperV3LegacyGetCommentsResponseAuthorObject2Avatar `json:"avatar" api:"required"`
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
func (r PaperV3LegacyGetCommentsResponseAuthorObject2) RawJSON() string { return r.JSON.raw }
func (r *PaperV3LegacyGetCommentsResponseAuthorObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3LegacyGetCommentsResponseAuthorObject2Avatar struct {
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
func (r PaperV3LegacyGetCommentsResponseAuthorObject2Avatar) RawJSON() string { return r.JSON.raw }
func (r *PaperV3LegacyGetCommentsResponseAuthorObject2Avatar) UnmarshalJSON(data []byte) error {
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
	Author          PaperV3LegacyGetCommentsResponseResponseAuthorUnion   `json:"author" api:"required"`
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

// PaperV3LegacyGetCommentsResponseResponseAuthorUnion contains all possible
// properties and values from
// [PaperV3LegacyGetCommentsResponseResponseAuthorObject],
// [PaperV3LegacyGetCommentsResponseResponseAuthorObject2].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type PaperV3LegacyGetCommentsResponseResponseAuthorUnion struct {
	// This field is a union of [string], [any]
	ID PaperV3LegacyGetCommentsResponseResponseAuthorUnionID `json:"id"`
	// This field is a union of
	// [[]PaperV3LegacyGetCommentsResponseResponseAuthorObjectAvatar],
	// [[]PaperV3LegacyGetCommentsResponseResponseAuthorObject2Avatar]
	Avatar           PaperV3LegacyGetCommentsResponseResponseAuthorUnionAvatar `json:"avatar"`
	BlueskyUsername  string                                                    `json:"blueskyUsername"`
	GitHubUsername   string                                                    `json:"githubUsername"`
	GoogleScholarID  string                                                    `json:"googleScholarId"`
	Institution      string                                                    `json:"institution"`
	LinkedinUsername string                                                    `json:"linkedinUsername"`
	OrcidID          string                                                    `json:"orcidId"`
	PublicEmail      string                                                    `json:"publicEmail"`
	RealName         string                                                    `json:"realName"`
	Reputation       float64                                                   `json:"reputation"`
	Role             string                                                    `json:"role"`
	Username         string                                                    `json:"username"`
	Verified         bool                                                      `json:"verified"`
	WeeklyReputation float64                                                   `json:"weeklyReputation"`
	XUsername        string                                                    `json:"xUsername"`
	JSON             struct {
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
		raw              string
	} `json:"-"`
}

func (u PaperV3LegacyGetCommentsResponseResponseAuthorUnion) AsPaperV3LegacyGetCommentsResponseResponseAuthorObject() (v PaperV3LegacyGetCommentsResponseResponseAuthorObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PaperV3LegacyGetCommentsResponseResponseAuthorUnion) AsPaperV3LegacyGetCommentsResponseResponseAuthorObject2() (v PaperV3LegacyGetCommentsResponseResponseAuthorObject2) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PaperV3LegacyGetCommentsResponseResponseAuthorUnion) RawJSON() string { return u.JSON.raw }

func (r *PaperV3LegacyGetCommentsResponseResponseAuthorUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PaperV3LegacyGetCommentsResponseResponseAuthorUnionID is an implicit subunion of
// [PaperV3LegacyGetCommentsResponseResponseAuthorUnion].
// PaperV3LegacyGetCommentsResponseResponseAuthorUnionID provides convenient access
// to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [PaperV3LegacyGetCommentsResponseResponseAuthorUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString
// OfPaperV3LegacyGetCommentsResponseResponseAuthorObject2ID]
type PaperV3LegacyGetCommentsResponseResponseAuthorUnionID struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfPaperV3LegacyGetCommentsResponseResponseAuthorObject2ID any `json:",inline"`
	JSON                                                      struct {
		OfString                                                  respjson.Field
		OfPaperV3LegacyGetCommentsResponseResponseAuthorObject2ID respjson.Field
		raw                                                       string
	} `json:"-"`
}

func (r *PaperV3LegacyGetCommentsResponseResponseAuthorUnionID) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PaperV3LegacyGetCommentsResponseResponseAuthorUnionAvatar is an implicit
// subunion of [PaperV3LegacyGetCommentsResponseResponseAuthorUnion].
// PaperV3LegacyGetCommentsResponseResponseAuthorUnionAvatar provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [PaperV3LegacyGetCommentsResponseResponseAuthorUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfPaperV3LegacyGetCommentsResponseResponseAuthorObjectAvatarArray
// OfPaperV3LegacyGetCommentsResponseResponseAuthorObject2AvatarArray]
type PaperV3LegacyGetCommentsResponseResponseAuthorUnionAvatar struct {
	// This field will be present if the value is a
	// [[]PaperV3LegacyGetCommentsResponseResponseAuthorObjectAvatar] instead of an
	// object.
	OfPaperV3LegacyGetCommentsResponseResponseAuthorObjectAvatarArray []PaperV3LegacyGetCommentsResponseResponseAuthorObjectAvatar `json:",inline"`
	// This field will be present if the value is a
	// [[]PaperV3LegacyGetCommentsResponseResponseAuthorObject2Avatar] instead of an
	// object.
	OfPaperV3LegacyGetCommentsResponseResponseAuthorObject2AvatarArray []PaperV3LegacyGetCommentsResponseResponseAuthorObject2Avatar `json:",inline"`
	JSON                                                               struct {
		OfPaperV3LegacyGetCommentsResponseResponseAuthorObjectAvatarArray  respjson.Field
		OfPaperV3LegacyGetCommentsResponseResponseAuthorObject2AvatarArray respjson.Field
		raw                                                                string
	} `json:"-"`
}

func (r *PaperV3LegacyGetCommentsResponseResponseAuthorUnionAvatar) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3LegacyGetCommentsResponseResponseAuthorObject struct {
	ID               string                                                       `json:"id" api:"required" format:"uuid"`
	Avatar           []PaperV3LegacyGetCommentsResponseResponseAuthorObjectAvatar `json:"avatar" api:"required"`
	BlueskyUsername  string                                                       `json:"blueskyUsername" api:"required"`
	GitHubUsername   string                                                       `json:"githubUsername" api:"required"`
	GoogleScholarID  string                                                       `json:"googleScholarId" api:"required"`
	Institution      string                                                       `json:"institution" api:"required"`
	LinkedinUsername string                                                       `json:"linkedinUsername" api:"required"`
	OrcidID          string                                                       `json:"orcidId" api:"required"`
	PublicEmail      string                                                       `json:"publicEmail" api:"required"`
	RealName         string                                                       `json:"realName" api:"required"`
	Reputation       float64                                                      `json:"reputation" api:"required"`
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
func (r PaperV3LegacyGetCommentsResponseResponseAuthorObject) RawJSON() string { return r.JSON.raw }
func (r *PaperV3LegacyGetCommentsResponseResponseAuthorObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3LegacyGetCommentsResponseResponseAuthorObjectAvatar struct {
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
func (r PaperV3LegacyGetCommentsResponseResponseAuthorObjectAvatar) RawJSON() string {
	return r.JSON.raw
}
func (r *PaperV3LegacyGetCommentsResponseResponseAuthorObjectAvatar) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3LegacyGetCommentsResponseResponseAuthorObject2 struct {
	ID               any                                                           `json:"id" api:"required"`
	Avatar           []PaperV3LegacyGetCommentsResponseResponseAuthorObject2Avatar `json:"avatar" api:"required"`
	BlueskyUsername  string                                                        `json:"blueskyUsername" api:"required"`
	GitHubUsername   string                                                        `json:"githubUsername" api:"required"`
	GoogleScholarID  string                                                        `json:"googleScholarId" api:"required"`
	Institution      string                                                        `json:"institution" api:"required"`
	LinkedinUsername string                                                        `json:"linkedinUsername" api:"required"`
	OrcidID          string                                                        `json:"orcidId" api:"required"`
	PublicEmail      string                                                        `json:"publicEmail" api:"required"`
	RealName         string                                                        `json:"realName" api:"required"`
	Reputation       float64                                                       `json:"reputation" api:"required"`
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
func (r PaperV3LegacyGetCommentsResponseResponseAuthorObject2) RawJSON() string { return r.JSON.raw }
func (r *PaperV3LegacyGetCommentsResponseResponseAuthorObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3LegacyGetCommentsResponseResponseAuthorObject2Avatar struct {
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
func (r PaperV3LegacyGetCommentsResponseResponseAuthorObject2Avatar) RawJSON() string {
	return r.JSON.raw
}
func (r *PaperV3LegacyGetCommentsResponseResponseAuthorObject2Avatar) UnmarshalJSON(data []byte) error {
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
