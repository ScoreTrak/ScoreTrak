# EntSchemaApi.PropertyApi

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**createProperty**](PropertyApi.md#createProperty) | **POST** /properties | Create a new Property
[**deleteProperty**](PropertyApi.md#deleteProperty) | **DELETE** /properties/{id} | Deletes a Property by ID
[**listProperty**](PropertyApi.md#listProperty) | **GET** /properties | List Properties
[**readProperty**](PropertyApi.md#readProperty) | **GET** /properties/{id} | Find a Property by ID
[**readPropertyCompetition**](PropertyApi.md#readPropertyCompetition) | **GET** /properties/{id}/competition | Find the attached Competition
[**readPropertyServices**](PropertyApi.md#readPropertyServices) | **GET** /properties/{id}/services | Find the attached Service
[**readPropertyTeam**](PropertyApi.md#readPropertyTeam) | **GET** /properties/{id}/team | Find the attached Team
[**updateProperty**](PropertyApi.md#updateProperty) | **PATCH** /properties/{id} | Updates a Property

<a name="createProperty"></a>
# **createProperty**
> PropertyCreate createProperty(body)

Create a new Property

Creates a new Property and persists it to storage.

### Example
```javascript
import {EntSchemaApi} from 'ent_schema_api';

let apiInstance = new EntSchemaApi.PropertyApi();
let body = new EntSchemaApi.PropertiesBody(); // PropertiesBody | Property to create

apiInstance.createProperty(body, (error, data, response) => {
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
 **body** | [**PropertiesBody**](PropertiesBody.md)| Property to create | 

### Return type

[**PropertyCreate**](PropertyCreate.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="deleteProperty"></a>
# **deleteProperty**
> deleteProperty(id)

Deletes a Property by ID

Deletes the Property with the requested ID.

### Example
```javascript
import {EntSchemaApi} from 'ent_schema_api';

let apiInstance = new EntSchemaApi.PropertyApi();
let id = 56; // Number | ID of the Property

apiInstance.deleteProperty(id, (error, data, response) => {
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
 **id** | **Number**| ID of the Property | 

### Return type

null (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

<a name="listProperty"></a>
# **listProperty**
> [PropertyList] listProperty(opts)

List Properties

List Properties.

### Example
```javascript
import {EntSchemaApi} from 'ent_schema_api';

let apiInstance = new EntSchemaApi.PropertyApi();
let opts = { 
  'page': 56, // Number | what page to render
  'itemsPerPage': 56 // Number | item count to render per page
};
apiInstance.listProperty(opts, (error, data, response) => {
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

[**[PropertyList]**](PropertyList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

<a name="readProperty"></a>
# **readProperty**
> PropertyRead readProperty(id)

Find a Property by ID

Finds the Property with the requested ID and returns it.

### Example
```javascript
import {EntSchemaApi} from 'ent_schema_api';

let apiInstance = new EntSchemaApi.PropertyApi();
let id = 56; // Number | ID of the Property

apiInstance.readProperty(id, (error, data, response) => {
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
 **id** | **Number**| ID of the Property | 

### Return type

[**PropertyRead**](PropertyRead.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

<a name="readPropertyCompetition"></a>
# **readPropertyCompetition**
> PropertyCompetitionRead readPropertyCompetition(id)

Find the attached Competition

Find the attached Competition of the Property with the given ID

### Example
```javascript
import {EntSchemaApi} from 'ent_schema_api';

let apiInstance = new EntSchemaApi.PropertyApi();
let id = 56; // Number | ID of the Property

apiInstance.readPropertyCompetition(id, (error, data, response) => {
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
 **id** | **Number**| ID of the Property | 

### Return type

[**PropertyCompetitionRead**](PropertyCompetitionRead.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

<a name="readPropertyServices"></a>
# **readPropertyServices**
> PropertyServicesRead readPropertyServices(id)

Find the attached Service

Find the attached Service of the Property with the given ID

### Example
```javascript
import {EntSchemaApi} from 'ent_schema_api';

let apiInstance = new EntSchemaApi.PropertyApi();
let id = 56; // Number | ID of the Property

apiInstance.readPropertyServices(id, (error, data, response) => {
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
 **id** | **Number**| ID of the Property | 

### Return type

[**PropertyServicesRead**](PropertyServicesRead.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

<a name="readPropertyTeam"></a>
# **readPropertyTeam**
> PropertyTeamRead readPropertyTeam(id)

Find the attached Team

Find the attached Team of the Property with the given ID

### Example
```javascript
import {EntSchemaApi} from 'ent_schema_api';

let apiInstance = new EntSchemaApi.PropertyApi();
let id = 56; // Number | ID of the Property

apiInstance.readPropertyTeam(id, (error, data, response) => {
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
 **id** | **Number**| ID of the Property | 

### Return type

[**PropertyTeamRead**](PropertyTeamRead.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

<a name="updateProperty"></a>
# **updateProperty**
> PropertyUpdate updateProperty(body, id)

Updates a Property

Updates a Property and persists changes to storage.

### Example
```javascript
import {EntSchemaApi} from 'ent_schema_api';

let apiInstance = new EntSchemaApi.PropertyApi();
let body = new EntSchemaApi.PropertiesIdBody(); // PropertiesIdBody | Property properties to update
let id = 56; // Number | ID of the Property

apiInstance.updateProperty(body, id, (error, data, response) => {
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
 **body** | [**PropertiesIdBody**](PropertiesIdBody.md)| Property properties to update | 
 **id** | **Number**| ID of the Property | 

### Return type

[**PropertyUpdate**](PropertyUpdate.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

