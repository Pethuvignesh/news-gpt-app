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

// ConversationsV1ConfigurationAddress struct for ConversationsV1ConfigurationAddress
type ConversationsV1ConfigurationAddress struct {
	// A 34 character string that uniquely identifies this resource.
	Sid *string `json:"sid,omitempty"`
	// The unique ID of the [Account](https://www.twilio.com/docs/iam/api/account) the address belongs to
	AccountSid *string `json:"account_sid,omitempty"`
	// Type of Address, value can be `whatsapp` or `sms`.
	Type *string `json:"type,omitempty"`
	// The unique address to be configured. The address can be a whatsapp address or phone number
	Address *string `json:"address,omitempty"`
	// The human-readable name of this configuration, limited to 256 characters. Optional.
	FriendlyName *string `json:"friendly_name,omitempty"`
	// Auto Creation configuration for the address.
	AutoCreation *interface{} `json:"auto_creation,omitempty"`
	// The date that this resource was created.
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The date that this resource was last updated.
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// An absolute API resource URL for this address configuration.
	Url *string `json:"url,omitempty"`
}