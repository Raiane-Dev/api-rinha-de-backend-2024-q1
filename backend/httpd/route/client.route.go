package route

import "rinha_api/backend/httpd/controller"

func (router *Router) clientPublic() {

	router.Clients.Post("/:id/transacoes", controller.SendTransaction)
	router.Clients.Get("/:id/extrato", controller.ConsultTransaction)
}
