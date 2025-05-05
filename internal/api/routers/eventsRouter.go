package routers

import (
	"TunaAPIGateway/internal/api/handlers"
	"github.com/julienschmidt/httprouter"
)

func EventsRouter() *httprouter.Router {
	router := httprouter.New()
	router.POST("/:event_type", handlers.PostEvent)
	return router
}
