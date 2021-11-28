package main

import (
	"foodcal/app/routes"
	userUsecase "foodcal/business/users"
	userController "foodcal/controllers/users"
	userRepo "foodcal/drivers/databases/users"

	foodUsecase "foodcal/business/foods"
	foodController "foodcal/controllers/foods"
	foodRepo "foodcal/drivers/databases/foods"

	transactionsUsecase "foodcal/business/transactions"
	transactionsController "foodcal/controllers/transactions"
	transactionsRepo "foodcal/drivers/databases/transactions"

	"foodcal/drivers/mysql"
	"log"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func init() {
	viper.SetConfigFile("app/config/config.json")
	err := viper.ReadInConfig()

	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}
}

func dbMigrate(db *gorm.DB) {
	db.AutoMigrate(&userRepo.User{}, &foodRepo.Food{}, &transactionsRepo.Transaction{})
}

func main() {
	configDb := mysql.ConfigDB{
		DB_Username: viper.GetString(`database.user`),
		DB_Password: viper.GetString(`database.pass`),
		DB_Host:     viper.GetString(`database.host`),
		DB_Port:     viper.GetString(`database.port`),
		DB_Database: viper.GetString(`database.name`),
	}

	db := configDb.InitialDb()
	dbMigrate(db)

	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second

	e := echo.New()
	userRepoInterface := userRepo.NewUserRepository(db)
	userUseCaseInterface := userUsecase.NewUseCase(userRepoInterface, timeoutContext)
	userControllerInterface := userController.NewUserController(userUseCaseInterface)

	foodRepoInterface := foodRepo.NewFoodRepository(db)
	foodUsecaseInterface := foodUsecase.NewUseCase(foodRepoInterface, timeoutContext)
	foodControllerInterface := foodController.NewFoodController(foodUsecaseInterface)

	transactionRepoInterface := transactionsRepo.NewTransactionRepository(db)
	transactionUsecaseInterface := transactionsUsecase.NewUseCase(transactionRepoInterface, timeoutContext)
	transactionControllerInterface := transactionsController.NewTransactionController(transactionUsecaseInterface)

	routesInit := routes.RouteControllerList{
		UserController:        *userControllerInterface,
		FoodController:        *foodControllerInterface,
		TransactionController: *transactionControllerInterface,
	}

	routesInit.RouteRegister(e)

	log.Fatal(e.Start(viper.GetString("server.address")))
}
