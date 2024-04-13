package operation

import (
	"GoAlgorithmShuntingYard/kit/constants"
	"GoAlgorithmShuntingYard/kit/logger"
	"context"
	"time"
)

const TitleOperationLogger = "---- LOGGER ENV-EXPRESS-MATH ----"

type Logger struct {
	next   IEvaluation
	logger logger.ILogger
}

func NewLogger(next IEvaluation, logger logger.ILogger) Logger {
	return Logger{
		next:   next,
		logger: logger,
	}
}

func (l Logger) Evaluate(ctx context.Context, request Request) (response Response, err error) {
	defer func(begin time.Time) {
		logTime := "begin" + begin.String() + " " + "end" + time.Now().String() + " " + "took" + time.Since(begin).String()

		if err != nil {
			l.logger.Info(TitleOperationLogger,
				"----------------- Input", request,
				"-----------------  OutPut: ", err.Error(),
				"Long Time: ", logTime,
				"ProcessID -> ", ctx.Value(constants.UUID))
		} else {
			l.logger.Info(TitleOperationLogger,
				"----------------- Input", request,
				"-----------------  OutPut: ", response,
				"Long Time: ", logTime,
				"ProcessID -> ", ctx.Value(constants.UUID))
		}

	}(time.Now())

	if l.next != nil {
		response, err = l.next.Evaluate(ctx, request)
		if err != nil {
			return Response{}, err
		}
		return response, nil
	}
	return Response{}, nil
}
