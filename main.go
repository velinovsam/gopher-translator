package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"gopher-translator/handlers/history_handler"
	"gopher-translator/handlers/translation_handler"
)

func main() {

	port := flag.Int("port", 1234, "HTTP service port")
	flag.Parse()

	router := gin.Default()

	router.POST("/word", translation_handler.HandleTranslateWordRequest)
	router.POST("/sentence", translation_handler.HandleTranslateSentenceRequest)
	router.GET("/history", history_handler.HandleHistoryRequest)

	_ = router.Run(fmt.Sprintf(":%d", *port))
}
