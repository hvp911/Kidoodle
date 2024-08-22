package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"

	contenthandler "content/internal/routes/contents/handler"
	contentservice "content/internal/routes/contents/service"
)

//var Version = "0.0.1"

const (
	// HOST would be overridden if environment variable POSTGRESQL_HOST is set. For Docker environment variable is already configured to: "postgres"
	HOST    = "localhost"
	PORT    = 5432
	USER    = "postgres"
	PASS    = "password123"
	DBName  = "content"
	APIAddr = ":5001"
)

func main() {
	engine := initEngine()
	err := http.ListenAndServe(APIAddr, engine)
	if err != nil {
		log.Fatalf("listen: %s\n", err)
	}
}

func initEngine() *gin.Engine {
	// DB settings - ideally would be derived from configuration file
	// Using environment variable as easy option for now for host and hardcoded values for rest configurations
	host := os.Getenv("POSTGRESQL_HOST")
	if len(host) == 0 {
		host = HOST
	}
	dbConnection := initDBConnection(host, PORT, USER, PASS, DBName)

	corsMiddleware := cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{http.MethodGet},
	})

	// Gin Engine
	engine := gin.New()
	engine.Use(gin.Recovery())
	engine.Use(corsMiddleware)
	// TODO: authentication middleware can be used on the engine if required. A middleware would intercept the request and validate tokens before sending over to the engine for further processing.

	// Build endpoint handlers
	contentService := contentservice.NewReaderService(dbConnection)
	contentHandler := contenthandler.New(contentService)

	// Attach handlers to the engine
	external := engine.Group("/api/v1")
	contentHandler.SetupRoutes(external)

	return engine
}

// Prepare db connection to be used for API endpoints
func initDBConnection(host string, port int, user string, password string, dbName string) *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("unable to connect to database: %s\n", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("unable to ping database: %s\n", err)
	}
	fmt.Println("Successfully connected!")
	return db
}
