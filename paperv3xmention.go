// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package alphaxivcat

import (
	"github.com/AlphaxivCat/alphaxiv_cat-go/option"
)

// PaperV3XMentionService contains methods and other services that help with
// interacting with the alphaxiv_cat API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPaperV3XMentionService] method instead.
type PaperV3XMentionService struct {
	options []option.RequestOption
}

// NewPaperV3XMentionService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPaperV3XMentionService(opts ...option.RequestOption) (r PaperV3XMentionService) {
	r = PaperV3XMentionService{}
	r.options = opts
	return
}
