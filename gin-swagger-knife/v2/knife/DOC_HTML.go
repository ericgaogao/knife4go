package knife


import (
	"github.com/ericgaogao/knife4go/gin-swagger-knife/constant"
	"github.com/ericgaogao/knife4go/gin-swagger-knife/utils"
	"github.com/gin-gonic/gin"

)


const (
	DOC_HTML_RELATIVE_PATH = constant.ROOT_PATH + "/doc.html"
	// 文件内容的16进制表示
	DOC_HTML_HEX_CONTENT = `<!DOCTYPE html><html lang="en"><head><meta charset="utf-8"><meta http-equiv="X-UA-Compatible" content="IE=edge"><meta name="viewport" content="width=device-width,initial-scale=1"><title></title><link href="webjars/css/chunk-1287e0be.d6530adc.css" rel="prefetch"><link href="webjars/css/chunk-f19628f6.9612b987.css" rel="prefetch"><link href="webjars/js/chunk-0d102d5a.41a24621.js" rel="prefetch"><link href="webjars/js/chunk-0fd67716.7c05a478.js" rel="prefetch"><link href="webjars/js/chunk-1287e0be.75c81db6.js" rel="prefetch"><link href="webjars/js/chunk-214218f0.e57955ea.js" rel="prefetch"><link href="webjars/js/chunk-232de5ac.484ce381.js" rel="prefetch"><link href="webjars/js/chunk-2d0af44e.705393a6.js" rel="prefetch"><link href="webjars/js/chunk-2d0bd799.fc5aa0a6.js" rel="prefetch"><link href="webjars/js/chunk-2d0da532.6221417b.js" rel="prefetch"><link href="webjars/js/chunk-3b888a65.4da69ee0.js" rel="prefetch"><link href="webjars/js/chunk-42d52f4e.32f2de93.js" rel="prefetch"><link href="webjars/js/chunk-589faee0.6360fbf7.js" rel="prefetch"><link href="webjars/js/chunk-6cbe5a9e.e96d6f5d.js" rel="prefetch"><link href="webjars/js/chunk-735c675c.82070437.js" rel="prefetch"><link href="webjars/js/chunk-9aa77aca.74ceb75d.js" rel="prefetch"><link href="webjars/js/chunk-a27fb06c.5945d122.js" rel="prefetch"><link href="webjars/js/chunk-adb9e944.bfa2cd65.js" rel="prefetch"><link href="webjars/js/chunk-f19628f6.5c66c88c.js" rel="prefetch"><link href="webjars/css/app.7c14bd7b.css" rel="preload" as="style"><link href="webjars/css/chunk-vendors.f24a310a.css" rel="preload" as="style"><link href="webjars/js/app.5b8b5fd8.js" rel="preload" as="script"><link href="webjars/js/chunk-vendors.1f73f3b9.js" rel="preload" as="script"><link href="webjars/css/chunk-vendors.f24a310a.css" rel="stylesheet"><link href="webjars/css/app.7c14bd7b.css" rel="stylesheet"></head><body><noscript><strong>Please ensure that the static resources of Knife4j are released to avoid being intercepted by the back-end security framework!. Enable it to continue.</strong></noscript><div id="app"></div><script src="webjars/js/chunk-vendors.1f73f3b9.js"></script><script src="webjars/js/app.5b8b5fd8.js"></script></body></html>`
)

func AddRouterOfDocHtml(router *gin.Engine) {
    
    utils.GetHtml(router, DOC_HTML_RELATIVE_PATH, DOC_HTML_HEX_CONTENT)
    
}







