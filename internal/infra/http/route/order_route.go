package route

import (
	"github.com/go-chi/chi/v5"
	"github.com/reangeline/go-clean-arch/internal/presentation/controller"
)

func InitializeOrderRoutes(controller *controller.OrderController, r chi.Router) {

	r.Route("/order", func(r chi.Router) {
		r.Get("/", controller.GetAllOrder)
		r.Post("/", controller.CreateOrder)
	})

}
