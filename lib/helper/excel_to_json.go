package helper

import (
	"encoding/json"
	"log"
	"os"

	"github.com/tealeg/xlsx"
)

func ExcelConvertJson(excelPath, jsonPath string) {
	// 打开Excel文件
	xlFile, err := xlsx.OpenFile(excelPath)
	if err != nil {
		log.Fatalf("无法打开Excel文件：%v", err)
	}

	// 获取第一个工作表
	sheet := xlFile.Sheets[0]

	// 创建一个切片来保存JSON数据
	var jsonData []map[string]interface{}

	// 获取第一行的标题
	var headers []string
	for _, cell := range sheet.Rows[0].Cells {
		headers = append(headers, cell.String())
	}

	// 遍历除第一行之外的所有行
	for rowIndex := 1; rowIndex < len(sheet.Rows); rowIndex++ {
		rowData := make(map[string]interface{})
		for colIndex, cell := range sheet.Rows[rowIndex].Cells {
			columnName := headers[colIndex]
			rowData[columnName] = cell.String()
		}
		jsonData = append(jsonData, rowData)
	}

	// 将JSON数据编码为字符串
	jsonBytes, err := json.Marshal(jsonData)
	if err != nil {
		log.Fatalf("无法编码为JSON：%v", err)
	}

	// 打印JSON字符串
	//fmt.Println(string(jsonBytes))

	// 保存JSON到文件（可选）
	jsonFile, err := os.Create(jsonPath)
	if err != nil {
		log.Fatalf("无法创建JSON文件：%v", err)
	}
	defer jsonFile.Close()

	_, err = jsonFile.Write(jsonBytes)
	if err != nil {
		log.Fatalf("无法写入JSON文件：%v", err)
	}
}
