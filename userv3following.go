// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package alphaxivcat

import (
	"github.com/AlphaxivCat/alphaxiv_cat-go/option"
)

// UserV3FollowingService contains methods and other services that help with
// interacting with the alphaxiv_cat API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserV3FollowingService] method instead.
type UserV3FollowingService struct {
	options       []option.RequestOption
	Topics        UserV3FollowingTopicService
	Organizations UserV3FollowingOrganizationService
}

// NewUserV3FollowingService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewUserV3FollowingService(opts ...option.RequestOption) (r UserV3FollowingService) {
	r = UserV3FollowingService{}
	r.options = opts
	r.Topics = NewUserV3FollowingTopicService(opts...)
	r.Organizations = NewUserV3FollowingOrganizationService(opts...)
	return
}
