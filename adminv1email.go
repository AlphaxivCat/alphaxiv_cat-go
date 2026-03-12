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

// AdminV1EmailService contains methods and other services that help with
// interacting with the alphaxiv_cat API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAdminV1EmailService] method instead.
type AdminV1EmailService struct {
	options []option.RequestOption
}

// NewAdminV1EmailService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAdminV1EmailService(opts ...option.RequestOption) (r AdminV1EmailService) {
	r = AdminV1EmailService{}
	r.options = opts
	return
}

// Queue monthly digest emails to users via Upstash
//
// Source file:
// `api-server/src/controllers/admin/v1/emails/send-monthly-digest.controller.ts`
func (r *AdminV1EmailService) SendMonthlyDigest(ctx context.Context, body AdminV1EmailSendMonthlyDigestParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "admin/v1/emails/send-monthly-digest"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return err
}

// Queue weekly digest emails to users via Upstash
//
// Source file:
// `api-server/src/controllers/admin/v1/emails/send-weekly-digest.controller.ts`
func (r *AdminV1EmailService) SendWeeklyDigest(ctx context.Context, body AdminV1EmailSendWeeklyDigestParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "admin/v1/emails/send-weekly-digest"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return err
}

type AdminV1EmailSendMonthlyDigestParams struct {
	// Filter by user role
	//
	// Any of "admin", "user".
	Role AdminV1EmailSendMonthlyDigestParamsRole `json:"role,omitzero"`
	paramObj
}

func (r AdminV1EmailSendMonthlyDigestParams) MarshalJSON() (data []byte, err error) {
	type shadow AdminV1EmailSendMonthlyDigestParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AdminV1EmailSendMonthlyDigestParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Filter by user role
type AdminV1EmailSendMonthlyDigestParamsRole string

const (
	AdminV1EmailSendMonthlyDigestParamsRoleAdmin AdminV1EmailSendMonthlyDigestParamsRole = "admin"
	AdminV1EmailSendMonthlyDigestParamsRoleUser  AdminV1EmailSendMonthlyDigestParamsRole = "user"
)

type AdminV1EmailSendWeeklyDigestParams struct {
	// Custom intro message
	IntroText param.Opt[string] `json:"introText,omitzero"`
	// Custom email subject
	Subject param.Opt[string] `json:"subject,omitzero"`
	// Custom events to include
	Events []AdminV1EmailSendWeeklyDigestParamsEvent `json:"events,omitzero"`
	// Filter by user role
	//
	// Any of "admin", "user".
	Role AdminV1EmailSendWeeklyDigestParamsRole `json:"role,omitzero"`
	paramObj
}

func (r AdminV1EmailSendWeeklyDigestParams) MarshalJSON() (data []byte, err error) {
	type shadow AdminV1EmailSendWeeklyDigestParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AdminV1EmailSendWeeklyDigestParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Date, Description, Link, Title are required.
type AdminV1EmailSendWeeklyDigestParamsEvent struct {
	Date         string            `json:"date" api:"required"`
	Description  string            `json:"description" api:"required"`
	Link         string            `json:"link" api:"required"`
	Title        string            `json:"title" api:"required"`
	EndTimeRaw   param.Opt[string] `json:"endTimeRaw,omitzero"`
	StartTimeRaw param.Opt[string] `json:"startTimeRaw,omitzero"`
	paramObj
}

func (r AdminV1EmailSendWeeklyDigestParamsEvent) MarshalJSON() (data []byte, err error) {
	type shadow AdminV1EmailSendWeeklyDigestParamsEvent
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AdminV1EmailSendWeeklyDigestParamsEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Filter by user role
type AdminV1EmailSendWeeklyDigestParamsRole string

const (
	AdminV1EmailSendWeeklyDigestParamsRoleAdmin AdminV1EmailSendWeeklyDigestParamsRole = "admin"
	AdminV1EmailSendWeeklyDigestParamsRoleUser  AdminV1EmailSendWeeklyDigestParamsRole = "user"
)
