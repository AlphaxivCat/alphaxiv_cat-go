// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package alphaxivcat

import (
	"github.com/AlphaxivCat/alphaxiv_cat-go/option"
)

// RetoolV1Service contains methods and other services that help with interacting
// with the alphaxiv_cat API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRetoolV1Service] method instead.
type RetoolV1Service struct {
	options []option.RequestOption
}

// NewRetoolV1Service generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRetoolV1Service(opts ...option.RequestOption) (r RetoolV1Service) {
	r = RetoolV1Service{}
	r.options = opts
	return
}
