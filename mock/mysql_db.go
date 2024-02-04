package mock

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type MySqlConfig struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
}

func GetMySqlDB(config MySqlConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/bookmark?charset=utf8mb4&parseTime=True&loc=Local", config.User, config.Password, config.Host, config.Port)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Warn),
	})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func ImportDataToMySql(db *gorm.DB, sql string) error {
	err := db.Exec(sql).Error
	if err != nil {
		return err
	}
	return nil
}
