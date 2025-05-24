// Copyright (c) 2021-2024 Onur Cinar.
// The source code is provided under GNU AGPLv3 License.
// https://github.com/cinar/indicator

package helper

import "math"

// Abs calculates the absolute value of each value in a channel of float64.
//
// Example:
//
//	abs := helper.Abs(helper.SliceToChan([]int{-10, 20, -4, -5}))
//	fmt.Println(helper.ChanToSlice(abs)) // [10, 20, 4, 5]

// 对通道中的每个数值计算绝对值：
// 将数据转为 float64 类型，然后计算绝对值，再转为原始类型。
func Abs[T Number](c <-chan T) <-chan T {
	return Apply(c, func(n T) T {
		return T(math.Abs(float64(n)))
	})
}
