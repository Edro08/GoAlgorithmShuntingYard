package math

import (
	"GoAlgorithmShuntingYard/kit/errorCatalog"
	"fmt"
	"regexp"
)

type ExpressionInfix struct {
	value string
}

func NewExpressionInfix(value string) (ExpressionInfix, error) {

	// Verificar que el value no este vacio
	if value == "" {
		return ExpressionInfix{}, fmt.Errorf("%w: %v", errorCatalog.ErrInvalidData, value)
	}

	// Expresión regular para buscar cualquier cosa entre paréntesis
	regexPa := regexp.MustCompile(`\((.*?)\)`)

	// Buscar la primera coincidencia en el texto
	coincidencia := regexPa.FindStringSubmatch(value)

	// Verificar si se encontró una coincidencia
	if len(coincidencia) > 1 {
		return ExpressionInfix{}, fmt.Errorf("%w: %v", errorCatalog.ErrInvalidData, value)
	}

	// Expresión regular para buscar números enteros positivos y cero, y los operadores aritméticos
	regex := regexp.MustCompile(`^ *[-+*/.0-9 ]* *$`)

	// Buscar si la cadena cumple con la expresión regular
	if !regex.MatchString(value) {
		return ExpressionInfix{}, fmt.Errorf("%w: %v", errorCatalog.ErrInvalidData, value)
	}

	return ExpressionInfix{
		value,
	}, nil
}

func (id ExpressionInfix) String() string {
	return id.value
}

type ExpressionPostfix struct {
	value string
}

func NewExpressionPostfix(value string) (ExpressionPostfix, error) {
	return ExpressionPostfix{
		value,
	}, nil
}

func (id ExpressionPostfix) String() string {
	return id.value
}
