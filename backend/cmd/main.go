package main

import (
	"log"
	"os"
	"rinha_api/backend/config"
	"rinha_api/backend/httpd/controller"
	"rinha_api/backend/util/logger"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

func init() {
	logger.Init()

	if err := config.ConnectInstance(); err != nil {
		log.Fatalf("err connect db %s", err.Error())
	}

	config.ExecMigration()
}

func main() {

	go func() {
		select {
		case err := <-config.DatabaseErr:
			logger.Error("connection refused", err)
			config.DatabaseInstance.Exec("VACUUM;")
			config.DatabaseInstance.Close()
			config.ConnectInstance()
			logger.Info("tratament go")
		}
	}()

	r := router.New()
	r.POST("/clientes/{id}/transacoes", controller.SendTransaction)
	r.GET("/clientes/{id}/extrato", controller.ConsultTransaction)

	s := &fasthttp.Server{
		Handler:          r.Handler,
		Concurrency:      fasthttp.DefaultConcurrency,
		DisableKeepalive: true,
	}

	port := os.Getenv("SERVER_PORT")
	if err := s.ListenAndServe(":" + port); err != nil {
		log.Fatalf("Error listening on port %s: %s", port, err)
	}

}
