package excel

import (
	"fmt"
	"github.com/tealeg/xlsx"
	"strings"
	"unicode"
)

const (
	inJson  = "inJson"
	outJson = "outJson"
	//outData      = "outData"
	fileInSystem = "C:\\Users\\Yura\\Desktop\\petProject\\mapping.xlsx"
)

type fieldAttributes struct {
	outData     string
	inData      string
	firstColumn *xlsx.Cell
	thirdColumn *xlsx.Cell
	color       string
}

var outData string
var inData string
var node *Node

func ReadFromXSL() {
	node = &Node{OutValue: "Root", InValue: "Root"}
	excelFileName := fileInSystem
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		fmt.Printf("Error opening Excel file: %s\n", err)
		return
	}
	for _, sheet := range xlFile.Sheets {
		for i, row := range sheet.Rows {
			field := fieldAttributes{
				outData:     outData,
				inData:      inData,
				firstColumn: row.Cells[0],
				thirdColumn: row.Cells[2],
				color:       row.Cells[0].GetStyle().Fill.FgColor,
			}
			if i == 0 {
				outData, inData = createData(node)
				if node.Children != nil {
					node = node.Children
				}
				continue
			} else {
				node = checkBlock(&field, node)
				if node.Children != nil {
					node = node.Children
				}
			}

		}
	}
}

// Проверка по цвету блока
func checkBlock(field *fieldAttributes, node *Node) *Node {
	if field.color != "FFFFFFFF" {
		node = createNewData(field, node)

	} else if field.firstColumn.String() != "" {
		fmt.Printf("SET %s.%s = %s.%s;\n", node.OutValue, field.firstColumn, node.InValue, field.thirdColumn)

	}
	return node
}

// Создание нового объекта
func createNewData(field *fieldAttributes, node *Node) *Node {
	//Проверяем если пустые ячейки
	if field.firstColumn.String() == "" {
		fmt.Println("empty")
		return node
	}
	// Проверяем на наличие массива
	if strings.Contains(field.firstColumn.String(), "Конец блока") {
		closeArray(field, node)
		node.DeleteNode()
		node = node.Parent
		return node
	} else if strings.Contains(field.firstColumn.String(), "[i]") {
		createArray(field, node)
		return node
	}

	in := []rune("in" + field.thirdColumn.String())
	in[2] = unicode.ToUpper(in[2])

	out := []rune("out" + field.firstColumn.String())
	out[3] = unicode.ToUpper(out[3])
	node.AddNode(string(out), string(in))
	fmt.Printf("DECLARE %s REFERENCE TO %s.%s;\n", string(in), inData, field.thirdColumn.String())
	fmt.Printf("CREATE LASTCHILD OF %s NAME '%s';\n", outData, field.firstColumn)
	fmt.Printf("DECLARE %s REFERENCE TO %s.%s;\n", string(out), outData, field.thirdColumn)
	return node
}

func createData(node *Node) (string, string) {
	node.AddNode("outData", "inJson")
	fmt.Printf("CREATE LASTCHILD OF %s NAME 'data';\n", outJson)
	fmt.Printf("DECLARE outData REFERENCE TO %s.data;\n", outJson)
	return "outData", "inJson"
}

func closeArray(field *fieldAttributes, node *Node) {
	fmt.Println("\tEND FOR;")
	fmt.Printf("END IF;\n")
}

func createArray(field *fieldAttributes, node *Node) {
	outBlock := strings.Replace(field.firstColumn.String(), "[i]", "", -1)
	inBlock := strings.Replace(field.thirdColumn.String(), "[i]", "", -1)
	in := []rune("in" + inBlock)
	in[2] = unicode.ToUpper(in[2])

	out := []rune("out" + outBlock)
	out[3] = unicode.ToUpper(out[3])
	fmt.Println()
	fmt.Printf("IF EXISTS(%s.%s[]) THEN\n", node.InValue, inBlock)
	fmt.Printf("\tCREATE LASTCHILD OF %s IDENTITY (JSON.Array)%s;\n", node.OutValue, outBlock)
	fmt.Printf("\tDECLARE %s REFERENCE TO %s.%s;\n", string(out), node.OutValue, outBlock)
	fmt.Printf("\tFOR %s AS %s.%s.Item[] DO\n", string(in), node.InValue, inBlock)
	fmt.Printf("\tCREATE LASTCHILD OF %s.%s AS %s NAME 'Item';\n", node.OutValue, outBlock, string(out))
	node.AddNode(string(out), string(in))
	fmt.Println()
}
