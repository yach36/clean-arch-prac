# REST API Reference

baseURL: http://localhost:8080

- ### GET /health
```
path param: null
body param: null
return: constant json message
```
return
```json
{
    "status": 200,
    "message": "ok"
}
```
- ### GET /users
```
path param: null
body param: null
return: all users
```
return
```json
[
    {
        "id": 1,
        "name": "taro",
        "age": 21
    },
    {
        "id": 2,
        "name": "kojiro",
        "age": 26
    },
    {
        "id": 3,
        "name": "momoko",
        "age": 29
    },
    {
        "id": 4,
        "name": "hanako",
        "age": 34
    }
]
```
- ### GET /users/:id
```
path param: id(int)
body param: null
return: the user with specified id
```
return
```json
{
    "id": 1,
    "name": "taro",
    "age": 21
}
```
- ### POST /users
```
path param: null
body param: json string with name(string) and age(int)
return: json message
```
body param
```json
{
    "name": "takashi",
    "age": 24
}
```
return
```json
{
    "status": 200,
    "message": "success"
}
```
- ### DELETE /users/:id
```
path param: id(int)
body param: null
return: json message
```
return
```json
{
    "status": 200,
    "message": "success"
}
```
