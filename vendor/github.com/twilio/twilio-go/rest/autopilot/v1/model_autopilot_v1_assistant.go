/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Autopilot
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

import (
	"time"
)

// AutopilotV1Assistant struct for AutopilotV1Assistant
type AutopilotV1Assistant struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Assistant resource.
	AccountSid *string `json:"account_sid,omitempty"`
	// The date and time in GMT when the resource was created specified in [RFC 2822](https://www.ietf.org/rfc/rfc2822.txt) format.
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The date and time in GMT when the resource was last updated specified in [RFC 2822](https://www.ietf.org/rfc/rfc2822.txt) format.
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// The string that you assigned to describe the resource. It is not unique and can be up to 255 characters long.
	FriendlyName *string `json:"friendly_name,omitempty"`
	// Reserved.
	LatestModelBuildSid *string `json:"latest_model_build_sid,omitempty"`
	// A list of the URLs of the Assistant's related resources.
	Links *map[string]interface{} `json:"links,omitempty"`
	// Whether queries should be logged and kept after training. Can be: `true` or `false` and defaults to `true`. If `true`, queries are stored for 30 days, and then deleted. If `false`, no queries are stored.
	LogQueries *bool `json:"log_queries,omitempty"`
	// A string describing the state of the assistant.
	DevelopmentStage *string `json:"development_stage,omitempty"`
	// Whether model needs to be rebuilt.
	NeedsModelBuild *bool `json:"needs_model_build,omitempty"`
	// The unique string that we created to identify the Assistant resource.
	Sid *string `json:"sid,omitempty"`
	// An application-defined string that uniquely identifies the resource. It can be used in place of the resource's `sid` in the URL to address the resource. It can be up to 64 characters long.
	UniqueName *string `json:"unique_name,omitempty"`
	// The absolute URL of the Assistant resource.
	Url *string `json:"url,omitempty"`
	// Reserved.
	CallbackUrl *string `json:"callback_url,omitempty"`
	// Reserved.
	CallbackEvents *string `json:"callback_events,omitempty"`
}
