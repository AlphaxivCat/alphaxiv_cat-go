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
)

// CommentV2Service contains methods and other services that help with interacting
// with the alphaxiv_cat API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCommentV2Service] method instead.
type CommentV2Service struct {
	options   []option.RequestOption
	Moderator CommentV2ModeratorService
}

// NewCommentV2Service generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCommentV2Service(opts ...option.RequestOption) (r CommentV2Service) {
	r = CommentV2Service{}
	r.options = opts
	r.Moderator = NewCommentV2ModeratorService(opts...)
	return
}

// Delete a comment by UUID
//
// Source file:
// `api-server/src/controllers/comments/v2/delete-comment.controller.ts`
func (r *CommentV2Service) Delete(ctx context.Context, comment string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if comment == "" {
		err = errors.New("missing required comment parameter")
		return err
	}
	path := fmt.Sprintf("comments/v2/%s", comment)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Downvote a comment by UUID
//
// Source file:
// `api-server/src/controllers/comments/v2/downvote-comment.controller.ts`
func (r *CommentV2Service) Downvote(ctx context.Context, comment string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if comment == "" {
		err = errors.New("missing required comment parameter")
		return err
	}
	path := fmt.Sprintf("comments/v2/%s/downvote", comment)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return err
}

// Flag a comment for moderator review
//
// Source file: `api-server/src/controllers/comments/v2/flag-comment.controller.ts`
func (r *CommentV2Service) Flag(ctx context.Context, comment string, body CommentV2FlagParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if comment == "" {
		err = errors.New("missing required comment parameter")
		return err
	}
	path := fmt.Sprintf("comments/v2/%s/flag", comment)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return err
}

// Toggle author endorsement of a comment
//
// Source file:
// `api-server/src/controllers/comments/v2/toggle-comment-endorse.controller.ts`
func (r *CommentV2Service) ToggleEndorse(ctx context.Context, comment string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if comment == "" {
		err = errors.New("missing required comment parameter")
		return err
	}
	path := fmt.Sprintf("comments/v2/%s/endorse", comment)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return err
}

// Upvote a comment by UUID
//
// Source file:
// `api-server/src/controllers/comments/v2/upvote-comment.controller.ts`
func (r *CommentV2Service) Upvote(ctx context.Context, comment string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if comment == "" {
		err = errors.New("missing required comment parameter")
		return err
	}
	path := fmt.Sprintf("comments/v2/%s/upvote", comment)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return err
}

type CommentV2FlagParams struct {
	Reason string `json:"reason" api:"required"`
	paramObj
}

func (r CommentV2FlagParams) MarshalJSON() (data []byte, err error) {
	type shadow CommentV2FlagParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CommentV2FlagParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
