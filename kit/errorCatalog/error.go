package errorCatalog

import "errors"

var (
	ErrInvalidData = errors.New("la expresión no es válida")
	ErrDefault     = errors.New("se ha producido algún error al procesar la expresión")
)
