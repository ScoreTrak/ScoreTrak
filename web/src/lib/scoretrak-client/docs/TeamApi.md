# EntSchemaApi.TeamApi

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**createTeam**](TeamApi.md#createTeam) | **POST** /teams | Create a new Team
[**deleteTeam**](TeamApi.md#deleteTeam) | **DELETE** /teams/{id} | Deletes a Team by ID
[**listTeam**](TeamApi.md#listTeam) | **GET** /teams | List Teams
[**listTeamHosts**](TeamApi.md#listTeamHosts) | **GET** /teams/{id}/hosts | List attached Hosts
[**listTeamUsers**](TeamApi.md#listTeamUsers) | **GET** /teams/{id}/users | List attached Users
[**readTeam**](TeamApi.md#readTeam) | **GET** /teams/{id} | Find a Team by ID
[**readTeamCompetition**](TeamApi.md#readTeamCompetition) | **GET** /teams/{id}/competition | Find the attached Competition
[**updateTeam**](TeamApi.md#updateTeam) | **PATCH** /teams/{id} | Updates a Team

<a name="createTeam"></a>
# **createTeam**
> TeamCreate createTeam(body)

Create a new Team

Creates a new Team and persists it to storage.

### Example
```javascript
import {EntSchemaApi} from 'ent_schema_api';

let apiInstance = new EntSchemaApi.TeamApi();
let body = new EntSchemaApi.TeamsBody(); // TeamsBody | Team to create

apiInstance.createTeam(body, (error, data, response) => {
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
 **body** | [**TeamsBody**](TeamsBody.md)| Team to create | 

### Return type

[**TeamCreate**](TeamCreate.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="deleteTeam"></a>
# **deleteTeam**
> deleteTeam(id)

Deletes a Team by ID

Deletes the Team with the requested ID.

### Example
```javascript
import {EntSchemaApi} from 'ent_schema_api';

let apiInstance = new EntSchemaApi.TeamApi();
let id = 56; // Number | ID of the Team

apiInstance.deleteTeam(id, (error, data, response) => {
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
 **id** | **Number**| ID of the Team | 

### Return type

null (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

<a name="listTeam"></a>
# **listTeam**
> [TeamList] listTeam(opts)

List Teams

List Teams.

### Example
```javascript
import {EntSchemaApi} from 'ent_schema_api';

let apiInstance = new EntSchemaApi.TeamApi();
let opts = { 
  'page': 56, // Number | what page to render
  'itemsPerPage': 56 // Number | item count to render per page
};
apiInstance.listTeam(opts, (error, data, response) => {
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

[**[TeamList]**](TeamList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

<a name="listTeamHosts"></a>
# **listTeamHosts**
> [TeamHostsList] listTeamHosts(id, opts)

List attached Hosts

List attached Hosts.

### Example
```javascript
import {EntSchemaApi} from 'ent_schema_api';

let apiInstance = new EntSchemaApi.TeamApi();
let id = 56; // Number | ID of the Team
let opts = { 
  'page': 56, // Number | what page to render
  'itemsPerPage': 56 // Number | item count to render per page
};
apiInstance.listTeamHosts(id, opts, (error, data, response) => {
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
 **id** | **Number**| ID of the Team | 
 **page** | **Number**| what page to render | [optional] 
 **itemsPerPage** | **Number**| item count to render per page | [optional] 

### Return type

[**[TeamHostsList]**](TeamHostsList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

<a name="listTeamUsers"></a>
# **listTeamUsers**
> [TeamUsersList] listTeamUsers(id, opts)

List attached Users

List attached Users.

### Example
```javascript
import {EntSchemaApi} from 'ent_schema_api';

let apiInstance = new EntSchemaApi.TeamApi();
let id = 56; // Number | ID of the Team
let opts = { 
  'page': 56, // Number | what page to render
  'itemsPerPage': 56 // Number | item count to render per page
};
apiInstance.listTeamUsers(id, opts, (error, data, response) => {
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
 **id** | **Number**| ID of the Team | 
 **page** | **Number**| what page to render | [optional] 
 **itemsPerPage** | **Number**| item count to render per page | [optional] 

### Return type

[**[TeamUsersList]**](TeamUsersList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

<a name="readTeam"></a>
# **readTeam**
> TeamRead readTeam(id)

Find a Team by ID

Finds the Team with the requested ID and returns it.

### Example
```javascript
import {EntSchemaApi} from 'ent_schema_api';

let apiInstance = new EntSchemaApi.TeamApi();
let id = 56; // Number | ID of the Team

apiInstance.readTeam(id, (error, data, response) => {
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
 **id** | **Number**| ID of the Team | 

### Return type

[**TeamRead**](TeamRead.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

<a name="readTeamCompetition"></a>
# **readTeamCompetition**
> TeamCompetitionRead readTeamCompetition(id)

Find the attached Competition

Find the attached Competition of the Team with the given ID

### Example
```javascript
import {EntSchemaApi} from 'ent_schema_api';

let apiInstance = new EntSchemaApi.TeamApi();
let id = 56; // Number | ID of the Team

apiInstance.readTeamCompetition(id, (error, data, response) => {
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
 **id** | **Number**| ID of the Team | 

### Return type

[**TeamCompetitionRead**](TeamCompetitionRead.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

<a name="updateTeam"></a>
# **updateTeam**
> TeamUpdate updateTeam(body, id)

Updates a Team

Updates a Team and persists changes to storage.

### Example
```javascript
import {EntSchemaApi} from 'ent_schema_api';

let apiInstance = new EntSchemaApi.TeamApi();
let body = new EntSchemaApi.TeamsIdBody(); // TeamsIdBody | Team properties to update
let id = 56; // Number | ID of the Team

apiInstance.updateTeam(body, id, (error, data, response) => {
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
 **body** | [**TeamsIdBody**](TeamsIdBody.md)| Team properties to update | 
 **id** | **Number**| ID of the Team | 

### Return type

[**TeamUpdate**](TeamUpdate.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

