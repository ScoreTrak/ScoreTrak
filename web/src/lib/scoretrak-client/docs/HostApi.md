# EntSchemaApi.HostApi

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**createHost**](HostApi.md#createHost) | **POST** /hosts | Create a new Host
[**deleteHost**](HostApi.md#deleteHost) | **DELETE** /hosts/{id} | Deletes a Host by ID
[**listHost**](HostApi.md#listHost) | **GET** /hosts | List Hosts
[**listHostServices**](HostApi.md#listHostServices) | **GET** /hosts/{id}/services | List attached Services
[**readHost**](HostApi.md#readHost) | **GET** /hosts/{id} | Find a Host by ID
[**readHostCompetition**](HostApi.md#readHostCompetition) | **GET** /hosts/{id}/competition | Find the attached Competition
[**readHostHostGroup**](HostApi.md#readHostHostGroup) | **GET** /hosts/{id}/host-group | Find the attached HostGroup
[**readHostTeam**](HostApi.md#readHostTeam) | **GET** /hosts/{id}/team | Find the attached Team
[**updateHost**](HostApi.md#updateHost) | **PATCH** /hosts/{id} | Updates a Host

<a name="createHost"></a>
# **createHost**
> HostCreate createHost(body)

Create a new Host

Creates a new Host and persists it to storage.

### Example
```javascript
import {EntSchemaApi} from 'ent_schema_api';

let apiInstance = new EntSchemaApi.HostApi();
let body = new EntSchemaApi.HostsBody(); // HostsBody | Host to create

apiInstance.createHost(body, (error, data, response) => {
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
 **body** | [**HostsBody**](HostsBody.md)| Host to create | 

### Return type

[**HostCreate**](HostCreate.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="deleteHost"></a>
# **deleteHost**
> deleteHost(id)

Deletes a Host by ID

Deletes the Host with the requested ID.

### Example
```javascript
import {EntSchemaApi} from 'ent_schema_api';

let apiInstance = new EntSchemaApi.HostApi();
let id = 56; // Number | ID of the Host

apiInstance.deleteHost(id, (error, data, response) => {
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
 **id** | **Number**| ID of the Host | 

### Return type

null (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

<a name="listHost"></a>
# **listHost**
> [HostList] listHost(opts)

List Hosts

List Hosts.

### Example
```javascript
import {EntSchemaApi} from 'ent_schema_api';

let apiInstance = new EntSchemaApi.HostApi();
let opts = { 
  'page': 56, // Number | what page to render
  'itemsPerPage': 56 // Number | item count to render per page
};
apiInstance.listHost(opts, (error, data, response) => {
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

[**[HostList]**](HostList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

<a name="listHostServices"></a>
# **listHostServices**
> [HostServicesList] listHostServices(id, opts)

List attached Services

List attached Services.

### Example
```javascript
import {EntSchemaApi} from 'ent_schema_api';

let apiInstance = new EntSchemaApi.HostApi();
let id = 56; // Number | ID of the Host
let opts = { 
  'page': 56, // Number | what page to render
  'itemsPerPage': 56 // Number | item count to render per page
};
apiInstance.listHostServices(id, opts, (error, data, response) => {
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
 **id** | **Number**| ID of the Host | 
 **page** | **Number**| what page to render | [optional] 
 **itemsPerPage** | **Number**| item count to render per page | [optional] 

### Return type

[**[HostServicesList]**](HostServicesList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

<a name="readHost"></a>
# **readHost**
> HostRead readHost(id)

Find a Host by ID

Finds the Host with the requested ID and returns it.

### Example
```javascript
import {EntSchemaApi} from 'ent_schema_api';

let apiInstance = new EntSchemaApi.HostApi();
let id = 56; // Number | ID of the Host

apiInstance.readHost(id, (error, data, response) => {
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
 **id** | **Number**| ID of the Host | 

### Return type

[**HostRead**](HostRead.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

<a name="readHostCompetition"></a>
# **readHostCompetition**
> HostCompetitionRead readHostCompetition(id)

Find the attached Competition

Find the attached Competition of the Host with the given ID

### Example
```javascript
import {EntSchemaApi} from 'ent_schema_api';

let apiInstance = new EntSchemaApi.HostApi();
let id = 56; // Number | ID of the Host

apiInstance.readHostCompetition(id, (error, data, response) => {
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
 **id** | **Number**| ID of the Host | 

### Return type

[**HostCompetitionRead**](HostCompetitionRead.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

<a name="readHostHostGroup"></a>
# **readHostHostGroup**
> HostHostGroupRead readHostHostGroup(id)

Find the attached HostGroup

Find the attached HostGroup of the Host with the given ID

### Example
```javascript
import {EntSchemaApi} from 'ent_schema_api';

let apiInstance = new EntSchemaApi.HostApi();
let id = 56; // Number | ID of the Host

apiInstance.readHostHostGroup(id, (error, data, response) => {
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
 **id** | **Number**| ID of the Host | 

### Return type

[**HostHostGroupRead**](HostHostGroupRead.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

<a name="readHostTeam"></a>
# **readHostTeam**
> HostTeamRead readHostTeam(id)

Find the attached Team

Find the attached Team of the Host with the given ID

### Example
```javascript
import {EntSchemaApi} from 'ent_schema_api';

let apiInstance = new EntSchemaApi.HostApi();
let id = 56; // Number | ID of the Host

apiInstance.readHostTeam(id, (error, data, response) => {
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
 **id** | **Number**| ID of the Host | 

### Return type

[**HostTeamRead**](HostTeamRead.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

<a name="updateHost"></a>
# **updateHost**
> HostUpdate updateHost(body, id)

Updates a Host

Updates a Host and persists changes to storage.

### Example
```javascript
import {EntSchemaApi} from 'ent_schema_api';

let apiInstance = new EntSchemaApi.HostApi();
let body = new EntSchemaApi.HostsIdBody(); // HostsIdBody | Host properties to update
let id = 56; // Number | ID of the Host

apiInstance.updateHost(body, id, (error, data, response) => {
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
 **body** | [**HostsIdBody**](HostsIdBody.md)| Host properties to update | 
 **id** | **Number**| ID of the Host | 

### Return type

[**HostUpdate**](HostUpdate.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

