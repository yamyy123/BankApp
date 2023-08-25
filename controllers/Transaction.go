package controllers

import (
	"BANKAPP/interfaces"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)


type TransactionController struct {
	TransactionService interfaces.ITransaction
}

func InitTransactionController(TransactionService interfaces.ITransaction) TransactionController {
	return TransactionController{TransactionService } //DI(dependency injection) pattern
}

func  (pc *TransactionController) Transfer(ctx *gin.Context) {

	type requestForm struct {
        From int `json:"from" bson:"from"`
        To int    `json:"to" bson:"to"`
		Amount int `json:"amount" bson:"amount"`
    }
	var requestData requestForm
	if err := ctx.ShouldBindJSON(&requestData); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
	

	results,err1:=pc.TransactionService.Transfer(requestData.From,requestData.To,requestData.Amount)
	if err1 != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail"})
		return
	}		
	for ind,val:=range results{
		ctx.JSON(http.StatusOK, gin.H{"index":ind,"value":val})
		fmt.Println(ind,val)
	   
	}
	
}