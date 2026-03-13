// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package alphaxivcat

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/AlphaxivCat/alphaxiv_cat-go/internal/apijson"
	"github.com/AlphaxivCat/alphaxiv_cat-go/internal/requestconfig"
	"github.com/AlphaxivCat/alphaxiv_cat-go/option"
	"github.com/AlphaxivCat/alphaxiv_cat-go/packages/respjson"
)

// UserV3FollowingService contains methods and other services that help with
// interacting with the alphaxiv_cat API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserV3FollowingService] method instead.
type UserV3FollowingService struct {
	options       []option.RequestOption
	Topics        UserV3FollowingTopicService
	Organizations UserV3FollowingOrganizationService
}

// NewUserV3FollowingService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewUserV3FollowingService(opts ...option.RequestOption) (r UserV3FollowingService) {
	r = UserV3FollowingService{}
	r.options = opts
	r.Topics = NewUserV3FollowingTopicService(opts...)
	r.Organizations = NewUserV3FollowingOrganizationService(opts...)
	return
}

// List the users that the given user is following
//
// Source file:
// `api-server/src/controllers/users/v3/get-following-users.controller.ts`
func (r *UserV3FollowingService) List(ctx context.Context, id string, opts ...option.RequestOption) (res *UserV3FollowingListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("users/v3/%s/following", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type UserV3FollowingListResponse struct {
	Users []UserV3FollowingListResponseUser `json:"users" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Users       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserV3FollowingListResponse) RawJSON() string { return r.JSON.raw }
func (r *UserV3FollowingListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserV3FollowingListResponseUser struct {
	ID               string                                  `json:"id" api:"required" format:"uuid"`
	Avatar           []UserV3FollowingListResponseUserAvatar `json:"avatar" api:"required"`
	GoogleScholarID  string                                  `json:"googleScholarId" api:"required"`
	Institution      string                                  `json:"institution" api:"required"`
	RealName         string                                  `json:"realName" api:"required"`
	Reputation       float64                                 `json:"reputation" api:"required"`
	Username         string                                  `json:"username" api:"required"`
	WeeklyReputation float64                                 `json:"weeklyReputation" api:"required"`
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
func (r UserV3FollowingListResponseUser) RawJSON() string { return r.JSON.raw }
func (r *UserV3FollowingListResponseUser) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserV3FollowingListResponseUserAvatar struct {
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
func (r UserV3FollowingListResponseUserAvatar) RawJSON() string { return r.JSON.raw }
func (r *UserV3FollowingListResponseUserAvatar) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
