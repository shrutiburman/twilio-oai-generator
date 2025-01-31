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

package com.twilio.rest.api.v2010.credential;

import com.twilio.base.Updater;
import com.twilio.converter.Promoter;
import com.twilio.exception.ApiConnectionException;
import com.twilio.exception.ApiException;
import com.twilio.exception.RestException;
import com.twilio.http.HttpMethod;
import com.twilio.http.Request;
import com.twilio.http.Response;
import com.twilio.http.TwilioRestClient;
import com.twilio.rest.Domains;

/*
    * Twilio - Accounts
    *
    * This is the public Twilio REST API.
    *
    * API version: 1.11.0
    * Contact: support@twilio.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.


public class AwsUpdater extends Updater<Aws>{
    private String Sid;
    private String TestString;

    public AwsUpdater(final String Sid){
        
            this.Sid = Sid;
        
        }

    public AwsUpdater setSid(final String Sid){
        this.Sid = Sid;
        return this;
    }
    public AwsUpdater setTestString(final String TestString){
        this.TestString = TestString;
        return this;
    }

    @Override
    public Aws update(final TwilioRestClient client){
        String path = "/v1/Credentials/AWS/{Sid}";
        path = path.replace("{"+"Sid"+"}", this.Sid.toString());

        Request request = new Request(
            HttpMethod.POST,
            Domains.API.toString(),
            path
        );

        addPostParams(request);
        Response response = client.request(request);
        if (response == null) {
            throw new ApiConnectionException("Aws update failed: Unable to connect to server");
        } else if (!TwilioRestClient.SUCCESS.test(response.getStatusCode())) {
            RestException restException = RestException.fromJson(response.getStream(), client.getObjectMapper());
            if (restException == null) {
                throw new ApiException("Server Error, no content");
            }
            throw new ApiException(restException);
        }

        return Aws.fromJson(response.getStream(), client.getObjectMapper());
    }


    private void addPostParams(final Request request) {
        if (Sid != null) {
            request.addPostParam("Sid", Sid.toString());
        }
        if (TestString != null) {
            request.addPostParam("TestString", TestString.toString());
        }
    }
}
