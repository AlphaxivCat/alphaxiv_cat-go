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
	"github.com/AlphaxivCat/alphaxiv_cat-go/packages/param"
	"github.com/AlphaxivCat/alphaxiv_cat-go/packages/respjson"
)

// UserV3FollowingTopicService contains methods and other services that help with
// interacting with the alphaxiv_cat API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserV3FollowingTopicService] method instead.
type UserV3FollowingTopicService struct {
	options []option.RequestOption
}

// NewUserV3FollowingTopicService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewUserV3FollowingTopicService(opts ...option.RequestOption) (r UserV3FollowingTopicService) {
	r = UserV3FollowingTopicService{}
	r.options = opts
	return
}

// List the topics followed by a user
//
// Source file:
// `api-server/src/controllers/users/v3/get-following-topics.controller.ts`
func (r *UserV3FollowingTopicService) List(ctx context.Context, id string, opts ...option.RequestOption) (res *UserV3FollowingTopicListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("users/v3/%s/following/topics", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Toggle following a topic for the current user
//
// Source file:
// `api-server/src/controllers/users/v3/toggle-following-topic.controller.ts`
func (r *UserV3FollowingTopicService) Toggle(ctx context.Context, id string, body UserV3FollowingTopicToggleParams, opts ...option.RequestOption) (res *UserV3FollowingTopicToggleResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("users/v3/%s/following/topics/toggle", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type UserV3FollowingTopicListResponse struct {
	Topics []UserV3FollowingTopicListResponseTopic `json:"topics" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Topics      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserV3FollowingTopicListResponse) RawJSON() string { return r.JSON.raw }
func (r *UserV3FollowingTopicListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserV3FollowingTopicListResponseTopic struct {
	ID    string `json:"id" api:"required" format:"uuid"`
	Topic string `json:"topic" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Topic       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserV3FollowingTopicListResponseTopic) RawJSON() string { return r.JSON.raw }
func (r *UserV3FollowingTopicListResponseTopic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserV3FollowingTopicToggleResponse struct {
	Following bool   `json:"following" api:"required"`
	Topic     string `json:"topic" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Following   respjson.Field
		Topic       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserV3FollowingTopicToggleResponse) RawJSON() string { return r.JSON.raw }
func (r *UserV3FollowingTopicToggleResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserV3FollowingTopicToggleParams struct {
	Topic string `json:"topic" api:"required"`
	paramObj
}

func (r UserV3FollowingTopicToggleParams) MarshalJSON() (data []byte, err error) {
	type shadow UserV3FollowingTopicToggleParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UserV3FollowingTopicToggleParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
