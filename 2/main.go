package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// MyBenchLoggerMiddleware は、リクエストごとのレイテンシ（応答時間）とステータスコードをログに記録するミドルウェアです。
func MyBenchLoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// リクエスト処理前に現在時刻を記録
		t := time.Now()

		// コンテキストにサンプル変数を設定（デモ用途）
		c.Set("example", "12345")

		// 次のハンドラまたはミドルウェアへ制御を渡す
		c.Next()

		// リクエスト処理後、経過時間（レイテンシ）を計算してログ出力
		latency := time.Since(t)
		log.Print(latency)

		// 最終的にレスポンスとして設定されたステータスコードをログ出力
		status := c.Writer.Status()
		log.Println(status)
	}
}

// AuthRequiredMiddleware は認証を要求するミドルウェアです。
// この例では認証処理はコメントアウトされており、常に認証が成功するようになっています。
func AuthRequiredMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 認証処理をここに実装
		// 以下は認証に失敗した場合の処理の例
		c.JSON(http.StatusUnauthorized, gin.H{"error": true, "from": "AuthRequiredMiddleware"})
		c.Abort() // 以降の処理を中断
	}
}

// benchmarkEndpoint はベンチマーク用のエンドポイントです。
func benchmarkEndpoint(c *gin.Context) {
	// MyBenchLoggerMiddlewareで設定されたサンプル変数を取得してログ出力
	example := c.MustGet("example").(string)
	log.Println(example)

	// レスポンスとしてJSONを返す
	c.JSON(http.StatusOK, gin.H{"error": false, "from": "benchmarkEndpoint"})
}

// meEndpoint は認証が必要なユーザー情報取得用のエンドポイントです。
func meEndpoint(c *gin.Context) {
	// レスポンスとしてJSONを返す
	c.JSON(http.StatusOK, gin.H{"error": false, "from": "meEndpoint"})
}

func main() {
	r := gin.New()
	r.Use(gin.Logger())   // ログ出力ミドルウェアを使用
	r.Use(gin.Recovery()) // パニック回復ミドルウェアを使用

	// ベンチマーク用エンドポイントのルーティング設定
	r.GET("/benchmark", MyBenchLoggerMiddleware(), benchmarkEndpoint)

	// 認証が必要なエンドポイントのグループ
	authorized := r.Group("/auth")
	authorized.Use(AuthRequiredMiddleware()) // 認証ミドルウェアを適用
	// ここの{}は必須ではない
	{
		authorized.GET("/me", meEndpoint) // ユーザー情報取得エンドポイント
	}

	// サーバーを8080ポートで起動
	r.Run(":8080")
}
