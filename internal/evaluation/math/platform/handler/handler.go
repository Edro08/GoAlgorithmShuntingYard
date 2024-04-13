package handler

import (
	"GoAlgorithmShuntingYard/internal/evaluation/math/operation"
	"GoAlgorithmShuntingYard/kit/constants"
	"GoAlgorithmShuntingYard/kit/errorCatalog"
	"GoAlgorithmShuntingYard/kit/logger"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"io"
	"net/http"
)

const TitleTransportEnvExprMath = "---- TRANSPORT ENV-EXPRESS-MATH ----"

type Handler struct {
	operation operation.IEvaluation
	logger    logger.ILogger
}

func NewEnvExprMathHandler(operation operation.IEvaluation, logger logger.ILogger) Handler {
	return Handler{
		operation: operation,
		logger:    logger,
	}
}

func (h Handler) ServerHTTP(w http.ResponseWriter, r *http.Request) {
	reqInterface, ctx, err := h.Decoder(r.Context(), r)

	if err != nil {
		h.EncoderError(ctx, w, err, nil)
		return
	}

	req := reqInterface.(Request)

	resp, err := h.operation.Evaluate(ctx, h.builderRequest(req))

	if err != nil {
		h.EncoderError(ctx, w, err, req)
		return
	}

	h.Encoder(ctx, w, h.builderResponse(resp))
}

func (h Handler) builderRequest(request Request) operation.Request {
	return operation.Request{
		Infix: request.Infix,
	}
}

func (h Handler) builderResponse(request operation.Response) Response {
	return Response{
		Infix:   request.Infix,
		Postfix: request.Postfix,
		Result:  request.Result,
	}
}

func (h Handler) Decoder(ctx context.Context, r *http.Request) (interface{}, context.Context, error) {
	processID := uuid.New()
	ctxNew := context.WithValue(ctx, constants.UUID, processID.String())

	ip := r.RemoteAddr
	ctxNew = context.WithValue(ctxNew, constants.IP, ip)

	decoder := json.NewDecoder(r.Body)
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(r.Body)

	var requestData Request
	err := decoder.Decode(&requestData)
	if err != nil {
		return nil, ctxNew, err
	}

	return requestData, ctxNew, nil
}

func (h Handler) Encoder(ctx context.Context, w http.ResponseWriter, response interface{}) {
	resp := response.(Response)
	statusCode := http.StatusOK
	h.logger.Info(TitleTransportEnvExprMath, "Status Code", statusCode, "Response", resp, "ProcessID", ctx.Value(constants.UUID))
	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(resp)
}

func (h Handler) EncoderError(ctx context.Context, w http.ResponseWriter, err error, request interface{}) {
	var r ErrResponse
	var statusCode int

	switch {
	case errors.Is(err, errorCatalog.ErrInvalidData):
		req := request.(Request)
		r = ErrResponse{
			Message: fmt.Sprintf("La expresión %v no es válida", req.Infix),
		}
		statusCode = http.StatusBadRequest
	default:
		r = ErrResponse{
			Message: errorCatalog.ErrDefault.Error(),
		}
		statusCode = http.StatusInternalServerError
	}

	h.logger.Info(TitleTransportEnvExprMath, "Status Code", statusCode, "Response", r, "ProcessID", ctx.Value(constants.UUID))
	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(r)
}
