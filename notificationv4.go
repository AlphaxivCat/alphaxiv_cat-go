// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package alphaxivcat

import (
	"github.com/AlphaxivCat/alphaxiv_cat-go/option"
)

// NotificationV4Service contains methods and other services that help with
// interacting with the alphaxiv_cat API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewNotificationV4Service] method instead.
type NotificationV4Service struct {
	options []option.RequestOption
	Archive NotificationV4ArchiveService
}

// NewNotificationV4Service generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewNotificationV4Service(opts ...option.RequestOption) (r NotificationV4Service) {
	r = NotificationV4Service{}
	r.options = opts
	r.Archive = NewNotificationV4ArchiveService(opts...)
	return
}
