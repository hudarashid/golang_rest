package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/gin-gonic/gin"
	"github.com/hudarashid/golang_rest/controllers"
	dbCon "github.com/hudarashid/golang_rest/db/sqlc"
	"github.com/hudarashid/golang_rest/routes"
	"github.com/hudarashid/golang_rest/util"
)

var (
	server *gin.Engine
	db     *dbCon.Queries
	ctx    context.Context

	ContactController controllers.ContactController
	ContactRoutes     routes.ContactRoutes
)

func init() {
	ctx = context.TODO()
	config, err := util.LoadConfig(".")

	if err != nil {
		log.Fatalf("could not loadconfig: %v", err)
	}

	conn, err := sql.Open(config.DbDriver, config.DbSource)
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	db = dbCon.New(conn)

	fmt.Println("PostgreSql connected successfully...")

	ContactController = *controllers.NewContactController(db, ctx)
	ContactRoutes = routes.NewRouteContact(ContactController)

	server = gin.Default()
}

func main() {
	config, err := util.LoadConfig(".")

	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	router := server.Group("/api")

	router.GET("/healthcheck", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "The contact APi is working fine"})
	})

	ContactRoutes.ContactRoute(router)

	server.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "failed", "message": fmt.Sprintf("The specified route %s not found", ctx.Request.URL)})
	})

	log.Fatal(server.Run(":" + config.ServerAddress))
}
