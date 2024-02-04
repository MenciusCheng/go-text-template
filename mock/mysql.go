package mock

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// 生成并导入测试数据
func GenAndImportNwithdrawalRecords() {

	db, err := GetMySqlDB(MySqlConfig{User: "root", Password: "", Host: "127.0.0.1", Port: "3306"})
	if err != nil {
		panic(err)
	}

	appliedAt := time.Date(2024, 1, 1, 0, 0, 0, 0, time.Local)
	count := 10000
	for i := 0; i < 365; i++ {
		if i > 0 {
			appliedAt = appliedAt.AddDate(0, 0, 1)
		}
		fmt.Println("appliedAt", appliedAt.Format("2006-01-02"))
		recordSql := GenRecordNwithdrawalRecords(count, appliedAt, "_1")
		err = ImportDataToMySql(db, recordSql)
		if err != nil {
			panic(err)
		}
	}
}

// 生成Mock数据SQL
func GenRecordNwithdrawalRecords(count int, appliedAt time.Time, post string) string {
	insertSql := fmt.Sprintf("INSERT INTO `record_nwithdrawal%s` (`app_id`, `type`, `order_id`, `account_id`, `nickname`, `is_delegated`, `id_card`, `real_name`, `phone`, `pre_amount`, `post_amount`, `token`, `bank_account`, `bank`, `bank_branch`, `bank_sub_branch`, `province`, `city`, `applied_at`, `status`, `remark`, `ext`, `created_at`, `updated_at`, `id_card_img_a`, `id_card_img_b`)", post)

	now := time.Now()
	values := make([]string, 0)
	for i := 0; i < count; i++ {
		value := fmt.Sprintf("(1,0,%s,%d,'单元测试',0,'%s','单元测试机器人','%s',20000,18400,20000,'282882288282','他cute','','建设银行','','','%s',2,'4297040 - 0 = 4277040','','%s','%s','beta/resources/UserCertification/1620301787700031254.jpg','beta/resources/UserCertification/1620301794243380175.jpg')",
			GenOrderId(appliedAt, i), GenAccountId(i), GenIdCard(i), GenPhone(i), TimeToStr(appliedAt), TimeToStr(now), TimeToStr(now))
		values = append(values, value)
	}

	return fmt.Sprintf("%s\n VALUES\n %s;", insertSql, strings.Join(values, ",\n"))
}

func GenOrderId(appliedAt time.Time, i int) string {
	startDate := time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local)

	pre := int(appliedAt.Sub(startDate).Hours() / 24)
	return fmt.Sprintf("%d%s", 1000+pre, strconv.Itoa(10000+i))
}

func GenAccountId(index int) int {
	ids := []int{1373836, 1180510, 1325632, 1991576, 1908865, 1760541, 1176978, 1676801, 1316201, 1264059, 1867649, 1117795, 1580701, 1460379, 1908031, 1278532, 1378116, 1046982, 1320226, 1495662, 1066504, 1417812, 1275667, 1043619, 1734675, 1370141, 1160694, 1662479, 1631074, 1047437, 1268674, 1640748, 1588690, 1582763, 1206387, 1925812, 1091804, 1410832, 1169371, 1225108, 1075815, 1841681, 1989312, 1318501, 1285231, 1324510, 1262164, 1672272, 1468953, 1266655, 1660021, 1162025, 1762258, 1425427, 1867004, 1927192, 1466868, 1729978, 1670354, 1128362, 1928177, 1737279, 1976685, 1340674, 1579844, 1718805, 1486834, 1142389, 1252354, 1223311, 1140206, 1945859, 1406462, 1745801, 1850961, 1734038, 1310977, 1171400, 1910741, 1446566, 1375952, 1711846, 1500205, 1910518, 1488144, 1612066, 1553734, 1605508, 1926622, 1239747, 1937279, 1067856, 1362721, 1642683, 1054813, 1843451, 1631691, 1257228, 1347524, 1494444}

	return ids[index%len(ids)]
}

func GenIdCard(index int) string {
	cards := []string{"440397154248485477", "440290963074646871", "440354105994430451", "440823054813290563", "440901443810086030", "440967578990005783", "440967961906518872", "440104776572688059", "440733729360854623", "440490018855913929", "440323040449280214", "440779917745327316", "440350825462003377", "440977976461151079", "440780943168662745", "440993937710678637", "440260269457844480", "440505209417333723", "440454792867424927", "440882997747358403", "440238958532221792", "440876442085243347", "440100636723441439", "440449738420290130", "440993715864719054", "440913633275407623", "440609758130710727", "440182033982035194", "440920536862784371", "440648881843381564", "440571394384985522", "440801131908093051", "440353856129623051", "440952113841575898", "440702768646685427", "440126635951759330", "440576035901470565", "440962069143087750", "440330292068624633", "440737840360840982", "440007115997469885", "440399007759526617", "440897682063285220", "440517089048801190", "440274877351275313", "440047018250731036", "440295619053789177", "440923538022330948", "440824376367756822", "440730366600922070", "440642878819453897", "440002086603764250", "440552643316929232", "440838656444679308", "440757786203916006", "440352147888659971", "440047323028117682", "440549195822100701", "440938601453487847", "440840702905059523", "440544757207961360", "440775719468693201", "440647739697316070", "440669492954144804", "440166431919526267", "440550151942487980", "440781818758445799", "440116032307816320", "440404384370165332", "440097739360792058", "440701818010545489", "440840410166758071", "440428083832518841", "440960659332002055", "440328627296206982", "440540182917517798", "440177092761188989", "440094819308232985", "440152420913457169", "440434351284263821", "440295765241509937", "440836860122174111", "440078047210609832", "440818203413950502", "440861803156348270", "440119859348008854", "440424618215152191", "440600205453815506", "440314816713750394", "440457552816105429", "440364656140241310", "440018559650763965", "440322775895763848", "440676795376905172", "440165039979731920", "440490926654724212", "440043051846657050", "440172122615203957", "440161496951002746", "440476721564223576"}

	return cards[index%len(cards)]
}

func GenPhone(index int) string {
	arr := []string{"15250175384", "11293963213", "14964164126", "13987091273", "14718908990", "16353100376", "10625361146", "19254306663", "19643857966", "14609643245", "17140391433", "19308875724", "19179875521", "11999355149", "17928434767", "17638580098", "17942310440", "15858380449", "16936624067", "18055990744", "17249941473", "18446014934", "16538751273", "19203600280", "11441617336", "19329989251", "19895130964", "15101085788", "19413520089", "12794833995", "16411309927", "12844894563", "17187275819", "11780862402", "15106933488", "11371769580", "13892241819", "12697458155", "14578213393", "15171538779", "16621915963", "12567136937", "11664225820", "18710502857", "11807607761", "12472610210", "15563063509", "16179731948", "18907219285", "13409319491", "12904408179", "13881388049", "15063985013", "15098576496", "15396241261", "12343129242", "16918821387", "18363721571", "11377917604", "18361059237", "11316276309", "14322412085", "17844797186", "12198141730", "16193091131", "19485680848", "16158660611", "14389978940", "15939243021", "14130753895", "17063056337", "18982342658", "16607318550", "14091005482", "15463221536", "10023675082", "13435529014", "19524447096", "11295241194", "16212250810", "16055063220", "15016982208", "11628200902", "11107411043", "19770743066", "16538713951", "17140644961", "18612383999", "17189303654", "15388036997", "18705005882", "11927333278", "15314160308", "18983278748", "17996804508", "17359447371", "13204932154", "10301173236", "13058030010", "16919203928"}

	return arr[index%len(arr)]
}

func TimeToStr(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}
