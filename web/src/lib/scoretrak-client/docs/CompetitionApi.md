# EntSchemaApi.CompetitionApi

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**createCompetition**](CompetitionApi.md#createCompetition) | **POST** /competitions | Create a new Competition
[**deleteCompetition**](CompetitionApi.md#deleteCompetition) | **DELETE** /competitions/{id} | Deletes a Competition by ID
[**listCompetition**](CompetitionApi.md#listCompetition) | **GET** /competitions | List Competitions
[**listCompetitionTeams**](CompetitionApi.md#listCompetitionTeams) | **GET** /competitions/{id}/teams | List attached Teams
[**listCompetitionUsers**](CompetitionApi.md#listCompetitionUsers) | **GET** /competitions/{id}/users | List attached Users
[**readCompetition**](CompetitionApi.md#readCompetition) | **GET** /competitions/{id} | Find a Competition by ID
[**updateCompetition**](CompetitionApi.md#updateCompetition) | **PATCH** /competitions/{id} | Updates a Competition

<a name="createCompetition"></a>
# **createCompetition**
> CompetitionCreate createCompetition(body)

Create a new Competition

Creates a new Competition and persists it to storage.

### Example
```javascript
import {EntSchemaApi} from 'ent_schema_api';

let apiInstance = new EntSchemaApi.CompetitionApi();
let body = new EntSchemaApi.CompetitionsBody(); // CompetitionsBody | Competition to create

apiInstance.createCompetition(body, (error, data, response) => {
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
 **body** | [**CompetitionsBody**](CompetitionsBody.md)| Competition to create | 

### Return type

[**CompetitionCreate**](CompetitionCreate.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="deleteCompetition"></a>
# **deleteCompetition**
> deleteCompetition(id)

Deletes a Competition by ID

Deletes the Competition with the requested ID.

### Example
```javascript
import {EntSchemaApi} from 'ent_schema_api';

let apiInstance = new EntSchemaApi.CompetitionApi();
let id = 56; // Number | ID of the Competition

apiInstance.deleteCompetition(id, (error, data, response) => {
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
 **id** | **Number**| ID of the Competition | 

### Return type

null (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

<a name="listCompetition"></a>
# **listCompetition**
> [CompetitionList] listCompetition(opts)

List Competitions

List Competitions.

### Example
```javascript
import {EntSchemaApi} from 'ent_schema_api';

let apiInstance = new EntSchemaApi.CompetitionApi();
let opts = { 
  'page': 56, // Number | what page to render
  'itemsPerPage': 56 // Number | item count to render per page
};
apiInstance.listCompetition(opts, (error, data, response) => {
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

[**[CompetitionList]**](CompetitionList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

<a name="listCompetitionTeams"></a>
# **listCompetitionTeams**
> [CompetitionTeamsList] listCompetitionTeams(id, opts)

List attached Teams

List attached Teams.

### Example
```javascript
import {EntSchemaApi} from 'ent_schema_api';

let apiInstance = new EntSchemaApi.CompetitionApi();
let id = 56; // Number | ID of the Competition
let opts = { 
  'page': 56, // Number | what page to render
  'itemsPerPage': 56 // Number | item count to render per page
};
apiInstance.listCompetitionTeams(id, opts, (error, data, response) => {
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
 **id** | **Number**| ID of the Competition | 
 **page** | **Number**| what page to render | [optional] 
 **itemsPerPage** | **Number**| item count to render per page | [optional] 

### Return type

[**[CompetitionTeamsList]**](CompetitionTeamsList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

<a name="listCompetitionUsers"></a>
# **listCompetitionUsers**
> [CompetitionUsersList] listCompetitionUsers(id, opts)

List attached Users

List attached Users.

### Example
```javascript
import {EntSchemaApi} from 'ent_schema_api';

let apiInstance = new EntSchemaApi.CompetitionApi();
let id = 56; // Number | ID of the Competition
let opts = { 
  'page': 56, // Number | what page to render
  'itemsPerPage': 56 // Number | item count to render per page
};
apiInstance.listCompetitionUsers(id, opts, (error, data, response) => {
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
 **id** | **Number**| ID of the Competition | 
 **page** | **Number**| what page to render | [optional] 
 **itemsPerPage** | **Number**| item count to render per page | [optional] 

### Return type

[**[CompetitionUsersList]**](CompetitionUsersList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

<a name="readCompetition"></a>
# **readCompetition**
> CompetitionRead readCompetition(id)

Find a Competition by ID

Finds the Competition with the requested ID and returns it.

### Example
```javascript
import {EntSchemaApi} from 'ent_schema_api';

let apiInstance = new EntSchemaApi.CompetitionApi();
let id = 56; // Number | ID of the Competition

apiInstance.readCompetition(id, (error, data, response) => {
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
 **id** | **Number**| ID of the Competition | 

### Return type

[**CompetitionRead**](CompetitionRead.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

<a name="updateCompetition"></a>
# **updateCompetition**
> CompetitionUpdate updateCompetition(body, id)

Updates a Competition

Updates a Competition and persists changes to storage.

### Example
```javascript
import {EntSchemaApi} from 'ent_schema_api';

let apiInstance = new EntSchemaApi.CompetitionApi();
let body = new EntSchemaApi.CompetitionsIdBody(); // CompetitionsIdBody | Competition properties to update
let id = 56; // Number | ID of the Competition

apiInstance.updateCompetition(body, id, (error, data, response) => {
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
 **body** | [**CompetitionsIdBody**](CompetitionsIdBody.md)| Competition properties to update | 
 **id** | **Number**| ID of the Competition | 

### Return type

[**CompetitionUpdate**](CompetitionUpdate.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

