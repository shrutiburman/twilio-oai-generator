/**
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

import { inspect, InspectOptions } from 'util';
import Page from '../../../../../base/Page';
import V2010 from '../../../V2010';


/**
 * Options to pass to create a FlexFlowInstance
 *
 * @property { string } channelType The channel type. One of &#x60;web&#x60;, &#x60;facebook&#x60;, &#x60;sms&#x60;, &#x60;whatsapp&#x60;, &#x60;line&#x60; or &#x60;custom&#x60;. By default, Studio’s Send to Flex widget passes it on to the Task attributes for Tasks created based on this Flex Flow. The Task attributes will be used by the Flex UI to render the respective Task as appropriate (applying channel-specific design and length limits). If &#x60;channelType&#x60; is &#x60;facebook&#x60;, &#x60;whatsapp&#x60; or &#x60;line&#x60;, the Send to Flex widget should set the Task Channel to Programmable Chat.
 * @property { string } chatServiceSid The SID of the chat service.
 * @property { string } friendlyName A descriptive string that you create to describe the Flex Flow resource.
 * @property { string } [contactIdentity] The channel contact\\\&#39;s Identity.
 * @property { boolean } [enabled] Whether the new Flex Flow is enabled.
 * @property { string } [integrationChannel] The Task Channel SID (TCXXXX) or unique name (e.g., &#x60;sms&#x60;) to use for the Task that will be created. Applicable and required when &#x60;integrationType&#x60; is &#x60;task&#x60;. The default value is &#x60;default&#x60;.
 * @property { boolean } [integrationCreationOnMessage] In the context of outbound messaging, defines whether to create a Task immediately (and therefore reserve the conversation to current agent), or delay Task creation until the customer sends the first response. Set to false to create immediately, true to delay Task creation. This setting is only applicable for outbound messaging.
 * @property { string } [integrationFlowSid] The SID of the Studio Flow. Required when &#x60;integrationType&#x60; is &#x60;studio&#x60;.
 * @property { number } [integrationPriority] The Task priority of a new Task. The default priority is 0. Optional when &#x60;integrationType&#x60; is &#x60;task&#x60;, not applicable otherwise.
 * @property { number } [integrationRetryCount] The number of times to retry the Studio Flow or webhook in case of failure. Takes integer values from 0 to 3 with the default being 3. Optional when &#x60;integrationType&#x60; is &#x60;studio&#x60; or &#x60;external&#x60;, not applicable otherwise.
 * @property { number } [integrationTimeout] The Task timeout in seconds for a new Task. Default is 86,400 seconds (24 hours). Optional when &#x60;integrationType&#x60; is &#x60;task&#x60;, not applicable otherwise.
 * @property { string } [integrationUrl] The URL of the external webhook. Required when &#x60;integrationType&#x60; is &#x60;external&#x60;.
 * @property { string } [integrationWorkflowSid] The Workflow SID for a new Task. Required when &#x60;integrationType&#x60; is &#x60;task&#x60;.
 * @property { string } [integrationWorkspaceSid] The Workspace SID for a new Task. Required when &#x60;integrationType&#x60; is &#x60;task&#x60;.
 * @property { string } [integrationType] The software that will handle inbound messages. [Integration Type](https://www.twilio.com/docs/flex/developer/messaging/manage-flows#integration-types) can be: &#x60;studio&#x60;, &#x60;external&#x60;, or &#x60;task&#x60;.
 * @property { boolean } [janitorEnabled] When enabled, the Messaging Channel Janitor will remove active Proxy sessions if the associated Task is deleted outside of the Flex UI. Defaults to &#x60;false&#x60;.
 * @property { boolean } [longLived] When enabled, Flex will keep the chat channel active so that it may be used for subsequent interactions with a contact identity. Defaults to &#x60;false&#x60;.
 */
export interface FlexFlowListInstanceCreateOptions {
    channelType: string;
    chatServiceSid: string;
    friendlyName: string;
    contactIdentity?: string;
    enabled?: boolean;
    integrationChannel?: string;
    integrationCreationOnMessage?: boolean;
    integrationFlowSid?: string;
    integrationPriority?: number;
    integrationRetryCount?: number;
    integrationTimeout?: number;
    integrationUrl?: string;
    integrationWorkflowSid?: string;
    integrationWorkspaceSid?: string;
    integrationType?: string;
    janitorEnabled?: boolean;
    longLived?: boolean;
}
/**
 * Options to pass to page a FlexFlowInstance
 *
 * @property { string } [friendlyName] The &#x60;friendly_name&#x60; of the Flex Flow resources to read.
 * @property { number } [pageSize] How many resources to return in each list page. The default is 50, and the maximum is 1000.
 */
export interface FlexFlowListInstancePageOptions {
    friendlyName?: string;
    pageSize?: number;
}

export interface FlexFlowListInstance {


    /**
     * Create a FlexFlowInstance
     *
     * @param { FlexFlowListInstanceCreateOptions } params - Parameter for request
     * @param { function } [callback] - Callback to handle processed record
     *
     * @returns { Promise } Resolves to processed FlexFlowInstance
     */
    create(params: FlexFlowListInstanceCreateOptions, callback?: (error: Error | null, item?: FlexFlowInstance) => any): Promise<FlexFlowInstance>;
    create(params: any, callback?: any): Promise<FlexFlowInstance>
;
    /**
     * Page a FlexFlowInstance
     *
     * @param { function } [callback] - Callback to handle processed record
     *
     * @returns { Promise } Resolves to processed FlexFlowInstance
     */
    page(callback?: (error: Error | null, item?: FlexFlowInstance) => any): Promise<FlexFlowInstance>;
    /**
     * Page a FlexFlowInstance
     *
     * @param { FlexFlowListInstancePageOptions } params - Parameter for request
     * @param { function } [callback] - Callback to handle processed record
     *
     * @returns { Promise } Resolves to processed FlexFlowInstance
     */
    page(params: FlexFlowListInstancePageOptions, callback?: (error: Error | null, item?: FlexFlowInstance) => any): Promise<FlexFlowInstance>;
    page(params?: any, callback?: any): Promise<FlexFlowInstance>
;
    /**
     * Provide a user-friendly representation
     */
    toJSON(): any;
    [inspect.custom](_depth: any, options: InspectOptions): any;
}


interface FlexFlowListInstanceImpl extends FlexFlowListInstance {}
class FlexFlowListInstanceImpl implements FlexFlowListInstance {
    _version?: V2010;
    _solution?: any;
    _uri?: string;

}

export function FlexFlowListInstance(version: V2010): FlexFlowListInstance {
    const instance = {} as FlexFlowListInstanceImpl;

    instance._version = version;
    instance._solution = {  };
    instance._uri = `/v1/Credentials/AWS/FlexFlows`;

    instance.create = function create(params: any, callback?: any): Promise<FlexFlowInstance> {
        if (params === null || params === undefined) {
            throw new Error('Required parameter "params" missing.');
        }

        if (params.channelType === null || params.channelType === undefined) {
            throw new Error('Required parameter "params.channelType" missing.');
        }

        if (params.chatServiceSid === null || params.chatServiceSid === undefined) {
            throw new Error('Required parameter "params.chatServiceSid" missing.');
        }

        if (params.friendlyName === null || params.friendlyName === undefined) {
            throw new Error('Required parameter "params.friendlyName" missing.');
        }

        const data: any = {};

        data['ChannelType'] = params.channelType;
        data['ChatServiceSid'] = params.chatServiceSid;
        if (params.contactIdentity !== undefined) data['ContactIdentity'] = params.contactIdentity;
        if (params.enabled !== undefined) data['Enabled'] = params.enabled;
        data['FriendlyName'] = params.friendlyName;
        if (params.integrationChannel !== undefined) data['Integration.Channel'] = params.integrationChannel;
        if (params.integrationCreationOnMessage !== undefined) data['Integration.CreationOnMessage'] = params.integrationCreationOnMessage;
        if (params.integrationFlowSid !== undefined) data['Integration.FlowSid'] = params.integrationFlowSid;
        if (params.integrationPriority !== undefined) data['Integration.Priority'] = params.integrationPriority;
        if (params.integrationRetryCount !== undefined) data['Integration.RetryCount'] = params.integrationRetryCount;
        if (params.integrationTimeout !== undefined) data['Integration.Timeout'] = params.integrationTimeout;
        if (params.integrationUrl !== undefined) data['Integration.Url'] = params.integrationUrl;
        if (params.integrationWorkflowSid !== undefined) data['Integration.WorkflowSid'] = params.integrationWorkflowSid;
        if (params.integrationWorkspaceSid !== undefined) data['Integration.WorkspaceSid'] = params.integrationWorkspaceSid;
        if (params.integrationType !== undefined) data['IntegrationType'] = params.integrationType;
        if (params.janitorEnabled !== undefined) data['JanitorEnabled'] = params.janitorEnabled;
        if (params.longLived !== undefined) data['LongLived'] = params.longLived;

        const headers: any = {};
        headers['Content-Type'] = 'application/x-www-form-urlencoded'


        let operationVersion = version,
            operationPromise = operationVersion.create({ uri: this._uri, method: 'POST', data, headers });

        operationPromise = operationPromise.then(payload => new FlexFlowInstance(operationVersion, payload));

        if (typeof callback === 'function') {
            operationPromise = operationPromise
                .then(value => callback(null, value))
                .catch(error => callback(error));
        }

        return operationPromise;

    }

    instance.page = function page(params?: any, callback?: any): Promise<FlexFlowInstance> {
        if (typeof params === 'function') {
            callback = params;
            params = {};
        } else {
            params = params || {};
        }

        const data: any = {};

        if (params.friendlyName !== undefined) data['FriendlyName'] = params.friendlyName;
        if (params.pageSize !== undefined) data['PageSize'] = params.pageSize;

        const headers: any = {};


        let operationVersion = version,
            operationPromise = operationVersion.page({ uri: this._uri, method: 'GET', data, headers });

        operationPromise = operationPromise.then(payload => new FlexFlowInstance(operationVersion, payload));

        if (typeof callback === 'function') {
            operationPromise = operationPromise
                .then(value => callback(null, value))
                .catch(error => callback(error));
        }

        return operationPromise;

    }

    instance.toJSON = function toJSON() {
        return this._solution;
    }

    instance[inspect.custom] = function inspectImpl(_depth: any, options: InspectOptions) {
        return inspect(this.toJSON(), options);
    }

    return instance;
}
