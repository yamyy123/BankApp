package interfaces

import (
	"BANKAPP/models"

	"go.mongodb.org/mongo-driver/bson"
	//"time"
	//"go.mongodb.org/mongo-driver/mongo"
)



type ICustomer interface{
	CreateCustomer(customer *models.Customer)(error)
	DeleteService(filter bson.M)(error)
	UpdateService(filter bson.M,update bson.M)(error)
	FindService(filter bson.M)([] models.Customer,error)
	// FindServiceSum(filter mongo.Pipeline) (float64, error)
	
}