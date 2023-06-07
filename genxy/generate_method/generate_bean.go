package generate_method

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type DbParam struct {
	User          string
	Password      string
	Host          string
	Port          string
	TableNameLike map[string]TableConfig
	TableSchema   []string
}

type Columns struct {
	TableName      string `gorm:"column:TABLE_NAME" json:"tableName"`
	ColumnName     string `gorm:"column:COLUMN_NAME" `
	DataType       string `gorm:"column:DATA_TYPE" `
	Ordinal        int    `gorm:"column:ORDINAL_POSITION" `
	COLUMN_COMMENT string `gorm:"column:COLUMN_COMMENT" `
	COLUMN_TYPE    string `gorm:"column:COLUMN_TYPE" `
}

type BeanColumn struct {
	Name           string
	Type           string
	Ordinal        int
	COLUMN_COMMENT string
	COLUMN_TYPE    string
	TableConfig    TableConfig
}

type TableConfig struct {
	Type    string
	Ordinal int
}

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile | log.LstdFlags)

}

var (
	generateControllerTableMap = map[string]TableConfig{
		`checking_strategy`: {},
	}

	generateDaoTableNotAutoIncrMap = map[string]struct{}{
		`xymh_project`: {},
	}

	generateDaoDesc = map[string]struct{}{
		`xymh_operate_log`: {},
	}
)

func Run() {

	var (
		DbParam1 = DbParam{}
	)

	DbParam1.TableNameLike = map[string]TableConfig{
		"checking_strategy": {
			Type:    "",
			Ordinal: 0,
		},
		"checking_strategy_channel": {
			Type:    "",
			Ordinal: 0,
		},
		"checking_strategy_history": {
			Type:    "",
			Ordinal: 0,
		},
		//"%voice_stream_block_record": {
		//	Type:    "",
		//	Ordinal: 0,
		//},
	}

	DbParam1.TableSchema = []string{
		"ai",
	}

	beans := GetBeans(DbParam1)

	config := Config{
		BeanDir: "/Users/a123456/xinyu/zhi_dun/src/dbbean/",
		//DaoDir:        "/Users/a123456/xinyu/zhi_dun/src/dbdao/",
		//ControllerDir: "/Users/a123456/xinyu/zhi_dun/src/dbcontroller/",
		//ServiceDir:    "/Users/a123456/xinyu/zhi_dun/src/dbservice/",

		//IsCoverController: true,
		IsCoverController: false,
	}

	WriteBean(beans, &config)
	WriteDaoOneTableProjectId(beans, &config)

	WriteController(beans, &config)
	//WriteService(beans, &config)

	//WriteControllerRegister(beans)

}

func GetBeans(param DbParam) map[string][]BeanColumn {
	path := fmt.Sprintf("%s:%s@tcp(%s:%s)/information_schema?charset=utf8&parseTime=True&loc=Local",
		param.User, param.Password, param.Host, param.Port,
	)

	db, err := gorm.Open("mysql", path)
	if err != nil {
		panic("failed to connect database")
	}
	db.LogMode(true)
	db.SingularTable(true)
	defer db.Close()

	var column []Columns

	whereStr2 := ` ( 1 = 0  `
	for _, v := range param.TableSchema {
		whereStr2 += `  or  TABLE_SCHEMA like '` + v + `' `
	}
	whereStr2 += ` ) `

	//whereStr1 := ` ( 1 = 0 `
	//for k := range param.TableNameLike {
	//	whereStr1 += `  or  TABLE_NAME like '` + k + `' `
	//}
	//whereStr1 += ` ) `

	var beans map[string][]BeanColumn = make(map[string][]BeanColumn)

	for k, config := range param.TableNameLike {
		whereStr1 := `     TABLE_NAME like '` + k + `' `

		db.
			Where(whereStr1).
			Where(whereStr2).
			Order(` TABLE_NAME ,  ORDINAL_POSITION `).
			Find(&column)

		for _, value := range column {
			beanc := BeanColumn{
				Name:           value.ColumnName,
				Type:           value.DataType,
				Ordinal:        value.Ordinal,
				COLUMN_COMMENT: value.COLUMN_COMMENT,
				COLUMN_TYPE:    value.COLUMN_TYPE,
				TableConfig:    config,
			}
			beans[value.TableName] = append(beans[value.TableName], beanc)
		}

	}

	return beans
}
