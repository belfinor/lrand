package lrand

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.000
// @date    2019-04-17

import (
	"bytes"
	"encoding/binary"
	"io"
	"math/rand"
	"testing"
)

// Test data from http://www.math.sci.hiroshima-u.ac.jp/~m-mat/MT/emt64.html
var expected = []uint64{
	7266447313870364031, 4946485549665804864, 16945909448695747420,
	16394063075524226720, 4873882236456199058, 14877448043947020171,
	6740343660852211943, 13857871200353263164, 5249110015610582907,
	10205081126064480383, 1235879089597390050, 17320312680810499042,
	16489141110565194782, 8942268601720066061, 13520575722002588570,
	14226945236717732373, 9383926873555417063, 15690281668532552105,
	11510704754157191257, 15864264574919463609, 6489677788245343319,
	5112602299894754389, 10828930062652518694, 15942305434158995996,
	15445717675088218264, 4764500002345775851, 14673753115101942098,
	236502320419669032, 13670483975188204088, 14931360615268175698,
	8904234204977263924, 12836915408046564963, 12120302420213647524,
	15755110976537356441, 5405758943702519480, 10951858968426898805,
	17251681303478610375, 4144140664012008120, 18286145806977825275,
	13075804672185204371, 10831805955733617705, 6172975950399619139,
	12837097014497293886, 12903857913610213846, 560691676108914154,
	1074659097419704618, 14266121283820281686, 11696403736022963346,
	13383246710985227247, 7132746073714321322, 10608108217231874211,
	9027884570906061560, 12893913769120703138, 15675160838921962454,
	2511068401785704737, 14483183001716371453, 3774730664208216065,
	5083371700846102796, 9583498264570933637, 17119870085051257224,
	5217910858257235075, 10612176809475689857, 1924700483125896976,
	7171619684536160599, 10949279256701751503, 15596196964072664893,
	14097948002655599357, 615821766635933047, 5636498760852923045,
	17618792803942051220, 580805356741162327, 425267967796817241,
	8381470634608387938, 13212228678420887626, 16993060308636741960,
	957923366004347591, 6210242862396777185, 1012818702180800310,
	15299383925974515757, 17501832009465945633, 17453794942891241229,
	15807805462076484491, 8407189590930420827, 974125122787311712,
	1861591264068118966, 997568339582634050, 18046771844467391493,
	17981867688435687790, 3809841506498447207, 9460108917638135678,
	16172980638639374310, 958022432077424298, 4393365126459778813,
	13408683141069553686, 13900005529547645957, 15773550354402817866,
	16475327524349230602, 6260298154874769264, 12224576659776460914,
	6405294864092763507, 7585484664713203306, 5187641382818981381,
	12435998400285353380, 13554353441017344755, 646091557254529188,
	11393747116974949255, 16797249248413342857, 15713519023537495495,
	12823504709579858843, 4738086532119935073, 4429068783387643752,
	585582692562183870, 1048280754023674130, 6788940719869959076,
	11670856244972073775, 2488756775360218862, 2061695363573180185,
	6884655301895085032, 3566345954323888697, 12784319933059041817,
	4772468691551857254, 6864898938209826895, 7198730565322227090,
	2452224231472687253, 13424792606032445807, 10827695224855383989,
	11016608897122070904, 14683280565151378358, 7077866519618824360,
	17487079941198422333, 3956319990205097495, 5804870313319323478,
	8017203611194497730, 3310931575584983808, 5009341981771541845,
	6930001938490791874, 14415278059151389495, 11001114762641844083,
	6715939435439735925, 411419160297131328, 4522402260441335284,
	3381955501804126859, 15935778656111987797, 4345051260540166684,
	13978444093099579683, 9219789505504949817, 9245142924137529075,
	11628184459157386459, 7242398879359936370, 8511401943157540109,
	11948130810477009827, 6865450671488705049, 13965005347172621081,
	15956599226522058336, 7737868921014130584, 2107342503741411693,
	15818996300425101108, 16399939197527488760, 13971145494081508107,
	3910681448359868691, 4249175367970221090, 9735751321242454020,
	12418107929362160460, 241792245481991138, 5806488997649497146,
	10724207982663648949, 1121862814449214435, 1326996977123564236,
	4902706567834759475, 12782714623891689967, 7306216312942796257,
	15681656478863766664, 957364844878149318, 5651946387216554503,
	8197027112357634782, 6302075516351125977, 13454588464089597862,
	15638309200463515550, 10116604639722073476, 12052913535387714920,
	2889379661594013754, 15383926144832314187, 7841953313015471731,
	17310575136995821873, 9820021961316981626, 15319619724109527290,
	15349724127275899898, 10511508162402504492, 6289553862380300393,
	15046218882019267110, 11772020174577005930, 3537640779967351792,
	6801855569284252424, 17687268231192623388, 12968358613633237218,
	1429775571144180123, 10427377732172208413, 12155566091986788996,
	16465954421598296115, 12710429690464359999, 9547226351541565595,
	12156624891403410342, 2985938688676214686, 18066917785985010959,
	5975570403614438776, 11541343163022500560, 11115388652389704592,
	9499328389494710074, 9247163036769651820, 3688303938005101774,
	2210483654336887556, 15458161910089693228, 6558785204455557683,
	1288373156735958118, 18433986059948829624, 3435082195390932486,
	16822351800343061990, 3120532877336962310, 16681785111062885568,
	7835551710041302304, 2612798015018627203, 15083279177152657491,
	6591467229462292195, 10592706450534565444, 7438147750787157163,
	323186165595851698, 7444710627467609883, 8473714411329896576,
	2782675857700189492, 3383567662400128329, 3200233909833521327,
	12897601280285604448, 3612068790453735040, 8324209243736219497,
	15789570356497723463, 1083312926512215996, 4797349136059339390,
	5556729349871544986, 18266943104929747076, 1620389818516182276,
	172225355691600141, 3034352936522087096, 1266779576738385285,
	3906668377244742888, 6961783143042492788, 17159706887321247572,
	4676208075243319061, 10315634697142985816, 13435140047933251189,
	716076639492622016, 13847954035438697558, 7195811275139178570,
	10815312636510328870, 6214164734784158515, 16412194511839921544,
	3862249798930641332, 1005482699535576005, 4644542796609371301,
	17600091057367987283, 4209958422564632034, 5419285945389823940,
	11453701547564354601, 9951588026679380114, 7425168333159839689,
	8436306210125134906, 11216615872596820107, 3681345096403933680,
	5770016989916553752, 11102855936150871733, 11187980892339693935,
	396336430216428875, 6384853777489155236, 7551613839184151117,
	16527062023276943109, 13429850429024956898, 9901753960477271766,
	9731501992702612259, 5217575797614661659, 10311708346636548706,
	15111747519735330483, 4353415295139137513, 1845293119018433391,
	11952006873430493561, 3531972641585683893, 16852246477648409827,
	15956854822143321380, 12314609993579474774, 16763911684844598963,
	16392145690385382634, 1545507136970403756, 17771199061862790062,
	12121348462972638971, 12613068545148305776, 954203144844315208,
	1257976447679270605, 3664184785462160180, 2747964788443845091,
	15895917007470512307, 15552935765724302120, 16366915862261682626,
	8385468783684865323, 10745343827145102946, 2485742734157099909,
	916246281077683950, 15214206653637466707, 12895483149474345798,
	1079510114301747843, 10718876134480663664, 1259990987526807294,
	8326303777037206221, 14104661172014248293, 15531278677382192198,
	3874303698666230242, 3611366553819264523, 1358753803061653874,
	1552102816982246938, 14492630642488100979, 15001394966632908727,
	2273140352787320862, 17843678642369606172, 2903980458593894032,
	16971437123015263604, 12969653681729206264, 3593636458822318001,
	9719758956915223015, 7437601263394568346, 3327758049015164431,
	17851524109089292731, 14769614194455139039, 8017093497335662337,
	12026985381690317404, 739616144640253634, 15535375191850690266,
	2418267053891303448, 15314073759564095878, 10333316143274529509,
	16565481511572123421, 16317667579273275294, 13991958187675987741,
	3753596784796798785, 9078249094693663275, 8459506356724650587,
	12579909555010529099, 7827737296967050903, 5489801927693999341,
	10995988997350541459, 14721747867313883304, 7915884580303296560,
	4105766302083365910, 12455549072515054554, 13602111324515032467,
	5205971628932290989, 5034622965420036444, 9134927878875794005,
	11319873529597990213, 14815445109496752058, 2266601052460299470,
	5696993487088103383, 6540200741841280242, 6631495948031875490,
	5328340585170897740, 17897267040961463930, 9030000260502624168,
	14285709137129830926, 12854071997824681544, 15408328651008978682,
	1063314403033437073, 13765209628446252802, 242013711116865605,
	4772374239432528212, 2515855479965038648, 5872624715703151235,
	14237704570091006662, 678604024776645862, 12329607334079533339,
	17570877682732917020, 2695443415284373666, 4312672841405514468,
	6454343485137106900, 8425658828390111343, 16335501385875554899,
	5551095603809016713, 11781094401885925035, 9395557946368382509,
	9765123360948816956, 18107191819981188154, 16049267500594757404,
	16349966108299794199, 1040405303135858246, 2366386386131378192,
	223761048139910454, 15375217587047847934, 15231693398695187454,
	12916726640254571028, 8878036960829635584, 1626201782473074365,
	5758998126998248293, 18077917959300292758, 10585588923088536745,
	15072345664541731497, 3559348759319842667, 12744591691872202375,
	2388494115860283059, 6414691845696331748, 3069528498807764495,
	8737958486926519702, 18059264986425101074, 3139684427605102737,
	12378931902986734693, 410666675039477949, 12139894855769838924,
	5780722552400398675, 7039346665375142557, 3020733445712569008,
	2612305843503943561, 13651771214166527665, 16478681918975800939,
	566088527565499576, 4715785502295754870, 6957318344287196220,
	11645756868405128885, 13139951104358618000, 17650948583490040612,
	18168787973649736637, 5486282999836125542, 6122201977153895166,
	17324241605502052782, 10063523107521105867, 17537430712468011382,
	10828407533637104262, 10294139354198325113, 12557151830240236401,
	16673044307512640231, 10918020421896090419, 11077531235278014145,
	5499571814940871256, 2334252435740638702, 18177461912527387031,
	2000007376901262542, 7968425560071444214, 1472650787501520648,
	3115849849651526279, 7980970700139577536, 12153253535907642097,
	8109716914843248719, 3154976533165008908, 5553369513523832559,
	10345792701798576501, 3677445364544507875, 10637177623943913351,
	7380255087060498096, 14479400372337014801, 15381362583330700960,
	204531043189704802, 13699106540959723942, 3817903465872254783,
	10972364467110284934, 2701394334530963810, 2931625600749229147,
	16428252083632828910, 11873166501966812913, 5566810080537233762,
	7840617383807795056, 10699413880206684652, 18259119259617231436,
	10332714341486317526, 10137911902863059694, 669146221352346842,
	8373571610024623455, 10620002450820868661, 12220730820779815970,
	5902974968095412898, 7931010481705150841, 16413777368097063650,
	11273457888324769727, 13719113891065284171, 8327795098009702553,
	10333342364827584837, 6202832891413866653, 9137034567886143162,
	14514450826524340059, 473610156015331016, 813689571029117640,
	13776316799690285717, 10429708855338427756, 8995290140880620858,
	2320123852041754384, 8082864073645003641, 6961777411740398590,
	10008644283003991179, 3239064015890722333, 16762634970725218787,
	16467281536733948427, 10563290046315192938, 5108560603794851559,
	15121667220761532906, 14155440077372845941, 10050536352394623377,
	15474881667376037792, 3448088038819200619, 3692020001240358871,
	6444847992258394902, 8687650838094264665, 3028124591188972359,
	16945232313401161629, 15547830510283682816, 3982930188609442149,
	14270781928849894661, 13768475593433447867, 13815150225221307677,
	8502397232429564693, 718377350715476994, 7459266877697905475,
	8353375565171101521, 7807281661994435472, 16924127046922196149,
	10157812396471387805, 2519858716882670232, 7384148884750265792,
	8077153156180046901, 3499231286164597752, 2700106282881469611,
	14679824700835879737, 14188324938219126828, 3016120398601032793,
	10858152824243889420, 9412371965669250534, 4857522662584941069,
	984331743838900386, 4094160040294753142, 2368635764350388458,
	15101240511397838657, 15584415763303953578, 7831857200208015446,
	1952643641639729063, 4184323302594028609, 16795120381104846695,
	3541559381538365280, 15408472870896842474, 5628362450757896366,
	16277348886873708846, 12437047172652330846, 10172715019035948149,
	1999700669649752791, 6217957085626135027, 11220551167830336823,
	16478747645632411810, 5437280487207382147, 11382378739613087836,
	15866932785489521505, 5502694314775516684, 16440179278067648435,
	15510104554374162846, 15722061259110909195, 10760687291786964354,
	10736868329920212671, 4166148127664495614, 14303518358120527892,
	9122250801678898571, 10028508179936801946, 216630713752669403,
	10655207865433859491, 4041437116174699233, 6280982262534375348,
	297501356638818866, 13976146806363377485, 13752396481560145603,
	11472199956603637419, 16393728429143900496, 14752844047515986640,
	1524477318846038424, 6596889774254235440, 1591982099532234960,
	8065146456116391065, 3964696017750868345, 17040425970526664920,
	11511165586176539991, 3443401252003315103, 16314977947073778249,
	16860120454903458341, 5370503221561340846, 15362920279125264094,
	2822458124714999779, 14575378304387898337, 9689406052675046032,
	2872149351415175149, 13019620945255883050, 14929026760148695825,
	8503417349692327218, 9677798905341573754, 828949921821462483,
	16110482368362750196, 15794218816553655671, 14942910774764855088,
	12026350906243760195, 13610867176871462505, 18324536557697872582,
	2658962269666727629, 327225403251576027, 9207535177029277544,
	8744129291351887858, 6129603385168921503, 18385497655031085907,
	13024478718952333892, 14547683159720717167, 5932119629366981711,
	325385464632594563, 3559879386019806291, 6629264948665231298,
	14358245326238118181, 15662449672706340765, 13975503159145803297,
	3609534220891499022, 4224273587485638227, 9274084767162416370,
	13156843921244091998, 18284750575626858789, 14664767920489118779,
	11292057742031803221, 13919998707305829132, 14473305049457001422,
	9696877879685767807, 1406758246007973837, 2429517644459056881,
	14361215588101587430, 11386164476149757528, 10474116023593331839,
	2921165656527786564, 15604610369733358953, 12955027028676000544,
	10314281035410779907, 3167047178514709947, 1088721329408346700,
	17930425515478182741, 7466411836095405617, 15534027454610690575,
	10879629128927506091, 11502219301371200635, 13915106894453889418,
	4226784327815861027, 12335222183627106346, 3648499746356007767,
	18441388887898023393, 18117929843327093625, 4237736098094830438,
	14229123019768296655, 3930112058127932690, 12663879236019645778,
	9281161952002617309, 4978473890680876319, 845759387067546611,
	1386164484606776333, 8008554770639925512, 11159581016793288971,
	18065390393740782906, 17647985458967631018, 9092379465737744314,
	2914678236848656327, 4376066698447630270, 16057186499919087528,
	3031333261848790078, 2926746602873431597, 7931945763526885287,
	147649915388326849, 15801792398814946230, 5265900391686545347,
	16173686275871890830, 7562781050481886043, 5853506575839330404,
	14957980734704564792, 10944286556353523404, 1783009880614150597,
	9529762028588888983, 822992871011696119, 2130074274744257510,
	8000279549284809219, 3514744284158856431, 128770032569293263,
	3737367602618100572, 16364836605077998543, 783266423471782696,
	4569418252658970391, 11093950688157406886, 14888808512267628166,
	4217786261273670948, 17047486076688645713, 14133826721458860485,
	17539744882220127106, 12394675039129853905, 5757634999463277090,
	9621947619435861331, 1182210208559436772, 14603391040490913939,
	17481976703660945893, 14063388816234683976, 2046622692581829572,
	8294969799792017441, 5293778434844788058, 17976364049306763808,
	399482430848083948, 16495545010129798933, 15241340958282367519,
	989828753826900814, 17616558773874893537, 2471817920909589004,
	11764082277667899978, 9618755269550400950, 1240014743757147125,
	1887649378641563002, 1842982574728131416, 13243531042427194002,
	7688268125537013927, 3080422097287486736, 2562894809975407783,
	12428984115620094788, 1355581933694478148, 9895969242586224966,
	8628445623963160889, 4298916726468199239, 12773165416305557280,
	5240726258301567487, 4975412836403427561, 1842172398579595303,
	7812151462958058676, 17974510987263071769, 14980707022065991200,
	18294903201142729875, 12911672684850242753, 8979482998667235743,
	16808468362384462073, 5981317232108359798, 12373702800369335100,
	16119707581920094765, 2782738549717633602, 15454155188515389391,
	16495638000603654629, 16348757069342790497, 7769562861984504567,
	17504300515449231559, 5557710032938318996, 11846125204788401203,
	13957316349928882624, 2738350683717432043, 15738068448047700954,
	6224714837294524999, 6081930777706411111, 11366312928059597928,
	4355315799925031482, 12393324728734964015, 15277140291994338591,
	1406052433297386355, 15859448364509213398, 1672805458341158435,
	2926095111610982994, 11056431822276774455, 12083767323511977430,
	3296968762229741153, 12312076899982286460, 17769284994682227273,
	15349428916826953443, 1056147296359223910, 18305757538706977431,
	6214378374180465222, 14279648441175008454, 17791306410319136644,
	956593013486324072, 2921235772936241950, 10002890515925652606,
	10399654693663712506, 6446247931049971441, 6380465770144534958,
	11439178472613251620, 10131486500045494660, 3692642123868351947,
	10972816599561388940, 4931112976348785580, 8213967169213816566,
	15336469859637867841, 15026830342847689383, 7524668622380765825,
	17309937346758783807, 372780684412666438, 5642417144539399955,
	18303842993081194577, 11085303253831702827, 15658163165983586950,
	8517521928922081563, 16091186344159989860, 17614656488010863910,
	4736067146481515156, 13449945221374241354, 17755469346196579408,
	13300502638545717375, 6611828134763118043, 14177591906740276597,
	9340430243077460347, 7499765399826404087, 3409518087967832469,
	9013253864026602045, 4444307427984430192, 3729283608700519712,
	13642048880719588383, 16486557958022946240, 2996465014991157904,
	10020049344596426576, 12302485648009883778, 8492591321344423126,
	17407986443716172520, 10530482934957373052, 15740662350540828750,
	1790629986901049436, 6305948377669917188, 15092985352503125323,
	928505047232899787, 14404651977039851607, 7564177565277805597,
	3411236815351677870, 7752718145953236134, 12315979971311483798,
	12477729506691004724, 14654956300924793305, 6689803038918974388,
	1540738812233000153, 13508351811701989957, 15864432023192136053,
	7990997967273843917, 7424300239290765161, 39585249496300263,
	3877436595063283319, 10710642254398044448, 4653804418844456375,
	1232267496410380283, 3690525514009038824, 15459770765077428485,
	13240346522153894145, 5674964360688390624, 16973644653010587289,
	15924280764204855206, 15196708627253442662, 17596174821341373274,
	16196745023027393691, 6980050627399795351, 17582264380857746637,
	18170372407506856324, 12108126025631005514, 15687749089493373169,
	5814107289258228434, 9381977959648494876, 15895601183088112734,
	16267869075651604263, 15228381979765852785, 11949618678312581999,
	4545324791131029438, 582725409406225185, 15282520250746126790,
	14758446535973412711, 7605613563088071833, 1111140641057375915,
	5364843095234852245, 218335432181198977, 4891472444796201742,
	4564628942836375772, 15500501278323817088, 4913946328556108657,
	2684786251736694229, 12090498456116310122, 5310885782157038567,
	5032788439854011923, 12627401038822728242, 11869662610126430929,
	17650156853043540226, 12126672500118808436, 10437658933435653256,
	13133995470637873311, 4601324715591152820, 1874350460376708372,
	5808688626286061164, 13777088437302430376, 5018451954762213522,
	2588296738534474754, 5503414509154170711, 5230497186769951796,
	13261090710400573914, 8515217303152165705, 11074538219737365303,
	15481562385740613213, 12705484409881007350, 14221931471178549498,
	12905633420087112297, 17337759164357146506, 14081997515778175224,
	17384320185513122939, 7131793076779216692, 17483217190312403109,
	900692047897995877, 14723287313048560400, 6132094372965340305,
	7572797575350925726, 12725160700431903514, 380860122911632449,
	1900504978569024571, 8423729759529914138, 7305587201606052334,
	12446871355267313320, 4615812356515386206, 3361817115406652303,
	17690418922000878428, 14632214537567910559, 2709702289926174775,
	3459675155951086144, 7788364399926538150, 16043992474431955950,
	15830963823784930267, 4216893617835797954, 538159724689093771,
	16029152738918251363, 14444848757576686696, 12941757045272633696,
	10900480525147953314, 12547307449905859302, 16001571796892398181,
	407942194622690676, 13873235372903944444, 18071603799493008777,
	1015646077646778622, 9387605808959554815, 11566702442022019410,
	7061722181092883183, 2629032108249254109, 5271820053177594520,
	12640880742139693547, 10098688629735675775, 5716304472850923064,
	3312674502353063071, 7295926377425759633, 833281439103466115,
	16316743519466861667, 9912050326606348167, 11651133878100804242,
	18026798122431692459, 6157758321723692663, 4856021830695749349,
	7074321707293278978, 10748097797809573561, 2949954440753264783,
	9813922580940661152, 9949237950172138336, 15643982711269455885,
	16078663425810239127, 12508044395364228880, 12920301578340189344,
	15368071871011048915, 1610400750626363239, 11994736084146033126,
	6042574085746186088, 4154587549267685807, 15915752367312946034,
	1191196620621769193, 467437822242538360, 2836463788873877488,
	10476401302029164984, 1716169985450737419, 5327734953288310341,
	3994170067185955262, 884431883768190063, 11019001754831208284,
	14322807384384895215, 161011537360955545, 1466223959660131656,
	5227048585229497539, 12410731857504225031, 2142243279080761103,
	17682826799106851430, 1792612570704179953, 14727410295243056025,
	1459567192481221274, 5669760721687603135, 17507918443756456845,
	10354471145847018200, 10362475129248202288, 13143844410150939443,
	6861184673150072028, 18396524361124732580, 543906666394301875,
	12476817828199026728, 11853496871128122868, 12747674713108891748,
	7986179867749890282, 9158195177777627533, 2217320706811118570,
	8631389005200569973, 5538133061362648855, 3369942850878700758,
	7813559982698427184, 509051590411815948, 10197035660403006684,
	13004818533162292132, 9831652587047067687, 7619315254749630976,
	994412663058993407,
}

func TestMersenne(t *testing.T) {
	mt := New()
	mt.SeedFromSlice([]uint64{0x12345, 0x23456, 0x34567, 0x45678})
	for i, want := range expected {
		have := mt.Uint64()
		if have != want {
			t.Errorf("wrong output %d: %d != %d", i, have, want)
		}
	}
}

func TestMersenneReader(t *testing.T) {
	expectedBytes := &bytes.Buffer{}
	for _, val := range expected {
		binary.Write(expectedBytes, binary.LittleEndian, val)
	}

	n := expectedBytes.Len() - 1

	mt := New()
	mt.SeedFromSlice([]uint64{0x12345, 0x23456, 0x34567, 0x45678})
	generatedBytes := make([]byte, n)
	nOut, err := mt.Read(generatedBytes)
	if nOut != n {
		t.Errorf("Read returned wrong length: expected %d, got %d", n, nOut)
	}
	if err != nil {
		t.Error("Read error", err)
	}
	if bytes.Compare(generatedBytes, expectedBytes.Bytes()[:n]) != 0 {
		t.Error("Wrong output")
	}
}

var _ rand.Source = &Mersenne{}

var _ io.Reader = &Mersenne{}

func TestNext(t *testing.T) {
	for i := 0; i < 1000; i++ {
		Next()
	}
}