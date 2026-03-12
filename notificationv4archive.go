// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package alphaxivcat

import (
	"context"
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

// NotificationV4ArchiveService contains methods and other services that help with
// interacting with the alphaxiv_cat API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewNotificationV4ArchiveService] method instead.
type NotificationV4ArchiveService struct {
	options []option.RequestOption
}

// NewNotificationV4ArchiveService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewNotificationV4ArchiveService(opts ...option.RequestOption) (r NotificationV4ArchiveService) {
	r = NotificationV4ArchiveService{}
	r.options = opts
	return
}

// Given an array of IDs to mark as read, move them into the notification archive.
// Invalid notifications are ignored. By checking the returned count, invalid
// client-side state can be detected and corrected.
//
// Source file:
// `api-server/src/controllers/notifications/v4/archive-notifications.controller.ts`
func (r *NotificationV4ArchiveService) New(ctx context.Context, body NotificationV4ArchiveNewParams, opts ...option.RequestOption) (res *NotificationV4ArchiveNewResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "notifications/v4/archive"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// The notification archive is paginated, requiring multiple calls to query long
// history.
//
// Source file:
// `api-server/src/controllers/notifications/v4/get-archived-notifications.controller.ts`
func (r *NotificationV4ArchiveService) List(ctx context.Context, query NotificationV4ArchiveListParams, opts ...option.RequestOption) (res *NotificationV4ArchiveListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "notifications/v4/archive"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type NotificationV4ArchiveNewResponse struct {
	Count float64 `json:"count" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationV4ArchiveNewResponse) RawJSON() string { return r.JSON.raw }
func (r *NotificationV4ArchiveNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NotificationV4ArchiveListResponse struct {
	Cursor        float64 `json:"cursor" api:"required"`
	Notifications []any   `json:"notifications" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Cursor        respjson.Field
		Notifications respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationV4ArchiveListResponse) RawJSON() string { return r.JSON.raw }
func (r *NotificationV4ArchiveListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NotificationV4ArchiveNewParams struct {
	IDs []string `json:"ids,omitzero" api:"required" format:"uuid"`
	paramObj
}

func (r NotificationV4ArchiveNewParams) MarshalJSON() (data []byte, err error) {
	type shadow NotificationV4ArchiveNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NotificationV4ArchiveNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NotificationV4ArchiveListParams struct {
	// Decimal UNIX time in milliseconds
	Before param.Opt[string] `query:"before,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [NotificationV4ArchiveListParams]'s query parameters as
// `url.Values`.
func (r NotificationV4ArchiveListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
