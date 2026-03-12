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

// ResearchService contains methods and other services that help with interacting
// with the alphaxiv_cat API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewResearchService] method instead.
type ResearchService struct {
	options []option.RequestOption
}

// NewResearchService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewResearchService(opts ...option.RequestOption) (r ResearchService) {
	r = ResearchService{}
	r.options = opts
	return
}

// Create a research project
//
// Source file:
// `api-server/src/controllers/research/v1/create-project.controller.ts`
func (r *ResearchService) NewProject(ctx context.Context, body ResearchNewProjectParams, opts ...option.RequestOption) (res *string, err error) {
	opts = slices.Concat(r.options, opts)
	path := "research/v1"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type ResearchNewProjectParams struct {
	Name     string            `json:"name" api:"required"`
	ParentID param.Opt[string] `json:"parentId,omitzero" format:"uuid"`
	// Prefill the project with a set of research papers.
	Folder ResearchNewProjectParamsFolder `json:"folder,omitzero"`
	// Add these papers to the group on init
	InitialPapers []string `json:"initialPapers,omitzero" format:"uuid"`
	paramObj
}

func (r ResearchNewProjectParams) MarshalJSON() (data []byte, err error) {
	type shadow ResearchNewProjectParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResearchNewProjectParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Prefill the project with a set of research papers.
//
// The properties ID, Delete are required.
type ResearchNewProjectParamsFolder struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// If set true, the folder is deleted after copying it
	Delete bool `json:"delete" api:"required"`
	paramObj
}

func (r ResearchNewProjectParamsFolder) MarshalJSON() (data []byte, err error) {
	type shadow ResearchNewProjectParamsFolder
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResearchNewProjectParamsFolder) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
