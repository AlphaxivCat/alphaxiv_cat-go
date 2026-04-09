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

// UserV3ByUsernameService contains methods and other services that help with
// interacting with the alphaxiv_cat API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserV3ByUsernameService] method instead.
type UserV3ByUsernameService struct {
	options []option.RequestOption
}

// NewUserV3ByUsernameService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewUserV3ByUsernameService(opts ...option.RequestOption) (r UserV3ByUsernameService) {
	r = UserV3ByUsernameService{}
	r.options = opts
	return
}

// This route is specifically for the Client's user page
//
// Source file:
// `api-server/src/controllers/users/v3/get-profile-page-by-username.controller.ts`
func (r *UserV3ByUsernameService) GetProfilePage(ctx context.Context, username string, opts ...option.RequestOption) (res *UserV3ByUsernameGetProfilePageResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if username == "" {
		err = errors.New("missing required username parameter")
		return nil, err
	}
	path := fmt.Sprintf("users/v3/by-username/%s/profile-page", url.PathEscape(username))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Retrieve a user's basic information given its username
//
// Source file:
// `api-server/src/controllers/users/v3/get-user-by-username.controller.ts`
func (r *UserV3ByUsernameService) GetUser(ctx context.Context, username string, opts ...option.RequestOption) (res *UserV3ByUsernameGetUserResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if username == "" {
		err = errors.New("missing required username parameter")
		return nil, err
	}
	path := fmt.Sprintf("users/v3/by-username/%s", url.PathEscape(username))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type UserV3ByUsernameGetProfilePageResponse struct {
	Activity      []UserV3ByUsernameGetProfilePageResponseActivity      `json:"activity" api:"required"`
	ClaimedPapers []UserV3ByUsernameGetProfilePageResponseClaimedPaper  `json:"claimedPapers" api:"required"`
	Featured      []UserV3ByUsernameGetProfilePageResponseFeaturedUnion `json:"featured" api:"required"`
	User          UserV3ByUsernameGetProfilePageResponseUser            `json:"user" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Activity      respjson.Field
		ClaimedPapers respjson.Field
		Featured      respjson.Field
		User          respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserV3ByUsernameGetProfilePageResponse) RawJSON() string { return r.JSON.raw }
func (r *UserV3ByUsernameGetProfilePageResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserV3ByUsernameGetProfilePageResponseActivity struct {
	Date  string                                             `json:"date" api:"required"`
	Item  UserV3ByUsernameGetProfilePageResponseActivityItem `json:"item" api:"required"`
	Liked float64                                            `json:"liked" api:"required"`
	// Any of "comment".
	Type                 string   `json:"type" api:"required"`
	PaperLikes           float64  `json:"paperLikes"`
	PaperPublicationDate string   `json:"paperPublicationDate"`
	PaperTopics          []string `json:"paperTopics"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Date                 respjson.Field
		Item                 respjson.Field
		Liked                respjson.Field
		Type                 respjson.Field
		PaperLikes           respjson.Field
		PaperPublicationDate respjson.Field
		PaperTopics          respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserV3ByUsernameGetProfilePageResponseActivity) RawJSON() string { return r.JSON.raw }
func (r *UserV3ByUsernameGetProfilePageResponseActivity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserV3ByUsernameGetProfilePageResponseActivityItem struct {
	ID              string                                                          `json:"id" api:"required" format:"uuid"`
	Annotation      UserV3ByUsernameGetProfilePageResponseActivityItemAnnotation    `json:"annotation" api:"required"`
	Author          UserV3ByUsernameGetProfilePageResponseActivityItemAuthor        `json:"author" api:"required"`
	Body            string                                                          `json:"body" api:"required"`
	Date            string                                                          `json:"date" api:"required"`
	Endorsements    []UserV3ByUsernameGetProfilePageResponseActivityItemEndorsement `json:"endorsements" api:"required"`
	HasDownvoted    bool                                                            `json:"hasDownvoted" api:"required"`
	HasFlagged      bool                                                            `json:"hasFlagged" api:"required"`
	HasUpvoted      bool                                                            `json:"hasUpvoted" api:"required"`
	IsAuthor        bool                                                            `json:"isAuthor" api:"required"`
	PaperGroupID    string                                                          `json:"paperGroupId" api:"required" format:"uuid"`
	PaperTitle      string                                                          `json:"paperTitle" api:"required"`
	PaperVersionID  string                                                          `json:"paperVersionId" api:"required" format:"uuid"`
	ParentCommentID string                                                          `json:"parentCommentId" api:"required" format:"uuid"`
	Tag             string                                                          `json:"tag" api:"required"`
	Title           string                                                          `json:"title" api:"required"`
	UniversalID     string                                                          `json:"universalId" api:"required"`
	Upvotes         float64                                                         `json:"upvotes" api:"required"`
	WasEdited       bool                                                            `json:"wasEdited" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		Annotation      respjson.Field
		Author          respjson.Field
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
		ParentCommentID respjson.Field
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
func (r UserV3ByUsernameGetProfilePageResponseActivityItem) RawJSON() string { return r.JSON.raw }
func (r *UserV3ByUsernameGetProfilePageResponseActivityItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserV3ByUsernameGetProfilePageResponseActivityItemAnnotation struct {
	AnchorPosition UserV3ByUsernameGetProfilePageResponseActivityItemAnnotationAnchorPosition  `json:"anchorPosition" api:"required"`
	FocusPosition  UserV3ByUsernameGetProfilePageResponseActivityItemAnnotationFocusPosition   `json:"focusPosition" api:"required"`
	HighlightRects []UserV3ByUsernameGetProfilePageResponseActivityItemAnnotationHighlightRect `json:"highlightRects" api:"required"`
	SelectedText   string                                                                      `json:"selectedText" api:"required"`
	// Any of "highlight".
	Type  string `json:"type" api:"required"`
	Color string `json:"color" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AnchorPosition respjson.Field
		FocusPosition  respjson.Field
		HighlightRects respjson.Field
		SelectedText   respjson.Field
		Type           respjson.Field
		Color          respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserV3ByUsernameGetProfilePageResponseActivityItemAnnotation) RawJSON() string {
	return r.JSON.raw
}
func (r *UserV3ByUsernameGetProfilePageResponseActivityItemAnnotation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserV3ByUsernameGetProfilePageResponseActivityItemAnnotationAnchorPosition struct {
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
func (r UserV3ByUsernameGetProfilePageResponseActivityItemAnnotationAnchorPosition) RawJSON() string {
	return r.JSON.raw
}
func (r *UserV3ByUsernameGetProfilePageResponseActivityItemAnnotationAnchorPosition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserV3ByUsernameGetProfilePageResponseActivityItemAnnotationFocusPosition struct {
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
func (r UserV3ByUsernameGetProfilePageResponseActivityItemAnnotationFocusPosition) RawJSON() string {
	return r.JSON.raw
}
func (r *UserV3ByUsernameGetProfilePageResponseActivityItemAnnotationFocusPosition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserV3ByUsernameGetProfilePageResponseActivityItemAnnotationHighlightRect struct {
	PageIndex float64                                                                         `json:"pageIndex" api:"required"`
	Rects     []UserV3ByUsernameGetProfilePageResponseActivityItemAnnotationHighlightRectRect `json:"rects" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PageIndex   respjson.Field
		Rects       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserV3ByUsernameGetProfilePageResponseActivityItemAnnotationHighlightRect) RawJSON() string {
	return r.JSON.raw
}
func (r *UserV3ByUsernameGetProfilePageResponseActivityItemAnnotationHighlightRect) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserV3ByUsernameGetProfilePageResponseActivityItemAnnotationHighlightRectRect struct {
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
func (r UserV3ByUsernameGetProfilePageResponseActivityItemAnnotationHighlightRectRect) RawJSON() string {
	return r.JSON.raw
}
func (r *UserV3ByUsernameGetProfilePageResponseActivityItemAnnotationHighlightRectRect) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserV3ByUsernameGetProfilePageResponseActivityItemAuthor struct {
	ID               string                                                           `json:"id" api:"required" format:"uuid"`
	Avatar           []UserV3ByUsernameGetProfilePageResponseActivityItemAuthorAvatar `json:"avatar" api:"required"`
	BlueskyUsername  string                                                           `json:"blueskyUsername" api:"required"`
	GitHubUsername   string                                                           `json:"githubUsername" api:"required"`
	GoogleScholarID  string                                                           `json:"googleScholarId" api:"required"`
	Institution      string                                                           `json:"institution" api:"required"`
	LinkedinUsername string                                                           `json:"linkedinUsername" api:"required"`
	OrcidID          string                                                           `json:"orcidId" api:"required"`
	PublicEmail      string                                                           `json:"publicEmail" api:"required"`
	RealName         string                                                           `json:"realName" api:"required"`
	Reputation       float64                                                          `json:"reputation" api:"required"`
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
func (r UserV3ByUsernameGetProfilePageResponseActivityItemAuthor) RawJSON() string { return r.JSON.raw }
func (r *UserV3ByUsernameGetProfilePageResponseActivityItemAuthor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserV3ByUsernameGetProfilePageResponseActivityItemAuthorAvatar struct {
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
func (r UserV3ByUsernameGetProfilePageResponseActivityItemAuthorAvatar) RawJSON() string {
	return r.JSON.raw
}
func (r *UserV3ByUsernameGetProfilePageResponseActivityItemAuthorAvatar) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserV3ByUsernameGetProfilePageResponseActivityItemEndorsement struct {
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
func (r UserV3ByUsernameGetProfilePageResponseActivityItemEndorsement) RawJSON() string {
	return r.JSON.raw
}
func (r *UserV3ByUsernameGetProfilePageResponseActivityItemEndorsement) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserV3ByUsernameGetProfilePageResponseClaimedPaper struct {
	ID               string   `json:"id" api:"required" format:"uuid"`
	Abstract         string   `json:"abstract" api:"required"`
	Authors          []string `json:"authors" api:"required"`
	Citations        float64  `json:"citations" api:"required"`
	GoogleCitationID string   `json:"google_citation_id" api:"required"`
	ImageURL         string   `json:"imageURL" api:"required"`
	// A versionless universal paper ID (e.g. 1706.03762)
	PaperID          string   `json:"paper_id" api:"required"`
	PublicTotalVotes float64  `json:"public_total_votes" api:"required"`
	PublicationDate  string   `json:"publication_date" api:"required"`
	Title            string   `json:"title" api:"required"`
	Topics           []string `json:"topics" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		Abstract         respjson.Field
		Authors          respjson.Field
		Citations        respjson.Field
		GoogleCitationID respjson.Field
		ImageURL         respjson.Field
		PaperID          respjson.Field
		PublicTotalVotes respjson.Field
		PublicationDate  respjson.Field
		Title            respjson.Field
		Topics           respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserV3ByUsernameGetProfilePageResponseClaimedPaper) RawJSON() string { return r.JSON.raw }
func (r *UserV3ByUsernameGetProfilePageResponseClaimedPaper) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UserV3ByUsernameGetProfilePageResponseFeaturedUnion contains all possible
// properties and values from
// [UserV3ByUsernameGetProfilePageResponseFeaturedObject],
// [UserV3ByUsernameGetProfilePageResponseFeaturedObject2].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type UserV3ByUsernameGetProfilePageResponseFeaturedUnion struct {
	// This field is a union of
	// [UserV3ByUsernameGetProfilePageResponseFeaturedObjectData],
	// [UserV3ByUsernameGetProfilePageResponseFeaturedObject2Data]
	Data UserV3ByUsernameGetProfilePageResponseFeaturedUnionData `json:"data"`
	Type string                                                  `json:"type"`
	JSON struct {
		Data respjson.Field
		Type respjson.Field
		raw  string
	} `json:"-"`
}

func (u UserV3ByUsernameGetProfilePageResponseFeaturedUnion) AsUserV3ByUsernameGetProfilePageResponseFeaturedObject() (v UserV3ByUsernameGetProfilePageResponseFeaturedObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UserV3ByUsernameGetProfilePageResponseFeaturedUnion) AsUserV3ByUsernameGetProfilePageResponseFeaturedObject2() (v UserV3ByUsernameGetProfilePageResponseFeaturedObject2) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u UserV3ByUsernameGetProfilePageResponseFeaturedUnion) RawJSON() string { return u.JSON.raw }

func (r *UserV3ByUsernameGetProfilePageResponseFeaturedUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UserV3ByUsernameGetProfilePageResponseFeaturedUnionData is an implicit subunion
// of [UserV3ByUsernameGetProfilePageResponseFeaturedUnion].
// UserV3ByUsernameGetProfilePageResponseFeaturedUnionData provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [UserV3ByUsernameGetProfilePageResponseFeaturedUnion].
type UserV3ByUsernameGetProfilePageResponseFeaturedUnionData struct {
	// This field is from variant
	// [UserV3ByUsernameGetProfilePageResponseFeaturedObjectData].
	ID string `json:"id"`
	// This field is from variant
	// [UserV3ByUsernameGetProfilePageResponseFeaturedObjectData].
	Date string `json:"date"`
	// This field is from variant
	// [UserV3ByUsernameGetProfilePageResponseFeaturedObjectData].
	Link string `json:"link"`
	// This field is from variant
	// [UserV3ByUsernameGetProfilePageResponseFeaturedObjectData].
	Organization string `json:"organization"`
	// This field is from variant
	// [UserV3ByUsernameGetProfilePageResponseFeaturedObjectData].
	Recording string `json:"recording"`
	// This field is from variant
	// [UserV3ByUsernameGetProfilePageResponseFeaturedObjectData].
	Speaker string `json:"speaker"`
	Title   string `json:"title"`
	// This field is from variant
	// [UserV3ByUsernameGetProfilePageResponseFeaturedObject2Data].
	Abstract string `json:"abstract"`
	// This field is from variant
	// [UserV3ByUsernameGetProfilePageResponseFeaturedObject2Data].
	PublicationDate string `json:"publication_date"`
	// This field is from variant
	// [UserV3ByUsernameGetProfilePageResponseFeaturedObject2Data].
	UniversalPaperID string `json:"universal_paper_id"`
	JSON             struct {
		ID               respjson.Field
		Date             respjson.Field
		Link             respjson.Field
		Organization     respjson.Field
		Recording        respjson.Field
		Speaker          respjson.Field
		Title            respjson.Field
		Abstract         respjson.Field
		PublicationDate  respjson.Field
		UniversalPaperID respjson.Field
		raw              string
	} `json:"-"`
}

func (r *UserV3ByUsernameGetProfilePageResponseFeaturedUnionData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserV3ByUsernameGetProfilePageResponseFeaturedObject struct {
	Data UserV3ByUsernameGetProfilePageResponseFeaturedObjectData `json:"data" api:"required"`
	// Any of "video".
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserV3ByUsernameGetProfilePageResponseFeaturedObject) RawJSON() string { return r.JSON.raw }
func (r *UserV3ByUsernameGetProfilePageResponseFeaturedObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserV3ByUsernameGetProfilePageResponseFeaturedObjectData struct {
	ID           string `json:"id" api:"required"`
	Date         string `json:"date" api:"required"`
	Link         string `json:"link" api:"required"`
	Organization string `json:"organization" api:"required"`
	Recording    string `json:"recording" api:"required"`
	Speaker      string `json:"speaker" api:"required"`
	Title        string `json:"title" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Date         respjson.Field
		Link         respjson.Field
		Organization respjson.Field
		Recording    respjson.Field
		Speaker      respjson.Field
		Title        respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserV3ByUsernameGetProfilePageResponseFeaturedObjectData) RawJSON() string { return r.JSON.raw }
func (r *UserV3ByUsernameGetProfilePageResponseFeaturedObjectData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserV3ByUsernameGetProfilePageResponseFeaturedObject2 struct {
	Data UserV3ByUsernameGetProfilePageResponseFeaturedObject2Data `json:"data" api:"required"`
	// Any of "paper".
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserV3ByUsernameGetProfilePageResponseFeaturedObject2) RawJSON() string { return r.JSON.raw }
func (r *UserV3ByUsernameGetProfilePageResponseFeaturedObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserV3ByUsernameGetProfilePageResponseFeaturedObject2Data struct {
	Abstract         string `json:"abstract" api:"required"`
	PublicationDate  string `json:"publication_date" api:"required"`
	Title            string `json:"title" api:"required"`
	UniversalPaperID string `json:"universal_paper_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Abstract         respjson.Field
		PublicationDate  respjson.Field
		Title            respjson.Field
		UniversalPaperID respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserV3ByUsernameGetProfilePageResponseFeaturedObject2Data) RawJSON() string {
	return r.JSON.raw
}
func (r *UserV3ByUsernameGetProfilePageResponseFeaturedObject2Data) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserV3ByUsernameGetProfilePageResponseUser struct {
	ID                   string                                             `json:"id" api:"required" format:"uuid"`
	Avatar               []UserV3ByUsernameGetProfilePageResponseUserAvatar `json:"avatar" api:"required"`
	Biography            string                                             `json:"biography" api:"required"`
	BlueskyUsername      string                                             `json:"blueskyUsername" api:"required"`
	Email                string                                             `json:"email" api:"required"`
	FirstLogin           bool                                               `json:"firstLogin" api:"required"`
	FollowerCount        float64                                            `json:"followerCount" api:"required"`
	FollowingCount       float64                                            `json:"followingCount" api:"required"`
	FollowingTopicsCount float64                                            `json:"followingTopicsCount" api:"required"`
	GitHubUsername       string                                             `json:"githubUsername" api:"required"`
	GoogleScholarID      string                                             `json:"googleScholarId" api:"required"`
	Institution          string                                             `json:"institution" api:"required"`
	LinkedinUsername     string                                             `json:"linkedinUsername" api:"required"`
	Location             string                                             `json:"location" api:"required"`
	OrcidID              string                                             `json:"orcidId" api:"required"`
	PublicEmail          string                                             `json:"publicEmail" api:"required"`
	RealName             string                                             `json:"realName" api:"required"`
	Reputation           float64                                            `json:"reputation" api:"required"`
	// Any of "user", "reviewer", "admin", "bot".
	Role                   string                                                    `json:"role" api:"required"`
	SemanticScholar        UserV3ByUsernameGetProfilePageResponseUserSemanticScholar `json:"semanticScholar" api:"required"`
	Username               string                                                    `json:"username" api:"required"`
	Verified               bool                                                      `json:"verified" api:"required"`
	WeeklyReputation       float64                                                   `json:"weeklyReputation" api:"required"`
	XUsername              string                                                    `json:"xUsername" api:"required"`
	Featured               []UserV3ByUsernameGetProfilePageResponseUserFeatured      `json:"featured"`
	FollowingOrganizations []string                                                  `json:"followingOrganizations" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                     respjson.Field
		Avatar                 respjson.Field
		Biography              respjson.Field
		BlueskyUsername        respjson.Field
		Email                  respjson.Field
		FirstLogin             respjson.Field
		FollowerCount          respjson.Field
		FollowingCount         respjson.Field
		FollowingTopicsCount   respjson.Field
		GitHubUsername         respjson.Field
		GoogleScholarID        respjson.Field
		Institution            respjson.Field
		LinkedinUsername       respjson.Field
		Location               respjson.Field
		OrcidID                respjson.Field
		PublicEmail            respjson.Field
		RealName               respjson.Field
		Reputation             respjson.Field
		Role                   respjson.Field
		SemanticScholar        respjson.Field
		Username               respjson.Field
		Verified               respjson.Field
		WeeklyReputation       respjson.Field
		XUsername              respjson.Field
		Featured               respjson.Field
		FollowingOrganizations respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserV3ByUsernameGetProfilePageResponseUser) RawJSON() string { return r.JSON.raw }
func (r *UserV3ByUsernameGetProfilePageResponseUser) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserV3ByUsernameGetProfilePageResponseUserAvatar struct {
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
func (r UserV3ByUsernameGetProfilePageResponseUserAvatar) RawJSON() string { return r.JSON.raw }
func (r *UserV3ByUsernameGetProfilePageResponseUserAvatar) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserV3ByUsernameGetProfilePageResponseUserSemanticScholar struct {
	ID            string  `json:"id" api:"required"`
	CitationCount float64 `json:"citationCount" api:"nullable"`
	ExternalIDs   []any   `json:"externalIds" api:"nullable"`
	HIndex        float64 `json:"hIndex" api:"nullable"`
	Homepage      string  `json:"homepage" api:"nullable"`
	PaperCount    float64 `json:"paperCount" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		CitationCount respjson.Field
		ExternalIDs   respjson.Field
		HIndex        respjson.Field
		Homepage      respjson.Field
		PaperCount    respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserV3ByUsernameGetProfilePageResponseUserSemanticScholar) RawJSON() string {
	return r.JSON.raw
}
func (r *UserV3ByUsernameGetProfilePageResponseUserSemanticScholar) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserV3ByUsernameGetProfilePageResponseUserFeatured struct {
	EventID        string `json:"eventId" api:"required"`
	PaperVersionID string `json:"paperVersionId" api:"required"`
	// Any of "video", "paper".
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EventID        respjson.Field
		PaperVersionID respjson.Field
		Type           respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserV3ByUsernameGetProfilePageResponseUserFeatured) RawJSON() string { return r.JSON.raw }
func (r *UserV3ByUsernameGetProfilePageResponseUserFeatured) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserV3ByUsernameGetUserResponse struct {
	ID                   string                                  `json:"id" api:"required" format:"uuid"`
	Avatar               []UserV3ByUsernameGetUserResponseAvatar `json:"avatar" api:"required"`
	Biography            string                                  `json:"biography" api:"required"`
	BlueskyUsername      string                                  `json:"blueskyUsername" api:"required"`
	Email                string                                  `json:"email" api:"required"`
	FirstLogin           bool                                    `json:"firstLogin" api:"required"`
	FollowerCount        float64                                 `json:"followerCount" api:"required"`
	FollowingCount       float64                                 `json:"followingCount" api:"required"`
	FollowingTopicsCount float64                                 `json:"followingTopicsCount" api:"required"`
	GitHubUsername       string                                  `json:"githubUsername" api:"required"`
	GoogleScholarID      string                                  `json:"googleScholarId" api:"required"`
	Institution          string                                  `json:"institution" api:"required"`
	LinkedinUsername     string                                  `json:"linkedinUsername" api:"required"`
	Location             string                                  `json:"location" api:"required"`
	OrcidID              string                                  `json:"orcidId" api:"required"`
	PublicEmail          string                                  `json:"publicEmail" api:"required"`
	RealName             string                                  `json:"realName" api:"required"`
	Reputation           float64                                 `json:"reputation" api:"required"`
	// Any of "user", "reviewer", "admin", "bot".
	Role                   UserV3ByUsernameGetUserResponseRole            `json:"role" api:"required"`
	SemanticScholar        UserV3ByUsernameGetUserResponseSemanticScholar `json:"semanticScholar" api:"required"`
	Username               string                                         `json:"username" api:"required"`
	Verified               bool                                           `json:"verified" api:"required"`
	WeeklyReputation       float64                                        `json:"weeklyReputation" api:"required"`
	XUsername              string                                         `json:"xUsername" api:"required"`
	Featured               []UserV3ByUsernameGetUserResponseFeatured      `json:"featured"`
	FollowingOrganizations []string                                       `json:"followingOrganizations" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                     respjson.Field
		Avatar                 respjson.Field
		Biography              respjson.Field
		BlueskyUsername        respjson.Field
		Email                  respjson.Field
		FirstLogin             respjson.Field
		FollowerCount          respjson.Field
		FollowingCount         respjson.Field
		FollowingTopicsCount   respjson.Field
		GitHubUsername         respjson.Field
		GoogleScholarID        respjson.Field
		Institution            respjson.Field
		LinkedinUsername       respjson.Field
		Location               respjson.Field
		OrcidID                respjson.Field
		PublicEmail            respjson.Field
		RealName               respjson.Field
		Reputation             respjson.Field
		Role                   respjson.Field
		SemanticScholar        respjson.Field
		Username               respjson.Field
		Verified               respjson.Field
		WeeklyReputation       respjson.Field
		XUsername              respjson.Field
		Featured               respjson.Field
		FollowingOrganizations respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserV3ByUsernameGetUserResponse) RawJSON() string { return r.JSON.raw }
func (r *UserV3ByUsernameGetUserResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserV3ByUsernameGetUserResponseAvatar struct {
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
func (r UserV3ByUsernameGetUserResponseAvatar) RawJSON() string { return r.JSON.raw }
func (r *UserV3ByUsernameGetUserResponseAvatar) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserV3ByUsernameGetUserResponseRole string

const (
	UserV3ByUsernameGetUserResponseRoleUser     UserV3ByUsernameGetUserResponseRole = "user"
	UserV3ByUsernameGetUserResponseRoleReviewer UserV3ByUsernameGetUserResponseRole = "reviewer"
	UserV3ByUsernameGetUserResponseRoleAdmin    UserV3ByUsernameGetUserResponseRole = "admin"
	UserV3ByUsernameGetUserResponseRoleBot      UserV3ByUsernameGetUserResponseRole = "bot"
)

type UserV3ByUsernameGetUserResponseSemanticScholar struct {
	ID            string  `json:"id" api:"required"`
	CitationCount float64 `json:"citationCount" api:"nullable"`
	ExternalIDs   []any   `json:"externalIds" api:"nullable"`
	HIndex        float64 `json:"hIndex" api:"nullable"`
	Homepage      string  `json:"homepage" api:"nullable"`
	PaperCount    float64 `json:"paperCount" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		CitationCount respjson.Field
		ExternalIDs   respjson.Field
		HIndex        respjson.Field
		Homepage      respjson.Field
		PaperCount    respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserV3ByUsernameGetUserResponseSemanticScholar) RawJSON() string { return r.JSON.raw }
func (r *UserV3ByUsernameGetUserResponseSemanticScholar) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserV3ByUsernameGetUserResponseFeatured struct {
	EventID        string `json:"eventId" api:"required"`
	PaperVersionID string `json:"paperVersionId" api:"required"`
	// Any of "video", "paper".
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EventID        respjson.Field
		PaperVersionID respjson.Field
		Type           respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserV3ByUsernameGetUserResponseFeatured) RawJSON() string { return r.JSON.raw }
func (r *UserV3ByUsernameGetUserResponseFeatured) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
