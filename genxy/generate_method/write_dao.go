package generate_method

import (
	"bytes"
	"fmt"
	"strings"
)

func WriteDao(beanMap map[string][]BeanColumn, config *Config) {

	for key, value := range beanMap {
		//		"%s":  bean.%s,
		var updateStr string
		for _, v2 := range value {
			if strings.Contains(v2.Name, `create_at`) || strings.Contains(v2.Name, `update_at`) {
				continue
			}

			if v2.Name == `id` {
				continue
			}

			updateStr = updateStr + `"` + RemoveUnderscoreCapitalize(v2.Name) + `": bean.` + RemoveUnderscoreCapitalize(v2.Name) + ",\n"
		}

		createIdEqual0 := ``
		_, ok := generateDaoTableNotAutoIncrMap[key]
		if ok {

		} else {
			createIdEqual0 = ` 	bean.Id = 0 `
		}

		var desc = ``
		_, ok = generateDaoDesc[key]
		if ok {
			desc = `desc`
		}

		var param = []interface{}{

			//前面三个
			RemoveUnderscoreLowercase(key),
			RemoveUnderscoreCapitalize(key) + `Dao`,
			RemoveUnderscoreLowercase(key),

			//第一个函数 get
			RemoveUnderscoreLowercase(key),
			RemoveUnderscoreCapitalize(key),
			RemoveUnderscoreCapitalize(key),
			RemoveUnderscoreCapitalize(key),
			RemoveUnderscoreCapitalize(key),
			desc,

			//第二个函数 update
			RemoveUnderscoreLowercase(key),
			RemoveUnderscoreCapitalize(key),
			updateStr,

			//第三个函数
			RemoveUnderscoreLowercase(key),
			RemoveUnderscoreCapitalize(key),
			createIdEqual0,

			//第四个函数
			RemoveUnderscoreLowercase(key),
			RemoveUnderscoreCapitalize(key),

			//第五个函数 getById
			RemoveUnderscoreLowercase(key),
			RemoveUnderscoreCapitalize(key),
			RemoveUnderscoreCapitalize(key),
			RemoveUnderscoreCapitalize(key),
			RemoveUnderscoreCapitalize(key),
			RemoveUnderscoreCapitalize(key),

			//第六个函数 GetByMapPage
			RemoveUnderscoreLowercase(key),
			RemoveUnderscoreCapitalize(key),
			RemoveUnderscoreCapitalize(key),
			RemoveUnderscoreCapitalize(key),
			RemoveUnderscoreCapitalize(key),
		}

		var template string

		template = fmt.Sprintf(`

package dbdao

import (
	"git.2tianxin.com/go_tech/zhi_dun/db"
	"git.2tianxin.com/go_tech/zhi_dun/dbbean"
	"git.2tianxin.com/platform/tars-go/kissgo/appzaplog"
	"git.2tianxin.com/platform/tars-go/kissgo/appzaplog/zap"
)

//这个文件是自动生成的，不要修改。当再次自动生成的时候，修改会被冲掉

type %s struct{}

var %s = new(%s)

func (dao *%s) Get(pageNum int32, pageSize int32) ([]dbbean.%s, int, error) {

	var count int
	err := db.Db.Model(&dbbean.%s{}).Count(&count).Error
	if err != nil {
		appzaplog.Error("Get error", zap.Error(err))
		return nil, 0, err
	}

	var beans []dbbean.%s

	err = db.Db.Model(&dbbean.%s{}).Offset((pageNum - 1) * pageSize).Limit(pageSize).Order("id %s").Find(&beans).Error
	if err != nil {
		appzaplog.Error("Get error", zap.Error(err))
		return nil, 0, err
	}

	return beans, count, nil
}

func (dao *%s) Update(bean *dbbean.%s) error {
	if bean.Id == 0 {
		return fmt.Errorf("id can not 0  ")
	}

	m := map[string]interface{}{
		%s
	}

	err := db.Db.Model(&bean).Updates(m).Error
	if err != nil {
		appzaplog.Error("Update error  ", zap.Error(err))
		return err
	}

	return nil
}

func (dao *%s) Create(bean *dbbean.%s) error {
	%s
	bean.CreateAt = time.Now()
	bean.UpdateAt = time.Now()

	err := db.Db.Create(&bean).Error
	if err != nil {
		appzaplog.Error("Create error  ", zap.Error(err))
		return err
	}

	return nil
}

func (dao *%s) Delete(id int) error {
	if id == 0 {
		return fmt.Errorf("id can not 0  ")
	}

	bean := dbbean.%s{
		Id: int(id),
	}

	err := db.Db.Delete(&bean).Error
	if err != nil {
		appzaplog.Error("Delete error  ", zap.Error(err))
		return err
	}

	return nil
}



func (dao *%s) GetById(id int) (dbbean.%s, error) {

	if id == 0 {
		return dbbean.%s{} , nil 
	}

	var retBean dbbean.%s
	retBean.Id = id

	err := db.Db.First(&retBean).Error
	switch {
	case err == gorm.ErrRecordNotFound:
		return dbbean.%s{}, nil
	case err != nil:
		appzaplog.Error("GetById error", zap.Error(err))
		return dbbean.%s{}, err
	default:
		return retBean, nil
	}
}




func (dao *%s) GetByMapPage(m map[string]interface{}, pageNum int32, pageSize int32) ([]dbbean.%s, int, error) {

	var count int
	err := db.Db.Model(&dbbean.%s{}).Where(m).Count(&count).Error
	if err != nil {
		appzaplog.Error("GetByMapPage error", zap.Error(err))
		return nil, 0, err
	}

	var beans []dbbean.%s

	err = db.Db.Model(&dbbean.%s{}).Where(m).Offset((pageNum - 1) * pageSize).Limit(pageSize).Order("id ").Find(&beans).Error
	if err != nil {
		appzaplog.Error("GetByMapPage error", zap.Error(err))
		return nil, 0, err
	}

	return beans, count, nil
}


`,
			param...,
		)

		//log.Println(template)

		//log.Println(key, value)
		var b bytes.Buffer

		b.WriteString(template)

		writeFile(config.DaoDir, key+`_dao_generate.go`, b.Bytes(), true)

	}
}
