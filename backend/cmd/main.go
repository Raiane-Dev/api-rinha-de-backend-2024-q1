package main

import (
	"fmt"
	"log"
	"os"
	"rinha_api/backend/config"
	"rinha_api/backend/httpd/controller"

	_ "net/http/pprof"
	"runtime/pprof"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

func init() {
	var (
		err error
	)

	config.WriterDB, err = config.ConnectInstance()
	if err != nil {
		log.Fatalf("err connect db %s", err.Error())
	}

	config.ReaderDB, err = config.ConnectInstance()
	if err != nil {
		log.Fatalf("err connect db %s", err.Error())
	}

	config.ExecMigration()
}

func main() {
	defer func() {
		f, _ := os.Create("profile.pprof")
		defer f.Close()
		pprof.WriteHeapProfile(f)

		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()

	go func() {
		select {
		case err := <-config.ErrWriterDB:

			log.Println("connection refused for instance WriterDB", err)
			config.WriterDB.Exec("VACUUM;")
			config.WriterDB.Close()
			config.WriterDB, err = config.ConnectInstance()
			log.Printf("retry connection, err: %s", err)

		case err := <-config.ErrReaderDB:
			log.Println("connection refused for instance ReaderDB", err)
			config.ReaderDB.Exec("VACUUM;")
			config.ReaderDB.Close()
			config.ReaderDB, err = config.ConnectInstance()
			log.Printf("retry connection, err: %s", err)

		}

	}()

	// if os.Getenv("DEBUG") == "debug" {
	// go func() {
	// log.Println(http.ListenAndServe("localhost:6060", nil))
	// }()
	// }

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
