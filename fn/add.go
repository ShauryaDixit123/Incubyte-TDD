package fn

import (
	"errors"
	"strconv"
	"strings"
)

func Add(s string) (int, error) {
	sap := ","
	if strings.Contains(s, "//") {
		sap = string(s[2])
		s = strings.Trim(s, "/")
	}
	replacer := strings.NewReplacer("\n", sap)
	s = replacer.Replace(s)
	ar := strings.Split(s, sap)
	total := 0
	for _, v := range ar {
		num := 0
		if len(v) == 0 {
			num = 0
		} else {
			cn, er := strconv.Atoi(v)
			if er != nil {
				return 0, er
			}
			if cn < 0 {
				return 0, errors.New("negetive number not allowed")
			}
			num = cn
		}
		total += num
	}

	return total, nil
}
