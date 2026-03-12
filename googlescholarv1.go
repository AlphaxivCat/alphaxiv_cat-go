// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package alphaxivcat

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"slices"

	"github.com/AlphaxivCat/alphaxiv_cat-go/internal/apijson"
	"github.com/AlphaxivCat/alphaxiv_cat-go/internal/requestconfig"
	"github.com/AlphaxivCat/alphaxiv_cat-go/option"
	"github.com/AlphaxivCat/alphaxiv_cat-go/packages/param"
	"github.com/AlphaxivCat/alphaxiv_cat-go/packages/respjson"
)

// GoogleScholarV1Service contains methods and other services that help with
// interacting with the alphaxiv_cat API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewGoogleScholarV1Service] method instead.
type GoogleScholarV1Service struct {
	options []option.RequestOption
}

// NewGoogleScholarV1Service generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewGoogleScholarV1Service(opts ...option.RequestOption) (r GoogleScholarV1Service) {
	r = GoogleScholarV1Service{}
	r.options = opts
	return
}

// Start connecting Google Scholar profile to user account, or replace a pending
// connection with a different profile
//
// Source file:
// `api-server/src/controllers/google-scholar/v1/connect.controller.ts`
func (r *GoogleScholarV1Service) Connect(ctx context.Context, body GoogleScholarV1ConnectParams, opts ...option.RequestOption) (res *GoogleScholarV1ConnectResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "google-scholar/v1"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Remove Google Scholar ID and queued papers from a user
//
// Source file: `api-server/src/controllers/google-scholar/v1/delete.controller.ts`
func (r *GoogleScholarV1Service) DeleteConnection(ctx context.Context, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "google-scholar/v1"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Get a full report of a user's Google Scholar sync (including lists of papers)
//
// Source file: `api-server/src/controllers/google-scholar/v1/report.controller.ts`
func (r *GoogleScholarV1Service) GetReport(ctx context.Context, opts ...option.RequestOption) (res *[]GoogleScholarV1GetReportResponseUnion, err error) {
	opts = slices.Concat(r.options, opts)
	path := "google-scholar/v1/report"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Get status of user's Google Scholar sync
//
// Source file: `api-server/src/controllers/google-scholar/v1/status.controller.ts`
func (r *GoogleScholarV1Service) GetStatus(ctx context.Context, opts ...option.RequestOption) (res *GoogleScholarV1GetStatusResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "google-scholar/v1/status"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Start a new Google Scholar sync for this user
//
// Source file: `api-server/src/controllers/google-scholar/v1/resync.controller.ts`
func (r *GoogleScholarV1Service) Resync(ctx context.Context, mode GoogleScholarV1ResyncParamsMode, opts ...option.RequestOption) (res *GoogleScholarV1ResyncResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := fmt.Sprintf("google-scholar/v1/resync/%v", mode)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Set the local-part of a user's Google Scholar institutional email and send a
// verification email to that address
//
// Source file:
// `api-server/src/controllers/google-scholar/v1/set-email.controller.ts`
func (r *GoogleScholarV1Service) SetEmail(ctx context.Context, body GoogleScholarV1SetEmailParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "google-scholar/v1/email"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return err
}

// Make some progress syncing a user's Google Scholar papers
//
// Source file: `api-server/src/controllers/google-scholar/v1/sync.controller.ts`
func (r *GoogleScholarV1Service) Sync(ctx context.Context, opts ...option.RequestOption) (res *GoogleScholarV1SyncResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "google-scholar/v1/sync"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Verify a user's Google Scholar email by entering the code sent to that email
//
// Source file: `api-server/src/controllers/google-scholar/v1/verify.controller.ts`
func (r *GoogleScholarV1Service) VerifyEmail(ctx context.Context, body GoogleScholarV1VerifyEmailParams, opts ...option.RequestOption) (res *GoogleScholarV1VerifyEmailResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "google-scholar/v1/verify"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type GoogleScholarV1ConnectResponse struct {
	// How many papers have been claimed since the current sync started
	ClaimedPapers float64 `json:"claimedPapers" api:"required"`
	EmailDomain   string  `json:"emailDomain" api:"required"`
	// How many papers had an error during processing
	ErroredPapers   float64 `json:"erroredPapers" api:"required"`
	GoogleScholarID string  `json:"googleScholarId" api:"required"`
	// How many papers were not found on arXiv
	NotFoundPapers float64 `json:"notFoundPapers" api:"required"`
	// How many papers are ready and will be claimed once the user verifies their email
	PendingPapers float64 `json:"pendingPapers" api:"required"`
	// How many papers have not been processed at all
	RemainingPapers   float64 `json:"remainingPapers" api:"required"`
	SyncStartedAtDate string  `json:"syncStartedAtDate" api:"required"`
	Verified          bool    `json:"verified" api:"required"`
	EmailLocalPart    string  `json:"emailLocalPart"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClaimedPapers     respjson.Field
		EmailDomain       respjson.Field
		ErroredPapers     respjson.Field
		GoogleScholarID   respjson.Field
		NotFoundPapers    respjson.Field
		PendingPapers     respjson.Field
		RemainingPapers   respjson.Field
		SyncStartedAtDate respjson.Field
		Verified          respjson.Field
		EmailLocalPart    respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GoogleScholarV1ConnectResponse) RawJSON() string { return r.JSON.raw }
func (r *GoogleScholarV1ConnectResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GoogleScholarV1GetReportResponseUnion contains all possible properties and
// values from [GoogleScholarV1GetReportResponseObject],
// [GoogleScholarV1GetReportResponseObject].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type GoogleScholarV1GetReportResponseUnion struct {
	// This field is from variant [GoogleScholarV1GetReportResponseObject].
	Status string `json:"status"`
	// This field is from variant [GoogleScholarV1GetReportResponseObject].
	Title string `json:"title"`
	// This field is from variant [GoogleScholarV1GetReportResponseObject].
	UniversalID string `json:"universalId"`
	// This field is from variant [GoogleScholarV1GetReportResponseObject].
	GoogleCitationID string `json:"googleCitationId"`
	JSON             struct {
		Status           respjson.Field
		Title            respjson.Field
		UniversalID      respjson.Field
		GoogleCitationID respjson.Field
		raw              string
	} `json:"-"`
}

func (u GoogleScholarV1GetReportResponseUnion) AsGoogleScholarV1GetReportResponseObject() (v GoogleScholarV1GetReportResponseObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u GoogleScholarV1GetReportResponseUnion) AsVariant2() (v GoogleScholarV1GetReportResponseObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u GoogleScholarV1GetReportResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *GoogleScholarV1GetReportResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GoogleScholarV1GetReportResponseObject struct {
	// Any of "claimed".
	Status string `json:"status" api:"required"`
	Title  string `json:"title" api:"required"`
	// A versionless universal paper ID (e.g. 1706.03762)
	UniversalID string `json:"universalId" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Status      respjson.Field
		Title       respjson.Field
		UniversalID respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GoogleScholarV1GetReportResponseObject) RawJSON() string { return r.JSON.raw }
func (r *GoogleScholarV1GetReportResponseObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GoogleScholarV1GetStatusResponse struct {
	// How many papers have been claimed since the current sync started
	ClaimedPapers float64 `json:"claimedPapers" api:"required"`
	EmailDomain   string  `json:"emailDomain" api:"required"`
	// How many papers had an error during processing
	ErroredPapers   float64 `json:"erroredPapers" api:"required"`
	GoogleScholarID string  `json:"googleScholarId" api:"required"`
	// How many papers were not found on arXiv
	NotFoundPapers float64 `json:"notFoundPapers" api:"required"`
	// How many papers are ready and will be claimed once the user verifies their email
	PendingPapers float64 `json:"pendingPapers" api:"required"`
	// How many papers have not been processed at all
	RemainingPapers   float64 `json:"remainingPapers" api:"required"`
	SyncStartedAtDate string  `json:"syncStartedAtDate" api:"required"`
	Verified          bool    `json:"verified" api:"required"`
	EmailLocalPart    string  `json:"emailLocalPart"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClaimedPapers     respjson.Field
		EmailDomain       respjson.Field
		ErroredPapers     respjson.Field
		GoogleScholarID   respjson.Field
		NotFoundPapers    respjson.Field
		PendingPapers     respjson.Field
		RemainingPapers   respjson.Field
		SyncStartedAtDate respjson.Field
		Verified          respjson.Field
		EmailLocalPart    respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GoogleScholarV1GetStatusResponse) RawJSON() string { return r.JSON.raw }
func (r *GoogleScholarV1GetStatusResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GoogleScholarV1ResyncResponse struct {
	// How many papers have been claimed since the current sync started
	ClaimedPapers float64 `json:"claimedPapers" api:"required"`
	EmailDomain   string  `json:"emailDomain" api:"required"`
	// How many papers had an error during processing
	ErroredPapers   float64 `json:"erroredPapers" api:"required"`
	GoogleScholarID string  `json:"googleScholarId" api:"required"`
	// How many papers were not found on arXiv
	NotFoundPapers float64 `json:"notFoundPapers" api:"required"`
	// How many papers are ready and will be claimed once the user verifies their email
	PendingPapers float64 `json:"pendingPapers" api:"required"`
	// How many papers have not been processed at all
	RemainingPapers   float64 `json:"remainingPapers" api:"required"`
	SyncStartedAtDate string  `json:"syncStartedAtDate" api:"required"`
	Verified          bool    `json:"verified" api:"required"`
	EmailLocalPart    string  `json:"emailLocalPart"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClaimedPapers     respjson.Field
		EmailDomain       respjson.Field
		ErroredPapers     respjson.Field
		GoogleScholarID   respjson.Field
		NotFoundPapers    respjson.Field
		PendingPapers     respjson.Field
		RemainingPapers   respjson.Field
		SyncStartedAtDate respjson.Field
		Verified          respjson.Field
		EmailLocalPart    respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GoogleScholarV1ResyncResponse) RawJSON() string { return r.JSON.raw }
func (r *GoogleScholarV1ResyncResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GoogleScholarV1SyncResponse struct {
	// How many papers have been claimed since the current sync started
	ClaimedPapers float64 `json:"claimedPapers" api:"required"`
	EmailDomain   string  `json:"emailDomain" api:"required"`
	// How many papers had an error during processing
	ErroredPapers   float64                                    `json:"erroredPapers" api:"required"`
	GoogleScholarID string                                     `json:"googleScholarId" api:"required"`
	NewPapers       []GoogleScholarV1SyncResponseNewPaperUnion `json:"newPapers" api:"required"`
	// How many papers were not found on arXiv
	NotFoundPapers float64 `json:"notFoundPapers" api:"required"`
	// How many papers are ready and will be claimed once the user verifies their email
	PendingPapers float64 `json:"pendingPapers" api:"required"`
	// How many papers have not been processed at all
	RemainingPapers   float64 `json:"remainingPapers" api:"required"`
	SyncStartedAtDate string  `json:"syncStartedAtDate" api:"required"`
	Verified          bool    `json:"verified" api:"required"`
	EmailLocalPart    string  `json:"emailLocalPart"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClaimedPapers     respjson.Field
		EmailDomain       respjson.Field
		ErroredPapers     respjson.Field
		GoogleScholarID   respjson.Field
		NewPapers         respjson.Field
		NotFoundPapers    respjson.Field
		PendingPapers     respjson.Field
		RemainingPapers   respjson.Field
		SyncStartedAtDate respjson.Field
		Verified          respjson.Field
		EmailLocalPart    respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GoogleScholarV1SyncResponse) RawJSON() string { return r.JSON.raw }
func (r *GoogleScholarV1SyncResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GoogleScholarV1SyncResponseNewPaperUnion contains all possible properties and
// values from [GoogleScholarV1SyncResponseNewPaperObject],
// [GoogleScholarV1SyncResponseNewPaperObject].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type GoogleScholarV1SyncResponseNewPaperUnion struct {
	// This field is from variant [GoogleScholarV1SyncResponseNewPaperObject].
	Status string `json:"status"`
	// This field is from variant [GoogleScholarV1SyncResponseNewPaperObject].
	Title string `json:"title"`
	// This field is from variant [GoogleScholarV1SyncResponseNewPaperObject].
	UniversalID string `json:"universalId"`
	// This field is from variant [GoogleScholarV1SyncResponseNewPaperObject].
	GoogleCitationID string `json:"googleCitationId"`
	JSON             struct {
		Status           respjson.Field
		Title            respjson.Field
		UniversalID      respjson.Field
		GoogleCitationID respjson.Field
		raw              string
	} `json:"-"`
}

func (u GoogleScholarV1SyncResponseNewPaperUnion) AsGoogleScholarV1SyncResponseNewPaperObject() (v GoogleScholarV1SyncResponseNewPaperObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u GoogleScholarV1SyncResponseNewPaperUnion) AsVariant2() (v GoogleScholarV1SyncResponseNewPaperObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u GoogleScholarV1SyncResponseNewPaperUnion) RawJSON() string { return u.JSON.raw }

func (r *GoogleScholarV1SyncResponseNewPaperUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GoogleScholarV1SyncResponseNewPaperObject struct {
	// Any of "claimed".
	Status string `json:"status" api:"required"`
	Title  string `json:"title" api:"required"`
	// A versionless universal paper ID (e.g. 1706.03762)
	UniversalID string `json:"universalId" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Status      respjson.Field
		Title       respjson.Field
		UniversalID respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GoogleScholarV1SyncResponseNewPaperObject) RawJSON() string { return r.JSON.raw }
func (r *GoogleScholarV1SyncResponseNewPaperObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GoogleScholarV1VerifyEmailResponse struct {
	GoogleScholarID string `json:"googleScholarId" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		GoogleScholarID respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GoogleScholarV1VerifyEmailResponse) RawJSON() string { return r.JSON.raw }
func (r *GoogleScholarV1VerifyEmailResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GoogleScholarV1ConnectParams struct {
	// Google Scholar ID
	GoogleScholarID string `json:"googleScholarId" api:"required"`
	paramObj
}

func (r GoogleScholarV1ConnectParams) MarshalJSON() (data []byte, err error) {
	type shadow GoogleScholarV1ConnectParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GoogleScholarV1ConnectParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GoogleScholarV1ResyncParamsMode string

const (
	GoogleScholarV1ResyncParamsModeAll GoogleScholarV1ResyncParamsMode = "all"
	GoogleScholarV1ResyncParamsModeNew GoogleScholarV1ResyncParamsMode = "new"
)

type GoogleScholarV1SetEmailParams struct {
	// Institutional email local-part
	LocalPart string `json:"localPart" api:"required"`
	// Send verification code to the user's primary email instead of the institutional
	// email. Admin only.
	VerifyAccountEmail bool `json:"verifyAccountEmail" api:"required"`
	paramObj
}

func (r GoogleScholarV1SetEmailParams) MarshalJSON() (data []byte, err error) {
	type shadow GoogleScholarV1SetEmailParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GoogleScholarV1SetEmailParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GoogleScholarV1VerifyEmailParams struct {
	Code string `json:"code" api:"required"`
	paramObj
}

func (r GoogleScholarV1VerifyEmailParams) MarshalJSON() (data []byte, err error) {
	type shadow GoogleScholarV1VerifyEmailParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GoogleScholarV1VerifyEmailParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
