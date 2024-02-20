package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"rinha_api/backend/config"
	"rinha_api/backend/entity/io/input"
	"rinha_api/backend/entity/io/output"
	"rinha_api/backend/model"
	"rinha_api/backend/util/logger"
	"strconv"
	"sync"

	"github.com/go-playground/validator"
	"github.com/valyala/fasthttp"
)

var (
	validate = validator.New()
)

func SendTransaction(ctx *fasthttp.RequestCtx) {
	var (
		mu               sync.Mutex
		wg               sync.WaitGroup
		transactionInput input.TransactionInput

		client model.ClientEntity
	)
	wg.Add(1)
	if err := json.Unmarshal(ctx.PostBody(), &transactionInput); err != nil {
		logger.Error(fmt.Sprintf("invalid json body %+v", transactionInput), err)
		ctx.SetStatusCode(http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(ctx.UserValue("id").(string))
	if err != nil {
		logger.Error("id not number. id: "+string(ctx.UserValue("id").(string)), err)
		ctx.SetStatusCode(http.StatusBadRequest)
		return
	}

	if err := validate.Struct(&transactionInput); err != nil {
		logger.Error("validation failed", err)
		ctx.SetStatusCode(http.StatusBadRequest)
		return
	}

	client, err = client.FindBy("id = ?", id)
	if err != nil {
		logger.Error("not found client", err)
		ctx.SetStatusCode(http.StatusNotFound)
		return
	}

	if transactionInput.Type == "d" {
		if (client.Balance + -transactionInput.Value) < -client.Limit {
			logger.Error(fmt.Sprintf("debit err | balance + value = %d", client.Balance+transactionInput.Value), err)
			ctx.SetStatusCode(http.StatusUnprocessableEntity)
			return
		}
	}

	client.Balance = client.Balance + transactionInput.Value
	newTransaction := model.TransactionEntity{
		ClientID:    id,
		Type:        transactionInput.Type,
		Value:       transactionInput.Value,
		Description: transactionInput.Description,
	}

	go func() {
		mu.Lock()
		defer mu.Unlock()
		defer wg.Done()

		if err := newTransaction.Create(); err != nil {
			config.DatabaseErr <- err
			logger.Error("not create transaction", err)
			ctx.SetStatusCode(http.StatusBadGateway)
			return
		}

		if err := client.Update("id = ?", id); err != nil {
			config.DatabaseErr <- err
			logger.Error("not update client", err)
			ctx.SetStatusCode(http.StatusBadGateway)
			return
		}

	}()
	resp := output.TransactionOutput{
		Limit:   client.Limit,
		Balance: client.Balance,
	}

	wg.Wait()
	resp_json, _ := json.Marshal(resp)
	ctx.SetContentType("application/json")
	ctx.SetStatusCode(http.StatusOK)
	ctx.Response.SetBody(resp_json)
}
