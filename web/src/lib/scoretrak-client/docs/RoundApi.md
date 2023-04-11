# EntSchemaApi.RoundApi

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**createRound**](RoundApi.md#createRound) | **POST** /rounds | Create a new Round
[**deleteRound**](RoundApi.md#deleteRound) | **DELETE** /rounds/{id} | Deletes a Round by ID
[**listRound**](RoundApi.md#listRound) | **GET** /rounds | List Rounds
[**listRoundChecks**](RoundApi.md#listRoundChecks) | **GET** /rounds/{id}/checks | List attached Checks
[**readRound**](RoundApi.md#readRound) | **GET** /rounds/{id} | Find a Round by ID
[**readRoundCompetition**](RoundApi.md#readRoundCompetition) | **GET** /rounds/{id}/competition | Find the attached Competition
[**updateRound**](RoundApi.md#updateRound) | **PATCH** /rounds/{id} | Updates a Round

<a name="createRound"></a>
# **createRound**
> RoundCreate createRound(body)

Create a new Round

Creates a new Round and persists it to storage.

### Example
```javascript
import {EntSchemaApi} from 'ent_schema_api';

let apiInstance = new EntSchemaApi.RoundApi();
let body = new EntSchemaApi.RoundsBody(); // RoundsBody | Round to create

apiInstance.createRound(body, (error, data, response) => {
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
 **body** | [**RoundsBody**](RoundsBody.md)| Round to create | 

### Return type

[**RoundCreate**](RoundCreate.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="deleteRound"></a>
# **deleteRound**
> deleteRound(id)

Deletes a Round by ID

Deletes the Round with the requested ID.

### Example
```javascript
import {EntSchemaApi} from 'ent_schema_api';

let apiInstance = new EntSchemaApi.RoundApi();
let id = 56; // Number | ID of the Round

apiInstance.deleteRound(id, (error, data, response) => {
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
 **id** | **Number**| ID of the Round | 

### Return type

null (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

<a name="listRound"></a>
# **listRound**
> [RoundList] listRound(opts)

List Rounds

List Rounds.

### Example
```javascript
import {EntSchemaApi} from 'ent_schema_api';

let apiInstance = new EntSchemaApi.RoundApi();
let opts = { 
  'page': 56, // Number | what page to render
  'itemsPerPage': 56 // Number | item count to render per page
};
apiInstance.listRound(opts, (error, data, response) => {
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

[**[RoundList]**](RoundList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

<a name="listRoundChecks"></a>
# **listRoundChecks**
> [RoundChecksList] listRoundChecks(id, opts)

List attached Checks

List attached Checks.

### Example
```javascript
import {EntSchemaApi} from 'ent_schema_api';

let apiInstance = new EntSchemaApi.RoundApi();
let id = 56; // Number | ID of the Round
let opts = { 
  'page': 56, // Number | what page to render
  'itemsPerPage': 56 // Number | item count to render per page
};
apiInstance.listRoundChecks(id, opts, (error, data, response) => {
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
 **id** | **Number**| ID of the Round | 
 **page** | **Number**| what page to render | [optional] 
 **itemsPerPage** | **Number**| item count to render per page | [optional] 

### Return type

[**[RoundChecksList]**](RoundChecksList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

<a name="readRound"></a>
# **readRound**
> RoundRead readRound(id)

Find a Round by ID

Finds the Round with the requested ID and returns it.

### Example
```javascript
import {EntSchemaApi} from 'ent_schema_api';

let apiInstance = new EntSchemaApi.RoundApi();
let id = 56; // Number | ID of the Round

apiInstance.readRound(id, (error, data, response) => {
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
 **id** | **Number**| ID of the Round | 

### Return type

[**RoundRead**](RoundRead.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

<a name="readRoundCompetition"></a>
# **readRoundCompetition**
> RoundCompetitionRead readRoundCompetition(id)

Find the attached Competition

Find the attached Competition of the Round with the given ID

### Example
```javascript
import {EntSchemaApi} from 'ent_schema_api';

let apiInstance = new EntSchemaApi.RoundApi();
let id = 56; // Number | ID of the Round

apiInstance.readRoundCompetition(id, (error, data, response) => {
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
 **id** | **Number**| ID of the Round | 

### Return type

[**RoundCompetitionRead**](RoundCompetitionRead.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

<a name="updateRound"></a>
# **updateRound**
> RoundUpdate updateRound(body, id)

Updates a Round

Updates a Round and persists changes to storage.

### Example
```javascript
import {EntSchemaApi} from 'ent_schema_api';

let apiInstance = new EntSchemaApi.RoundApi();
let body = new EntSchemaApi.RoundsIdBody(); // RoundsIdBody | Round properties to update
let id = 56; // Number | ID of the Round

apiInstance.updateRound(body, id, (error, data, response) => {
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
 **body** | [**RoundsIdBody**](RoundsIdBody.md)| Round properties to update | 
 **id** | **Number**| ID of the Round | 

### Return type

[**RoundUpdate**](RoundUpdate.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

