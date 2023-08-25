package main

import (
	"BANKAPP/config"
	"BANKAPP/constants"
	"BANKAPP/controllers"
	"BANKAPP/routes"
	"BANKAPP/service"
	"context"
	"fmt"
	"log"

	//	"rest-api/services"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)


var (
	mongoclient *mongo.Client
	ctx         context.Context
	server         *gin.Engine
)
func initRoutes(){
	routes.Default(server)
}

func initApp(mongoClient *mongo.Client){
	
	//customer collection
	ctx = context.TODO()
	CustomerCollection := mongoClient.Database(constants.Dbname).Collection("customer")
	CustomerService := service.CustomerServiceInit(CustomerCollection, ctx)
	CustomerController := controllers.InitCustomerController(CustomerService)
	routes.CustomerRoute(server,CustomerController)


	//transaction collection
	ctx = context.TODO()
	TransactionCollection := mongoClient.Database(constants.Dbname).Collection("transaction")
	TransactionService := service.TransactionServiceInit(mongoClient,CustomerCollection,TransactionCollection , ctx)
	TransactionController := controllers.InitTransactionController(TransactionService)
	routes.Transactionroutes(server,TransactionController)
}

func initAcc(mongoClient *mongo.Client){
	ctx = context.TODO()
	accCollection := mongoClient.Database(constants.Dbname).Collection("account")
	accService := service.InitAccount(accCollection, ctx)
	accController := controllers.InitAccController(accService)
	routes.AccRoute(server,accController)
}

func initBank(mongoClient *mongo.Client){
	ctx = context.TODO()
	bCollection := mongoClient.Database(constants.Dbname).Collection("bank")
	bService := service.InitBank(bCollection, ctx)
	bController := controllers.InitBankController(bService)
	routes.BankRoute(server,bController)
}
func initLoan(mongoClient *mongo.Client){
	ctx = context.TODO()
	lCollection := mongoClient.Database(constants.Dbname).Collection("loan")
	lService := service.InitLoan(lCollection, ctx)
	lController := controllers.InitLoanController(lService)
	routes.LoanRoute(server,lController)
}

func main(){
	server = gin.Default()
	mongoclient,err :=config.ConnectDataBase()
	defer   mongoclient.Disconnect(ctx)
	if err!=nil{
		panic(err)
	}
	initRoutes()
	initApp(mongoclient)
	initAcc(mongoclient)
	initBank(mongoclient)
	initLoan(mongoclient)
	fmt.Println("server running on port",constants.Port)
	log.Fatal(server.Run(constants.Port))
}

