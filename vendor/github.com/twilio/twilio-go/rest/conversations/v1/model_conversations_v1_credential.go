/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Conversations
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

// ConversationsV1Credential struct for ConversationsV1Credential
type ConversationsV1Credential struct {
	// A 34 character string that uniquely identifies this resource.
	Sid *string `json:"sid,omitempty"`
	// The unique ID of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this credential.
	AccountSid *string `json:"account_sid,omitempty"`
	// The human-readable name of this credential, limited to 64 characters. Optional.
	FriendlyName *string `json:"friendly_name,omitempty"`
	Type         *string `json:"type,omitempty"`
	// [APN only] Whether to send the credential to sandbox APNs. Can be `true` to send to sandbox APNs or `false` to send to production.
	Sandbox *string `json:"sandbox,omitempty"`
	// The date that this resource was created.
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The date that this resource was last updated.
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// An absolute API resource URL for this credential.
	Url *string `json:"url,omitempty"`
}
