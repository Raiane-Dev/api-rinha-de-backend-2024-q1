package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"rinha_api/backend/config"
	"rinha_api/backend/entity/io/input"
	"rinha_api/backend/entity/io/output"
	"rinha_api/backend/model"

	"github.com/go-playground/validator"
	"github.com/valyala/fasthttp"
)

var (
	validate = validator.New()
)

func SendTransaction(ctx *fasthttp.RequestCtx) {
	var (
		// mu               sync.Mutex
		// wg               sync.WaitGroup
		transactionInput input.TransactionInput
	)
	// wg.Add(1)
	if err := json.Unmarshal(ctx.PostBody(), &transactionInput); err != nil {
		log.Printf("invalid json body %+v | err %+v", transactionInput, err)
		ctx.SetStatusCode(http.StatusBadRequest)
		return
	}

	id := ctx.UserValue("id").(string)
	if err := validate.Struct(&transactionInput); err != nil {
		log.Println("validation failed", err)
		ctx.SetStatusCode(http.StatusBadRequest)
		return
	}

	client, err := model.FindByClient("id = ?", id)
	if err != nil {
		log.Println("not found client", err)
		config.ErrReaderDB <- err
		ctx.SetStatusCode(http.StatusNotFound)
		return
	}

	if transactionInput.Type == "d" {
		if (client.Balance + -transactionInput.Value) < -client.Limit {
			log.Printf("debit err | balance + value = %d | err %+v", client.Balance+transactionInput.Value, err)
			ctx.SetStatusCode(http.StatusUnprocessableEntity)
			return
		}
	}

	client.Balance = client.Balance + transactionInput.Value
	newTransaction := model.TransactionEntity{
		Type:        transactionInput.Type,
		Value:       transactionInput.Value,
		Description: transactionInput.Description,
	}

	// go func() {
	// mu.Lock()
	// defer mu.Unlock()
	// defer wg.Done()

	if err := newTransaction.Create(); err != nil {
		log.Println("not create transaction", err)
		config.ErrWriterDB <- err
		ctx.SetStatusCode(http.StatusBadGateway)
		return
	}

	if err := client.Update("id = ?", id); err != nil {
		log.Println("not update client", err)
		config.ErrWriterDB <- err
		ctx.SetStatusCode(http.StatusBadGateway)
		return
	}

	// }()

	// wg.Wait()
	resp_json, _ := json.Marshal(output.TransactionOutput{
		Limit:   client.Limit,
		Balance: client.Balance,
	})
	ctx.SetContentType("application/json")
	ctx.SetStatusCode(http.StatusOK)
	ctx.Response.SetBody(resp_json)
}
