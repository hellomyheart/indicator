// Package helper contains the helper functions.
//
// This package belongs to the Indicator project. Indicator is
// a Golang module that supplies a variety of technical
// indicators, strategies, and a backtesting framework
// for analysis.
//
// # License
//
//	Copyright (c) 2021-2024 Onur Cinar.
//	The source code is provided under GNU AGPLv3 License.
//	https://github.com/cinar/indicator
//
// # Disclaimer
//
// The information provided on this project is strictly for
// informational purposes and is not to be construed as
// advice or solicitation to buy or sell any security.
// 帮助类包
package helper

// Integer refers to any integer type.
// 定义了 Integer 接口，表示可以接受任意一种类型
type Integer interface {
	int | int8 | int16 | int32 | int64
}

// Float refers to any float type.
// 定义了Float 接口，表示可以接受任意一种类型
type Float interface {
	float32 | float64
}

// Number refers to any numeric type.
// 定义了Number 接口，表示可以接受任意一种类型
type Number interface {
	Integer | Float
}
