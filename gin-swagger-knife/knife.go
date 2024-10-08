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
)

type InitArgs struct {
	SwaggerJson string // swagger.json文件里的内容
	ServiceJson string // service.json文件里的内容
}

// @param content string swag int 命令生成的swagger.json文件里的内容
func InitSwaggerKnife(router *gin.Engine, args *InitArgs) {
	v2.AddApiDocRouter(router, args.SwaggerJson)
	v2.AddSwaggerServicesRouter(router, args.ServiceJson)
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

	js.AddRouterOfApp5b8b5fd8Js(router)

	js.AddRouterOfApp5b8b5fd8JsGz(router)

	js.AddRouterOfChunk0d102d5a41a24621Js(router)

	js.AddRouterOfChunk0d102d5a41a24621JsGz(router)

	js.AddRouterOfChunk0fd677167c05a478Js(router)

	js.AddRouterOfChunk0fd677167c05a478JsGz(router)

	js.AddRouterOfChunk1287e0be75c81db6Js(router)

	js.AddRouterOfChunk1287e0be75c81db6JsGz(router)

	js.AddRouterOfChunk214218f0E57955eaJs(router)

	js.AddRouterOfChunk214218f0E57955eaJsGz(router)

	js.AddRouterOfChunk232de5ac484ce381Js(router)

	js.AddRouterOfChunk232de5ac484ce381JsGz(router)

	js.AddRouterOfChunk2d0af44e705393a6Js(router)

	js.AddRouterOfChunk2d0bd799Fc5aa0a6Js(router)

	js.AddRouterOfChunk2d0da5326221417bJs(router)

	js.AddRouterOfChunk3b888a654da69ee0Js(router)

	js.AddRouterOfChunk3b888a654da69ee0JsGz(router)

	js.AddRouterOfChunk42d52f4e32f2de93Js(router)

	js.AddRouterOfChunk42d52f4e32f2de93JsGz(router)

	js.AddRouterOfChunk589faee06360fbf7Js(router)

	js.AddRouterOfChunk589faee06360fbf7JsGz(router)

	js.AddRouterOfChunk6cbe5a9eE96d6f5dJs(router)

	js.AddRouterOfChunk6cbe5a9eE96d6f5dJsGz(router)

	js.AddRouterOfChunk735c675c82070437Js(router)

	js.AddRouterOfChunk735c675c82070437JsGz(router)

	js.AddRouterOfChunk9aa77aca74ceb75dJs(router)

	js.AddRouterOfChunk9aa77aca74ceb75dJsGz(router)

	js.AddRouterOfChunkA27fb06c5945d122Js(router)

	js.AddRouterOfChunkA27fb06c5945d122JsGz(router)

	js.AddRouterOfChunkAdb9e944Bfa2cd65Js(router)

	js.AddRouterOfChunkAdb9e944Bfa2cd65JsGz(router)

	js.AddRouterOfChunkF19628f65c66c88cJs(router)

	js.AddRouterOfChunkF19628f65c66c88cJsGz(router)

	js.AddRouterOfChunkVendors1f73f3b9Js(router)

	js.AddRouterOfChunkVendors1f73f3b9JsGz(router)

}
