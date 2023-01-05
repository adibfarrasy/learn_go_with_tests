package service

import "hello/service/data"

func Hello(name *string, language *data.Language) string {
	var greet string
	var inner_name string

	if language == nil {
		greet = "Hello"
	} else {
		switch *language {
		case data.English:
			greet = "Hello"
		case data.Spanish:
			greet = "Hola"
		}
	}

	if name == nil || *name == "" {
		if language == nil {
			inner_name = "world"
		} else {
			switch *language {
			case data.English:
				inner_name = "world"
			case data.Spanish:
				inner_name = "mundo"
			}
		}
	} else {
		inner_name = *name
	}

	return greet + ", " + inner_name
}
