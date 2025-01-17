package router_v1

import (
	"github.com/gin-gonic/gin"
)

func Register(router *gin.Engine) {
	v1 := router.Group("/api/v1")

	RegisterCommonRouter(v1.Group(""))

	RegisterReportRouter(v1.Group("/rp"))

	RegisterAdminRouter(v1.Group("/rp/admin"))

	RegisterProcurementSearchRouter(v1.Group("/procurement"))
}
