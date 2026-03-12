// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package alphaxivcat

import (
	"github.com/stainless-sdks/alphaxiv_cat-go/option"
)

// GoogleScholarService contains methods and other services that help with
// interacting with the alphaxiv_cat API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewGoogleScholarService] method instead.
type GoogleScholarService struct {
	options []option.RequestOption
	V1      GoogleScholarV1Service
}

// NewGoogleScholarService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewGoogleScholarService(opts ...option.RequestOption) (r GoogleScholarService) {
	r = GoogleScholarService{}
	r.options = opts
	r.V1 = NewGoogleScholarV1Service(opts...)
	return
}
