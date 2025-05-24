// Copyright (c) 2021-2024 Onur Cinar.
// The source code is provided under GNU AGPLv3 License.
// https://github.com/cinar/indicator

package helper

import "fmt"

// numericReportColumn is the number report column struct.
type numericReportColumn[T Number] struct {
	ReportColumn          // 接口组合 接口组合：通过匿名嵌入实现接口方法继承
	name         string   //列名称
	values       <-chan T //数值流通道
}

// NewNumericReportColumn returns a new instance of a numeric data column for a report.
// NewNumericReportColumn返回报告的数字数据列的新实例。
func NewNumericReportColumn[T Number](name string, values <-chan T) ReportColumn {
	return &numericReportColumn[T]{
		name:   name,
		values: values,
	}
}

// Name returns the name of the report column.
// 返回列名
func (c *numericReportColumn[T]) Name() string {
	return c.name
}

// Type returns number as the data type.
// 返回列类型
func (*numericReportColumn[T]) Type() string {
	return "number"
}

// Role returns the role of the report column.
// 返回列角色
func (*numericReportColumn[T]) Role() string {
	return "data"
}

// Value returns the next data value for the report column.
// 获取列的下一个数据值
func (c *numericReportColumn[T]) Value() string {
	return fmt.Sprintf("%v", <-c.values)
}
