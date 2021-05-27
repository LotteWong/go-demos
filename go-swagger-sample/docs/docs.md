# Title For Go-Swagger-Sample Api Docs
Description For Go-Swagger-Sample Api Docs

## Version: 1.0.0

### /users

#### GET
##### Summary

查询用户列表

##### Description

查询用户列表描述

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| username | query | 登录名称 | No | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | success | [models.Users](#modelsusers) |

#### POST
##### Summary

创建用户

##### Description

创建用户

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| createBody | body | 创建用户请求主体 | Yes | [models.User](#modelsuser) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | success | [models.User](#modelsuser) |

### /users/{id}

#### GET
##### Summary

查询用户详情

##### Description

查询用户详情描述

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| id | path | 用户标识 | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | success | [models.User](#modelsuser) |

### Models

#### models.User

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| id | string | 用户标识<br>_Example:_ `"1"` | No |
| password | string | 登录密码<br>_Example:_ `"123456"` | No |
| username | string | 登录名称<br>_Example:_ `"admin"` | No |

#### models.Users

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| items | [ [models.User](#modelsuser) ] | 用户列表 | No |
| totalcount | integer | 共计条数<br>_Example:_ `0` | No |
