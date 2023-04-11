package main

import (
	"comp4913-backend/configs"
	"comp4913-backend/internal/controllers"
	"comp4913-backend/internal/db"
	"comp4913-backend/internal/services"
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"os"
)

func init() {
	if err := godotenv.Load("./env/.env.local"); err != nil {
		log.Fatal(err)
	}
}

func main() {
	orm, err := db.New(db.GetDBConfig()...)
	if err != nil {
		log.Fatal(err)
	}

	//Service
	userSvc := services.NewUserService(orm)
	ipfsClient := services.NewIPFSClient(os.Getenv("IPFS_PRJ_ID"), os.Getenv("IPFS_API_SECRET"))
	ethClient := services.NewEthClient(fmt.Sprintf("ws://%s:%d", configs.RPC, configs.PORT))
	bookSvc := services.NewBookService(orm, ethClient)
	ethWatcherSvc := services.NewEthWatcherService(orm, ethClient, bookSvc, context.Background())
	go ethWatcherSvc.BookIndexing(configs.ERC721FactoryAddress)
	go ethWatcherSvc.SubscribeBook(configs.ERC721FactoryAddress)
	go ethWatcherSvc.BookTransferIndexing()
	go ethWatcherSvc.SubscribeBookTransfer()

	//Controller
	userController := controllers.NewUserController(userSvc, orm)
	bookController := controllers.NewBookController(bookSvc, ipfsClient, orm)

	e := echo.New()
	e.HideBanner = true
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Pre(middleware.AddTrailingSlash())
	e.Use(middleware.CORS())

	userRoute := e.Group("/users")
	userRoute.GET("/:address/", userController.GetUser)
	userRoute.GET("/:address/books/", userController.GetUserWithBooks)
	userRoute.GET("/:address/publishes/", userController.GetUserWithPublishes)

	bookRoute := e.Group("/books")
	bookRoute.GET("/", bookController.GetBooks)
	bookRoute.GET("/:address/", bookController.GetBookByAddress)
	bookRoute.POST("/", bookController.UploadBook)
	//bookRoute.POST("/:address:/request/", bookController.RequestBook)
	//bookRoute.POST("/:address:/approve/", bookController.ApproveRequest)
	//bookRoute.POST("/approve/", bookController.ApproveAllRequest)

	if err := e.Start(":8080"); err != nil {
		log.Fatal(err)
	}
}
