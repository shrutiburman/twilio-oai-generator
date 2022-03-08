/*
    * Twilio - Accounts
    *
    * This is the public Twilio REST API.
    *
    * API version: 1.11.0
    * Contact: support@twilio.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi
// ApiV2010Message struct for ApiV2010Message
type ApiV2010Message struct {
        // The SID of the Account that created the resource
    AccountSid *string `json:"account_sid,omitempty"`
        // The message text
    Body *string `json:"body,omitempty"`
        // The direction of the message
    Direction *string `json:"direction,omitempty"`
        // The phone number that initiated the message
    From *string `json:"from,omitempty"`
        // The unique string that identifies the resource
    Sid *string `json:"sid,omitempty"`
        // The status of the message
    Status *string `json:"status,omitempty"`
        // The phone number that received the message
    To *string `json:"to,omitempty"`
        // The URI of the resource, relative to `https://api.twilio.com`
    Uri *string `json:"uri,omitempty"`
}
