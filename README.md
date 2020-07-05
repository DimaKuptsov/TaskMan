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

### Projects:

__Get all:__

_Method:_ GET

_URL:_ `{host}/projects/all`

__Get by ID:__

_Method:_ GET

_URL:_ `{host}/projects/{id}`

_Params:_
* __id__ - uuid, required

__Create:__

_Method:_ POST

_URL:_ `{host}/projects/create`

_Params:_
* __name__ - string, min length 1, max length 500, required
* __description__ - string, max length 1000

__Update:__

_Method:_ PUT

_URL:_ `{host}/projects/update`

_Params:_
* __id__ - uuid, required
* __name__ - string, min length 1, max length 500
* __description__ - string, max length 1000

__Delete:__

_Method:_ DELETE

_URL:_ `{host}/projects/delete/{id}`

_Params:_
* __id__ - uuid, required

### Columns:

__Get for the project:__

_Method:_ GET

_URL:_ `{host}/columns/project/{projectId}`

_Params:_
* __projectId__ - uuid, required

__Get by ID:__

_Method:_ GET

_URL:_ `{host}/columns/{id}`

_Params:_
* __id__ - uuid, required

__Create:__

_Method:_ POST

_URL:_ `{host}/columns/create`

_Params:_
* __projectId__ - uuid, required
* __name__ - string, min length 1, max length 255, required

__Update:__

_Method:_ PUT

_URL:_ `{host}/columns/update`

_Params:_
* __id__ - uuid, required
* __name__ - string, min length 1, max length 255
* __priority__ - integer, min 1

__Delete:__

_Method:_ DELETE

_URL:_ `{host}/columns/delete/{id}`

_Params:_
* __id__ - uuid, required

### Tasks:

__Get for the column:__

_Method:_ GET

_URL:_ `{host}/tasks/column/{columnId}`

_Params:_
* __columnId__ - uuid, required

__Get by ID:__

_Method:_ GET

_URL:_ `{host}/tasks/{id}`

_Params:_
* __id__ - uuid, required

__Create:__

_Method:_ POST

_URL:_ `{host}/tasks/create`

_Params:_
* __columnId__ - uuid, required
* __name__ - string, min length 1, max length 500, required
* __description__ - string, max length 5000

__Update:__

_Method:_ PUT

_URL:_ `{host}/tasks/update`

_Params:_
* __id__ - uuid, required
* __name__ - string, min length 1, max length 500
* __description__ - string, max length 5000
* __priority__ - integer, min 1

__Delete:__

_Method:_ DELETE

_URL:_ `{host}/tasks/delete/{id}`

_Params:_
* __id__ - uuid, required

### Comments:

__Get for the task:__

_Method:_ GET

_URL:_ `{host}/comments/task/{taskId}`

_Params:_
* __taskId__ - uuid, required

__Get by ID:__

_Method:_ GET

_URL:_ `{host}/comments/{id}`

_Params:_
* __id__ - uuid, required

__Create:__

_Method:_ POST

_URL:_ `{host}/comments/create`

_Params:_
* __taskId__ - uuid, required
* __text__ - string, min length 1, max length 5000, required

__Update:__

_Method:_ PUT

_URL:_ `{host}/comments/update`

_Params:_
* __id__ - uuid, required
* __text__ - string, min length 1, max length 5000, required

__Delete:__

_Method:_ DELETE

_URL:_ `{host}/comments/delete/{id}`

_Params:_
* __id__ - uuid, required
