package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getOrderController(c *gin.Context) {
	var orders []OrderDto = makeOrderFactory(30)
	c.JSON(http.StatusOK, orders)
}
