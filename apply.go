package validator

import "errors"

func Apply(title, text string) error {
	if title == "" {
		return errors.New("empty title")
	}
	if len(title) >= 100 {
		return errors.New("title too long")
	}
	if text == "" {
		return errors.New("empty text")
	}
	if len(text) >= 500 {
		return errors.New("text too long")
	}
	return nil
}

//Название не должно быть пустым.
//Название должно быть короче 100 символов.
//Текст объявления не должен быть пустым.
//Текст объявления должен быть короче 500 символов.
