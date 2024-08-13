package fn

import (
	"errors"
	"strconv"
	"strings"
)

func Add(s string) (int, error) {
	sap := "," // saperator
	if strings.Contains(s, "//") {
		if len(s) < 3 { // min len for saperator should be provided
			return 0, errors.New("minimun length required is 3")
		}
		sap = string(s[2])       // reinitialize saperator value
		s = strings.Trim(s, "/") // remove / from string
	}
	replacer := strings.NewReplacer("\n", sap) // initialzing to replace our \n with sap
	s = replacer.Replace(s)
	ar := strings.Split(s, sap)
	total := 0 // total till now
	for _, v := range ar {
		num := 0
		if len(v) == 0 { // if there is not value
			num = 0
		} else {
			cn, er := strconv.Atoi(v) // convt value to number
			if er != nil {
				return 0, er
			}
			if cn < 0 { // checks for -ve values
				return 0, errors.New("negetive number not allowed")
			}
			num = cn
		}
		total += num // add
	}

	return total, nil
}
