// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package alphaxivcat

import (
	"github.com/stainless-sdks/alphaxiv_cat-go/option"
)

// RetrieveService contains methods and other services that help with interacting
// with the alphaxiv_cat API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRetrieveService] method instead.
type RetrieveService struct {
	options []option.RequestOption
	V1      RetrieveV1Service
}

// NewRetrieveService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRetrieveService(opts ...option.RequestOption) (r RetrieveService) {
	r = RetrieveService{}
	r.options = opts
	r.V1 = NewRetrieveV1Service(opts...)
	return
}
