package css


import (
	"github.com/ericgaogao/knife4go/gin-swagger-knife/constant"
	"github.com/ericgaogao/knife4go/gin-swagger-knife/utils"
	"github.com/gin-gonic/gin"

)


const (
	CHUNK_1287E0BE_D6530ADC_CSS_RELATIVE_PATH = constant.ROOT_PATH + "/webjars/css/chunk-1287e0be.d6530adc.css"
	// 文件内容的16进制表示
	CHUNK_1287E0BE_D6530ADC_CSS_HEX_CONTENT = `.api-tab[data-v-c1c1f512]{margin-top:15px}.api-tab .ant-tag[data-v-c1c1f512]{height:32px;line-height:32px}.api-basic[data-v-c1c1f512]{padding:11px}.api-basic-title[data-v-c1c1f512]{font-size:14px;font-weight:700}.api-basic-body[data-v-c1c1f512]{font-size:14px;font-family:-webkit-body}.api-description[data-v-c1c1f512]{border-left:4px solid #ddd;line-height:30px}.api-body-desc[data-v-c1c1f512]{padding:10px;min-height:35px;-webkit-box-sizing:border-box;box-sizing:border-box;border:1px solid #e8e8e8}.ant-card-body[data-v-c1c1f512]{padding:5px}.api-title[data-v-c1c1f512]{margin-top:10px;margin-bottom:5px;font-size:16px;font-weight:600;height:30px;line-height:30px;border-left:4px solid #00ab6d;text-indent:8px}.content-line[data-v-c1c1f512]{height:25px;line-height:25px}.content-line-count[data-v-c1c1f512]{height:35px;line-height:35px}.divider[data-v-c1c1f512]{margin:4px 0}`
)

func AddRouterOfChunk1287e0beD6530adcCss(router *gin.Engine) {
    
    utils.GetCss(router, CHUNK_1287E0BE_D6530ADC_CSS_RELATIVE_PATH, CHUNK_1287E0BE_D6530ADC_CSS_HEX_CONTENT)
	
}







