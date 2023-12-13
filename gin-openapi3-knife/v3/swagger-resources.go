package v3

import (
	"gitee.com/youbeiwuhuan/knife4go/gin-openapi3-knife/constant"
	"gitee.com/youbeiwuhuan/knife4go/gin-openapi3-knife/utils"
	"github.com/gin-gonic/gin"
)

const (
	// TODO 路径要改
	SWAGGER_RESOURCES_CONTENT     = `{"configUrl": "/v3/api-docs/swagger-config","oauth2RedirectUrl": "` + constant.ROOT_PATH + `/swagger-ui/oauth2-redirect.html","url": "/v3/api-docs","validatorUrl": ""}`
	SWAGGER_RESOURCES_CONFIG_PATH = constant.ROOT_PATH + "/swagger-config"
)

func AddSwaggerResourcesRouter(router *gin.Engine) {
	utils.GetJson(router, SWAGGER_RESOURCES_CONFIG_PATH, SWAGGER_RESOURCES_CONTENT)
}
