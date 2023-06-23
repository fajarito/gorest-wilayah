package main

import (
	"fmt"
	"gorest-krs-wilayah/handler"
	"gorest-krs-wilayah/kabupaten"
	"gorest-krs-wilayah/kecamatan"
	"gorest-krs-wilayah/kelurahan"
	"gorest-krs-wilayah/provinsi"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	pagination "github.com/webstradev/gin-pagination"

	_ "gorest-krs-wilayah/docs"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// swagger embed files
// @title BKKBN Digital Service - Wilayah
// @version         2.2.0
// @description     Digital Service BKKBN for Integration - Wilayah Service
// @termsOfService  https://bkkbn.go.id
// @contact.name   Direktorat Teknologi Informasi dan Data
// @contact.url    https://bkkbn.go.id
// @contact.email  dittifdok@bkkbn.go.id
// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
// @host      localhost:8080
// @BasePath  /v1/api/2022/wilayah
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db_host := os.Getenv("DB_HOST")
	db_port := os.Getenv("DB_PORT")
	db_database := os.Getenv("DB_DATABASE")
	db_username := os.Getenv("DB_USERNAME")
	db_password := os.Getenv("DB_PASSWORD")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", db_host, db_username, db_password, db_database, db_port)
	// dsn := "host=36.37.120.165 user=usr_bkkbn password=bkkbn$1 dbname=BKKBN port=5435 sslmode=disable TimeZone=Asia/Shanghai"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger:                 logger.Default.LogMode(logger.Info),
		SkipDefaultTransaction: true,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "sch_bkkbn.", // schema name
			SingularTable: false,
		}, PrepareStmt: true,
	})

	if err != nil {
		log.Fatal("DB Connection Error")
	}

	provinsiRepository := provinsi.NewRepository(db)
	provinsiService := provinsi.NewService(provinsiRepository)
	provinsiHandler := handler.NewProvinsiHandler(provinsiService)
	kabupatenRepository := kabupaten.NewRepository(db)
	kabupatenService := kabupaten.NewService(kabupatenRepository)
	kabupatenHandler := handler.NewKabupatenHandler(kabupatenService)
	kecamatanRepository := kecamatan.NewRepository(db)
	kecamatanService := kecamatan.NewService(kecamatanRepository)
	kecamatanHandler := handler.NewKecamatanHandler(kecamatanService)
	kelurahanRepository := kelurahan.NewRepository(db)
	kelurahanService := kelurahan.NewService(kelurahanRepository)
	kelurahanHandler := handler.NewKelurahanHandler(kelurahanService)

	fmt.Println("==================v3==================v3==================v3==================v")

	r := gin.Default()

	corsConfig := cors.DefaultConfig()

	// corsConfig.AllowOrigins = []string{"https://example.com"}
	corsConfig.AllowAllOrigins = true
	// To be able to send tokens to the server.
	corsConfig.AllowCredentials = true

	// OPTIONS method for ReactJS
	corsConfig.AddAllowMethods("OPTIONS")

	// Register the middleware
	r.Use(cors.New(corsConfig))

	r.Use(pagination.Default())

	v3 := r.Group("/v1/api")
	v3.GET("/2022/wilayah/showprovinsi", provinsiHandler.GetProvinsi)
	v3.GET("/2022/wilayah/showkabupaten", kabupatenHandler.GetKabupaten)
	v3.GET("/2022/wilayah/showkecamatan", kecamatanHandler.GetKecamatan)
	v3.GET("/2022/wilayah/showkelurahan", kelurahanHandler.GetKelurahan)
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8080")
}
