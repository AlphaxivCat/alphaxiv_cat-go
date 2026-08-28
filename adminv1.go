// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package alphaxivcat

import (
	"context"
	"encoding/json"
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

// AdminV1Service contains methods and other services that help with interacting
// with the alphaxiv_cat API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAdminV1Service] method instead.
type AdminV1Service struct {
	options []option.RequestOption
	Emails  AdminV1EmailService
}

// NewAdminV1Service generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAdminV1Service(opts ...option.RequestOption) (r AdminV1Service) {
	r = AdminV1Service{}
	r.options = opts
	r.Emails = NewAdminV1EmailService(opts...)
	return
}

// Get page of comments for moderator feed
//
// Source file:
// `api-server/file:/app/api-server/src/controllers/admin/v1/get-moderator-feed.controller.ts`
func (r *AdminV1Service) GetModeratorFeed(ctx context.Context, query AdminV1GetModeratorFeedParams, opts ...option.RequestOption) (res *[]AdminV1GetModeratorFeedResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "admin/v1/moderator-feed"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Look up a user by email address (admin only)
//
// Source file:
// `api-server/file:/app/api-server/src/controllers/admin/v1/lookup-user-by-email.controller.ts`
func (r *AdminV1Service) LookupUserByEmail(ctx context.Context, query AdminV1LookupUserByEmailParams, opts ...option.RequestOption) (res *AdminV1LookupUserByEmailResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "admin/v1/lookup-user-by-email"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type AdminV1GetModeratorFeedResponse struct {
	ID              string                                     `json:"id" api:"required" format:"uuid"`
	Author          AdminV1GetModeratorFeedResponseAuthorUnion `json:"author" api:"required"`
	Body            string                                     `json:"body" api:"required"`
	Date            string                                     `json:"date" api:"required"`
	FlagCount       float64                                    `json:"flagCount" api:"required"`
	IsAddressed     bool                                       `json:"isAddressed" api:"required"`
	IsClosed        bool                                       `json:"isClosed" api:"required"`
	IsFlagAddressed bool                                       `json:"isFlagAddressed" api:"required"`
	PaperTitle      string                                     `json:"paperTitle" api:"required"`
	SelectedText    string                                     `json:"selectedText" api:"required"`
	Title           string                                     `json:"title" api:"required"`
	UniversalID     string                                     `json:"universalId" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		Author          respjson.Field
		Body            respjson.Field
		Date            respjson.Field
		FlagCount       respjson.Field
		IsAddressed     respjson.Field
		IsClosed        respjson.Field
		IsFlagAddressed respjson.Field
		PaperTitle      respjson.Field
		SelectedText    respjson.Field
		Title           respjson.Field
		UniversalID     respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AdminV1GetModeratorFeedResponse) RawJSON() string { return r.JSON.raw }
func (r *AdminV1GetModeratorFeedResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AdminV1GetModeratorFeedResponseAuthorUnion contains all possible properties and
// values from [AdminV1GetModeratorFeedResponseAuthorObject],
// [AdminV1GetModeratorFeedResponseAuthorObject2].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type AdminV1GetModeratorFeedResponseAuthorUnion struct {
	// This field is a union of [string], [any]
	ID AdminV1GetModeratorFeedResponseAuthorUnionID `json:"id"`
	// This field is a union of [[]AdminV1GetModeratorFeedResponseAuthorObjectAvatar],
	// [[]AdminV1GetModeratorFeedResponseAuthorObject2Avatar]
	Avatar           AdminV1GetModeratorFeedResponseAuthorUnionAvatar `json:"avatar"`
	BlueskyUsername  string                                           `json:"blueskyUsername"`
	GitHubUsername   string                                           `json:"githubUsername"`
	GoogleScholarID  string                                           `json:"googleScholarId"`
	Institution      string                                           `json:"institution"`
	LinkedinUsername string                                           `json:"linkedinUsername"`
	OrcidID          string                                           `json:"orcidId"`
	PublicEmail      string                                           `json:"publicEmail"`
	RealName         string                                           `json:"realName"`
	Reputation       float64                                          `json:"reputation"`
	ResearcherSlug   string                                           `json:"researcherSlug"`
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
		ResearcherSlug   respjson.Field
		Role             respjson.Field
		Username         respjson.Field
		Verified         respjson.Field
		WeeklyReputation respjson.Field
		XUsername        respjson.Field
		raw              string
	} `json:"-"`
}

func (u AdminV1GetModeratorFeedResponseAuthorUnion) AsAdminV1GetModeratorFeedResponseAuthorObject() (v AdminV1GetModeratorFeedResponseAuthorObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AdminV1GetModeratorFeedResponseAuthorUnion) AsAdminV1GetModeratorFeedResponseAuthorObject2() (v AdminV1GetModeratorFeedResponseAuthorObject2) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AdminV1GetModeratorFeedResponseAuthorUnion) RawJSON() string { return u.JSON.raw }

func (r *AdminV1GetModeratorFeedResponseAuthorUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AdminV1GetModeratorFeedResponseAuthorUnionID is an implicit subunion of
// [AdminV1GetModeratorFeedResponseAuthorUnion].
// AdminV1GetModeratorFeedResponseAuthorUnionID provides convenient access to the
// sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [AdminV1GetModeratorFeedResponseAuthorUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfAdminV1GetModeratorFeedResponseAuthorObject2ID]
type AdminV1GetModeratorFeedResponseAuthorUnionID struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfAdminV1GetModeratorFeedResponseAuthorObject2ID any `json:",inline"`
	JSON                                             struct {
		OfString                                         respjson.Field
		OfAdminV1GetModeratorFeedResponseAuthorObject2ID respjson.Field
		raw                                              string
	} `json:"-"`
}

func (r *AdminV1GetModeratorFeedResponseAuthorUnionID) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AdminV1GetModeratorFeedResponseAuthorUnionAvatar is an implicit subunion of
// [AdminV1GetModeratorFeedResponseAuthorUnion].
// AdminV1GetModeratorFeedResponseAuthorUnionAvatar provides convenient access to
// the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [AdminV1GetModeratorFeedResponseAuthorUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfAdminV1GetModeratorFeedResponseAuthorObjectAvatarArray
// OfAdminV1GetModeratorFeedResponseAuthorObject2AvatarArray]
type AdminV1GetModeratorFeedResponseAuthorUnionAvatar struct {
	// This field will be present if the value is a
	// [[]AdminV1GetModeratorFeedResponseAuthorObjectAvatar] instead of an object.
	OfAdminV1GetModeratorFeedResponseAuthorObjectAvatarArray []AdminV1GetModeratorFeedResponseAuthorObjectAvatar `json:",inline"`
	// This field will be present if the value is a
	// [[]AdminV1GetModeratorFeedResponseAuthorObject2Avatar] instead of an object.
	OfAdminV1GetModeratorFeedResponseAuthorObject2AvatarArray []AdminV1GetModeratorFeedResponseAuthorObject2Avatar `json:",inline"`
	JSON                                                      struct {
		OfAdminV1GetModeratorFeedResponseAuthorObjectAvatarArray  respjson.Field
		OfAdminV1GetModeratorFeedResponseAuthorObject2AvatarArray respjson.Field
		raw                                                       string
	} `json:"-"`
}

func (r *AdminV1GetModeratorFeedResponseAuthorUnionAvatar) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AdminV1GetModeratorFeedResponseAuthorObject struct {
	ID               string                                              `json:"id" api:"required" format:"uuid"`
	Avatar           []AdminV1GetModeratorFeedResponseAuthorObjectAvatar `json:"avatar" api:"required"`
	BlueskyUsername  string                                              `json:"blueskyUsername" api:"required"`
	GitHubUsername   string                                              `json:"githubUsername" api:"required"`
	GoogleScholarID  string                                              `json:"googleScholarId" api:"required"`
	Institution      string                                              `json:"institution" api:"required"`
	LinkedinUsername string                                              `json:"linkedinUsername" api:"required"`
	OrcidID          string                                              `json:"orcidId" api:"required"`
	PublicEmail      string                                              `json:"publicEmail" api:"required"`
	RealName         string                                              `json:"realName" api:"required"`
	Reputation       float64                                             `json:"reputation" api:"required"`
	ResearcherSlug   string                                              `json:"researcherSlug" api:"required"`
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
func (r AdminV1GetModeratorFeedResponseAuthorObject) RawJSON() string { return r.JSON.raw }
func (r *AdminV1GetModeratorFeedResponseAuthorObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AdminV1GetModeratorFeedResponseAuthorObjectAvatar struct {
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
func (r AdminV1GetModeratorFeedResponseAuthorObjectAvatar) RawJSON() string { return r.JSON.raw }
func (r *AdminV1GetModeratorFeedResponseAuthorObjectAvatar) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AdminV1GetModeratorFeedResponseAuthorObject2 struct {
	ID               any                                                  `json:"id" api:"required"`
	Avatar           []AdminV1GetModeratorFeedResponseAuthorObject2Avatar `json:"avatar" api:"required"`
	BlueskyUsername  string                                               `json:"blueskyUsername" api:"required"`
	GitHubUsername   string                                               `json:"githubUsername" api:"required"`
	GoogleScholarID  string                                               `json:"googleScholarId" api:"required"`
	Institution      string                                               `json:"institution" api:"required"`
	LinkedinUsername string                                               `json:"linkedinUsername" api:"required"`
	OrcidID          string                                               `json:"orcidId" api:"required"`
	PublicEmail      string                                               `json:"publicEmail" api:"required"`
	RealName         string                                               `json:"realName" api:"required"`
	Reputation       float64                                              `json:"reputation" api:"required"`
	ResearcherSlug   string                                               `json:"researcherSlug" api:"required"`
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
func (r AdminV1GetModeratorFeedResponseAuthorObject2) RawJSON() string { return r.JSON.raw }
func (r *AdminV1GetModeratorFeedResponseAuthorObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AdminV1GetModeratorFeedResponseAuthorObject2Avatar struct {
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
func (r AdminV1GetModeratorFeedResponseAuthorObject2Avatar) RawJSON() string { return r.JSON.raw }
func (r *AdminV1GetModeratorFeedResponseAuthorObject2Avatar) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AdminV1LookupUserByEmailResponse struct {
	ID          string `json:"id" api:"required" format:"uuid"`
	Email       string `json:"email" api:"required"`
	Institution string `json:"institution" api:"required"`
	RealName    string `json:"realName" api:"required"`
	// Any of "user", "reviewer", "admin", "bot".
	Role      AdminV1LookupUserByEmailResponseRole `json:"role" api:"required"`
	UpdatedAt string                               `json:"updatedAt" api:"required"`
	Username  string                               `json:"username" api:"required"`
	Verified  bool                                 `json:"verified" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Email       respjson.Field
		Institution respjson.Field
		RealName    respjson.Field
		Role        respjson.Field
		UpdatedAt   respjson.Field
		Username    respjson.Field
		Verified    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AdminV1LookupUserByEmailResponse) RawJSON() string { return r.JSON.raw }
func (r *AdminV1LookupUserByEmailResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AdminV1LookupUserByEmailResponseRole string

const (
	AdminV1LookupUserByEmailResponseRoleUser     AdminV1LookupUserByEmailResponseRole = "user"
	AdminV1LookupUserByEmailResponseRoleReviewer AdminV1LookupUserByEmailResponseRole = "reviewer"
	AdminV1LookupUserByEmailResponseRoleAdmin    AdminV1LookupUserByEmailResponseRole = "admin"
	AdminV1LookupUserByEmailResponseRoleBot      AdminV1LookupUserByEmailResponseRole = "bot"
)

type AdminV1GetModeratorFeedParams struct {
	Page     param.Opt[string] `query:"page,omitzero" json:"-"`
	PageSize param.Opt[string] `query:"pageSize,omitzero" json:"-"`
	// Any of "all", "flagged", "recent".
	FeedType AdminV1GetModeratorFeedParamsFeedType `query:"feedType,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AdminV1GetModeratorFeedParams]'s query parameters as
// `url.Values`.
func (r AdminV1GetModeratorFeedParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AdminV1GetModeratorFeedParamsFeedType string

const (
	AdminV1GetModeratorFeedParamsFeedTypeAll     AdminV1GetModeratorFeedParamsFeedType = "all"
	AdminV1GetModeratorFeedParamsFeedTypeFlagged AdminV1GetModeratorFeedParamsFeedType = "flagged"
	AdminV1GetModeratorFeedParamsFeedTypeRecent  AdminV1GetModeratorFeedParamsFeedType = "recent"
)

type AdminV1LookupUserByEmailParams struct {
	Email string `query:"email" api:"required" format:"email" json:"-"`
	paramObj
}

// URLQuery serializes [AdminV1LookupUserByEmailParams]'s query parameters as
// `url.Values`.
func (r AdminV1LookupUserByEmailParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
