// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package alphaxivcat

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/AlphaxivCat/alphaxiv_cat-go/internal/apijson"
	"github.com/AlphaxivCat/alphaxiv_cat-go/internal/requestconfig"
	"github.com/AlphaxivCat/alphaxiv_cat-go/option"
	"github.com/AlphaxivCat/alphaxiv_cat-go/packages/param"
	"github.com/AlphaxivCat/alphaxiv_cat-go/packages/respjson"
)

// PaperV2Service contains methods and other services that help with interacting
// with the alphaxiv_cat API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPaperV2Service] method instead.
type PaperV2Service struct {
	options []option.RequestOption
}

// NewPaperV2Service generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewPaperV2Service(opts ...option.RequestOption) (r PaperV2Service) {
	r = PaperV2Service{}
	r.options = opts
	return
}

// Add comment to paper version
//
// Source file:
// `api-server/file:/app/api-server/src/controllers/papers/v2/add-comment.controller.ts`
func (r *PaperV2Service) Comment(ctx context.Context, version string, body PaperV2CommentParams, opts ...option.RequestOption) (res *PaperV2CommentResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if version == "" {
		err = errors.New("missing required version parameter")
		return nil, err
	}
	path := fmt.Sprintf("papers/v2/%s/comment", version)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type PaperV2CommentResponse struct {
	ID              string                              `json:"id" api:"required" format:"uuid"`
	Annotation      PaperV2CommentResponseAnnotation    `json:"annotation" api:"required"`
	Author          PaperV2CommentResponseAuthorUnion   `json:"author" api:"required"`
	AuthorResponded bool                                `json:"authorResponded" api:"required"`
	Body            string                              `json:"body" api:"required"`
	Date            string                              `json:"date" api:"required"`
	Endorsements    []PaperV2CommentResponseEndorsement `json:"endorsements" api:"required"`
	HasDownvoted    bool                                `json:"hasDownvoted" api:"required"`
	HasFlagged      bool                                `json:"hasFlagged" api:"required"`
	HasUpvoted      bool                                `json:"hasUpvoted" api:"required"`
	IsAuthor        bool                                `json:"isAuthor" api:"required"`
	PaperGroupID    string                              `json:"paperGroupId" api:"required" format:"uuid"`
	PaperTitle      string                              `json:"paperTitle" api:"required"`
	PaperVersionID  string                              `json:"paperVersionId" api:"required" format:"uuid"`
	Responses       []PaperV2CommentResponseResponse    `json:"responses" api:"required"`
	Tag             string                              `json:"tag" api:"required"`
	Title           string                              `json:"title" api:"required"`
	UniversalID     string                              `json:"universalId" api:"required"`
	Upvotes         float64                             `json:"upvotes" api:"required"`
	WasEdited       bool                                `json:"wasEdited" api:"required"`
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
func (r PaperV2CommentResponse) RawJSON() string { return r.JSON.raw }
func (r *PaperV2CommentResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV2CommentResponseAnnotation struct {
	AnchorPosition PaperV2CommentResponseAnnotationAnchorPosition  `json:"anchorPosition" api:"required"`
	FocusPosition  PaperV2CommentResponseAnnotationFocusPosition   `json:"focusPosition" api:"required"`
	HighlightRects []PaperV2CommentResponseAnnotationHighlightRect `json:"highlightRects" api:"required"`
	SelectedText   string                                          `json:"selectedText" api:"required"`
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
func (r PaperV2CommentResponseAnnotation) RawJSON() string { return r.JSON.raw }
func (r *PaperV2CommentResponseAnnotation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV2CommentResponseAnnotationAnchorPosition struct {
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
func (r PaperV2CommentResponseAnnotationAnchorPosition) RawJSON() string { return r.JSON.raw }
func (r *PaperV2CommentResponseAnnotationAnchorPosition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV2CommentResponseAnnotationFocusPosition struct {
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
func (r PaperV2CommentResponseAnnotationFocusPosition) RawJSON() string { return r.JSON.raw }
func (r *PaperV2CommentResponseAnnotationFocusPosition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV2CommentResponseAnnotationHighlightRect struct {
	PageIndex float64                                             `json:"pageIndex" api:"required"`
	Rects     []PaperV2CommentResponseAnnotationHighlightRectRect `json:"rects" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PageIndex   respjson.Field
		Rects       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV2CommentResponseAnnotationHighlightRect) RawJSON() string { return r.JSON.raw }
func (r *PaperV2CommentResponseAnnotationHighlightRect) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV2CommentResponseAnnotationHighlightRectRect struct {
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
func (r PaperV2CommentResponseAnnotationHighlightRectRect) RawJSON() string { return r.JSON.raw }
func (r *PaperV2CommentResponseAnnotationHighlightRectRect) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PaperV2CommentResponseAuthorUnion contains all possible properties and values
// from [PaperV2CommentResponseAuthorObject],
// [PaperV2CommentResponseAuthorObject2].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type PaperV2CommentResponseAuthorUnion struct {
	// This field is a union of [string], [any]
	ID PaperV2CommentResponseAuthorUnionID `json:"id"`
	// This field is a union of [[]PaperV2CommentResponseAuthorObjectAvatar],
	// [[]PaperV2CommentResponseAuthorObject2Avatar]
	Avatar           PaperV2CommentResponseAuthorUnionAvatar `json:"avatar"`
	BlueskyUsername  string                                  `json:"blueskyUsername"`
	GitHubUsername   string                                  `json:"githubUsername"`
	GoogleScholarID  string                                  `json:"googleScholarId"`
	Institution      string                                  `json:"institution"`
	LinkedinUsername string                                  `json:"linkedinUsername"`
	OrcidID          string                                  `json:"orcidId"`
	PublicEmail      string                                  `json:"publicEmail"`
	RealName         string                                  `json:"realName"`
	Reputation       float64                                 `json:"reputation"`
	Role             string                                  `json:"role"`
	Username         string                                  `json:"username"`
	Verified         bool                                    `json:"verified"`
	WeeklyReputation float64                                 `json:"weeklyReputation"`
	XUsername        string                                  `json:"xUsername"`
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

func (u PaperV2CommentResponseAuthorUnion) AsPaperV2CommentResponseAuthorObject() (v PaperV2CommentResponseAuthorObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PaperV2CommentResponseAuthorUnion) AsPaperV2CommentResponseAuthorObject2() (v PaperV2CommentResponseAuthorObject2) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PaperV2CommentResponseAuthorUnion) RawJSON() string { return u.JSON.raw }

func (r *PaperV2CommentResponseAuthorUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PaperV2CommentResponseAuthorUnionID is an implicit subunion of
// [PaperV2CommentResponseAuthorUnion]. PaperV2CommentResponseAuthorUnionID
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [PaperV2CommentResponseAuthorUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfPaperV2CommentResponseAuthorObject2ID]
type PaperV2CommentResponseAuthorUnionID struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfPaperV2CommentResponseAuthorObject2ID any `json:",inline"`
	JSON                                    struct {
		OfString                                respjson.Field
		OfPaperV2CommentResponseAuthorObject2ID respjson.Field
		raw                                     string
	} `json:"-"`
}

func (r *PaperV2CommentResponseAuthorUnionID) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PaperV2CommentResponseAuthorUnionAvatar is an implicit subunion of
// [PaperV2CommentResponseAuthorUnion]. PaperV2CommentResponseAuthorUnionAvatar
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [PaperV2CommentResponseAuthorUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfPaperV2CommentResponseAuthorObjectAvatarArray
// OfPaperV2CommentResponseAuthorObject2AvatarArray]
type PaperV2CommentResponseAuthorUnionAvatar struct {
	// This field will be present if the value is a
	// [[]PaperV2CommentResponseAuthorObjectAvatar] instead of an object.
	OfPaperV2CommentResponseAuthorObjectAvatarArray []PaperV2CommentResponseAuthorObjectAvatar `json:",inline"`
	// This field will be present if the value is a
	// [[]PaperV2CommentResponseAuthorObject2Avatar] instead of an object.
	OfPaperV2CommentResponseAuthorObject2AvatarArray []PaperV2CommentResponseAuthorObject2Avatar `json:",inline"`
	JSON                                             struct {
		OfPaperV2CommentResponseAuthorObjectAvatarArray  respjson.Field
		OfPaperV2CommentResponseAuthorObject2AvatarArray respjson.Field
		raw                                              string
	} `json:"-"`
}

func (r *PaperV2CommentResponseAuthorUnionAvatar) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV2CommentResponseAuthorObject struct {
	ID               string                                     `json:"id" api:"required" format:"uuid"`
	Avatar           []PaperV2CommentResponseAuthorObjectAvatar `json:"avatar" api:"required"`
	BlueskyUsername  string                                     `json:"blueskyUsername" api:"required"`
	GitHubUsername   string                                     `json:"githubUsername" api:"required"`
	GoogleScholarID  string                                     `json:"googleScholarId" api:"required"`
	Institution      string                                     `json:"institution" api:"required"`
	LinkedinUsername string                                     `json:"linkedinUsername" api:"required"`
	OrcidID          string                                     `json:"orcidId" api:"required"`
	PublicEmail      string                                     `json:"publicEmail" api:"required"`
	RealName         string                                     `json:"realName" api:"required"`
	Reputation       float64                                    `json:"reputation" api:"required"`
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
func (r PaperV2CommentResponseAuthorObject) RawJSON() string { return r.JSON.raw }
func (r *PaperV2CommentResponseAuthorObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV2CommentResponseAuthorObjectAvatar struct {
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
func (r PaperV2CommentResponseAuthorObjectAvatar) RawJSON() string { return r.JSON.raw }
func (r *PaperV2CommentResponseAuthorObjectAvatar) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV2CommentResponseAuthorObject2 struct {
	ID               any                                         `json:"id" api:"required"`
	Avatar           []PaperV2CommentResponseAuthorObject2Avatar `json:"avatar" api:"required"`
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
func (r PaperV2CommentResponseAuthorObject2) RawJSON() string { return r.JSON.raw }
func (r *PaperV2CommentResponseAuthorObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV2CommentResponseAuthorObject2Avatar struct {
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
func (r PaperV2CommentResponseAuthorObject2Avatar) RawJSON() string { return r.JSON.raw }
func (r *PaperV2CommentResponseAuthorObject2Avatar) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV2CommentResponseEndorsement struct {
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
func (r PaperV2CommentResponseEndorsement) RawJSON() string { return r.JSON.raw }
func (r *PaperV2CommentResponseEndorsement) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV2CommentResponseResponse struct {
	ID              string                                      `json:"id" api:"required" format:"uuid"`
	Annotation      PaperV2CommentResponseResponseAnnotation    `json:"annotation" api:"required"`
	Author          PaperV2CommentResponseResponseAuthorUnion   `json:"author" api:"required"`
	AuthorResponded bool                                        `json:"authorResponded" api:"required"`
	Body            string                                      `json:"body" api:"required"`
	Date            string                                      `json:"date" api:"required"`
	Endorsements    []PaperV2CommentResponseResponseEndorsement `json:"endorsements" api:"required"`
	HasDownvoted    bool                                        `json:"hasDownvoted" api:"required"`
	HasFlagged      bool                                        `json:"hasFlagged" api:"required"`
	HasUpvoted      bool                                        `json:"hasUpvoted" api:"required"`
	IsAuthor        bool                                        `json:"isAuthor" api:"required"`
	PaperGroupID    string                                      `json:"paperGroupId" api:"required" format:"uuid"`
	PaperTitle      string                                      `json:"paperTitle" api:"required"`
	PaperVersionID  string                                      `json:"paperVersionId" api:"required" format:"uuid"`
	Tag             string                                      `json:"tag" api:"required"`
	Title           string                                      `json:"title" api:"required"`
	UniversalID     string                                      `json:"universalId" api:"required"`
	Upvotes         float64                                     `json:"upvotes" api:"required"`
	WasEdited       bool                                        `json:"wasEdited" api:"required"`
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
func (r PaperV2CommentResponseResponse) RawJSON() string { return r.JSON.raw }
func (r *PaperV2CommentResponseResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV2CommentResponseResponseAnnotation struct {
	AnchorPosition PaperV2CommentResponseResponseAnnotationAnchorPosition  `json:"anchorPosition" api:"required"`
	FocusPosition  PaperV2CommentResponseResponseAnnotationFocusPosition   `json:"focusPosition" api:"required"`
	HighlightRects []PaperV2CommentResponseResponseAnnotationHighlightRect `json:"highlightRects" api:"required"`
	SelectedText   string                                                  `json:"selectedText" api:"required"`
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
func (r PaperV2CommentResponseResponseAnnotation) RawJSON() string { return r.JSON.raw }
func (r *PaperV2CommentResponseResponseAnnotation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV2CommentResponseResponseAnnotationAnchorPosition struct {
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
func (r PaperV2CommentResponseResponseAnnotationAnchorPosition) RawJSON() string { return r.JSON.raw }
func (r *PaperV2CommentResponseResponseAnnotationAnchorPosition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV2CommentResponseResponseAnnotationFocusPosition struct {
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
func (r PaperV2CommentResponseResponseAnnotationFocusPosition) RawJSON() string { return r.JSON.raw }
func (r *PaperV2CommentResponseResponseAnnotationFocusPosition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV2CommentResponseResponseAnnotationHighlightRect struct {
	PageIndex float64                                                     `json:"pageIndex" api:"required"`
	Rects     []PaperV2CommentResponseResponseAnnotationHighlightRectRect `json:"rects" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PageIndex   respjson.Field
		Rects       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV2CommentResponseResponseAnnotationHighlightRect) RawJSON() string { return r.JSON.raw }
func (r *PaperV2CommentResponseResponseAnnotationHighlightRect) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV2CommentResponseResponseAnnotationHighlightRectRect struct {
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
func (r PaperV2CommentResponseResponseAnnotationHighlightRectRect) RawJSON() string {
	return r.JSON.raw
}
func (r *PaperV2CommentResponseResponseAnnotationHighlightRectRect) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PaperV2CommentResponseResponseAuthorUnion contains all possible properties and
// values from [PaperV2CommentResponseResponseAuthorObject],
// [PaperV2CommentResponseResponseAuthorObject2].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type PaperV2CommentResponseResponseAuthorUnion struct {
	// This field is a union of [string], [any]
	ID PaperV2CommentResponseResponseAuthorUnionID `json:"id"`
	// This field is a union of [[]PaperV2CommentResponseResponseAuthorObjectAvatar],
	// [[]PaperV2CommentResponseResponseAuthorObject2Avatar]
	Avatar           PaperV2CommentResponseResponseAuthorUnionAvatar `json:"avatar"`
	BlueskyUsername  string                                          `json:"blueskyUsername"`
	GitHubUsername   string                                          `json:"githubUsername"`
	GoogleScholarID  string                                          `json:"googleScholarId"`
	Institution      string                                          `json:"institution"`
	LinkedinUsername string                                          `json:"linkedinUsername"`
	OrcidID          string                                          `json:"orcidId"`
	PublicEmail      string                                          `json:"publicEmail"`
	RealName         string                                          `json:"realName"`
	Reputation       float64                                         `json:"reputation"`
	Role             string                                          `json:"role"`
	Username         string                                          `json:"username"`
	Verified         bool                                            `json:"verified"`
	WeeklyReputation float64                                         `json:"weeklyReputation"`
	XUsername        string                                          `json:"xUsername"`
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

func (u PaperV2CommentResponseResponseAuthorUnion) AsPaperV2CommentResponseResponseAuthorObject() (v PaperV2CommentResponseResponseAuthorObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PaperV2CommentResponseResponseAuthorUnion) AsPaperV2CommentResponseResponseAuthorObject2() (v PaperV2CommentResponseResponseAuthorObject2) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PaperV2CommentResponseResponseAuthorUnion) RawJSON() string { return u.JSON.raw }

func (r *PaperV2CommentResponseResponseAuthorUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PaperV2CommentResponseResponseAuthorUnionID is an implicit subunion of
// [PaperV2CommentResponseResponseAuthorUnion].
// PaperV2CommentResponseResponseAuthorUnionID provides convenient access to the
// sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [PaperV2CommentResponseResponseAuthorUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfPaperV2CommentResponseResponseAuthorObject2ID]
type PaperV2CommentResponseResponseAuthorUnionID struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfPaperV2CommentResponseResponseAuthorObject2ID any `json:",inline"`
	JSON                                            struct {
		OfString                                        respjson.Field
		OfPaperV2CommentResponseResponseAuthorObject2ID respjson.Field
		raw                                             string
	} `json:"-"`
}

func (r *PaperV2CommentResponseResponseAuthorUnionID) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PaperV2CommentResponseResponseAuthorUnionAvatar is an implicit subunion of
// [PaperV2CommentResponseResponseAuthorUnion].
// PaperV2CommentResponseResponseAuthorUnionAvatar provides convenient access to
// the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [PaperV2CommentResponseResponseAuthorUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfPaperV2CommentResponseResponseAuthorObjectAvatarArray
// OfPaperV2CommentResponseResponseAuthorObject2AvatarArray]
type PaperV2CommentResponseResponseAuthorUnionAvatar struct {
	// This field will be present if the value is a
	// [[]PaperV2CommentResponseResponseAuthorObjectAvatar] instead of an object.
	OfPaperV2CommentResponseResponseAuthorObjectAvatarArray []PaperV2CommentResponseResponseAuthorObjectAvatar `json:",inline"`
	// This field will be present if the value is a
	// [[]PaperV2CommentResponseResponseAuthorObject2Avatar] instead of an object.
	OfPaperV2CommentResponseResponseAuthorObject2AvatarArray []PaperV2CommentResponseResponseAuthorObject2Avatar `json:",inline"`
	JSON                                                     struct {
		OfPaperV2CommentResponseResponseAuthorObjectAvatarArray  respjson.Field
		OfPaperV2CommentResponseResponseAuthorObject2AvatarArray respjson.Field
		raw                                                      string
	} `json:"-"`
}

func (r *PaperV2CommentResponseResponseAuthorUnionAvatar) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV2CommentResponseResponseAuthorObject struct {
	ID               string                                             `json:"id" api:"required" format:"uuid"`
	Avatar           []PaperV2CommentResponseResponseAuthorObjectAvatar `json:"avatar" api:"required"`
	BlueskyUsername  string                                             `json:"blueskyUsername" api:"required"`
	GitHubUsername   string                                             `json:"githubUsername" api:"required"`
	GoogleScholarID  string                                             `json:"googleScholarId" api:"required"`
	Institution      string                                             `json:"institution" api:"required"`
	LinkedinUsername string                                             `json:"linkedinUsername" api:"required"`
	OrcidID          string                                             `json:"orcidId" api:"required"`
	PublicEmail      string                                             `json:"publicEmail" api:"required"`
	RealName         string                                             `json:"realName" api:"required"`
	Reputation       float64                                            `json:"reputation" api:"required"`
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
func (r PaperV2CommentResponseResponseAuthorObject) RawJSON() string { return r.JSON.raw }
func (r *PaperV2CommentResponseResponseAuthorObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV2CommentResponseResponseAuthorObjectAvatar struct {
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
func (r PaperV2CommentResponseResponseAuthorObjectAvatar) RawJSON() string { return r.JSON.raw }
func (r *PaperV2CommentResponseResponseAuthorObjectAvatar) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV2CommentResponseResponseAuthorObject2 struct {
	ID               any                                                 `json:"id" api:"required"`
	Avatar           []PaperV2CommentResponseResponseAuthorObject2Avatar `json:"avatar" api:"required"`
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
func (r PaperV2CommentResponseResponseAuthorObject2) RawJSON() string { return r.JSON.raw }
func (r *PaperV2CommentResponseResponseAuthorObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV2CommentResponseResponseAuthorObject2Avatar struct {
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
func (r PaperV2CommentResponseResponseAuthorObject2Avatar) RawJSON() string { return r.JSON.raw }
func (r *PaperV2CommentResponseResponseAuthorObject2Avatar) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV2CommentResponseResponseEndorsement struct {
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
func (r PaperV2CommentResponseResponseEndorsement) RawJSON() string { return r.JSON.raw }
func (r *PaperV2CommentResponseResponseEndorsement) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV2CommentParams struct {
	ParentCommentID param.Opt[string]                   `json:"parentCommentId,omitzero" api:"required" format:"uuid"`
	SelectedText    param.Opt[string]                   `json:"selectedText,omitzero" api:"required"`
	Title           param.Opt[string]                   `json:"title,omitzero" api:"required"`
	AnchorPosition  PaperV2CommentParamsAnchorPosition  `json:"anchorPosition,omitzero" api:"required"`
	FocusPosition   PaperV2CommentParamsFocusPosition   `json:"focusPosition,omitzero" api:"required"`
	HighlightRects  []PaperV2CommentParamsHighlightRect `json:"highlightRects,omitzero" api:"required"`
	// Any of "anonymous", "general", "personal", "research", "resources".
	Tag            PaperV2CommentParamsTag `json:"tag,omitzero" api:"required"`
	Body           string                  `json:"body" api:"required"`
	HighlightColor param.Opt[string]       `json:"highlightColor,omitzero"`
	paramObj
}

func (r PaperV2CommentParams) MarshalJSON() (data []byte, err error) {
	type shadow PaperV2CommentParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PaperV2CommentParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Offset, PageIndex, SpanIndex are required.
type PaperV2CommentParamsAnchorPosition struct {
	Offset    float64 `json:"offset" api:"required"`
	PageIndex float64 `json:"pageIndex" api:"required"`
	SpanIndex float64 `json:"spanIndex" api:"required"`
	paramObj
}

func (r PaperV2CommentParamsAnchorPosition) MarshalJSON() (data []byte, err error) {
	type shadow PaperV2CommentParamsAnchorPosition
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PaperV2CommentParamsAnchorPosition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Offset, PageIndex, SpanIndex are required.
type PaperV2CommentParamsFocusPosition struct {
	Offset    float64 `json:"offset" api:"required"`
	PageIndex float64 `json:"pageIndex" api:"required"`
	SpanIndex float64 `json:"spanIndex" api:"required"`
	paramObj
}

func (r PaperV2CommentParamsFocusPosition) MarshalJSON() (data []byte, err error) {
	type shadow PaperV2CommentParamsFocusPosition
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PaperV2CommentParamsFocusPosition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties PageIndex, Rects are required.
type PaperV2CommentParamsHighlightRect struct {
	PageIndex float64                                 `json:"pageIndex" api:"required"`
	Rects     []PaperV2CommentParamsHighlightRectRect `json:"rects,omitzero" api:"required"`
	paramObj
}

func (r PaperV2CommentParamsHighlightRect) MarshalJSON() (data []byte, err error) {
	type shadow PaperV2CommentParamsHighlightRect
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PaperV2CommentParamsHighlightRect) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties X1, X2, Y1, Y2 are required.
type PaperV2CommentParamsHighlightRectRect struct {
	X1 float64 `json:"x1" api:"required"`
	X2 float64 `json:"x2" api:"required"`
	Y1 float64 `json:"y1" api:"required"`
	Y2 float64 `json:"y2" api:"required"`
	paramObj
}

func (r PaperV2CommentParamsHighlightRectRect) MarshalJSON() (data []byte, err error) {
	type shadow PaperV2CommentParamsHighlightRectRect
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PaperV2CommentParamsHighlightRectRect) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV2CommentParamsTag string

const (
	PaperV2CommentParamsTagAnonymous PaperV2CommentParamsTag = "anonymous"
	PaperV2CommentParamsTagGeneral   PaperV2CommentParamsTag = "general"
	PaperV2CommentParamsTagPersonal  PaperV2CommentParamsTag = "personal"
	PaperV2CommentParamsTagResearch  PaperV2CommentParamsTag = "research"
	PaperV2CommentParamsTagResources PaperV2CommentParamsTag = "resources"
)
