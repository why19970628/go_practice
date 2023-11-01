package sqlx

import (
	"context"
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"testing"
)

type ActualTable struct {
	ID   uint
	Name string
}

func TestAddStressSuffixTest(t *testing.T) {
	// 连接数据库
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database")
	}

	// 注册 Hooks 钩子函数以修改 SQL 语句
	err = db.Callback().Create().Before("gorm:before_create").Register("add_stress_suffix", AddGormStressSuffix)
	if err != nil {
		panic(err)
		return
	}

	// 在需要进行压测的地方，设置压测标识
	st := NewStressTest(
		WithTableSuffix("_stress_v2"),
		WithStressKey("P_STRESS"),
		WithStressVal("true"),
	)
	ctx := st.SetStressContext(context.Background())

	fmt.Println(st.GetStressKey())

	// 创建或更新数据时，GORM 将在影子表中存储数据
	db.WithContext(ctx).Create(&ActualTable{Name: "Test"})
}
