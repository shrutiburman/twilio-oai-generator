/*
 * Twilio - Accounts
 * This is the public Twilio REST API.
 *
 * The version of the OpenAPI document: 1.11.0
 * Contact: support@twilio.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package com.twilio.rest.api.v2010.account;

import com.twilio.base.Creator;
import com.twilio.converter.Promoter;
import com.twilio.exception.ApiConnectionException;
import com.twilio.exception.ApiException;
import com.twilio.exception.RestException;
import com.twilio.http.HttpMethod;
import com.twilio.http.Request;
import com.twilio.http.Response;
import com.twilio.http.TwilioRestClient;
import com.twilio.rest.Domains;
import java.math.BigDecimal;
import java.util.List;
import java.util.Map;
import java.time.LocalDate;
import com.twilio.converter.Converter;
import java.time.ZonedDateTime;

import java.net.URI;

/*
    * Twilio - Accounts
    *
    * This is the public Twilio REST API.
    *
    * API version: 1.11.0
    * Contact: support@twilio.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.


public class CallCreator extends Creator<Call>{
    private String AccountSid;
    private String RequiredStringProperty;
    private List<String> TestArrayOfStrings;
    public CallCreator(final String AccountSid, final String RequiredStringProperty) {
    
        this.AccountSid = AccountSid;
    
        this.RequiredStringProperty = RequiredStringProperty;
    
    }

    public CallCreator setAccountSid(final String AccountSid){
        this.AccountSid = AccountSid;
        return this;
    }
    public CallCreator setRequiredStringProperty(final String RequiredStringProperty){
        this.RequiredStringProperty = RequiredStringProperty;
        return this;
    }
    public CallCreator setTestArrayOfStrings(final List<String> TestArrayOfStrings){
        this.TestArrayOfStrings = TestArrayOfStrings;
        return this;
    }

    @Override
    public Call create(final TwilioRestClient client){
        String path = "/2010-04-01/Accounts/{AccountSid}/Calls.json";
        path = path.replace("{"+"AccountSid"+"}", this.AccountSid.toString());
        path = path.replace("{"+"RequiredStringProperty"+"}", this.RequiredStringProperty.toString());

        Request request = new Request(
            HttpMethod.POST,
            Domains.API.toString(),
            path
        );

        addPostParams(request);
        Response response = client.request(request);
        if (response == null) {
            throw new ApiConnectionException("Call creation failed: Unable to connect to server");
        } else if (!TwilioRestClient.SUCCESS.test(response.getStatusCode())) {
            RestException restException = RestException.fromJson(response.getStream(), client.getObjectMapper());
            if (restException == null) {
                throw new ApiException("Server Error, no content");
            }
            throw new ApiException(restException);
        }

        return Call.fromJson(response.getStream(), client.getObjectMapper());
    }


    private void addPostParams(final Request request) {
        if (AccountSid != null) {
        request.addPostParam("AccountSid", AccountSid.toString());
        }

        if (RequiredStringProperty != null) {
        request.addPostParam("RequiredStringProperty", RequiredStringProperty.toString());
        }

        if (TestArrayOfStrings != null) {
        request.addPostParam("TestArrayOfStrings", TestArrayOfStrings.toString());
        }

    }
}
