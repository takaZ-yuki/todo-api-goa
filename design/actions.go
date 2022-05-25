package design

import (
	. "goa.design/goa/v3/dsl"
)

// Service describes a service
var _ = Service("user_controller", func() {
	Description("ユーザ関連のエンドポイント")
	HTTP(func() {
		Path("/users")
	})

	Method("GetUsers", func() {
		Description("ユーザ一覧の検索")
		Error("NotFound")
		Error("BadRequest")

		Result(ArrayOf(User))

		HTTP(func() {
			GET("/")
			Response(StatusOK)
			Response("NotFound", StatusNotFound)
			Response("BadRequest", StatusBadRequest)
		})
	})

	Method("GetUser", func() {
		Description("ユーザ検索")
		Error("NotFound")
		Error("BadRequest")

		Payload(func() {
			Attribute("id", Int, "ユーザーID")
			Required("id")
		})

		Result(User)

		HTTP(func() {
			GET("/{id}")
			Response(StatusOK)
			Response("NotFound", StatusNotFound)
			Response("BadRequest", StatusBadRequest)
		})
	})

	Method("UpdateUser", func() {
		Description("ユーザ更新")
		Payload(User)
		Result(Empty)
		Error("NotFound")
		Error("BadRequest")

		HTTP(func() {
			PUT("/{id}")
			Body(func() {
				Attribute("name")
				Attribute("email")
				//TODO: 必須設定
				// Required("name")
			})

			Response(StatusNoContent)
			Response("NotFound", StatusNotFound)
			Response("BadRequest", StatusBadRequest)
		})
	})

	Method("CreateUser", func() {
		Description("ユーザ登録")
		Payload(User)
		Result(Empty)
		Error("NotFound")
		Error("BadRequest")

		HTTP(func() {
			POST("/")
			Body(func() {
				Attribute("name")
				Attribute("email")
				//TODO: 必須設定
				// Required("name")
			})

			Response(StatusNoContent)
			Response("NotFound", StatusNotFound)
			Response("BadRequest", StatusBadRequest)
		})
	})

	Method("DeleteUser", func() {
		Description("ユーザ削除")
		Payload(func() {
			Attribute("id", Int, "ユーザーID")
			Required("id")
		})
		Result(Empty)
		Error("NotFound")
		Error("BadRequest")

		HTTP(func() {
			DELETE("/{id}")
			Response(StatusNoContent)
			Response("NotFound", StatusNotFound)
			Response("BadRequest", StatusBadRequest)
		})
	})
})