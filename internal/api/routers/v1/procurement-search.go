package router_v1

import (
	"web-api/internal/api/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterProcurementSearchRouter(router *gin.RouterGroup) {
	router.GET("/get", controllers.Procurement.GetAll)
	router.GET("/average", controllers.Procurement.GetAverage)
}
