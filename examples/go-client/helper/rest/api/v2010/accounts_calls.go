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


// Optional parameters for the method 'CreateCall'
type CreateCallParams struct {
    // 
    PathAccountSid *string `json:"PathAccountSid,omitempty"`
    // 
    RequiredStringProperty *string `json:"RequiredStringProperty,omitempty"`
    // 
    TestArrayOfStrings *[]string `json:"TestArrayOfStrings,omitempty"`
}

func (params *CreateCallParams) SetPathAccountSid(PathAccountSid string) (*CreateCallParams){
    params.PathAccountSid = &PathAccountSid
    return params
}
func (params *CreateCallParams) SetRequiredStringProperty(RequiredStringProperty string) (*CreateCallParams){
    params.RequiredStringProperty = &RequiredStringProperty
    return params
}
func (params *CreateCallParams) SetTestArrayOfStrings(TestArrayOfStrings []string) (*CreateCallParams){
    params.TestArrayOfStrings = &TestArrayOfStrings
    return params
}

func (c *ApiService) CreateCall(params *CreateCallParams) (*TestResponseObject, error) {
    path := "/2010-04-01/Accounts/{AccountSid}/Calls.json"
    if params != nil && params.PathAccountSid != nil {
    path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
} else {
    path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
}

data := url.Values{}
headers := make(map[string]interface{})

if params != nil && params.RequiredStringProperty != nil {
    data.Set("RequiredStringProperty", *params.RequiredStringProperty)
}
if params != nil && params.TestArrayOfStrings != nil {
    for _, item  := range *params.TestArrayOfStrings {
        data.Add("TestArrayOfStrings", item)
    }
}



    resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    ps := &TestResponseObject{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }

    return ps, err
}

// Optional parameters for the method 'DeleteCall'
type DeleteCallParams struct {
    // 
    PathAccountSid *string `json:"PathAccountSid,omitempty"`
}

func (params *DeleteCallParams) SetPathAccountSid(PathAccountSid string) (*DeleteCallParams){
    params.PathAccountSid = &PathAccountSid
    return params
}

func (c *ApiService) DeleteCall(Sid string, params *DeleteCallParams) (error) {
    path := "/2010-04-01/Accounts/{AccountSid}/Calls/{Sid}.json"
    if params != nil && params.PathAccountSid != nil {
    path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
} else {
    path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
}
    path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

data := url.Values{}
headers := make(map[string]interface{})




    resp, err := c.requestHandler.Delete(c.baseURL+path, data, headers)
    if err != nil {
        return err
    }

    defer resp.Body.Close()

    return nil
}

// Optional parameters for the method 'FetchCall'
type FetchCallParams struct {
    // 
    PathAccountSid *string `json:"PathAccountSid,omitempty"`
}

func (params *FetchCallParams) SetPathAccountSid(PathAccountSid string) (*FetchCallParams){
    params.PathAccountSid = &PathAccountSid
    return params
}

func (c *ApiService) FetchCall(Sid string, params *FetchCallParams) (*TestResponseObject, error) {
    path := "/2010-04-01/Accounts/{AccountSid}/Calls/{Sid}.json"
    if params != nil && params.PathAccountSid != nil {
    path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
} else {
    path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
}
    path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

data := url.Values{}
headers := make(map[string]interface{})




    resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    ps := &TestResponseObject{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }

    return ps, err
}

// Optional parameters for the method 'ListCall'
type ListCallParams struct {
    // 
    PathAccountSid *string `json:"PathAccountSid,omitempty"`
    // 
    To *string `json:"To,omitempty"`
    // 
    From *string `json:"From,omitempty"`
    // 
    PageSize *int `json:"PageSize,omitempty"`
    // Max number of records to return.
    Limit *int `json:"limit,omitempty"`
}

func (params *ListCallParams) SetPathAccountSid(PathAccountSid string) (*ListCallParams){
    params.PathAccountSid = &PathAccountSid
    return params
}
func (params *ListCallParams) SetTo(To string) (*ListCallParams){
    params.To = &To
    return params
}
func (params *ListCallParams) SetFrom(From string) (*ListCallParams){
    params.From = &From
    return params
}
func (params *ListCallParams) SetPageSize(PageSize int) (*ListCallParams){
    params.PageSize = &PageSize
    return params
}
func (params *ListCallParams) SetLimit(Limit int) (*ListCallParams){
    params.Limit = &Limit
    return params
}

// Retrieve a single page of Call records from the API. Request is executed immediately.
func (c *ApiService) PageCall(params *ListCallParams, pageToken, pageNumber string) (*ListCallResponse, error) {
    path := "/2010-04-01/Accounts/{AccountSid}/Calls.json"

    if params != nil && params.PathAccountSid != nil {
    path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
} else {
    path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
}

data := url.Values{}
headers := make(map[string]interface{})

if params != nil && params.To != nil {
    data.Set("To", *params.To)
}
if params != nil && params.From != nil {
    data.Set("From", *params.From)
}
if params != nil && params.PageSize != nil {
    data.Set("PageSize", fmt.Sprint(*params.PageSize))
}

    if pageToken != "" {
        data.Set("PageToken", pageToken)
    }
    if pageNumber != "" {
        data.Set("Page", pageNumber)
    }

    resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    ps := &ListCallResponse{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }

    return ps, err
}

// Lists Call records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListCall(params *ListCallParams) ([]TestResponseObject, error) {
    response, err := c.StreamCall(params)
    if err != nil {
        return nil, err
    }

    records := make([]TestResponseObject, 0)

    for record := range response {
        records = append(records, record)
    }

    return records, err
}

// Streams Call records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamCall(params *ListCallParams) (chan TestResponseObject, error) {
    if params == nil {
        params = &ListCallParams{}
    }
    params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

    response, err := c.PageCall(params, "", "")
    if err != nil {
        return nil, err
    }

    curRecord := 1
    //set buffer size of the channel to 1
    channel := make(chan TestResponseObject, 1)

    go func() {
        for response != nil {
            responseRecords := response.Calls
            for item := range responseRecords {
                channel <- responseRecords[item]
                curRecord += 1
                if params.Limit != nil && *params.Limit < curRecord {
                    close(channel)
                    return
                }
            }

            var record interface{}
            if record, err = client.GetNext(c.baseURL, response, c.getNextListCallResponse); record == nil || err != nil {
                close(channel)
                return
            }

            response = record.(*ListCallResponse)
        }
        close(channel)
    }()

    return channel, err
}

func (c *ApiService) getNextListCallResponse(nextPageUrl string) (interface{}, error) {
    if nextPageUrl == "" {
        return nil, nil
    }
    resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    ps := &ListCallResponse{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }
    return ps, nil
}

