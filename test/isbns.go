package test

// ISBNs is ISBN codes including "1234"
var ISBNs = []string{
	"9780061234002",
	"9780071123419",
	"9780123403520",
	"9780123423207",
	"9780123472250",
	"9780123493538",
	"9780198123484",
	"9780262012638",
	"9780312345617",
	"9780333123492",
	"9780333613559",
	"9780471234425",
	"9780471234852",
	"9780521123440",
	"9780521234160",
	"9780521234375",
	"9780521234382",
	"9780521234474",
	"9780521234733",
	"9780521234740",
	"9780521234801",
	"9780521234887",
	"9780521234900",
	"9780521234917",
	"9780521234979",
	"9780571234059",
	"9780571234967",
	"9780612347311",
	"9780631234302",
	"9780631234432",
	"9780631234517",
	"9780669123418",
	"9780691123424",
	"9780702123474",
	"9780712346467",
	"9780712346474",
	"9780785123415",
	"9780785123422",
	"9780785123477",
	"9780785123491",
	"9780792312505",
	"9780802123459",
	"9780806123462",
	"9780817913311",
	"9780836915228",
	"9780837123417",
	"9780859641333",
	"9781433813290",
	"9781443813150",
	"9781578512362",
	"9782010123412",
	"9782275012391",
	"9783428123445",
	"9783442123414",
	"9783880123472",
	"9784000112345",
	"9784000261234",
	"9784001112344",
	"9784003231234",
	"9784004312345",
	"9784022512345",
	"9784022612342",
	"9784023312340",
	"9784041012345",
	"9784046012340",
	"9784047123403",
	"9784047123410",
	"9784047123427",
	"9784047123434",
	"9784047123441",
	"9784047123458",
	"9784047123465",
	"9784047123472",
	"9784047123489",
	"9784047123496",
	"9784047271234",
	"9784047341234",
	"9784048912341",
	"9784055012348",
	"9784061581234",
	"9784062123426",
	"9784062123433",
	"9784062123440",
	"9784062123457",
	"9784062123471",
	"9784062191234",
	"9784062711234",
	"9784062881234",
	"9784063123401",
	"9784063123418",
	"9784063123425",
	"9784063123432",
	"9784063123449",
	"9784063123456",
	"9784063123463",
	"9784063123470",
	"9784063123487",
	"9784063123494",
	"9784063412345",
	"9784063491234",
	"9784063612349",
	"9784063631234",
	"9784063712346",
	"9784063871234",
	"9784074123469",
	"9784083212345",
	"9784086012348",
	"9784086191234",
	"9784088481234",
	"9784088551234",
	"9784091234018",
	"9784091234025",
	"9784091234032",
	"9784091234049",
	"9784091234056",
	"9784091234063",
	"9784091234070",
	"9784091234087",
	"9784091234094",
	"9784091234100",
	"9784091234223",
	"9784091234230",
	"9784091234247",
	"9784091234254",
	"9784091234278",
	"9784091234285",
	"9784091234292",
	"9784091234308",
	"9784091234346",
	"9784091234353",
	"9784091234360",
	"9784091234384",
	"9784091234391",
	"9784091234407",
	"9784091234438",
	"9784091234445",
	"9784091234452",
	"9784091234469",
	"9784091234476",
	"9784091234483",
	"9784091234490",
	"9784091234506",
	"9784091234612",
	"9784091234629",
	"9784091234636",
	"9784091234643",
	"9784091234650",
	"9784091234667",
	"9784091234674",
	"9784091234681",
	"9784091234704",
	"9784091234735",
	"9784091234759",
	"9784091234766",
	"9784091234773",
	"9784091234780",
	"9784091234797",
	"9784091234803",
	"9784091234810",
	"9784091234827",
	"9784091234834",
	"9784091234841",
	"9784091234858",
	"9784091234865",
	"9784091234872",
	"9784091234889",
	"9784091234896",
	"9784091234902",
	"9784091312341",
	"9784091351234",
	"9784091412348",
	"9784091512345",
	"9784091591234",
	"9784091812346",
	"9784091912343",
	"9784092812345",
	"9784093881234",
	"9784094512342",
	"9784101234137",
	"9784101234151",
	"9784101234175",
	"9784101234205",
	"9784101234229",
	"9784101234236",
	"9784101481234",
	"9784107712349",
	"9784121012340",
	"9784125012346",
	"9784130261234",
	"9784130912341",
	"9784140161234",
	"9784140912348",
	"9784150312343",
	"9784150412340",
	"9784173521234",
	"9784181234157",
	"9784191234239",
	"9784199501234",
	"9784251012340",
	"9784255001234",
	"9784260021234",
	"9784262123417",
	"9784262123424",
	"9784262123486",
	"9784262123493",
	"9784272412341",
	"9784272612345",
	"9784276112346",
	"9784276123410",
	"9784284401234",
	"9784286012346",
	"9784309412344",
	"9784309621234",
	"9784312345196",
	"9784315512342",
	"9784320012349",
	"9784320112346",
	"9784320123403",
	"9784320123410",
	"9784320123427",
	"9784320123434",
	"9784320123441",
	"9784320123458",
	"9784320123465",
	"9784320123472",
	"9784320123489",
	"9784320123496",
	"9784322123425",
	"9784322123432",
	"9784322123463",
	"9784322123470",
	"9784322123487",
	"9784322123494",
	"9784334012342",
	"9784335501234",
	"9784338261234",
	"9784344012349",
	"9784381012340",
	"9784384012347",
	"9784384701234",
	"9784393312346",
	"9784395012343",
	"9784396112349",
	"9784396821234",
	"9784397501234",
	"9784403501234",
	"9784403671234",
	"9784408171234",
	"9784411012340",
	"9784411801234",
	"9784416212349",
	"9784416312346",
	"9784418141234",
	"9784419061234",
	"9784422812342",
	"9784425211234",
	"9784425212347",
	"9784433212346",
	"9784434212345",
	"9784469012347",
	"9784471123444",
	"9784471123451",
	"9784475012348",
	"9784477001234",
	"9784496051234",
	"9784498123458",
	"9784498123496",
	"9784502512346",
	"9784532123406",
	"9784532131234",
	"9784534041234",
	"9784535512344",
	"9784537212341",
	"9784540121234",
	"9784542123403",
	"9784561234210",
	"9784561234418",
	"9784561234463",
	"9784561234524",
	"9784561234609",
	"9784561234944",
	"9784564431234",
	"9784569812342",
	"9784571121234",
	"9784573012349",
	"9784573912342",
	"9784579211234",
	"9784579212347",
	"9784580812345",
	"9784582212341",
	"9784585221234",
	"9784621301234",
	"9784623071234",
	"9784641031234",
	"9784641123434",
	"9784641123441",
	"9784641123458",
	"9784641123472",
	"9784641123489",
	"9784641123496",
	"9784750311234",
	"9784750312347",
	"9784750512341",
	"9784754712341",
	"9784757101234",
	"9784757123410",
	"9784757123427",
	"9784757123434",
	"9784757123441",
	"9784757123458",
	"9784757123465",
	"9784757123472",
	"9784757123489",
	"9784757512344",
	"9784757712348",
	"9784758312349",
	"9784759251234",
	"9784760112340",
	"9784760123407",
	"9784760123414",
	"9784760123438",
	"9784761512347",
	"9784765012348",
	"9784766123449",
	"9784766412345",
	"9784768471234",
	"9784772612340",
	"9784774112343",
	"9784774123400",
	"9784774123417",
	"9784774123424",
	"9784774123479",
	"9784774791234",
	"9784775301234",
	"9784775312346",
	"9784775912348",
	"9784776612346",
	"9784779112348",
	"9784779121234",
	"9784780712346",
	"9784780912340",
	"9784781612348",
	"9784782851234",
	"9784784071234",
	"9784785201234",
	"9784785751234",
	"9784787712349",
	"9784788341234",
	"9784788512344",
	"9784788912342",
	"9784789712347",
	"9784789912341",
	"9784790712343",
	"9784791761234",
	"9784793121234",
	"9784794212344",
	"9784797212341",
	"9784798012346",
	"9784798123424",
	"9784798123431",
	"9784798123448",
	"9784798612348",
	"9784798912349",
	"9784801512344",
	"9784803912340",
	"9784804612348",
	"9784805851234",
	"9784806123453",
	"9784809412349",
	"9784812212349",
	"9784816912344",
	"9784818651234",
	"9784819112345",
	"9784827112344",
	"9784827712346",
	"9784829123409",
	"9784829123454",
	"9784829123478",
	"9784829161234",
	"9784829921234",
	"9784830112348",
	"9784840123419",
	"9784840123440",
	"9784840123464",
	"9784840812344",
	"9784842912349",
	"9784845211234",
	"9784846012342",
	"9784860112349",
	"9784860371234",
	"9784860511234",
	"9784862041234",
	"9784863101234",
	"9784863271234",
	"9784863712348",
	"9784864641234",
	"9784864711234",
	"9784864712347",
	"9784864912341",
	"9784877512347",
	"9784877821234",
	"9784877912345",
	"9784879212344",
	"9784880861234",
	"9784882081234",
	"9784883521234",
	"9784884751234",
	"9784885612343",
	"9784885912344",
	"9784886212344",
	"9784888881234",
	"9784892712340",
	"9784893812346",
	"9784894341234",
	"9784894891234",
	"9784895881234",
	"9784896012347",
	"9784896251234",
	"9784896412345",
	"9784896912340",
	"9784899841234",
	"9784900612341",
	"9784901234009",
	"9784901234061",
	"9784901234115",
	"9784901234139",
	"9784901234160",
	"9784901234214",
	"9784901234313",
	"9784901234344",
	"9784901234375",
	"9784901234405",
	"9784901234436",
	"9784901234498",
	"9784901234504",
	"9784901234511",
	"9784901234573",
	"9784901234696",
	"9784901234719",
	"9784901234740",
	"9784901234849",
	"9784901234887",
	"9784901234900",
	"9784901234962",
	"9784901331234",
	"9784902212341",
	"9784902251234",
	"9784903212340",
	"9784904231234",
	"9784904611234",
	"9784906451234",
	"9784906691234",
	"9784915361234",
	"9784915811234",
	"9784916123404",
	"9784916123435",
	"9784916123442",
	"9784938681234",
	"9784939051234",
	"9784990123406",
	"9784990123413",
	"9784990123420",
	"9784990123437",
	"9784990123444",
	"9784990123475",
	"9784990123482",
	"9784990123499",
	"9785891234666",
	"9787226012345",
	"9787801234452",
	"9789041112347",
	"9789048123490",
	"9789264112346",
	"9789264123410",
	"9789264123496",
	"9789812347992",
}
