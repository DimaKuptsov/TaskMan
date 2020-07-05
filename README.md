# TaskMan
The back-end part of the simple task manager.

## Installation
* Clone the repository

`git clone git@github.com:DimaKuptsov/task-man.git`

* Fill docker environment variables (.env)
* Add Go configuration file (task-man/config.json)
```
  "app_name": "task-man",
  "env": "dev",
  "listen_url": "host:port",
  "postgres": {
    "host": postrgers host,
    "port": postgres port,
    "database": database name,
    "test_database": database for test name,
    "user": postgres user,
    "password": "postgres user pass",
    "pool_size": postgres pool size,
    "max_retries": postgres retries number,
    "read_timeout": max read timeout,
    "write_timeout": max write timeout
  }
```
* Build docker

`docker-compose up --build`

* Go into golang container

`docker exec -ti task-man-go bash`

* Run migrations

` migrate -database postgres://{db_user}:{db_pass}@{db_host}:{db_port}/{db_name}?sslmode=disable -path db/postgres/migrations/ up`

## API

[Postman collection](https://www.getpostman.com/collections/40aa541aa0bbb2b7ff88) for test

####Projects:
#####Get all:
__Method:__ GET

__URL:__ {host}/projects/all

#####Get by ID:
*Method:* GET

*URL:* {host}/projects/{id}

*Params:*
* id - uuid, required

#####Create:
*Method:* POST

*URL:* {host}/projects/create

*Params:*
* name - string, min length 1, max length 500, required
* description - string, max length 1000

#####Update:
*Method:* PUT

*URL:* {host}/projects/update

*Params:*
* id - uuid, required
* name - string, min length 1, max length 500
* description - string, max length 1000

#####Delete:
*Method:* DELETE

*URL:* {host}/projects/delete/{id}

*Params:*
* id - uuid, required

####Columns:
#####Get for the project:
__Method:__ GET

__URL:__ {host}/columns/project/{projectId}

*Params:*
* projectId - uuid, required

#####Get by ID:
*Method:* GET

*URL:* {host}/columns/{id}

*Params:*
* id - uuid, required

#####Create:
*Method:* POST

*URL:* {host}/columns/create

*Params:*
* projectId - uuid, required
* name - string, min length 1, max length 255, required

#####Update:
*Method:* PUT

*URL:* {host}/columns/update

*Params:*
* id - uuid, required
* name - string, min length 1, max length 255
* priority - integer, min 1

#####Delete:
*Method:* DELETE

*URL:* {host}/columns/delete/{id}

*Params:*
* id - uuid, required

####Tasks:
#####Get for the column:
__Method:__ GET

__URL:__ {host}/tasks/column/{columnId}

*Params:*
* columnId - uuid, required

#####Get by ID:
*Method:* GET

*URL:* {host}/tasks/{id}

*Params:*
* id - uuid, required

#####Create:
*Method:* POST

*URL:* {host}/tasks/create

*Params:*
* columnId - uuid, required
* name - string, min length 1, max length 500, required
* description - string, max length 5000

#####Update:
*Method:* PUT

*URL:* {host}/tasks/update

*Params:*
* id - uuid, required
* name - string, min length 1, max length 500
* description - string, max length 5000
* priority - integer, min 1

#####Delete:
*Method:* DELETE

*URL:* {host}/tasks/delete/{id}

*Params:*
* id - uuid, required

####Comments:
#####Get for the task:
__Method:__ GET

__URL:__ {host}/comments/task/{taskId}

*Params:*
* taskId - uuid, required

#####Get by ID:
*Method:* GET

*URL:* {host}/comments/{id}

*Params:*
* id - uuid, required

#####Create:
*Method:* POST

*URL:* {host}/comments/create

*Params:*
* taskId - uuid, required
* text - string, min length 1, max length 5000, required

#####Update:
*Method:* PUT

*URL:* {host}/comments/update

*Params:*
* id - uuid, required
* text - string, min length 1, max length 5000, required

#####Delete:
*Method:* DELETE

*URL:* {host}/comments/delete/{id}

*Params:*
* id - uuid, required
