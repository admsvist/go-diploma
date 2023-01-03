package repository

type FileReader interface {
	ReadFile(string) ([]byte, error)
}

type CountryCodes interface {
	Contains(code string) bool
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
