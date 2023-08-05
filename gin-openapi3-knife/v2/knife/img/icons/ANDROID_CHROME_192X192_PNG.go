package icons

import (
	"gitee.com/youbeiwuhuan/knife4go/gin-openapi3-knife/constant"
	"gitee.com/youbeiwuhuan/knife4go/gin-openapi3-knife/utils"
	"github.com/gin-gonic/gin"
)

const (
	ANDROID_CHROME_192X192_PNG_RELATIVE_PATH = constant.ROOT_PATH + "/img/icons/android-chrome-192x192.png"
	// 文件内容的16进制表示
	ANDROID_CHROME_192X192_PNG_HEX_CONTENT = `89504e470d0a1a0a0000000d49484452000000c0000000c0080600000052dc6c070000000467414d410000b18f0bfc6105000000206348524d00007a26000080840000fa00000080e8000075300000ea6000003a98000017709cba513c00000006624b474400ff00ff00ffa0bda7930000000774494d4507e10113070b39d7bef400000023694944415478daed9d69931cc799dfff59475777574fcf00c445516b89020f10c06025c2e1b02ddb726cc4bed9432badd6a414e18db0236c44d82fedcfe12fa0177ee5084710a4246b25ed4a1b41727980047110208021a6670633185c73608eeaaeeeae3bfda2babaabaaabafe9237366ea173131008999cecaca7c9ecc279fe79f404a4a4a4a4a4a4a4a4a4a4a4a4a4a4a4a4a4a4a4a4a4a4a4a4a4a4a4a4a4a4a4a4a4aca4180f4fb0f2f5dbad4f7bf3d68fcfce73fa7acdbb017d277d69b8e1dd4a5f30e53a776ec44de26458fc17ee8df59a7f7d5d631b18e241dfedd61e950dae1efcdffce7a22a4ef2bc2c0ef2bd231a1ce241dbe127fee804213fe4c13be984d821eef0b48df59cff7d5ec9458670a8daff89f3b75ee412469d07ba12f1afa3ef14990f0bec2ef2a7d5fd1f715ff73f37d1120b133c50e5fc1ff6ffeec0126e8d0f0a07713be82ce9dd82408bdaf60a027bdabf064000ed7fba2487e572e6293400afd82b0051101c8dbdf3f7996b834435c2a11978ac4a32228048092c6c74da65309017d6ec259d027da99b431b8298547015716b179248ba700ecc6171a9d4a00d04b972e91714f8284658f084002200390770c7cc376719c002221fefb2413f602d2ab0590630a4027e614a9ff648482c0a30271a9485c2a12878ac43afae9fa5cf3dff93427417c02843b34631fcbfe4f4f12fea4f9b38d413f71534200ef5816e6b2015a7158d8320a801001376411ffa32063132deb1bd0ecd4c9f448cb5001c80050741bc7770dfc2feae162d0e649f7129992a09c998630234fae375a1fef7f2341170182e3bd0fe0bf21ba3422884d80b04b151aff3d93d9a8bf6bbe98ff3754202781961d61b1e3234732105f9f827d6387c1a7379edbc30f34037f5d90f17f427d16ee0eefd2a54b635b0a35ac7f78f04bf0077f16405e33f063cfc30f10356c938300d2eb53204732682db0d9413cba9ed9a8bfdbe8a3600924c03756040015c2ff1eadce9500c8335737e7a48afd3bb68fd16a9df4c6148463193633d047b25cfcc7b2855701e40128f02d70b03f8aefa9464687a58fdc6843be6ce155cbc1df82d5e0a780702c03e98d29e6033f40aad8bf9bb9ba39d7e82709d1e00080d606294cb383894725f5feee1f04cb5b61fd30a000514548e7a701915d0f7b1e4e970dfc8c02050039f8d645426c128c89b6652a801c050a65033ff3284e33eb1891403a3f0da28a2c0d5413c1f256d4fbbb7f201eedfa6e842ebf830020b947d52d65bdfe7b50b8ac1f0a14905ece43fc668e69275b2efe7cc72017e17b812c5a5ea0696146e905424b9ff8da3f0b20bf63908b968b3f67d6211410bf9983f4729e8bc10f0a5759afff3ef7a8ba851e464968fbd1680c9502a085b99dcfc5aa3dcffab900001901d2856990acc0acb329c551dda43f733cccc01f84194427c1b8886f7cb38e8719dda43fa31447d9740640b2fe3b41669c8fde3f62d59e2fcced7c8e84b18cd8a8095a1c1ff8c186c101604b15bb927fa8ff2371a9c1fae14001f1c52cc4d305a6cd705c7c7fdb20ff1e512f10ac33058cc80bc436becdfd59e333f3db06f981e3e2fb2cfb423c5d80f862960beb4f5c6ae41feaff2855ec0afc50b583f63380e644084fd9f804b00158004c00a6ba50be23ef9ab7583f201aad96ce17418a123b2f00646b367dabeee0145a5ea06dbd39cc24e812f3cf00c8d61d9caad9f46dea7f3e934e204509d2f9e278fdde00c8bbe62d75a17c078d710b7f0cdb689f0000da9b1d9c783a8d1f320118006ac4f12aea42f97dc1f276593f2428201c91219d2d328d38b82eceed18f80bb4bcc03836c489831f407ec7c05fb82ece31eb000248678b108e4c3ee69f846079bbea42f97de278150035f863d744cb1304e3bbf53340f3083fc9030413a00ea0967d5c5dc86cd43f63fda001d26b05082714969d2f98367ed4088bc62342cd83b2bd78810ee929cdc84fd9c2aba68d1f8195eda580704281f41adba56898cc46fdb3ece3ea02fcc15f477402443c40705693d479612f60357e490d401540b550d23e126bce13d60f0b0a90bc08f9c23420310c8b52bc5436f037145011f502e1a4b4bd125ffb67006429a0960dfc8d47f112b3079708e40bd320793ec29e62cd795228691fa1314ed1f200163a587fc0b72a00801b376ee0e2c58bc15fe361b7e0054862dd75ddbc24da4795b32084f9ca4f28caf0b66dd01d9bd972c8a3788982cce7653c456bd315596f5ebc7811376edce8ebf7c5c29ee103af1c80c2769dfcab9a8dff0a866b7ff1db2ae4efcdb0488c49680f75f20ff5dfe657f45b0034006500c132c844ec9d844fea3b0de0601914f7023a005d9dd7ae4a65bbc4fab901b42c518e9d25a214d3558bbe6d7b388a642fb057daacbfede168d5a23fa514d36c1e162039f69e378c54b64beabc76158df18976eb1f0cfe36c4f05f625e0068f704220049702881402ceb78ee020422b3ee00529040eb2ebc7593a51738e552f2b490c112a29ddeecf87ebc4042be4f33dd01c0d4f31af933c3c6db88bdbb49229d2b423a33c5eae3231097d60bf3da7bca86b18896f50f2641d8fa47d6fe019dac5378431c4484c27b013dffa07227b365dc64dd017e2f00f2d929d6d1884cdda66fd51cbc047fa9a260c00d71978daf02205773f052dda6ff01be37983c14106664c867f9c9f7c96c1937f30f2a77e00ffaf0da3f1cf9693b000b689b00a11992742660343ea44a3cea87454d778b75278002645a86748e6d3cdaf5f0da6e1d3f843f01e24ba17ec3a2f1a4c420ec99dbade387ae87d7993d6070fe32cd49d8d374b7d485f2fbc4a315b436bfc1d2271ef949ccd2ed365ce25e209800f5c607e9cab3da0365ad7e055c740720be5280708ae98924311dfca566e22cfc411bce16edea053a58ff60f993d54c9c351dfc2558d95e0a08a7b2105fe126ec4995b5fa15e559ed015ad63f087d86233f1dad3fd061022478814e6151bd50d23e16abce2aebde087252e40bd380cc342c7aaa6ce22d8f620aed8763bdc2a289b9fe1ec554e3779e62f6607223d8c030072b8c5875560b25ed63b42f7de261cfae35db1d37527d8445fd0db1e97a5e5624f651e51c0861b6310b108a12a8e6c0dbb2d86d883d7cc32364212fe331fa088b26843d83757f1e4061cb20ffb66ee13f83e1da5f7aa5e01b170ec29ec4a3767eb9f2ebdca3ea1df81b5f0dc91bdfb6b0679c7e8a27c27b81b817c802c8aa0be5ebe6c9dc9bf611e502ebce814820cd16e13ea983565d2693800285aa457f3a95c1578a987814efa151431cfab1a48d6fd67471ac6ad19f36ea0f983c0c514548b345bf0e8303eb2f69d6d7ea42f93a3a5bffc8dabf1b5d2d76c30bc45f525b5894b85400605a27b2b310081b2b156e645e024c0fde33836558f4844bb159c8a084e48cc448939190e70f60ea790d7f653af83118863de5d96948aff2b1f6278ea717ee6bef669e1b0f10b5fe75b4873d7b96a7f61b33095c77d80b98086d88f32b957b99e7c675d61de4f752a37cf205a6e593b2e1e027ba8d6fa19527d45638032416ba6400e4741bdf321cfca4f1df260f058417f82a73cc3c37aee7572af710ddf806599f91987f3fbfafe704886d88c393209800fe8698425717ca1f0aa6bbc9ba93027502e97c9169f9a4ebe13b9a811f21392c1a9e046d27be00729a811fb91ebec3ec0144e2873da7d8a59d87114c77535d287f081a59fa2459febec5cafa72ab091be2e07b64432c551dcb9d9273f68cf23a087b9b418a32e8a609aa3191520100e252bc04823b39091b48b650413f362d3f80c28e8137750bff1d7e92dde4699439ca178f80303422a1f678b947d53fa80be5ab0076e12f7dc2f93e91b87fbfca1c7b3936eab421ae02a8aa25ed53a96a2fb3ee2f00204aa37c5261773a46298e552cbcedfab93b4961d1b6b0a74b51ac58789b521c3bcc7d1746aadacb6a49fb149db33d3be6fb74a3ef8d55429e10909427647994ca82671fcbce7211162d48a01507de26bbb028f5f0a207b2acca5845545211882e7df200a6b6eae44f0c0b7f0b866b7fe9b502e473455ec29e96ba58fe55f6496d0e2deb1f0e7b06d63f31dfa71b7b99de9d4e889b8763ea62f9a6bc63dd61dd7100b858c752205fb3e9db868b13889e1087bf140059c3c5894699639e556379d83f859177ac3bea62f926ba1f7af5bdf10d339085ee11160d5cba443c2a128fd6cd13b959884461dd81242f020e85f7945d4d7f232cba3d95c1d76859aba0df82b5ffd4660d3f311dfc10acb29a0820fff134a4d32a171b5f627be5a9b9ddcbf28ef910adb0673009e2919f8155f9f6dac9f18850902dda0c8be656f57965d3b8caba0303a4d7a7201c635a3e291a367e5cb1701abe750f2243d9c69ff3150ba70d9b61cc9f02c23105d2eb7ca43a0380b2695ccdadeaf368cff7092b3eecc9fa037be8e89817087b83f8864e12eb6ec53a997f83ca02f31e2519019008dcd53acba5d08c43e14d29f892b43c800c5fdd6d6ab386ffc454e2442490ffc5116e244ec4babb56fc6afbb25873d6d0caf50fc7fec3cb9f3d69b2eec9d274c9130a6bd54b62cd71dc8294b18f286f80b017ce10a665785b16e82edbf24942703f27610bade54f7ed7c0c5aa85ff02df1b4c1e0a88dfca23f3e6115f509d35146e7e55fffbfc83ca4d4437be550c11f68c338a41995433100e8b7e2655ec45b6bdd980838c464a71a462e12d87e2050045004587e2858a85b728c511368de22393368c54b117d592f6199237bee1c13f147b5e6b76299f8c668bda1e2012d73a969d8540d82817871b599040ab2ebc0d76e593d4c3294ac87a5ec60680ecb641fe5dddc25f8395b23300e9cc14a4378aac3e3e0271a95158d07ea9acd5e711b5fe89a7bec3c8d18fc203248545c37942d5fc52f996bc63de66d6a3b12796ce1541a699864595aa45ffca72f14dcbc537ab16fd21f5c3a04c1a43a625e6d57461e41df3767ea97c0bad43af78becf9ec39e71868a3624844581f6c331915088824dabd6c9dc792a12366bdc700373228807b84feaccdae0511ca180693878d9b0f1afc12c5903c87c6f06e2b7f95076162c6f67eaeece3bb2663d826ffd8322f7916d7c239f37a2767752956bee05b24faa8bca3a3faa72e26b050827d9aaca552dfc69d5c29f82a5badb49052247ea6eca7afdb3ec93ea22a2290f89ea6ea3f8bca1e3cd5dc2a2f11b2725b1e694cd53b9335416982f36892c806404b80f6b4c9742cc963e00201164fee551d6f2924dc49af3b8787bfbb268b8eb6859ff6009145ffe8ce42aaa911cb8f45b3e291aaee7e6385395dbb141b7d98545994101f16515f27767b8c8f701a54e7e45ff4dfea17e1bc9f93e7d97390ec2a80761a7f2c9e609b15ad2be90cb9c5cb6c199bee5c4e04457358c5cb6e7d592f6053a2b3cf45de63808233b72efb37c52141c2a50422ceb786e960b553955020c17de1abbb0280be473456e521e88436beabcf69eb2692c22baf14d52771be92d9ce35886f456955baedccd6cf3a32ac793c6fdd8e1e46e8530996de3667eb972177b54771b8691265d75098b86f7027e58d47475f364ee3c95043669bfe10666058000ee637661d1892100f23f3f02f18fd85e34d86c8ee96e15ef6cbf2355ecc768dff88e3cec99d01d63a19baa5c0d405559ab2f2b6bf54f4079780d7cdd73353638b95f2dd41eaaacd53f51d6eacb88d6f88e24d7bf1f463e01fa5095aba221635d98d73e116bcec3713cd840446e3ae4645d300e3284f90d9b61c49af3b030af7d8296ac7958dbb36f75b7a1da308e07eb912ddabc6c43305dea2922ec1794f35c944f4e49a06507de7376e593632350779be548dded41e5d7b9c7d5bb18f0528b5132ee587c3c2c1a3e21f62fdb58d0ae4b9a3537e676f40767b79d8f8c40ddedfc3437658e9266cda90b5aa0ee16afef1d5bd833ced8acee40aa7214063faa72226079f0d6d85f893cba876aa8bb71a2ec4c1c4f2f7cadbd9bd9329711bdd4624fea6ec33089d3d84eaa72adb0e8c3ca5c66d3b83681b6f4860f55b9d1c1a3badba6712dffb0328768d873cfea6ec330d609d04355ae79d90628f44249fb4030dc8d713f704fa85f33c0932ac25004aa18054ed4dd0c77a350d23e08a9bb051bdf3dabbb0dd53de37ee07e55e5c49a63bb0559b16794335ca8ca4dcba0cf99aaca0d0f05c43f6aa8bbf151e6e8e51e55ff415d2a5fc308d5dd8661d20969dd54e574b5a45d9174fbc184db9408c9f0a58cb6a76708d4dd327c3c83a4db0fd4927605bd65cd27c644428f03a8caa1a12a771e847df9a4302583ea6c55e5f60c05a4d70b90cfb1b94d350ef1a8a92e967f957d5afb1ac9975aec59dd6d18266d1a7aabca2d95bfe446552eb8148e1375e4be09abbbf161fc7d75b7a5f2971883badb304cecf0a9435834f81e559573bd9a7532374b45c2e626f47023f322e0b255951bbcd180fcdd6948dfe143dd4db03dad706fe7b2bc6325a9bb4d34ecd9d63606fd118e0805e5935155b947d5f9cc469d2f55b9e37c544df58402c271bed4dd321bf5abb947d54eea6e61b1e089f7f044274087b068a027d4ba6cc3d713fa27b1ee3c9d7487b4d1281e91668bdc148f7445f2ef48e3a5c847ac3b4fd592f64f88d6f806258e91753f3059eb0f30d0a0ec234fc80f8bd65dc75565d93ea29ce54255ae28c3dbb64077382e9f0cd4ddbec791badb43fd77f9e588ba1bd3b0671cd603ab97aadc55a9622d306ea30f07aa725de152ddcd5a504bda55245f6a313275b7616092813990aa9c20d8d6f12c37e593b4c65655ae1bd21b5390ce3017dc00001097d60ba5f22f94f57a096356771b06d61ea05be18c9f27f4a0fc95bccd8faa9c7cae0832c359f924f54fae659ed4ddb6cddbf907e5afd0b9c89d49d8330eb31cfc3ecb27255f55ceab5a27f3b3fca8ca51a6aa726d0840e6cd1988dfe248ddeddef63bb2663f42e76ccfb116ba0cd075cc89abcab5658b669fd69694f5da1570f17a01f1d50284939c944f5240389985c8c945d600a8b25ebb927d525b4272b6e7c8d5dd8681e90448289f8c678bb64e88e7b58fc59af398657b8396929c0879b6c8c766532690678b20394ec29e35e7b13aaf7d8cf613dfb62a2f80adf507182e810206509573bdac2458479573dca8caed325695a380f81d15f21fcf7051e6088f3aeab2fe77b9d5c9aabb0d03f38114a293aa5c2b2cbaa05d93cbf67dd60d05e0abcacd3254950bd4dd66b95277bbaf2e68d7d0fb2e5fe6033f80b90700faba7dd2d71372280188d5088bb22f9f5425c064a72a279f2f427a8d8f9407e278d5c27dedbd0c0375b761e0c90300bdb345abf995caddcc96798375430134ca278b108e4e382c4a01e1a8ecdfe8c287f14766cbbc915fa9dc4567ebcfcdc6370c3713a0c78638c813d2e1d18abaa07d2098ee73d66d06054851f2d51626d99302209d9b0629f291a62d98ee737541fb001eada065f5e3075edc6c7c236d67dd80308d8e49d2160d5fb9a42b6bf515e51927aa7214104fab10bf3121a9410a88dfc8417c858f54675050e559fd1365adbe82e8a15758e2844beb0f7036014224a9ca85b345f54269f713b16aafb06e2810941e168149941e66fccfe2a55453acda2b85d26ea0ee16cff69c88badb50ed67dd80385d8ae8c3aa72a2607af03222ec179473fca8cad9e35595a380f48aca93ba9b955faafc3af7b87a0fad8d6f3cdbb33901781bfc00bf1e00e8ac2dda3a1c5bd4ae4bbbd63dd60d05e0cb8fcc4e8f4f7e24906b99e548dd6dd7baa72e36d5dd3a6d7cb90a7bc6616e3993e812168dcaacbb54201e35cc13b9598884dd5d5b4123f322605378cfc6503e4900f9c234a4d32aebc7f49b637b95a9af772f67b6cd65b4dfe61817b8e26ee913c0b307009255e52261d1dcaa7e5f79ce91aadc9931a8ca05ea6e67f85177539e1bd772abfa7d743ff4e272e31b86db09d045552e3c09745054d492f6a1e0df2cc8160a908238fa654a7379c547be8f60b8eb6a49fb1014e1b067c7c1cfabf507385d0205f4a12a272150955325c53ea2bcc183aa9c5094e06d59a0bb23c813a280f8cff2c8bc39c34b99a3975bd5ff415daa5c47b2ac39f332c741e0d60324d0294fc8bf6cc357955b62dd4800230d554e34c4da07926e2f155aea6e9dac3fd332c741e0da030003a8cad91ea8445cfb5876160207aa720519b43a64f924f5f714f2596eca1c4d7541fb65f659fd3ea2852e4cd5dd86810fb3d21fddf284fccb36962a5fca3be657ac1b0aa091ae50dc7bba429066c15399e38ef995ba5409d4dd92ac3fb727be9de0de03003d55e55ae5931e15894b6bd6c9dc797e54e5006f2fe5930490bf3703e9654eca1c6d6fb730b77359deb556c1f8528b913e17eb060c4892aa5c244f28f7a85acaacd73f67ddd000e9b502841303aaca514038a1407a8d9b324764d6eb9fe71e554b48cef761aaee360cfb6602f4a12a17c4a3f54249fb48ac3b4f58b779cf452bac8b6d628875e749a1a47d04442eb5e046dd6da86763dd8041e8503e09c4cb27ebaeebe625c93e9ae547556ec7ee4f558e02e2b755c8df9de125ece9e6572abfcdafe85fa2256c1b843d0d7058e63808cc07c710f4b86ca37c552a5b25d68d04d07fe13a6f05f700a4b255524be5abe0e8528b51b2af3c00d055552e1a16753c0291d8d6f11c1faa728586aadc7af7b0a874b6e85f68c701beba9bf69eb25e5f40f2a5165ca8bb0dc37ef600498533112f905faa7c256f195fb26e28003faa736eaab3aa1c05c88c0cf91c3ff93ef296f1657ea91256770bac3ff7852efdb22f274087f2c9b0c06e15409578b4a22e96df172c6f9b759b7bca1706b28bd37cc82e0a96b7ad2e96df277e9963f836c7b8b0edbedbf8469e93750386a4a7b668f6696d4959e34c55ee544c558e02c229ced4ddd66a57b24f23ea6e5c6a7b0ecbbe9d00095ea0a392845ad23e16abce23d66d6e4a98cfc624cce546d89313e975b1ea3c524bdac7e8aef0b0efad3fb00f37c161065095f3bcac48eca31c954f6a0ebc2d0b00209d2ef8bafe7c9439daf9e5cadfe556ab5f619fa8bb0dc3bef500317aa9cae9ea42f99ac491aa5c708d116fd72f4965fbbeba50be86de61cf7d3df003985bc361e9a12ad7ca1372a900c0b08e672f70a32a6779108e29903859fbfbea6ebb9733cf8d0788863df775be4f370e8a0700922bc7a279422bfa5ce6b9799d754301f8e59317a6215d98e626ec99796e5ecfade87368cff7d957658e83b0ef3d00d0e60592a45482cb3624c17475f354ee1c9504e6d5e54422209c2c7d04d3dd2cded97e47d2eda768cff68c477e0e84f5070ec80400fa2a9f14014852d5b1dc293967cf28aff3503ec905145eee51f50fea42f92a38bdcd715c1ca42550985e61d12b1227aa723c2055ed15d52f73ec15f63c701c180f000c503e697994ca02b58f65cff310166509f1a8a52e967f957d520bd4dde261cf7d57e6380807d503f43a21aee697ca37a55deb2eeb86b246dab5eee697ca3711cdf33f7027be9d3870d66f20553997d6cd13b90b3ca8cab180d85e656a6ef77266db5cc13e56771b8683ea0180deaa727a6e55bfaf6c1a5759379415caa671b5a1eed6edd0ebc05a7fe0804e803e55e51a8533da8762dd5d63dde6492318ee9a5ad23e449f83ff205a7fe0002e81027ac8ac07f78e4962cd719d82243754e50ea4416883c2cdadea7faf3ea8dc4474e35bc5010f7bc6391c2f3cb9662074d986f69954e144556e0248157ba950d23e43b2f50f0ffe03cf81f50040d7f2c9b0b6a828d81e2011c7e244556e9c10971a8545ed17cab3fa3ca2975a24e5fb1c68eb0f1c1e0f100e8bba68e509b56e9f5caadc9677ccdbac1b3a6ee41df3767ea9721bd143af20df279ce979a0077ec081f60040625814683f1c13894745c1a135eb446e960755b9712058deeed4bda6badb2edad5dd822b8d0ef4c637d227ac1b3041e25e2050956b7a81ece3ea4266a3fe19eb868e8bcc46fdb3ece3ea02daad7fbcc6f7c00ffc8003ef01808ed9a240ec700c7eb25cc53c953b4365810f49e61121d69cc7d3b7b6df15fd8b4476e15bffa0d03d58fb1f2aeb0f1c920900f4a52ae75fb661b8ae9b9324fba87216841c0c0f49a9935fd17f9b7fa8df42fba516fb5edd6d180ec60b1e9cae976da825edaa5cb6f950951b0172d92ea9252d5077dbf7975a8c9243e301804154e528a102b17851951b06e2d29a3aafbda76c188b48bed4e250853de31c560fd04d554e07a0e71f54ee647851951b82cc96f165fe41e50edaadff8151771b8643e50180ae61d168b6288524589e6e9ecc9da7929067ddeebd2098ee56f1eece3b52c57e82e8c6377ce875e836be913e62dd008674ab19a801d09567b507ca5afd53d07d681d29a8b256ff5479567b8096e53f54b9fefd702827c000aa727aa1a47d22d69c55d66d1e14b1e6ac164ada27e89ced7960d4dd86ea27d60d6045bfaa7282e97a9e2212fb053e54e5fa8178d4ce3fa8fc3af7a87a07dd65cd0f5dd833cea1f40031fa5095d3ae499a35c7baa1fd2269d69cbaa01d1a75b761d817166d5c0cac2a77223bcb83aa5c3788e3e985fbdabb992d73198744dd6d18520fe0d353552ebf5299cb3c37aeb16e682f32cf8debf995caa152771b86433f01ba944f8627401514ba5a2a7f2818ee06eb36774230dc0db554fe00b4ed36c7b603afd8b31f5a0ef5122860005539db9d92b3f68c72863b55390a2ff7b8fa7b75b1fc050e99badb301c7a0f90408fdb27b52b926e2fb36e641c49b7977ba8bb1dda7c9f6ea41ea0c180aa725e43558e8bf249e251535d2cff2afbb4368743a8ee360ca90768a7d30971cb0b2c956fca3bfca8cac93bd65dd55777eb76e8956e7c13483d40881e6151118dba01e2519178b4de289f64aa2a27d85eb930b77b59de311fa215f60c26c1a150771b86d40324138f0805d9a2adcb3656f5f9cc067b55b9cc867135b7aacfa3fd36c720db330d7b76219d00313a844513f5841aaa72cf58b555acbbcfbaa8bb45d6fdb1674b69902e8112e89227145595ab3b8eab325295a370f30ff5dfe5975375b761483d406fe2aa7241ba7470d9c6e752c55e9c74a3a48abda896b4cf911cf63c54ea6ec3907a800ef4a12ae787456d0f1027ab2a475c6a1416b45f286b4d75b76ef93ea9f5ef42ea017a9314168da449e497cab7e56df3d6a41a246f9bb7f24be5b0ba5b3cdf270d7bf649ea01bad0b7aa1c8528385ed53a999fa522c98db34d82e5ed4cdddb7e47d6ec474856773bf4658e83907a80fee8a52aa7679fd49694f5f1abca29ebf5cfb24f6a4b88467e0eb5badb30a41ea0071d0ec782efcdbd00fcbb06cae6a9dceb5416a6c7d116b1e63c2ade6eaabb051227c112283df4da03e904e8837ecb2745c3f5dc9c24d8479473235795f3a8935fd17f937fa87f85e47c9fb4cc710fa44ba0c1e8543ed93c212e94b42fe4b27d7fd41f2c97edf94249fb02ed27be6999e310a41ea04ffa2c9f1489430908b1ace3b90ba35295230ead15e6b577339bc622a21bdf2475b774e93300a907189c6eaa727e5874b97237b365dc1cd50766b68c9bf9e5ca5db41f7aa5ea6e43927a8001e857550e149260b9ba7932776e585539c1749f17ef6ca7ea6e6322f5007ba3e74df4ca5a7d796855b940dd6dadbe8c437a93fbb8493dc080f4d80b4492e524dd2e9ba772afd28c78642f9f2556ed95e95b5bef09a6b781e8da3f98002ed2b5ff50a413600ff4c8166d5eb621981ef514117b51956ba8bbfdbfdce3ea3db45f6a91863d4744ba041a8e7858b44d665d5dd06eec45554ed2ac7bea8276039d65cdd3b0e708483dc01ee9b2148acaacbb5424148679227ba15f5539e278fad4d7dae586badb2e92f37dd2a5cf08483dc0f07452956be609e51e56bece6c1a5ff4fb0b339bc617b98795afd19eef93aabb8d9874020c410f55b9d65288a25228691f0a7e0e4f5704c35d2f94b40f41514174e993aabb8d81740934247da8ca058972b65be8a12a47e1e51e557faf2e95af2179e39b96398e98d4038c965eb74f5e91747ba9d30f4bbabdd450774b6f739c10a90718017daaca890d5539d77e217b3e5e3e495c6aa88be55f669fd6be4672be4faaee3606520f305a7aabca2d966fcabbd6edf80fcabbd66d753155779b34a90718117d14ce080044e251423caa9b277217d0289f24b6b73335b7fb7fe51d7315d1bb7cd3b0e798493dc0e8094784c2e59375f883ba925bd5ef65b68c8f821fc86c191fe556f57b4033f2130c7c1bd1748774e08f98d4038c909817e8962f04c1f036cc177317884bb5e2ed9dff2d559d4df813a082f632c734db734ca41360c474c91342e83b95aa8eee4c676a92eedc2d9422290f4935be69d8734c70a16f7f4089ab4804fb00d2f87f76f1d6d66f1aff367c70160cfeb8ca43ca18483dc018e8503813d08c1211979ac4a5417e7f1dd1b57fbaf49900a907182fc1600742cb1ff896dd422b08e1c11ff4c1579aed3921f8bae8ed8071e9d2a5a07f0544532382a299f00408a23d81e50fbe52eb3f46d20930661a9320494122bc390e874ec3075ee9c677cca4136002843c413001c2833f207c8a9caefb27443a0126446c1284bf078453abd3c13f21d209304142932020bc046a920efec9914e0046c427433ae853525252525252525252525252525252525252525252525252525252525246c1ff0796795db7b24596500000002574455874646174653a63726561746500323031372d30312d31395430373a31313a35372b30313a3030f45bc1630000002574455874646174653a6d6f6469667900323031372d30312d31395430373a31313a35372b30313a3030850679df000000577a5458745261772070726f66696c65207479706520697074630000789ce3f20c0871562828ca4fcbcc49e5520003230b2e630b1323134b9314031320448034c3640323b35420cbd8d4c8c4ccc41cc407cb8048a04a2e00ea171174f24235950000000049454e44ae426082`
)

func AddRouterOfAndroidChrome192x192Png(router *gin.Engine) {

	utils.GetOther(router, ANDROID_CHROME_192X192_PNG_RELATIVE_PATH, ANDROID_CHROME_192X192_PNG_HEX_CONTENT)

}
