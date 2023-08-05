package js

import (
	"gitee.com/youbeiwuhuan/knife4go/gin-openapi3-knife/constant"
	"gitee.com/youbeiwuhuan/knife4go/gin-openapi3-knife/utils"
	"github.com/gin-gonic/gin"
)

const (
	CHUNK_069EB437_2486111C_JS_LICENSE_TXT_RELATIVE_PATH = constant.ROOT_PATH + "/webjars/js/chunk-069eb437.2486111c.js.LICENSE.txt"
	// 文件内容的16进制表示
	CHUNK_069EB437_2486111C_JS_LICENSE_TXT_HEX_CONTENT = `2f2a2a0d0a202a20766b4265617574696679202d206a61766173637269707420706c7567696e20746f207072657474792d7072696e74206f72206d696e696679207465787420696e20584d4c2c204a534f4e2c2043535320616e642053514c20666f726d6174732e0d0a202a2068747470733a2f2f6769746875622e636f6d2f6161626c7565647261676f6e2f766b62656175746966790d0a202a20436f7079726967687420286329203230313220566164696d204b697279756b68696e0d0a202a20766b697279756b68696e204020676d61696c2e636f6d0d0a202a20687474703a2f2f7777772e65736c696e7374727563746f722e6e65742f766b62656175746966792f0d0a202a0d0a202a204475616c206c6963656e73656420756e64657220746865204d495420616e642047504c206c6963656e7365733a0d0a202a202020687474703a2f2f7777772e6f70656e736f757263652e6f72672f6c6963656e7365732f6d69742d6c6963656e73652e7068700d0a202a202020687474703a2f2f7777772e676e752e6f72672f6c6963656e7365732f67706c2e68746d6c0d0a202a0d0a202a202020507265747479207072696e740d0a202a0d0a202a2020202020202020766b62656175746966792e786d6c2874657874205b2c696e64656e745f7061747465726e5d293b0d0a202a2020202020202020766b62656175746966792e6a736f6e2874657874205b2c696e64656e745f7061747465726e5d293b0d0a202a2020202020202020766b62656175746966792e6373732874657874205b2c696e64656e745f7061747465726e5d293b0d0a202a2020202020202020766b62656175746966792e73716c2874657874205b2c696e64656e745f7061747465726e5d293b0d0a202a0d0a202a20202020202020204074657874202d20537472696e673b207465787420746f20626561747566793b0d0a202a202020202020202040696e64656e745f7061747465726e202d20496e7465676572207c20537472696e673b0d0a202a20202020202020202020202020202020496e74656765723a20206e756d626572206f66207768697465207370616365733b0d0a202a20202020202020202020202020202020537472696e673a20202063686172616374657220737472696e6720746f2076697375616c697a6520696e64656e746174696f6e20282063616e20616c736f206265206120736574206f662077686974652073706163657320290d0a202a2020204d696e6966790d0a202a0d0a202a2020202020202020766b62656175746966792e786d6c6d696e2874657874205b2c70726573657276655f636f6d6d656e74735d293b0d0a202a2020202020202020766b62656175746966792e6a736f6e6d696e2874657874293b0d0a202a2020202020202020766b62656175746966792e6373736d696e2874657874205b2c70726573657276655f636f6d6d656e74735d293b0d0a202a2020202020202020766b62656175746966792e73716c6d696e2874657874293b0d0a202a0d0a202a20202020202020204074657874202d20537472696e673b207465787420746f206d696e6966793b0d0a202a20202020202020204070726573657276655f636f6d6d656e7473202d20426f6f6c3b205b6f7074696f6e616c5d3b0d0a202a20202020202020202020202020202020536574207468697320666c616720746f207472756520746f2070726576656e742072656d6f76696e6720636f6d6d656e74732066726f6d2040746578742028206d696e786d6c20616e64206d696e6373732066756e6374696f6e73206f6e6c792e20290d0a202a0d0a202a2020204578616d706c65733a0d0a202a2020202020202020766b62656175746966792e786d6c2874657874293b202f2f20707265747479207072696e7420584d4c0d0a202a2020202020202020766b62656175746966792e6a736f6e28746578742c203420293b202f2f20707265747479207072696e74204a534f4e0d0a202a2020202020202020766b62656175746966792e63737328746578742c20272e202e202e202e27293b202f2f20707265747479207072696e74204353530d0a202a2020202020202020766b62656175746966792e73716c28746578742c20272d2d2d2d27293b202f2f20707265747479207072696e742053514c0d0a202a0d0a202a2020202020202020766b62656175746966792e786d6c6d696e28746578742c2074727565293b2f2f206d696e69667920584d4c2c20707265736572766520636f6d6d656e74730d0a202a2020202020202020766b62656175746966792e6a736f6e6d696e2874657874293b2f2f206d696e696679204a534f4e0d0a202a2020202020202020766b62656175746966792e6373736d696e2874657874293b2f2f206d696e696679204353532c2072656d6f766520636f6d6d656e747320282064656661756c7420290d0a202a2020202020202020766b62656175746966792e73716c6d696e2874657874293b2f2f206d696e6966792053514c0d0a202a0d0a202a2f0d0a`
)

func AddRouterOfChunk069eb4372486111cJsLICENSETxt(router *gin.Engine) {

	utils.GetOther(router, CHUNK_069EB437_2486111C_JS_LICENSE_TXT_RELATIVE_PATH, CHUNK_069EB437_2486111C_JS_LICENSE_TXT_HEX_CONTENT)

}
