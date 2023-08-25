package routes

import (
	"BANKAPP/controllers"

	"github.com/gin-gonic/gin"
)


func Transactionroutes(router *gin.Engine, controller controllers.TransactionController) {
	router.POST("/api/customer/created", controller.Transfer)
	
}