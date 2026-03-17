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
	"github.com/AlphaxivCat/alphaxiv_cat-go/internal/requestconfig"
	"github.com/AlphaxivCat/alphaxiv_cat-go/option"
	"github.com/AlphaxivCat/alphaxiv_cat-go/packages/param"
	"github.com/AlphaxivCat/alphaxiv_cat-go/packages/respjson"
)

// UserV3Service contains methods and other services that help with interacting
// with the alphaxiv_cat API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserV3Service] method instead.
type UserV3Service struct {
	options         []option.RequestOption
	Following       UserV3FollowingService
	ByUsername      UserV3ByUsernameService
	SemanticScholar UserV3SemanticScholarService
	Citations       UserV3CitationService
}

// NewUserV3Service generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewUserV3Service(opts ...option.RequestOption) (r UserV3Service) {
	r = UserV3Service{}
	r.options = opts
	r.Following = NewUserV3FollowingService(opts...)
	r.ByUsername = NewUserV3ByUsernameService(opts...)
	r.SemanticScholar = NewUserV3SemanticScholarService(opts...)
	r.Citations = NewUserV3CitationService(opts...)
	return
}

// Generate a biography and institution for a user using their claimed papers
//
// Source file:
// `api-server/src/controllers/users/v3/autocomplete-profile.controller.ts`
//
// Deprecated: deprecated
func (r *UserV3Service) AutocompleteProfile(ctx context.Context, body UserV3AutocompleteProfileParams, opts ...option.RequestOption) (res *UserV3AutocompleteProfileResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "users/v3/autocomplete-profile"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Delete the given banner
//
// Source file: `api-server/src/controllers/users/v3/delete-banner.controller.ts`
func (r *UserV3Service) DeleteBanner(ctx context.Context, bannerID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if bannerID == "" {
		err = errors.New("missing required bannerId parameter")
		return err
	}
	path := fmt.Sprintf("users/v3/banners/%s", bannerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Deletes the user's account
//
// Source file: `api-server/src/controllers/users/v3/delete-own-user.controller.ts`
func (r *UserV3Service) DeleteOwnUser(ctx context.Context, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "users/v3"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Retrieve public activity timeline for a user
//
// Source file: `api-server/src/controllers/users/v3/get-activity.controller.ts`
func (r *UserV3Service) GetActivity(ctx context.Context, id string, query UserV3GetActivityParams, opts ...option.RequestOption) (res *[]UserV3GetActivityResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("users/v3/%s/activity", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Retrieve the claimed papers for a user
//
// Source file:
// `api-server/src/controllers/users/v3/get-claimed-papers.controller.ts`
func (r *UserV3Service) GetClaimedPapers(ctx context.Context, id string, query UserV3GetClaimedPapersParams, opts ...option.RequestOption) (res *[]UserV3GetClaimedPapersResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("users/v3/%s/claimed-papers", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Retrieve information about yourself
//
// Source file:
// `api-server/src/controllers/users/v3/get-current-user.controller.ts`
func (r *UserV3Service) GetCurrentUser(ctx context.Context, opts ...option.RequestOption) (res *UserV3GetCurrentUserResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "users/v3"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Retrieve highlighted activity for a user
//
// Source file: `api-server/src/controllers/users/v3/get-featured.controller.ts`
func (r *UserV3Service) GetFeaturedActivity(ctx context.Context, id string, opts ...option.RequestOption) (res *[]UserV3GetFeaturedActivityResponseUnion, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("users/v3/%s/featured", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// List the users following the specified user
//
// Source file: `api-server/src/controllers/users/v3/get-followers.controller.ts`
func (r *UserV3Service) GetFollowers(ctx context.Context, id string, opts ...option.RequestOption) (res *UserV3GetFollowersResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("users/v3/%s/followers", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Retrieve weekly and all-time leaderboards for users ranked by reputation
//
// Source file: `api-server/src/controllers/users/v3/get-leaderboard.controller.ts`
func (r *UserV3Service) GetLeaderboard(ctx context.Context, opts ...option.RequestOption) (res *UserV3GetLeaderboardResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "users/v3/leaderboard"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Retrieve a user's basic information given its UUID
//
// Source file: `api-server/src/controllers/users/v3/get-user.controller.ts`
func (r *UserV3Service) GetUserByUuid(ctx context.Context, id string, opts ...option.RequestOption) (res *UserV3GetUserByUuidResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("users/v3/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Retrieve the view history for the current user
//
// Source file:
// `api-server/src/controllers/users/v3/get-viewed-history.controller.ts`
func (r *UserV3Service) GetViewedHistory(ctx context.Context, query UserV3GetViewedHistoryParams, opts ...option.RequestOption) (res *[]UserV3GetViewedHistoryResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "users/v3/viewed-history"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Send a notification digest email for the given user when necessary.
//
// Source file:
// `api-server/src/controllers/users/v3/process-notification-email.controller.ts`
//
// Deprecated: deprecated
func (r *UserV3Service) ProcessNotificationEmail(ctx context.Context, id string, opts ...option.RequestOption) (res *UserV3ProcessNotificationEmailResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("users/v3/%s/process-notification-email", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Search for users by name, username, or institution
//
// Source file: `api-server/src/controllers/users/v3/search-users.controller.ts`
func (r *UserV3Service) Search(ctx context.Context, query UserV3SearchParams, opts ...option.RequestOption) (res *UserV3SearchResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "users/v3/search"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Follow or unfollow another user
//
// Source file:
// `api-server/src/controllers/users/v3/toggle-follow-user.controller.ts`
func (r *UserV3Service) ToggleFollowUser(ctx context.Context, id string, opts ...option.RequestOption) (res *UserV3ToggleFollowUserResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("users/v3/%s/follow", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Update base or banner preferences for the authenticated user
//
// Source file:
// `api-server/src/controllers/users/v3/update-preferences.controller.ts`
func (r *UserV3Service) UpdatePreferences(ctx context.Context, body UserV3UpdatePreferencesParams, opts ...option.RequestOption) (res *UserV3UpdatePreferencesResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "users/v3/preferences"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Update profile details for the authenticated user
//
// Source file: `api-server/src/controllers/users/v3/update-profile.controller.ts`
func (r *UserV3Service) UpdateProfile(ctx context.Context, body UserV3UpdateProfileParams, opts ...option.RequestOption) (res *UserV3UpdateProfileResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "users/v3/profile"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Upload or remove the authenticated user's avatar image.
//
// Source file: `api-server/src/controllers/users/v3/upload-avatar.controller.ts`
func (r *UserV3Service) UploadAvatar(ctx context.Context, opts ...option.RequestOption) (res *UserV3UploadAvatarResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "users/v3/avatar"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

type UserV3AutocompleteProfileResponse struct {
	Bio         string `json:"bio" api:"required"`
	Institution string `json:"institution" api:"required"`
	Message     string `json:"message" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Bio         respjson.Field
		Institution respjson.Field
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserV3AutocompleteProfileResponse) RawJSON() string { return r.JSON.raw }
func (r *UserV3AutocompleteProfileResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserV3GetActivityResponse struct {
	Date  string                        `json:"date" api:"required"`
	Item  UserV3GetActivityResponseItem `json:"item" api:"required"`
	Liked float64                       `json:"liked" api:"required"`
	// Any of "comment".
	Type                 UserV3GetActivityResponseType `json:"type" api:"required"`
	PaperLikes           float64                       `json:"paperLikes"`
	PaperPublicationDate string                        `json:"paperPublicationDate"`
	PaperTopics          []string                      `json:"paperTopics"`
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
func (r UserV3GetActivityResponse) RawJSON() string { return r.JSON.raw }
func (r *UserV3GetActivityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserV3GetActivityResponseItem struct {
	ID              string                                     `json:"id" api:"required" format:"uuid"`
	Annotation      UserV3GetActivityResponseItemAnnotation    `json:"annotation" api:"required"`
	Author          UserV3GetActivityResponseItemAuthor        `json:"author" api:"required"`
	Body            string                                     `json:"body" api:"required"`
	Date            string                                     `json:"date" api:"required"`
	Endorsements    []UserV3GetActivityResponseItemEndorsement `json:"endorsements" api:"required"`
	HasDownvoted    bool                                       `json:"hasDownvoted" api:"required"`
	HasFlagged      bool                                       `json:"hasFlagged" api:"required"`
	HasUpvoted      bool                                       `json:"hasUpvoted" api:"required"`
	IsAuthor        bool                                       `json:"isAuthor" api:"required"`
	PaperGroupID    string                                     `json:"paperGroupId" api:"required" format:"uuid"`
	PaperTitle      string                                     `json:"paperTitle" api:"required"`
	PaperVersionID  string                                     `json:"paperVersionId" api:"required" format:"uuid"`
	ParentCommentID string                                     `json:"parentCommentId" api:"required" format:"uuid"`
	Tag             string                                     `json:"tag" api:"required"`
	Title           string                                     `json:"title" api:"required"`
	UniversalID     string                                     `json:"universalId" api:"required"`
	Upvotes         float64                                    `json:"upvotes" api:"required"`
	WasEdited       bool                                       `json:"wasEdited" api:"required"`
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
func (r UserV3GetActivityResponseItem) RawJSON() string { return r.JSON.raw }
func (r *UserV3GetActivityResponseItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserV3GetActivityResponseItemAnnotation struct {
	AnchorPosition UserV3GetActivityResponseItemAnnotationAnchorPosition  `json:"anchorPosition" api:"required"`
	FocusPosition  UserV3GetActivityResponseItemAnnotationFocusPosition   `json:"focusPosition" api:"required"`
	HighlightRects []UserV3GetActivityResponseItemAnnotationHighlightRect `json:"highlightRects" api:"required"`
	SelectedText   string                                                 `json:"selectedText" api:"required"`
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
func (r UserV3GetActivityResponseItemAnnotation) RawJSON() string { return r.JSON.raw }
func (r *UserV3GetActivityResponseItemAnnotation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserV3GetActivityResponseItemAnnotationAnchorPosition struct {
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
func (r UserV3GetActivityResponseItemAnnotationAnchorPosition) RawJSON() string { return r.JSON.raw }
func (r *UserV3GetActivityResponseItemAnnotationAnchorPosition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserV3GetActivityResponseItemAnnotationFocusPosition struct {
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
func (r UserV3GetActivityResponseItemAnnotationFocusPosition) RawJSON() string { return r.JSON.raw }
func (r *UserV3GetActivityResponseItemAnnotationFocusPosition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserV3GetActivityResponseItemAnnotationHighlightRect struct {
	PageIndex float64                                                    `json:"pageIndex" api:"required"`
	Rects     []UserV3GetActivityResponseItemAnnotationHighlightRectRect `json:"rects" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PageIndex   respjson.Field
		Rects       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserV3GetActivityResponseItemAnnotationHighlightRect) RawJSON() string { return r.JSON.raw }
func (r *UserV3GetActivityResponseItemAnnotationHighlightRect) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserV3GetActivityResponseItemAnnotationHighlightRectRect struct {
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
func (r UserV3GetActivityResponseItemAnnotationHighlightRectRect) RawJSON() string { return r.JSON.raw }
func (r *UserV3GetActivityResponseItemAnnotationHighlightRectRect) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserV3GetActivityResponseItemAuthor struct {
	ID               string                                      `json:"id" api:"required" format:"uuid"`
	Avatar           []UserV3GetActivityResponseItemAuthorAvatar `json:"avatar" api:"required"`
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
func (r UserV3GetActivityResponseItemAuthor) RawJSON() string { return r.JSON.raw }
func (r *UserV3GetActivityResponseItemAuthor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserV3GetActivityResponseItemAuthorAvatar struct {
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
func (r UserV3GetActivityResponseItemAuthorAvatar) RawJSON() string { return r.JSON.raw }
func (r *UserV3GetActivityResponseItemAuthorAvatar) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserV3GetActivityResponseItemEndorsement struct {
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
func (r UserV3GetActivityResponseItemEndorsement) RawJSON() string { return r.JSON.raw }
func (r *UserV3GetActivityResponseItemEndorsement) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserV3GetActivityResponseType string

const (
	UserV3GetActivityResponseTypeComment UserV3GetActivityResponseType = "comment"
)

type UserV3GetClaimedPapersResponse struct {
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
func (r UserV3GetClaimedPapersResponse) RawJSON() string { return r.JSON.raw }
func (r *UserV3GetClaimedPapersResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserV3GetCurrentUserResponse struct {
	Badges      UserV3GetCurrentUserResponseBadges      `json:"badges" api:"required"`
	Preferences UserV3GetCurrentUserResponsePreferences `json:"preferences" api:"required"`
	User        UserV3GetCurrentUserResponseUser        `json:"user" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Badges      respjson.Field
		Preferences respjson.Field
		User        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserV3GetCurrentUserResponse) RawJSON() string { return r.JSON.raw }
func (r *UserV3GetCurrentUserResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserV3GetCurrentUserResponseBadges struct {
	Site int64 `json:"site" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Site        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserV3GetCurrentUserResponseBadges) RawJSON() string { return r.JSON.raw }
func (r *UserV3GetCurrentUserResponseBadges) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserV3GetCurrentUserResponsePreferences struct {
	Banners []UserV3GetCurrentUserResponsePreferencesBanner `json:"banners" api:"required"`
	Base    UserV3GetCurrentUserResponsePreferencesBase     `json:"base" api:"required"`
	Email   UserV3GetCurrentUserResponsePreferencesEmail    `json:"email" api:"required"`
	Folders []UserV3GetCurrentUserResponsePreferencesFolder `json:"folders" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Banners     respjson.Field
		Base        respjson.Field
		Email       respjson.Field
		Folders     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserV3GetCurrentUserResponsePreferences) RawJSON() string { return r.JSON.raw }
func (r *UserV3GetCurrentUserResponsePreferences) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserV3GetCurrentUserResponsePreferencesBanner struct {
	ID   string `json:"id" api:"required" format:"uuid"`
	Link string `json:"link" api:"required"`
	Name string `json:"name" api:"required"`
	// Any of "success", "info", "warning", "error".
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Link        respjson.Field
		Name        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserV3GetCurrentUserResponsePreferencesBanner) RawJSON() string { return r.JSON.raw }
func (r *UserV3GetCurrentUserResponsePreferencesBanner) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserV3GetCurrentUserResponsePreferencesBase struct {
	AssistantCustomStyles   []UserV3GetCurrentUserResponsePreferencesBaseAssistantCustomStyle `json:"assistantCustomStyles" api:"required"`
	AssistantStyleSelection string                                                            `json:"assistantStyleSelection" api:"required"`
	// Any of "assistant", "notes", "similar".
	DefaultPrivatePaperSidebarTab string `json:"defaultPrivatePaperSidebarTab" api:"required"`
	// Any of "comments", "assistant", "similar", "notes", "social".
	DefaultPublicPaperSidebarTab string `json:"defaultPublicPaperSidebarTab" api:"required"`
	// Any of "Hot", "Comments", "Views", "Likes", "GitHub", "Twitter (X)",
	// "Recommended".
	FeedSort                string `json:"feedSort" api:"required"`
	IsDarkModeEnabled       bool   `json:"isDarkModeEnabled" api:"required"`
	IsDebugModeEnabled      bool   `json:"isDebugModeEnabled" api:"required"`
	IsMembersSidebarVisible bool   `json:"isMembersSidebarVisible" api:"required"`
	// Any of "am", "ar", "az", "bg", "bn", "ca", "cs", "da", "de", "el", "en", "es",
	// "et", "fa", "fi", "fr", "gu", "ha", "he", "hi", "hr", "hu", "id", "it", "ja",
	// "ka", "kn", "ko", "lt", "lv", "ml", "mr", "ms", "my", "ne", "nl", "no", "pa",
	// "pl", "pt", "ro", "ru", "si", "sk", "sl", "sr", "sv", "sw", "ta", "te", "th",
	// "tl", "tr", "uk", "ur", "uz", "vi", "yo", "zh".
	PreferredLanguage  string  `json:"preferredLanguage" api:"required"`
	PreferredLlmModel  string  `json:"preferredLlmModel" api:"required"`
	ReadingModeEnabled bool    `json:"readingModeEnabled" api:"required"`
	ShowModelThinking  bool    `json:"showModelThinking" api:"required"`
	ToolingPaneWidth   float64 `json:"toolingPaneWidth" api:"required"`
	// Any of "off", "full".
	WebSearch string `json:"webSearch" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AssistantCustomStyles         respjson.Field
		AssistantStyleSelection       respjson.Field
		DefaultPrivatePaperSidebarTab respjson.Field
		DefaultPublicPaperSidebarTab  respjson.Field
		FeedSort                      respjson.Field
		IsDarkModeEnabled             respjson.Field
		IsDebugModeEnabled            respjson.Field
		IsMembersSidebarVisible       respjson.Field
		PreferredLanguage             respjson.Field
		PreferredLlmModel             respjson.Field
		ReadingModeEnabled            respjson.Field
		ShowModelThinking             respjson.Field
		ToolingPaneWidth              respjson.Field
		WebSearch                     respjson.Field
		ExtraFields                   map[string]respjson.Field
		raw                           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserV3GetCurrentUserResponsePreferencesBase) RawJSON() string { return r.JSON.raw }
func (r *UserV3GetCurrentUserResponsePreferencesBase) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserV3GetCurrentUserResponsePreferencesBaseAssistantCustomStyle struct {
	ID           string `json:"id" api:"required" format:"uuid"`
	Instructions string `json:"instructions" api:"required"`
	Name         string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Instructions respjson.Field
		Name         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserV3GetCurrentUserResponsePreferencesBaseAssistantCustomStyle) RawJSON() string {
	return r.JSON.raw
}
func (r *UserV3GetCurrentUserResponsePreferencesBaseAssistantCustomStyle) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserV3GetCurrentUserResponsePreferencesEmail struct {
	Bounced             bool `json:"bounced" api:"required"`
	DirectNotifications bool `json:"directNotifications" api:"required"`
	RelevantActivity    bool `json:"relevantActivity" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Bounced             respjson.Field
		DirectNotifications respjson.Field
		RelevantActivity    respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserV3GetCurrentUserResponsePreferencesEmail) RawJSON() string { return r.JSON.raw }
func (r *UserV3GetCurrentUserResponsePreferencesEmail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserV3GetCurrentUserResponsePreferencesFolder struct {
	FolderID string `json:"folderId" api:"required" format:"uuid"`
	Opened   bool   `json:"opened" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FolderID    respjson.Field
		Opened      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserV3GetCurrentUserResponsePreferencesFolder) RawJSON() string { return r.JSON.raw }
func (r *UserV3GetCurrentUserResponsePreferencesFolder) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserV3GetCurrentUserResponseUser struct {
	ID                       string                                             `json:"id" api:"required" format:"uuid"`
	Avatar                   []UserV3GetCurrentUserResponseUserAvatar           `json:"avatar" api:"required"`
	Biography                string                                             `json:"biography" api:"required"`
	BlueskyUsername          string                                             `json:"blueskyUsername" api:"required"`
	Email                    string                                             `json:"email" api:"required"`
	FirstLogin               bool                                               `json:"firstLogin" api:"required"`
	FollowerCount            float64                                            `json:"followerCount" api:"required"`
	FollowingCount           float64                                            `json:"followingCount" api:"required"`
	FollowingTopics          []string                                           `json:"followingTopics" api:"required"`
	FollowingTopicsCount     float64                                            `json:"followingTopicsCount" api:"required"`
	GitHubUsername           string                                             `json:"githubUsername" api:"required"`
	GoogleScholarID          string                                             `json:"googleScholarId" api:"required"`
	Institution              string                                             `json:"institution" api:"required"`
	LinkedinUsername         string                                             `json:"linkedinUsername" api:"required"`
	Location                 string                                             `json:"location" api:"required"`
	OrcidID                  string                                             `json:"orcidId" api:"required"`
	Preferences              UserV3GetCurrentUserResponseUserPreferences        `json:"preferences" api:"required"`
	PublicEmail              string                                             `json:"publicEmail" api:"required"`
	PushSubscriptions        []UserV3GetCurrentUserResponseUserPushSubscription `json:"pushSubscriptions" api:"required"`
	RealName                 string                                             `json:"realName" api:"required"`
	Reputation               float64                                            `json:"reputation" api:"required"`
	RequestedImplementations []string                                           `json:"requestedImplementations" api:"required" format:"uuid"`
	// Any of "user", "reviewer", "admin", "bot".
	Role                   string                                          `json:"role" api:"required"`
	SemanticScholar        UserV3GetCurrentUserResponseUserSemanticScholar `json:"semanticScholar" api:"required"`
	Username               string                                          `json:"username" api:"required"`
	Verified               bool                                            `json:"verified" api:"required"`
	VotedPaperGroups       []string                                        `json:"votedPaperGroups" api:"required" format:"uuid"`
	WeeklyReputation       float64                                         `json:"weeklyReputation" api:"required"`
	XUsername              string                                          `json:"xUsername" api:"required"`
	Featured               []UserV3GetCurrentUserResponseUserFeatured      `json:"featured"`
	FollowingOrganizations []string                                        `json:"followingOrganizations" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                       respjson.Field
		Avatar                   respjson.Field
		Biography                respjson.Field
		BlueskyUsername          respjson.Field
		Email                    respjson.Field
		FirstLogin               respjson.Field
		FollowerCount            respjson.Field
		FollowingCount           respjson.Field
		FollowingTopics          respjson.Field
		FollowingTopicsCount     respjson.Field
		GitHubUsername           respjson.Field
		GoogleScholarID          respjson.Field
		Institution              respjson.Field
		LinkedinUsername         respjson.Field
		Location                 respjson.Field
		OrcidID                  respjson.Field
		Preferences              respjson.Field
		PublicEmail              respjson.Field
		PushSubscriptions        respjson.Field
		RealName                 respjson.Field
		Reputation               respjson.Field
		RequestedImplementations respjson.Field
		Role                     respjson.Field
		SemanticScholar          respjson.Field
		Username                 respjson.Field
		Verified                 respjson.Field
		VotedPaperGroups         respjson.Field
		WeeklyReputation         respjson.Field
		XUsername                respjson.Field
		Featured                 respjson.Field
		FollowingOrganizations   respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserV3GetCurrentUserResponseUser) RawJSON() string { return r.JSON.raw }
func (r *UserV3GetCurrentUserResponseUser) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserV3GetCurrentUserResponseUserAvatar struct {
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
func (r UserV3GetCurrentUserResponseUserAvatar) RawJSON() string { return r.JSON.raw }
func (r *UserV3GetCurrentUserResponseUserAvatar) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserV3GetCurrentUserResponseUserPreferences struct {
	Banners []UserV3GetCurrentUserResponseUserPreferencesBanner `json:"banners" api:"required"`
	Base    UserV3GetCurrentUserResponseUserPreferencesBase     `json:"base" api:"required"`
	Email   UserV3GetCurrentUserResponseUserPreferencesEmail    `json:"email" api:"required"`
	Folders []UserV3GetCurrentUserResponseUserPreferencesFolder `json:"folders" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Banners     respjson.Field
		Base        respjson.Field
		Email       respjson.Field
		Folders     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserV3GetCurrentUserResponseUserPreferences) RawJSON() string { return r.JSON.raw }
func (r *UserV3GetCurrentUserResponseUserPreferences) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserV3GetCurrentUserResponseUserPreferencesBanner struct {
	ID   string `json:"id" api:"required" format:"uuid"`
	Link string `json:"link" api:"required"`
	Name string `json:"name" api:"required"`
	// Any of "success", "info", "warning", "error".
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Link        respjson.Field
		Name        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserV3GetCurrentUserResponseUserPreferencesBanner) RawJSON() string { return r.JSON.raw }
func (r *UserV3GetCurrentUserResponseUserPreferencesBanner) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserV3GetCurrentUserResponseUserPreferencesBase struct {
	AssistantCustomStyles   []UserV3GetCurrentUserResponseUserPreferencesBaseAssistantCustomStyle `json:"assistantCustomStyles" api:"required"`
	AssistantStyleSelection string                                                                `json:"assistantStyleSelection" api:"required"`
	// Any of "assistant", "notes", "similar".
	DefaultPrivatePaperSidebarTab string `json:"defaultPrivatePaperSidebarTab" api:"required"`
	// Any of "comments", "assistant", "similar", "notes", "social".
	DefaultPublicPaperSidebarTab string `json:"defaultPublicPaperSidebarTab" api:"required"`
	// Any of "Hot", "Comments", "Views", "Likes", "GitHub", "Twitter (X)",
	// "Recommended".
	FeedSort                string `json:"feedSort" api:"required"`
	IsDarkModeEnabled       bool   `json:"isDarkModeEnabled" api:"required"`
	IsDebugModeEnabled      bool   `json:"isDebugModeEnabled" api:"required"`
	IsMembersSidebarVisible bool   `json:"isMembersSidebarVisible" api:"required"`
	// Any of "am", "ar", "az", "bg", "bn", "ca", "cs", "da", "de", "el", "en", "es",
	// "et", "fa", "fi", "fr", "gu", "ha", "he", "hi", "hr", "hu", "id", "it", "ja",
	// "ka", "kn", "ko", "lt", "lv", "ml", "mr", "ms", "my", "ne", "nl", "no", "pa",
	// "pl", "pt", "ro", "ru", "si", "sk", "sl", "sr", "sv", "sw", "ta", "te", "th",
	// "tl", "tr", "uk", "ur", "uz", "vi", "yo", "zh".
	PreferredLanguage  string  `json:"preferredLanguage" api:"required"`
	PreferredLlmModel  string  `json:"preferredLlmModel" api:"required"`
	ReadingModeEnabled bool    `json:"readingModeEnabled" api:"required"`
	ShowModelThinking  bool    `json:"showModelThinking" api:"required"`
	ToolingPaneWidth   float64 `json:"toolingPaneWidth" api:"required"`
	// Any of "off", "full".
	WebSearch string `json:"webSearch" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AssistantCustomStyles         respjson.Field
		AssistantStyleSelection       respjson.Field
		DefaultPrivatePaperSidebarTab respjson.Field
		DefaultPublicPaperSidebarTab  respjson.Field
		FeedSort                      respjson.Field
		IsDarkModeEnabled             respjson.Field
		IsDebugModeEnabled            respjson.Field
		IsMembersSidebarVisible       respjson.Field
		PreferredLanguage             respjson.Field
		PreferredLlmModel             respjson.Field
		ReadingModeEnabled            respjson.Field
		ShowModelThinking             respjson.Field
		ToolingPaneWidth              respjson.Field
		WebSearch                     respjson.Field
		ExtraFields                   map[string]respjson.Field
		raw                           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserV3GetCurrentUserResponseUserPreferencesBase) RawJSON() string { return r.JSON.raw }
func (r *UserV3GetCurrentUserResponseUserPreferencesBase) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserV3GetCurrentUserResponseUserPreferencesBaseAssistantCustomStyle struct {
	ID           string `json:"id" api:"required" format:"uuid"`
	Instructions string `json:"instructions" api:"required"`
	Name         string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Instructions respjson.Field
		Name         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserV3GetCurrentUserResponseUserPreferencesBaseAssistantCustomStyle) RawJSON() string {
	return r.JSON.raw
}
func (r *UserV3GetCurrentUserResponseUserPreferencesBaseAssistantCustomStyle) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserV3GetCurrentUserResponseUserPreferencesEmail struct {
	Bounced             bool `json:"bounced" api:"required"`
	DirectNotifications bool `json:"directNotifications" api:"required"`
	RelevantActivity    bool `json:"relevantActivity" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Bounced             respjson.Field
		DirectNotifications respjson.Field
		RelevantActivity    respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserV3GetCurrentUserResponseUserPreferencesEmail) RawJSON() string { return r.JSON.raw }
func (r *UserV3GetCurrentUserResponseUserPreferencesEmail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserV3GetCurrentUserResponseUserPreferencesFolder struct {
	FolderID string `json:"folderId" api:"required" format:"uuid"`
	Opened   bool   `json:"opened" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FolderID    respjson.Field
		Opened      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserV3GetCurrentUserResponseUserPreferencesFolder) RawJSON() string { return r.JSON.raw }
func (r *UserV3GetCurrentUserResponseUserPreferencesFolder) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserV3GetCurrentUserResponseUserPushSubscription struct {
	AuthKey   string `json:"authKey" api:"required"`
	Endpoint  string `json:"endpoint" api:"required"`
	P256dhKey string `json:"p256dhKey" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AuthKey     respjson.Field
		Endpoint    respjson.Field
		P256dhKey   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserV3GetCurrentUserResponseUserPushSubscription) RawJSON() string { return r.JSON.raw }
func (r *UserV3GetCurrentUserResponseUserPushSubscription) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserV3GetCurrentUserResponseUserSemanticScholar struct {
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
func (r UserV3GetCurrentUserResponseUserSemanticScholar) RawJSON() string { return r.JSON.raw }
func (r *UserV3GetCurrentUserResponseUserSemanticScholar) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserV3GetCurrentUserResponseUserFeatured struct {
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
func (r UserV3GetCurrentUserResponseUserFeatured) RawJSON() string { return r.JSON.raw }
func (r *UserV3GetCurrentUserResponseUserFeatured) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UserV3GetFeaturedActivityResponseUnion contains all possible properties and
// values from [UserV3GetFeaturedActivityResponseObject],
// [UserV3GetFeaturedActivityResponseObject].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type UserV3GetFeaturedActivityResponseUnion struct {
	// This field is from variant [UserV3GetFeaturedActivityResponseObject].
	Data UserV3GetFeaturedActivityResponseObjectData `json:"data"`
	// This field is from variant [UserV3GetFeaturedActivityResponseObject].
	Type string `json:"type"`
	JSON struct {
		Data respjson.Field
		Type respjson.Field
		raw  string
	} `json:"-"`
}

func (u UserV3GetFeaturedActivityResponseUnion) AsUserV3GetFeaturedActivityResponseObject() (v UserV3GetFeaturedActivityResponseObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UserV3GetFeaturedActivityResponseUnion) AsVariant2() (v UserV3GetFeaturedActivityResponseObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u UserV3GetFeaturedActivityResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *UserV3GetFeaturedActivityResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserV3GetFeaturedActivityResponseObject struct {
	Data UserV3GetFeaturedActivityResponseObjectData `json:"data" api:"required"`
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
func (r UserV3GetFeaturedActivityResponseObject) RawJSON() string { return r.JSON.raw }
func (r *UserV3GetFeaturedActivityResponseObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserV3GetFeaturedActivityResponseObjectData struct {
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
func (r UserV3GetFeaturedActivityResponseObjectData) RawJSON() string { return r.JSON.raw }
func (r *UserV3GetFeaturedActivityResponseObjectData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserV3GetFollowersResponse struct {
	Followers []UserV3GetFollowersResponseFollower `json:"followers" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Followers   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserV3GetFollowersResponse) RawJSON() string { return r.JSON.raw }
func (r *UserV3GetFollowersResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserV3GetFollowersResponseFollower struct {
	ID               string                                     `json:"id" api:"required" format:"uuid"`
	Avatar           []UserV3GetFollowersResponseFollowerAvatar `json:"avatar" api:"required"`
	GoogleScholarID  string                                     `json:"googleScholarId" api:"required"`
	Institution      string                                     `json:"institution" api:"required"`
	RealName         string                                     `json:"realName" api:"required"`
	Reputation       float64                                    `json:"reputation" api:"required"`
	Username         string                                     `json:"username" api:"required"`
	WeeklyReputation float64                                    `json:"weeklyReputation" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		Avatar           respjson.Field
		GoogleScholarID  respjson.Field
		Institution      respjson.Field
		RealName         respjson.Field
		Reputation       respjson.Field
		Username         respjson.Field
		WeeklyReputation respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserV3GetFollowersResponseFollower) RawJSON() string { return r.JSON.raw }
func (r *UserV3GetFollowersResponseFollower) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserV3GetFollowersResponseFollowerAvatar struct {
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
func (r UserV3GetFollowersResponseFollowerAvatar) RawJSON() string { return r.JSON.raw }
func (r *UserV3GetFollowersResponseFollowerAvatar) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserV3GetLeaderboardResponse struct {
	AllTime []UserV3GetLeaderboardResponseAllTime `json:"allTime" api:"required"`
	Weekly  []UserV3GetLeaderboardResponseWeekly  `json:"weekly" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AllTime     respjson.Field
		Weekly      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserV3GetLeaderboardResponse) RawJSON() string { return r.JSON.raw }
func (r *UserV3GetLeaderboardResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserV3GetLeaderboardResponseAllTime struct {
	ID               string                                      `json:"id" api:"required" format:"uuid"`
	Avatar           []UserV3GetLeaderboardResponseAllTimeAvatar `json:"avatar" api:"required"`
	GoogleScholarID  string                                      `json:"googleScholarId" api:"required"`
	Institution      string                                      `json:"institution" api:"required"`
	RealName         string                                      `json:"realName" api:"required"`
	Reputation       float64                                     `json:"reputation" api:"required"`
	Username         string                                      `json:"username" api:"required"`
	WeeklyReputation float64                                     `json:"weeklyReputation" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		Avatar           respjson.Field
		GoogleScholarID  respjson.Field
		Institution      respjson.Field
		RealName         respjson.Field
		Reputation       respjson.Field
		Username         respjson.Field
		WeeklyReputation respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserV3GetLeaderboardResponseAllTime) RawJSON() string { return r.JSON.raw }
func (r *UserV3GetLeaderboardResponseAllTime) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserV3GetLeaderboardResponseAllTimeAvatar struct {
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
func (r UserV3GetLeaderboardResponseAllTimeAvatar) RawJSON() string { return r.JSON.raw }
func (r *UserV3GetLeaderboardResponseAllTimeAvatar) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserV3GetLeaderboardResponseWeekly struct {
	ID               string                                     `json:"id" api:"required" format:"uuid"`
	Avatar           []UserV3GetLeaderboardResponseWeeklyAvatar `json:"avatar" api:"required"`
	GoogleScholarID  string                                     `json:"googleScholarId" api:"required"`
	Institution      string                                     `json:"institution" api:"required"`
	RealName         string                                     `json:"realName" api:"required"`
	Reputation       float64                                    `json:"reputation" api:"required"`
	Username         string                                     `json:"username" api:"required"`
	WeeklyReputation float64                                    `json:"weeklyReputation" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		Avatar           respjson.Field
		GoogleScholarID  respjson.Field
		Institution      respjson.Field
		RealName         respjson.Field
		Reputation       respjson.Field
		Username         respjson.Field
		WeeklyReputation respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserV3GetLeaderboardResponseWeekly) RawJSON() string { return r.JSON.raw }
func (r *UserV3GetLeaderboardResponseWeekly) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserV3GetLeaderboardResponseWeeklyAvatar struct {
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
func (r UserV3GetLeaderboardResponseWeeklyAvatar) RawJSON() string { return r.JSON.raw }
func (r *UserV3GetLeaderboardResponseWeeklyAvatar) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserV3GetUserByUuidResponse struct {
	ID                   string                              `json:"id" api:"required" format:"uuid"`
	Avatar               []UserV3GetUserByUuidResponseAvatar `json:"avatar" api:"required"`
	Biography            string                              `json:"biography" api:"required"`
	BlueskyUsername      string                              `json:"blueskyUsername" api:"required"`
	Email                string                              `json:"email" api:"required"`
	FirstLogin           bool                                `json:"firstLogin" api:"required"`
	FollowerCount        float64                             `json:"followerCount" api:"required"`
	FollowingCount       float64                             `json:"followingCount" api:"required"`
	FollowingTopicsCount float64                             `json:"followingTopicsCount" api:"required"`
	GitHubUsername       string                              `json:"githubUsername" api:"required"`
	GoogleScholarID      string                              `json:"googleScholarId" api:"required"`
	Institution          string                              `json:"institution" api:"required"`
	LinkedinUsername     string                              `json:"linkedinUsername" api:"required"`
	Location             string                              `json:"location" api:"required"`
	OrcidID              string                              `json:"orcidId" api:"required"`
	PublicEmail          string                              `json:"publicEmail" api:"required"`
	RealName             string                              `json:"realName" api:"required"`
	Reputation           float64                             `json:"reputation" api:"required"`
	// Any of "user", "reviewer", "admin", "bot".
	Role                   UserV3GetUserByUuidResponseRole            `json:"role" api:"required"`
	SemanticScholar        UserV3GetUserByUuidResponseSemanticScholar `json:"semanticScholar" api:"required"`
	Username               string                                     `json:"username" api:"required"`
	Verified               bool                                       `json:"verified" api:"required"`
	WeeklyReputation       float64                                    `json:"weeklyReputation" api:"required"`
	XUsername              string                                     `json:"xUsername" api:"required"`
	Featured               []UserV3GetUserByUuidResponseFeatured      `json:"featured"`
	FollowingOrganizations []string                                   `json:"followingOrganizations" format:"uuid"`
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
func (r UserV3GetUserByUuidResponse) RawJSON() string { return r.JSON.raw }
func (r *UserV3GetUserByUuidResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserV3GetUserByUuidResponseAvatar struct {
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
func (r UserV3GetUserByUuidResponseAvatar) RawJSON() string { return r.JSON.raw }
func (r *UserV3GetUserByUuidResponseAvatar) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserV3GetUserByUuidResponseRole string

const (
	UserV3GetUserByUuidResponseRoleUser     UserV3GetUserByUuidResponseRole = "user"
	UserV3GetUserByUuidResponseRoleReviewer UserV3GetUserByUuidResponseRole = "reviewer"
	UserV3GetUserByUuidResponseRoleAdmin    UserV3GetUserByUuidResponseRole = "admin"
	UserV3GetUserByUuidResponseRoleBot      UserV3GetUserByUuidResponseRole = "bot"
)

type UserV3GetUserByUuidResponseSemanticScholar struct {
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
func (r UserV3GetUserByUuidResponseSemanticScholar) RawJSON() string { return r.JSON.raw }
func (r *UserV3GetUserByUuidResponseSemanticScholar) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserV3GetUserByUuidResponseFeatured struct {
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
func (r UserV3GetUserByUuidResponseFeatured) RawJSON() string { return r.JSON.raw }
func (r *UserV3GetUserByUuidResponseFeatured) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserV3GetViewedHistoryResponse struct {
	ID       string `json:"id" api:"required" format:"uuid"`
	Abstract string `json:"abstract" api:"required"`
	ImageURL string `json:"imageUrl" api:"required"`
	// A versionless universal paper ID (e.g. 1706.03762)
	PaperID          string   `json:"paperId" api:"required"`
	PublicationDate  string   `json:"publicationDate" api:"required"`
	PublicTotalVotes float64  `json:"publicTotalVotes" api:"required"`
	Title            string   `json:"title" api:"required"`
	Topics           []string `json:"topics" api:"required"`
	ViewedAt         string   `json:"viewedAt" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		Abstract         respjson.Field
		ImageURL         respjson.Field
		PaperID          respjson.Field
		PublicationDate  respjson.Field
		PublicTotalVotes respjson.Field
		Title            respjson.Field
		Topics           respjson.Field
		ViewedAt         respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserV3GetViewedHistoryResponse) RawJSON() string { return r.JSON.raw }
func (r *UserV3GetViewedHistoryResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserV3ProcessNotificationEmailResponse struct {
	Message string `json:"message" api:"required"`
	Success bool   `json:"success" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserV3ProcessNotificationEmailResponse) RawJSON() string { return r.JSON.raw }
func (r *UserV3ProcessNotificationEmailResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserV3SearchResponse struct {
	Users []UserV3SearchResponseUser `json:"users" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Users       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserV3SearchResponse) RawJSON() string { return r.JSON.raw }
func (r *UserV3SearchResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserV3SearchResponseUser struct {
	ID               string                           `json:"id" api:"required" format:"uuid"`
	Avatar           []UserV3SearchResponseUserAvatar `json:"avatar" api:"required"`
	GoogleScholarID  string                           `json:"googleScholarId" api:"required"`
	Institution      string                           `json:"institution" api:"required"`
	RealName         string                           `json:"realName" api:"required"`
	Reputation       float64                          `json:"reputation" api:"required"`
	Username         string                           `json:"username" api:"required"`
	WeeklyReputation float64                          `json:"weeklyReputation" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		Avatar           respjson.Field
		GoogleScholarID  respjson.Field
		Institution      respjson.Field
		RealName         respjson.Field
		Reputation       respjson.Field
		Username         respjson.Field
		WeeklyReputation respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserV3SearchResponseUser) RawJSON() string { return r.JSON.raw }
func (r *UserV3SearchResponseUser) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserV3SearchResponseUserAvatar struct {
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
func (r UserV3SearchResponseUserAvatar) RawJSON() string { return r.JSON.raw }
func (r *UserV3SearchResponseUserAvatar) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserV3ToggleFollowUserResponse struct {
	Following bool `json:"following" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Following   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserV3ToggleFollowUserResponse) RawJSON() string { return r.JSON.raw }
func (r *UserV3ToggleFollowUserResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserV3UpdatePreferencesResponse struct {
	Banners []UserV3UpdatePreferencesResponseBanner `json:"banners" api:"required"`
	Base    UserV3UpdatePreferencesResponseBase     `json:"base" api:"required"`
	Email   UserV3UpdatePreferencesResponseEmail    `json:"email" api:"required"`
	Folders []UserV3UpdatePreferencesResponseFolder `json:"folders" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Banners     respjson.Field
		Base        respjson.Field
		Email       respjson.Field
		Folders     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserV3UpdatePreferencesResponse) RawJSON() string { return r.JSON.raw }
func (r *UserV3UpdatePreferencesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserV3UpdatePreferencesResponseBanner struct {
	ID   string `json:"id" api:"required" format:"uuid"`
	Link string `json:"link" api:"required"`
	Name string `json:"name" api:"required"`
	// Any of "success", "info", "warning", "error".
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Link        respjson.Field
		Name        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserV3UpdatePreferencesResponseBanner) RawJSON() string { return r.JSON.raw }
func (r *UserV3UpdatePreferencesResponseBanner) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserV3UpdatePreferencesResponseBase struct {
	AssistantCustomStyles   []UserV3UpdatePreferencesResponseBaseAssistantCustomStyle `json:"assistantCustomStyles" api:"required"`
	AssistantStyleSelection string                                                    `json:"assistantStyleSelection" api:"required"`
	// Any of "assistant", "notes", "similar".
	DefaultPrivatePaperSidebarTab string `json:"defaultPrivatePaperSidebarTab" api:"required"`
	// Any of "comments", "assistant", "similar", "notes", "social".
	DefaultPublicPaperSidebarTab string `json:"defaultPublicPaperSidebarTab" api:"required"`
	// Any of "Hot", "Comments", "Views", "Likes", "GitHub", "Twitter (X)",
	// "Recommended".
	FeedSort                string `json:"feedSort" api:"required"`
	IsDarkModeEnabled       bool   `json:"isDarkModeEnabled" api:"required"`
	IsDebugModeEnabled      bool   `json:"isDebugModeEnabled" api:"required"`
	IsMembersSidebarVisible bool   `json:"isMembersSidebarVisible" api:"required"`
	// Any of "am", "ar", "az", "bg", "bn", "ca", "cs", "da", "de", "el", "en", "es",
	// "et", "fa", "fi", "fr", "gu", "ha", "he", "hi", "hr", "hu", "id", "it", "ja",
	// "ka", "kn", "ko", "lt", "lv", "ml", "mr", "ms", "my", "ne", "nl", "no", "pa",
	// "pl", "pt", "ro", "ru", "si", "sk", "sl", "sr", "sv", "sw", "ta", "te", "th",
	// "tl", "tr", "uk", "ur", "uz", "vi", "yo", "zh".
	PreferredLanguage  string  `json:"preferredLanguage" api:"required"`
	PreferredLlmModel  string  `json:"preferredLlmModel" api:"required"`
	ReadingModeEnabled bool    `json:"readingModeEnabled" api:"required"`
	ShowModelThinking  bool    `json:"showModelThinking" api:"required"`
	ToolingPaneWidth   float64 `json:"toolingPaneWidth" api:"required"`
	// Any of "off", "full".
	WebSearch string `json:"webSearch" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AssistantCustomStyles         respjson.Field
		AssistantStyleSelection       respjson.Field
		DefaultPrivatePaperSidebarTab respjson.Field
		DefaultPublicPaperSidebarTab  respjson.Field
		FeedSort                      respjson.Field
		IsDarkModeEnabled             respjson.Field
		IsDebugModeEnabled            respjson.Field
		IsMembersSidebarVisible       respjson.Field
		PreferredLanguage             respjson.Field
		PreferredLlmModel             respjson.Field
		ReadingModeEnabled            respjson.Field
		ShowModelThinking             respjson.Field
		ToolingPaneWidth              respjson.Field
		WebSearch                     respjson.Field
		ExtraFields                   map[string]respjson.Field
		raw                           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserV3UpdatePreferencesResponseBase) RawJSON() string { return r.JSON.raw }
func (r *UserV3UpdatePreferencesResponseBase) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserV3UpdatePreferencesResponseBaseAssistantCustomStyle struct {
	ID           string `json:"id" api:"required" format:"uuid"`
	Instructions string `json:"instructions" api:"required"`
	Name         string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Instructions respjson.Field
		Name         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserV3UpdatePreferencesResponseBaseAssistantCustomStyle) RawJSON() string { return r.JSON.raw }
func (r *UserV3UpdatePreferencesResponseBaseAssistantCustomStyle) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserV3UpdatePreferencesResponseEmail struct {
	Bounced             bool `json:"bounced" api:"required"`
	DirectNotifications bool `json:"directNotifications" api:"required"`
	RelevantActivity    bool `json:"relevantActivity" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Bounced             respjson.Field
		DirectNotifications respjson.Field
		RelevantActivity    respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserV3UpdatePreferencesResponseEmail) RawJSON() string { return r.JSON.raw }
func (r *UserV3UpdatePreferencesResponseEmail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserV3UpdatePreferencesResponseFolder struct {
	FolderID string `json:"folderId" api:"required" format:"uuid"`
	Opened   bool   `json:"opened" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FolderID    respjson.Field
		Opened      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserV3UpdatePreferencesResponseFolder) RawJSON() string { return r.JSON.raw }
func (r *UserV3UpdatePreferencesResponseFolder) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserV3UpdateProfileResponse struct {
	ID                   string                              `json:"id" api:"required" format:"uuid"`
	Avatar               []UserV3UpdateProfileResponseAvatar `json:"avatar" api:"required"`
	Biography            string                              `json:"biography" api:"required"`
	BlueskyUsername      string                              `json:"blueskyUsername" api:"required"`
	Email                string                              `json:"email" api:"required"`
	FirstLogin           bool                                `json:"firstLogin" api:"required"`
	FollowerCount        float64                             `json:"followerCount" api:"required"`
	FollowingCount       float64                             `json:"followingCount" api:"required"`
	FollowingTopicsCount float64                             `json:"followingTopicsCount" api:"required"`
	GitHubUsername       string                              `json:"githubUsername" api:"required"`
	GoogleScholarID      string                              `json:"googleScholarId" api:"required"`
	Institution          string                              `json:"institution" api:"required"`
	LinkedinUsername     string                              `json:"linkedinUsername" api:"required"`
	Location             string                              `json:"location" api:"required"`
	OrcidID              string                              `json:"orcidId" api:"required"`
	PublicEmail          string                              `json:"publicEmail" api:"required"`
	RealName             string                              `json:"realName" api:"required"`
	Reputation           float64                             `json:"reputation" api:"required"`
	// Any of "user", "reviewer", "admin", "bot".
	Role                   UserV3UpdateProfileResponseRole            `json:"role" api:"required"`
	SemanticScholar        UserV3UpdateProfileResponseSemanticScholar `json:"semanticScholar" api:"required"`
	Username               string                                     `json:"username" api:"required"`
	Verified               bool                                       `json:"verified" api:"required"`
	WeeklyReputation       float64                                    `json:"weeklyReputation" api:"required"`
	XUsername              string                                     `json:"xUsername" api:"required"`
	Featured               []UserV3UpdateProfileResponseFeatured      `json:"featured"`
	FollowingOrganizations []string                                   `json:"followingOrganizations" format:"uuid"`
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
func (r UserV3UpdateProfileResponse) RawJSON() string { return r.JSON.raw }
func (r *UserV3UpdateProfileResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserV3UpdateProfileResponseAvatar struct {
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
func (r UserV3UpdateProfileResponseAvatar) RawJSON() string { return r.JSON.raw }
func (r *UserV3UpdateProfileResponseAvatar) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserV3UpdateProfileResponseRole string

const (
	UserV3UpdateProfileResponseRoleUser     UserV3UpdateProfileResponseRole = "user"
	UserV3UpdateProfileResponseRoleReviewer UserV3UpdateProfileResponseRole = "reviewer"
	UserV3UpdateProfileResponseRoleAdmin    UserV3UpdateProfileResponseRole = "admin"
	UserV3UpdateProfileResponseRoleBot      UserV3UpdateProfileResponseRole = "bot"
)

type UserV3UpdateProfileResponseSemanticScholar struct {
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
func (r UserV3UpdateProfileResponseSemanticScholar) RawJSON() string { return r.JSON.raw }
func (r *UserV3UpdateProfileResponseSemanticScholar) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserV3UpdateProfileResponseFeatured struct {
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
func (r UserV3UpdateProfileResponseFeatured) RawJSON() string { return r.JSON.raw }
func (r *UserV3UpdateProfileResponseFeatured) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserV3UploadAvatarResponse struct {
	Data []UserV3UploadAvatarResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserV3UploadAvatarResponse) RawJSON() string { return r.JSON.raw }
func (r *UserV3UploadAvatarResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserV3UploadAvatarResponseData struct {
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
func (r UserV3UploadAvatarResponseData) RawJSON() string { return r.JSON.raw }
func (r *UserV3UploadAvatarResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserV3AutocompleteProfileParams struct {
	UserID string `json:"userId" api:"required" format:"uuid"`
	paramObj
}

func (r UserV3AutocompleteProfileParams) MarshalJSON() (data []byte, err error) {
	type shadow UserV3AutocompleteProfileParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UserV3AutocompleteProfileParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserV3GetActivityParams struct {
	// Any of "date", "liked".
	Sort UserV3GetActivityParamsSort `query:"sort,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [UserV3GetActivityParams]'s query parameters as
// `url.Values`.
func (r UserV3GetActivityParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type UserV3GetActivityParamsSort string

const (
	UserV3GetActivityParamsSortDate  UserV3GetActivityParamsSort = "date"
	UserV3GetActivityParamsSortLiked UserV3GetActivityParamsSort = "liked"
)

type UserV3GetClaimedPapersParams struct {
	// Any of "date", "liked", "citations".
	Sort UserV3GetClaimedPapersParamsSort `query:"sort,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [UserV3GetClaimedPapersParams]'s query parameters as
// `url.Values`.
func (r UserV3GetClaimedPapersParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type UserV3GetClaimedPapersParamsSort string

const (
	UserV3GetClaimedPapersParamsSortDate      UserV3GetClaimedPapersParamsSort = "date"
	UserV3GetClaimedPapersParamsSortLiked     UserV3GetClaimedPapersParamsSort = "liked"
	UserV3GetClaimedPapersParamsSortCitations UserV3GetClaimedPapersParamsSort = "citations"
)

type UserV3GetViewedHistoryParams struct {
	Offset param.Opt[string] `query:"offset,omitzero" json:"-"`
	Search param.Opt[string] `query:"search,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [UserV3GetViewedHistoryParams]'s query parameters as
// `url.Values`.
func (r UserV3GetViewedHistoryParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type UserV3SearchParams struct {
	Q     string            `query:"q" api:"required" json:"-"`
	Limit param.Opt[string] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [UserV3SearchParams]'s query parameters as `url.Values`.
func (r UserV3SearchParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type UserV3UpdatePreferencesParams struct {
	Banners []UserV3UpdatePreferencesParamsBanner `json:"banners,omitzero"`
	Base    UserV3UpdatePreferencesParamsBase     `json:"base,omitzero"`
	paramObj
}

func (r UserV3UpdatePreferencesParams) MarshalJSON() (data []byte, err error) {
	type shadow UserV3UpdatePreferencesParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UserV3UpdatePreferencesParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Link, Name, Type are required.
type UserV3UpdatePreferencesParamsBanner struct {
	Link param.Opt[string] `json:"link,omitzero" api:"required"`
	Name string            `json:"name" api:"required"`
	// Any of "success", "info", "warning", "error".
	Type string `json:"type,omitzero" api:"required"`
	paramObj
}

func (r UserV3UpdatePreferencesParamsBanner) MarshalJSON() (data []byte, err error) {
	type shadow UserV3UpdatePreferencesParamsBanner
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UserV3UpdatePreferencesParamsBanner) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[UserV3UpdatePreferencesParamsBanner](
		"type", "success", "info", "warning", "error",
	)
}

type UserV3UpdatePreferencesParamsBase struct {
	PreferredLlmModel       param.Opt[string]  `json:"preferredLlmModel,omitzero"`
	ToolingPaneWidth        param.Opt[float64] `json:"toolingPaneWidth,omitzero"`
	IsDarkModeEnabled       param.Opt[bool]    `json:"isDarkModeEnabled,omitzero"`
	IsDebugModeEnabled      param.Opt[bool]    `json:"isDebugModeEnabled,omitzero"`
	IsMembersSidebarVisible param.Opt[bool]    `json:"isMembersSidebarVisible,omitzero"`
	ReadingModeEnabled      param.Opt[bool]    `json:"readingModeEnabled,omitzero"`
	ShowModelThinking       param.Opt[bool]    `json:"showModelThinking,omitzero"`
	// Any of "assistant", "notes", "similar".
	DefaultPrivatePaperSidebarTab string `json:"defaultPrivatePaperSidebarTab,omitzero"`
	// Any of "comments", "assistant", "similar", "notes", "social".
	DefaultPublicPaperSidebarTab string `json:"defaultPublicPaperSidebarTab,omitzero"`
	// Any of "am", "ar", "az", "bg", "bn", "ca", "cs", "da", "de", "el", "en", "es",
	// "et", "fa", "fi", "fr", "gu", "ha", "he", "hi", "hr", "hu", "id", "it", "ja",
	// "ka", "kn", "ko", "lt", "lv", "ml", "mr", "ms", "my", "ne", "nl", "no", "pa",
	// "pl", "pt", "ro", "ru", "si", "sk", "sl", "sr", "sv", "sw", "ta", "te", "th",
	// "tl", "tr", "uk", "ur", "uz", "vi", "yo", "zh".
	PreferredLanguage       string                                                  `json:"preferredLanguage,omitzero"`
	AssistantCustomStyles   []UserV3UpdatePreferencesParamsBaseAssistantCustomStyle `json:"assistantCustomStyles,omitzero"`
	AssistantStyleSelection string                                                  `json:"assistantStyleSelection,omitzero"`
	// Any of "Hot", "Comments", "Views", "Likes", "GitHub", "Twitter (X)",
	// "Recommended".
	FeedSort string `json:"feedSort,omitzero"`
	// Any of "off", "full".
	WebSearch string `json:"webSearch,omitzero"`
	paramObj
}

func (r UserV3UpdatePreferencesParamsBase) MarshalJSON() (data []byte, err error) {
	type shadow UserV3UpdatePreferencesParamsBase
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UserV3UpdatePreferencesParamsBase) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[UserV3UpdatePreferencesParamsBase](
		"defaultPrivatePaperSidebarTab", "assistant", "notes", "similar",
	)
	apijson.RegisterFieldValidator[UserV3UpdatePreferencesParamsBase](
		"defaultPublicPaperSidebarTab", "comments", "assistant", "similar", "notes", "social",
	)
	apijson.RegisterFieldValidator[UserV3UpdatePreferencesParamsBase](
		"feedSort", "Hot", "Comments", "Views", "Likes", "GitHub", "Twitter (X)", "Recommended",
	)
	apijson.RegisterFieldValidator[UserV3UpdatePreferencesParamsBase](
		"preferredLanguage", "am", "ar", "az", "bg", "bn", "ca", "cs", "da", "de", "el", "en", "es", "et", "fa", "fi", "fr", "gu", "ha", "he", "hi", "hr", "hu", "id", "it", "ja", "ka", "kn", "ko", "lt", "lv", "ml", "mr", "ms", "my", "ne", "nl", "no", "pa", "pl", "pt", "ro", "ru", "si", "sk", "sl", "sr", "sv", "sw", "ta", "te", "th", "tl", "tr", "uk", "ur", "uz", "vi", "yo", "zh",
	)
	apijson.RegisterFieldValidator[UserV3UpdatePreferencesParamsBase](
		"webSearch", "off", "full",
	)
}

// The properties ID, Instructions, Name are required.
type UserV3UpdatePreferencesParamsBaseAssistantCustomStyle struct {
	ID           string `json:"id" api:"required" format:"uuid"`
	Instructions string `json:"instructions" api:"required"`
	Name         string `json:"name" api:"required"`
	paramObj
}

func (r UserV3UpdatePreferencesParamsBaseAssistantCustomStyle) MarshalJSON() (data []byte, err error) {
	type shadow UserV3UpdatePreferencesParamsBaseAssistantCustomStyle
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UserV3UpdatePreferencesParamsBaseAssistantCustomStyle) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserV3UpdateProfileParams struct {
	Biography        param.Opt[string] `json:"biography,omitzero"`
	BlueskyUsername  param.Opt[string] `json:"blueskyUsername,omitzero"`
	GitHubUsername   param.Opt[string] `json:"githubUsername,omitzero"`
	Institution      param.Opt[string] `json:"institution,omitzero"`
	LinkedinUsername param.Opt[string] `json:"linkedinUsername,omitzero"`
	Location         param.Opt[string] `json:"location,omitzero"`
	PublicEmail      param.Opt[string] `json:"publicEmail,omitzero" format:"email"`
	XUsername        param.Opt[string] `json:"xUsername,omitzero"`
	RealName         param.Opt[string] `json:"realName,omitzero"`
	Username         param.Opt[string] `json:"username,omitzero"`
	paramObj
}

func (r UserV3UpdateProfileParams) MarshalJSON() (data []byte, err error) {
	type shadow UserV3UpdateProfileParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UserV3UpdateProfileParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
