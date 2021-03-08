package main

import (
	"net/http"
	"os"

	"github.com/go-kit/kit/log"
	"github.com/lucabecci/auth-microservices-go/user"
)

func main() {
	var logger log.Logger
	logger = log.NewLogfmtLogger(os.Stderr)
	logger = log.With(logger, "listen", "8081", "caller", log.DefaultCaller)

	svc := user.NewLoggingMiddleware(logger, user.NewService())
	r := user.NewHTTPServer(svc, logger)
	logger.Log("msg", "HTTP", "addr", "8081")
	logger.Log("err", http.ListenAndServe(":8081", r))
}
