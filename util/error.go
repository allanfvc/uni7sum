package util

import "strings"

func FormatErrors(err []error) {
	if err != nil && len(err) > 0 {
		message := []string{}
		for _, element := range err {
			message = append(message, element.Error())
		}
		panic(strings.Join(message, ", "))
	}
}
