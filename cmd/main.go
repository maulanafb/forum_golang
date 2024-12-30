package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/maulanafb/forum_golang/internal/configs"
	"github.com/maulanafb/forum_golang/internal/handlers/memberships"
	membershipRepo "github.com/maulanafb/forum_golang/internal/repository/memberships"
	membershipSvc "github.com/maulanafb/forum_golang/internal/service/memberships"
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
		log.Fatal("Gagal membaca konfigurasi" + err.Error())
	}
	cfg = configs.Get()

	db, err := internalsql.Connect(cfg.Database.DataSourceName)

	if err != nil {
		log.Fatal("Gagal terhubung ke database" + err.Error())
	}

	membershipRepo := membershipRepo.NewRepository(db)
	membershipService := membershipSvc.NewService(membershipRepo)
	membershipHandler := memberships.NewHandler(r, membershipService)

	membershipHandler.RegisterRoutes()
	r.Run(cfg.Service.Port)
}
