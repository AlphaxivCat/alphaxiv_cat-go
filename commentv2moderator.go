// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package alphaxivcat

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/stainless-sdks/alphaxiv_cat-go/internal/apijson"
	"github.com/stainless-sdks/alphaxiv_cat-go/internal/requestconfig"
	"github.com/stainless-sdks/alphaxiv_cat-go/option"
	"github.com/stainless-sdks/alphaxiv_cat-go/packages/param"
	"github.com/stainless-sdks/alphaxiv_cat-go/packages/respjson"
)

// CommentV2ModeratorService contains methods and other services that help with
// interacting with the alphaxiv_cat API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCommentV2ModeratorService] method instead.
type CommentV2ModeratorService struct {
	options []option.RequestOption
}

// NewCommentV2ModeratorService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCommentV2ModeratorService(opts ...option.RequestOption) (r CommentV2ModeratorService) {
	r = CommentV2ModeratorService{}
	r.options = opts
	return
}

// Send moderator feedback to comment author
//
// Source file:
// `api-server/src/controllers/comments/v2/moderator-send-feedback.controller.ts`
func (r *CommentV2ModeratorService) SendFeedback(ctx context.Context, comment string, body CommentV2ModeratorSendFeedbackParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if comment == "" {
		err = errors.New("missing required comment parameter")
		return err
	}
	path := fmt.Sprintf("comments/v2/%s/moderator/send-feedback", comment)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return err
}

// Toggle one (or more) of a set of comment moderation flags
//
// Source file:
// `api-server/src/controllers/comments/v2/toggle-comment-moderation-flag.controller.ts`
func (r *CommentV2ModeratorService) ToggleFlags(ctx context.Context, comment string, body CommentV2ModeratorToggleFlagsParams, opts ...option.RequestOption) (res *CommentV2ModeratorToggleFlagsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if comment == "" {
		err = errors.New("missing required comment parameter")
		return nil, err
	}
	path := fmt.Sprintf("comments/v2/%s/moderator/toggle-flags", comment)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type CommentV2ModeratorToggleFlagsResponse struct {
	Addressed     bool `json:"addressed" api:"required"`
	Closed        bool `json:"closed" api:"required"`
	FlagAddressed bool `json:"flagAddressed" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Addressed     respjson.Field
		Closed        respjson.Field
		FlagAddressed respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CommentV2ModeratorToggleFlagsResponse) RawJSON() string { return r.JSON.raw }
func (r *CommentV2ModeratorToggleFlagsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CommentV2ModeratorSendFeedbackParams struct {
	Message string `json:"message" api:"required"`
	paramObj
}

func (r CommentV2ModeratorSendFeedbackParams) MarshalJSON() (data []byte, err error) {
	type shadow CommentV2ModeratorSendFeedbackParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CommentV2ModeratorSendFeedbackParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CommentV2ModeratorToggleFlagsParams struct {
	Addressed     param.Opt[bool] `json:"addressed,omitzero"`
	Closed        param.Opt[bool] `json:"closed,omitzero"`
	FlagAddressed param.Opt[bool] `json:"flagAddressed,omitzero"`
	paramObj
}

func (r CommentV2ModeratorToggleFlagsParams) MarshalJSON() (data []byte, err error) {
	type shadow CommentV2ModeratorToggleFlagsParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CommentV2ModeratorToggleFlagsParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
