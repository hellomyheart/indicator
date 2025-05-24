// Copyright (c) 2021-2024 Onur Cinar.
// The source code is provided under GNU AGPLv3 License.
// https://github.com/cinar/indicator

package helper

import "fmt"

// annotationReportColumn is the annotation report column struct.
// annotationReportColumn是注释报告列结构。
type annotationReportColumn struct {
	ReportColumn               // 接口组合 接口组合：通过匿名嵌入实现接口方法继承
	values       <-chan string // 数据源通道 通道驱动：使用只读通道实现实时数据流处理
}

// NewAnnotationReportColumn returns a new instance of an annotation column for a report.
// NewAnnotationReportColumn返回一个报告注释列的新实例。
// 工厂模式：封装实例创建逻辑
// 依赖注入：通过参数传递数据通道
func NewAnnotationReportColumn(values <-chan string) ReportColumn {
	return &annotationReportColumn{
		values: values,
	}
}

// Name returns the name of the report column.
//Name返回报表列的名称。
func (*annotationReportColumn) Name() string {
	return ""
}

// Type returns number as the data type.

// Type返回string作为数据类型。
func (*annotationReportColumn) Type() string {
	return "string"
}

// Role returns the role of the report column.
// Role返回报表列的角色。
func (*annotationReportColumn) Role() string {
	return "annotation"
}

// Value returns the next data value for the report column.
// Value返回报表列的下一个数据值。
// 空值处理：返回"null"兼容JSON格式
// 字符串安全：使用%q自动处理特殊字符
func (c *annotationReportColumn) Value() string {
	value := <-c.values

	if value != "" {
		return fmt.Sprintf("%q", value)
	}

	return "null"
}
