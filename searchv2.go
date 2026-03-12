// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package alphaxivcat

import (
	"github.com/AlphaxivCat/alphaxiv_cat-go/option"
)

// SearchV2Service contains methods and other services that help with interacting
// with the alphaxiv_cat API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSearchV2Service] method instead.
type SearchV2Service struct {
	options []option.RequestOption
	Paper   SearchV2PaperService
}

// NewSearchV2Service generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSearchV2Service(opts ...option.RequestOption) (r SearchV2Service) {
	r = SearchV2Service{}
	r.options = opts
	r.Paper = NewSearchV2PaperService(opts...)
	return
}
