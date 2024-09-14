package knife

import (
	v2 "github.com/ericgaogao/knife4go/gin-swagger-knife/v2"
	"github.com/gin-gonic/gin"

	css "github.com/ericgaogao/knife4go/gin-swagger-knife/v2/knife/webjars/css"

	fonts "github.com/ericgaogao/knife4go/gin-swagger-knife/v2/knife/webjars/fonts"

	icons "github.com/ericgaogao/knife4go/gin-swagger-knife/v2/knife/img/icons"

	img "github.com/ericgaogao/knife4go/gin-swagger-knife/v2/knife/webjars/img"

	js "github.com/ericgaogao/knife4go/gin-swagger-knife/v2/knife/webjars/js"

	knife "github.com/ericgaogao/knife4go/gin-swagger-knife/v2/knife"

	oauth "github.com/ericgaogao/knife4go/gin-swagger-knife/v2/knife/oauth"
	//oauth0 "github.com/ericgaogao/knife4go/gin-swagger-knife/v2/knife/webjars/oauth"
	//oauth1 "github.com/ericgaogao/knife4go/gin-swagger-knife/v2/knife/webjars/oauth"
)

// @param content string swag int 命令生成的swagger.json文件里的内容
func InitSwaggerKnife(router *gin.Engine, swaggerJson string) {
	v2.AddApiDocRouter(router, swaggerJson)
	v2.AddSwaggerResourcesRouter(router)

	knife.AddRouterOfDocHtml(router)

	knife.AddRouterOfFaviconIco(router)

	icons.AddRouterOfAndroidChrome192x192Png(router)

	icons.AddRouterOfAndroidChrome512x512Png(router)

	icons.AddRouterOfAppleTouchIcon120x120Png(router)

	icons.AddRouterOfAppleTouchIcon152x152Png(router)

	icons.AddRouterOfAppleTouchIcon180x180Png(router)

	icons.AddRouterOfAppleTouchIcon60x60Png(router)

	icons.AddRouterOfAppleTouchIcon76x76Png(router)

	icons.AddRouterOfAppleTouchIconPng(router)

	icons.AddRouterOfFavicon16x16Png(router)

	icons.AddRouterOfFavicon32x32Png(router)

	icons.AddRouterOfMsapplicationIcon144x144Png(router)

	icons.AddRouterOfMstile150x150Png(router)

	icons.AddRouterOfSafariPinnedTabSvg(router)

	oauth.AddRouterOfAxiosMinJs(router)

	oauth.AddRouterOfAxiosMinJsGz(router)

	oauth.AddRouterOfOauth2Html(router)

	knife.AddRouterOfRobotsTxt(router)

	css.AddRouterOfApp7c14bd7bCss(router)

	css.AddRouterOfApp7c14bd7bCssGz(router)

	css.AddRouterOfChunk1287e0beD6530adcCss(router)

	css.AddRouterOfChunkF19628f69612b987Css(router)

	css.AddRouterOfChunkVendorsF24a310aCss(router)

	css.AddRouterOfChunkVendorsF24a310aCssGz(router)

	fonts.AddRouterOfFontawesomeWebfont706450d7Ttf(router)

	fonts.AddRouterOfFontawesomeWebfont97493d3fWoff2(router)

	fonts.AddRouterOfFontawesomeWebfontD9ee23d5Woff(router)

	fonts.AddRouterOfFontawesomeWebfontF7c2b4b7Eot(router)

	fonts.AddRouterOfIconfont4ca3d0c0Ttf(router)

	fonts.AddRouterOfIconfontE2d2b98eEot(router)

	img.AddRouterOfEditormdLogo84b6c2a9Svg(router)

	img.AddRouterOfFontawesomeWebfont139e74e2Svg(router)

	img.AddRouterOfIconfontDd63dc33Svg(router)

	img.AddRouterOfLoadingC929501eGif(router)

	img.AddRouterOfLoading2x695405a9Gif(router)

	img.AddRouterOfLoading3x65eacf61Gif(router)

	js.AddRouterOfApp6f22e053Js(router)

	js.AddRouterOfApp6f22e053JsGz(router)

	js.AddRouterOfChunk0d102d5a6eea8501Js(router)

	js.AddRouterOfChunk0d102d5a6eea8501JsGz(router)

	js.AddRouterOfChunk0fd677167c05a478Js(router)

	js.AddRouterOfChunk0fd677167c05a478JsGz(router)

	js.AddRouterOfChunk1287e0beBc9d4cf1Js(router)

	js.AddRouterOfChunk1287e0beBc9d4cf1JsGz(router)

	js.AddRouterOfChunk214218f0E57955eaJs(router)

	js.AddRouterOfChunk214218f0E57955eaJsGz(router)

	js.AddRouterOfChunk232de5ac93e72917Js(router)

	js.AddRouterOfChunk232de5ac93e72917JsGz(router)

	js.AddRouterOfChunk2d0af44e88b13b03Js(router)

	js.AddRouterOfChunk2d0bd79908bfaf06Js(router)

	js.AddRouterOfChunk2d0da532Dc4ba7b3Js(router)

	js.AddRouterOfChunk3b888a654da69ee0Js(router)

	js.AddRouterOfChunk3b888a654da69ee0JsGz(router)

	js.AddRouterOfChunk42d52f4e32f2de93Js(router)

	js.AddRouterOfChunk42d52f4e32f2de93JsGz(router)

	js.AddRouterOfChunk589faee0Cdfbd926Js(router)

	js.AddRouterOfChunk589faee0Cdfbd926JsGz(router)

	js.AddRouterOfChunk6cbe5a9eE96d6f5dJs(router)

	js.AddRouterOfChunk6cbe5a9eE96d6f5dJsGz(router)

	js.AddRouterOfChunk735c675c28cc9be2Js(router)

	js.AddRouterOfChunk735c675c28cc9be2JsGz(router)

	js.AddRouterOfChunk9aa77acaFe25e739Js(router)

	js.AddRouterOfChunk9aa77acaFe25e739JsGz(router)

	js.AddRouterOfChunkA27fb06c5945d122Js(router)

	js.AddRouterOfChunkA27fb06c5945d122JsGz(router)

	js.AddRouterOfChunkAdb9e94400075de1Js(router)

	js.AddRouterOfChunkAdb9e94400075de1JsGz(router)

	js.AddRouterOfChunkF19628f65b2ed48bJs(router)

	js.AddRouterOfChunkF19628f65b2ed48bJsGz(router)

	js.AddRouterOfChunkVendors1f73f3b9Js(router)

	js.AddRouterOfChunkVendors1f73f3b9JsGz(router)
}
