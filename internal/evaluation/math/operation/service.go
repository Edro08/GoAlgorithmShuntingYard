package operation

import (
	"GoAlgorithmShuntingYard/internal/evaluation/math"
	"GoAlgorithmShuntingYard/kit/config"
	"GoAlgorithmShuntingYard/kit/constants"
	"GoAlgorithmShuntingYard/kit/errorCatalog"
	"GoAlgorithmShuntingYard/kit/logger"
	"context"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

const TitleServiceEnvExprMath = "---- SERVICE ENV-EXPRESS-MATH ----"

type Service struct {
	config config.IConfig
	logger logger.ILogger
}

func NewEnvExprMathService(config config.IConfig, logger logger.ILogger) Service {
	return Service{
		config: config,
		logger: logger,
	}
}

func (h Service) Evaluate(ctx context.Context, request Request) (response Response, err error) {
	infix, err := math.NewExpressionInfix(request.Infix)
	if err != nil {
		h.logger.Error(TitleServiceEnvExprMath, "error", err, "ProcessID", ctx.Value(constants.UUID))
		return Response{}, err
	}

	postfix, err := h.createExpressionPostfix(ctx, infix)
	if err != nil {
		return Response{}, err
	}

	result, err := h.evaluateExpressionPostfix(ctx, postfix)
	if err != nil {
		return Response{}, err
	}

	response.Infix = infix.String()
	response.Postfix = postfix.String()
	response.Result = result

	return response, nil
}

func (h Service) createExpressionPostfix(ctx context.Context, infix math.ExpressionInfix) (math.ExpressionPostfix, error) {
	precedence, found := h.config.GetMapInt("config.envExpMath.precedence")
	if !found {
		wrapper := fmt.Errorf("%w: config.envExpMath.precedence no definido", errorCatalog.ErrDefault)
		h.logger.Error(TitleServiceEnvExprMath, "error", wrapper, "ProcessID", ctx.Value(constants.UUID))
		return math.ExpressionPostfix{}, wrapper
	}

	// Definir salida y pila
	var salida []string
	var pila []string

	// Recorrer lista de tokens y ordenar segun Algoritmo Shunting Yard
	for _, token := range h.getTokens(ctx, infix) {
		switch {
		case isNumber(token):
			salida = append(salida, token)
		case precedence[token] > 0:
			for len(pila) > 0 && (precedence[token] <= precedence[pila[len(pila)-1]]) {
				salida = append(salida, pila[len(pila)-1])
				pila = pila[:len(pila)-1]
			}
			pila = append(pila, token)
		}
	}

	for len(pila) > 0 {
		salida = append(salida, pila[len(pila)-1])
		pila = pila[:len(pila)-1]
	}

	postFix, err := math.NewExpressionPostfix(strings.Join(salida, " "))
	if err != nil {
		h.logger.Error(TitleServiceEnvExprMath, "error", err, "ProcessID", ctx.Value(constants.UUID))
		return math.ExpressionPostfix{}, err
	}

	return postFix, nil
}

func (h Service) evaluateExpressionPostfix(ctx context.Context, expression math.ExpressionPostfix) (float64, error) {
	var pila []float64
	for _, token := range strings.Fields(expression.String()) {
		if isNumber(token) {
			number, err := strconv.ParseFloat(token, 64)
			if err != nil {
				h.logger.Error(TitleServiceEnvExprMath, "error", err, "ProcessID", ctx.Value(constants.UUID))
				return 0, err
			}
			pila = append(pila, number)
		} else {
			if len(pila) < 2 {
				wrapper := fmt.Errorf("%w: no hay suficientes operandos para el operador %s", errorCatalog.ErrDefault, token)
				h.logger.Error(TitleServiceEnvExprMath, "error", wrapper.Error(), "ProcessID", ctx.Value(constants.UUID))
				return 0, wrapper
			}

			op2 := pila[len(pila)-1]
			op1 := pila[len(pila)-2]
			pila = pila[:len(pila)-2]

			var result float64
			switch token {
			case "+":
				result = op1 + op2
			case "-":
				result = op1 - op2
			case "*":
				result = op1 * op2
			case "/":
				if op2 == 0 {
					wrapper := fmt.Errorf("%w: división por cero", errorCatalog.ErrDefault)
					h.logger.Error(TitleServiceEnvExprMath, "error", wrapper, "ProcessID", ctx.Value(constants.UUID))
					return 0, wrapper
				}
				result = op1 / op2
			}
			pila = append(pila, result)
		}
	}

	if len(pila) != 1 {
		wrapper := fmt.Errorf("%w: quedaron operandos sin operar", errorCatalog.ErrDefault)
		h.logger.Error(TitleServiceEnvExprMath, "error", wrapper.Error(), "ProcessID", ctx.Value(constants.UUID))
		return 0, wrapper
	}

	return pila[0], nil
}

func isNumber(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

func (h Service) getTokens(ctx context.Context, infix math.ExpressionInfix) []string {
	// Convertir la expresión infija en una lista de tokens
	infixReplace := strings.Replace(infix.String(), " ", "", -1)
	var tokens []string
	var currentToken string

	for _, char := range infixReplace {
		if unicode.IsDigit(char) || char == '.' {
			currentToken += string(char)
		} else {
			if currentToken != "" {
				tokens = append(tokens, currentToken)
				currentToken = ""
			}
			tokens = append(tokens, string(char))
		}
	}

	// Asegurarse de agregar el último token si aún no se ha agregado
	if currentToken != "" {
		tokens = append(tokens, currentToken)
	}
	return tokens
}
