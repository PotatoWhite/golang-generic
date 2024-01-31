## 0. prerequisite

- uuid : file_id를 생성하기 위해 사용
- gin : http server를 구성하기 위해 사용
- gorm : db를 구성하기 위해 사용
- postgres : db를 구성하기 위해 사용

### environment

#### docker & postgres

```shell
docker run --name user_postgres -e POSTGRES_PASSWORD=password -e POSTGRES_DB=user-service -p 5432:5432 -d postgres

➜  ~ docker ps                                                                            
CONTAINER ID   IMAGE      COMMAND                  CREATED         STATUS         PORTS                                       NAMES
0a19222b3450   postgres   "docker-entrypoint.s…"   3 seconds ago   Up 3 seconds   0.0.0.0:5432->5432/tcp, :::5432->5432/tcp   user_postgres

```

### uuid

```shell
go get github.com/google/uuid
```

### gin

```shell
go get github.com/gin-gonic/gin
```

### gorm & gorm postgres

```shell
go get gorm.io/gorm

go get gorm.io/driver/postgres
```

### postgres

```shell
go get github.com/lib/pq
```

### easywalk

```shell
go get github.com/easywalk/go-restful
```

# 1. user-service 구현

### model 생성 (user.go)

```go
package model

import (
	"github.com/google/uuid"
)

type User struct {
	ID   uuid.UUID `gorm:"type:uuid;primary_key;"`
	Name string    `gorm:"type:varchar(255);not null;unique;"`
	Age  int
}

func (u *User) TableName() string {
	return "users"
}

func (u *User) GetID() uuid.UUID {
	return u.ID
}

func (u *User) SetID(id uuid.UUID) {
	u.ID = id
}
```

### easywalk를 이용한 http server 구성 (main.go)

```go
package main

import (
	"github.com/easywalk/go-restful/easywalk/handler"
	"github.com/easywalk/go-restful/easywalk/repository"
	"github.com/easywalk/go-restful/easywalk/service"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"restfule-api-generic/pkg/model"
)

func main() {
	// config
	username := "postgres"
	password := "password"
	database := "user-service"
	dsn := "host=localhost user=" + username + " password=" + password + " dbname=" + database + " port=5432 sslmode=disable TimeZone=Asia/Seoul"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("데이터베이스 연결 실패: " + err.Error())
	}

	// 서비스 초기화
	repo := repository.NewSimplyRepository[*model.User](db)
	svc := service.NewGenericService[*model.User](repo)

	// Gin 라우터 설정
	r := gin.Default()
	group := r.Group("/users")

	// 핸들러 설정
	handler.NewHandler[*model.User](group, svc)

	r.Run() // listen and serve on 0.0.0.0:8080
}
```

# 2. user-service 테스트
```http request
### 생성
< {%
    request.variables.clearAll()
    request.variables.set("base_name", "test");
%}
POST localhost:8080/users
Content-Type: application/json

{
    "name": "{{base_name}}-{{$uuid}}",
    "age": 30
}
> {%
    client.test("Status code is 200", function () {
        client.assert(response.status === 200, "Response status is not 200");
        client.assert(response.body.ID !== "", "Response body ID is empty");

        client.global.set("id", response.body.ID);
        client.global.set("name", response.body.Name);
        client.global.set("age", response.body.Age);
    });
%}

### 생성 검증
GET localhost:8080/users/{{id}}
Content-Type: application/json

> {%
    client.test("Status code is 200", function () {
        client.assert(response.status === 200, "Response status is not 200");
        client.assert(response.body.Name === client.global.get("name"), "Response body Name is not same :" + response.body.Name + " : " + client.global.get("name"));
        client.assert(response.body.Age == client.global.get("age"), "Response body Age is not same :" + response.body.Age + " : " + client.global.get("age"));
    });
%}


### 재생성 (실패 - 이름 중복)
POST localhost:8080/users
Content-Type: application/json

{
    "name": "{{name}}",
    "age": 30
}
> {%
    client.test("Status code is 500", function () {
        client.assert(response.status === 500, "Response status is not 200");
    });
%}

### 이름 수정
PATCH localhost:8080/users/{{id}}
Content-Type: application/json

{
    "name": "{{base_name}}-{{$uuid}}"
}

> {%
    client.test("Status code is 200", function () {
        client.assert(response.status === 200, "Response status is not 200");
        client.assert(response.body.ID === client.global.get("id"), "Response body ID is not same :" + response.body.ID + " : " + client.global.get("id"));
        client.assert(response.body.Name != client.global.get("name"), "Response body Name is not same :" + response.body.Name + " : " + client.global.get("name"));
        client.assert(response.body.Age == client.global.get("age"), "Response body Age is not same :" + response.body.Age + " : " + client.global.get("age"));

        client.global.set("id", response.body.ID);
        client.global.set("name", response.body.Name);
        client.global.set("age", response.body.Age);
    });
%}

### 이름 수정 검증
GET localhost:8080/users/{{id}}
Content-Type: application/json

> {%
    client.test("Status code is 200", function () {
        client.assert(response.status === 200, "Response status is not 200");
        client.assert(response.body.Name === client.global.get("name"), "Response body Name is not same :" + response.body.Name + " : " + client.global.get("name"));
        client.assert(response.body.Age == client.global.get("age"), "Response body Age is not same :" + response.body.Age + " : " + client.global.get("age"));
    });
%}

### 삭제
DELETE localhost:8080/users/{{id}}
Content-Type: application/json

> {%
    client.test("Status code is 204", function () {
        client.assert(response.status === 204, "Response status is not 204" + response.status);
    });
%}

### 없는 것 또 삭제 (idempotent)
DELETE localhost:8080/users/{{id}}
Content-Type: application/json

> {%
    client.test("Status code is 204", function () {
        client.assert(response.status === 204, "Response status is not 204" + response.status);
    });
%}

### 삭제 검증
GET localhost:8080/users/{{id}}
Content-Type: application/json

> {%
    client.test("Status code is 204", function () {
        client.assert(response.status === 204, "Response status is not 404" + response.status);
    });
%}
```