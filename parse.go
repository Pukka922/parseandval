package parseandval

type parserLib interface {
	BodyParser(out interface{}) error
}

func parse[T any](lib parserLib) (*T, error) {
	form := new(T)

	if err := lib.BodyParser(form); err != nil {
		return nil, err
	}

	return form, nil
}
