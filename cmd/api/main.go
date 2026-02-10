package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/naeemjr06-prog/user-access-api/internal/platform/config"
	"github.com/naeemjr06-prog/user-access-api/internal/platform/postgres"
	"github.com/naeemjr06-prog/user-access-api/internal/user/application"
	"github.com/naeemjr06-prog/user-access-api/internal/user/infrastructure/crypto"
	"github.com/naeemjr06-prog/user-access-api/internal/user/infrastructure/http"
	"github.com/naeemjr06-prog/user-access-api/internal/user/infrastructure/persistence"
)

func main() {
	config.Load()

	db, err := postgres.NewPool(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	userRepo := persistence.NewPostgresUserRepository(db)
	hasher := crypto.NewArgonHasher()

	registerUserUC := application.NewRegisterUserHandler(userRepo, hasher)
	registerAdminUC := application.NewRegisterUserAdminHandler(userRepo, hasher)

	registerHandler := http.NewRegisterHandler(registerUserUC, registerAdminUC)

	r := gin.Default()

	r.POST("/auth/register", registerHandler.Register)
	r.POST("/auth/register/admin", registerHandler.RegisterAdmin)

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	r.Run(":8080")
}
