/*
 * Ent Schema API
 * This is an auto generated API description made out of an Ent schema definition
 *
 * OpenAPI spec version: 0.1.0
 *
 * NOTE: This class is auto generated by the swagger code generator program.
 * https://github.com/swagger-api/swagger-codegen.git
 *
 * Swagger Codegen version: 3.0.33
 *
 * Do not edit the class manually.
 *
 */
import {ApiClient} from "../ApiClient";
import {InlineResponse400} from '../model/InlineResponse400';
import {ServiceChecksList} from '../model/ServiceChecksList';
import {ServiceCompetitionRead} from '../model/ServiceCompetitionRead';
import {ServiceCreate} from '../model/ServiceCreate';
import {ServiceHostsRead} from '../model/ServiceHostsRead';
import {ServiceList} from '../model/ServiceList';
import {ServicePropertiesList} from '../model/ServicePropertiesList';
import {ServiceRead} from '../model/ServiceRead';
import {ServiceTeamRead} from '../model/ServiceTeamRead';
import {ServiceUpdate} from '../model/ServiceUpdate';
import {ServicesBody} from '../model/ServicesBody';
import {ServicesIdBody} from '../model/ServicesIdBody';

/**
* Service service.
* @module api/ServiceApi
* @version 0.1.0
*/
export class ServiceApi {

    /**
    * Constructs a new ServiceApi. 
    * @alias module:api/ServiceApi
    * @class
    * @param {module:ApiClient} [apiClient] Optional API client implementation to use,
    * default to {@link module:ApiClient#instanc
    e} if unspecified.
    */
    constructor(apiClient) {
        this.apiClient = apiClient || ApiClient.instance;
    }

    /**
     * Callback function to receive the result of the createService operation.
     * @callback moduleapi/ServiceApi~createServiceCallback
     * @param {String} error Error message, if any.
     * @param {module:model/ServiceCreate{ data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Create a new Service
     * Creates a new Service and persists it to storage.
     * @param {module:model/ServicesBody} body Service to create
     * @param {module:api/ServiceApi~createServiceCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link <&vendorExtensions.x-jsdoc-type>}
     */
    createService(body, callback) {
      
      let postBody = body;
      // verify the required parameter 'body' is set
      if (body === undefined || body === null) {
        throw new Error("Missing the required parameter 'body' when calling createService");
      }

      let pathParams = {
        
      };
      let queryParams = {
        
      };
      let headerParams = {
        
      };
      let formParams = {
        
      };

      let authNames = [];
      let contentTypes = ['application/json'];
      let accepts = ['application/json'];
      let returnType = ServiceCreate;

      return this.apiClient.callApi(
        '/services', 'POST',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, callback
      );
    }
    /**
     * Callback function to receive the result of the deleteService operation.
     * @callback moduleapi/ServiceApi~deleteServiceCallback
     * @param {String} error Error message, if any.
     * @param data This operation does not return a value.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Deletes a Service by ID
     * Deletes the Service with the requested ID.
     * @param {Number} id ID of the Service
     * @param {module:api/ServiceApi~deleteServiceCallback} callback The callback function, accepting three arguments: error, data, response
     */
    deleteService(id, callback) {
      
      let postBody = null;
      // verify the required parameter 'id' is set
      if (id === undefined || id === null) {
        throw new Error("Missing the required parameter 'id' when calling deleteService");
      }

      let pathParams = {
        'id': id
      };
      let queryParams = {
        
      };
      let headerParams = {
        
      };
      let formParams = {
        
      };

      let authNames = [];
      let contentTypes = [];
      let accepts = ['application/json'];
      let returnType = null;

      return this.apiClient.callApi(
        '/services/{id}', 'DELETE',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, callback
      );
    }
    /**
     * Callback function to receive the result of the listService operation.
     * @callback moduleapi/ServiceApi~listServiceCallback
     * @param {String} error Error message, if any.
     * @param {Array.<module:model/ServiceList>{ data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * List Services
     * List Services.
     * @param {Object} opts Optional parameters
     * @param {Number} opts.page what page to render
     * @param {Number} opts.itemsPerPage item count to render per page
     * @param {module:api/ServiceApi~listServiceCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link <&vendorExtensions.x-jsdoc-type>}
     */
    listService(opts, callback) {
      opts = opts || {};
      let postBody = null;

      let pathParams = {
        
      };
      let queryParams = {
        'page': opts['page'],'itemsPerPage': opts['itemsPerPage']
      };
      let headerParams = {
        
      };
      let formParams = {
        
      };

      let authNames = [];
      let contentTypes = [];
      let accepts = ['application/json'];
      let returnType = [ServiceList];

      return this.apiClient.callApi(
        '/services', 'GET',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, callback
      );
    }
    /**
     * Callback function to receive the result of the listServiceChecks operation.
     * @callback moduleapi/ServiceApi~listServiceChecksCallback
     * @param {String} error Error message, if any.
     * @param {Array.<module:model/ServiceChecksList>{ data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * List attached Checks
     * List attached Checks.
     * @param {Number} id ID of the Service
     * @param {Object} opts Optional parameters
     * @param {Number} opts.page what page to render
     * @param {Number} opts.itemsPerPage item count to render per page
     * @param {module:api/ServiceApi~listServiceChecksCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link <&vendorExtensions.x-jsdoc-type>}
     */
    listServiceChecks(id, opts, callback) {
      opts = opts || {};
      let postBody = null;
      // verify the required parameter 'id' is set
      if (id === undefined || id === null) {
        throw new Error("Missing the required parameter 'id' when calling listServiceChecks");
      }

      let pathParams = {
        'id': id
      };
      let queryParams = {
        'page': opts['page'],'itemsPerPage': opts['itemsPerPage']
      };
      let headerParams = {
        
      };
      let formParams = {
        
      };

      let authNames = [];
      let contentTypes = [];
      let accepts = ['application/json'];
      let returnType = [ServiceChecksList];

      return this.apiClient.callApi(
        '/services/{id}/checks', 'GET',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, callback
      );
    }
    /**
     * Callback function to receive the result of the listServiceProperties operation.
     * @callback moduleapi/ServiceApi~listServicePropertiesCallback
     * @param {String} error Error message, if any.
     * @param {Array.<module:model/ServicePropertiesList>{ data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * List attached Properties
     * List attached Properties.
     * @param {Number} id ID of the Service
     * @param {Object} opts Optional parameters
     * @param {Number} opts.page what page to render
     * @param {Number} opts.itemsPerPage item count to render per page
     * @param {module:api/ServiceApi~listServicePropertiesCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link <&vendorExtensions.x-jsdoc-type>}
     */
    listServiceProperties(id, opts, callback) {
      opts = opts || {};
      let postBody = null;
      // verify the required parameter 'id' is set
      if (id === undefined || id === null) {
        throw new Error("Missing the required parameter 'id' when calling listServiceProperties");
      }

      let pathParams = {
        'id': id
      };
      let queryParams = {
        'page': opts['page'],'itemsPerPage': opts['itemsPerPage']
      };
      let headerParams = {
        
      };
      let formParams = {
        
      };

      let authNames = [];
      let contentTypes = [];
      let accepts = ['application/json'];
      let returnType = [ServicePropertiesList];

      return this.apiClient.callApi(
        '/services/{id}/properties', 'GET',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, callback
      );
    }
    /**
     * Callback function to receive the result of the readService operation.
     * @callback moduleapi/ServiceApi~readServiceCallback
     * @param {String} error Error message, if any.
     * @param {module:model/ServiceRead{ data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Find a Service by ID
     * Finds the Service with the requested ID and returns it.
     * @param {Number} id ID of the Service
     * @param {module:api/ServiceApi~readServiceCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link <&vendorExtensions.x-jsdoc-type>}
     */
    readService(id, callback) {
      
      let postBody = null;
      // verify the required parameter 'id' is set
      if (id === undefined || id === null) {
        throw new Error("Missing the required parameter 'id' when calling readService");
      }

      let pathParams = {
        'id': id
      };
      let queryParams = {
        
      };
      let headerParams = {
        
      };
      let formParams = {
        
      };

      let authNames = [];
      let contentTypes = [];
      let accepts = ['application/json'];
      let returnType = ServiceRead;

      return this.apiClient.callApi(
        '/services/{id}', 'GET',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, callback
      );
    }
    /**
     * Callback function to receive the result of the readServiceCompetition operation.
     * @callback moduleapi/ServiceApi~readServiceCompetitionCallback
     * @param {String} error Error message, if any.
     * @param {module:model/ServiceCompetitionRead{ data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Find the attached Competition
     * Find the attached Competition of the Service with the given ID
     * @param {Number} id ID of the Service
     * @param {module:api/ServiceApi~readServiceCompetitionCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link <&vendorExtensions.x-jsdoc-type>}
     */
    readServiceCompetition(id, callback) {
      
      let postBody = null;
      // verify the required parameter 'id' is set
      if (id === undefined || id === null) {
        throw new Error("Missing the required parameter 'id' when calling readServiceCompetition");
      }

      let pathParams = {
        'id': id
      };
      let queryParams = {
        
      };
      let headerParams = {
        
      };
      let formParams = {
        
      };

      let authNames = [];
      let contentTypes = [];
      let accepts = ['application/json'];
      let returnType = ServiceCompetitionRead;

      return this.apiClient.callApi(
        '/services/{id}/competition', 'GET',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, callback
      );
    }
    /**
     * Callback function to receive the result of the readServiceHosts operation.
     * @callback moduleapi/ServiceApi~readServiceHostsCallback
     * @param {String} error Error message, if any.
     * @param {module:model/ServiceHostsRead{ data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Find the attached Host
     * Find the attached Host of the Service with the given ID
     * @param {Number} id ID of the Service
     * @param {module:api/ServiceApi~readServiceHostsCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link <&vendorExtensions.x-jsdoc-type>}
     */
    readServiceHosts(id, callback) {
      
      let postBody = null;
      // verify the required parameter 'id' is set
      if (id === undefined || id === null) {
        throw new Error("Missing the required parameter 'id' when calling readServiceHosts");
      }

      let pathParams = {
        'id': id
      };
      let queryParams = {
        
      };
      let headerParams = {
        
      };
      let formParams = {
        
      };

      let authNames = [];
      let contentTypes = [];
      let accepts = ['application/json'];
      let returnType = ServiceHostsRead;

      return this.apiClient.callApi(
        '/services/{id}/hosts', 'GET',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, callback
      );
    }
    /**
     * Callback function to receive the result of the readServiceTeam operation.
     * @callback moduleapi/ServiceApi~readServiceTeamCallback
     * @param {String} error Error message, if any.
     * @param {module:model/ServiceTeamRead{ data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Find the attached Team
     * Find the attached Team of the Service with the given ID
     * @param {Number} id ID of the Service
     * @param {module:api/ServiceApi~readServiceTeamCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link <&vendorExtensions.x-jsdoc-type>}
     */
    readServiceTeam(id, callback) {
      
      let postBody = null;
      // verify the required parameter 'id' is set
      if (id === undefined || id === null) {
        throw new Error("Missing the required parameter 'id' when calling readServiceTeam");
      }

      let pathParams = {
        'id': id
      };
      let queryParams = {
        
      };
      let headerParams = {
        
      };
      let formParams = {
        
      };

      let authNames = [];
      let contentTypes = [];
      let accepts = ['application/json'];
      let returnType = ServiceTeamRead;

      return this.apiClient.callApi(
        '/services/{id}/team', 'GET',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, callback
      );
    }
    /**
     * Callback function to receive the result of the updateService operation.
     * @callback moduleapi/ServiceApi~updateServiceCallback
     * @param {String} error Error message, if any.
     * @param {module:model/ServiceUpdate{ data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Updates a Service
     * Updates a Service and persists changes to storage.
     * @param {module:model/ServicesIdBody} body Service properties to update
     * @param {Number} id ID of the Service
     * @param {module:api/ServiceApi~updateServiceCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link <&vendorExtensions.x-jsdoc-type>}
     */
    updateService(body, id, callback) {
      
      let postBody = body;
      // verify the required parameter 'body' is set
      if (body === undefined || body === null) {
        throw new Error("Missing the required parameter 'body' when calling updateService");
      }
      // verify the required parameter 'id' is set
      if (id === undefined || id === null) {
        throw new Error("Missing the required parameter 'id' when calling updateService");
      }

      let pathParams = {
        'id': id
      };
      let queryParams = {
        
      };
      let headerParams = {
        
      };
      let formParams = {
        
      };

      let authNames = [];
      let contentTypes = ['application/json'];
      let accepts = ['application/json'];
      let returnType = ServiceUpdate;

      return this.apiClient.callApi(
        '/services/{id}', 'PATCH',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, callback
      );
    }

}