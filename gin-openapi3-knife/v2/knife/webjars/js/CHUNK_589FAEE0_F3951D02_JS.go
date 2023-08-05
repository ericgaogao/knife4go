package js


import (
	"gitee.com/youbeiwuhuan/knife4go/gin-swagger-knife/constant"
	"gitee.com/youbeiwuhuan/knife4go/gin-swagger-knife/utils"
	"github.com/gin-gonic/gin"

)


const (
	CHUNK_589FAEE0_F3951D02_JS_RELATIVE_PATH = constant.ROOT_PATH + "/webjars/js/chunk-589faee0.f3951d02.js"
	// 文件内容的16进制表示
	CHUNK_589FAEE0_F3951D02_JS_HEX_CONTENT = `2f2a2120466f72206c6963656e736520696e666f726d6174696f6e20706c6561736520736565206368756e6b2d35383966616565302e66333935316430322e6a732e4c4943454e53452e747874202a2f0d0a2877696e646f772e7765627061636b4a736f6e703d77696e646f772e7765627061636b4a736f6e707c7c5b5d292e70757368285b5b226368756e6b2d3538396661656530225d2c7b2233633630223a66756e6374696f6e28652c742c6e297b2275736520737472696374223b6e2e722874293b766172206f3d286e28223431363022292c6e28226131356222292c6e28226236383022292c6e28226433623722292c6e28226163316622292c6e28223235663022292c6e28223533313922292c6e28223135396222292c6e2822623163372229292c723d6e28226233313122292c693d6e2e6e2872292c733d7b70726f70733a7b6170693a7b747970653a4f626a6563742c72657175697265643a21307d2c73776167676572496e7374616e63653a7b747970653a4f626a6563742c72657175697265643a21307d2c646562756753656e643a7b747970653a426f6f6c65616e2c64656661756c743a21317d2c726573706f6e7365486561646572733a7b747970653a41727261797d2c726573706f6e7365526177546578743a7b747970653a537472696e672c64656661756c743a22227d2c726573706f6e73654375726c546578743a7b747970653a537472696e672c64656661756c743a22227d2c726573706f6e73655374617475733a7b747970653a4f626a6563747d2c726573706f6e7365436f6e74656e743a7b747970653a4f626a6563747d2c726573706f6e73654669656c644465736372697074696f6e436865636b65643a7b747970653a426f6f6c65616e2c64656661756c743a21307d7d2c636f6d706f6e656e74733a7b456469746f72446562756753686f773a66756e6374696f6e28297b72657475726e2050726f6d6973652e616c6c285b6e2e6528226368756e6b2d336238383861363522292c6e2e6528226368756e6b2d306664363737313622292c6e2e6528226368756e6b2d336563346161613822292c6e2e6528226368756e6b2d373335633637356322295d292e7468656e286e2e62696e64286e756c6c2c22623365652229297d7d2c646174613a66756e6374696f6e28297b72657475726e7b706167696e6174696f6e3a21312c6931386e3a6e756c6c2c626173653634496d6167653a21312c6465627567526573706f6e73653a21302c726573706f6e7365486561646572436f6c756d6e3a5b5d7d7d2c77617463683a7b6c616e67756167653a66756e6374696f6e28652c74297b746869732e696e69744931386e28297d7d2c636f6d70757465643a7b6c616e67756167653a66756e6374696f6e28297b72657475726e20746869732e2473746f72652e73746174652e676c6f62616c732e6c616e67756167657d2c726573706f6e736553697a65546578743a66756e6374696f6e28297b76617220653d22302042222c743d746869732e726573706f6e73655374617475733b6966286e756c6c213d7426266e756c6c213d74297b766172206e3d742e73697a652c6f3d286e2f31303234292e746f46697865642832292c723d286e2f313032342f31303234292e746f46697865642832293b653d6f3e313f6f2b22204b42223a723e313f722b22204d42223a6e2b222042227d72657475726e20657d7d2c637265617465643a66756e6374696f6e28297b746869732e696e69744931386e28292c746869732e636f70795261775465787428292c746869732e636f70794375726c5465787428297d2c6d6574686f64733a7b67657443757272656e744931386e496e7374616e63653a66756e6374696f6e28297b72657475726e20746869732e246931386e2e6d657373616765735b746869732e6c616e67756167655d7d2c626173653634496e69743a66756e6374696f6e28297b76617220653d6f2e612e67657456616c756528746869732e726573706f6e7365436f6e74656e742c22626173653634222c22222c2130293b6f2e612e7374724e6f74426c616e6b286529262628746869732e626173653634496d6167653d2130297d2c696e69744931386e3a66756e6374696f6e28297b746869732e6931386e3d746869732e67657443757272656e744931386e496e7374616e636528292c746869732e726573706f6e7365486561646572436f6c756d6e3d746869732e6931386e2e7461626c652e6465627567526573706f6e7365486561646572436f6c756d6e737d2c636f7079526177546578743a66756e6374696f6e28297b76617220653d746869732c743d2262746e4465627567436f7079526177222b746869732e6170692e69642c6e3d6e657720692e61282223222b742c7b746578743a66756e6374696f6e28297b72657475726e20652e726573706f6e7365526177546578747d7d292c6f3d746869732e6931386e2e6d6573736167652e636f70792e7261772e737563636573732c723d746869732e6931386e2e6d6573736167652e636f70792e7261772e6661696c3b6e2e6f6e282273756363657373222c2866756e6374696f6e2874297b652e246d6573736167652e696e666f286f297d29292c6e2e6f6e28226572726f72222c2866756e6374696f6e2874297b652e246d6573736167652e696e666f2872297d29297d2c636f70794375726c546578743a66756e6374696f6e28297b76617220653d746869732c743d2262746e4465627567436f70794375726c222b746869732e6170692e69642c6e3d6e657720692e61282223222b742c7b746578743a66756e6374696f6e28297b72657475726e20652e726573706f6e73654375726c546578747d7d292c6f3d746869732e6931386e2e6d6573736167652e636f70792e6375726c2e737563636573732c723d746869732e6931386e2e6d6573736167652e636f70792e6375726c2e6661696c3b6e2e6f6e282273756363657373222c2866756e6374696f6e2874297b652e246d6573736167652e696e666f286f297d29292c6e2e6f6e28226572726f72222c2866756e6374696f6e2874297b652e246d6573736167652e696e666f2872297d29297d2c7265736574526573706f6e7365436f6e74656e743a66756e6374696f6e28297b6966286e756c6c213d746869732e726573706f6e7365436f6e74656e742626226a736f6e223d3d746869732e726573706f6e7365436f6e74656e742e6d6f6465297b76617220653d746869732e726573706f6e7365436f6e74656e742e746578743b746869732e726573706f6e7365436f6e74656e742e746578743d6f2e612e6a736f6e35737472696e67696679286f2e612e6a736f6e357061727365286529297d7d2c73686f774669656c644465734368616e67653a66756e6374696f6e2865297b76617220743d652e7461726765742e636865636b65643b746869732e24656d69742822646562756753686f774669656c644465736372697074696f6e4368616e6765222c74292c746869732e746f67676c654669656c644465736372697074696f6e2874297d2c6465627567456469746f724368616e67653a66756e6374696f6e2865297b746869732e24656d697428226465627567456469746f724368616e6765222c65297d2c746f67676c654669656c644465736372697074696f6e3a66756e6374696f6e2865297b76617220743d22726573706f6e7365456469746f72436f6e74656e74222b746869732e6170692e69642c6e3d646f63756d656e742e676574456c656d656e74427949642874292e676574456c656d656e74734279436c6173734e616d6528226b6e696665346a2d64656275672d656469746f722d6669656c642d6465736372697074696f6e22293b6f2e612e6172724e6f74456d707479286e293f6e2e666f7245616368282866756e6374696f6e2874297b742e7374796c652e646973706c61793d653f22626c6f636b223a226e6f6e65227d29293a746869732e73686f77456469746f724669656c64416e7957617928297d2c73686f77456469746f724669656c644465736372697074696f6e3a66756e6374696f6e2865297b76617220743d746869733b6f2e612e636865636b556e646566696e656428652926267061727365496e742865293c3d323030262673657454696d656f7574282866756e6374696f6e28297b742e73686f77456469746f724669656c645761697428297d292c313030297d2c73686f77456469746f724669656c64576169743a66756e6374696f6e28297b746869732e646562756753656e642626746869732e726573706f6e73654669656c644465736372697074696f6e436865636b65642626226a736f6e223d3d746869732e726573706f6e7365436f6e74656e742e6d6f64652626746869732e73686f77456469746f724669656c64416e7957617928297d2c73686f77456469746f724669656c64416e795761793a66756e6374696f6e28297b76617220653d746869732e73776167676572496e7374616e63652c743d746869732e6170692e6765744874747053756363657373436f64654f626a65637428292c6e3d22726573706f6e7365456469746f72436f6e74656e74222b746869732e6170692e69642c723d646f63756d656e742e676574456c656d656e7442794964286e292c693d5b5d2c733d722e676574456c656d656e74734279436c6173734e616d6528226163655f746578742d6c6179657222292c613d302c633d722e717565727953656c6563746f7228222e6163655f7072696e742d6d617267696e22293b6966286f2e612e636865636b556e646566696e656428632926266f2e612e636865636b556e646566696e656428632e7374796c6529262628613d632e7374796c652e6c656674292c732e6c656e6774683e3029666f722876617220753d735b305d2e676574456c656d656e74734279436c6173734e616d6528226163655f6c696e6522292c6c3d303b6c3c752e6c656e6774683b6c2b2b297b76617220643d755b6c5d2c703d642e676574456c656d656e74734279436c6173734e616d6528226163655f7661726961626c6522292c663d6e756c6c3b6966286f2e612e6172724e6f74456d707479287029297b663d6f2e612e746f537472696e6728705b305d2e696e6e657248544d4c2c2222292e7265706c616365282f5e22282e2a2922242f672c22243122293b76617220683d642e676574456c656d656e74734279436c6173734e616d6528226b6e696665346a2d64656275672d656469746f722d6669656c642d6465736372697074696f6e22293b696628216f2e612e6172724e6f74456d707479286829297b76617220793d646f63756d656e742e637265617465456c656d656e7428227370616e22293b792e636c6173734e616d653d226b6e696665346a2d64656275672d656469746f722d6669656c642d6465736372697074696f6e222c792e696e6e657248544d4c3d742e726573706f6e73654465736372697074696f6e46696e6428692c662c65292c792e7374796c652e6c6566743d612c642e617070656e644368696c642879297d7d76617220673d642e676574456c656d656e74734279436c6173734e616d6528226163655f706172656e22293b6966286f2e612e6172724e6f74456d707479286729297b666f722876617220623d5b5d2c6d3d303b6d3c672e6c656e6774683b6d2b2b29622e7075736828675b6d5d2e696e6e657248544d4c293b73776974636828622e6a6f696e28222229297b63617365225b223a63617365227b223a692e7075736828667c7c30293b627265616b3b63617365227d223a63617365225d223a692e706f7028297d7d7d7d7d7d2c613d6e28223238373722292c633d4f626a65637428612e612928732c2866756e6374696f6e28297b76617220653d746869732c743d652e24637265617465456c656d656e742c6e3d652e5f73656c662e5f637c7c743b72657475726e206e2822612d726f77222c7b737461746963436c6173733a226b6e696665346a2d64656275672d726573706f6e7365227d2c5b652e646562756753656e643f6e2822612d726f77222c5b6e2822612d74616273222c7b61747472733a7b64656661756c744163746976654b65793a226465627567526573706f6e7365227d7d2c5b6e282274656d706c617465222c7b736c6f743a227461624261724578747261436f6e74656e74227d2c5b652e726573706f6e73655374617475733f6e2822612d726f77222c7b737461746963436c6173733a226b6e696665346a2d64656275672d737461747573227d2c5b6e28227370616e222c5b6e2822612d636865636b626f78222c7b61747472733a7b64656661756c74436865636b65643a652e726573706f6e73654669656c644465736372697074696f6e436865636b65647d2c6f6e3a7b6368616e67653a652e73686f774669656c644465734368616e67657d7d2c5b6e28227370616e222c7b7374617469635374796c653a7b636f6c6f723a2223393139313931227d2c646f6d50726f70733a7b696e6e657248544d4c3a652e5f7328652e2474282264656275672e726573706f6e73652e73686f774465732229297d7d2c5b652e5f762822e698bee7a4bae8afb4e6988e22295d295d295d2c31292c6e28227370616e222c7b737461746963436c6173733a226b6579222c646f6d50726f70733a7b696e6e657248544d4c3a652e5f7328652e2474282264656275672e726573706f6e73652e636f64652229297d7d2c5b652e5f762822e5938de5ba94e7a0813a22295d292c6e28227370616e222c7b737461746963436c6173733a2276616c7565227d2c5b652e5f7628652e5f7328652e726573706f6e73655374617475732e636f646529295d292c6e28227370616e222c7b737461746963436c6173733a226b6579222c646f6d50726f70733a7b696e6e657248544d4c3a652e5f7328652e2474282264656275672e726573706f6e73652e636f73742229297d7d2c5b652e5f762822e88097e697b63a22295d292c6e28227370616e222c7b737461746963436c6173733a2276616c7565227d2c5b652e5f7628652e5f7328652e726573706f6e73655374617475732e636f737429295d292c6e28227370616e222c7b737461746963436c6173733a226b6579222c646f6d50726f70733a7b696e6e657248544d4c3a652e5f7328652e2474282264656275672e726573706f6e73652e73697a652229297d7d2c5b652e5f762822e5a4a7e5b08f3a22295d292c6e28227370616e222c7b737461746963436c6173733a2276616c7565227d2c5b652e5f7628652e5f7328652e726573706f6e736553697a6554657874292b222022295d295d293a652e5f6528295d2c31292c6e2822612d7461622d70616e65222c7b6b65793a226465627567526573706f6e7365222c61747472733a7b7461623a652e6931386e2e64656275672e726573706f6e73652e636f6e74656e747d7d2c5b652e726573706f6e7365436f6e74656e743f6e2822612d726f77222c5b652e726573706f6e7365436f6e74656e742e626c6f62466c61673f6e2822612d726f77222c5b652e726573706f6e7365436f6e74656e742e696d616765466c61673f6e2822646976222c5b6e2822696d67222c7b61747472733a7b7372633a652e726573706f6e7365436f6e74656e742e626c6f6255726c7d7d295d293a6e2822646976222c5b6e2822612d627574746f6e222c7b61747472733a7b747970653a226c696e6b222c687265663a652e726573706f6e7365436f6e74656e742e626c6f6255726c2c646f776e6c6f61643a652e726573706f6e7365436f6e74656e742e626c6f6246696c654e616d657d2c646f6d50726f70733a7b696e6e657248544d4c3a652e5f7328652e2474282264656275672e726573706f6e73652e646f776e6c6f61642229297d7d2c5b652e5f762822e4b88be8bdbde69687e4bbb622295d295d2c31295d293a6e2822612d726f77222c7b61747472733a7b69643a22726573706f6e7365456469746f72436f6e74656e74222b652e6170692e69647d7d2c5b6e2822656469746f722d64656275672d73686f77222c7b61747472733a7b6465627567526573706f6e73653a652e6465627567526573706f6e73652c76616c75653a652e726573706f6e7365436f6e74656e742e746578742c6d6f64653a652e726573706f6e7365436f6e74656e742e6d6f64657d2c6f6e3a7b73686f774465736372697074696f6e3a652e73686f77456469746f724669656c644465736372697074696f6e2c6465627567456469746f724368616e67653a652e6465627567456469746f724368616e67657d7d295d2c31295d2c31293a652e5f6528295d2c31292c6e2822612d7461622d70616e65222c7b6b65793a226465627567526177222c61747472733a7b7461623a22526177222c666f72636552656e6465723a22227d7d2c5b6e2822612d726f77222c7b737461746963436c6173733a226b6e696665346a2d64656275672d726573706f6e73652d6d74227d2c5b6e2822612d627574746f6e222c7b61747472733a7b69643a2262746e4465627567436f7079526177222b652e6170692e69642c747970653a227072696d617279227d7d2c5b6e2822612d69636f6e222c7b61747472733a7b747970653a22636f7079227d7d292c652e5f7628222022292c6e28227370616e222c7b646f6d50726f70733a7b696e6e657248544d4c3a652e5f7328652e2474282264656275672e726573706f6e73652e636f70792229297d7d2c5b652e5f762822e5a48de588b622295d295d2c31295d2c31292c6e2822612d726f77222c7b737461746963436c6173733a226b6e696665346a2d64656275672d726573706f6e73652d6d74227d2c5b6e2822612d7465787461726561222c7b61747472733a7b726f77733a31302c76616c75653a652e726573706f6e7365526177546578747d7d295d2c31295d2c31292c6e2822612d7461622d70616e65222c7b6b65793a22646562756748656164657273222c61747472733a7b7461623a2248656164657273227d7d2c5b6e2822612d726f77222c7b737461746963436c6173733a226b6e696665346a2d64656275672d726573706f6e73652d6d74227d2c5b6e2822612d7461626c65222c7b61747472733a7b626f7264657265643a22222c73697a653a22736d616c6c222c636f6c756d6e733a652e726573706f6e7365486561646572436f6c756d6e2c706167696e6174696f6e3a652e706167696e6174696f6e2c64617461536f757263653a652e726573706f6e7365486561646572732c726f774b65793a226964227d7d295d2c31295d2c31292c6e2822612d7461622d70616e65222c7b6b65793a2264656275674375726c222c61747472733a7b7461623a224375726c227d7d2c5b6e2822612d726f77222c7b737461746963436c6173733a226b6e696665346a2d64656275672d726573706f6e73652d6d74227d2c5b6e2822612d627574746f6e222c7b61747472733a7b69643a2262746e4465627567436f70794375726c222b652e6170692e69642c747970653a227072696d617279227d7d2c5b6e2822612d69636f6e222c7b61747472733a7b747970653a22636f7079227d7d292c652e5f7628222022292c6e28227370616e222c7b646f6d50726f70733a7b696e6e657248544d4c3a652e5f7328652e2474282264656275672e726573706f6e73652e636f70792229297d7d2c5b652e5f762822e5a48de588b622295d295d2c31295d2c31292c6e2822612d726f77222c7b737461746963436c6173733a226b6e696665346a2d64656275672d726573706f6e73652d6d74227d2c5b6e2822707265222c7b737461746963436c6173733a226b6e696665346a2d64656275672d726573706f6e73652d6375726c227d2c5b652e5f7628652e5f7328652e726573706f6e73654375726c5465787429295d295d295d2c31292c6e756c6c213d652e726573706f6e7365436f6e74656e7426266e756c6c213d652e726573706f6e7365436f6e74656e742e62617365363426262222213d652e726573706f6e7365436f6e74656e742e6261736536343f6e2822612d7461622d70616e65222c7b6b65793a226465627567426173653634496d67222c61747472733a7b7461623a22426173653634496d67227d7d2c5b6e2822612d726f77222c7b737461746963436c6173733a226b6e696665346a2d64656275672d726573706f6e73652d6d74227d2c5b6e2822696d67222c7b61747472733a7b7372633a652e726573706f6e7365436f6e74656e742e6261736536347d7d295d295d2c31293a652e5f6528295d2c32295d2c31293a6e2822612d726f7722295d2c31297d292c5b5d2c21312c6e756c6c2c6e756c6c2c6e756c6c293b742e64656661756c743d632e6578706f7274737d2c623331313a66756e6374696f6e28652c742c6e297b652e6578706f7274733d66756e6374696f6e2865297b76617220743d7b7d3b66756e6374696f6e206e286f297b696628745b6f5d2972657475726e20745b6f5d2e6578706f7274733b76617220723d745b6f5d3d7b693a6f2c6c3a21312c6578706f7274733a7b7d7d3b72657475726e20655b6f5d2e63616c6c28722e6578706f7274732c722c722e6578706f7274732c6e292c722e6c3d21302c722e6578706f7274737d72657475726e206e2e6d3d652c6e2e633d742c6e2e643d66756e6374696f6e28652c742c6f297b6e2e6f28652c74297c7c4f626a6563742e646566696e6550726f706572747928652c742c7b656e756d657261626c653a21302c6765743a6f7d297d2c6e2e723d66756e6374696f6e2865297b22756e646566696e656422213d747970656f662053796d626f6c262653796d626f6c2e746f537472696e6754616726264f626a6563742e646566696e6550726f706572747928652c53796d626f6c2e746f537472696e675461672c7b76616c75653a224d6f64756c65227d292c4f626a6563742e646566696e6550726f706572747928652c225f5f65734d6f64756c65222c7b76616c75653a21307d297d2c6e2e743d66756e6374696f6e28652c74297b696628312674262628653d6e286529292c3826742972657475726e20653b6966283426742626226f626a656374223d3d747970656f6620652626652626652e5f5f65734d6f64756c652972657475726e20653b766172206f3d4f626a6563742e637265617465286e756c6c293b6966286e2e72286f292c4f626a6563742e646566696e6550726f7065727479286f2c2264656661756c74222c7b656e756d657261626c653a21302c76616c75653a657d292c322674262622737472696e6722213d747970656f66206529666f7228766172207220696e2065296e2e64286f2c722c66756e6374696f6e2874297b72657475726e20655b745d7d2e62696e64286e756c6c2c7229293b72657475726e206f7d2c6e2e6e3d66756e6374696f6e2865297b76617220743d652626652e5f5f65734d6f64756c653f66756e6374696f6e28297b72657475726e20652e64656661756c747d3a66756e6374696f6e28297b72657475726e20657d3b72657475726e206e2e6428742c2261222c74292c747d2c6e2e6f3d66756e6374696f6e28652c74297b72657475726e204f626a6563742e70726f746f747970652e6861734f776e50726f70657274792e63616c6c28652c74297d2c6e2e703d22222c6e286e2e733d36297d285b66756e6374696f6e28652c74297b652e6578706f7274733d66756e6374696f6e2865297b76617220743b6966282253454c454354223d3d3d652e6e6f64654e616d6529652e666f63757328292c743d652e76616c75653b656c73652069662822494e505554223d3d3d652e6e6f64654e616d657c7c225445585441524541223d3d3d652e6e6f64654e616d65297b766172206e3d652e6861734174747269627574652822726561646f6e6c7922293b6e7c7c652e7365744174747269627574652822726561646f6e6c79222c2222292c652e73656c65637428292c652e73657453656c656374696f6e52616e676528302c652e76616c75652e6c656e677468292c6e7c7c652e72656d6f76654174747269627574652822726561646f6e6c7922292c743d652e76616c75657d656c73657b652e6861734174747269627574652822636f6e74656e746564697461626c6522292626652e666f63757328293b766172206f3d77696e646f772e67657453656c656374696f6e28292c723d646f63756d656e742e63726561746552616e676528293b722e73656c6563744e6f6465436f6e74656e74732865292c6f2e72656d6f7665416c6c52616e67657328292c6f2e61646452616e67652872292c743d6f2e746f537472696e6728297d72657475726e20747d7d2c66756e6374696f6e28652c74297b66756e6374696f6e206e28297b7d6e2e70726f746f747970653d7b6f6e3a66756e6374696f6e28652c742c6e297b766172206f3d746869732e657c7c28746869732e653d7b7d293b72657475726e286f5b655d7c7c286f5b655d3d5b5d29292e70757368287b666e3a742c6374783a6e7d292c746869737d2c6f6e63653a66756e6374696f6e28652c742c6e297b766172206f3d746869733b66756e6374696f6e207228297b6f2e6f666628652c72292c742e6170706c79286e2c617267756d656e7473297d72657475726e20722e5f3d742c746869732e6f6e28652c722c6e297d2c656d69743a66756e6374696f6e2865297b666f722876617220743d5b5d2e736c6963652e63616c6c28617267756d656e74732c31292c6e3d2828746869732e657c7c28746869732e653d7b7d29295b655d7c7c5b5d292e736c69636528292c6f3d302c723d6e2e6c656e6774683b6f3c723b6f2b2b296e5b6f5d2e666e2e6170706c79286e5b6f5d2e6374782c74293b72657475726e20746869737d2c6f66663a66756e6374696f6e28652c74297b766172206e3d746869732e657c7c28746869732e653d7b7d292c6f3d6e5b655d2c723d5b5d3b6966286f26267429666f722876617220693d302c733d6f2e6c656e6774683b693c733b692b2b296f5b695d2e666e213d3d7426266f5b695d2e666e2e5f213d3d742626722e70757368286f5b695d293b72657475726e20722e6c656e6774683f6e5b655d3d723a64656c657465206e5b655d2c746869737d7d2c652e6578706f7274733d6e2c652e6578706f7274732e54696e79456d69747465723d6e7d2c66756e6374696f6e28652c742c6e297b766172206f3d6e2833292c723d6e2834293b652e6578706f7274733d66756e6374696f6e28652c742c6e297b6966282165262621742626216e297468726f77206e6577204572726f7228224d697373696e6720726571756972656420617267756d656e747322293b696628216f2e737472696e67287429297468726f77206e657720547970654572726f7228225365636f6e6420617267756d656e74206d757374206265206120537472696e6722293b696628216f2e666e286e29297468726f77206e657720547970654572726f722822546869726420617267756d656e74206d75737420626520612046756e6374696f6e22293b6966286f2e6e6f64652865292972657475726e2066756e6374696f6e28652c742c6e297b72657475726e20652e6164644576656e744c697374656e657228742c6e292c7b64657374726f793a66756e6374696f6e28297b652e72656d6f76654576656e744c697374656e657228742c6e297d7d7d28652c742c6e293b6966286f2e6e6f64654c6973742865292972657475726e2066756e6374696f6e28652c742c6e297b72657475726e2041727261792e70726f746f747970652e666f72456163682e63616c6c28652c2866756e6374696f6e2865297b652e6164644576656e744c697374656e657228742c6e297d29292c7b64657374726f793a66756e6374696f6e28297b41727261792e70726f746f747970652e666f72456163682e63616c6c28652c2866756e6374696f6e2865297b652e72656d6f76654576656e744c697374656e657228742c6e297d29297d7d7d28652c742c6e293b6966286f2e737472696e672865292972657475726e2066756e6374696f6e28652c742c6e297b72657475726e207228646f63756d656e742e626f64792c652c742c6e297d28652c742c6e293b7468726f77206e657720547970654572726f722822466972737420617267756d656e74206d757374206265206120537472696e672c2048544d4c456c656d656e742c2048544d4c436f6c6c656374696f6e2c206f72204e6f64654c69737422297d7d2c66756e6374696f6e28652c74297b742e6e6f64653d66756e6374696f6e2865297b72657475726e20766f69642030213d3d6526266520696e7374616e63656f662048544d4c456c656d656e742626313d3d3d652e6e6f6465547970657d2c742e6e6f64654c6973743d66756e6374696f6e2865297b766172206e3d4f626a6563742e70726f746f747970652e746f537472696e672e63616c6c2865293b72657475726e20766f69642030213d3d65262628225b6f626a656374204e6f64654c6973745d223d3d3d6e7c7c225b6f626a6563742048544d4c436f6c6c656374696f6e5d223d3d3d6e292626226c656e67746822696e2065262628303d3d3d652e6c656e6774687c7c742e6e6f646528655b305d29297d2c742e737472696e673d66756e6374696f6e2865297b72657475726e22737472696e67223d3d747970656f6620657c7c6520696e7374616e63656f6620537472696e677d2c742e666e3d66756e6374696f6e2865297b72657475726e225b6f626a6563742046756e6374696f6e5d223d3d3d4f626a6563742e70726f746f747970652e746f537472696e672e63616c6c2865297d7d2c66756e6374696f6e28652c742c6e297b766172206f3d6e2835293b66756e6374696f6e207228652c742c6e2c6f2c72297b76617220733d692e6170706c7928746869732c617267756d656e7473293b72657475726e20652e6164644576656e744c697374656e6572286e2c732c72292c7b64657374726f793a66756e6374696f6e28297b652e72656d6f76654576656e744c697374656e6572286e2c732c72297d7d7d66756e6374696f6e206928652c742c6e2c72297b72657475726e2066756e6374696f6e286e297b6e2e64656c65676174655461726765743d6f286e2e7461726765742c74292c6e2e64656c65676174655461726765742626722e63616c6c28652c6e297d7d652e6578706f7274733d66756e6374696f6e28652c742c6e2c6f2c69297b72657475726e2266756e6374696f6e223d3d747970656f6620652e6164644576656e744c697374656e65723f722e6170706c79286e756c6c2c617267756d656e7473293a2266756e6374696f6e223d3d747970656f66206e3f722e62696e64286e756c6c2c646f63756d656e74292e6170706c79286e756c6c2c617267756d656e7473293a2822737472696e67223d3d747970656f662065262628653d646f63756d656e742e717565727953656c6563746f72416c6c286529292c41727261792e70726f746f747970652e6d61702e63616c6c28652c2866756e6374696f6e2865297b72657475726e207228652c742c6e2c6f2c69297d2929297d7d2c66756e6374696f6e28652c74297b69662822756e646566696e656422213d747970656f6620456c656d656e74262621456c656d656e742e70726f746f747970652e6d617463686573297b766172206e3d456c656d656e742e70726f746f747970653b6e2e6d6174636865733d6e2e6d61746368657353656c6563746f727c7c6e2e6d6f7a4d61746368657353656c6563746f727c7c6e2e6d734d61746368657353656c6563746f727c7c6e2e6f4d61746368657353656c6563746f727c7c6e2e7765626b69744d61746368657353656c6563746f727d652e6578706f7274733d66756e6374696f6e28652c74297b666f72283b65262639213d3d652e6e6f6465547970653b297b6966282266756e6374696f6e223d3d747970656f6620652e6d6174636865732626652e6d6174636865732874292972657475726e20653b653d652e706172656e744e6f64657d7d7d2c66756e6374696f6e28652c742c6e297b2275736520737472696374223b6e2e722874293b766172206f3d6e2830292c723d6e2e6e286f292c693d2266756e6374696f6e223d3d747970656f662053796d626f6c26262273796d626f6c223d3d747970656f662053796d626f6c2e6974657261746f723f66756e6374696f6e2865297b72657475726e20747970656f6620657d3a66756e6374696f6e2865297b72657475726e206526262266756e6374696f6e223d3d747970656f662053796d626f6c2626652e636f6e7374727563746f723d3d3d53796d626f6c262665213d3d53796d626f6c2e70726f746f747970653f2273796d626f6c223a747970656f6620657d2c733d66756e6374696f6e28297b66756e6374696f6e206528652c74297b666f7228766172206e3d303b6e3c742e6c656e6774683b6e2b2b297b766172206f3d745b6e5d3b6f2e656e756d657261626c653d6f2e656e756d657261626c657c7c21312c6f2e636f6e666967757261626c653d21302c2276616c756522696e206f2626286f2e7772697461626c653d2130292c4f626a6563742e646566696e6550726f706572747928652c6f2e6b65792c6f297d7d72657475726e2066756e6374696f6e28742c6e2c6f297b72657475726e206e26266528742e70726f746f747970652c6e292c6f26266528742c6f292c747d7d28292c613d66756e6374696f6e28297b66756e6374696f6e20652874297b2866756e6374696f6e28652c74297b69662821286520696e7374616e63656f66207429297468726f77206e657720547970654572726f72282243616e6e6f742063616c6c206120636c61737320617320612066756e6374696f6e22297d2928746869732c65292c746869732e7265736f6c76654f7074696f6e732874292c746869732e696e697453656c656374696f6e28297d72657475726e207328652c5b7b6b65793a227265736f6c76654f7074696f6e73222c76616c75653a66756e6374696f6e28297b76617220653d617267756d656e74732e6c656e6774683e302626766f69642030213d3d617267756d656e74735b305d3f617267756d656e74735b305d3a7b7d3b746869732e616374696f6e3d652e616374696f6e2c746869732e636f6e7461696e65723d652e636f6e7461696e65722c746869732e656d69747465723d652e656d69747465722c746869732e7461726765743d652e7461726765742c746869732e746578743d652e746578742c746869732e747269676765723d652e747269676765722c746869732e73656c6563746564546578743d22227d7d2c7b6b65793a22696e697453656c656374696f6e222c76616c75653a66756e6374696f6e28297b746869732e746578743f746869732e73656c65637446616b6528293a746869732e7461726765742626746869732e73656c65637454617267657428297d7d2c7b6b65793a2273656c65637446616b65222c76616c75653a66756e6374696f6e28297b76617220653d746869732c743d2272746c223d3d646f63756d656e742e646f63756d656e74456c656d656e742e676574417474726962757465282264697222293b746869732e72656d6f766546616b6528292c746869732e66616b6548616e646c657243616c6c6261636b3d66756e6374696f6e28297b72657475726e20652e72656d6f766546616b6528297d2c746869732e66616b6548616e646c65723d746869732e636f6e7461696e65722e6164644576656e744c697374656e65722822636c69636b222c746869732e66616b6548616e646c657243616c6c6261636b297c7c21302c746869732e66616b65456c656d3d646f63756d656e742e637265617465456c656d656e742822746578746172656122292c746869732e66616b65456c656d2e7374796c652e666f6e7453697a653d2231327074222c746869732e66616b65456c656d2e7374796c652e626f726465723d2230222c746869732e66616b65456c656d2e7374796c652e70616464696e673d2230222c746869732e66616b65456c656d2e7374796c652e6d617267696e3d2230222c746869732e66616b65456c656d2e7374796c652e706f736974696f6e3d226162736f6c757465222c746869732e66616b65456c656d2e7374796c655b743f227269676874223a226c656674225d3d222d393939397078223b766172206e3d77696e646f772e70616765594f66667365747c7c646f63756d656e742e646f63756d656e74456c656d656e742e7363726f6c6c546f703b746869732e66616b65456c656d2e7374796c652e746f703d6e2b227078222c746869732e66616b65456c656d2e7365744174747269627574652822726561646f6e6c79222c2222292c746869732e66616b65456c656d2e76616c75653d746869732e746578742c746869732e636f6e7461696e65722e617070656e644368696c6428746869732e66616b65456c656d292c746869732e73656c6563746564546578743d72282928746869732e66616b65456c656d292c746869732e636f70795465787428297d7d2c7b6b65793a2272656d6f766546616b65222c76616c75653a66756e6374696f6e28297b746869732e66616b6548616e646c6572262628746869732e636f6e7461696e65722e72656d6f76654576656e744c697374656e65722822636c69636b222c746869732e66616b6548616e646c657243616c6c6261636b292c746869732e66616b6548616e646c65723d6e756c6c2c746869732e66616b6548616e646c657243616c6c6261636b3d6e756c6c292c746869732e66616b65456c656d262628746869732e636f6e7461696e65722e72656d6f76654368696c6428746869732e66616b65456c656d292c746869732e66616b65456c656d3d6e756c6c297d7d2c7b6b65793a2273656c656374546172676574222c76616c75653a66756e6374696f6e28297b746869732e73656c6563746564546578743d72282928746869732e746172676574292c746869732e636f70795465787428297d7d2c7b6b65793a22636f707954657874222c76616c75653a66756e6374696f6e28297b76617220653d766f696420303b7472797b653d646f63756d656e742e65786563436f6d6d616e6428746869732e616374696f6e297d63617463682874297b653d21317d746869732e68616e646c65526573756c742865297d7d2c7b6b65793a2268616e646c65526573756c74222c76616c75653a66756e6374696f6e2865297b746869732e656d69747465722e656d697428653f2273756363657373223a226572726f72222c7b616374696f6e3a746869732e616374696f6e2c746578743a746869732e73656c6563746564546578742c747269676765723a746869732e747269676765722c636c65617253656c656374696f6e3a746869732e636c65617253656c656374696f6e2e62696e642874686973297d297d7d2c7b6b65793a22636c65617253656c656374696f6e222c76616c75653a66756e6374696f6e28297b746869732e747269676765722626746869732e747269676765722e666f63757328292c646f63756d656e742e616374697665456c656d656e742e626c757228292c77696e646f772e67657453656c656374696f6e28292e72656d6f7665416c6c52616e67657328297d7d2c7b6b65793a2264657374726f79222c76616c75653a66756e6374696f6e28297b746869732e72656d6f766546616b6528297d7d2c7b6b65793a22616374696f6e222c7365743a66756e6374696f6e28297b76617220653d617267756d656e74732e6c656e6774683e302626766f69642030213d3d617267756d656e74735b305d3f617267756d656e74735b305d3a22636f7079223b696628746869732e5f616374696f6e3d652c22636f707922213d3d746869732e5f616374696f6e26262263757422213d3d746869732e5f616374696f6e297468726f77206e6577204572726f722827496e76616c69642022616374696f6e222076616c75652c20757365206569746865722022636f707922206f7220226375742227297d2c6765743a66756e6374696f6e28297b72657475726e20746869732e5f616374696f6e7d7d2c7b6b65793a22746172676574222c7365743a66756e6374696f6e2865297b696628766f69642030213d3d65297b69662821657c7c226f626a65637422213d3d28766f696420303d3d3d653f22756e646566696e6564223a69286529297c7c31213d3d652e6e6f646554797065297468726f77206e6577204572726f722827496e76616c69642022746172676574222076616c75652c2075736520612076616c696420456c656d656e7427293b69662822636f7079223d3d3d746869732e616374696f6e2626652e686173417474726962757465282264697361626c65642229297468726f77206e6577204572726f722827496e76616c6964202274617267657422206174747269627574652e20506c65617365207573652022726561646f6e6c792220696e7374656164206f66202264697361626c6564222061747472696275746527293b69662822637574223d3d3d746869732e616374696f6e262628652e6861734174747269627574652822726561646f6e6c7922297c7c652e686173417474726962757465282264697361626c6564222929297468726f77206e6577204572726f722827496e76616c6964202274617267657422206174747269627574652e20596f752063616e5c27742063757420746578742066726f6d20656c656d656e747320776974682022726561646f6e6c7922206f72202264697361626c656422206174747269627574657327293b746869732e5f7461726765743d657d7d2c6765743a66756e6374696f6e28297b72657475726e20746869732e5f7461726765747d7d5d292c657d28292c633d6e2831292c753d6e2e6e2863292c6c3d6e2832292c643d6e2e6e286c292c703d2266756e6374696f6e223d3d747970656f662053796d626f6c26262273796d626f6c223d3d747970656f662053796d626f6c2e6974657261746f723f66756e6374696f6e2865297b72657475726e20747970656f6620657d3a66756e6374696f6e2865297b72657475726e206526262266756e6374696f6e223d3d747970656f662053796d626f6c2626652e636f6e7374727563746f723d3d3d53796d626f6c262665213d3d53796d626f6c2e70726f746f747970653f2273796d626f6c223a747970656f6620657d2c663d66756e6374696f6e28297b66756e6374696f6e206528652c74297b666f7228766172206e3d303b6e3c742e6c656e6774683b6e2b2b297b766172206f3d745b6e5d3b6f2e656e756d657261626c653d6f2e656e756d657261626c657c7c21312c6f2e636f6e666967757261626c653d21302c2276616c756522696e206f2626286f2e7772697461626c653d2130292c4f626a6563742e646566696e6550726f706572747928652c6f2e6b65792c6f297d7d72657475726e2066756e6374696f6e28742c6e2c6f297b72657475726e206e26266528742e70726f746f747970652c6e292c6f26266528742c6f292c747d7d28292c683d66756e6374696f6e2865297b66756e6374696f6e207428652c6e297b2166756e6374696f6e28652c74297b69662821286520696e7374616e63656f66207429297468726f77206e657720547970654572726f72282243616e6e6f742063616c6c206120636c61737320617320612066756e6374696f6e22297d28746869732c74293b766172206f3d66756e6374696f6e28652c74297b6966282165297468726f77206e6577205265666572656e63654572726f72282274686973206861736e2774206265656e20696e697469616c69736564202d2073757065722829206861736e2774206265656e2063616c6c656422293b72657475726e21747c7c226f626a65637422213d747970656f66207426262266756e6374696f6e22213d747970656f6620743f653a747d28746869732c28742e5f5f70726f746f5f5f7c7c4f626a6563742e67657450726f746f747970654f66287429292e63616c6c287468697329293b72657475726e206f2e7265736f6c76654f7074696f6e73286e292c6f2e6c697374656e436c69636b2865292c6f7d72657475726e2066756e6374696f6e28652c74297b6966282266756e6374696f6e22213d747970656f66207426266e756c6c213d3d74297468726f77206e657720547970654572726f72282253757065722065787072657373696f6e206d75737420656974686572206265206e756c6c206f7220612066756e6374696f6e2c206e6f7420222b747970656f662074293b652e70726f746f747970653d4f626a6563742e63726561746528742626742e70726f746f747970652c7b636f6e7374727563746f723a7b76616c75653a652c656e756d657261626c653a21312c7772697461626c653a21302c636f6e666967757261626c653a21307d7d292c742626284f626a6563742e73657450726f746f747970654f663f4f626a6563742e73657450726f746f747970654f6628652c74293a652e5f5f70726f746f5f5f3d74297d28742c65292c6628742c5b7b6b65793a227265736f6c76654f7074696f6e73222c76616c75653a66756e6374696f6e28297b76617220653d617267756d656e74732e6c656e6774683e302626766f69642030213d3d617267756d656e74735b305d3f617267756d656e74735b305d3a7b7d3b746869732e616374696f6e3d2266756e6374696f6e223d3d747970656f6620652e616374696f6e3f652e616374696f6e3a746869732e64656661756c74416374696f6e2c746869732e7461726765743d2266756e6374696f6e223d3d747970656f6620652e7461726765743f652e7461726765743a746869732e64656661756c745461726765742c746869732e746578743d2266756e6374696f6e223d3d747970656f6620652e746578743f652e746578743a746869732e64656661756c74546578742c746869732e636f6e7461696e65723d226f626a656374223d3d3d7028652e636f6e7461696e6572293f652e636f6e7461696e65723a646f63756d656e742e626f64797d7d2c7b6b65793a226c697374656e436c69636b222c76616c75653a66756e6374696f6e2865297b76617220743d746869733b746869732e6c697374656e65723d64282928652c22636c69636b222c2866756e6374696f6e2865297b72657475726e20742e6f6e436c69636b2865297d29297d7d2c7b6b65793a226f6e436c69636b222c76616c75653a66756e6374696f6e2865297b76617220743d652e64656c65676174655461726765747c7c652e63757272656e745461726765743b746869732e636c6970626f617264416374696f6e262628746869732e636c6970626f617264416374696f6e3d6e756c6c292c746869732e636c6970626f617264416374696f6e3d6e65772061287b616374696f6e3a746869732e616374696f6e2874292c7461726765743a746869732e7461726765742874292c746578743a746869732e746578742874292c636f6e7461696e65723a746869732e636f6e7461696e65722c747269676765723a742c656d69747465723a746869737d297d7d2c7b6b65793a2264656661756c74416374696f6e222c76616c75653a66756e6374696f6e2865297b72657475726e20792822616374696f6e222c65297d7d2c7b6b65793a2264656661756c74546172676574222c76616c75653a66756e6374696f6e2865297b76617220743d792822746172676574222c65293b696628742972657475726e20646f63756d656e742e717565727953656c6563746f722874297d7d2c7b6b65793a2264656661756c7454657874222c76616c75653a66756e6374696f6e2865297b72657475726e2079282274657874222c65297d7d2c7b6b65793a2264657374726f79222c76616c75653a66756e6374696f6e28297b746869732e6c697374656e65722e64657374726f7928292c746869732e636c6970626f617264416374696f6e262628746869732e636c6970626f617264416374696f6e2e64657374726f7928292c746869732e636c6970626f617264416374696f6e3d6e756c6c297d7d5d2c5b7b6b65793a226973537570706f72746564222c76616c75653a66756e6374696f6e28297b76617220653d617267756d656e74732e6c656e6774683e302626766f69642030213d3d617267756d656e74735b305d3f617267756d656e74735b305d3a5b22636f7079222c22637574225d2c743d22737472696e67223d3d747970656f6620653f5b655d3a652c6e3d2121646f63756d656e742e7175657279436f6d6d616e64537570706f727465643b72657475726e20742e666f7245616368282866756e6374696f6e2865297b6e3d6e26262121646f63756d656e742e7175657279436f6d6d616e64537570706f727465642865297d29292c6e7d7d5d292c747d28752e61293b66756e6374696f6e207928652c74297b766172206e3d22646174612d636c6970626f6172642d222b653b696628742e686173417474726962757465286e292972657475726e20742e676574417474726962757465286e297d742e64656661756c743d687d5d292e64656661756c747d7d5d293b`
)

func AddRouterOfChunk589faee0F3951d02Js(router *gin.Engine) {
    
    utils.GetJs(router, CHUNK_589FAEE0_F3951D02_JS_RELATIVE_PATH, CHUNK_589FAEE0_F3951D02_JS_HEX_CONTENT)
    
}







