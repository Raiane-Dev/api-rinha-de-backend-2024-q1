package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"rinha_api/backend/entity/io/output"
	"rinha_api/backend/model"
	"time"

	"github.com/valyala/fasthttp"
)

func ConsultTransaction(ctx *fasthttp.RequestCtx) {

	id := ctx.UserValue("id").(string)
	client, err := model.FindByClient("id = ?", id)
	if err != nil {
		// config.ErrReaderDB <- err
		log.Println("not found client", err)
		ctx.SetStatusCode(http.StatusNotFound)
		return
	}

	resp_json, _ := json.Marshal(output.ExtractOutput[map[string]string]{
		Balance: output.Balance{
			Total:       client.Balance,
			ExtractDate: time.Now().Format(time.RFC3339),
			Limit:       client.Limit,
		},
		LatestTransaction: model.FindManyTransactions("cliente_id = ?", id),
	})
	ctx.SetContentType("application/json")
	ctx.SetStatusCode(http.StatusOK)
	ctx.Response.SetBody(resp_json)

}
