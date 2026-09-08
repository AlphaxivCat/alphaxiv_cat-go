// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package alphaxivcat

import (
	"github.com/AlphaxivCat/alphaxiv_cat-go/option"
)

// UserV3SemanticScholarService contains methods and other services that help with
// interacting with the alphaxiv_cat API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserV3SemanticScholarService] method instead.
type UserV3SemanticScholarService struct {
	options []option.RequestOption
}

// NewUserV3SemanticScholarService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewUserV3SemanticScholarService(opts ...option.RequestOption) (r UserV3SemanticScholarService) {
	r = UserV3SemanticScholarService{}
	r.options = opts
	return
}
