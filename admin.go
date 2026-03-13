// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package alphaxivcat

import (
	"github.com/AlphaxivCat/alphaxiv_cat-go/option"
)

// AdminService contains methods and other services that help with interacting with
// the alphaxiv_cat API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAdminService] method instead.
type AdminService struct {
	options []option.RequestOption
	V1      AdminV1Service
}

// NewAdminService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAdminService(opts ...option.RequestOption) (r AdminService) {
	r = AdminService{}
	r.options = opts
	r.V1 = NewAdminV1Service(opts...)
	return
}
