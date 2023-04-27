package ejercicios

import "strconv"

func Uno(text string) (int, string) {
	value, err := strconv.Atoi(text)
	if err != nil {
		return 0, "Hubo un error " + err.Error()
	}
	if value > 100 {
		return value, "Es mayor a 100"
	} else {
		return value, "Es menos a 100"
	}
}
