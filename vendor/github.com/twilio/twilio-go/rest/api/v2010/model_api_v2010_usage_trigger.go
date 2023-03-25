/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Api
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

// ApiV2010UsageTrigger struct for ApiV2010UsageTrigger
type ApiV2010UsageTrigger struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that the trigger monitors.
	AccountSid *string `json:"account_sid,omitempty"`
	// The API version used to create the resource.
	ApiVersion *string `json:"api_version,omitempty"`
	// The HTTP method we use to call `callback_url`. Can be: `GET` or `POST`.
	CallbackMethod *string `json:"callback_method,omitempty"`
	// The URL we call using the `callback_method` when the trigger fires.
	CallbackUrl *string `json:"callback_url,omitempty"`
	// The current value of the field the trigger is watching.
	CurrentValue *string `json:"current_value,omitempty"`
	// The date and time in GMT that the resource was created specified in [RFC 2822](https://www.ietf.org/rfc/rfc2822.txt) format.
	DateCreated *string `json:"date_created,omitempty"`
	// The date and time in GMT that the trigger was last fired specified in [RFC 2822](https://www.ietf.org/rfc/rfc2822.txt) format.
	DateFired *string `json:"date_fired,omitempty"`
	// The date and time in GMT that the resource was last updated specified in [RFC 2822](https://www.ietf.org/rfc/rfc2822.txt) format.
	DateUpdated *string `json:"date_updated,omitempty"`
	// The string that you assigned to describe the trigger.
	FriendlyName *string `json:"friendly_name,omitempty"`
	Recurring    *string `json:"recurring,omitempty"`
	// The unique string that that we created to identify the UsageTrigger resource.
	Sid       *string `json:"sid,omitempty"`
	TriggerBy *string `json:"trigger_by,omitempty"`
	// The value at which the trigger will fire.  Must be a positive, numeric value.
	TriggerValue *string `json:"trigger_value,omitempty"`
	// The URI of the resource, relative to `https://api.twilio.com`.
	Uri           *string `json:"uri,omitempty"`
	UsageCategory *string `json:"usage_category,omitempty"`
	// The URI of the [UsageRecord](https://www.twilio.com/docs/usage/api/usage-record) resource this trigger watches, relative to `https://api.twilio.com`.
	UsageRecordUri *string `json:"usage_record_uri,omitempty"`
}
