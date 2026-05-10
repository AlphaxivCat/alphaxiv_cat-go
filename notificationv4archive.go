// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package alphaxivcat

import (
	"github.com/AlphaxivCat/alphaxiv_cat-go/option"
)

// NotificationV4ArchiveService contains methods and other services that help with
// interacting with the alphaxiv_cat API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewNotificationV4ArchiveService] method instead.
type NotificationV4ArchiveService struct {
	options []option.RequestOption
}

// NewNotificationV4ArchiveService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewNotificationV4ArchiveService(opts ...option.RequestOption) (r NotificationV4ArchiveService) {
	r = NotificationV4ArchiveService{}
	r.options = opts
	return
}
