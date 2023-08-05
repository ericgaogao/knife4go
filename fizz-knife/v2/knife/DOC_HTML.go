package knife


import (
	"gitee.com/youbeiwuhuan/knife4go/gin-swagger-knife/constant"
	"gitee.com/youbeiwuhuan/knife4go/gin-swagger-knife/utils"
	"github.com/gin-gonic/gin"

)


const (
	DOC_HTML_RELATIVE_PATH = constant.ROOT_PATH + "/doc.html"
	// 文件内容的16进制表示
	DOC_HTML_HEX_CONTENT = `<!DOCTYPE html><html lang=en><head><meta charset=utf-8><meta http-equiv=X-UA-Compatible content="IE=edge"><meta name=viewport content="width=device-width,initial-scale=1"><link rel=icon href=favicon.ico><title></title><link href=webjars/css/chunk-296622eb.20e6d994.css rel=prefetch><link href=webjars/css/chunk-d7d5f59c.a9ffbfcb.css rel=prefetch><link href=webjars/js/chunk-069eb437.2486111c.js rel=prefetch><link href=webjars/js/chunk-0d102d5a.144e3dee.js rel=prefetch><link href=webjars/js/chunk-0fd67716.d57e2c41.js rel=prefetch><link href=webjars/js/chunk-296622eb.e4e811f0.js rel=prefetch><link href=webjars/js/chunk-2d0af44e.edce7802.js rel=prefetch><link href=webjars/js/chunk-2d0bd799.d2d95c02.js rel=prefetch><link href=webjars/js/chunk-2d0da532.0b13c746.js rel=prefetch><link href=webjars/js/chunk-3b888a65.8737ce4f.js rel=prefetch><link href=webjars/js/chunk-3ec4aaa8.a79d19f8.js rel=prefetch><link href=webjars/js/chunk-589faee0.f3951d02.js rel=prefetch><link href=webjars/js/chunk-735c675c.ecb8aa58.js rel=prefetch><link href=webjars/js/chunk-adb9e944.55c41d5b.js rel=prefetch><link href=webjars/js/chunk-d7d5f59c.42190458.js rel=prefetch><link href=webjars/js/chunk-f876db6c.4965fec9.js rel=prefetch><link href=webjars/css/app.b27647f3.css rel=preload as=style><link href=webjars/css/chunk-vendors.f24a310a.css rel=preload as=style><link href=webjars/js/app.06135c03.js rel=preload as=script><link href=webjars/js/chunk-vendors.d59642c5.js rel=preload as=script><link href=webjars/css/chunk-vendors.f24a310a.css rel=stylesheet><link href=webjars/css/app.b27647f3.css rel=stylesheet></head><body><noscript><strong>We're sorry but knife4j-vue doesn't work properly without JavaScript enabled. Please enable it to continue.</strong></noscript><div id=app></div><script src=webjars/js/chunk-vendors.d59642c5.js></script><script src=webjars/js/app.06135c03.js></script></body></html>`
)

func AddRouterOfDocHtml(router *gin.Engine) {
    
    utils.GetHtml(router, DOC_HTML_RELATIVE_PATH, DOC_HTML_HEX_CONTENT)
    
}







