package app

import (
	"GoAlgorithmShuntingYard/internal/evaluation/math/operation"
	"GoAlgorithmShuntingYard/internal/evaluation/math/platform/handler"
	"GoAlgorithmShuntingYard/kit/config"
	"GoAlgorithmShuntingYard/kit/logger"
	"github.com/gorilla/mux"
	"net/http"
)

func RunEndpointEnvExpMath(router *mux.Router, config config.IConfig, newLogger logger.ILogger) {
	url, _ := config.GetString("service.endpoints.envExpMath")
	// Agregar CORS Permitidos
	router.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Origins", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Allow-Methods", "POST,OPTIONS")
		w.Header().Set("Access-Control-Max-Age", "86400")
		w.Header().Set("Connection", "Keep-Alive")
	}).Methods(http.MethodOptions)

	service := operation.NewEnvExprMathService(config, newLogger)
	loggerService := operation.NewLogger(service, newLogger)
	transport := handler.NewEnvExprMathHandler(loggerService, newLogger)

	router.HandleFunc(url, transport.ServerHTTP).Methods(http.MethodPost)
}
