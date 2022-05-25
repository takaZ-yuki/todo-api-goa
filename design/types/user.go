package types

import . "goa.design/goa/v3/dsl"

// ユーザー情報
var User = Type("User", func() {
	Field(1, "ID", Int)
	Field(2, "mail", String)
	Field(3, "nickname", String)
})