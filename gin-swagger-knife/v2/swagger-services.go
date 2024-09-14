package v2

import (
	"github.com/ericgaogao/knife4go/gin-swagger-knife/constant"
	"github.com/ericgaogao/knife4go/gin-swagger-knife/utils"
	"github.com/gin-gonic/gin"
)

const (
	SWAGGER_SERVICES_RELATIVE_PATH = constant.ROOT_PATH + "/services.json"
)

func AddSwaggerServicesRouter(router *gin.Engine, services string) {
	utils.GetJson(router, SWAGGER_SERVICES_RELATIVE_PATH, services)
}
