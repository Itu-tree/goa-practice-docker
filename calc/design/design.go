package design

import (
	. "goa.design/goa/v3/dsl"
)

// API 定義
var _ = API("calc", func() {
	// API の説明（タイトルと説明）
	Title("Calculator Service")
	Description("Service for adding numbers, a Goa teaser")

	// サーバ定義
	Server("calc", func() {
		Host("localhost", func() {
			URI("http://localhost:8000") // HTTP REST API
			URI("grpc://localhost:8080") // gRPC
		})
	})
})

// サービス定義
var _ = Service("calc", func() {
	// 説明
	Description("The calc service performs operations on numbers.")

	// メソッド (HTTPでいうところのエンドポントに相当)
	Method("add", func() {
		// ペイロード定義
		Payload(func() {
			// 整数型の属性 `a` これは左の被演算子
			Attribute("a", Int, func() {
				Description("Left operand") // 説明
				Meta("rpc:tag", "1")        // gRPC 用のメタ情報。タグ定義
			})
			// 整数型の属性 `b` これは右の被演算子
			Attribute("b", Int, func() {
				Description("Right operand") // 説明
				Meta("rpc:tag", "2")         // gRPC 用のメタ情報。タグ定義
			})
			Required("a", "b") // a と b は required な属性であることの指定
		})

		Result(Int) // メソッドの返値（整数を返す）

		// HTTP トランスポート用の定義
		HTTP(func() {
			GET("/add/{a}/{b}") // GET エンドポイント
			Response(StatusOK)  // レスポンスのステータスは Status OK = 200 を返す
		})

		// GRPC トランスポート用の定義
		GRPC(func() {
			Response(CodeOK) // レスポンスのステータスは CodeOK を返す
		})
	})

	Method("sub", func() {
		// input
		Payload(func() {
			Field(1, "a", Int, "Left operand")
			Field(2, "b", Int, "Right operand")
			Required("a", "b")
		})

		// output
		Result(Int)

		// HTTP トランスポート用の定義
		HTTP(func() {
			GET("/sub/{a}/{b}") // GET エンドポイント
			Response(StatusOK)  // レスポンスのステータスは Status OK = 200 を返す
		})

		// GRPC トランスポート用の定義
		GRPC(func() {
			Response(CodeOK) // レスポンスのステータスは CodeOK を返す
		})
	})

	Method("mul", func() {
		// input
		Payload(func() {
			Field(1, "a", Int, "Left operand")
			Field(2, "b", Int, "Right operand")
			Required("a", "b")
		})
		// output
		Result(Int)

		// HTTP トランスポート用の定義
		HTTP(func() {
			GET("/mul/{a}/{b}") // GET エンドポイント
			Response(StatusOK)  // レスポンスのステータスは Status OK = 200 を返す
		})

		// GRPC トランスポート用の定義
		GRPC(func() {
			Response(CodeOK) // レスポンスのステータスは CodeOK を返す
		})

	})

	//
	Method("divide", func() {
		Description("Divide returns the integral division of two integers.")
		// input
		Payload(func() {
			Field(1, "a", Int, "Left operand")
			Field(2, "b", Int, "Right operand")
			Required("a", "b")
		})
		// output
		Result(Int)

		// error
		Error("DivByZero")

		// HTTP トランスポートへの割り当て
		HTTP(func() {
			GET("/div/{a}/{b}")                     // 入力
			Response(StatusOK)                      // 正常系
			Response("DivByZero", StatusBadRequest) // エラー
		})

		GRPC(func() {
			Response(CodeOK)
			Response("DivByZero", CodeInvalidArgument)
		})
	})
})
