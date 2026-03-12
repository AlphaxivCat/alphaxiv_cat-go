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

// FolderV3Service contains methods and other services that help with interacting
// with the alphaxiv_cat API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFolderV3Service] method instead.
type FolderV3Service struct {
	options []option.RequestOption
	Shared  FolderV3SharedService
}

// NewFolderV3Service generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewFolderV3Service(opts ...option.RequestOption) (r FolderV3Service) {
	r = FolderV3Service{}
	r.options = opts
	r.Shared = NewFolderV3SharedService(opts...)
	return
}

// Create a new folder for the current user
//
// Source file: `api-server/src/controllers/folders/v3/create-folder.controller.ts`
func (r *FolderV3Service) New(ctx context.Context, body FolderV3NewParams, opts ...option.RequestOption) (res *FolderV3NewResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "folders/v3"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Get all folders for the current user
//
// Source file: `api-server/src/controllers/folders/v3/get-folders.controller.ts`
func (r *FolderV3Service) List(ctx context.Context, opts ...option.RequestOption) (res *[]FolderV3ListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "folders/v3"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Delete a folder
//
// Source file: `api-server/src/controllers/folders/v3/delete-folder.controller.ts`
func (r *FolderV3Service) Delete(ctx context.Context, folderID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if folderID == "" {
		err = errors.New("missing required folderId parameter")
		return err
	}
	path := fmt.Sprintf("folders/v3/%s", folderID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Add papers to a folder (without removing from other folders)
//
// Source file:
// `api-server/src/controllers/folders/v3/add-papers-to-folder.controller.ts`
func (r *FolderV3Service) AddPapers(ctx context.Context, folderID string, body FolderV3AddPapersParams, opts ...option.RequestOption) (res *FolderV3AddPapersResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if folderID == "" {
		err = errors.New("missing required folderId parameter")
		return nil, err
	}
	path := fmt.Sprintf("folders/v3/%s/add-papers", folderID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Move papers from source folder to destination folder
//
// Source file:
// `api-server/src/controllers/folders/v3/move-papers-to-folder.controller.ts`
func (r *FolderV3Service) MovePapers(ctx context.Context, folderID string, body FolderV3MovePapersParams, opts ...option.RequestOption) (res *FolderV3MovePapersResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if folderID == "" {
		err = errors.New("missing required folderId parameter")
		return nil, err
	}
	path := fmt.Sprintf("folders/v3/%s/move-papers", folderID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Remove papers from a specific folder
//
// Source file: `api-server/src/controllers/folders/v3/remove-papers.controller.ts`
func (r *FolderV3Service) RemovePapers(ctx context.Context, folderID string, body FolderV3RemovePapersParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if folderID == "" {
		err = errors.New("missing required folderId parameter")
		return err
	}
	path := fmt.Sprintf("folders/v3/%s/remove-papers", folderID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return err
}

// Toggle whether a folder and all its descendant folders are shared or not
//
// Source file:
// `api-server/src/controllers/folders/v3/toggle-folder-sharing.controller.ts`
func (r *FolderV3Service) ToggleSharing(ctx context.Context, folderID string, body FolderV3ToggleSharingParams, opts ...option.RequestOption) (res *FolderV3ToggleSharingResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if folderID == "" {
		err = errors.New("missing required folderId parameter")
		return nil, err
	}
	path := fmt.Sprintf("folders/v3/%s/sharing", folderID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Rename a folder
//
// Source file:
// `api-server/src/controllers/folders/v3/update-folder-name.controller.ts`
func (r *FolderV3Service) UpdateName(ctx context.Context, folderID string, body FolderV3UpdateNameParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if folderID == "" {
		err = errors.New("missing required folderId parameter")
		return err
	}
	path := fmt.Sprintf("folders/v3/%s", folderID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, nil, opts...)
	return err
}

// Update a folder's parent (for nesting)
//
// Source file:
// `api-server/src/controllers/folders/v3/update-folder-parent.controller.ts`
func (r *FolderV3Service) UpdateParent(ctx context.Context, folderID string, body FolderV3UpdateParentParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if folderID == "" {
		err = errors.New("missing required folderId parameter")
		return err
	}
	path := fmt.Sprintf("folders/v3/%s/parent", folderID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, nil, opts...)
	return err
}

type FolderV3NewResponse struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FolderV3NewResponse) RawJSON() string { return r.JSON.raw }
func (r *FolderV3NewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FolderV3ListResponse struct {
	ID       string                      `json:"id" api:"required" format:"uuid"`
	Name     string                      `json:"name" api:"required"`
	Order    float64                     `json:"order" api:"required"`
	Papers   []FolderV3ListResponsePaper `json:"papers" api:"required"`
	ParentID string                      `json:"parentId" api:"required" format:"uuid"`
	// Any of "not_shared", "shared".
	SharingStatus FolderV3ListResponseSharingStatus `json:"sharingStatus" api:"required"`
	Type          string                            `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		Name          respjson.Field
		Order         respjson.Field
		Papers        respjson.Field
		ParentID      respjson.Field
		SharingStatus respjson.Field
		Type          respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FolderV3ListResponse) RawJSON() string { return r.JSON.raw }
func (r *FolderV3ListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FolderV3ListResponsePaper struct {
	Abstract        string                                  `json:"abstract" api:"required"`
	AddedAt         string                                  `json:"addedAt" api:"required"`
	Authors         []FolderV3ListResponsePaperAuthor       `json:"authors" api:"required"`
	Citation        string                                  `json:"citation" api:"required"`
	ImageURL        string                                  `json:"imageUrl" api:"required"`
	Organizations   []FolderV3ListResponsePaperOrganization `json:"organizations" api:"required"`
	PaperGroupID    string                                  `json:"paperGroupId" api:"required" format:"uuid"`
	PublicationDate string                                  `json:"publicationDate" api:"required"`
	Title           string                                  `json:"title" api:"required"`
	Topics          []string                                `json:"topics" api:"required"`
	// Any of "private", "community", "public".
	Type             string                                `json:"type" api:"required"`
	UniversalPaperID string                                `json:"universalPaperId" api:"required"`
	UserAuthors      []FolderV3ListResponsePaperUserAuthor `json:"userAuthors" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Abstract         respjson.Field
		AddedAt          respjson.Field
		Authors          respjson.Field
		Citation         respjson.Field
		ImageURL         respjson.Field
		Organizations    respjson.Field
		PaperGroupID     respjson.Field
		PublicationDate  respjson.Field
		Title            respjson.Field
		Topics           respjson.Field
		Type             respjson.Field
		UniversalPaperID respjson.Field
		UserAuthors      respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FolderV3ListResponsePaper) RawJSON() string { return r.JSON.raw }
func (r *FolderV3ListResponsePaper) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FolderV3ListResponsePaperAuthor struct {
	ID       string `json:"id" api:"required" format:"uuid"`
	FullName string `json:"full_name" api:"required"`
	UserID   string `json:"user_id" api:"required" format:"uuid"`
	Username string `json:"username" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		FullName    respjson.Field
		UserID      respjson.Field
		Username    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FolderV3ListResponsePaperAuthor) RawJSON() string { return r.JSON.raw }
func (r *FolderV3ListResponsePaperAuthor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FolderV3ListResponsePaperOrganization struct {
	Image string `json:"image" api:"required"`
	Name  string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Image       respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FolderV3ListResponsePaperOrganization) RawJSON() string { return r.JSON.raw }
func (r *FolderV3ListResponsePaperOrganization) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FolderV3ListResponsePaperUserAuthor struct {
	ID               string                                      `json:"id" api:"required" format:"uuid"`
	Avatar           []FolderV3ListResponsePaperUserAuthorAvatar `json:"avatar" api:"required"`
	BlueskyUsername  string                                      `json:"blueskyUsername" api:"required"`
	GitHubUsername   string                                      `json:"githubUsername" api:"required"`
	GoogleScholarID  string                                      `json:"googleScholarId" api:"required"`
	Institution      string                                      `json:"institution" api:"required"`
	LinkedinUsername string                                      `json:"linkedinUsername" api:"required"`
	OrcidID          string                                      `json:"orcidId" api:"required"`
	PublicEmail      string                                      `json:"publicEmail" api:"required"`
	RealName         string                                      `json:"realName" api:"required"`
	Reputation       float64                                     `json:"reputation" api:"required"`
	// Any of "user", "reviewer", "admin", "bot".
	Role             string  `json:"role" api:"required"`
	Username         string  `json:"username" api:"required"`
	Verified         bool    `json:"verified" api:"required"`
	WeeklyReputation float64 `json:"weeklyReputation" api:"required"`
	XUsername        string  `json:"xUsername" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		Avatar           respjson.Field
		BlueskyUsername  respjson.Field
		GitHubUsername   respjson.Field
		GoogleScholarID  respjson.Field
		Institution      respjson.Field
		LinkedinUsername respjson.Field
		OrcidID          respjson.Field
		PublicEmail      respjson.Field
		RealName         respjson.Field
		Reputation       respjson.Field
		Role             respjson.Field
		Username         respjson.Field
		Verified         respjson.Field
		WeeklyReputation respjson.Field
		XUsername        respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FolderV3ListResponsePaperUserAuthor) RawJSON() string { return r.JSON.raw }
func (r *FolderV3ListResponsePaperUserAuthor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FolderV3ListResponsePaperUserAuthorAvatar struct {
	// Any of "full_size", "thumbnail".
	Type string `json:"type" api:"required"`
	URL  string `json:"url" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FolderV3ListResponsePaperUserAuthorAvatar) RawJSON() string { return r.JSON.raw }
func (r *FolderV3ListResponsePaperUserAuthorAvatar) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FolderV3ListResponseSharingStatus string

const (
	FolderV3ListResponseSharingStatusNotShared FolderV3ListResponseSharingStatus = "not_shared"
	FolderV3ListResponseSharingStatusShared    FolderV3ListResponseSharingStatus = "shared"
)

type FolderV3AddPapersResponse struct {
	AddedCount     float64 `json:"addedCount" api:"required"`
	DuplicateCount float64 `json:"duplicateCount" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AddedCount     respjson.Field
		DuplicateCount respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FolderV3AddPapersResponse) RawJSON() string { return r.JSON.raw }
func (r *FolderV3AddPapersResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FolderV3MovePapersResponse struct {
	DuplicateCount float64 `json:"duplicateCount" api:"required"`
	MovedCount     float64 `json:"movedCount" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DuplicateCount respjson.Field
		MovedCount     respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FolderV3MovePapersResponse) RawJSON() string { return r.JSON.raw }
func (r *FolderV3MovePapersResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FolderV3ToggleSharingResponse struct {
	// Any of "not_shared", "shared".
	SharingStatus FolderV3ToggleSharingResponseSharingStatus `json:"sharingStatus" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SharingStatus respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FolderV3ToggleSharingResponse) RawJSON() string { return r.JSON.raw }
func (r *FolderV3ToggleSharingResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FolderV3ToggleSharingResponseSharingStatus string

const (
	FolderV3ToggleSharingResponseSharingStatusNotShared FolderV3ToggleSharingResponseSharingStatus = "not_shared"
	FolderV3ToggleSharingResponseSharingStatusShared    FolderV3ToggleSharingResponseSharingStatus = "shared"
)

type FolderV3NewParams struct {
	FolderName string `json:"folderName" api:"required"`
	paramObj
}

func (r FolderV3NewParams) MarshalJSON() (data []byte, err error) {
	type shadow FolderV3NewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FolderV3NewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FolderV3AddPapersParams struct {
	PaperGroupIDs []string `json:"paperGroupIds,omitzero" format:"uuid"`
	UniversalIDs  []string `json:"universalIds,omitzero"`
	paramObj
}

func (r FolderV3AddPapersParams) MarshalJSON() (data []byte, err error) {
	type shadow FolderV3AddPapersParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FolderV3AddPapersParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FolderV3MovePapersParams struct {
	PaperGroupIDs []string          `json:"paperGroupIds,omitzero" api:"required" format:"uuid"`
	FromFolderID  param.Opt[string] `json:"fromFolderId,omitzero" format:"uuid"`
	paramObj
}

func (r FolderV3MovePapersParams) MarshalJSON() (data []byte, err error) {
	type shadow FolderV3MovePapersParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FolderV3MovePapersParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FolderV3RemovePapersParams struct {
	PaperGroupIDs []string `json:"paperGroupIds,omitzero" api:"required" format:"uuid"`
	paramObj
}

func (r FolderV3RemovePapersParams) MarshalJSON() (data []byte, err error) {
	type shadow FolderV3RemovePapersParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FolderV3RemovePapersParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FolderV3ToggleSharingParams struct {
	// Any of "not_shared", "shared".
	SharingStatus FolderV3ToggleSharingParamsSharingStatus `json:"sharingStatus,omitzero" api:"required"`
	paramObj
}

func (r FolderV3ToggleSharingParams) MarshalJSON() (data []byte, err error) {
	type shadow FolderV3ToggleSharingParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FolderV3ToggleSharingParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FolderV3ToggleSharingParamsSharingStatus string

const (
	FolderV3ToggleSharingParamsSharingStatusNotShared FolderV3ToggleSharingParamsSharingStatus = "not_shared"
	FolderV3ToggleSharingParamsSharingStatusShared    FolderV3ToggleSharingParamsSharingStatus = "shared"
)

type FolderV3UpdateNameParams struct {
	Name string `json:"name" api:"required"`
	paramObj
}

func (r FolderV3UpdateNameParams) MarshalJSON() (data []byte, err error) {
	type shadow FolderV3UpdateNameParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FolderV3UpdateNameParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FolderV3UpdateParentParams struct {
	ParentID param.Opt[string] `json:"parentId,omitzero" api:"required" format:"uuid"`
	paramObj
}

func (r FolderV3UpdateParentParams) MarshalJSON() (data []byte, err error) {
	type shadow FolderV3UpdateParentParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FolderV3UpdateParentParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
