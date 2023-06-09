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

// AutopilotV1Query struct for AutopilotV1Query
type AutopilotV1Query struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Query resource.
	AccountSid *string `json:"account_sid,omitempty"`
	// The date and time in GMT when the resource was created specified in [RFC 2822](https://www.ietf.org/rfc/rfc2822.txt) format.
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The date and time in GMT when the resource was last updated specified in [RFC 2822](https://www.ietf.org/rfc/rfc2822.txt) format.
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// The natural language analysis results that include the [Task](https://www.twilio.com/docs/autopilot/api/task) recognized and a list of identified [Fields](https://www.twilio.com/docs/autopilot/api/task-field).
	Results *interface{} `json:"results,omitempty"`
	// The [ISO language-country](https://docs.oracle.com/cd/E13214_01/wli/docs92/xref/xqisocodes.html) string that specifies the language used by the Query. For example: `en-US`.
	Language *string `json:"language,omitempty"`
	// The SID of the [Model Build](https://www.twilio.com/docs/autopilot/api/model-build) queried.
	ModelBuildSid *string `json:"model_build_sid,omitempty"`
	// The end-user's natural language input.
	Query *string `json:"query,omitempty"`
	// The SID of an optional reference to the [Sample](https://www.twilio.com/docs/autopilot/api/task-sample) created from the query.
	SampleSid *string `json:"sample_sid,omitempty"`
	// The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the resource.
	AssistantSid *string `json:"assistant_sid,omitempty"`
	// The unique string that we created to identify the Query resource.
	Sid *string `json:"sid,omitempty"`
	// The status of the Query. Can be: `pending-review`, `reviewed`, or `discarded`
	Status *string `json:"status,omitempty"`
	// The absolute URL of the Query resource.
	Url *string `json:"url,omitempty"`
	// The communication channel from where the end-user input came.
	SourceChannel *string `json:"source_channel,omitempty"`
	// The SID of the [Dialogue](https://www.twilio.com/docs/autopilot/api/dialogue).
	DialogueSid *string `json:"dialogue_sid,omitempty"`
}
