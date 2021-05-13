package utils

import "log"

func CheckErr(errs ...error) bool {
	if len(errs) > 0 {
		if nil != errs[0] {
			log.Println("ERROR:", errs)
			return false
		}
	}
	return true
}
