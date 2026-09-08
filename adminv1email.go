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

// Queue weekly digest emails to users
//
// Source file:
// `api-server/file:/app/api-server/src/controllers/admin/v1/emails/send-weekly-digest.controller.ts`
func (r *AdminV1EmailService) SendWeeklyDigest(ctx context.Context, body AdminV1EmailSendWeeklyDigestParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "admin/v1/emails/send-weekly-digest"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return err
}

type AdminV1EmailSendWeeklyDigestParams struct {
	// Test mode: page size override, to exercise batching
	TestBatchSize param.Opt[int64] `json:"testBatchSize,omitzero"`
	// Text overrides for copy variant A
	A AdminV1EmailSendWeeklyDigestParamsA `json:"a,omitzero"`
	// Text overrides for copy variant B
	B AdminV1EmailSendWeeklyDigestParamsB `json:"b,omitzero"`
	// Custom events to include, both variants
	Events []AdminV1EmailSendWeeklyDigestParamsEvent `json:"events,omitzero"`
	// Filter by user role
	//
	// Any of "admin", "user".
	Role AdminV1EmailSendWeeklyDigestParamsRole `json:"role,omitzero"`
	// Test mode: only these addresses can receive the digest
	TestEmails []string `json:"testEmails,omitzero"`
	paramObj
}

func (r AdminV1EmailSendWeeklyDigestParams) MarshalJSON() (data []byte, err error) {
	type shadow AdminV1EmailSendWeeklyDigestParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AdminV1EmailSendWeeklyDigestParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Text overrides for copy variant A
type AdminV1EmailSendWeeklyDigestParamsA struct {
	// Custom intro message
	IntroText param.Opt[string] `json:"introText,omitzero"`
	// Custom email subject
	Subject param.Opt[string] `json:"subject,omitzero"`
	paramObj
}

func (r AdminV1EmailSendWeeklyDigestParamsA) MarshalJSON() (data []byte, err error) {
	type shadow AdminV1EmailSendWeeklyDigestParamsA
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AdminV1EmailSendWeeklyDigestParamsA) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Text overrides for copy variant B
type AdminV1EmailSendWeeklyDigestParamsB struct {
	// Custom intro message
	IntroText param.Opt[string] `json:"introText,omitzero"`
	// Custom email subject
	Subject param.Opt[string] `json:"subject,omitzero"`
	paramObj
}

func (r AdminV1EmailSendWeeklyDigestParamsB) MarshalJSON() (data []byte, err error) {
	type shadow AdminV1EmailSendWeeklyDigestParamsB
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AdminV1EmailSendWeeklyDigestParamsB) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Date, Description, Link, Title are required.
type AdminV1EmailSendWeeklyDigestParamsEvent struct {
	Date         string            `json:"date" api:"required"`
	Description  string            `json:"description" api:"required"`
	Link         string            `json:"link" api:"required"`
	Title        string            `json:"title" api:"required"`
	CtaText      param.Opt[string] `json:"ctaText,omitzero"`
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
