package css


import (
	"github.com/ericgaogao/knife4go/gin-swagger-knife/constant"
	"github.com/ericgaogao/knife4go/gin-swagger-knife/utils"
	"github.com/gin-gonic/gin"

)


const (
	CHUNK_F19628F6_9612B987_CSS_RELATIVE_PATH = constant.ROOT_PATH + "/webjars/css/chunk-f19628f6.9612b987.css"
	// 文件内容的16进制表示
	CHUNK_F19628F6_9612B987_CSS_HEX_CONTENT = `.api-title[data-v-d082656a]{margin-top:10px;margin-bottom:5px;font-size:16px;font-weight:600;height:30px;line-height:30px;border-left:4px solid #00ab6d;text-indent:8px}`
)

func AddRouterOfChunkF19628f69612b987Css(router *gin.Engine) {
    
    utils.GetCss(router, CHUNK_F19628F6_9612B987_CSS_RELATIVE_PATH, CHUNK_F19628F6_9612B987_CSS_HEX_CONTENT)
	
}







