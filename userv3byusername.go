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

// Retrieve a user's basic information given its username
//
// Source file:
// `api-server/file:/app/api-server/src/controllers/users/v3/get-user-by-username.controller.ts`
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

type UserV3ByUsernameGetUserResponse struct {
	ID                   string                                  `json:"id" api:"required" format:"uuid"`
	Avatar               []UserV3ByUsernameGetUserResponseAvatar `json:"avatar" api:"required"`
	Biography            string                                  `json:"biography" api:"required"`
	BlueskyUsername      string                                  `json:"blueskyUsername" api:"required"`
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
	ResearcherSlug       string                                  `json:"researcherSlug" api:"required"`
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
		ResearcherSlug         respjson.Field
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
