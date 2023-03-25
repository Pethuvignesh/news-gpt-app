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

// ApiV2010Message struct for ApiV2010Message
type ApiV2010Message struct {
	// The message text. Can be up to 1,600 characters long.
	Body *string `json:"body,omitempty"`
	// The number of segments that make up the complete message. A message body that is too large to be sent in a single SMS message is segmented and charged as multiple messages. Inbound messages over 160 characters are reassembled when the message is received. Note: When using a Messaging Service to send messages, `num_segments` will always be 0 in Twilio's response to your API request.
	NumSegments *string `json:"num_segments,omitempty"`
	Direction   *string `json:"direction,omitempty"`
	// The phone number (in [E.164](https://en.wikipedia.org/wiki/E.164) format), [alphanumeric sender ID](https://www.twilio.com/docs/sms/send-messages#use-an-alphanumeric-sender-id), or [Wireless SIM](https://www.twilio.com/docs/wireless/tutorials/communications-guides/how-to-send-and-receive-text-messages) that initiated the message. For incoming messages, this will be the number of the sending phone. For outgoing messages, this value will be one of your Twilio phone numbers or the alphanumeric sender ID used.
	From *string `json:"from,omitempty"`
	// The phone number in [E.164](https://en.wikipedia.org/wiki/E.164) format that received the message. For incoming messages, this will be one of your Twilio phone numbers. For outgoing messages, this will be the sending phone.
	To *string `json:"to,omitempty"`
	// The date and time in GMT that the resource was last updated specified in [RFC 2822](https://www.ietf.org/rfc/rfc2822.txt) format.
	DateUpdated *string `json:"date_updated,omitempty"`
	// The amount billed for the message, in the currency specified by `price_unit`.  Note that your account is charged for each segment we send to the handset. Populated after the message has been sent. May not be immediately available.
	Price *string `json:"price,omitempty"`
	// The description of the `error_code` if your message `status` is `failed` or `undelivered`. If the message was successful, this value is null.
	ErrorMessage *string `json:"error_message,omitempty"`
	// The URI of the resource, relative to `https://api.twilio.com`.
	Uri *string `json:"uri,omitempty"`
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that sent the message that created the resource.
	AccountSid *string `json:"account_sid,omitempty"`
	// The number of media files associated with the message. A message can send up to 10 media files.
	NumMedia *string `json:"num_media,omitempty"`
	Status   *string `json:"status,omitempty"`
	// The SID of the [Messaging Service](https://www.twilio.com/docs/sms/services/api) used with the message. The value is null if a Messaging Service was not used.
	MessagingServiceSid *string `json:"messaging_service_sid,omitempty"`
	// The unique string that that we created to identify the Message resource.
	Sid *string `json:"sid,omitempty"`
	// The date and time in GMT that the resource was sent specified in [RFC 2822](https://www.ietf.org/rfc/rfc2822.txt) format. For outgoing messages, this is when we sent the message. For incoming messages, this is when we made the HTTP request to your application.
	DateSent *string `json:"date_sent,omitempty"`
	// The date and time in GMT that the resource was created specified in [RFC 2822](https://www.ietf.org/rfc/rfc2822.txt) format.
	DateCreated *string `json:"date_created,omitempty"`
	// The error code returned if your message `status` is `failed` or `undelivered`. The error_code provides more information about the failure. If the message was successful, this value is null.
	ErrorCode *int `json:"error_code,omitempty"`
	// The currency in which `price` is measured, in [ISO 4127](https://www.iso.org/iso/home/standards/currency_codes.htm) format (e.g. `usd`, `eur`, `jpy`).
	PriceUnit *string `json:"price_unit,omitempty"`
	// The API version used to process the message.
	ApiVersion *string `json:"api_version,omitempty"`
	// A list of related resources identified by their URIs relative to `https://api.twilio.com`
	SubresourceUris *map[string]interface{} `json:"subresource_uris,omitempty"`
}
