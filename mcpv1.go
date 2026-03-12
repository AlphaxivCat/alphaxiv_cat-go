// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package alphaxivcat

import (
	"context"
	"encoding/json"
	"net/http"
	"slices"

	shimjson "github.com/AlphaxivCat/alphaxiv_cat-go/internal/encoding/json"
	"github.com/AlphaxivCat/alphaxiv_cat-go/internal/requestconfig"
	"github.com/AlphaxivCat/alphaxiv_cat-go/option"
	"github.com/AlphaxivCat/alphaxiv_cat-go/packages/ssestream"
)

// McpV1Service contains methods and other services that help with interacting with
// the alphaxiv_cat API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMcpV1Service] method instead.
type McpV1Service struct {
	options []option.RequestOption
}

// NewMcpV1Service generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewMcpV1Service(opts ...option.RequestOption) (r McpV1Service) {
	r = McpV1Service{}
	r.options = opts
	return
}

// Establish SSE connection for server-to-client MCP messages
//
// Source file: `api-server/src/controllers/mcp/v1/get.controller.ts`
func (r *McpV1Service) EstablishConnectionStreaming(ctx context.Context, opts ...option.RequestOption) (stream *ssestream.Stream[McpV1EstablishConnectionResponse]) {
	var (
		raw *http.Response
		err error
	)
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/event-stream")}, opts...)
	path := "mcp/v1"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &raw, opts...)
	return ssestream.NewStream[McpV1EstablishConnectionResponse](ssestream.NewDecoder(raw), err)
}

// Send a message to the MCP server
//
// Source file: `api-server/src/controllers/mcp/v1/post.controller.ts`
func (r *McpV1Service) SendMessage(ctx context.Context, body McpV1SendMessageParams, opts ...option.RequestOption) (res *McpV1SendMessageResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "mcp/v1"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Terminate an MCP session
//
// Source file: `api-server/src/controllers/mcp/v1/delete.controller.ts`
func (r *McpV1Service) TerminateSession(ctx context.Context, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "mcp/v1"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

type McpV1EstablishConnectionResponse = any

type McpV1SendMessageResponse = any

type McpV1SendMessageParams struct {
	Body any
	paramObj
}

func (r McpV1SendMessageParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *McpV1SendMessageParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}
