package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// Ginのデフォルトのルーターを生成
	r := gin.Default()

	// ルートURLに対してGETリクエストをハンドル
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello World!")
	})

	// サーバーをポート8001で開始
	r.Run(":8001")
}
