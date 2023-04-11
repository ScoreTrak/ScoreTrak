# EntSchemaApi.UserApi

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**createUser**](UserApi.md#createUser) | **POST** /users | Create a new User
[**deleteUser**](UserApi.md#deleteUser) | **DELETE** /users/{id} | Deletes a User by ID
[**listUser**](UserApi.md#listUser) | **GET** /users | List Users
[**listUserCompetitions**](UserApi.md#listUserCompetitions) | **GET** /users/{id}/competitions | List attached Competitions
[**listUserTeams**](UserApi.md#listUserTeams) | **GET** /users/{id}/teams | List attached Teams
[**readUser**](UserApi.md#readUser) | **GET** /users/{id} | Find a User by ID
[**updateUser**](UserApi.md#updateUser) | **PATCH** /users/{id} | Updates a User

<a name="createUser"></a>
# **createUser**
> UserCreate createUser(body)

Create a new User

Creates a new User and persists it to storage.

### Example
```javascript
import {EntSchemaApi} from 'ent_schema_api';

let apiInstance = new EntSchemaApi.UserApi();
let body = new EntSchemaApi.UsersBody(); // UsersBody | User to create

apiInstance.createUser(body, (error, data, response) => {
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
 **body** | [**UsersBody**](UsersBody.md)| User to create | 

### Return type

[**UserCreate**](UserCreate.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="deleteUser"></a>
# **deleteUser**
> deleteUser(id)

Deletes a User by ID

Deletes the User with the requested ID.

### Example
```javascript
import {EntSchemaApi} from 'ent_schema_api';

let apiInstance = new EntSchemaApi.UserApi();
let id = 56; // Number | ID of the User

apiInstance.deleteUser(id, (error, data, response) => {
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
 **id** | **Number**| ID of the User | 

### Return type

null (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

<a name="listUser"></a>
# **listUser**
> [UserList] listUser(opts)

List Users

List Users.

### Example
```javascript
import {EntSchemaApi} from 'ent_schema_api';

let apiInstance = new EntSchemaApi.UserApi();
let opts = { 
  'page': 56, // Number | what page to render
  'itemsPerPage': 56 // Number | item count to render per page
};
apiInstance.listUser(opts, (error, data, response) => {
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

[**[UserList]**](UserList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

<a name="listUserCompetitions"></a>
# **listUserCompetitions**
> [UserCompetitionsList] listUserCompetitions(id, opts)

List attached Competitions

List attached Competitions.

### Example
```javascript
import {EntSchemaApi} from 'ent_schema_api';

let apiInstance = new EntSchemaApi.UserApi();
let id = 56; // Number | ID of the User
let opts = { 
  'page': 56, // Number | what page to render
  'itemsPerPage': 56 // Number | item count to render per page
};
apiInstance.listUserCompetitions(id, opts, (error, data, response) => {
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
 **id** | **Number**| ID of the User | 
 **page** | **Number**| what page to render | [optional] 
 **itemsPerPage** | **Number**| item count to render per page | [optional] 

### Return type

[**[UserCompetitionsList]**](UserCompetitionsList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

<a name="listUserTeams"></a>
# **listUserTeams**
> [UserTeamsList] listUserTeams(id, opts)

List attached Teams

List attached Teams.

### Example
```javascript
import {EntSchemaApi} from 'ent_schema_api';

let apiInstance = new EntSchemaApi.UserApi();
let id = 56; // Number | ID of the User
let opts = { 
  'page': 56, // Number | what page to render
  'itemsPerPage': 56 // Number | item count to render per page
};
apiInstance.listUserTeams(id, opts, (error, data, response) => {
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
 **id** | **Number**| ID of the User | 
 **page** | **Number**| what page to render | [optional] 
 **itemsPerPage** | **Number**| item count to render per page | [optional] 

### Return type

[**[UserTeamsList]**](UserTeamsList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

<a name="readUser"></a>
# **readUser**
> UserRead readUser(id)

Find a User by ID

Finds the User with the requested ID and returns it.

### Example
```javascript
import {EntSchemaApi} from 'ent_schema_api';

let apiInstance = new EntSchemaApi.UserApi();
let id = 56; // Number | ID of the User

apiInstance.readUser(id, (error, data, response) => {
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
 **id** | **Number**| ID of the User | 

### Return type

[**UserRead**](UserRead.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

<a name="updateUser"></a>
# **updateUser**
> UserUpdate updateUser(body, id)

Updates a User

Updates a User and persists changes to storage.

### Example
```javascript
import {EntSchemaApi} from 'ent_schema_api';

let apiInstance = new EntSchemaApi.UserApi();
let body = new EntSchemaApi.UsersIdBody(); // UsersIdBody | User properties to update
let id = 56; // Number | ID of the User

apiInstance.updateUser(body, id, (error, data, response) => {
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
 **body** | [**UsersIdBody**](UsersIdBody.md)| User properties to update | 
 **id** | **Number**| ID of the User | 

### Return type

[**UserUpdate**](UserUpdate.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

