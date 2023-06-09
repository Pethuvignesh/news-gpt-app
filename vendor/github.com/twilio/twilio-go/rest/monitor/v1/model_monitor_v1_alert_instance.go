/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Monitor
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

// MonitorV1AlertInstance struct for MonitorV1AlertInstance
type MonitorV1AlertInstance struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Alert resource.
	AccountSid *string `json:"account_sid,omitempty"`
	// The text of the alert.
	AlertText *string `json:"alert_text,omitempty"`
	// The API version used when the alert was generated.  Can be empty for events that don't have a specific API version.
	ApiVersion *string `json:"api_version,omitempty"`
	// The date and time in GMT when the resource was created specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The date and time in GMT when the alert was generated specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601#UTC) format.  Due to buffering, this can be different than `date_created`.
	DateGenerated *time.Time `json:"date_generated,omitempty"`
	// The date and time in GMT when the resource was last updated specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// The error code for the condition that generated the alert. See the [Error Dictionary](https://www.twilio.com/docs/api/errors) for possible causes and solutions to the error.
	ErrorCode *string `json:"error_code,omitempty"`
	// The log level.  Can be: `error`, `warning`, `notice`, or `debug`.
	LogLevel *string `json:"log_level,omitempty"`
	// The URL of the page in our [Error Dictionary](https://www.twilio.com/docs/api/errors) with more information about the error condition.
	MoreInfo *string `json:"more_info,omitempty"`
	// The method used by the request that generated the alert. If the alert was generated by a request we made to your server, this is the method we used. If the alert was generated by a request from your application to our API, this is the method your application used.
	RequestMethod *string `json:"request_method,omitempty"`
	// The URL of the request that generated the alert. If the alert was generated by a request we made to your server, this is the URL on your server that generated the alert. If the alert was generated by a request from your application to our API, this is the URL of the resource requested.
	RequestUrl *string `json:"request_url,omitempty"`
	// The variables passed in the request that generated the alert. This value is only returned when a single Alert resource is fetched.
	RequestVariables *string `json:"request_variables,omitempty"`
	// The SID of the resource for which the alert was generated.  For instance, if your server failed to respond to an HTTP request during the flow of a particular call, this value would be the SID of the server.  This value is empty if the alert was not generated for a particular resource.
	ResourceSid *string `json:"resource_sid,omitempty"`
	// The response body of the request that generated the alert. This value is only returned when a single Alert resource is fetched.
	ResponseBody *string `json:"response_body,omitempty"`
	// The response headers of the request that generated the alert. This value is only returned when a single Alert resource is fetched.
	ResponseHeaders *string `json:"response_headers,omitempty"`
	// The unique string that we created to identify the Alert resource.
	Sid *string `json:"sid,omitempty"`
	// The absolute URL of the Alert resource.
	Url *string `json:"url,omitempty"`
	// The request headers of the request that generated the alert. This value is only returned when a single Alert resource is fetched.
	RequestHeaders *string `json:"request_headers,omitempty"`
	// The SID of the service or resource that generated the alert. Can be `null`.
	ServiceSid *string `json:"service_sid,omitempty"`
}
