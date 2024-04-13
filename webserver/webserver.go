package webserver

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/leonlatsch/fred/jsinject"
)

func SetupWebServer() {
	ginRouter := gin.New()
	ginRouter.Static("", ".")
	log.Println("fred now serving at http://localhost:8000")
	http.ListenAndServe(":8000", &jsinject.WsInjectorMiddleware{Next: ginRouter})
}
