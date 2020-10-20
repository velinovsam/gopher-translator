package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"gopher-translator/handlers/historyHandler"
	"gopher-translator/handlers/translationHandler"
)

func main() {

	port := flag.Int("port", 1234, "HTTP service port")
	flag.Parse()

	router := gin.Default()

	router.POST("/word", translationHandler.HandleTranslateWordRequest)
	router.POST("/sentence", translationHandler.HandleTranslateSentenceRequest)
	router.GET("/history", historyHandler.HandleHistoryRequest)

	_ = router.Run(fmt.Sprintf(":%d", *port))
}
