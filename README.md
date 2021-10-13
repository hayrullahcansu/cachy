# cachy

# ![bashtop](logo.png)

# Description

Cachy is a lightweight in-memory cache api.

# Features

* No thirdparty libraries.
* Docker image.

# Configurability

All options changeable from within UI.

#### settings.json 

```json
{
  "backup_file_path": "TIMESTAMP-data.json", 
  "backup_interval": "30"   //by sec
}
```

| Param Name | Variable Type | Requirement | Description                         | Value                                                                             |
|------------|---------------|-------------|-------------------------------------|-------------------------------------------------------------------------------------|
| backup_file_path       | `:string`     |    no`*`   | Path for backing up                       |   Described below  |
| backup_interval      | `:string`     |    no`*`   | Interval for backing up         |  `10`, `30`, `60` ... by any seconds  |

#### default values for backup_file_path. 
```
darwin   ->  "/tmp/TIMESTAMP-data.json"
linux    ->  "/tmp/TIMESTAMP-data.json"
windows  ->  "TIMESTAMP-data.json"
```

# Rest API

All endpoints here
- Cache Resource `/api/v1/cache`
  - GET `/` List of all cache entries
  - GET `/{cache_key}` returns cache entry
    - Response body: 
    ```json
    {
        "key": "cachy_test_key",
        "value": "cachy_test_value",
        "expire_at": "13-10-2021 17:57:30"
    }
    ```
  - POST `/{cache_key}` sets new cache entry
    - Request body: 
    ```json
    {
        "time_span": 10,
        "data":"cachy_test_value"
    }
    ```
    - Response body: 
    ```json
    {
        "key": "cachy_test_key",
        "value": "cachy_test_value",
        "expire_at": "13-10-2021 17:57:30"
    }
    ```
  - PUT `/{cache_key}` updates existing cache entry
     - Request body: 
    ```json
    {
        "time_span": 20,
        "data":"cachy_test_value"
    }
    ```
      - Response body: 
    ```json
    {
        "key": "cachy_test_key",
        "value": "cachy_test_value",
        "expire_at": "13-10-2021 18:00:18"
    }
    ```

  - DELETE `/{cache_key}` removes cache entry
    
  - DELETE `/flush` removes all cache entries
    

# Docker

- Create docker image `docker build -t cachy-api --squash .` 
- Create/Run a docker container `docker run -d --rm -p 8080:8080 --name cachy-api-1 cachy-api` 
  
# TODO

- [x] Design Patters
- [x] Readme File
- [ ] Go Doc
- [x] ApiDoc
- [ ] Web UI
- [x] Tests
- [x] Logging for Http Requests that are incoming in `server.log` file
- [x] Docker Support
- [ ] Deploy (Heroku, aws)