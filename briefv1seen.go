// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package alphaxivcat

import (
	"github.com/AlphaxivCat/alphaxiv_cat-go/option"
)

// BriefV1SeenService contains methods and other services that help with
// interacting with the alphaxiv_cat API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBriefV1SeenService] method instead.
type BriefV1SeenService struct {
	options []option.RequestOption
}

// NewBriefV1SeenService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBriefV1SeenService(opts ...option.RequestOption) (r BriefV1SeenService) {
	r = BriefV1SeenService{}
	r.options = opts
	return
}
