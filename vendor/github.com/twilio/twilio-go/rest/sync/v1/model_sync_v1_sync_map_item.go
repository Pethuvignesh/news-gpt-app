/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Sync
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

// SyncV1SyncMapItem struct for SyncV1SyncMapItem
type SyncV1SyncMapItem struct {
	// The unique, user-defined key for the Map Item.
	Key *string `json:"key,omitempty"`
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Map Item resource.
	AccountSid *string `json:"account_sid,omitempty"`
	// The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) the resource is associated with.
	ServiceSid *string `json:"service_sid,omitempty"`
	// The SID of the Sync Map that contains the Map Item.
	MapSid *string `json:"map_sid,omitempty"`
	// The absolute URL of the Map Item resource.
	Url *string `json:"url,omitempty"`
	// The current revision of the Map Item, represented as a string.
	Revision *string `json:"revision,omitempty"`
	// An arbitrary, schema-less object that the Map Item stores. Can be up to 16 KiB in length.
	Data *interface{} `json:"data,omitempty"`
	// The date and time in GMT when the Map Item expires and will be deleted, specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. If the Map Item does not expire, this value is `null`.  The Map Item might not be deleted immediately after it expires.
	DateExpires *time.Time `json:"date_expires,omitempty"`
	// The date and time in GMT when the resource was created specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The date and time in GMT when the resource was last updated specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// The identity of the Map Item's creator. If the Map Item is created from the client SDK, the value matches the Access Token's `identity` field. If the Map Item was created from the REST API, the value is `system`.
	CreatedBy *string `json:"created_by,omitempty"`
}