package icons


import (
	"gitee.com/youbeiwuhuan/knife4go/gin-openapi3-knife/constant"
	"gitee.com/youbeiwuhuan/knife4go/gin-openapi3-knife/utils"
	"github.com/gin-gonic/gin"

)


const (
	MSAPPLICATION_ICON_144X144_PNG_RELATIVE_PATH = constant.ROOT_PATH + "/img/icons/msapplication-icon-144x144.png"
	// 文件内容的16进制表示
	MSAPPLICATION_ICON_144X144_PNG_HEX_CONTENT = `89504e470d0a1a0a0000000d4948445200000090000000900803000000d098128a000000097048597300000b1300000b1301009a9c1800000048504c544547704c6465665f746c5b89746a5660695b624454666662656667664fb384569b7b4db1844a9a7c4e5b66537c704346614dba874354664bcc8d4bc28a4141603f4e65456a6d4782758dd2aec30000001074524e53006086ad2840e8180bead1fcfeb6c8e95b45ba80000003d34944415478daed9aeb929b300c8583836308c92cc6b0fbfe6f5adaa6b922599664c34cd16f2e429c2fe768d9c361afbdf6da6bafbd52ab719c8a5d9575d1e6efa92d525f27a02abc9f0a3aefeb82dcedf698c72140357c8f7eb1a6133a23779a96cf1bbf919b1d6f27db76e8811aba09b8f264b0860c74d6d4c1f76aadbb69c88007f5c30f34225fc3fd580fd4f883dccadc34348fe80c1e167a0f7574851bba422fcccf1784fa395b77c7cc04f8b80bd0909f405d571334a00bfce4e13ea079443532a200ebba017e476045230f7eaeddd32f51d5714664d214ed2764405de59e9eafa98f88d8a011f965f431e4e19b1ceb977933d1bfa6281a433eb4d6bdd98709e9e87b6fd5917f8c08d3750afae0807ca020ffa46b1df479c857ee8358c7d47543461e55f4021e28fadd48b434d0c4c60e457ee9e70c1d11d1d26a3f3114fd863c05fdd9d248e8c38a464cacb50e488e98a591d0e720ff6262efba96a20fbb3cd1c43ed0ef45e8b390ef2b68405c4bbba75915137b47bf0dfc346b58265639742312a4591ef2c635e8b680a759147d16f267eb224b232fcdcee85b696ee5a08f589a3cb72aa3af8ebc34cd9ebc36f2b2340b1569558dbd340efa3ab995857e5247a3047921fa199027a10feb1af8e38b0079e122bba468f2aa2a419fae6b4e6ecd89fe8c7c20afaa715dcb47a481bc6c9115adaa515d8bd14f5d5563e8b32c4db4aa463a622db2a255352ffa7ac8cb1659d1aa2ab1b4d88878ab6a9634ab915b55d3ac4a6e554db32ab9151811d3d21472ab6e9ad5343195455623b76aa6599ddcaa877e16e425e8e7419e8f7e26e4d969361bf25cf4f321cf4bb3a3666ed548b3bab9559e66b5732b2bcdbe34a49c5ba569563fb772d2ec03fd69eab556559d45161b50304e6d40d4cf3293cf68621c4bcb6b62e9e897419e86fe6f5da39f58f490a7a35f0a79ea229b615595a5d99cb995857e39e489e817449ea4eb108a2a3abec8c29545d184349b39b7a6a7d9ccb995817e51e429e817459e847e51e449e897449e64694591e7a19f57d1e9e867459eb4c8665a55056936736e4d4fb39973ab0cfdfcc8a7a11f5a5b6640444b0b25904fb1b4cc26966c69d94d2c15fd52c83fd0c7753db455d17ea2963698922f2c8efe500e7912fa4591a7a05f147902fa85918f5b5a3913a32db2a1bca271f407b3c60b83d36c91dc9a80fe2ac863e8af823c926693fe1bb004fa2b210f5adad05ab7663f1f9f65d654f412fa2b22bf68692b99189c664be7d618fac556552afa6b23ff8e7e68ed1606f4407f7de45f2d6d5d13fb447f13c83feb7a2b8afe9766fbdebac361431d19b39d17f6a721576f694073474d73d86bafbdf6daebffa85fdc8607936645129b0000000049454e44ae426082`
)

func AddRouterOfMsapplicationIcon144x144Png(router *gin.Engine) {
    
	utils.GetOther(router, MSAPPLICATION_ICON_144X144_PNG_RELATIVE_PATH, MSAPPLICATION_ICON_144X144_PNG_HEX_CONTENT)
	
}







