// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package alphaxivcat

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/AlphaxivCat/alphaxiv_cat-go/internal/apijson"
	"github.com/AlphaxivCat/alphaxiv_cat-go/internal/apiquery"
	"github.com/AlphaxivCat/alphaxiv_cat-go/internal/requestconfig"
	"github.com/AlphaxivCat/alphaxiv_cat-go/option"
	"github.com/AlphaxivCat/alphaxiv_cat-go/packages/param"
	"github.com/AlphaxivCat/alphaxiv_cat-go/packages/respjson"
	"github.com/AlphaxivCat/alphaxiv_cat-go/packages/ssestream"
)

// AssistantV2Service contains methods and other services that help with
// interacting with the alphaxiv_cat API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAssistantV2Service] method instead.
type AssistantV2Service struct {
	options  []option.RequestOption
	Messages AssistantV2MessageService
}

// NewAssistantV2Service generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAssistantV2Service(opts ...option.RequestOption) (r AssistantV2Service) {
	r = AssistantV2Service{}
	r.options = opts
	r.Messages = NewAssistantV2MessageService(opts...)
	return
}

// Send a message to the AI assistant and receive streaming responses
//
// Source file: `api-server/src/controllers/assistant/v2/chat.controller.ts`
func (r *AssistantV2Service) ChatStreaming(ctx context.Context, body AssistantV2ChatParams, opts ...option.RequestOption) (stream *ssestream.Stream[AssistantV2ChatResponse]) {
	var (
		raw *http.Response
		err error
	)
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/event-stream")}, opts...)
	path := "assistant/v2/chat"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &raw, opts...)
	return ssestream.NewStream[AssistantV2ChatResponse](ssestream.NewDecoder(raw), err)
}

// Delete an llm chat by id
//
// Source file: `api-server/src/controllers/assistant/v2/delete-chat.controller.ts`
func (r *AssistantV2Service) DeleteChat(ctx context.Context, llmChat string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if llmChat == "" {
		err = errors.New("missing required llmChat parameter")
		return err
	}
	path := fmt.Sprintf("assistant/v2/%s", llmChat)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Updates properties on an LlmChat. Currently only supports title
//
// Source file: `api-server/src/controllers/assistant/v2/edit-chat.controller.ts`
func (r *AssistantV2Service) EditChat(ctx context.Context, llmChat string, body AssistantV2EditChatParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if llmChat == "" {
		err = errors.New("missing required llmChat parameter")
		return err
	}
	path := fmt.Sprintf("assistant/v2/%s", llmChat)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, nil, opts...)
	return err
}

// Get llm chats for this user, filtered by variant, and optionally by paper
// version
//
// Source file: `api-server/src/controllers/assistant/v2/get-chats.controller.ts`
func (r *AssistantV2Service) GetChats(ctx context.Context, query AssistantV2GetChatsParams, opts ...option.RequestOption) (res *[]AssistantV2GetChatsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "assistant/v2"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Fetch metadata (title and favicon) from a given URL
//
// Source file:
// `api-server/src/controllers/assistant/v2/get-url-metadata.controller.ts`
func (r *AssistantV2Service) GetURLMetadata(ctx context.Context, query AssistantV2GetURLMetadataParams, opts ...option.RequestOption) (res *AssistantV2GetURLMetadataResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "assistant/v2/url-metadata"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type AssistantV2ChatResponse = any

type AssistantV2GetChatsResponse struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// UNIX milliseconds
	NewestMessage float64 `json:"newestMessage" api:"required"`
	Title         string  `json:"title" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		NewestMessage respjson.Field
		Title         respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AssistantV2GetChatsResponse) RawJSON() string { return r.JSON.raw }
func (r *AssistantV2GetChatsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AssistantV2GetURLMetadataResponse struct {
	Favicon string `json:"favicon" api:"required"`
	Title   string `json:"title" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Favicon     respjson.Field
		Title       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AssistantV2GetURLMetadataResponse) RawJSON() string { return r.JSON.raw }
func (r *AssistantV2GetURLMetadataResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AssistantV2ChatParams struct {
	LlmChatID          param.Opt[string]                  `json:"llmChatId,omitzero" api:"required" format:"uuid"`
	PaperVersionID     param.Opt[string]                  `json:"paperVersionId,omitzero" api:"required" format:"uuid"`
	ParentMessageID    param.Opt[string]                  `json:"parentMessageId,omitzero" api:"required" format:"uuid"`
	SelectionPageRange []int64                            `json:"selectionPageRange,omitzero" api:"required"`
	Thinking           AssistantV2ChatParamsThinkingUnion `json:"thinking,omitzero" api:"required"`
	Files              []AssistantV2ChatParamsFile        `json:"files,omitzero" api:"required"`
	Message            string                             `json:"message" api:"required"`
	// Any of "off", "full".
	WebSearch       AssistantV2ChatParamsWebSearch `json:"webSearch,omitzero" api:"required"`
	FolderAddPapers param.Opt[bool]                `json:"folderAddPapers,omitzero"`
	FolderID        param.Opt[string]              `json:"folderId,omitzero" format:"uuid"`
	Signature       param.Opt[string]              `json:"signature,omitzero"`
	// Any of "homepage", "paper", "folder", "landing", "folder-add-papers".
	AssistantVariant AssistantV2ChatParamsAssistantVariant `json:"assistantVariant,omitzero"`
	// Any of "claude-opus-4.5", "claude-opus-4.6", "claude-opus-4.7",
	// "claude-sonnet-4.5", "claude-sonnet-4.6", "gemini-2.5-flash", "gemini-2.5-pro",
	// "gemini-3-flash", "gemini-3.1-pro", "glm-5-turbo", "glm-5.1", "gpt-5",
	// "gpt-5.1", "gpt-5.2", "gpt-5.4", "gpt-5.4-mini", "gpt-5.4-nano", "kimi-k2.5",
	// "kimi-k2.6", "mercury-2", "minimax-m2.5", "minimax-m2.7", "qwen-3.5", "fast",
	// "smart", "pro".
	Model AssistantV2ChatParamsModel `json:"model,omitzero"`
	// Any of "free", "pro".
	Plan AssistantV2ChatParamsPlan `json:"plan,omitzero"`
	paramObj
}

func (r AssistantV2ChatParams) MarshalJSON() (data []byte, err error) {
	type shadow AssistantV2ChatParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AssistantV2ChatParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ContentType, URL are required.
type AssistantV2ChatParamsFile struct {
	ContentType string `json:"contentType" api:"required"`
	URL         string `json:"url" api:"required" format:"uri"`
	paramObj
}

func (r AssistantV2ChatParamsFile) MarshalJSON() (data []byte, err error) {
	type shadow AssistantV2ChatParamsFile
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AssistantV2ChatParamsFile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AssistantV2ChatParamsThinkingUnion struct {
	OfBool   param.Opt[bool]   `json:",omitzero,inline"`
	OfString param.Opt[string] `json:",omitzero,inline"`
	paramUnion
}

func (u AssistantV2ChatParamsThinkingUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfString)
}
func (u *AssistantV2ChatParamsThinkingUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

type AssistantV2ChatParamsWebSearch string

const (
	AssistantV2ChatParamsWebSearchOff  AssistantV2ChatParamsWebSearch = "off"
	AssistantV2ChatParamsWebSearchFull AssistantV2ChatParamsWebSearch = "full"
)

type AssistantV2ChatParamsAssistantVariant string

const (
	AssistantV2ChatParamsAssistantVariantHomepage        AssistantV2ChatParamsAssistantVariant = "homepage"
	AssistantV2ChatParamsAssistantVariantPaper           AssistantV2ChatParamsAssistantVariant = "paper"
	AssistantV2ChatParamsAssistantVariantFolder          AssistantV2ChatParamsAssistantVariant = "folder"
	AssistantV2ChatParamsAssistantVariantLanding         AssistantV2ChatParamsAssistantVariant = "landing"
	AssistantV2ChatParamsAssistantVariantFolderAddPapers AssistantV2ChatParamsAssistantVariant = "folder-add-papers"
)

type AssistantV2ChatParamsModel string

const (
	AssistantV2ChatParamsModelClaudeOpus4_5   AssistantV2ChatParamsModel = "claude-opus-4.5"
	AssistantV2ChatParamsModelClaudeOpus4_6   AssistantV2ChatParamsModel = "claude-opus-4.6"
	AssistantV2ChatParamsModelClaudeOpus4_7   AssistantV2ChatParamsModel = "claude-opus-4.7"
	AssistantV2ChatParamsModelClaudeSonnet4_5 AssistantV2ChatParamsModel = "claude-sonnet-4.5"
	AssistantV2ChatParamsModelClaudeSonnet4_6 AssistantV2ChatParamsModel = "claude-sonnet-4.6"
	AssistantV2ChatParamsModelGemini2_5Flash  AssistantV2ChatParamsModel = "gemini-2.5-flash"
	AssistantV2ChatParamsModelGemini2_5Pro    AssistantV2ChatParamsModel = "gemini-2.5-pro"
	AssistantV2ChatParamsModelGemini3Flash    AssistantV2ChatParamsModel = "gemini-3-flash"
	AssistantV2ChatParamsModelGemini3_1Pro    AssistantV2ChatParamsModel = "gemini-3.1-pro"
	AssistantV2ChatParamsModelGlm5Turbo       AssistantV2ChatParamsModel = "glm-5-turbo"
	AssistantV2ChatParamsModelGlm5_1          AssistantV2ChatParamsModel = "glm-5.1"
	AssistantV2ChatParamsModelGpt5            AssistantV2ChatParamsModel = "gpt-5"
	AssistantV2ChatParamsModelGpt5_1          AssistantV2ChatParamsModel = "gpt-5.1"
	AssistantV2ChatParamsModelGpt5_2          AssistantV2ChatParamsModel = "gpt-5.2"
	AssistantV2ChatParamsModelGpt5_4          AssistantV2ChatParamsModel = "gpt-5.4"
	AssistantV2ChatParamsModelGpt5_4Mini      AssistantV2ChatParamsModel = "gpt-5.4-mini"
	AssistantV2ChatParamsModelGpt5_4Nano      AssistantV2ChatParamsModel = "gpt-5.4-nano"
	AssistantV2ChatParamsModelKimiK2_5        AssistantV2ChatParamsModel = "kimi-k2.5"
	AssistantV2ChatParamsModelKimiK2_6        AssistantV2ChatParamsModel = "kimi-k2.6"
	AssistantV2ChatParamsModelMercury2        AssistantV2ChatParamsModel = "mercury-2"
	AssistantV2ChatParamsModelMinimaxM2_5     AssistantV2ChatParamsModel = "minimax-m2.5"
	AssistantV2ChatParamsModelMinimaxM2_7     AssistantV2ChatParamsModel = "minimax-m2.7"
	AssistantV2ChatParamsModelQwen3_5         AssistantV2ChatParamsModel = "qwen-3.5"
	AssistantV2ChatParamsModelFast            AssistantV2ChatParamsModel = "fast"
	AssistantV2ChatParamsModelSmart           AssistantV2ChatParamsModel = "smart"
	AssistantV2ChatParamsModelPro             AssistantV2ChatParamsModel = "pro"
)

type AssistantV2ChatParamsPlan string

const (
	AssistantV2ChatParamsPlanFree AssistantV2ChatParamsPlan = "free"
	AssistantV2ChatParamsPlanPro  AssistantV2ChatParamsPlan = "pro"
)

type AssistantV2EditChatParams struct {
	Title string `json:"title" api:"required"`
	paramObj
}

func (r AssistantV2EditChatParams) MarshalJSON() (data []byte, err error) {
	type shadow AssistantV2EditChatParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AssistantV2EditChatParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AssistantV2GetChatsParams struct {
	Folder       param.Opt[string] `query:"folder,omitzero" format:"uuid" json:"-"`
	PaperVersion param.Opt[string] `query:"paperVersion,omitzero" format:"uuid" json:"-"`
	// Any of "homepage", "paper", "folder".
	Variant AssistantV2GetChatsParamsVariant `query:"variant,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AssistantV2GetChatsParams]'s query parameters as
// `url.Values`.
func (r AssistantV2GetChatsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AssistantV2GetChatsParamsVariant string

const (
	AssistantV2GetChatsParamsVariantHomepage AssistantV2GetChatsParamsVariant = "homepage"
	AssistantV2GetChatsParamsVariantPaper    AssistantV2GetChatsParamsVariant = "paper"
	AssistantV2GetChatsParamsVariantFolder   AssistantV2GetChatsParamsVariant = "folder"
)

type AssistantV2GetURLMetadataParams struct {
	// The URL to fetch metadata from
	URL string `query:"url" api:"required" format:"uri" json:"-"`
	paramObj
}

// URLQuery serializes [AssistantV2GetURLMetadataParams]'s query parameters as
// `url.Values`.
func (r AssistantV2GetURLMetadataParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
