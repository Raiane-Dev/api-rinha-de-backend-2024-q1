package controller

import (
	"encoding/json"
	"net/http"
	"rinha_api/backend/entity/io/output"
	"rinha_api/backend/model"
	"rinha_api/backend/util/logger"
	"strconv"
	"time"

	"github.com/valyala/fasthttp"
)

var (
	transaction model.TransactionEntity
)

func ConsultTransaction(ctx *fasthttp.RequestCtx) {
	var (
		client model.ClientEntity
	)
	id, err := strconv.Atoi(ctx.UserValue("id").(string))
	if err != nil {
		logger.Error("id not number", err)
		ctx.SetStatusCode(http.StatusBadRequest)
		return
	}

	client, err = client.FindBy("id = ?", id)
	if err != nil {
		logger.Error("not found client", err)
		ctx.SetStatusCode(http.StatusNotFound)
		return
	}

	transactions, err := transaction.FindMany("cliente_id = ?", id)
	if err != nil {
		logger.Error("not found transaction", err)
		ctx.SetStatusCode(http.StatusBadRequest)
		return
	}

	out := output.ExtractOutput[model.TransactionEntity]{
		Balance: output.Balance{
			Total:       client.Balance,
			ExtractDate: time.Now().Format(time.RFC3339),
			Limit:       client.Limit,
		},
		LatestTransaction: transactions,
	}

	resp_json, _ := json.Marshal(out)
	ctx.SetContentType("application/json")
	ctx.SetStatusCode(http.StatusOK)
	ctx.Response.SetBody(resp_json)
}
