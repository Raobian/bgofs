package common

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var units = map[string]uint64{
	"B":  1 << 0,
	"KB": 1 << 10,
	"K":  1 << 10,
	"MB": 1 << 20,
	"M":  1 << 20,
	"GB": 1 << 30,
	"G":  1 << 30,
	"TB": 1 << 40,
	"T":  1 << 40,
	"PB": 1 << 50,
	"P":  1 << 50,
}
var units_kes = []string{"B", "KB", "MB", "GB", "TB", "PB"}

func StrToSize(sizeStr string) uint64 {
	reg := regexp.MustCompile(`^([\d]+)([\w]*)`)
	allstr := reg.FindStringSubmatch(sizeStr)
	if allstr == nil {
		return 0
	}
	s, _ := strconv.ParseUint(allstr[1], 10, 64)
	if len(allstr) == 2 {
		return s
	}

	unit := strings.ToUpper(allstr[len(allstr)-1])
	return s * units[unit]
}

func SizeToStr(size uint64) string {
	for i := len(units_kes) - 1; i > 0; i-- {
		key := units_kes[i]
		if size >= units[key] {
			return fmt.Sprintf("%.2f%s", float64(size)/float64(units[key]), key)
		}
	}
	return fmt.Sprintf("%v", size)
}
