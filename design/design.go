package design

import (
	. "goa.design/goa/v3/dsl"
)

// API describes the global properties of the API server.
var _ = API("TODO API", func() {
    Title("TODOアプリ用のAPI")
    Description("")
		Version("1.0")

		//TODO: コントローラーをcotrollerフォルダに移動したい

    // Server はクライアントのリクエストを受け付ける単一のプロセスを記述します
    // DSLは、サーバーがホストする一連のサービスとホストの詳細を定義します
		Server("controller", func() {
      Description("")

			// このサーバーによってホストされているサービスを列挙します
			Services("user_controller")
			// Services("todo_controller")

			// ホストとそのトランスポート URL を列挙します。
			Host("localhost", func() {
				// トランスポート固有の URL, サポートされているスキーマは以下です：
				// 'http', 'https', 'grpc' and 'grpcs' とそれぞれのデフォルト
				// ポート： 80, 443, 8080, 8443
				URI("http://0.0.0.0:8088")
				// Variable は URI の変数を記述します
				Variable("version", String, "API version", func() {
					// URL パラメータにはデフォルト値か Enum バリデーション（あるいはその両方）が必要です
					Default("v1")
				})
			})

			// ホストとそのトランスポート URL を列挙します。
			// Host("development", func() {
			// 	Description("Development hosts.")
			// 	// トランスポート固有の URL, サポートされているスキーマは以下です：
			// 	// 'http', 'https', 'grpc' and 'grpcs' とそれぞれのデフォルト
			// 	// ポート： 80, 443, 8080, 8443
			// 	URI("http://0.0.0.0:8088")

			// 	// Variable は URI の変数を記述します
			// 	Variable("version", String, "API version", func() {
			// 		// URL パラメータにはデフォルト値か Enum バリデーション（あるいはその両方）が必要です
			// 		Default("v1")
			// 	})
			// })
    })
})