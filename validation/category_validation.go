package validation

import "strings"

func CategoryValidation(txterror string) string {
	if strings.Contains(strings.ToLower(txterror), "name") {
		return "parameter name tidak boleh kosong"
	} else {
		return "gagal terhubung keserver"
	}
}
