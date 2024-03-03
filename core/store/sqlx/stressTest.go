package sqlx

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"strings"
)

type (
	StressTest        struct{ opt *StressTestOptions }
	Option            func(opt *StressTestOptions)
	StressTestOptions struct{ stressKey, stressValue, tableSuffix string }
)

var StressMp = map[string]any{}

func NewStressTest(opts ...Option) *StressTest {
	o := &StressTestOptions{
		stressKey: "GO_STRESS",
		//stressValue: "true",
		tableSuffix: "_stress",
	}
	for _, opt := range opts {
		opt(o)
	}

	return &StressTest{opt: o}
}

func (s *StressTest) Register() {
	StressMp[s.GetStressKey()] = s.GetStressValue()
}

func (s *StressTest) SetStressContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, s.GetStressKey(), true)
}

func WithStressKey(val string) Option {
	return func(opt *StressTestOptions) {
		opt.stressKey = val
	}
}

func WithStressVal(val string) Option {
	return func(opt *StressTestOptions) {
		opt.stressValue = val
	}
}

func WithTableSuffix(val string) Option {
	return func(opt *StressTestOptions) {
		opt.tableSuffix = val
	}
}

func (s *StressTest) GetTableSuffix() string {
	return s.opt.tableSuffix
}

func (s *StressTest) GetStressKey() string {
	return s.opt.stressKey
}
func (s *StressTest) GetStressValue() string {
	return s.opt.stressValue
}

// IsPreviewRequest 判断是否是压测请求
func (s *StressTest) isTestRequest(ctx context.Context) bool {
	if val, ok := ctx.Value(s.GetStressKey()).(bool); ok && val {
		return true
	}
	return false
}

func AddGormStressSuffix(db *gorm.DB) {
	st := NewStressTest()
	if st.isTestRequest(db.Statement.Context) {
		db.Statement.Table = fmt.Sprintf("%s%s", db.Statement.Table, st.GetTableSuffix())
	}
}

func (s *StressTest) stressTestCallback(readMode bool) func(db *gorm.DB) {
	return func(db *gorm.DB) {
		if db.DryRun || !s.isTestRequest(db.Statement.Context) {
			return
		}
		if readMode && shouldSkipStressTestsForRead(db.Statement.Context) {
			return
		}
		rawSQL := db.Statement.SQL.String()
		if shouldSkipStressTestsForRead(db.Statement.Context) && isSelectSql(rawSQL) {
			return
		}
		// Use shadow table if rawSQL != ""
		if rawSQL != "" {
			// select * from users as u where id in (select id from t)
			db.Statement.SQL.Reset()
			db.Statement.SQL.WriteString(replaceWithShadowTable(db, rawSQL))
			return
		}
		if db.Statement.TableExpr != nil {
			// Irregular Table, for example:
			// DB.Table("users as u").Find(&users)
			// DB.Table(" users  as u").Find(&users)
			// db.Model(User{}).Select("Name")),Find(&users)
			db.Statement.TableExpr.SQL = replaceWithShadowTable(db, db.Statement.TableExpr.SQL)
			return
		}
		if db.Statement.Table != "" {
			db.Statement.Table = getSuffixedTableName(db, db.Statement.Table, s.GetTableSuffix())
		}
	}

}

func isSelectSql(rawSQL string) bool {
	return len(rawSQL) > 6 && strings.EqualFold(rawSQL[:6], "select")

}

func shouldSkipStressTestsForRead(context context.Context) bool {
	// Implementation of shouldSkipStressTestsForRead function
	return false
}

func replaceWithShadowTable(db *gorm.DB, sql string) string {
	// Implementation of replaceWithShadowTable function
	return ""
}

func getSuffixedTableName(db *gorm.DB, table string, suffix string) string {
	return fmt.Sprintf("%s_%s", table, suffix)
}
