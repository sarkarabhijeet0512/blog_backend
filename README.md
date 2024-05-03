# Blog Backend
Blogger engineering backend

Note:Tables Get auto Created when project is started
need to add the resource rows.
## Run the project in local 
cd cmd
go run .

## To generate new errors using stringer
cd er
go generate

# Authentication
## User Registration 
### PUT http://localhost:8765/v1/auth/api/user_registration 

```
{
    "email":"abc@abc.com",
    "mobile":"7000005555",
    "first_name":"Abhijeet",
    "last_name":"sarkar",
    "password":"12345"
}
```

## User Login

### POST http://localhost:8765/v1/auth/api/user_login

```
{
    "email":"abc@abc.com",
    "password":"12345"
}
```

# Authorization 
## Create UserRole
### PUT http://localhost:8765/v1/auth/api/user_role

```
{
    "role_name": "admin",
    "role_permission": [
        {
            "permission": 
                {
                    "read": true,
                    "write": false,
                    "edit": true,
                    "remove": false
                },
            "resourse_id": 1 //which components they have access to.
        }
    ]
}
```

## Get role list by id
### GET http://localhost:8765/v1/auth/api/role_details/1

## Get all roles
### GET http://localhost:8765/v1/auth/api/role_list

## Assign roles to user
### POST http://localhost:8765/v1/auth/api/assign_role
```
{
    "role_id":[3],
    "user_id":1
}
```

## Get role assigned to user 
### GET http://localhost:8765/v1/auth/api/user/assigned_role
NOTE: no need to send user id as it gets fetched from jwt token when login is done 

