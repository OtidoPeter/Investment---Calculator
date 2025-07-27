package fileops

import (
	"fmt"
	"os"
)

func WriteFloatToFile(fileName string, fv float64, rfv float64) {
	valueText := fmt.Sprintf("Future Value: %.1f\nReal Future Value: %.1f\n", fv, rfv)
	os.WriteFile(fileName, []byte(valueText), 0644)
}
