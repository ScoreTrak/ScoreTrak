# EntSchemaApi.ServiceApi

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**createService**](ServiceApi.md#createService) | **POST** /services | Create a new Service
[**deleteService**](ServiceApi.md#deleteService) | **DELETE** /services/{id} | Deletes a Service by ID
[**listService**](ServiceApi.md#listService) | **GET** /services | List Services
[**listServiceChecks**](ServiceApi.md#listServiceChecks) | **GET** /services/{id}/checks | List attached Checks
[**listServiceProperties**](ServiceApi.md#listServiceProperties) | **GET** /services/{id}/properties | List attached Properties
[**readService**](ServiceApi.md#readService) | **GET** /services/{id} | Find a Service by ID
[**readServiceCompetition**](ServiceApi.md#readServiceCompetition) | **GET** /services/{id}/competition | Find the attached Competition
[**readServiceHosts**](ServiceApi.md#readServiceHosts) | **GET** /services/{id}/hosts | Find the attached Host
[**readServiceTeam**](ServiceApi.md#readServiceTeam) | **GET** /services/{id}/team | Find the attached Team
[**updateService**](ServiceApi.md#updateService) | **PATCH** /services/{id} | Updates a Service

<a name="createService"></a>
# **createService**
> ServiceCreate createService(body)

Create a new Service

Creates a new Service and persists it to storage.

### Example
```javascript
import {EntSchemaApi} from 'ent_schema_api';

let apiInstance = new EntSchemaApi.ServiceApi();
let body = new EntSchemaApi.ServicesBody(); // ServicesBody | Service to create

apiInstance.createService(body, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ServicesBody**](ServicesBody.md)| Service to create | 

### Return type

[**ServiceCreate**](ServiceCreate.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="deleteService"></a>
# **deleteService**
> deleteService(id)

Deletes a Service by ID

Deletes the Service with the requested ID.

### Example
```javascript
import {EntSchemaApi} from 'ent_schema_api';

let apiInstance = new EntSchemaApi.ServiceApi();
let id = 56; // Number | ID of the Service

apiInstance.deleteService(id, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully.');
  }
});
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Number**| ID of the Service | 

### Return type

null (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

<a name="listService"></a>
# **listService**
> [ServiceList] listService(opts)

List Services

List Services.

### Example
```javascript
import {EntSchemaApi} from 'ent_schema_api';

let apiInstance = new EntSchemaApi.ServiceApi();
let opts = { 
  'page': 56, // Number | what page to render
  'itemsPerPage': 56 // Number | item count to render per page
};
apiInstance.listService(opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **Number**| what page to render | [optional] 
 **itemsPerPage** | **Number**| item count to render per page | [optional] 

### Return type

[**[ServiceList]**](ServiceList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

<a name="listServiceChecks"></a>
# **listServiceChecks**
> [ServiceChecksList] listServiceChecks(id, opts)

List attached Checks

List attached Checks.

### Example
```javascript
import {EntSchemaApi} from 'ent_schema_api';

let apiInstance = new EntSchemaApi.ServiceApi();
let id = 56; // Number | ID of the Service
let opts = { 
  'page': 56, // Number | what page to render
  'itemsPerPage': 56 // Number | item count to render per page
};
apiInstance.listServiceChecks(id, opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Number**| ID of the Service | 
 **page** | **Number**| what page to render | [optional] 
 **itemsPerPage** | **Number**| item count to render per page | [optional] 

### Return type

[**[ServiceChecksList]**](ServiceChecksList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

<a name="listServiceProperties"></a>
# **listServiceProperties**
> [ServicePropertiesList] listServiceProperties(id, opts)

List attached Properties

List attached Properties.

### Example
```javascript
import {EntSchemaApi} from 'ent_schema_api';

let apiInstance = new EntSchemaApi.ServiceApi();
let id = 56; // Number | ID of the Service
let opts = { 
  'page': 56, // Number | what page to render
  'itemsPerPage': 56 // Number | item count to render per page
};
apiInstance.listServiceProperties(id, opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Number**| ID of the Service | 
 **page** | **Number**| what page to render | [optional] 
 **itemsPerPage** | **Number**| item count to render per page | [optional] 

### Return type

[**[ServicePropertiesList]**](ServicePropertiesList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

<a name="readService"></a>
# **readService**
> ServiceRead readService(id)

Find a Service by ID

Finds the Service with the requested ID and returns it.

### Example
```javascript
import {EntSchemaApi} from 'ent_schema_api';

let apiInstance = new EntSchemaApi.ServiceApi();
let id = 56; // Number | ID of the Service

apiInstance.readService(id, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Number**| ID of the Service | 

### Return type

[**ServiceRead**](ServiceRead.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

<a name="readServiceCompetition"></a>
# **readServiceCompetition**
> ServiceCompetitionRead readServiceCompetition(id)

Find the attached Competition

Find the attached Competition of the Service with the given ID

### Example
```javascript
import {EntSchemaApi} from 'ent_schema_api';

let apiInstance = new EntSchemaApi.ServiceApi();
let id = 56; // Number | ID of the Service

apiInstance.readServiceCompetition(id, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Number**| ID of the Service | 

### Return type

[**ServiceCompetitionRead**](ServiceCompetitionRead.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

<a name="readServiceHosts"></a>
# **readServiceHosts**
> ServiceHostsRead readServiceHosts(id)

Find the attached Host

Find the attached Host of the Service with the given ID

### Example
```javascript
import {EntSchemaApi} from 'ent_schema_api';

let apiInstance = new EntSchemaApi.ServiceApi();
let id = 56; // Number | ID of the Service

apiInstance.readServiceHosts(id, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Number**| ID of the Service | 

### Return type

[**ServiceHostsRead**](ServiceHostsRead.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

<a name="readServiceTeam"></a>
# **readServiceTeam**
> ServiceTeamRead readServiceTeam(id)

Find the attached Team

Find the attached Team of the Service with the given ID

### Example
```javascript
import {EntSchemaApi} from 'ent_schema_api';

let apiInstance = new EntSchemaApi.ServiceApi();
let id = 56; // Number | ID of the Service

apiInstance.readServiceTeam(id, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Number**| ID of the Service | 

### Return type

[**ServiceTeamRead**](ServiceTeamRead.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

<a name="updateService"></a>
# **updateService**
> ServiceUpdate updateService(body, id)

Updates a Service

Updates a Service and persists changes to storage.

### Example
```javascript
import {EntSchemaApi} from 'ent_schema_api';

let apiInstance = new EntSchemaApi.ServiceApi();
let body = new EntSchemaApi.ServicesIdBody(); // ServicesIdBody | Service properties to update
let id = 56; // Number | ID of the Service

apiInstance.updateService(body, id, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ServicesIdBody**](ServicesIdBody.md)| Service properties to update | 
 **id** | **Number**| ID of the Service | 

### Return type

[**ServiceUpdate**](ServiceUpdate.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

