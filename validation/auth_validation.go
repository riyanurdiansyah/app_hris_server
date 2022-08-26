package validation

import "strings"

func SignUpValidation(txterror string) string {
	if strings.Contains(strings.ToLower(txterror), "name") {
		return "parameter name tidak boleh kosong"
	} else if strings.Contains(strings.ToLower(txterror), "password") {
		return "parameter password tidak boleh kosong"
	} else if strings.Contains(strings.ToLower(txterror), "username") {
		return "parameter username tidak boleh kosong"
	} else if strings.Contains(strings.ToLower(txterror), "signup") {
		return "parameter register_by tidak boleh kosong"
	} else if strings.Contains(strings.ToLower(txterror), "created") {
		return "parameter created_at tidak boleh kosong"
	} else if strings.Contains(strings.ToLower(txterror), "updated") {
		return "parameter updated_at tidak boleh kosong"
	} else {
		return "gagal terhubung keserver"
	}
}
