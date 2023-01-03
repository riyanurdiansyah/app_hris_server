package validation

import "strings"

func TextValidation(txterror string) string {
	if strings.Contains(strings.ToLower(txterror), "name") {
		return "name tidak boleh kosong"
	} else if strings.Contains(strings.ToLower(txterror), "password") {
		return "password tidak boleh kosong"
	} else if strings.Contains(strings.ToLower(txterror), "username") {
		return "username tidak boleh kosong"
	} else if strings.Contains(strings.ToLower(txterror), "signup") {
		return "register_by tidak boleh kosong"
	} else if strings.Contains(strings.ToLower(txterror), "role") {
		return "role tidak boleh kosong"
	} else if strings.Contains(strings.ToLower(txterror), "employeeid") {
		return "employee_id tidak boleh kosong"
	} else if strings.Contains(strings.ToLower(txterror), "value") {
		return "value tidak boleh kosong"
	} else {
		return txterror
	}
}
