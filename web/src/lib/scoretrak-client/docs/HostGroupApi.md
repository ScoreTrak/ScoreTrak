# EntSchemaApi.HostGroupApi

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**createHostGroup**](HostGroupApi.md#createHostGroup) | **POST** /host-groups | Create a new HostGroup
[**deleteHostGroup**](HostGroupApi.md#deleteHostGroup) | **DELETE** /host-groups/{id} | Deletes a HostGroup by ID
[**listHostGroup**](HostGroupApi.md#listHostGroup) | **GET** /host-groups | List HostGroups
[**listHostGroupHosts**](HostGroupApi.md#listHostGroupHosts) | **GET** /host-groups/{id}/hosts | List attached Hosts
[**readHostGroup**](HostGroupApi.md#readHostGroup) | **GET** /host-groups/{id} | Find a HostGroup by ID
[**readHostGroupCompetition**](HostGroupApi.md#readHostGroupCompetition) | **GET** /host-groups/{id}/competition | Find the attached Competition
[**readHostGroupTeam**](HostGroupApi.md#readHostGroupTeam) | **GET** /host-groups/{id}/team | Find the attached Team
[**updateHostGroup**](HostGroupApi.md#updateHostGroup) | **PATCH** /host-groups/{id} | Updates a HostGroup

<a name="createHostGroup"></a>
# **createHostGroup**
> HostGroupCreate createHostGroup(body)

Create a new HostGroup

Creates a new HostGroup and persists it to storage.

### Example
```javascript
import {EntSchemaApi} from 'ent_schema_api';

let apiInstance = new EntSchemaApi.HostGroupApi();
let body = new EntSchemaApi.HostgroupsBody(); // HostgroupsBody | HostGroup to create

apiInstance.createHostGroup(body, (error, data, response) => {
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
 **body** | [**HostgroupsBody**](HostgroupsBody.md)| HostGroup to create | 

### Return type

[**HostGroupCreate**](HostGroupCreate.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="deleteHostGroup"></a>
# **deleteHostGroup**
> deleteHostGroup(id)

Deletes a HostGroup by ID

Deletes the HostGroup with the requested ID.

### Example
```javascript
import {EntSchemaApi} from 'ent_schema_api';

let apiInstance = new EntSchemaApi.HostGroupApi();
let id = 56; // Number | ID of the HostGroup

apiInstance.deleteHostGroup(id, (error, data, response) => {
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
 **id** | **Number**| ID of the HostGroup | 

### Return type

null (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

<a name="listHostGroup"></a>
# **listHostGroup**
> [HostGroupList] listHostGroup(opts)

List HostGroups

List HostGroups.

### Example
```javascript
import {EntSchemaApi} from 'ent_schema_api';

let apiInstance = new EntSchemaApi.HostGroupApi();
let opts = { 
  'page': 56, // Number | what page to render
  'itemsPerPage': 56 // Number | item count to render per page
};
apiInstance.listHostGroup(opts, (error, data, response) => {
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

[**[HostGroupList]**](HostGroupList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

<a name="listHostGroupHosts"></a>
# **listHostGroupHosts**
> [HostGroupHostsList] listHostGroupHosts(id, opts)

List attached Hosts

List attached Hosts.

### Example
```javascript
import {EntSchemaApi} from 'ent_schema_api';

let apiInstance = new EntSchemaApi.HostGroupApi();
let id = 56; // Number | ID of the HostGroup
let opts = { 
  'page': 56, // Number | what page to render
  'itemsPerPage': 56 // Number | item count to render per page
};
apiInstance.listHostGroupHosts(id, opts, (error, data, response) => {
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
 **id** | **Number**| ID of the HostGroup | 
 **page** | **Number**| what page to render | [optional] 
 **itemsPerPage** | **Number**| item count to render per page | [optional] 

### Return type

[**[HostGroupHostsList]**](HostGroupHostsList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

<a name="readHostGroup"></a>
# **readHostGroup**
> HostGroupRead readHostGroup(id)

Find a HostGroup by ID

Finds the HostGroup with the requested ID and returns it.

### Example
```javascript
import {EntSchemaApi} from 'ent_schema_api';

let apiInstance = new EntSchemaApi.HostGroupApi();
let id = 56; // Number | ID of the HostGroup

apiInstance.readHostGroup(id, (error, data, response) => {
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
 **id** | **Number**| ID of the HostGroup | 

### Return type

[**HostGroupRead**](HostGroupRead.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

<a name="readHostGroupCompetition"></a>
# **readHostGroupCompetition**
> HostGroupCompetitionRead readHostGroupCompetition(id)

Find the attached Competition

Find the attached Competition of the HostGroup with the given ID

### Example
```javascript
import {EntSchemaApi} from 'ent_schema_api';

let apiInstance = new EntSchemaApi.HostGroupApi();
let id = 56; // Number | ID of the HostGroup

apiInstance.readHostGroupCompetition(id, (error, data, response) => {
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
 **id** | **Number**| ID of the HostGroup | 

### Return type

[**HostGroupCompetitionRead**](HostGroupCompetitionRead.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

<a name="readHostGroupTeam"></a>
# **readHostGroupTeam**
> HostGroupTeamRead readHostGroupTeam(id)

Find the attached Team

Find the attached Team of the HostGroup with the given ID

### Example
```javascript
import {EntSchemaApi} from 'ent_schema_api';

let apiInstance = new EntSchemaApi.HostGroupApi();
let id = 56; // Number | ID of the HostGroup

apiInstance.readHostGroupTeam(id, (error, data, response) => {
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
 **id** | **Number**| ID of the HostGroup | 

### Return type

[**HostGroupTeamRead**](HostGroupTeamRead.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

<a name="updateHostGroup"></a>
# **updateHostGroup**
> HostGroupUpdate updateHostGroup(body, id)

Updates a HostGroup

Updates a HostGroup and persists changes to storage.

### Example
```javascript
import {EntSchemaApi} from 'ent_schema_api';

let apiInstance = new EntSchemaApi.HostGroupApi();
let body = new EntSchemaApi.HostgroupsIdBody(); // HostgroupsIdBody | HostGroup properties to update
let id = 56; // Number | ID of the HostGroup

apiInstance.updateHostGroup(body, id, (error, data, response) => {
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
 **body** | [**HostgroupsIdBody**](HostgroupsIdBody.md)| HostGroup properties to update | 
 **id** | **Number**| ID of the HostGroup | 

### Return type

[**HostGroupUpdate**](HostGroupUpdate.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

