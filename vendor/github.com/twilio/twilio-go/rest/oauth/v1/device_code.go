/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Oauth
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

import (
	"encoding/json"
	"net/url"
)

// Optional parameters for the method 'CreateDeviceCode'
type CreateDeviceCodeParams struct {
	// A 34 character string that uniquely identifies this OAuth App.
	ClientSid *string `json:"ClientSid,omitempty"`
	// An Array of scopes for authorization request
	Scopes *[]string `json:"Scopes,omitempty"`
	// An array of intended audiences for token requests
	Audiences *[]string `json:"Audiences,omitempty"`
}

func (params *CreateDeviceCodeParams) SetClientSid(ClientSid string) *CreateDeviceCodeParams {
	params.ClientSid = &ClientSid
	return params
}
func (params *CreateDeviceCodeParams) SetScopes(Scopes []string) *CreateDeviceCodeParams {
	params.Scopes = &Scopes
	return params
}
func (params *CreateDeviceCodeParams) SetAudiences(Audiences []string) *CreateDeviceCodeParams {
	params.Audiences = &Audiences
	return params
}

// Issues a new Access token (optionally identity_token & refresh_token) in exchange of Oauth grant
func (c *ApiService) CreateDeviceCode(params *CreateDeviceCodeParams) (*OauthV1DeviceCode, error) {
	path := "/v1/device/code"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.ClientSid != nil {
		data.Set("ClientSid", *params.ClientSid)
	}
	if params != nil && params.Scopes != nil {
		for _, item := range *params.Scopes {
			data.Add("Scopes", item)
		}
	}
	if params != nil && params.Audiences != nil {
		for _, item := range *params.Audiences {
			data.Add("Audiences", item)
		}
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &OauthV1DeviceCode{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
