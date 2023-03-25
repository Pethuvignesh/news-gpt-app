/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Verify
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

// VerifyV2Service struct for VerifyV2Service
type VerifyV2Service struct {
	// The unique string that we created to identify the Service resource.
	Sid *string `json:"sid,omitempty"`
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Service resource.
	AccountSid *string `json:"account_sid,omitempty"`
	// The string that you assigned to describe the verification service. **This value should not contain PII.**
	FriendlyName *string `json:"friendly_name,omitempty"`
	// The length of the verification code to generate.
	CodeLength *int `json:"code_length,omitempty"`
	// Whether to perform a lookup with each verification started and return info about the phone number.
	LookupEnabled *bool `json:"lookup_enabled,omitempty"`
	// Whether to pass PSD2 transaction parameters when starting a verification.
	Psd2Enabled *bool `json:"psd2_enabled,omitempty"`
	// Whether to skip sending SMS verifications to landlines. Requires `lookup_enabled`.
	SkipSmsToLandlines *bool `json:"skip_sms_to_landlines,omitempty"`
	// Whether to ask the user to press a number before delivering the verify code in a phone call.
	DtmfInputRequired *bool `json:"dtmf_input_required,omitempty"`
	// The name of an alternative text-to-speech service to use in phone calls. Applies only to TTS languages.
	TtsName *string `json:"tts_name,omitempty"`
	// Whether to add a security warning at the end of an SMS verification body. Disabled by default and applies only to SMS. Example SMS body: `Your AppName verification code is: 1234. Don’t share this code with anyone; our employees will never ask for the code`
	DoNotShareWarningEnabled *bool `json:"do_not_share_warning_enabled,omitempty"`
	// Whether to allow sending verifications with a custom code instead of a randomly generated one. Not available for all customers.
	CustomCodeEnabled *bool `json:"custom_code_enabled,omitempty"`
	// Configurations for the Push factors (channel) created under this Service.
	Push *interface{} `json:"push,omitempty"`
	// Configurations for the TOTP factors (channel) created under this Service.
	Totp               *interface{} `json:"totp,omitempty"`
	DefaultTemplateSid *string      `json:"default_template_sid,omitempty"`
	// The date and time in GMT when the resource was created specified in [RFC 2822](https://www.ietf.org/rfc/rfc2822.txt) format.
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The date and time in GMT when the resource was last updated specified in [RFC 2822](https://www.ietf.org/rfc/rfc2822.txt) format.
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// The absolute URL of the resource.
	Url *string `json:"url,omitempty"`
	// The URLs of related resources.
	Links *map[string]interface{} `json:"links,omitempty"`
}
