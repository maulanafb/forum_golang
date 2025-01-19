package main

import (
	"fmt"
	posts "github.com/maulanafb/forum_golang/internal/handlers/posts"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/maulanafb/forum_golang/internal/configs"
	"github.com/maulanafb/forum_golang/internal/handlers/memberships"
	membershipRepo "github.com/maulanafb/forum_golang/internal/repository/memberships"
	postRepo "github.com/maulanafb/forum_golang/internal/repository/posts"
	membershipSvc "github.com/maulanafb/forum_golang/internal/service/memberships"
	postSvc "github.com/maulanafb/forum_golang/internal/service/posts"
	"github.com/maulanafb/forum_golang/pkg/internalsql"
)

func main() {
	r := gin.Default()

	var (
		cfg *configs.Config
	)

	err := configs.Init(
		configs.WithConfigFolder([]string{"./internal/configs/"}),
		configs.WithConfigFile("config"),
		configs.WithConfigType("yaml"),
	)

	if err != nil {
		log.Fatal("Gagal membaca konfigurasi: " + err.Error())
	}
	cfg = configs.Get()

	db, err := internalsql.Connect(cfg.Database.DataSourceName)
	if err != nil {
		log.Fatal("Gagal terhubung ke database: " + err.Error())
	}

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	postRepository := postRepo.NewRepository(db)

	postService := postSvc.NewService(cfg, postRepository)
	postHandler := posts.NewHandler(r, postService)
	postHandler.RegisterRoutes()
	membershipRepository := membershipRepo.NewRepository(db)

	membershipService := membershipSvc.NewService(cfg, membershipRepository)
	membershipHandler := memberships.NewHandler(r, membershipService)

	membershipHandler.RegisterRoutes()

	fmt.Println("Registered Routes:")
	for _, route := range r.Routes() {
		fmt.Printf("%s %s\n", route.Method, route.Path)
	}

	r.Run(cfg.Service.Port)
}
