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

import (
	"encoding/json"
	"fmt"
	"net/url"

    "github.com/twilio/twilio-go/client"
)


// Optional parameters for the method 'CreateCallFeedbackSummary'
type CreateCallFeedbackSummaryParams struct {
    // The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource.
    PathAccountSid *string `json:"PathAccountSid,omitempty"`
    // Only include feedback given on or before this date. Format is `YYYY-MM-DD` and specified in UTC.
    EndDate *string `json:"EndDate,omitempty"`
    // Only include feedback given on or after this date. Format is `YYYY-MM-DD` and specified in UTC.
    StartDate *string `json:"StartDate,omitempty"`
}

func (params *CreateCallFeedbackSummaryParams) SetPathAccountSid(PathAccountSid string) (*CreateCallFeedbackSummaryParams){
    params.PathAccountSid = &PathAccountSid
    return params
}
func (params *CreateCallFeedbackSummaryParams) SetEndDate(EndDate string) (*CreateCallFeedbackSummaryParams){
    params.EndDate = &EndDate
    return params
}
func (params *CreateCallFeedbackSummaryParams) SetStartDate(StartDate string) (*CreateCallFeedbackSummaryParams){
    params.StartDate = &StartDate
    return params
}

// Create a FeedbackSummary resource for a call
func (c *ApiService) CreateCallFeedbackSummary(params *CreateCallFeedbackSummaryParams) (*ApiV2010CallFeedbackSummary, error) {
    path := "/2010-04-01/Accounts/{AccountSid}/Calls/FeedbackSummary.json"
    if params != nil && params.PathAccountSid != nil {
    path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
} else {
    path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
}

data := url.Values{}
headers := make(map[string]interface{})

if params != nil && params.EndDate != nil {
    data.Set("EndDate", fmt.Sprint(*params.EndDate))
}
if params != nil && params.StartDate != nil {
    data.Set("StartDate", fmt.Sprint(*params.StartDate))
}



    resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    ps := &ApiV2010CallFeedbackSummary{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }

    return ps, err
}
