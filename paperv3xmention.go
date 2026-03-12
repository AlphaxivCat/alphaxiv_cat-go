// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package alphaxivcat

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/stainless-sdks/alphaxiv_cat-go/internal/apijson"
	"github.com/stainless-sdks/alphaxiv_cat-go/internal/requestconfig"
	"github.com/stainless-sdks/alphaxiv_cat-go/option"
	"github.com/stainless-sdks/alphaxiv_cat-go/packages/param"
	"github.com/stainless-sdks/alphaxiv_cat-go/packages/respjson"
)

// PaperV3XMentionService contains methods and other services that help with
// interacting with the alphaxiv_cat API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPaperV3XMentionService] method instead.
type PaperV3XMentionService struct {
	options []option.RequestOption
}

// NewPaperV3XMentionService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPaperV3XMentionService(opts ...option.RequestOption) (r PaperV3XMentionService) {
	r = PaperV3XMentionService{}
	r.options = opts
	return
}

// Search for X (Twitter) mentions of a paper, filter for quality, and save to
// database
//
// Source file:
// `api-server/src/controllers/papers/v3/get-paper-x-mentions.controller.ts`
//
// Deprecated: deprecated
func (r *PaperV3XMentionService) Update(ctx context.Context, paperGroupID string, body PaperV3XMentionUpdateParams, opts ...option.RequestOption) (res *PaperV3XMentionUpdateResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if paperGroupID == "" {
		err = errors.New("missing required paperGroupId parameter")
		return nil, err
	}
	path := fmt.Sprintf("papers/v3/x-mentions/%s", paperGroupID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Delete all X (Twitter) mentions for a paper
//
// Source file:
// `api-server/src/controllers/papers/v3/delete-paper-x-mentions.controller.ts`
func (r *PaperV3XMentionService) Delete(ctx context.Context, paperGroupID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if paperGroupID == "" {
		err = errors.New("missing required paperGroupId parameter")
		return err
	}
	path := fmt.Sprintf("papers/v3/x-mentions/%s", paperGroupID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

type PaperV3XMentionUpdateResponse struct {
	Mentions           []PaperV3XMentionUpdateResponseMention `json:"mentions" api:"required"`
	QualityTweetsCount float64                                `json:"qualityTweetsCount" api:"required"`
	TotalTweetsFound   float64                                `json:"totalTweetsFound" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Mentions           respjson.Field
		QualityTweetsCount respjson.Field
		TotalTweetsFound   respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3XMentionUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *PaperV3XMentionUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3XMentionUpdateResponseMention struct {
	AuthorAvatar   string                                      `json:"authorAvatar" api:"required"`
	AuthorName     string                                      `json:"authorName" api:"required"`
	AuthorUsername string                                      `json:"authorUsername" api:"required"`
	ConversationID string                                      `json:"conversationId" api:"required"`
	CreatedAt      string                                      `json:"createdAt" api:"required"`
	Likes          float64                                     `json:"likes" api:"required"`
	Quality        PaperV3XMentionUpdateResponseMentionQuality `json:"quality" api:"required"`
	Replies        float64                                     `json:"replies" api:"required"`
	Retweets       float64                                     `json:"retweets" api:"required"`
	Text           string                                      `json:"text" api:"required"`
	TweetID        string                                      `json:"tweetId" api:"required"`
	TweetURL       string                                      `json:"tweetUrl" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AuthorAvatar   respjson.Field
		AuthorName     respjson.Field
		AuthorUsername respjson.Field
		ConversationID respjson.Field
		CreatedAt      respjson.Field
		Likes          respjson.Field
		Quality        respjson.Field
		Replies        respjson.Field
		Retweets       respjson.Field
		Text           respjson.Field
		TweetID        respjson.Field
		TweetURL       respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3XMentionUpdateResponseMention) RawJSON() string { return r.JSON.raw }
func (r *PaperV3XMentionUpdateResponseMention) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3XMentionUpdateResponseMentionQuality struct {
	ContainsProfanity    bool   `json:"containsProfanity" api:"required"`
	HasUniquePerspective bool   `json:"hasUniquePerspective" api:"required"`
	IsBot                bool   `json:"isBot" api:"required"`
	IsEnglish            bool   `json:"isEnglish" api:"required"`
	IsGrok               bool   `json:"isGrok" api:"required"`
	IsMeaningful         bool   `json:"isMeaningful" api:"required"`
	IsRelevant           bool   `json:"isRelevant" api:"required"`
	Reason               string `json:"reason" api:"required"`
	ShouldInclude        bool   `json:"shouldInclude" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ContainsProfanity    respjson.Field
		HasUniquePerspective respjson.Field
		IsBot                respjson.Field
		IsEnglish            respjson.Field
		IsGrok               respjson.Field
		IsMeaningful         respjson.Field
		IsRelevant           respjson.Field
		Reason               respjson.Field
		ShouldInclude        respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperV3XMentionUpdateResponseMentionQuality) RawJSON() string { return r.JSON.raw }
func (r *PaperV3XMentionUpdateResponseMentionQuality) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperV3XMentionUpdateParams struct {
	// If true, only logs results without saving to database
	DryRun param.Opt[bool] `json:"dryRun,omitzero"`
	paramObj
}

func (r PaperV3XMentionUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow PaperV3XMentionUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PaperV3XMentionUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
