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
	"github.com/stainless-sdks/alphaxiv_cat-go/packages/respjson"
)

// FolderV3SharedService contains methods and other services that help with
// interacting with the alphaxiv_cat API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFolderV3SharedService] method instead.
type FolderV3SharedService struct {
	options []option.RequestOption
}

// NewFolderV3SharedService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewFolderV3SharedService(opts ...option.RequestOption) (r FolderV3SharedService) {
	r = FolderV3SharedService{}
	r.options = opts
	return
}

// Get a folder that has been shared publicly, including nested child folders
//
// Source file:
// `api-server/src/controllers/folders/v3/get-shared-folder.controller.ts`
func (r *FolderV3SharedService) Get(ctx context.Context, folderID string, opts ...option.RequestOption) (res *FolderV3SharedGetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if folderID == "" {
		err = errors.New("missing required folderId parameter")
		return nil, err
	}
	path := fmt.Sprintf("folders/v3/shared/%s", folderID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Copy a shared folder and all nested folders to your own library
//
// Source file:
// `api-server/src/controllers/folders/v3/copy-shared-folder.controller.ts`
func (r *FolderV3SharedService) Copy(ctx context.Context, folderID string, opts ...option.RequestOption) (res *FolderV3SharedCopyResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if folderID == "" {
		err = errors.New("missing required folderId parameter")
		return nil, err
	}
	path := fmt.Sprintf("folders/v3/shared/%s/copy", folderID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

type FolderV3SharedGetResponse struct {
	ChildFolders []FolderV3SharedGetResponseChildFolder `json:"childFolders" api:"required"`
	Folder       FolderV3SharedGetResponseFolder        `json:"folder" api:"required"`
	OwnerName    string                                 `json:"ownerName" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChildFolders respjson.Field
		Folder       respjson.Field
		OwnerName    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FolderV3SharedGetResponse) RawJSON() string { return r.JSON.raw }
func (r *FolderV3SharedGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FolderV3SharedGetResponseChildFolder struct {
	ID       string                                      `json:"id" api:"required" format:"uuid"`
	Name     string                                      `json:"name" api:"required"`
	Order    float64                                     `json:"order" api:"required"`
	Papers   []FolderV3SharedGetResponseChildFolderPaper `json:"papers" api:"required"`
	ParentID string                                      `json:"parentId" api:"required" format:"uuid"`
	// Any of "not_shared", "shared".
	SharingStatus string `json:"sharingStatus" api:"required"`
	Type          string `json:"type" api:"required"`
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
func (r FolderV3SharedGetResponseChildFolder) RawJSON() string { return r.JSON.raw }
func (r *FolderV3SharedGetResponseChildFolder) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FolderV3SharedGetResponseChildFolderPaper struct {
	Abstract        string                                                  `json:"abstract" api:"required"`
	AddedAt         string                                                  `json:"addedAt" api:"required"`
	Authors         []FolderV3SharedGetResponseChildFolderPaperAuthor       `json:"authors" api:"required"`
	Citation        string                                                  `json:"citation" api:"required"`
	ImageURL        string                                                  `json:"imageUrl" api:"required"`
	Organizations   []FolderV3SharedGetResponseChildFolderPaperOrganization `json:"organizations" api:"required"`
	PaperGroupID    string                                                  `json:"paperGroupId" api:"required" format:"uuid"`
	PublicationDate string                                                  `json:"publicationDate" api:"required"`
	Title           string                                                  `json:"title" api:"required"`
	Topics          []string                                                `json:"topics" api:"required"`
	// Any of "private", "community", "public".
	Type             string                                                `json:"type" api:"required"`
	UniversalPaperID string                                                `json:"universalPaperId" api:"required"`
	UserAuthors      []FolderV3SharedGetResponseChildFolderPaperUserAuthor `json:"userAuthors" api:"required"`
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
func (r FolderV3SharedGetResponseChildFolderPaper) RawJSON() string { return r.JSON.raw }
func (r *FolderV3SharedGetResponseChildFolderPaper) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FolderV3SharedGetResponseChildFolderPaperAuthor struct {
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
func (r FolderV3SharedGetResponseChildFolderPaperAuthor) RawJSON() string { return r.JSON.raw }
func (r *FolderV3SharedGetResponseChildFolderPaperAuthor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FolderV3SharedGetResponseChildFolderPaperOrganization struct {
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
func (r FolderV3SharedGetResponseChildFolderPaperOrganization) RawJSON() string { return r.JSON.raw }
func (r *FolderV3SharedGetResponseChildFolderPaperOrganization) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FolderV3SharedGetResponseChildFolderPaperUserAuthor struct {
	ID               string                                                      `json:"id" api:"required" format:"uuid"`
	Avatar           []FolderV3SharedGetResponseChildFolderPaperUserAuthorAvatar `json:"avatar" api:"required"`
	BlueskyUsername  string                                                      `json:"blueskyUsername" api:"required"`
	GitHubUsername   string                                                      `json:"githubUsername" api:"required"`
	GoogleScholarID  string                                                      `json:"googleScholarId" api:"required"`
	Institution      string                                                      `json:"institution" api:"required"`
	LinkedinUsername string                                                      `json:"linkedinUsername" api:"required"`
	OrcidID          string                                                      `json:"orcidId" api:"required"`
	PublicEmail      string                                                      `json:"publicEmail" api:"required"`
	RealName         string                                                      `json:"realName" api:"required"`
	Reputation       float64                                                     `json:"reputation" api:"required"`
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
func (r FolderV3SharedGetResponseChildFolderPaperUserAuthor) RawJSON() string { return r.JSON.raw }
func (r *FolderV3SharedGetResponseChildFolderPaperUserAuthor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FolderV3SharedGetResponseChildFolderPaperUserAuthorAvatar struct {
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
func (r FolderV3SharedGetResponseChildFolderPaperUserAuthorAvatar) RawJSON() string {
	return r.JSON.raw
}
func (r *FolderV3SharedGetResponseChildFolderPaperUserAuthorAvatar) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FolderV3SharedGetResponseFolder struct {
	ID       string                                 `json:"id" api:"required" format:"uuid"`
	Name     string                                 `json:"name" api:"required"`
	Order    float64                                `json:"order" api:"required"`
	Papers   []FolderV3SharedGetResponseFolderPaper `json:"papers" api:"required"`
	ParentID string                                 `json:"parentId" api:"required" format:"uuid"`
	// Any of "not_shared", "shared".
	SharingStatus string `json:"sharingStatus" api:"required"`
	Type          string `json:"type" api:"required"`
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
func (r FolderV3SharedGetResponseFolder) RawJSON() string { return r.JSON.raw }
func (r *FolderV3SharedGetResponseFolder) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FolderV3SharedGetResponseFolderPaper struct {
	Abstract        string                                             `json:"abstract" api:"required"`
	AddedAt         string                                             `json:"addedAt" api:"required"`
	Authors         []FolderV3SharedGetResponseFolderPaperAuthor       `json:"authors" api:"required"`
	Citation        string                                             `json:"citation" api:"required"`
	ImageURL        string                                             `json:"imageUrl" api:"required"`
	Organizations   []FolderV3SharedGetResponseFolderPaperOrganization `json:"organizations" api:"required"`
	PaperGroupID    string                                             `json:"paperGroupId" api:"required" format:"uuid"`
	PublicationDate string                                             `json:"publicationDate" api:"required"`
	Title           string                                             `json:"title" api:"required"`
	Topics          []string                                           `json:"topics" api:"required"`
	// Any of "private", "community", "public".
	Type             string                                           `json:"type" api:"required"`
	UniversalPaperID string                                           `json:"universalPaperId" api:"required"`
	UserAuthors      []FolderV3SharedGetResponseFolderPaperUserAuthor `json:"userAuthors" api:"required"`
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
func (r FolderV3SharedGetResponseFolderPaper) RawJSON() string { return r.JSON.raw }
func (r *FolderV3SharedGetResponseFolderPaper) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FolderV3SharedGetResponseFolderPaperAuthor struct {
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
func (r FolderV3SharedGetResponseFolderPaperAuthor) RawJSON() string { return r.JSON.raw }
func (r *FolderV3SharedGetResponseFolderPaperAuthor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FolderV3SharedGetResponseFolderPaperOrganization struct {
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
func (r FolderV3SharedGetResponseFolderPaperOrganization) RawJSON() string { return r.JSON.raw }
func (r *FolderV3SharedGetResponseFolderPaperOrganization) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FolderV3SharedGetResponseFolderPaperUserAuthor struct {
	ID               string                                                 `json:"id" api:"required" format:"uuid"`
	Avatar           []FolderV3SharedGetResponseFolderPaperUserAuthorAvatar `json:"avatar" api:"required"`
	BlueskyUsername  string                                                 `json:"blueskyUsername" api:"required"`
	GitHubUsername   string                                                 `json:"githubUsername" api:"required"`
	GoogleScholarID  string                                                 `json:"googleScholarId" api:"required"`
	Institution      string                                                 `json:"institution" api:"required"`
	LinkedinUsername string                                                 `json:"linkedinUsername" api:"required"`
	OrcidID          string                                                 `json:"orcidId" api:"required"`
	PublicEmail      string                                                 `json:"publicEmail" api:"required"`
	RealName         string                                                 `json:"realName" api:"required"`
	Reputation       float64                                                `json:"reputation" api:"required"`
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
func (r FolderV3SharedGetResponseFolderPaperUserAuthor) RawJSON() string { return r.JSON.raw }
func (r *FolderV3SharedGetResponseFolderPaperUserAuthor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FolderV3SharedGetResponseFolderPaperUserAuthorAvatar struct {
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
func (r FolderV3SharedGetResponseFolderPaperUserAuthorAvatar) RawJSON() string { return r.JSON.raw }
func (r *FolderV3SharedGetResponseFolderPaperUserAuthorAvatar) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FolderV3SharedCopyResponse struct {
	CopiedCount   float64 `json:"copiedCount" api:"required"`
	CopiedFolders float64 `json:"copiedFolders" api:"required"`
	FolderID      string  `json:"folderId" api:"required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CopiedCount   respjson.Field
		CopiedFolders respjson.Field
		FolderID      respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FolderV3SharedCopyResponse) RawJSON() string { return r.JSON.raw }
func (r *FolderV3SharedCopyResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
