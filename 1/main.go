package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 仮のDBとして
var db = make(map[string]string)

// setupRouter関数は、Ginルーターの設定を行い、それを返します。
func setupRouter() *gin.Engine {
	// Ginのデフォルトルーターを初期化します。
	r := gin.Default()

	// GETリクエストを"/ping"パスにマップします。リクエストがあったら"pong"と応答します。
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// "/user/:name"へのGETリクエストを処理します。:nameはパラメータとして機能し、
	// URLの一部としてユーザー名を受け取ることができます。
	r.GET("/user/:name", func(c *gin.Context) {
		user := c.Params.ByName("name")
		value, ok := db[user]
		if ok {
			// ユーザーがdbマップに存在する場合、その値をJSONレスポンスとして返します。
			c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
		} else {
			// ユーザーがdbマップに存在しない場合、"no value"ステータスを返します。
			c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
		}
	})

	// Basic認証を必要とするルートグループを作成します。
	// このグループ内のルートにアクセスするには、認証が必要です。
	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"foo":  "bar", // ユーザー名"foo"に対するパスワードは"bar"です。
		"manu": "123", // ユーザー名"manu"に対するパスワードは"123"です。
	}))

	// "/admin"へのPOSTリクエストを処理します。このルートはBasic認証を必要とします。
	authorized.POST("admin", func(c *gin.Context) {
		user := c.MustGet(gin.AuthUserKey).(string) // 認証済みユーザー名を取得します。

		var json struct {
			Value string `json:"value" binding:"required"` // リクエストから"value"フィールドをバインドします。
		}

		if c.Bind(&json) == nil {
			db[user] = json.Value                        // バインドされた値をdbマップに保存します。
			c.JSON(http.StatusOK, gin.H{"status": "ok"}) // 状態を"ok"として応答します。
		}
	})

	return r // 設定されたルーターを返します。
}

func main() {
	r := setupRouter() // ルーターを設定します。
	r.Run(":8080")     // サーバーを8080ポートで起動します。
}
