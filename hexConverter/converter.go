package hexConverter

import (
	"math"
	"strconv"
	"strings"
)

var hexTable = map[byte]int{
	'0': 0,
	'1': 1,
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
	'A': 10,
	'B': 11,
	'C': 12,
	'D': 13,
	'E': 14,
	'F': 15,
}

type HexConverter interface { // (*) назвал бы IHexConverter (шарпист :)), а имплементацию HexConverter
	Convert(hexString string) string //из сигнатуры следует, что hexString - валидная строка
}

type HexConverterImpl struct {
}

func New() *HexConverterImpl {
	return &HexConverterImpl{}
}

func Convert(hexString string) string {
	s := 0
	hexString = strings.ToUpper(hexString)
	for i := 0; i < len(hexString); i++ {
		s += powInt(16, i) * hexTable[hexString[len(hexString)-i-1]]
	}
	return strconv.Itoa(s)
}

func powInt(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}
