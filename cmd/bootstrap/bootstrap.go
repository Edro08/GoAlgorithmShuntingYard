package bootstrap

import (
	"GoAlgorithmShuntingYard/cmd/bootstrap/app"
	"GoAlgorithmShuntingYard/internal/health"
	"GoAlgorithmShuntingYard/kit/config"
	"GoAlgorithmShuntingYard/kit/logger"
	"github.com/gorilla/mux"
	"net/http"
)

const TitleBootstrap = "---- BOOTSTRAP ----"

func Run() {
	// Inicializar configuraciones y logger de servidor
	newConfig := config.NewConfig("application.yaml")
	newLogger := logger.NewLogger()

	serverName, found := newConfig.GetString("server.name")
	if !found {
		newLogger.Fatal(TitleBootstrap, "error", "server name not found")
	}

	port, found := newConfig.GetString("server.port")
	if !found {
		newLogger.Fatal(TitleBootstrap, "error", "server port not found")
	}

	// Crear instancia de Router Mux
	router := mux.NewRouter()

	// Endpoint Health Check
	router.Handle("/health", health.NewHealthChecker(serverName).CheckHandlerCustom()).Methods(http.MethodGet)

	// Endpoint CORS EnvExpMath
	app.RunEndpointEnvExpMath(router, newConfig, newLogger)

	// Funcion de encedido de servidor
	ServerTurnOn(router, serverName, port, newLogger)
}
