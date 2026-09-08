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
	"github.com/AlphaxivCat/alphaxiv_cat-go/packages/respjson"
)

// EmailService contains methods and other services that help with interacting with
// the alphaxiv_cat API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEmailService] method instead.
type EmailService struct {
	options []option.RequestOption
}

// NewEmailService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewEmailService(opts ...option.RequestOption) (r EmailService) {
	r = EmailService{}
	r.options = opts
	return
}

// Receives bounce notifications from AWS SES via SNS
//
// Source file:
// `api-server/file:/app/api-server/src/controllers/v1/emails/capture-bounced-emails.controller.ts`
func (r *EmailService) CaptureBouncedEmails(ctx context.Context, body EmailCaptureBouncedEmailsParams, opts ...option.RequestOption) (res *EmailCaptureBouncedEmailsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/emails/capture-bounced-emails"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Receives bounce notifications from Resend
//
// Source file:
// `api-server/file:/app/api-server/src/controllers/v1/emails/capture-resend-bounced-emails.controller.ts`
func (r *EmailService) CaptureResendBouncedEmail(ctx context.Context, body EmailCaptureResendBouncedEmailParams, opts ...option.RequestOption) (res *EmailCaptureResendBouncedEmailResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/emails/capture-resend-bounced-email"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Process a bounced email and update user preferences
//
// Source file:
// `api-server/file:/app/api-server/src/controllers/v1/emails/process-bounced-email.controller.ts`
//
// Deprecated: deprecated
func (r *EmailService) ProcessBouncedEmail(ctx context.Context, body EmailProcessBouncedEmailParams, opts ...option.RequestOption) (res *EmailProcessBouncedEmailResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/emails/process-bounced-email"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type EmailCaptureBouncedEmailsResponse struct {
	Data EmailCaptureBouncedEmailsResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailCaptureBouncedEmailsResponse) RawJSON() string { return r.JSON.raw }
func (r *EmailCaptureBouncedEmailsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailCaptureBouncedEmailsResponseData struct {
	Message string `json:"message" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailCaptureBouncedEmailsResponseData) RawJSON() string { return r.JSON.raw }
func (r *EmailCaptureBouncedEmailsResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailCaptureResendBouncedEmailResponse struct {
	Data EmailCaptureResendBouncedEmailResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailCaptureResendBouncedEmailResponse) RawJSON() string { return r.JSON.raw }
func (r *EmailCaptureResendBouncedEmailResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailCaptureResendBouncedEmailResponseData struct {
	Message string `json:"message" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailCaptureResendBouncedEmailResponseData) RawJSON() string { return r.JSON.raw }
func (r *EmailCaptureResendBouncedEmailResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailProcessBouncedEmailResponse struct {
	Data EmailProcessBouncedEmailResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailProcessBouncedEmailResponse) RawJSON() string { return r.JSON.raw }
func (r *EmailProcessBouncedEmailResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailProcessBouncedEmailResponseData struct {
	Message string `json:"message" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailProcessBouncedEmailResponseData) RawJSON() string { return r.JSON.raw }
func (r *EmailProcessBouncedEmailResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailCaptureBouncedEmailsParams struct {
	// Stringified JSON message containing bounce/complaint data
	Message string `json:"Message" api:"required"`
	// SNS notification type
	Type string `json:"Type" api:"required"`
	paramObj
}

func (r EmailCaptureBouncedEmailsParams) MarshalJSON() (data []byte, err error) {
	type shadow EmailCaptureBouncedEmailsParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EmailCaptureBouncedEmailsParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailCaptureResendBouncedEmailParams struct {
	// Event data containing bounced emails
	Data EmailCaptureResendBouncedEmailParamsData `json:"data,omitzero" api:"required"`
	// Event type from Resend
	Type string `json:"type" api:"required"`
	paramObj
}

func (r EmailCaptureResendBouncedEmailParams) MarshalJSON() (data []byte, err error) {
	type shadow EmailCaptureResendBouncedEmailParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EmailCaptureResendBouncedEmailParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Event data containing bounced emails
type EmailCaptureResendBouncedEmailParamsData struct {
	// Bounced email addresses
	To []string `json:"to,omitzero"`
	paramObj
}

func (r EmailCaptureResendBouncedEmailParamsData) MarshalJSON() (data []byte, err error) {
	type shadow EmailCaptureResendBouncedEmailParamsData
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EmailCaptureResendBouncedEmailParamsData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailProcessBouncedEmailParams struct {
	Email string `json:"email" api:"required" format:"email"`
	paramObj
}

func (r EmailProcessBouncedEmailParams) MarshalJSON() (data []byte, err error) {
	type shadow EmailProcessBouncedEmailParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EmailProcessBouncedEmailParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
