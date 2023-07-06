package database

import (
	"fmt"

	_ "github.com/lib/pq"
	"github.com/syamozipc/web_app/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// sql.DBが接続とコネクションプールを持ち、レースコンディションを調整している
// よって、一つのインスタンスを使い回す
// gormの実装者によれば、closeは無くても問題なさそう
// https://github.com/go-gorm/gorm/issues/3145#issuecomment-682502842
var pool *gorm.DB

// TODO: 関数の外の値を書き換えない形にしたい（contextにつめる？）
func Open() error {
	cfg, err := config.New()
	if err != nil {
		return err
	}

	dsn := fmt.Sprintf("%s://%s:%s@%s:%d/%s?sslmode=disable", cfg.Driver, cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Name)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	pool = db

	return nil
}

// TODO: 戻り値の正当性が外部処理に依存するので、要修正
func Pool() *gorm.DB {
	return pool
}
