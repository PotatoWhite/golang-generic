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
