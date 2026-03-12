// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package alphaxivcat

import (
	"github.com/AlphaxivCat/alphaxiv_cat-go/option"
)

// ArxivService contains methods and other services that help with interacting with
// the alphaxiv_cat API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewArxivService] method instead.
type ArxivService struct {
	options []option.RequestOption
	V1      ArxivV1Service
}

// NewArxivService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewArxivService(opts ...option.RequestOption) (r ArxivService) {
	r = ArxivService{}
	r.options = opts
	r.V1 = NewArxivV1Service(opts...)
	return
}
