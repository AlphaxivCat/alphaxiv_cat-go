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

// NotificationService contains methods and other services that help with
// interacting with the alphaxiv_cat API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewNotificationService] method instead.
type NotificationService struct {
	options []option.RequestOption
	V4      NotificationV4Service
}

// NewNotificationService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewNotificationService(opts ...option.RequestOption) (r NotificationService) {
	r = NotificationService{}
	r.options = opts
	r.V4 = NewNotificationV4Service(opts...)
	return
}

// Queues notification emails for all users with unseen notifications
//
// Source file:
// `api-server/src/controllers/v2/notifications/kickoff-notification-emails.controller.ts`
func (r *NotificationService) SendKickoffNotificationEmails(ctx context.Context, opts ...option.RequestOption) (res *NotificationSendKickoffNotificationEmailsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v2/notifications/kickoff-notification-emails"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

type NotificationSendKickoffNotificationEmailsResponse struct {
	Success bool `json:"success" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationSendKickoffNotificationEmailsResponse) RawJSON() string { return r.JSON.raw }
func (r *NotificationSendKickoffNotificationEmailsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
