package generate_method

import (
	"bytes"
	"fmt"
	"strings"
)

//
func WriteController(beanMap map[string][]BeanColumn, config *Config) {

	for key, value := range beanMap {

		if _, ok := generateControllerTableMap[key]; !ok {
			continue
		}

		var updateStr string
		for _, v2 := range value {
			if strings.Contains(v2.Name, `create_at`) || strings.Contains(v2.Name, `update_at`) {
				continue
			}
			updateStr = updateStr + `"` + RemoveUnderscoreCapitalize(v2.Name) + `": bean.` + RemoveUnderscoreCapitalize(v2.Name) + ",\n"
		}

		var param = []interface{}{

			//前面
			RemoveUnderscoreCapitalize(key) + `Controller`,

			//第一个函数
			RemoveUnderscoreCapitalize(key) + `Controller`,
			RemoveUnderscoreCapitalize(key) + `Dao`,

			//第二个函数
			RemoveUnderscoreCapitalize(key) + `Controller`,
			RemoveUnderscoreCapitalize(key),
			RemoveUnderscoreCapitalize(key) + `Dao`,

			//第三个函数
			RemoveUnderscoreCapitalize(key) + `Controller`,
			RemoveUnderscoreCapitalize(key),
			RemoveUnderscoreCapitalize(key) + `Dao`,

			//第四个函数
			RemoveUnderscoreCapitalize(key) + `Controller`,
			RemoveUnderscoreCapitalize(key) + `Dao`,
		}

		var template string

		template = fmt.Sprintf(`

package dbcontroller

import (
	"encoding/json"
	"strconv"

	"git.2tianxin.com/platform/tars-go/kissgo/appzaplog"
	"git.2tianxin.com/platform/tars-go/kissgo/appzaplog/zap"

	"git.2tianxin.com/go_tech/zhi_dun/errorenum"
	"git.2tianxin.com/go_tech/zhi_dun/model"
	"git.2tianxin.com/go_tech/zhi_dun/util"
	"git.2tianxin.com/go_tech/zhi_dun/util/audit_httpcall"
	"git.2tianxin.com/go_tech/zhi_dun/util/common"
)

type %s struct {
	Name string
}

func (*%s) Get(param map[string]interface{}) *model.ResData {

	appzaplog.Info(" Get  start ", zap.Any("param", param))

	projectId := util.GetIntValue(param, common.ProjectId)
	if projectId <= 0 {
		return &model.ResData{Msg: "Get error, projectId invalid"}
	}

	pageNum := util.GetInt32Value(param, common.PageNum)
	pageSize := util.GetInt32Value(param, common.PageSize)
	if pageNum <= 0 {
		return &model.ResData{Msg: "pageNum not int or <= 0"}
	}

	if pageSize <= 0 {
		return &model.ResData{Msg: "pageNum not int or <= 0"}
	}

	m := map[string]interface{}{}
	m["project_id"] = projectId

	beans, count, err := dbdao.%s.Get(
		m,
		pageNum, pageSize)
	if err != nil {
		appzaplog.Error(" Get error", zap.Error(err))
		return &model.ResData{Msg: err.Error()}
	}

	return &model.ResData{Datas: beans, Total: int32(count)}
}

func (*%s) Create(param map[string]interface{}) *model.ResData {

	var bean dbbean.%s

	err := util.MapToStruct(param, &bean)
	if err != nil {
		return &model.ResData{Msg: " util.MapToStruct(param, &bean) " + err.Error()}
	}

	err = dbdao.%s.Create(&bean)
	if err != nil {
		return &model.ResData{Msg: "Create error  " + err.Error()}
	}
	return &model.ResData{}

}

func (*%s) Update(param map[string]interface{}) *model.ResData {

	var bean dbbean.%s

	err := util.MapToStruct(param, &bean)
	if err != nil {
		return &model.ResData{Msg: " util.MapToStruct(param, &bean) " + err.Error()}
	}

	err = dbdao.%s.Update(&bean)
	if err != nil {
		return &model.ResData{Msg: "Update error  " + err.Error()}
	}
	return &model.ResData{}
}

func (*%s) Delete(param map[string]interface{}) *model.ResData {

	id := util.GetInt32Value(param, "id")
	if id == 0 {
		return &model.ResData{Msg: "id can not 0  "}
	}

	err := dbdao.%s.Delete(int(id))
	if err != nil {
		return &model.ResData{Msg: "Delete error  " + err.Error()}
	}
	return &model.ResData{}
}



`,
			param...,
		)

		//log.Println(template)

		//log.Println(key, value)
		var b bytes.Buffer

		b.WriteString(template)

		writeFile(config.ControllerDir, key+`_controller.go`, b.Bytes(), config.IsCoverController)

	}

}
