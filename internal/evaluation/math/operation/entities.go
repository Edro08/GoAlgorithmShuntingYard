package operation

import (
	"context"
)

type IEvaluation interface {
	Evaluate(ctx context.Context, request Request) (response Response, err error)
}

type Request struct {
	Infix string
}

type Response struct {
	Infix   string
	Postfix string
	Result  float64
}

type ErrResponse struct {
	Message string
}
