package operation

import (
	"context"
	"errors"
)

type ServiceMock struct {
	WantError bool
}

func (e ServiceMock) Evaluate(ctx context.Context, request Request) (response Response, err error) {
	if e.WantError {
		return Response{}, errors.New("error mock")
	}

	return Response{
		Infix:   "5+4",
		Postfix: "5 4 +",
		Result:  1,
	}, nil
}
