package img

import (
	"gitee.com/youbeiwuhuan/knife4go/gin-openapi3-knife/constant"
	"gitee.com/youbeiwuhuan/knife4go/gin-openapi3-knife/utils"
	"github.com/gin-gonic/gin"
)

const (
	EDITORMD_LOGO_53EA80E2_SVG_RELATIVE_PATH = constant.ROOT_PATH + "/webjars/img/editormd-logo.53ea80e2.svg"
	// 文件内容的16进制表示
	EDITORMD_LOGO_53EA80E2_SVG_HEX_CONTENT = `3c3f786d6c2076657273696f6e3d22312e3022207374616e64616c6f6e653d226e6f223f3e0d0a3c21444f435459504520737667205055424c494320222d2f2f5733432f2f4454442053564720312e312f2f454e222022687474703a2f2f7777772e77332e6f72672f47726170686963732f5356472f312e312f4454442f73766731312e64746422203e0d0a3c73766720786d6c6e733d22687474703a2f2f7777772e77332e6f72672f323030302f737667223e0d0a3c6d657461646174613e47656e6572617465642062792049636f4d6f6f6e3c2f6d657461646174613e0d0a3c646566733e0d0a3c666f6e742069643d2269636f6d6f6f6e2220686f72697a2d6164762d783d2231303234223e0d0a3c666f6e742d6661636520756e6974732d7065722d656d3d22313032342220617363656e743d22393630222064657363656e743d222d363422202f3e0d0a3c6d697373696e672d676c79706820686f72697a2d6164762d783d223130323422202f3e0d0a3c676c79706820756e69636f64653d2226237832303b2220643d222220686f72697a2d6164762d783d2235313222202f3e0d0a3c676c79706820756e69636f64653d2226237865313938373b2220643d224d3732362e3935342036382e3233366c2d39312e3835352d35362e3331392d32312e353137203130362e373438203131332e3337312d35302e34337a4d3837362e323933203730392e3439336c31322e3530322032382e31303663362e3436392031342e3534352032332e3635392032312e3134372033382e3230312031342e3638316c36302e3635322d32362e3938346331342e3534362d362e3436382032312e3134392d32332e3636312031342e36382d33382e3230316c2d31322e3530322d32382e3130362d3131332e3533362035302e3530357a4d3732302e323336203432342e3437386c3131362e303431203236302e38362d372e3230392036392e303139682d3133302e3234386c2d3233332e3733362d3532322e3335382d3234352e343736203532322e333538682d3133332e3532386c2d37312e3236362d3734322e3434326838322e3436326c34372e373835203536322e343938203236342e3034372d3536322e3439386834332e3134316c3235322e3835203536322e3439382031352e31342d3134392e3933397a4d3736312e3839312031312e3931356c2d362e3036382036302e303934203131372e303330203236332e3039372033332e3735372d3332332e3139322d3134342e37313920302e3030317a4d3632312e363338203133372e3030376c3131332e35342d35302e353033203234362e343836203535342e3131312d3131332e3533362035302e3530362d3234362e3438392d3535342e3131347a2220686f72697a2d6164762d783d223130313722202f3e0d0a3c2f666f6e743e3c2f646566733e3c2f7376673e`
)

func AddRouterOfEditormdLogo53ea80e2Svg(router *gin.Engine) {

	utils.GetOther(router, EDITORMD_LOGO_53EA80E2_SVG_RELATIVE_PATH, EDITORMD_LOGO_53EA80E2_SVG_HEX_CONTENT)

}
