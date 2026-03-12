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
)

// CommentService contains methods and other services that help with interacting
// with the alphaxiv_cat API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCommentService] method instead.
type CommentService struct {
	options []option.RequestOption
	V2      CommentV2Service
}

// NewCommentService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewCommentService(opts ...option.RequestOption) (r CommentService) {
	r = CommentService{}
	r.options = opts
	r.V2 = NewCommentV2Service(opts...)
	return
}

// Edit a comment by UUID
//
// Source file: `api-server/src/controllers/comments/v1/edit-comment.controller.ts`
func (r *CommentService) Edit(ctx context.Context, comment string, body CommentEditParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if comment == "" {
		err = errors.New("missing required comment parameter")
		return err
	}
	path := fmt.Sprintf("comments/v1/%s", comment)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, nil, opts...)
	return err
}

type CommentEditParams struct {
	HighlightColor param.Opt[string]                `json:"highlightColor,omitzero" api:"required"`
	SelectedText   param.Opt[string]                `json:"selectedText,omitzero" api:"required"`
	Title          param.Opt[string]                `json:"title,omitzero" api:"required"`
	AnchorPosition CommentEditParamsAnchorPosition  `json:"anchorPosition,omitzero" api:"required"`
	FocusPosition  CommentEditParamsFocusPosition   `json:"focusPosition,omitzero" api:"required"`
	HighlightRects []CommentEditParamsHighlightRect `json:"highlightRects,omitzero" api:"required"`
	Body           string                           `json:"body" api:"required"`
	paramObj
}

func (r CommentEditParams) MarshalJSON() (data []byte, err error) {
	type shadow CommentEditParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CommentEditParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Offset, PageIndex, SpanIndex are required.
type CommentEditParamsAnchorPosition struct {
	Offset    float64 `json:"offset" api:"required"`
	PageIndex float64 `json:"pageIndex" api:"required"`
	SpanIndex float64 `json:"spanIndex" api:"required"`
	paramObj
}

func (r CommentEditParamsAnchorPosition) MarshalJSON() (data []byte, err error) {
	type shadow CommentEditParamsAnchorPosition
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CommentEditParamsAnchorPosition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Offset, PageIndex, SpanIndex are required.
type CommentEditParamsFocusPosition struct {
	Offset    float64 `json:"offset" api:"required"`
	PageIndex float64 `json:"pageIndex" api:"required"`
	SpanIndex float64 `json:"spanIndex" api:"required"`
	paramObj
}

func (r CommentEditParamsFocusPosition) MarshalJSON() (data []byte, err error) {
	type shadow CommentEditParamsFocusPosition
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CommentEditParamsFocusPosition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties PageIndex, Rects are required.
type CommentEditParamsHighlightRect struct {
	PageIndex float64                              `json:"pageIndex" api:"required"`
	Rects     []CommentEditParamsHighlightRectRect `json:"rects,omitzero" api:"required"`
	paramObj
}

func (r CommentEditParamsHighlightRect) MarshalJSON() (data []byte, err error) {
	type shadow CommentEditParamsHighlightRect
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CommentEditParamsHighlightRect) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties X1, X2, Y1, Y2 are required.
type CommentEditParamsHighlightRectRect struct {
	X1 float64 `json:"x1" api:"required"`
	X2 float64 `json:"x2" api:"required"`
	Y1 float64 `json:"y1" api:"required"`
	Y2 float64 `json:"y2" api:"required"`
	paramObj
}

func (r CommentEditParamsHighlightRectRect) MarshalJSON() (data []byte, err error) {
	type shadow CommentEditParamsHighlightRectRect
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CommentEditParamsHighlightRectRect) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
