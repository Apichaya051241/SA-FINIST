package main

import (
	"context"
	"log"

	"github.com/Yui/app/controllers"
	_ "github.com/Yui/app/docs"
	"github.com/Yui/app/ent"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Users struct {
	User []User
}
type User struct {
	USER_NAME     string
	USER_EMAIL    string
	USER_PASSWORD string
}

type Stores struct {
	Store []Store
}
type Store struct {
	STORE_NAME string
}

type Drugs struct {
	Drug []Drug
}
type Drug struct {
	DRUG_NAME string
}

// type Registerstores struct {
// 	Registerstore []Registerstore
// }
// type Registerstore struct {
// 	AMOUNT int
// }

// @title SUT SA Example API
// @version 1.0
// @description This is a sample server for SUT SE 2563
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @securitydefinitions.oauth2.application OAuth2Application
// @tokenUrl https://example.com/oauth/token
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.implicit OAuth2Implicit
// @authorizationUrl https://example.com/oauth/authorize
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl https://example.com/oauth/token
// @scope.read Grants read access
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.accessCode OAuth2AccessCode
// @tokenUrl https://example.com/oauth/token
// @authorizationUrl https://example.com/oauth/authorize
// @scope.admin Grants read and write access to administrative information
func main() {
	router := gin.Default()
	router.Use(cors.Default())

	client, err := ent.Open("sqlite3", "file:ent.db?cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("fail to open sqlite3: %v", err)
	}
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	v1 := router.Group("/api/v1")
	controllers.NewUserController(v1, client)
	controllers.NewStoreController(v1, client)
	controllers.NewDrugController(v1, client)
	controllers.NewRegisterstoreController(v1, client)

	// Set Users Data
	users := Users{
		User: []User{
			User{"สมชาย หมั่นเพียร", "somchai@gmail.com", "12345"},
			User{"วรรณี ใจชื่น", "anee@gmail.com", "12345"},
		},
	}

	for _, u := range users.User {
		client.User.
			Create().
			SetEmail(u.USER_EMAIL).
			SetName(u.USER_NAME).
			SetPassword(u.USER_PASSWORD).
			Save(context.Background())
	}

	drugs := Drugs{
		Drug: []Drug{
			Drug{"Aspirin"},
			Drug{"Loperamide"},
		},
	}

	for _, d := range drugs.Drug {
		client.Drug.
			Create().
			SetName(d.DRUG_NAME).
			Save(context.Background())
	}

	// Set Stores Data
	stores := Stores{
		Store: []Store{
			Store{"คลังยาภายใน"},
			Store{"คลังยาภายนอก"},
		},
	}
	for _, s := range stores.Store{
		client.Store.
			Create().
			SetName(s.STORE_NAME).
			Save(context.Background())
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run()
}
