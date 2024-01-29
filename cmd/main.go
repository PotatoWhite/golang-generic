package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"restfule-api-generic/internal/easywalk/handler"
	"restfule-api-generic/internal/easywalk/repository"
	"restfule-api-generic/internal/easywalk/service"
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

	// 서비스 및 핸들러 초기화
	repo := repository.NewSimplyRepository[*model.User](db)
	svc := service.NewGenericService[*model.User](repo)

	// Gin 라우터 설정
	r := gin.Default()
	group := r.Group("/users")

	handler.NewHandler[*model.User](group, svc)

	r.Run() // listen and serve on 0.0.0.0:8080
}
