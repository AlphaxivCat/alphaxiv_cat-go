// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package alphaxivcat

import (
	"github.com/stainless-sdks/alphaxiv_cat-go/option"
)

// RetoolService contains methods and other services that help with interacting
// with the alphaxiv_cat API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRetoolService] method instead.
type RetoolService struct {
	options []option.RequestOption
	V1      RetoolV1Service
}

// NewRetoolService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewRetoolService(opts ...option.RequestOption) (r RetoolService) {
	r = RetoolService{}
	r.options = opts
	r.V1 = NewRetoolV1Service(opts...)
	return
}
