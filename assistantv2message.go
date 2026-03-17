// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package alphaxivcat

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/AlphaxivCat/alphaxiv_cat-go/internal/apijson"
	"github.com/AlphaxivCat/alphaxiv_cat-go/internal/requestconfig"
	"github.com/AlphaxivCat/alphaxiv_cat-go/option"
	"github.com/AlphaxivCat/alphaxiv_cat-go/packages/param"
	"github.com/AlphaxivCat/alphaxiv_cat-go/packages/respjson"
)

// AssistantV2MessageService contains methods and other services that help with
// interacting with the alphaxiv_cat API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAssistantV2MessageService] method instead.
type AssistantV2MessageService struct {
	options []option.RequestOption
}

// NewAssistantV2MessageService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAssistantV2MessageService(opts ...option.RequestOption) (r AssistantV2MessageService) {
	r = AssistantV2MessageService{}
	r.options = opts
	return
}

// Get all messages for an llm chat
//
// Source file:
// `api-server/src/controllers/assistant/v2/get-chat-messages.controller.ts`
func (r *AssistantV2MessageService) List(ctx context.Context, llmChat string, opts ...option.RequestOption) (res *[]AssistantV2MessageListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if llmChat == "" {
		err = errors.New("missing required llmChat parameter")
		return nil, err
	}
	path := fmt.Sprintf("assistant/v2/%s/messages", llmChat)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Select an llm chat message by id
//
// Source file:
// `api-server/src/controllers/assistant/v2/select-message.controller.ts`
func (r *AssistantV2MessageService) Select(ctx context.Context, message string, body AssistantV2MessageSelectParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if body.LlmChat == "" {
		err = errors.New("missing required llmChat parameter")
		return err
	}
	if message == "" {
		err = errors.New("missing required message parameter")
		return err
	}
	path := fmt.Sprintf("assistant/v2/%s/messages/%s/select", body.LlmChat, message)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return err
}

// Set or update feedback for a message
//
// Source file:
// `api-server/src/controllers/assistant/v2/set-message-feedback.controller.ts`
func (r *AssistantV2MessageService) SetFeedback(ctx context.Context, messageID string, body AssistantV2MessageSetFeedbackParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if messageID == "" {
		err = errors.New("missing required messageId parameter")
		return err
	}
	path := fmt.Sprintf("assistant/v2/messages/%s/feedback", messageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return err
}

type AssistantV2MessageListResponse struct {
	ID               string `json:"id" api:"required" format:"uuid"`
	Content          string `json:"content" api:"required"`
	FeedbackCategory string `json:"feedbackCategory" api:"required"`
	FeedbackDetails  string `json:"feedbackDetails" api:"required"`
	// Any of "upvote", "downvote".
	FeedbackType    AssistantV2MessageListResponseFeedbackType `json:"feedbackType" api:"required"`
	Kind            string                                     `json:"kind" api:"required"`
	Model           string                                     `json:"model" api:"required"`
	ParentMessageID string                                     `json:"parentMessageId" api:"required" format:"uuid"`
	SelectedAt      string                                     `json:"selectedAt" api:"required"`
	ToolUseID       string                                     `json:"toolUseId" api:"required"`
	// Any of "tool_use", "tool_result_text", "tool_result_file", "input_file",
	// "input_text", "output_reasoning", "output_text".
	Type  AssistantV2MessageListResponseType `json:"type" api:"required"`
	Trace string                             `json:"trace" api:"nullable" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		Content          respjson.Field
		FeedbackCategory respjson.Field
		FeedbackDetails  respjson.Field
		FeedbackType     respjson.Field
		Kind             respjson.Field
		Model            respjson.Field
		ParentMessageID  respjson.Field
		SelectedAt       respjson.Field
		ToolUseID        respjson.Field
		Type             respjson.Field
		Trace            respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AssistantV2MessageListResponse) RawJSON() string { return r.JSON.raw }
func (r *AssistantV2MessageListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AssistantV2MessageListResponseFeedbackType string

const (
	AssistantV2MessageListResponseFeedbackTypeUpvote   AssistantV2MessageListResponseFeedbackType = "upvote"
	AssistantV2MessageListResponseFeedbackTypeDownvote AssistantV2MessageListResponseFeedbackType = "downvote"
)

type AssistantV2MessageListResponseType string

const (
	AssistantV2MessageListResponseTypeToolUse         AssistantV2MessageListResponseType = "tool_use"
	AssistantV2MessageListResponseTypeToolResultText  AssistantV2MessageListResponseType = "tool_result_text"
	AssistantV2MessageListResponseTypeToolResultFile  AssistantV2MessageListResponseType = "tool_result_file"
	AssistantV2MessageListResponseTypeInputFile       AssistantV2MessageListResponseType = "input_file"
	AssistantV2MessageListResponseTypeInputText       AssistantV2MessageListResponseType = "input_text"
	AssistantV2MessageListResponseTypeOutputReasoning AssistantV2MessageListResponseType = "output_reasoning"
	AssistantV2MessageListResponseTypeOutputText      AssistantV2MessageListResponseType = "output_text"
)

type AssistantV2MessageSelectParams struct {
	LlmChat string `path:"llmChat" api:"required" format:"uuid" json:"-"`
	paramObj
}

type AssistantV2MessageSetFeedbackParams struct {
	Feedback AssistantV2MessageSetFeedbackParamsFeedback `json:"feedback,omitzero" api:"required"`
	paramObj
}

func (r AssistantV2MessageSetFeedbackParams) MarshalJSON() (data []byte, err error) {
	type shadow AssistantV2MessageSetFeedbackParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AssistantV2MessageSetFeedbackParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Type is required.
type AssistantV2MessageSetFeedbackParamsFeedback struct {
	// Any of "upvote", "downvote".
	Type     string            `json:"type,omitzero" api:"required"`
	Category param.Opt[string] `json:"category,omitzero"`
	Details  param.Opt[string] `json:"details,omitzero"`
	paramObj
}

func (r AssistantV2MessageSetFeedbackParamsFeedback) MarshalJSON() (data []byte, err error) {
	type shadow AssistantV2MessageSetFeedbackParamsFeedback
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AssistantV2MessageSetFeedbackParamsFeedback) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AssistantV2MessageSetFeedbackParamsFeedback](
		"type", "upvote", "downvote",
	)
}
