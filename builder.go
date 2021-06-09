package route_builder

import (
	"strings"
)

type Builder struct {
	// 内部builder
	builder strings.Builder
	// 根路径
	rootPath string
	// 路由组
	group string
	// 版本
	version string
	// path
	path string
	// param 参数
	param string
}

func New() Builder {
	return Builder{}
}

func WithRootPath(builder Builder, rootPath string) Builder {
	builder.rootPath = rootPath
	return builder
}

func WithGroup(builder Builder, group string) Builder {
	builder.group = group
	return builder
}

func WithVersion(builder Builder, version string) Builder {
	builder.version = version
	return builder
}

func WithPath(builder Builder, path string) Builder {
	builder.path = path
	return builder
}

func WithParam(builder Builder, param string) Builder {
	builder.param = param
	return builder
}

func (b Builder) String() string {
	b.builder.WriteString("/")
	defer b.builder.Reset()
	if len(b.rootPath) > 0 {
		// rootPath
		b.builder.WriteString(b.rootPath + "/")
	}
	if len(b.group) > 0 {
		// group
		b.builder.WriteString(b.group + "/")
	}
	if len(b.version) > 0 {
		// version
		b.builder.WriteString(b.version + "/")
	}

	if len(b.path) > 0 {
		// path
		b.builder.WriteString(b.path + "/")
	}
	if len(b.param) > 0 {
		b.builder.WriteString(b.param)
	}

	res := b.builder.String()

	return res
}
