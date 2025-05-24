// Copyright (c) 2021-2024 Onur Cinar.
// The source code is provided under GNU AGPLv3 License.
// https://github.com/cinar/indicator

package helper

// Apply applies the given transformation function to each element in the
// input channel and returns a new channel containing the transformed
// values. The transformation function takes a float64 value as input
// and returns a float64 value as output.
//
// Example:
//
//	timesTwo := helper.Apply(c, func(n int) int {
//		return n * 2
//	})

// apply函数， 将chan的每一个元素都执行 f函数， 然后返回一个新的chan

// Apply将给定的转换函数应用到元素中的每个元素
//输入通道并返回一个包含转换后的新通道
//值。转换函数接受一个T值作为输入
//返回一个T值作为输出。
//
//例如:
//
// timetwo:= helper。应用(c, func(n int) int {
//返回n * 2
//})

func Apply[T Number](c <-chan T, f func(T) T) <-chan T {
	ac := make(chan T)

	go func() {
		defer close(ac)

		for n := range c {
			ac <- f(n)
		}
	}()

	return ac
}

// 三、设计模式分析
// 1. 通道流水线模式
// Input Chan --> Apply --> Output Chan
// 非阻塞式处理
// 支持背压控制
// 可组合性高
// 2. 泛型编程
// go
// // 支持各种数值类型
// Apply[int](...)      // 整数处理
// Apply[float64](...)  // 浮点处理
// 3. 函数式编程
// go
// // 高阶函数设计
// Apply(c, func(n T) T {
//     return n * 2 // 转换逻辑
// })

// 九、设计哲学
// 1. Go 式并发模型
// 每个 Apply 调用启动一个协程
// 通道作为主要通信机制
// 符合 Go 的 "不要通过共享内存来通信" 原则
// 2. 组合优于继承
// 通过通道链接构建复杂处理流程
// 支持链式调用
// 与 SliceToChan 等辅助函数无缝集成
// 3. 零成本抽象
// 泛型在编译时展开
// 不产生运行时开销
// 保持原生通道性能
// 十、总结
// 这个 Apply 函数是 Go 并发编程的典范实现：

// 核心价值
// 通用性：支持所有数值类型
// 灵活性：可组合任意转换函数
// 高效性：基于通道的流式处理
// 技术亮点
// 泛型编程与并发的完美结合
// 函数式编程风格
// 符合 Go 的并发设计哲学
// 使用建议
// 优先使用缓冲通道提升性能
// 注意转换函数的 panic 处理
// 控制协程数量防止资源耗尽
// 这个函数是构建复杂数据处理流水线的基础组件，特别适合需要实时处理数据流的场景。
