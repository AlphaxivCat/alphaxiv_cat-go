// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package alphaxivcat

import (
	"github.com/AlphaxivCat/alphaxiv_cat-go/option"
)

// ArxivV1LabService contains methods and other services that help with interacting
// with the alphaxiv_cat API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewArxivV1LabService] method instead.
type ArxivV1LabService struct {
	options []option.RequestOption
}

// NewArxivV1LabService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewArxivV1LabService(opts ...option.RequestOption) (r ArxivV1LabService) {
	r = ArxivV1LabService{}
	r.options = opts
	return
}
