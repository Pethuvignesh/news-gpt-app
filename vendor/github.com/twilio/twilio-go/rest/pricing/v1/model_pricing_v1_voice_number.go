/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Pricing
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

// PricingV1VoiceNumber struct for PricingV1VoiceNumber
type PricingV1VoiceNumber struct {
	// The phone number.
	Number *string `json:"number,omitempty"`
	// The name of the country.
	Country *string `json:"country,omitempty"`
	// The [ISO country code](http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2).
	IsoCountry        *string                                     `json:"iso_country,omitempty"`
	OutboundCallPrice *PricingV1VoiceVoiceNumberOutboundCallPrice `json:"outbound_call_price,omitempty"`
	InboundCallPrice  *PricingV1VoiceVoiceNumberInboundCallPrice  `json:"inbound_call_price,omitempty"`
	// The currency in which prices are measured, specified in [ISO 4127](http://www.iso.org/iso/home/standards/currency_codes.htm) format (e.g. `usd`, `eur`, `jpy`).
	PriceUnit *string `json:"price_unit,omitempty"`
	// The absolute URL of the resource.
	Url *string `json:"url,omitempty"`
}
