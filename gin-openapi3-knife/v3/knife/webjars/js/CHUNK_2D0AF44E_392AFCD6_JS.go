package js


import (
	"gitee.com/youbeiwuhuan/knife4go/gin-openapi3-knife/constant"
	"gitee.com/youbeiwuhuan/knife4go/gin-openapi3-knife/utils"
	"github.com/gin-gonic/gin"

)


const (
	CHUNK_2D0AF44E_392AFCD6_JS_RELATIVE_PATH = constant.ROOT_PATH + "/webjars/js/chunk-2d0af44e.392afcd6.js"
	// 文件内容的16进制表示
	CHUNK_2D0AF44E_392AFCD6_JS_HEX_CONTENT = `2877696e646f772e7765627061636b4a736f6e703d77696e646f772e7765627061636b4a736f6e707c7c5b5d292e70757368285b5b226368756e6b2d3264306166343465225d2c7b2230653336223a66756e6374696f6e28652c742c69297b2275736520737472696374223b692e722874293b766172206e3d7b6e616d653a22456469746f7253686f77222c636f6d706f6e656e74733a7b656469746f723a6928223763396522297d2c70726f70733a7b76616c75653a7b747970653a537472696e672c72657175697265643a21312c64656661756c743a22227d2c786d6c4d6f64653a7b747970653a426f6f6c65616e2c64656661756c743a21312c72657175697265643a21317d7d2c646174613a66756e6374696f6e28297b72657475726e7b6c616e673a226a736f6e222c656469746f723a6e756c6c2c656469746f724865696768743a3230307d7d2c6d6574686f64733a7b6368616e67653a66756e6374696f6e2865297b746869732e24656d697428226368616e6765222c65297d2c7265736574456469746f724865696768743a66756e6374696f6e28297b76617220653d746869733b73657454696d656f7574282866756e6374696f6e28297b76617220743d652e656469746f722e73657373696f6e2e6765744c656e67746828293b313d3d74262628743d3130293b76617220693d31362a743b652e656469746f724865696768743d697d292c333030297d2c656469746f72496e69743a66756e6374696f6e2865297b76617220743d746869733b746869732e656469746f723d652c6928223230393922292c6928223831386222292c6928223036393622292c746869732e786d6c4d6f6465262628746869732e6c616e673d22786d6c22292c6928223164323922292c746869732e7265736574456469746f7248656967687428292c746869732e656469746f722e72656e64657265722e6f6e2822616674657252656e646572222c2866756e6374696f6e28297b742e24656d6974282273686f774465736372697074696f6e222c2231323322297d29297d7d7d2c6f3d6928223238373722292c723d4f626a656374286f2e6129286e2c2866756e6374696f6e28297b76617220653d746869732c743d652e24637265617465456c656d656e742c693d652e5f73656c662e5f637c7c743b72657475726e20692822646976222c5b692822656469746f72222c7b61747472733a7b76616c75653a652e76616c75652c6c616e673a652e6c616e672c7468656d653a2265636c69707365222c77696474683a2231303025222c6865696768743a652e656469746f724865696768747d2c6f6e3a7b696e69743a652e656469746f72496e69742c696e7075743a652e6368616e67657d7d295d2c31297d292c5b5d2c21312c6e756c6c2c6e756c6c2c6e756c6c293b742e64656661756c743d722e6578706f7274737d7d5d293b`
)

func AddRouterOfChunk2d0af44e392afcd6Js(router *gin.Engine) {
    
    utils.GetJs(router, CHUNK_2D0AF44E_392AFCD6_JS_RELATIVE_PATH, CHUNK_2D0AF44E_392AFCD6_JS_HEX_CONTENT)
    
}







