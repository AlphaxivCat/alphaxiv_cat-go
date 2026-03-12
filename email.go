// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package alphaxivcat

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/stainless-sdks/alphaxiv_cat-go/internal/apijson"
	"github.com/stainless-sdks/alphaxiv_cat-go/internal/requestconfig"
	"github.com/stainless-sdks/alphaxiv_cat-go/option"
	"github.com/stainless-sdks/alphaxiv_cat-go/packages/param"
	"github.com/stainless-sdks/alphaxiv_cat-go/packages/respjson"
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
// `api-server/src/controllers/v1/emails/capture-bounced-emails.controller.ts`
func (r *EmailService) CaptureBouncedEmails(ctx context.Context, body EmailCaptureBouncedEmailsParams, opts ...option.RequestOption) (res *EmailCaptureBouncedEmailsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/emails/capture-bounced-emails"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Receives bounce notifications from Resend
//
// Source file:
// `api-server/src/controllers/v1/emails/capture-resend-bounced-emails.controller.ts`
func (r *EmailService) CaptureResendBouncedEmail(ctx context.Context, body EmailCaptureResendBouncedEmailParams, opts ...option.RequestOption) (res *EmailCaptureResendBouncedEmailResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/emails/capture-resend-bounced-email"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Kicks off a background job to send comment digest emails to users
//
// Source file:
// `api-server/src/controllers/v1/emails/kickoff-comment-updates.controller.ts`
func (r *EmailService) KickoffCommentUpdate(ctx context.Context, custom EmailKickoffCommentUpdateParamsCustom, body EmailKickoffCommentUpdateParams, opts ...option.RequestOption) (res *EmailKickoffCommentUpdateResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if body.Role == "" {
		err = errors.New("missing required role parameter")
		return nil, err
	}
	if body.Window == "" {
		err = errors.New("missing required window parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/emails/kickoff-comment-update/%s/%s/%v", url.PathEscape(body.Role), url.PathEscape(body.Window), custom)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Kicks off a background job to send general digest emails to users
//
// Source file:
// `api-server/src/controllers/v1/emails/kickoff-general-updates.controller.ts`
func (r *EmailService) KickoffGeneralUpdate(ctx context.Context, role string, opts ...option.RequestOption) (res *EmailKickoffGeneralUpdateResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if role == "" {
		err = errors.New("missing required role parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/emails/kickoff-general-update/%s", url.PathEscape(role))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Process a bounced email and update user preferences
//
// Source file:
// `api-server/src/controllers/v1/emails/process-bounced-email.controller.ts`
//
// Deprecated: deprecated
func (r *EmailService) ProcessBouncedEmail(ctx context.Context, body EmailProcessBouncedEmailParams, opts ...option.RequestOption) (res *EmailProcessBouncedEmailResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/emails/process-bounced-email"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Process comment digest email for a user
//
// Source file:
// `api-server/src/controllers/v1/emails/process-user-comment-update.controller.ts`
func (r *EmailService) ProcessCommentUpdate(ctx context.Context, body EmailProcessCommentUpdateParams, opts ...option.RequestOption) (res *EmailProcessCommentUpdateResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/emails/process-comment-update"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Process general digest email for a user
//
// Source file:
// `api-server/src/controllers/v1/emails/process-user-general-update.controller.ts`
func (r *EmailService) ProcessGeneralUpdate(ctx context.Context, body EmailProcessGeneralUpdateParams, opts ...option.RequestOption) (res *EmailProcessGeneralUpdateResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/emails/process-general-update"
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

type EmailKickoffCommentUpdateResponse struct {
	Data EmailKickoffCommentUpdateResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailKickoffCommentUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *EmailKickoffCommentUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailKickoffCommentUpdateResponseData struct {
	Message string `json:"message" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailKickoffCommentUpdateResponseData) RawJSON() string { return r.JSON.raw }
func (r *EmailKickoffCommentUpdateResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailKickoffGeneralUpdateResponse struct {
	Data EmailKickoffGeneralUpdateResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailKickoffGeneralUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *EmailKickoffGeneralUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailKickoffGeneralUpdateResponseData struct {
	Message string `json:"message" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailKickoffGeneralUpdateResponseData) RawJSON() string { return r.JSON.raw }
func (r *EmailKickoffGeneralUpdateResponseData) UnmarshalJSON(data []byte) error {
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

type EmailProcessCommentUpdateResponse struct {
	Data EmailProcessCommentUpdateResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailProcessCommentUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *EmailProcessCommentUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailProcessCommentUpdateResponseData struct {
	ProcessedUser string `json:"processedUser" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ProcessedUser respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailProcessCommentUpdateResponseData) RawJSON() string { return r.JSON.raw }
func (r *EmailProcessCommentUpdateResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailProcessGeneralUpdateResponse struct {
	Data EmailProcessGeneralUpdateResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailProcessGeneralUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *EmailProcessGeneralUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailProcessGeneralUpdateResponseData struct {
	ProcessedUser string `json:"processedUser" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ProcessedUser respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailProcessGeneralUpdateResponseData) RawJSON() string { return r.JSON.raw }
func (r *EmailProcessGeneralUpdateResponseData) UnmarshalJSON(data []byte) error {
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

type EmailKickoffCommentUpdateParams struct {
	// User role to filter by
	Role string `path:"role" api:"required" json:"-"`
	// Time window in hours
	Window string `path:"window" api:"required" json:"-"`
	paramObj
}

// Whether to use custom digest
type EmailKickoffCommentUpdateParamsCustom string

const (
	EmailKickoffCommentUpdateParamsCustomTrue  EmailKickoffCommentUpdateParamsCustom = "true"
	EmailKickoffCommentUpdateParamsCustomFalse EmailKickoffCommentUpdateParamsCustom = "false"
)

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

type EmailProcessCommentUpdateParams struct {
	UserID        string                                       `json:"userId" api:"required" format:"uuid"`
	CustomContent EmailProcessCommentUpdateParamsCustomContent `json:"customContent,omitzero"`
	paramObj
}

func (r EmailProcessCommentUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow EmailProcessCommentUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EmailProcessCommentUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailProcessCommentUpdateParamsCustomContent struct {
	IntroText param.Opt[string]                                   `json:"introText,omitzero"`
	Subject   param.Opt[string]                                   `json:"subject,omitzero"`
	Events    []EmailProcessCommentUpdateParamsCustomContentEvent `json:"events,omitzero"`
	paramObj
}

func (r EmailProcessCommentUpdateParamsCustomContent) MarshalJSON() (data []byte, err error) {
	type shadow EmailProcessCommentUpdateParamsCustomContent
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EmailProcessCommentUpdateParamsCustomContent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Date, Description, Link, Title are required.
type EmailProcessCommentUpdateParamsCustomContentEvent struct {
	Date         string            `json:"date" api:"required"`
	Description  string            `json:"description" api:"required"`
	Link         string            `json:"link" api:"required"`
	Title        string            `json:"title" api:"required"`
	EndTimeRaw   param.Opt[string] `json:"endTimeRaw,omitzero"`
	StartTimeRaw param.Opt[string] `json:"startTimeRaw,omitzero"`
	paramObj
}

func (r EmailProcessCommentUpdateParamsCustomContentEvent) MarshalJSON() (data []byte, err error) {
	type shadow EmailProcessCommentUpdateParamsCustomContentEvent
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EmailProcessCommentUpdateParamsCustomContentEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailProcessGeneralUpdateParams struct {
	UserID string `json:"userId" api:"required" format:"uuid"`
	paramObj
}

func (r EmailProcessGeneralUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow EmailProcessGeneralUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EmailProcessGeneralUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
