package icons


import (
	"gitee.com/youbeiwuhuan/knife4go/gin-swagger-knife/constant"
	"gitee.com/youbeiwuhuan/knife4go/gin-swagger-knife/utils"
	"github.com/gin-gonic/gin"

)


const (
	FAVICON_16X16_PNG_RELATIVE_PATH = constant.ROOT_PATH + "/img/icons/favicon-16x16.png"
	// 文件内容的16进制表示
	FAVICON_16X16_PNG_HEX_CONTENT = `89504e470d0a1a0a0000000d4948445200000010000000100803000000282d0f530000000467414d410000b18f0bfc6105000000206348524d00007a26000080840000fa00000080e8000075300000ea6000003a98000017709cba513c000000bd504c54454354664353664dba870000004dba874db886456d6e4352654354664354664dba874dba874dbb874888774353664354664354664dba874dba874354664354664dba874dba874354664dba874dba874354664dba874dba874dba874dba874dba874dba874dba874dba874dba874dba874dba874dba874dba874dba874dba874dba874dba874dba874dbb874cac8245646b4353664dba87498f794355664354664cb385466e6e4a9c7d4359674db8864779724ba681445e69488676ffffff0fb1e5500000002d74524e53000000001919191918079ed6d4d4d4d4644cecd62104928c28dcf474fd16c555f50aaa3ae9028b236bfc12c145e3d0e08c00000001624b47443e496400e30000000774494d4507e10113070b364701e991000000c34944415418d34dccd91a81601485e16dfd99ca3c64888ca1d8fc9b4286fbbf2d291ed6d9f71e2c328c7ca1582a13954bc542de30c8b42ad55abd91cb35eab56ac532a9d9e2fd41b73b9db63eecb9d5a4ae2d7c3c85bd5e783ab2d85d427fc0d1590f87fa1cf1a00f029c115fae6178bdf0c8014861ec4a7cd3fa168b3b864a009329df1f8f3b4f277883c26c2ef27c8acc67502960b1942892e5021928782b49b6f2924e41616d8958eb777f001bdfdfe0070ac1761ba4fd05ec76f887e4d7cb5abd00156e15d7d8c5461d0000002574455874646174653a63726561746500323031372d30312d31395430373a31313a35342b30313a3030c5b3dbfe0000002574455874646174653a6d6f6469667900323031372d30312d31395430373a31313a35342b30313a3030b4ee6342000000577a5458745261772070726f66696c65207479706520697074630000789ce3f20c0871562828ca4fcbcc49e5520003230b2e630b1323134b9314031320448034c3640323b35420cbd8d4c8c4ccc41cc407cb8048a04a2e00ea171174f24235950000000049454e44ae426082`
)

func AddRouterOfFavicon16x16Png(router *gin.Engine) {
    
	utils.GetOther(router, FAVICON_16X16_PNG_RELATIVE_PATH, FAVICON_16X16_PNG_HEX_CONTENT)
	
}







