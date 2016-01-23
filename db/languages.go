package db

func FindLanguages() ([]string, error) {
	query := "select name from gophergala_languages"
	langs := []string{}
	err := db.Select(&langs, query)
	if err != nil {
		return []string{}, err
	}
	return langs, nil
}
