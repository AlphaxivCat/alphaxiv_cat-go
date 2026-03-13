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

// EventService contains methods and other services that help with interacting with
// the alphaxiv_cat API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEventService] method instead.
type EventService struct {
	options []option.RequestOption
}

// NewEventService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewEventService(opts ...option.RequestOption) (r EventService) {
	r = EventService{}
	r.options = opts
	return
}

// Retrieve events for the homepage
//
// Source file: `api-server/src/controllers/events/v1/get-events.controller.ts`
func (r *EventService) List(ctx context.Context, opts ...option.RequestOption) (res *[]EventListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "events/v1"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type EventListResponse struct {
	ID           string `json:"id" api:"required" format:"uuid"`
	Date         string `json:"date" api:"required"`
	Link         string `json:"link" api:"required"`
	Organization string `json:"organization" api:"required"`
	Recording    string `json:"recording" api:"required"`
	Speaker      string `json:"speaker" api:"required"`
	Title        string `json:"title" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Date         respjson.Field
		Link         respjson.Field
		Organization respjson.Field
		Recording    respjson.Field
		Speaker      respjson.Field
		Title        respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EventListResponse) RawJSON() string { return r.JSON.raw }
func (r *EventListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
