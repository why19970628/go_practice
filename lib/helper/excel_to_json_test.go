package helper

import (
	"testing"
)

func TestExcelConvertJson(t *testing.T) {
	ExcelConvertJson("包袋属性默认值8.30.xlsx", "1.json")
}
