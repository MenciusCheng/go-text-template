package generate_method

import (
	"bytes"
	"fmt"
)

func WriteService(beanMap map[string][]BeanColumn, config *Config) {

	for key := range beanMap {
		//var updateStr string
		//for _, v2 := range value {
		//	if strings.Contains(v2.Name, `create_at`) || strings.Contains(v2.Name, `update_at`) {
		//		continue
		//	}
		//
		//	if v2.Name == `id` {
		//		continue
		//	}
		//
		//	updateStr = updateStr + `"` + RemoveUnderscoreCapitalize(v2.Name) + `": bean.` + RemoveUnderscoreCapitalize(v2.Name) + ",\n"
		//}

		var finalStr string

		finalStr += fmt.Sprintf(`
package dbservice

import (
	"encoding/json"
	"fmt"
	"git.2tianxin.com/go_tech/zhi_dun/db"
	"git.2tianxin.com/go_tech/zhi_dun/dbbean"
	"git.2tianxin.com/go_tech/zhi_dun/dbdao"
	"time"

	"git.2tianxin.com/platform/tars-go/kissgo/appzaplog"
	"git.2tianxin.com/platform/tars-go/kissgo/appzaplog/zap"
	"github.com/go-redis/redis"
)

//这个文件是自动生成的，不要修改。当再次自动生成的时候，修改会被冲掉

type %s struct{}

var %s = new(%s)
`,

			//
			RemoveUnderscoreLowercase(key)+`CacheStruct`,
			RemoveUnderscoreCapitalize(key)+`Cache`,
			RemoveUnderscoreLowercase(key)+`CacheStruct`,
		)

		finalStr += fmt.Sprintf(`
const (
	%s = "%s"
)
`,
			RemoveUnderscoreCapitalize(key)+`IdCacheKey`,
			RemoveUnderscoreCapitalize(key)+`IdCacheKey:id:%d`,
		)

		//
		finalStr += fmt.Sprintf(`

func (this *%s) GetFormRedisOrDbById(id int) (*dbbean.%s, error) {

	key := fmt.Sprintf(%s, id)

	byteSlice, err := db.Redis.Get(key).Bytes()
	switch {
	case err == redis.Nil:
		retBean, err := dbdao.%s.GetById(id)
		if err != nil {
			appzaplog.Error("GetFormRedisOrDbById - this.GetById error",
				zap.Error(err),
				zap.Any("id", id),
			)
			return nil, err
		}
		appzaplog.Debug("GetFormRedisOrDbById - GetById get from db ",
			zap.Any("id", id),
			zap.Any("key", key),
		)

		byteSlice2, err := json.Marshal(&retBean)
		if err != nil {
			appzaplog.Error("GetFormRedisOrDbById - Marshal error",
				zap.Error(err),
				zap.Any("retBean", retBean),
			)
			return nil, err
		}

		db.RedisCheck.Set(key, string(byteSlice2), time.Hour*20)
		return &retBean, nil

	case err != nil:
		appzaplog.Error("GetFormRedisOrDbById - Get ",
			zap.Error(err),
		)

		this.DeleteFormRedisById(id)
		return nil, err
	default:
		var retBean dbbean.%s
		err = json.Unmarshal(byteSlice, &retBean)
		if err != nil {
			appzaplog.Error("GetFormRedisOrDbById - Unmarshal error ",
				zap.Any("byteSlice", byteSlice),
				zap.Any("key", key),
				zap.Error(err),
				// zap.Any("ctx", ctx),
			)
			return nil, err
		}
		appzaplog.Debug("GetFormRedisOrDbById get from redis ",
			zap.Any("retBean", retBean),
			zap.Any("key", key),
			zap.Any("string(byteSlice)", string(byteSlice)),
		)
		return &retBean, nil
	}
}

`,
			//
			RemoveUnderscoreLowercase(key)+`CacheStruct`,

			RemoveUnderscoreCapitalize(key),
			RemoveUnderscoreCapitalize(key)+`IdCacheKey`,
			RemoveUnderscoreCapitalize(key)+`Dao`,
			RemoveUnderscoreCapitalize(key),
		)

		finalStr += fmt.Sprintf(`

func (%s) DeleteFormRedisById(id int) {
	key := fmt.Sprintf("%s", id)
	db.Redis.Del(key)
}
`,

			RemoveUnderscoreLowercase(key)+`CacheStruct`,
			RemoveUnderscoreCapitalize(key)+`IdCacheKey:id:%d`,
		)

		var b bytes.Buffer

		b.WriteString(finalStr)

		writeFile(config.ServiceDir, key+`_service_generate.go`, b.Bytes(), true)

	}

}
