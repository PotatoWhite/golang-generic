# file을 관리하는 file-service

- 목적 : file의 등록 및 삭제, 조회, 수정을 위한 서비스

## application의 목적

- file의 등록 및 삭제, 조회, 수정을 위한 서비스

## 서비스를 제공하는 방법

- REST API

## endpoint

- file 등록 : POST /files
- file 삭제 : DELETE /files/{file_id}
- file 조회 : GET /files/{file_id}
- file 수정 : PUT /files/{file_id}
- file 목록 조회 : GET /files

## file properties

- file_id : UUID
- name : String
- size : Long
- type : String
- extension : String
- created_at : LocalDateTime
- updated_at : LocalDateTime

# 진행 방법

## 1. application의 spec을 정의 한다.

- application의 spec은 application의 목적, 서비스를 제공하는 방법, endpoint, properties로 구성된다.

## 0. prerequisite

- uuid : file_id를 생성하기 위해 사용
- gin : http server를 구성하기 위해 사용
- gorm : db를 구성하기 위해 사용
- postgres : db를 구성하기 위해 사용

### environment

#### docker & postgres

```shell
docker run --name file_postgres -e POSTGRES_PASSWORD=password -e POSTGRES_DB=file-service -p 5432:5432 -d postgres

➜  ~ docker ps                                                                            
CONTAINER ID   IMAGE      COMMAND                  CREATED         STATUS         PORTS                                       NAMES
0a19222b3450   postgres   "docker-entrypoint.s…"   3 seconds ago   Up 3 seconds   0.0.0.0:5432->5432/tcp, :::5432->5432/tcp   file_postgres

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

## 1.1 /pkg/files/spec.go

```go
package files

import "time"

type File struct {
	ID        uuid.UUID
	Name      string
	Size      int64
	Type      string
	Extension string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type FileService interface {
	CreateFile(file *File) (*File, error)
	DeleteFile(fileID string) error
	GetFile(fileID string) (*File, error)
	UpdateFile(fileID string, file *File) (*File, error)
	ListFiles() ([]*File, error)
}
```
