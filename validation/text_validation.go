package validation

import "strings"

func TextValidation(txterror string) string {
	if strings.Contains(strings.ToLower(txterror), "name") {
		return "parameter name tidak boleh kosong"
	} else if strings.Contains(strings.ToLower(txterror), "password") {
		return "parameter password tidak boleh kosong"
	} else if strings.Contains(strings.ToLower(txterror), "username") {
		return "parameter username tidak boleh kosong"
	} else if strings.Contains(strings.ToLower(txterror), "signup") {
		return "parameter register_by tidak boleh kosong"
	} else if strings.Contains(strings.ToLower(txterror), "role") {
		return "parameter role tidak boleh kosong"
	} else {
		return "gagal terhubung keserver"
	}
}
