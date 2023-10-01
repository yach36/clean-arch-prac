# grpc Server Reference

## client setup
install evans for https://github.com/ktr0731/evans#installation

```bash
# connect to grpc server
evans --host localhost --port 50051 -r
```
enter user_grpc.UserService
```
> package user_grpc
> service UserService
```

## rpc specification
- ### GetUserList
call
```
> call GetUserList
```
return
```json
{
  "Users": [
    {
      "Age": "21",
      "ID": "1",
      "Name": "taro"
    },
    {
      "Age": "26",
      "ID": "2",
      "Name": "kojiro"
    },
    {
      "Age": "29",
      "ID": "3",
      "Name": "momoko"
    },
    {
      "Age": "34",
      "ID": "4",
      "Name": "hanako"
    }
  ]
}
```

- ### GetUser
call
```
> call GetUser
id (TYPE_INT64) => 1
```
return
```json
{
  "Age": "21",
  "ID": "1",
  "Name": "taro"
}
```

- ### RegisterUser
call
```
> call RegisterUser
Name (TYPE_STRING) => takashi
Age (TYPE_INT64) => 24
```
return
```json
{
  "message": "success",
  "status": "200"
}
```

- ### DeleteUser
call
```
> call DeleteUser
id (TYPE_INT64) => 2
```
return
```json
{
  "message": "success",
  "status": "200"
}
```
