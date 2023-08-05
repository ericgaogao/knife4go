package fonts

import (
	"gitee.com/youbeiwuhuan/knife4go/gin-openapi3-knife/constant"
	"gitee.com/youbeiwuhuan/knife4go/gin-openapi3-knife/utils"
	"github.com/gin-gonic/gin"
)

const (
	ICONFONT_E2D2B98E_EOT_RELATIVE_PATH = constant.ROOT_PATH + "/webjars/fonts/iconfont.e2d2b98e.eot"
	// 文件内容的16进制表示
	ICONFONT_E2D2B98E_EOT_HEX_CONTENT = `f01200004812000001000200000000000200050300000000000001009001000000004c500000000000000000000000000000000001000000000000008572812a0000000000000000000000000000000000001000690063006f006e0066006f006e00740000000e0052006500670075006c006100720000001600560065007200730069006f006e00200031002e00300000001000690063006f006e0066006f006e0074000000000000010000000b00800003003047535542b0feb3ed00000138000000424f532f323cbb49250000017c00000056636d6170025f06320000020000000236676c796620258bcd0000045000000adc6865616417a20069000000e00000003668686561085b040c000000bc00000024686d74782c800000000001d40000002c6c6f63610cb80f4600000438000000186d617870011e00ac00000118000000206e616d653e54fe7d00000f2c0000026d706f737470faaeb00000119c000000a9000100000380ff80005c048000000000047d00010000000000000000000000000000000b00010000000100002a8172855f0f3cf5000b040000000000da02de5800000000da02de580000ff3c047d038000000008000200000000000000010000000b00a000090000000000020000000a000a000000ff00000000000000010000000a001e002c000144464c54000800040000000000000001000000016c69676100080000000100000001000400040000000100080001000600000001000000000001040c019000050008028902cc0000008f028902cc000001eb0032010800000200050300000000000000000000000000000000000000000000506645640040e601e7270380ff80005c038000c400000001000000000000040000000400000004000000040000000400000004000000040000000400000004000000048000000400000000000005000000030000002c00000004000001ae00010000000000a8000300010000002c0003000a000001ae0004007c00000016001000030006e601e60ee619e62ae650e652e668e6a4e6f5e727ffff0000e601e60ee619e62ae650e652e668e6a4e6f5e727ffff00000000000000000000000000000000000000000001001600160016001600160016001600160016001600000004000a0002000300080009000500070006000100000106000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000c000000000088000000000000000a0000e6010000e601000000040000e60e0000e60e0000000a0000e6190000e619000000020000e62a0000e62a000000030000e6500000e650000000080000e6520000e652000000090000e6680000e668000000050000e6a40000e6a4000000070000e6f50000e6f5000000060000e7270000e7270000000100000000000000b200f6014a0194023802de037e04700506056e00080000ff3c03a7031c0005000b0026002e003d005b006c0070000001110901110901250511052501321721363332171406232227011406222635343637272635343617060717321f0125052e0134363337330716151407170705321614062322270714161514062226353437273717321737343637273703321615140717232e01270723372634361327331703a6fe5cfe5801a8017afe86fe82017e017afd52210d02020b2127101c1b150efe0d1f2d1f110c0c201f4709120e04020701d7feaf151a1b1c06300e0a0f0526016a15181a1b0b0ca9021f2c1f133a273c0b08a704094326b0151f093c33091a0f252c280f1d821233050211fe36fef5010b01ca010bfeddefeffe66efef01a019192c1a220dfeeb151f1d15101908dc1023151d46130adc0204ff87021f281f14250f121b0c17153e1f2b1f067c04060215212115190ec213c7087e0f1908881501531f150e0f741338105b6c0e2c1ffed4231000000000020000ff8203cf02be001c002600001301363217011e010e012f01110e0107212e0135110706272227263637011533353626270e01153f01871a431a01840a0310180b0c023f27fd9c252b0e090a0f0908030a0166b7012d2f302c017701331414fecd08181503080afe8d28330101312a01740b07010b0b1808fe8940ab1c2d01012d1c00000000040000ff8004000380000b00180021003000000526002736003716001706001323353626270e011715231121011e01071521352636131e01151406071523352e013534360200d9fedf06060121d9d901210606fedf464807439b9b430748023efee1683503fecc0335681b22100e3e0e102280060121d9d901210606fedfd9d9fedf02605d0dbc0e0ebc0d5dfe67028f0779195d5d1979fe9701221b111c086d6d081c111b220000020000ffb703c703490014002900000901262207010e011f011e0137090116363f01362625011114163b013236371133111e013b013236371103b6fe5f09150afe5d0c050802081b0c018f018d0c1b08020805fe43fe84140edb0f1301ae01130fdb0f1301021d01250606fede091b0c020b05080114fee908050b020c1bb0fefdfe190e14140e0118fee80e14140e01e70000000700000000040002c7001e0043004b004e005e0069006d0000351e011f01213705373e01372e01272e01270e01072e01230e010714170e01372726353e013732161f01373e01371e011715171e01150e01072131233121272e012734361733373307333723172337252e012b010733373332373e04070e022b013733321e013707333701594620013d06013d2046590101463905ab805b902612311c3649010b303a9f0f06012e280e1c0a322a2573456599052d272e024830fea906fec515304f023b99271c490123032a05363901080516134a47261b1818090f14120f0528020e161d0c141b13090567472646f447640d020101020d64473e5e147ca30301594b13160147341b1817585f280c0d1b36020c0a23393741010492602c0c0e3d272f570a0209562f2558832828b6704a120b09b6450103080f1a1d120a10083504082ab6b6000000080000ff800380038000130019002b003900470055006300700000012122060f010e0115111e013321323637112e01051514062b01010e01072122263511333e013d01213216150723220607141617333236352e010721220607141633213236352e010721220607141617213236352e010721220607141617213236352e01072122061514163321323634260340fe460e1b0ac0090a012619026d1a37020126fe0d110f8c02a6011e14fd930f11a61e2201ba0f1186e0060d010a0ae6050e040c09fe4c050d010a0901ba050e040c09fe4c050d010a0901ba050e040c09fe4c050d010a0901ba050e040c09fe46050e090a01b3060d070380100ac00a180afd461d23231d03801926398c0f11fd260a1501110f02ba012619a6110f9a090a050d010a090a0993090a050e090a0a09a60a09060d010a0a090a930a0a050d010a090a0aa7090a050e09140900090000ff8803a5037800140026002c003800450051005d005e0067000025353f01132e0127210e0107111e011721323f01010522263511343637211e011511230e01071537073534363701213e013426232122061416052e0123212206141617213e0107230e0114163b013236342603230e0114163b013236342605231416323634262206039c040401014030fda530400101403101970a09050112fd3e0e13130e025b0e139b304001d787130efe81016711161611fe9911171701a0011710fe991117171101671116bdd111171711d111161603c206090906c2070808ff000f080d08080d08b201040d0242304001014030fcf430400104040120d8130e030c0e130101130efdee024030a9cb8f6d0e130101bd011622171722168311171722160101167101162217172216feba01080d08080d080e0708080d0808000000080000ff7f04000380005b00670089008f00930097009b009f00002507161407171e010e012f01060717160e01262f010607150e0122263d012627070e012e013f01262707062e01363f01263437272e013e011f01363727263e01161f01363735343632161d011617373e011e010f01161737361e0106250e01071e01173e01372e0101212e01351134363321151e0117331516173527370121220e0215111e011721261317232226350533352337231533033335232521152103e22902032a0f0f0b1b102a0d13190a061a1e0b191a1b011621161c19190b1e1b05091a130e29101c0a0f0f29020229100e0a1c102a0d13190a051b1e0b19191c1622161b1a190b1e1b050a19130e29101b0b0efef33a4d01014d3a3a4d01014dfe9bfeaa151d1c1601bf01392ac71a180901fecdfe1f14251c0f01382b01861b86daa8151dfe59f9f9f9f9f9f9f9f90154feac0154bf0c0d1c0d0c051c1e0d040d1714220c1f12050c220c05291015151029050c220c05121f0c2214170d040d1e1c040d0d1c0d0c061a1e0e040e1814220c1f12050d210c05291015151029050c210d05131d0d2214170d040d1f1a56024a37384902024938374afee7011d1402ea161ce02a3801250d135e0809014b0f1c2514fd162a3801160340e81d1495329532fe2a32953200070000ffb3047d03480035003d004100490051005500620000010e01070e01071e0117333236342627232e01273e013f023e01371e011f01331e01170e0107230e0114163b013e01372e01272e01013313232723072337332723373332142b011523371533323634262337113311033e012e010f0117163e012f01020e7bb41c57680104b5891b0f15150f1b6a8d0301594a1403119265618f1506226a8d03038d6a950f14140f9588b60403a88022aefef130612d176b172d51522802936a5d5d402a2a3d1c1a1a1c842a4d08020d13098d8e0b1c030c570348028b7025955d82ac03141c13010286654d771b07145e7602016e591b02866565860201131d1303ac827ca90b667cfe98ff00404063702d9e62dc56142e1424ff000100fe8b07141003066c68090b1e094000020000ff8c03ff03740042004800000114062b01140717161406222f010e0423112311222e022f01070623222e013f01263523222634363b013527263436321f012137363216140f01153332161501213436321603fe18118f2a840c17230b7e030d292a3f1e5220412e260909750d120f18020b81258f111818118f6f0c18220c6e021a6e0c22180c6f8f1118fecefe6877aa77013911186d4c850c22180c7e030b1a1411023cfdc41218190909840e16220c914966182118bc6e0c21190c6f6f0c19210c6ebc1810016f557777000000001200de00010000000000000015000000010000000000010008001500010000000000020007001d00010000000000030008002400010000000000040008002c0001000000000005000b003400010000000000060008003f000100000000000a002b0047000100000000000b001300720003000104090000002a00850003000104090001001000af0003000104090002000e00bf0003000104090003001000cd0003000104090004001000dd0003000104090005001600ed000300010409000600100103000300010409000a00560113000300010409000b002601690a437265617465642062792069636f6e666f6e740a69636f6e666f6e74526567756c617269636f6e666f6e7469636f6e666f6e7456657273696f6e20312e3069636f6e666f6e7447656e65726174656420627920737667327474662066726f6d20466f6e74656c6c6f2070726f6a6563742e687474703a2f2f666f6e74656c6c6f2e636f6d000a0043007200650061007400650064002000620079002000690063006f006e0066006f006e0074000a00690063006f006e0066006f006e00740052006500670075006c0061007200690063006f006e0066006f006e007400690063006f006e0066006f006e007400560065007200730069006f006e00200031002e003000690063006f006e0066006f006e007400470065006e00650072006100740065006400200062007900200073007600670032007400740066002000660072006f006d00200046006f006e00740065006c006c006f002000700072006f006a006500630074002e0068007400740070003a002f002f0066006f006e00740065006c006c006f002e0063006f006d0000000002000000000000000a00000000000000000000000000000000000000000000000b01020103010401050106010701080109010a010b010c00086d6f64656c696e6704686f6d651461757468656e7469636174696f6e73797374656d057a687579650a41504977656e64616e670777656e64616e671974756269616f7a68697a756f6d6f62616e796968756966752d057a646c78620b41504977656e64616e67310564656275670000000000`
)

func AddRouterOfIconfontE2d2b98eEot(router *gin.Engine) {

	utils.GetOther(router, ICONFONT_E2D2B98E_EOT_RELATIVE_PATH, ICONFONT_E2D2B98E_EOT_HEX_CONTENT)

}
