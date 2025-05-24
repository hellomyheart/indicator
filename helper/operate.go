// Copyright (c) 2021-2024 Onur Cinar.
// The source code is provided under GNU AGPLv3 License.
// https://github.com/cinar/indicator

package helper

// Operate applies the provided operate function to corresponding values from two
// numeric input channels and sends the resulting values to an output channel.
//
// Example:
//
//	add := helper.Operate(ac, bc, func(a, b int) int {
//	  return a + b
//	})

// chan 操作函数，将两个chan中中的数据进行操作（o函数执行），
// 将o函数执行结果合在一起返回chan
func Operate[A any, B any, R any](ac <-chan A, bc <-chan B, o func(A, B) R) <-chan R {
	oc := make(chan R)

	go func() {
		// 自动关闭输出通道
		// 这个放在前面是因为：
		//		优势：
		//			覆盖所有退出路径：无论是正常完成、break退出还是panic，都能确保关闭
		//			及时释放资源：在goroutine结束时立即触发关闭
		//			符合最佳实践：Go官方推荐在goroutine入口处设置defer关闭
		defer close(oc)

		for {
			an, ok := <-ac
			if !ok {
				// chan异常，关闭输出通道
				Drain(bc)
				break
			}

			bn, ok := <-bc
			if !ok {
				// chan异常，关闭输出通道
				Drain(ac)
				break
			}
			// o结果放入 oc chan
			oc <- o(an, bn)
		}
	}()

	return oc
}
