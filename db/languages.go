package db

func FindLanguages() ([]string, error) {
	query := "select name from gophergala_languages order by name"
	langs := []string{}
	err := db.Select(&langs, query)
	if err != nil {
		return []string{}, err
	}
	return langs, nil
}
