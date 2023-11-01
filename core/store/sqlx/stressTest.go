package sqlx

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type (
	StressTest        struct{ opt *StressTestOptions }
	Option            func(opt *StressTestOptions)
	StressTestOptions struct{ stressKey, stressValue, tableSuffix string }
)

var StressMp = map[string]any{}

func NewStressTest(opts ...Option) *StressTest {
	o := &StressTestOptions{
		stressKey:   "GO_STRESS",
		stressValue: "true",
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
	return context.WithValue(ctx, s.GetStressKey(), s.GetStressValue())
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
func (s *StressTest) IsPreviewRequest(ctx context.Context) bool {
	if value := ctx.Value(s.GetStressKey()); value != nil {
		return value.(string) == s.GetStressValue()
	}
	return false
}

func AddGormStressSuffix(db *gorm.DB) {
	st := NewStressTest()
	if st.IsPreviewRequest(db.Statement.Context) {
		db.Statement.Table = fmt.Sprintf("%s%s", db.Statement.Table, st.GetTableSuffix())
	}
}
