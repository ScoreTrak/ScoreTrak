# EntSchemaApi.CheckApi

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**createCheck**](CheckApi.md#createCheck) | **POST** /checks | Create a new Check
[**deleteCheck**](CheckApi.md#deleteCheck) | **DELETE** /checks/{id} | Deletes a Check by ID
[**listCheck**](CheckApi.md#listCheck) | **GET** /checks | List Checks
[**readCheck**](CheckApi.md#readCheck) | **GET** /checks/{id} | Find a Check by ID
[**readCheckCompetition**](CheckApi.md#readCheckCompetition) | **GET** /checks/{id}/competition | Find the attached Competition
[**readCheckRounds**](CheckApi.md#readCheckRounds) | **GET** /checks/{id}/rounds | Find the attached Round
[**readCheckServices**](CheckApi.md#readCheckServices) | **GET** /checks/{id}/services | Find the attached Service
[**updateCheck**](CheckApi.md#updateCheck) | **PATCH** /checks/{id} | Updates a Check

<a name="createCheck"></a>
# **createCheck**
> CheckCreate createCheck(body)

Create a new Check

Creates a new Check and persists it to storage.

### Example
```javascript
import {EntSchemaApi} from 'ent_schema_api';

let apiInstance = new EntSchemaApi.CheckApi();
let body = new EntSchemaApi.ChecksBody(); // ChecksBody | Check to create

apiInstance.createCheck(body, (error, data, response) => {
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
 **body** | [**ChecksBody**](ChecksBody.md)| Check to create | 

### Return type

[**CheckCreate**](CheckCreate.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="deleteCheck"></a>
# **deleteCheck**
> deleteCheck(id)

Deletes a Check by ID

Deletes the Check with the requested ID.

### Example
```javascript
import {EntSchemaApi} from 'ent_schema_api';

let apiInstance = new EntSchemaApi.CheckApi();
let id = 56; // Number | ID of the Check

apiInstance.deleteCheck(id, (error, data, response) => {
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
 **id** | **Number**| ID of the Check | 

### Return type

null (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

<a name="listCheck"></a>
# **listCheck**
> [CheckList] listCheck(opts)

List Checks

List Checks.

### Example
```javascript
import {EntSchemaApi} from 'ent_schema_api';

let apiInstance = new EntSchemaApi.CheckApi();
let opts = { 
  'page': 56, // Number | what page to render
  'itemsPerPage': 56 // Number | item count to render per page
};
apiInstance.listCheck(opts, (error, data, response) => {
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

[**[CheckList]**](CheckList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

<a name="readCheck"></a>
# **readCheck**
> CheckRead readCheck(id)

Find a Check by ID

Finds the Check with the requested ID and returns it.

### Example
```javascript
import {EntSchemaApi} from 'ent_schema_api';

let apiInstance = new EntSchemaApi.CheckApi();
let id = 56; // Number | ID of the Check

apiInstance.readCheck(id, (error, data, response) => {
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
 **id** | **Number**| ID of the Check | 

### Return type

[**CheckRead**](CheckRead.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

<a name="readCheckCompetition"></a>
# **readCheckCompetition**
> CheckCompetitionRead readCheckCompetition(id)

Find the attached Competition

Find the attached Competition of the Check with the given ID

### Example
```javascript
import {EntSchemaApi} from 'ent_schema_api';

let apiInstance = new EntSchemaApi.CheckApi();
let id = 56; // Number | ID of the Check

apiInstance.readCheckCompetition(id, (error, data, response) => {
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
 **id** | **Number**| ID of the Check | 

### Return type

[**CheckCompetitionRead**](CheckCompetitionRead.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

<a name="readCheckRounds"></a>
# **readCheckRounds**
> CheckRoundsRead readCheckRounds(id)

Find the attached Round

Find the attached Round of the Check with the given ID

### Example
```javascript
import {EntSchemaApi} from 'ent_schema_api';

let apiInstance = new EntSchemaApi.CheckApi();
let id = 56; // Number | ID of the Check

apiInstance.readCheckRounds(id, (error, data, response) => {
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
 **id** | **Number**| ID of the Check | 

### Return type

[**CheckRoundsRead**](CheckRoundsRead.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

<a name="readCheckServices"></a>
# **readCheckServices**
> CheckServicesRead readCheckServices(id)

Find the attached Service

Find the attached Service of the Check with the given ID

### Example
```javascript
import {EntSchemaApi} from 'ent_schema_api';

let apiInstance = new EntSchemaApi.CheckApi();
let id = 56; // Number | ID of the Check

apiInstance.readCheckServices(id, (error, data, response) => {
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
 **id** | **Number**| ID of the Check | 

### Return type

[**CheckServicesRead**](CheckServicesRead.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

<a name="updateCheck"></a>
# **updateCheck**
> CheckUpdate updateCheck(body, id)

Updates a Check

Updates a Check and persists changes to storage.

### Example
```javascript
import {EntSchemaApi} from 'ent_schema_api';

let apiInstance = new EntSchemaApi.CheckApi();
let body = new EntSchemaApi.ChecksIdBody(); // ChecksIdBody | Check properties to update
let id = 56; // Number | ID of the Check

apiInstance.updateCheck(body, id, (error, data, response) => {
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
 **body** | [**ChecksIdBody**](ChecksIdBody.md)| Check properties to update | 
 **id** | **Number**| ID of the Check | 

### Return type

[**CheckUpdate**](CheckUpdate.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

