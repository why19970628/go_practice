package sqlx

import "gorm.io/gorm"

type Handler func(*gorm.DB)
type Interceptor func(dsn *DSN, op string, options *Config, next Handler) Handler

type processor interface {
	Get(name string) func(*gorm.DB)
	Replace(name string, handler func(*gorm.DB)) error
}

func registerInterceptor(db *gorm.DB, options *Config, interceptors ...Interceptor) {
	dsn, err := parseDSN(options.DSN)
	if err != nil {
		panic(err)
	}
	var processors = []struct {
		Name      string
		Processor processor
	}{
		{"gorm:create", db.Callback().Create()},
		{"gorm:query", db.Callback().Query()},
		{"gorm:delete", db.Callback().Delete()},
		{"gorm:update", db.Callback().Update()},
		{"gorm:row", db.Callback().Row()},
		{"gorm:raw", db.Callback().Raw()},
	}

	for _, interceptor := range interceptors {
		for _, processor := range processors {
			handler := processor.Processor.Get(processor.Name)
			handler = interceptor(dsn, processor.Name, options, handler)

			if err := processor.Processor.Replace(processor.Name, handler); err != nil {
				panic(err)
			}
		}
	}
}
