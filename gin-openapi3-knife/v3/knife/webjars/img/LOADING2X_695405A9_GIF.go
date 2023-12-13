package img


import (
	"gitee.com/youbeiwuhuan/knife4go/gin-openapi3-knife/constant"
	"gitee.com/youbeiwuhuan/knife4go/gin-openapi3-knife/utils"
	"github.com/gin-gonic/gin"

)


const (
	LOADING2X_695405A9_GIF_RELATIVE_PATH = constant.ROOT_PATH + "/webjars/img/loading2x.695405a9.gif"
	// 文件内容的16进制表示
	LOADING2X_695405A9_GIF_HEX_CONTENT = `47494638396140004000c600000496e48ce2fc34cefcc4fefca4e6fc64f2fc5cd2fc84f6fc3cc6fc1cb2fce4fefcacf6fc5cdefc4ccefc34c2fc94e6fcd4f6fc94f6fcb4f6fc2cbafc74defc7cfefc4cdafc64fafc64dafcd4fefcb4fefc84eafc44d2fcb4f2fc84fefcf4fefc94eefc94fefc34bafc74e6fcccfafc6ccefc5ccafc24bafca4fefc64defc4cd6fc34cafcdcfafcbcf6fc4ce2fc6cfefcbceefc1caefc8ce6fc3ccefc6cf2fc64cefc8cf6fc24b6fcecfefc5ce2fc54cefc34c6fc9ceafcd4fafc9cf6fc2cbefc7cdefc54dafc64fefc6cdafcdcfefcbcfefc8ceafc44d6fc8cfefcfcfefc9ceefc9cfefc34befc7ce6fcccfefc6ccafcacfefcbcfafcbcf2fcffffff00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000021ff0b4e45545343415045322e30030100000021f90409090053002c00000000400040000007fe80538283848586878245053403888e8f90918e032e952e1992999a9b11962e3e9ba1a287369e11a3a885441e2f4b1f8ea596a788382333393da9881f2f42174221b0a68e292b3b2b1c49ba864417cece988707c3872dc6d7d1cb8338bedd15ca86b195b385492ad7c738da8528cf2f1745a4d4854ae82b1beb8538bdeeaf85e22ec86d9b71ec98007ff9066970d60d45b879839a1434b62061b90a17de655450689aac423daefd1070049cc5414e9e3943f20f621206c756c86c74b210926ece9c1002285042c1631496255180f090028defbe0de239e803876b32893c52e0a4a82a0f153c68b04a0885af5f17340c82622342041b62058198f8c386231cfe4312c4f8010151880a78b33a314988d757211ef81afa6001ea11ae533e3cb811a3718c0688b04ac61b42aaa122192bec8d9444c2911d3b2a1aea80c0b1dc042710a1c8cbba020a75849254802218d2870d0c04b3d061da710c0cbb964c1eae812f62494571004970ba7703d8883220698dd78387cdeb92f09820d7b7dc1d12242519603d6b5eac4b6a8bc281a07b8ce6098c1cdf05a53aebb4da8634769f2005f4500a2c615f08ea8d82c30fdd25c00109da4887974e09f1d0d80fa265570481167db0830cf3a552a02e1fd624e288244e35c08945a0a8e200456473d20335c428e38c313e21c514192c11828e3cf6b8e31234d5640200441669a49104a0feb0448f3e2eb96390273d71e494454ea0e48e5836b92384359540259513e4f8e38f5a4269d1905f1e49c0140aa898e29b279ee8a24501d068670d4fc050e29e7cf6394888a8003a4a12191421e826491441c4a199285084061a70941011901621a92e380ca001149532ca59119c423a40878f7ce004a4a8723a672aa7421aaa06d8659204a5ae6e8aaaa1da245aabad902e2a89a3a806dbaba790cc2aecab45fc67c878a9daaa2a5fc4c616dba9afbe5a15229af28aeaa8e5b4486c124e0c20d8079a560b2997851cab81a587d06aa9a0433d1aa9518f861a2a22f5a2eaebb2c2e2ea88aebcaa676cb0666e836aac8664b0ab655310e144061938c1b0c2c1b2aeda14b5c9fe8b43871f685b1c4ab56693c4ae5b958a43b486646beba553500ba98bb4da8aae880aec6a66ab9bbac8ecabca5a04b0ad56b9acc19c3834eb6f4d945a6b08ce50aceac4ab1a309c61b31f1722f4aa1d3b5b04a9a328fcaad420db6a31c5a85aac4bc7c8aaf7f4c187e8ca29a7283f5274ad2c13d2aaaa88d41c6ccf20f23ab3ddc1fe3d48b990c62df7a3085bfdf6d08e98ba2edf7de20cab9fa8dccd38e5a2e8bd29e4984342448a75af13080021f90409090059002c0000000040004000860496e484dafc3ccefcc4f6fc5ceafcaceafc64cefc34befc84f2fce4fefc4ce2fca4fafc84fefc14a6ecd4f6fc74f2fc64defc4cd6fc4cc2fcbcf6fc24aef464fafc74dafc9ceefc94fafc44d6fcc4fafc6ceefcb4eefc0c9eec84eafc44cefca4f2fc6cdafc34cafcf4fefc5ce2fc1caefcd4fefcb4fefc9cfafc94e2fc64eafc6cd2fc44befc8cf6fc54e6fcacfefc1caaf47cfefc7ce6fc54dafc5ccafc2cb2f494f6fcccfefc049ae48cdefc3cd2fcccf2fcaceefc64d2fc3cbefc84f6fcecfefc4ce6fca4fefc8cfefc14a6f4d4fafc7cf2fc6ce2fc4cdafc54cafcbcfafc24b2fc7cdafc9cf2fc94fefcc4fefc6cf2fc0ca2ec44d2fcacf2fcfcfefcdcfefcbcfefc9cfefc64eefcffffff00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000007fe8059828384858687824f313137888e8f90918e261595155592999a9b5796152f9ba1a2879d9642a3a885094e0c2f238ea595a788403f2e5098a987230c8b31b3a49ec086460ac60454ba8655bdbd31b986b115c3844fc6d7d0ca8240bebd4ec9d1c2875458d70a41afda8427be8b4fc1a6872fe60a36ea85bc31cd0ce984d2d45980b83047a0dfbd44ed629c0817af500b7a0b0f12a2e2449fc50485fe15aa426f033889834cb4637025a3b841541ed06b04b290108b8b4cf83b29c80a3d04caa82430682881b35e433e6a14448580b920181d01b9f1b1a7902b4e9ef05cd72dc6bb9a5784bcb86265100a7a0b94ca1021424a11442f9c5c597bc544d341febc9cc5f8f6881cc1a959a8341140962c09446bd5aacd9ab490a2b96e2351b1a2c25857434a22f4edab03d109a881035b015288a295b790466078f0b6ca91c90744a4968168c4e5c1b0af3c698a5792c1111efaa69e0ca1b6a02a4f05b35d5b05342a2a53744c269b3a830649546eb0150ef585f1502366a09e8ce27a6b2b98855fd53656f564199c4701d98a5988f75040a44c9e21539989a76d254e212b65c2412a4f5827d108485ce05b2aef1dd7d2820c36a89e093740286184149a908d441758a0e1861c6ac8c400595471c288249658224b2d8550c28a2cb6d82208569cf0828c34ce68e388f5b5c4848b3cb2f8418c260639638e2059c022113db2fe28808837367922832a26e92208592460c28413dea0a5850c6618c2976086198205fe3968e6996812d55282a35061c267032e11c081a124006461eaa40000002514a00d104fd078029cea8c00c39e7bfae0002a23dc60e290f7f48028a23818905e26548838e28d2312aa0b10354c3a691439b0592590419e50dc3d54e4d080a8882ec1415d816e5ae30989a1a4e02040f48003ac7b4a70a961b69628556756e48ae90db31552840fc00290846525ce68059e83687a2d9b3add794801878ada002231ce38e2aa865061a2a7e30069ae7123e480c3af7b1e400b894c4d522c345544082134269848a42a34fcdac00e8e5001c481ae993be2478ed298a3ba82a49e70600213d01949ad34e2e9a8b9446a4a238a0c2650ec7882442c239100563c2c488b9138234f1fe3580810320f6aea284ce26b88ca2f0c9cc50d0eabbaa06bc5825633ae86206da3151a8b12b0c317a66cabd059045c22d6ba344c22bb8310ed73bae5d2b8732638db8a6dd857236272892fe72428c985a87c02ddd5087a36da31e68b88d890b6e6e86669fefc64e1a2441c38e29bbc2d63dc8c475285154facad4d200021f90409090058002c0000000040004000860496e484dafc4ccefcc4f6fca4e6fc34bafc5ceafc84f2fc1caef4e4fefc6cd2fca4f6fc7ceafcd4fafc4cc6fc24bafc94f6fcb4f6fc84e6fc4ce2fc34c6fcc4fefc64fafc2cb2f414a6ec24b6fc6cdafc74fefc94e6fc64d6fcb4f2fc34c2fc6ceefc84fefc24b2fcf4fefca4fefcdcfafc94fefcb4fefc0c9eecccf2fc3cbefc64f2fc1cb2fc74f2fc5ccafc2cbafc5ce2fc3cc6fcccfafc7cdefc9ce6fc049ae494e2fc54cefcaceafc34befc64eefc8cf6fc1caefcecfefc7cd6fcacfafcd4fefc54c6fc94fafcbcf6fc8ce6fc54e6fc34cafc6cfefc1caaf474defc7cfefc64dafcbceefc8cfefcfcfefcacfefcdcfefc9cfefcbcfefcccf6fc7cf2fc2cbefcccfefc9ceafcffffff00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000007fe80588283848586878256262640888e8f90918e504a954a5092999a9b4f4a219e279ba2a3879d96a1a4a984094f5152238e9d9f21a8872326472109aa884e51bf2652b1964ab5864d16c91b4ebc86095126c0bba59ea08856c9d998cd84b7d0bf24cc86a695c683234ad91647b0dc8452c0bf8de3959fe68227ea1624ee8523f1d1c41122f1a99ca11e47d46d68d76f909578512ad0ab762f4ab684c21a12724202628f42e48a1582a24f89408d82a044fb66ce94bd8d21d6694369e8c4b768db0491a345a882be28cd9cf46078a807c470836419143462c3c523d310f50072b2508f135881102d5401a09541159e607d2251a7cc644f1cf5d8316182819cfe85a4603df1440a94aa4c213ec1bb315db685879c3c29d2b66d0b4473c56295f2d1d0c32824ee467262c5e991b25c75142e6ce09a62ba732b6c155c81af238e21aa26a0b279f30e44232a8c9d9df5a4e94844470869bdf9c0d64209e4829e6df7b62827270cf06e0be2eb642089a397e63602c4f20945f66ea21c1dab736e6c7943f89da947d8b9528c6f1aa1bc708ba8aa828b85cbed445b1d98b93901925ea393153f9017144dead164e081084a0584150b36c8e08340d0d7d002125468e185151221031650cce5e187de1dc880112496682289141811815c62b5089a8b27cc43131114d468e38d29a6388170207a28968c284970e28939520043872e26fee9e177283180e38d46a418e50f58240084830e5aa165840752c880040c84296698334830408268a6a9e620058ed2e671fcbd99c9083110212029c161051f3757f0c0c3071e7063de70fdb933420e2cf8c983000da43282151ffed84f123c20c082a53cb09044639938d1e16c9f15ca0b9d89629028027e3e70459b79f6788264ee38c1c1032cb060eaa57ec630c4234ec8062a68546de428213d24516ba2999aca43079c72455b68a34911ac2623745040550d0890ecb17e2e81988f2748b1674a8b2550a0130420010000341ce24115b5f2a0ec0388b03817ac8520e7a1a8b688b0eeba1890370211c7da2a80547359711b74c3e50405830be6e4c3bfebc9fae048021a245ac514a70df5c808c3612510a4c00ed283baffd650c2235054d02c29be82061fa48a014903c50004816602b4e58705c97401e94401383381a0133c8ab515cd59153200ce22c8290a92093b96189082dc80b30d06824c1b5e4cc7e84c0d1463f0b23bd07d26e1cfb3612dc80c382b8012c8a1f265c5674c9e8300c5359ccdcb55b38d9b48db88e080f3860d21a758de84007d02e383a8f06fc028f52097c28edc5d3222093850430e29ac6975d5a28f42b2a4a58bc2f3587ea70e0914525420783381000021f90409090050002c0000000040004000860496e484e2fcc4f6fc3ccefca4e6fc6cf2fc64d2fca4fefc3cc6fc8cf6fce4fefc54cefc1cb2fcd4fafcb4f2fc4cdafca4eefc64defc94f6fc34c2fc94e2fcc4fefc4cd6fc74fefcbcf6fc44d6fc5ccafc2cbafc7cdefc84eafcccf6fc44cefc64fafcf4fefc54dafcaceefc94fefcb4fefc8ce2fcaceafc4ccafc84fefcdcfefc34bafc74e6fc6cfefc1caefcc4fafc3cd2fca4eafc6cd6fcacfefcecfefc5cd6fc24b6fcd4fefcbcf2fc4ce2fca4f2fc6ce2fc9cf2fc34cafc94e6fcccfefc7cfefcbcfafc6ccafc2cbefcccfafc44d2fc64fefcfcfefc5ce2fcacf2fc9cfefcbcfefc8ce6fc54cafc8cfefc7ce6fcffffff00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000007fe8050828384858687823733332a888e8f90918e0a4a24950a92999a9b4b954a4a159ba2a3879d9ea1a4a984344b253f478e4b9695a8872107294a98aa874725bf25b5869db3c2854a40c94eb0bc8534bf33bfbbc39f96c68337c929c9d3cd8221c0bf4bcc85b2a7b64e40dbea21de853f25d1d18dd4e7c3ebc925ee85e0e125e483ccd1321462dd3627edf61152110fd80f6ac50ccd4846f1a14242475a352c9130603550855468534702e045280a80453366ce1a46641481d03b49a8c2c612dd880d1cf463648a03cd8ed0e8682bdc8c711eed4139928e620a1a8f68a830b90ad80da28514859bf96349850a4b6e0c2aa1aeec1247340e18017161663995fe4ba6f682160fa9a32324622e9bdb020488b52910a904b604aa2186259628a06ae8c80d27db2cbe03b276eddf0b887e44a3fbab02d68c3718df2da1c4a4022796fd1ae9ab0451089b1b3797088d3115d15b7dff82c80d8204d6420a34ca8e1657f4a88c1752f306223692e3c1742b18d75430b572bb998ec0e3ec6a9f12d57e771ff89d894605d9d89b85b8b0dbaf936ea982cfbb58c12f10c9de1ca7cf0fe41fcde9a90048d380041688080d37fc90e0820a3678835b179520c184145638210f6221e6cf86dd0d98400e208628a2884bb4224f43276ed61c4d128ce862880568c4e1602b9ec4c38b2f16c0508a288683df451fe238e2590adcc020833f24fef920813358e8e484351a28e594530a78dc7f37ec17d4033c58994970d25ca4430f133c808137e66da4a56d19f440660f114048dd7634eed3c1046ee6d943078665a7e189e278f9480822bc89e7a130e860259833cae5ce113ac0a0679e1388f0c22347c0d66334b40d22283fab74f0e61093b2d06721b0c9d6d967617d1a02071f98a44204941aca82604625069f20882926e011234ce0820b101c12449b93c2808889c0385a882fe1ac394808083030ac0b439077040f6e92da0312073a64dc0d9ccd6442134208d1840f830430acb50c0430c9136ec2d0c05d430d7a937f826800c0bf0068e0e910d7bac0809c8428f003799b681a4d37fe022cf020239fc06b700d52a6b4913111ff3bb120477c50b00b671298d16058350170c08510f1eeb0137cbac98e0e19d231cb854470adb5310cd8cf4654ddfcf120343060adc1369cba0fb9b221acb2c4873061f0bb1c9c040e7a8c09ddcb10471bac7433cfd0b56bbf2b0f4d8803535bfb63501bad3dc8d31e3bb2c0b53630ac0a2bae18a7f5812230300011541eb277e0a3dc2c04e1a49cb0320e889342c10643104053200021f90409090053002c0000000040004000860496e484defc34cefcc4eefc5ceafca4eefc64cefc84f6fc44befcacf6fc5cdefce4fefc4ce2fc1caefcb4f6fc4cd6fcc4fefc74f2fc94f6fc74d6fc84eafc74e6fcd4fafc44cefc64fafc2cb6f4b4fefc0c9eec6ceefcb4eefc64dafc84fefc4cc2fca4fefc6ce6fcf4fefc94fefc94eefc64f2fcaceafc64defc2cb2f4bcf6fc54dafcccfafc7cfefc7cdafcdcfefc44d6fc049ae48cdefc3ccefcc4f6fc64eefc6ccefc8cf6fc34cafcacfafc5ce2fcecfefc54e6fc24aef44cdafc7cf2fc94fafc74dafc8ceafc7ce6fcd4fefc44d2fc2cbafcbcfefc0ca2ec6cdafc8cfefc54c6fcacfefcfcfefc9cfefc9ceefcaceefcbcfafcccfefcffffff00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000007fe8053828384858687820b47470b888e8f90918e0b1a951a3b92999a9b521a4c95449ba2a3879d954ca1a4aa843b471a524d8e9d9f1aa987231a4e4c98ab874d961a10b296a88eb94e4e21b1bd853ba7958da59ea0882f244e24d7bccc83b8c047cb85a69eb6844d21d8c82423dce2d39f2fd2b4e58352d7e9c2ed84dec0e1f5c4f4a68c5087cc893f7d535e4cab24c5d02c6a858edcbb16505f13570bd91132f589de8274d794213444e959be7ab45e996342d049b49185202cd4f0720ac75a8488a4c3768459931d1a11f1f304ee9fa58682ceb5dc766bc14156968804354424a58678825e489142440a569bd7c222bdc5a4450b2535096124fae2e9945ffeef8a3a527a4fa4a126103eb4d0dbc2092262968e301da4500323b7859abc4047a222112566239f452425654a0853df1e218218115e26077784d84b9a2f13a13217a6e46c4e953f5c92257f0891b990a267cf8eb4b52845afefd224be3e6a4204b0c9ce9a469030cb37f20708c8ef8e4b39965959bebe35d4ceb403424ab9dc46408eec929b227823a5ec5552d13511f0ed9a28810ef36dfde8f5f3ebdf8f6847d7ff5c05f89f70231d11c281082688603c8501e3e051fa3981c1841456582104ae7ca2e13b0bb5c74d08168648e1076b3d089887cc802862881f28b4e18bc0543792842b5ad8d002440008e0565c118890810a06c9848ffc1569247ff88dfe92a428c4c1c74d1326e4b0a424b7d13492060c3050c349ab74b79093ab8c404096594640a424238cf30c8a9b0041e69b126c3757831b5602e6282370f0e69b3c6890649526ee161f1363ee99650d3226961a871dbda6ca54234860289907c8999a55c164d6440613c82914050a1cb4c00f933270c35fc07cc2c821320000400f27e0d784030fe080830387405083a1042092a12582368384abae1a81dc083ed86a6b11db3591030f6f46d0df519dd91003b1000430480949b8e08207056c6bab11b696e0c80e376449c0996f01f58805d712db43504934606f0341745384b2b6a6d58c54cc20802d00e10e52efbdf90e9200bf385450e4090323508807f7a7e26b8e0238906b2b973035d1c0c02c1472b0bd090fc282b2e4ae30a528010c6c802114236c48050c2790df021b608bc460538cdc401286ec3003bf3378ba8a0103cb7048ccf6026dc813e3da4a014c2f0cdc835b3e3b5dc808b5f2cb732f2a0cdc01224cff8c48140cb329ca0ec3ba0a8223656b6d080a19e330c3ca920c60440c4b7c2d48d6e78a20800e161cb974c5721bbe89cf252bbe490715abe038294fcc3043c12305020021f90409090062002c0000000040004000860496e484defc54cefcc4eefca4e6fc5ceafc34b6fc84f2fc6ccefce4fefca4f6fc1caaf474f2fcc4fafc94f6fc4cc2fcb4f6fc24b6fc94e6fc4ce2fc64fafcd4fafc3cc2fc74defcb4eefc0ca2ec84e6fc6ceefc84fefc6cdafca4fefc2cb2f494fefc5ce2fcf4fefc1cb2fc74fefcccfafcb4fefc2cb6fc4ce6fcdcfafc44cafcbceefc0c9eec8cdefc64d6fcc4f6fc64eafc44befc8cf6fcacf6fc24aef49cfafc9ce6fc6cfefc34cafc7cd6fc049ae484e2fc5cd6fcaceafc34bafc84f6fc74d2fcecfefca4fafc1caefc7cf2fcc4fefc94fafc5ccafcbcf6fc24bafc54defc64fefcd4fefc3cc6fc7ce6fcb4f2fc14a6ec8ce6fc6cf2fc8cfefc74d6fcacfefc9cfefcfcfefc7cfefcccfefcbcfefc2cbafc54e6fcdcfefcbcf2fcccf6fc64eefc9ceafcffffff00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000007fe806282838485868782095a5a09888e8f90918e092695264192999a9b592655954c9ba2a3879d9555a1a4aa84415a2659578e9d9f26a987229d4598ab8757962645b296a88e459e95b1bc8541a7958da5c7b58894bfbbca8322bf265ac985a69eb684575a9fb422d7dec79f5dd0b4e1834cdaefe8d9dadd83dfc485f5e526f7e882ba447b656816a882dad8012c34ae998973844c7d7a17a4dfb67f0bc5503b162c222d8283ae181b662d2321631f9f0992286d90c06859945d09021151bd53dcf0fd8a2968dcc72a356f05c1388819a8a085987c34a1504c972c599864691aafd93c6c5aac80f05092902b9c5d88fa5297d351c38f62b35859abb50aa261fe96b4740d58891151868a9c1deae2612d08ad6b1165f969a908d2714cee22ba92a5c83f666cfffeb5a2c5e6485a1f138b53754f4411c093273f7c9477e0272d61015e61e22134602b1e543e5a0db799e36b57aab87ead5913e3da20af65ddad05a9a620453e96c5dd777295b9a314adcbc844ab87a6b8992c4777c5032c936214af120fbebcf9f38e82485d1fb5fd7aec192f3bc4ccd4a9b6fb967882af82a5bfffff5870d09f76ea1468d155e89810e0820236e81f07207c851f5c085ea32080000a88050802f5e3e14ee655e1e088ff711053024cb0c71e5451c1b7907c1f5ae2227a34d6881e79a4e0388a08399ca0a3262260e18f6a042c00000036c4feb7040558e8b7ca173e1c79240dc6714602054b5230856c9b247084946052019007592eb9a407553a22420b2c8009e609696e12249614d449c10d776ee7480f34b829251436fc18c938242c89279d4b6051a11841c4e0e7913a7410549c904ceac1a1761e0a429a0f3c0ac0031530d4c40e941a2282110cfc93c01466dab9a41588e8e0270d4f1c12c61043e080018ee3c030c104951584459d659280c8096066d002514124812bae4d88770518bffe5a4095e3e079e81488ac9001003a20c02521010c3102ae23443188024e44a18113330c5243b5bfd6a0a6157592302e3648ec0bcf080b8c00c50816d4e4040e08e3a0013605d08b82bf8204911a2f02c4983bf008f10e7270c20b0fa205bd131c40e313e88e308200851cbc05c21df7c400c84e9a748505e75e5c42ca092b9c14c81b082a8a04258fd081211be3b045cb83c800b209e525e0acc023243157d13a2fc305bd0540a74c07254311c6211bacccf2214280ec8049099c3bc4c038104535d2d8505b2d0aa58ad280b9268f50eb211baf0cf72045803ce32a22383b30ca88bcedc8cbd6fa2cc90b2a24c183d662288e48103f7021c5e0356e9c44d5369242b5baa1938244ce0d94aeca0c4a28b1774681000021f9040909004a002c0000000040004000860496e484defcc4f6fc3ccefca4eefc6cf2fc44cafc8cf6fc64cefc1cb2fce4fefca4fefc34c2fc54cefc84eafc94f6fcd4fefcb4eefc6ce2fc2cbafc94e2fcc4fefc74fefcb4fefc54dafc94fefc4cdafc6ccafcf4fefc34cafc7cdefc8ce2fcccfafcaceefc64fafc84fefc24b2fc94eefc9cf2fcbcf6fc74e2fc34befcbcfafc5cdefc1caefc84e2fcc4fafc3cd2fca4f2fc4ccafc6cdafcecfefcacfefc3cc6fc5cd2fcdcfefcb4f2fc2cbefc94e6fcccfefc7cfefc54defc9cfefc4ce2fc7ccaf4fcfefc3ccafc6cfefc8cfefc24b6fc9ceafc9cf6fc74e6fcbcfefcffffff00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000007fe804a828384858687820a49490a888e8f90918e0a1795173392999a9b3b173495109ba2a3879d9534a1a4aa843349173b418e9d9f17a9871c9d1598ab8741961715b296a88e159e95b1bc8533a7958da5c7b58894bfbbca831cbf1749c985a69eb68441499fb41cd7dec79f37d0b4e18310daefe8d9dadd83dfc485f5e517f7e8826e447b656816a882dad8012c34aed98573844c7d7a37a3dfb67f0b95503b162c222d828382181b662d2321631f9f0992286d90c0683b940599011151bd53dcf0fd8a2968dc471a356fcdc0388819a8a085207cbca050c98d1d3b20ec681aafd93c6c551f227285f306515fea723a6af8f1eb4b5a1d0d0db394a4a4cb4afe8c883254e4ec10dd5f349220daf1d35205a4e320c845146447857f1c46aaabc4d35062871f058b53752f48d5a5950e3fbadb3789578041ceaecdabf291e5b5cd34cbe4eab0d2674d855183bc6670182c55332a7c147b8daca7bfd714adcbf892d142cbbcd1051ebc8a396593d0a34b5716610382ebd8b35f37221d42920a15be87072fbe42230a00d2ab5fbf1e81f4243e32f8883f5f3e7df93748b0dfaf1e88f40af505785f7d3424c01f7ffe4557817c0cd237a00f34a077207bee45079f7d0f06184a75da75880005dd91379e88e29536dd8928a6181274ce91c2c10742b4b8491044a8d65b083924c0020119edc0030f445c250a0831e8c842020cc8fe28090744f030c28f3eb8a5c90c32b060a59109b400d0053ffef8240f5a69c2810e451c69a5992c0881948b1938d9e59723dc16090e35a069240b4598a06424228df0a59b3f66d0d421333470e6a109248042496b66724f367f3ef9e5028d2a61c39d3aead8c03b41605042a5842d30c23f33f8f0a7973cd0800896673280c3213074d0810627b4289205228890162110b419e90888087166a29fdef202031d20db037341f090eb102258d0689f80fa80880045242a8194833820ebb7260c42c3010f9870c005e2e6aaaeaa36d1e0e40826129258bc84dc202bb20c6850d3013ff4fbc303d8e02a02b443d08bcd6bab4880ecb7e80ec2afbf000f5201b4eaa66680a20a0b77908304853cdc6fc43d8da06eae422aa7c1bdb2bee3f1bf85dcf06cae3cec298a091977e080212b833c880f230fa11774330c80f20b52e66cc80c143f0bea2ade661c2ecefeb2acd6c0cf2e609202df6aac0151461b1284c04af798ecb72a20d27541238b30e8351cbc3036c76647ad732144506d81cc9280d0430748702bc8d98e6530c4086babf8b7dc86abb2f21189935241d48d35bec9050514f0734681000021f90409090055002c0000000040004000860496e484d6fcc4f2fc3ccefc8ceefc54eafca4eefc6ccefc34b6f4e4fefca4fefc1caaf44ce2fc84fefcd4fefc9ce6fc6ce2fc6ceefc4cc2fcc4fafc4cd6fcbceefc74d6fcb4fefc0ca2ec94eefc34c6fc2cb2f494fefc84eafcf4fefc5cdefcaceafc7cfefcccfafc7cd6fc0c9eec8cdefc44d6fc64fafcacf6fc44befcacfafc24aef454e6fcdcfafca4e6fc7cf2fc54dafcb4f6fcbcfafc94f6fc9cfafc049ae484dafcccf6fc44d2fc84f6fc64eefca4f2fc6cd2fc34bafcecfefc1caefc54defc8cfefc9ceafc7ce6fc6cf2fc5ccafcc4fefc4cdafcbcf2fc74dafc14a6ec9cf2fc34cafc2cb6fcfcfefcccfefc7cdafcacfefcdcfefcbcfefc9cfefcffffff00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000007fe80558283848586878209535309888e8f90918e091795173e92999a9b4f1751950e9ba2a3879d9551a1a4aa843e53174f4e8e9d9f17a9871e9d4698ab874e961746b296a88e469e95b1bc853ea7958da5c7b58894bfbbca831ebf1753c985a69eb6844e539fb41ed7dec79f52d0b4e1830edaefe8d9dadd83dfc485f5e517f7e882a4447b656816a882dad8012c34aed98573844c7d7ae7a3dfb67f0bab503b162c222d82839c181b662d2321631f9f0992286d90c0684f947998a0f2d6af28dcf0fd8a2968dcc728101179f08171900025006a1c2899eea3c22a529e3c71f0e469bc66f3b05d7d88a80980af00309428eaab594e470d3f927d49aba3a11afe60c1ae40724820ce0445192a727668efcd298824c44dfa55420b86531ce445e4e489917f1e46aaabc4d350821483c1d6b010b4ca6248f79c5cfd68e9f123100b327f55f2e0f326276c87796224c9830d0caa0120e8acaaa16c4b525c174a5044350f800687c152250001dcaf0b78af2acb51ba28272e90027091f1256d801ea06c10fe3a317951d6959d37c9bebd7b4748a024994fbffe7c03eea550dd3fb5ff7e4c42fc20e080041298847bdf68a360021a14e8e08036b8174f3f14aa6344830f3a18617b0e90a6a025530498618107b6c75285fd348244121024f1c17c2fba489f10f9f967e37e0e30f5de8e3cf6d8137beb69e2410640041989130a2cfe07500c2668c0440c19394005150a3cb58a031f38e9240546a2a50007537270892a3e0cc18496673291014053500126986e1a911e634b0cc0449a6816798d0751c0e9e694532a26890c14e089e7002874c9989470bee966143519e2030468a2c944079dcd091a2191fdd9e894fe200281a15a4260a56711d010a42f54fcc34c98b07200d82177566a825b8444c100033a9806da1341841042658448a1c09fc82202049a03a87a0b0bbbee4ac4674e70206c080d04615d63b05211052222d8c9c4103a0ec241b4bbaa30489b0a1c3b6b1553607bedbb8678d026070a942b082e91162205babc7643c509049fa04048c1cadb80befbe2a5cc0b00e33a70c107a1e373adb0dfee6804c02f143231c115f74485bcc29e9a91133a007ceac7060f276c03d872a0a8242a003c83212c873c4814170fdb9e0fd0465b807439d7db00cc486baaca0c005f7048d186c47b6d034e67e403c0111405354341c0fcb2d2a23c0130b1846c5d88032487d02f2f1e14106d0e8e98edb1bc41ccec880344b090c39c7273aa400354aced63df3e6ac272c6858ff244c1276495b824461c4d363a81000021f9040909005c002c0000000040004000860496e484defc54cefcc4eefca4eafc54eafc34befc84f2fca4f6fce4fefc6ceefc1caefc6ccefc4cc2fcc4fafcbcf6fc94e6fc34cafc94f6fc2cbafc4ce2fcb4eefc6cfefcd4fafc84e6fc64fafca4fefc2cb2f4a4f2fc7cf2fc0c9eec8cdefc5cd2fcaceafc64eefc44befc84fefcacf6fcf4fefc24b6fc7cdefcccfafc44cafc94fefc3cc2fc74f2fc24aef46cd6fc4ccafcb4fefc3cc6fc54e6fcdcfefc049ae484e2fc4cdafcccf6fca4eefc5ceafc34c2fc8cf6fca4fafcecfefc6cf2fc1cb2fcc4fefc9ceafc9cfafc2cbefc4ce6fcb4f2fc7cfefcd4fefc8ce6fc64fefcacfefc2cb6f47cf6fc0ca2ec8ce2fc64cefcaceefc64f2fc8cfefcacfafcfcfefcccfefc9cfefc74d6fc54cafcbcfefc3ccafcffffff00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000007fe805c82838485868782095a5a09888e8f90918e093195313e92999a9b56314b95489ba2a3879d954ba1a4aa8403131e599888a69fa987269d41b1ab863e4e00bf0d8e9d9f31b58641a75a55bb860fbfcf15b296a88894c497cc850935cf002ecb8656d7c683555ad33126d98559dd351fa5a7c587489e96e4eb3e1edd4eba83c395ac183261c91eb8758402740300255c4181856655a28170a00b773822da8b017110c182ca2a1a0ab17084467bc6aa04f9c4d29fc84123168620642a20214a0f99dd6ae4e8c2c20507c5a12c27ef933a47267c1c3484e3041020585c1282b230c0201a569020b1425150bd82f83cd63c6a68cb82054f4f08593aa857b709fe6c0b99b3f7292e972a340ac6088208c8d9b37e653c38f4e1970b0276e52aaac453dbb9a25a1081f0ebf7ef02105dcb3179415652152b41d89a4026cf52476d0200a35d0dc446e7ce9b9656f97a8da595c4824a10a9ac7a01911cb845e12dbd518b545b499ef2aeac0236a9b9e82cd1086e28c10bdeab51200438edb62a075b781371aecac4399641c8c7ceb17b811091796330aae8c34604eac2b520c13f8abff09700062820330e2481c181082678600902d2b0d5835945f82026084460e185186288818035e9e56102376428e285100888044b74a5f88916211211818b30be2863042506d8a187207130a288446c6863512adad3880336286824060c06fee8a0844c4e38e093504679887fb10168420f5250e99916d355a4850e1450105945f1cdc70c0d2d8419a6085a3e321731e9a9e2030f6ad63904422746e79d26552c31439d75fed0e6948f5d231d7f418800a89a3a84b4ce70405662dc233e74b06898450cf19a2ab2dd788d15ea1d50c4a50734264815472ca1a54aaa1632da461b9d46c8a8808a20ab205a2891c1117bba4983062b5c91d94de79586c80f75ce10835d26589081ae4a90105c151a5c11ec151ad805a9257cd1f3270512a8c785064a389b4106310cb2d2291d5971ed0a2bdc2a96279322620212c70d9280aee7ee7ad01247047c44baa7027bc5c157e42b485283163245bfce9e06300901a613ec95b5185bcca1aee64e5108c04750ac719f185b6bea4ba8f65beec920575c880fc15edb6a8031407bae06864c1cb2c6b81e1c7358d9f860aeb31638d7f2c003f98c70c39090db6f06637e1cb0c8a5bc7b45b7f49d0badb4871ccd73c1255f212e2948707c2ed03a93f0b5203428bd82c2a434dbef0a8e785d37c6d9be44030916ac202ec854dbabc5154bc0fda4dd52aab204c581273e0a12021f31ace39c040bf42e81000021f9040909004b002c0000000040004000860496e484e2fcc4f6fc3ccefca4eafc64eefc84fefc6ccefc3cc2fce4fefca4fefc54cefc94e6fcd4fafc1cb2fc94f6fcb4eafc64fafc6ce2fc5ccafcb4f6fc84eafcc4fefc4ce2fc54dafc34cafc94fefc74fefc4cd6fc54cafcf4fefc94eefc2cb6fcb4fefcccfafc44d6fca4f2fc64d6fc44cafc9ce6fcdcfefc9cf2fcb4f2fc6cfefc7ce6fc8ceefc5cdefc1caefc8ce2fcc4fafc44d2fcaceefc8cfefc3cc6fcecfefcacfefc5cd6fcd4fefc24b6fc64fefc74defc6ccafcbcfafc84eefcccfefc54defc9cfefc7cfefcfcfefc2cbafcbcfefc6cdafc9ceafc9cf6fcbcf2fcffffff00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000007fe804b82838485868782273a2004888e8f90918e100095004a92999a9b1396003d9ba1a2879d9613a3a88502263a25368ea595a78844402116afa9871e452f0e2f38b09eb387b637212144ba86312fcece2a88b100c48536c8d8b9cb8309bfde35ca86d3d583444621c7c7e1db8347cfce0c871dc38728d8c839ec85093ade2f3ada068d334404dbb164fa0ac1f8e5cb010f43f34c19ca81ec208a84853c2060e8ab41a181843c544466641d46413338be585028a22c42442cdcbb11f0a4a005fe5ecc20045250829121802cf390a36621110c7f21f020909e207306433075e4c186c9a3323264a830f5908467bf600c0ad0a1c7810ef104513c18221f220ffeb690752d1444abd60124ae0ab2d1cf9983017a619e43872c30117bf72c20b25bc42e86188718fc2a9237129104832f1ac23c131922097643679090a010910c2ce65a0662e1aa07994091093d64c345e8c6765bcc55adc924118ab18fb57ee4838368bb3228041e75383661234677a560ac15b70bdea2a046b5b81c910d16c7b7ea33762f687748406c23c7ce7cb0f0e8a1540cd01a2d218a63464aebf3f081c3f9ec46a0f01f73360d68d3810826584c120f34e8e0830d8680200a390051e185166658d412215ce0e1872082f800827195676208091410e28a1f9280e05a84a513233a16a8c8e28a2e1e58e289f79c73e38a23ea38a38c32a2b804101026fe19e1841836a9217c0a46292582068652e526220d71a524e6089810101bec108162f691a49f2e28182066043b68a90f540759c05e241e08b1c30a11e4c9e60dfae4c0d641406c59109e7a4640a8015bd2e29e895e4602c41062ae49e806c3b1d3dc486c41f7880734ac9927a12bdc3057a2e5c0b42361b2b1a7419e62121a8106673ea541499bd452e920afa18aea6c8648aae7106e1562c110436890c3809761132b37e71c74102290eab902ad8678608001c4167b9e76e8506b1a62d89069080a2be0a9c09c21648badb8391861810546040b233e5419a3e95b2840b9840d43604bac06eb18a18110048b5b1050732e5115a98508e12fb6c12e61841003a46b20ee123f19c4eb8bea0e21442113532cc4c570669320111af44bac01cb0a4cb0c5d644e5ed49c3663b8484208b3cb22140b01582662759bb72bfbc4d5cf1c582880414c390dca072bf480b6274c1f5f81c313b36f8fb6f60210f1cf51225cb85110a431b70f5202e7b8d48c684e99b9dd60a38d2f5ce8858c016d38e24a08101a2ca2d32cc6fd96244c20a4e4df794a3ccfd35e292a050b110cb32ae899f37009d50200021f9040909005c002c0000000040004000860496e484dafc3ccefcc4f6fca4e6fc54eafc6cd2fc34bafc84f2fca4fafce4fefc84fefc14a6ec6ceefc4ce2fcd4fefc6ce2fcbcf6fc24b2fcbceefc4cc6fcc4fafc74dafc94eefc94fafc1caef47cfefc4cd6fca4eefc64fafc7cf2fc5cdefc84eafc64eefc34cafcacfafcf4fefc4ce6fcb4fefcccfafc94f6fc9cfafc0c9eec64dafc1caaf474f2fcdcfafc74e2fc5ccafc7cd6fc9ceefc1cb2fc54dafcaceefc049ae494e2fc3cd2fcccf2fcaceafc5ceafc3cbefc8cf6fca4fefcecfefc8cfefc14aaf46cf2fc54defcbcfafc2cb2f4bcf2fc54c6fcc4fefc94fefc1caefc4cdafca4f2fc7cf6fc9ce6fc64f2fc3ccafcacfefcfcfefc54e6fcbcfefcccfefc9cfefc6cdafcdcfefc74e6fc7cdafc9cf2fcffffff00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000007fe805c828384858687821c505013888e8f90918e464a4a413303929a9b9c579733415a9ca3a4875796333316a5ac85271f0259248e2b4aa0aa8e3f142a0739ad88243822c3108e9e3396a2883c00cd41b3bf8555c3d499a6b7b8873acddc27d1850ac307c31b5287b5b7ab862419dc00363fdf8559d4c34c8706b7a18701ee0006f20afd10504f40bc4257525d524748810a770c0e061cb445c4b861200cadf8b4afd0117f37261692b284dab8070851651b34c05f1173220945a827a21821740b094939e0af51cc422b6846b8692be72027fe284423f140e2212ce2c8c1e4720c94322e2458b8b3e1e211891f530d3dd8e1c0010a688740d09431e8c20a2dfe1656b01564c11fc375554ce8455b4848d9b2054c841df44398451143061722210122df4152b0e89d8c04d1dfcb0daa1caa280287e048529c6805e0e490022a93a3988842051182cb97112818b9e4c26348240c1c184c0249eac926341ffae101f6df1229f8ded6c457ca03bdaa7f23514c084908e3657750a15e2a32f4df26a838752425c114ec0e8470e744023578bd58d60fef811d43c0bcab535791ff084b8bcb212cd71d15d1458184809c50419603ad4d245978b3054442024ff0378a14543c6021291b5ef8d387208658ca033e9468e2892636f813160f54d1e28b2ec6d81417547460e38d38e2e80388f801e7e3640a6890e390379a00e273aa25fe999f92060a49e490467ed8e38f3e5251e39339ee28e5774c2e69c26c0f4481e2983e54f6211632a6f9e28c22b6e9e69b8274c8899c9b6098049da05111df445500a181062839a8171511fea280157f2ea0c19d0149f19d0907b242421489fef9a78ad14c19dd7e7322a1a8a2956a60058275ba171d707b46f2809f966a002a109ccae31d979389e7950faee60aea029f41c64a584cf9786a1502e2daaaa53e8c278581744a51c5748b21716a7ec121f2a9ae496071c8034958e143aa8f48719a5e85364420788824d1ea02d0ae63c5bbef46211f868f52e7dda9661aa2c0a72620888415dd762b1c17683ed0a2b6823c075ca0873015dd783a6141ea0ff0b876ebc354f8a9c6b0a3d4f6eb1558d19810f0bb080b9257740c13fcdbc02062f16ec0510ed26314294be1dba9104f24850f00c33bdec97aa5ccc50fbf6dc763c556e42bf377427351c5b4258b44c2cbf0def6f4642c0b4202975490ca0a12ddf69cb5c94c3fe563d34b8d9c84bc874c393617182ae9f14487525d2e2140576b9a8f39b72245cf56c46c48c626a02dc8cd7ae11989025158d1b52357ebf5b6d679d90aa734655f4eca94866bfe8802dff5ed797f4810fa53200021f9040909005c002c0000000040004000860496e484dafc44cefcc4f6fca4eafc34befc54eafc1caefce4fefca4f6fc84f6fc64cefcd4f6fc4cc2fc6ceefc2cbafcb4f6fc0ca2ecc4fafc2cb2f494f6fc7cdefc84e2fc4ce2fcb4eefc34cafc24b6fc6cfefc24aef4a4fefcaceafc44befc64fafcf4fefc6cd6fcd4fefc5ccafcb4fefc1caaf4ccfafc94fefc94e2fc0c9eec54cefcccf2fc3cbefcacf6fc84fefc6cd2fc4ccafc7cf2fc14a6ec74e6fc8ce6fc54e6fc7cfefc24b2f4049ae48cdefca4eefc34c2fc64eefc1cb2fcecfefc64d2fcd4fafc6cf2fc2cbefcbcf6fcc4fefc9cfafc84e6fc54defcb4f2fc3cc6fc74fefcacfefcaceefc64fefcfcfefcdcfefcbcfefcccfefc9cfefc5cd6fcccf6fcacfafc8cfefc54cafc14a6f47ce6fc24b2fcffffff00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000007fe805c8283848586878210484844888e8f90918e121995192792999a9b5a954319359ba2a3879d9647a4a9852342360a218ea695a8883f541a020caa88210617bf0a8e34199f19b4872b07ca43b0bbabbfd05288b243c78549ca3ecad2ce843fd0bf0e4fa596c687213cca593e3e3fdd850a36e04ce4c4e6863507ecec15ef8521f2a019684648d63d420834a8f3a1c19d3f4256c05da060c8a0354122b4ed0bf590d0931e12a114ea54ec62956ceb328ceb48a8884419232d55f328209b8f2c8d58169221b148c172d6762cf441c55908280e114191d8632517539f8e8518a2efa60f918e1048206808ca1227203a702d4441a2954146681c3942c3c5200bfefaaa6ac94a220780192c10dd000102ec86284e09f5022724b0a1104a16323ba7430580c7000a20020b62435f1037461c2af1ab4711c3879eec18c26ec7210f1c20439e81080565be9641a0482ae8a311d08f4254106038480bd5aab1f0ba029b2f58271b98701dab89eb0f18c05537a06d48caded7c6977cf6f7244596e88f396090f4a484e5e37cfbbec02deac704f00022e8605f6b4aecd861fd91889e6301754d501067dc0df4b567826a2d54d14d114bf0e5d343293c6682070f3dc1c47a1d85804300cc7553602a1dea24e2882446024509283291e28a2530c18d4e508c20858c34ce68e308ee1471c38e3cf6d86309234a81e290441689000a37bcfe90e4924a36996414238ed0e2942a5649651157f8a8258f508a28659160a21885754e96c92490224a61e59a5496800017274ed926952fb214e38d78d2f85f897cf659e287a4003aca13527420a8264f4401c5a1924031c5a39a3d84408b51bce9cc0f4c3c8ac21486725764112146124214536cbae9a30f76f3a58a434ac12817843e5aaaa6a532f12a2289b22a278a8b36da01adb33edaeb3b4fc4b922ab51ec39480825c86aaaa6db0d72abb41e09196609d11ad26cb09a963056a223bc4a68b6828450049b28d64908ada63261692171560ae813084481e2bb08d97b6c97863071eaa6e19e5324608f24ca226ec57e8a08021d6c0aaa23abf23a081433bbca8815175f0e19e939423291ac23f486ca05b3453a65ad8a1b97b7abc821fc306d2145b4896fc7286e0ce7902aaaabd3a444a62a88b52ddafc44ccba2a4bacbd568e4573093673f103ce25102ca2b1e9568773d35ca849e4c5191e5b02684b634d729551889c4ac6125b9d22d63713c9f62ecc5a2975216187e6f5cb8e3c8db3b22733bd709146072aa7cf74cbf93617e7528a77def6baea08d04c1c1e829045045ea2d635fba94adf876b0e09cf285aeeb92350445104d7fe0402003b`
)

func AddRouterOfLoading2x695405a9Gif(router *gin.Engine) {
    
	utils.GetOther(router, LOADING2X_695405A9_GIF_RELATIVE_PATH, LOADING2X_695405A9_GIF_HEX_CONTENT)
	
}







