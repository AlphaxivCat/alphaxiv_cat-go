// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package alphaxivcat

import (
	"context"
	"net/http"
	"slices"

	"github.com/AlphaxivCat/alphaxiv_cat-go/internal/apijson"
	"github.com/AlphaxivCat/alphaxiv_cat-go/internal/requestconfig"
	"github.com/AlphaxivCat/alphaxiv_cat-go/option"
	"github.com/AlphaxivCat/alphaxiv_cat-go/packages/respjson"
)

// RetoolV1Service contains methods and other services that help with interacting
// with the alphaxiv_cat API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRetoolV1Service] method instead.
type RetoolV1Service struct {
	options []option.RequestOption
}

// NewRetoolV1Service generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRetoolV1Service(opts ...option.RequestOption) (r RetoolV1Service) {
	r = RetoolV1Service{}
	r.options = opts
	return
}

// Get cumulative user count over time (all time)
//
// Source file:
// `api-server/src/controllers/retool/v1/get-cumulative-users.controller.ts`
func (r *RetoolV1Service) GetCumulativeUsers(ctx context.Context, opts ...option.RequestOption) (res *[]RetoolV1GetCumulativeUsersResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "retool/v1/cumulative-users"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Get daily conversation counts by variant (all time)
//
// Source file:
// `api-server/src/controllers/retool/v1/get-daily-conversations.controller.ts`
func (r *RetoolV1Service) GetDailyConversations(ctx context.Context, opts ...option.RequestOption) (res *[]RetoolV1GetDailyConversationsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "retool/v1/daily-conversations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Get daily new account counts (all time)
//
// Source file:
// `api-server/src/controllers/retool/v1/get-daily-new-accounts.controller.ts`
func (r *RetoolV1Service) GetDailyNewAccounts(ctx context.Context, opts ...option.RequestOption) (res *[]RetoolV1GetDailyNewAccountsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "retool/v1/daily-new-accounts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Get daily user chat message counts by variant (all time)
//
// Source file:
// `api-server/src/controllers/retool/v1/get-daily-user-chat-messages.controller.ts`
func (r *RetoolV1Service) GetDailyUserChatMessages(ctx context.Context, opts ...option.RequestOption) (res *[]RetoolV1GetDailyUserChatMessagesResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "retool/v1/daily-user-chat-messages"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Get total count of all comments
//
// Source file:
// `api-server/src/controllers/retool/v1/get-total-comment-count.controller.ts`
func (r *RetoolV1Service) GetTotalCommentCount(ctx context.Context, opts ...option.RequestOption) (res *float64, err error) {
	opts = slices.Concat(r.options, opts)
	path := "retool/v1/total-comment-count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Get total count of public, non-blocked papers
//
// Source file:
// `api-server/src/controllers/retool/v1/get-total-paper-count.controller.ts`
func (r *RetoolV1Service) GetTotalPaperCount(ctx context.Context, opts ...option.RequestOption) (res *float64, err error) {
	opts = slices.Concat(r.options, opts)
	path := "retool/v1/total-paper-count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Get total count of private notes (comments with tag=personal)
//
// Source file:
// `api-server/src/controllers/retool/v1/get-total-private-notes-count.controller.ts`
func (r *RetoolV1Service) GetTotalPrivateNotesCount(ctx context.Context, opts ...option.RequestOption) (res *float64, err error) {
	opts = slices.Concat(r.options, opts)
	path := "retool/v1/total-private-notes-count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Get total count of non-deleted users
//
// Source file:
// `api-server/src/controllers/retool/v1/get-total-user-count.controller.ts`
func (r *RetoolV1Service) GetTotalUserCount(ctx context.Context, opts ...option.RequestOption) (res *float64, err error) {
	opts = slices.Concat(r.options, opts)
	path := "retool/v1/total-user-count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Get assistant input message counts per user for the last week, sorted by count
// descending
//
// Source file:
// `api-server/src/controllers/retool/v1/get-weekly-message-counts-by-user.controller.ts`
func (r *RetoolV1Service) GetWeeklyMessageCountsByUser(ctx context.Context, opts ...option.RequestOption) (res *[]RetoolV1GetWeeklyMessageCountsByUserResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "retool/v1/weekly-message-counts-by-user"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Get weekly private note counts (all time)
//
// Source file:
// `api-server/src/controllers/retool/v1/get-weekly-private-notes.controller.ts`
func (r *RetoolV1Service) GetWeeklyPrivateNotes(ctx context.Context, opts ...option.RequestOption) (res *[]RetoolV1GetWeeklyPrivateNotesResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "retool/v1/weekly-private-notes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Get weekly public comment counts (all time)
//
// Source file:
// `api-server/src/controllers/retool/v1/get-weekly-public-comments.controller.ts`
func (r *RetoolV1Service) GetWeeklyPublicComments(ctx context.Context, opts ...option.RequestOption) (res *[]RetoolV1GetWeeklyPublicCommentsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "retool/v1/weekly-public-comments"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type RetoolV1GetCumulativeUsersResponse struct {
	Count float64 `json:"count" api:"required"`
	Date  string  `json:"date" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Date        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RetoolV1GetCumulativeUsersResponse) RawJSON() string { return r.JSON.raw }
func (r *RetoolV1GetCumulativeUsersResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RetoolV1GetDailyConversationsResponse struct {
	Count float64 `json:"count" api:"required"`
	Date  string  `json:"date" api:"required"`
	// Any of "homepage", "paper".
	Variant RetoolV1GetDailyConversationsResponseVariant `json:"variant" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Date        respjson.Field
		Variant     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RetoolV1GetDailyConversationsResponse) RawJSON() string { return r.JSON.raw }
func (r *RetoolV1GetDailyConversationsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RetoolV1GetDailyConversationsResponseVariant string

const (
	RetoolV1GetDailyConversationsResponseVariantHomepage RetoolV1GetDailyConversationsResponseVariant = "homepage"
	RetoolV1GetDailyConversationsResponseVariantPaper    RetoolV1GetDailyConversationsResponseVariant = "paper"
)

type RetoolV1GetDailyNewAccountsResponse struct {
	Count float64 `json:"count" api:"required"`
	Date  string  `json:"date" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Date        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RetoolV1GetDailyNewAccountsResponse) RawJSON() string { return r.JSON.raw }
func (r *RetoolV1GetDailyNewAccountsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RetoolV1GetDailyUserChatMessagesResponse struct {
	Count float64 `json:"count" api:"required"`
	Date  string  `json:"date" api:"required"`
	// Any of "homepage", "paper".
	Variant RetoolV1GetDailyUserChatMessagesResponseVariant `json:"variant" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Date        respjson.Field
		Variant     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RetoolV1GetDailyUserChatMessagesResponse) RawJSON() string { return r.JSON.raw }
func (r *RetoolV1GetDailyUserChatMessagesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RetoolV1GetDailyUserChatMessagesResponseVariant string

const (
	RetoolV1GetDailyUserChatMessagesResponseVariantHomepage RetoolV1GetDailyUserChatMessagesResponseVariant = "homepage"
	RetoolV1GetDailyUserChatMessagesResponseVariantPaper    RetoolV1GetDailyUserChatMessagesResponseVariant = "paper"
)

type RetoolV1GetWeeklyMessageCountsByUserResponse struct {
	Count    float64 `json:"count" api:"required"`
	Email    string  `json:"email" api:"required"`
	UserID   string  `json:"userId" api:"required"`
	Username string  `json:"username" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Email       respjson.Field
		UserID      respjson.Field
		Username    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RetoolV1GetWeeklyMessageCountsByUserResponse) RawJSON() string { return r.JSON.raw }
func (r *RetoolV1GetWeeklyMessageCountsByUserResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RetoolV1GetWeeklyPrivateNotesResponse struct {
	Count float64 `json:"count" api:"required"`
	Week  string  `json:"week" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Week        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RetoolV1GetWeeklyPrivateNotesResponse) RawJSON() string { return r.JSON.raw }
func (r *RetoolV1GetWeeklyPrivateNotesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RetoolV1GetWeeklyPublicCommentsResponse struct {
	Count float64 `json:"count" api:"required"`
	Week  string  `json:"week" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Week        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RetoolV1GetWeeklyPublicCommentsResponse) RawJSON() string { return r.JSON.raw }
func (r *RetoolV1GetWeeklyPublicCommentsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
