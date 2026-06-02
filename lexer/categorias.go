package lexer

import (
	"unicode"
)

func esNumero(s string) bool {
	for _, r := range s {
		if !unicode.IsDigit(r){
			return false
		}
	}
	return true
}

func esReservadaIF(s string) bool {
	if s == "if"{
		return true
	}
	return false
}

func esReservadaWHILE(s string) bool {
	if s == "while"{
		return true
	}
	return false
}

func esBooleano(s string) bool {
	if s == "true" || s == "false"{
		return true
	}
	return false
}

func esIdentificador(s string) bool {
	r := rune(s[0])

	if !unicode.IsLetter(r){
		return false
	}
	for _, r := range s[1:] {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r){
			return false
		}
	}
	return true
}

func Clasificar(s string) string{
	if esReservadaIF(s){
		return "RESERVADA_IF"
	}
	if esReservadaWHILE(s){
		return "RESERVADA_WHILE"
	}
	if esBooleano(s){
		return "BOOLEANO"
	}
	if esNumero(s){
		return "NUMERO"
	}
	if esIdentificador(s){
		return "IDENTIFICADOR"
	}
	return "DESCONOCIDO"
}