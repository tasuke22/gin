package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Request はクライアントからのリクエストデータを格納するための構造体です。
type Request struct {
	ID    int    `uri:"id" json:"id" binding:"required"` // URIから取得するID (必須)
	Title string `form:"title" json:"title"`             // クエリパラメータまたはリクエストボディから取得するタイトル
	Score int    `form:"score" json:"score"`             // クエリパラメータまたはリクエストボディから取得するスコア
}

// sampleHandler はHTTPリクエストを処理するハンドラー関数です。
func sampleHandler(c *gin.Context) {
	var request Request

	// URIからIDをバインド。バインドに失敗した場合は400エラーを返します。
	if err := c.ShouldBindUri(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// クエリパラメータからタイトルとスコアをバインド。
	// ShouldBindUriの後にBindQueryを呼び出すことで、同一リクエスト内で複数のバインド方法を組み合わせる
	if err := c.BindQuery(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// リクエストボディからデータをバインド。
	// バインドに失敗した場合は400エラーを返します。
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// バインドされたデータをレスポンスとして返します。
	c.JSON(http.StatusOK, request)
}

func main() {
	r := gin.New()        // 新しいGinルーターのインスタンスを作成
	r.Use(gin.Recovery()) // パニックから回復するためのミドルウェアを追加

	// "/sample/:id"のパスにPOSTメソッドでリクエストがあった場合、sampleHandlerを実行
	r.POST("/sample/:id", sampleHandler)

	r.Run(":8080") // サーバーを8080ポートで起動
}
