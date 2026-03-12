// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package alphaxivcat

import (
	"context"
	"net/http"
	"slices"

	"github.com/AlphaxivCat/alphaxiv_cat-go/internal/apijson"
	"github.com/AlphaxivCat/alphaxiv_cat-go/internal/requestconfig"
	"github.com/AlphaxivCat/alphaxiv_cat-go/option"
	"github.com/AlphaxivCat/alphaxiv_cat-go/packages/param"
)

// NotificationV4Service contains methods and other services that help with
// interacting with the alphaxiv_cat API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewNotificationV4Service] method instead.
type NotificationV4Service struct {
	options []option.RequestOption
	Archive NotificationV4ArchiveService
}

// NewNotificationV4Service generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewNotificationV4Service(opts ...option.RequestOption) (r NotificationV4Service) {
	r = NotificationV4Service{}
	r.options = opts
	r.Archive = NewNotificationV4ArchiveService(opts...)
	return
}

// Returns all active notifications.
//
// Source file:
// `api-server/src/controllers/notifications/v4/get-notifications.controller.ts`
func (r *NotificationV4Service) ListNotifications(ctx context.Context, opts ...option.RequestOption) (res *[]NotificationV4ListNotificationsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "notifications/v4"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Subscribes the current user to push notifications
//
// Source file:
// `api-server/src/controllers/notifications/v4/subscribe.controller.ts`
func (r *NotificationV4Service) Subscribe(ctx context.Context, body NotificationV4SubscribeParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "notifications/v4/subscribe"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return err
}

// Unsubscribes the current user to push notifications
//
// Source file:
// `api-server/src/controllers/notifications/v4/unsubscribe.controller.ts`
func (r *NotificationV4Service) Unsubscribe(ctx context.Context, body NotificationV4UnsubscribeParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "notifications/v4/unsubscribe"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return err
}

type NotificationV4ListNotificationsResponse = any

type NotificationV4SubscribeParams struct {
	Subscription NotificationV4SubscribeParamsSubscription `json:"subscription,omitzero" api:"required"`
	paramObj
}

func (r NotificationV4SubscribeParams) MarshalJSON() (data []byte, err error) {
	type shadow NotificationV4SubscribeParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NotificationV4SubscribeParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Endpoint, Keys are required.
type NotificationV4SubscribeParamsSubscription struct {
	Endpoint string                                        `json:"endpoint" api:"required" format:"uri"`
	Keys     NotificationV4SubscribeParamsSubscriptionKeys `json:"keys,omitzero" api:"required"`
	paramObj
}

func (r NotificationV4SubscribeParamsSubscription) MarshalJSON() (data []byte, err error) {
	type shadow NotificationV4SubscribeParamsSubscription
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NotificationV4SubscribeParamsSubscription) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Auth, P256dh are required.
type NotificationV4SubscribeParamsSubscriptionKeys struct {
	Auth   string `json:"auth" api:"required"`
	P256dh string `json:"p256dh" api:"required"`
	paramObj
}

func (r NotificationV4SubscribeParamsSubscriptionKeys) MarshalJSON() (data []byte, err error) {
	type shadow NotificationV4SubscribeParamsSubscriptionKeys
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NotificationV4SubscribeParamsSubscriptionKeys) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NotificationV4UnsubscribeParams struct {
	Endpoint string `json:"endpoint" api:"required" format:"uri"`
	paramObj
}

func (r NotificationV4UnsubscribeParams) MarshalJSON() (data []byte, err error) {
	type shadow NotificationV4UnsubscribeParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NotificationV4UnsubscribeParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
