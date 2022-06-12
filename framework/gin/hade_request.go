package gin

import (
	"github.com/spf13/cast"
	"mime/multipart"
)

type IRequest interface {
	//请求地址url中带的参数

	DefaultQueryInt(key string, def int) (int, bool)
	DefaultQueryInt64(key string, def int64) (int64, bool)
	DefaultQueryFloat64(key string, def float64) (float64, bool)
	DefaultQueryFloat32(key string, def float32) (float32, bool)
	DefaultQueryBool(key string, def bool) (bool, bool)
	DefaultQueryString(key string, def string) (string, bool)
	DefaultQueryStringSlice(key string, def []string) ([]string, bool)

	//路由匹配中带的参数
	//形如 /book/:id

	DefaultParamInt(key string, def int) (int, bool)
	DefaultParamInt64(key string, def int64) (int64, bool)
	DefaultParamFloat64(key string, def float64) (float64, bool)
	DefaultParamFloat32(key string, def float32) (float32, bool)
	DefaultParmaBool(key string, def bool) (bool, bool)
	DefaultParamString(key string, def string) (string, bool)
	DefaultParam(key string) interface{}

	// form表单中带的参数
	DefaultFormInt(key string, def int) (int, bool)
	DefaultFormInt64(key string, def int64) (int64, bool)
	DefaultFormFloat64(key string, def float64) (float64, bool)
	DefaultFormFloat32(key string, def float32) (float32, bool)
	DefaultFormBool(key string, def bool) (bool, bool)
	DefaultString(key string, def string) (string, bool)
	DefaultStringSlice(key string, def []string) ([]string, bool)
	DefaultFile(key string) (*multipart.File, error)
	DefaultForm(key string) interface{}

	// BindJson json body
	BindJson(obj interface{}) error

	// BindXml xml body
	BindXml(obj interface{}) error

	// GetRawData 其他格式
	GetRawData() ([]byte, error)

	//基础信息
	Uri() string
	Method() string
	Host() string
	ClientIP() string

	//header
	Headers() map[string]string
	Header(key string) (string, error)

	//cookie
	Cookies() map[string]string
	Cookie(key string) (string, bool)
}

func (ctx *Context) QueryAll() map[string][]string {
	ctx.initQueryCache()
	return ctx.queryCache

}

func (ctx *Context) DefaultQueryInt(key string, def int) (int, bool) {
	params := ctx.QueryAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return cast.ToInt(vals[0]), true
		}
	}
	return def, false
}

func (ctx *Context) DefaultQueryInt64(key string, def int64) (int64, bool) {
	params := ctx.QueryAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return cast.ToInt64(vals[0]), true
		}
	}
	return def, false
}

func (ctx *Context) DefaultQueryFloat64(key string, def float64) (float64, bool) {
	params := ctx.QueryAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return cast.ToFloat64(key[0]), true
		}
	}
	return def, false
}

func (ctx *Context) DefaultQueryFloat32(key string, def float32) (float32, bool) {
	params := ctx.QueryAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return cast.ToFloat32(vals[0]), true
		}
	}
	return def, false
}

func (ctx *Context) DefaultQueryBool(key string, def bool) (bool, bool) {
	params := ctx.QueryAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return cast.ToBool(vals[0]), true
		}
	}
	return def, false
}

func (ctx *Context) DefaultQueryString(key string, def string) (string, bool) {
	params := ctx.QueryAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return vals[0], true
		}
	}
	return def, false
}

func (ctx *Context) DefaultQueryStringSlice(key string, def []string) ([]string, bool) {
	params := ctx.QueryAll()
	if vals, ok := params[key]; ok {
		return vals, true
	}
	return def, false
}

//路由匹配中带的参数

func (ctx *Context) HadeParam(key string) interface{} {
	if val, ok := ctx.Params.Get(key); ok {
		return val
	}
	return nil
}

func (ctx *Context) DefaultParamInt(key string, def int) (int, bool) {
	if val := ctx.HadeParam(key); val != nil {
		return cast.ToInt(val), true
	}
	return def, false
}

func (ctx *Context) DefaultParamInt64(key string, def int64) (int64, bool) {
	if val := ctx.HadeParam(key); val != nil {
		return cast.ToInt64(val), true
	}
	return def, false
}

func (ctx *Context) DefaultParamFloat64(key string, def float64) (float64, bool) {
	if val := ctx.HadeParam(key); val != nil {
		return cast.ToFloat64(val), true
	}
	return def, false
}

func (ctx *Context) DefaultParamFloat32(key string, def float32) (float32, bool) {
	if val := ctx.HadeParam(key); val != nil {
		return cast.ToFloat32(val), true
	}
	return def, false
}

func (ctx *Context) DefaultParmaBool(key string, def bool) (bool, bool) {
	if val := ctx.HadeParam(key); val != nil {
		return cast.ToBool(val), true
	}
	return def, false
}

func (ctx *Context) DefaultParamString(key string, def string) (string, bool) {
	if val := ctx.HadeParam(key); val != nil {
		return cast.ToString(val), true
	}
	return def, false
}

func (ctx *Context) FormAll() map[string][]string {
	ctx.initFormCache()
	return ctx.formCache

}

func (ctx *Context) DefaultFormInt(key string, def int) (int, bool) {
	params := ctx.FormAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return cast.ToInt(vals[0]), true
		}
	}
	return def, false
}

func (ctx *Context) DefaultFormInt64(key string, def int64) (int64, bool) {
	params := ctx.FormAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return cast.ToInt64(vals[0]), true
		}
	}
	return def, false
}

func (ctx *Context) DefaultFormFloat64(key string, def float64) (float64, bool) {
	params := ctx.FormAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return cast.ToFloat64(vals[0]), true
		}
	}
	return def, false
}

func (ctx *Context) DefaultFormFloat32(key string, def float32) (float32, bool) {
	parmas := ctx.FormAll()
	if vals, ok := parmas[key]; ok {
		if len(vals) > 0 {
			return cast.ToFloat32(vals[0]), true
		}
	}
	return def, false
}

func (ctx *Context) DefaultFormBool(key string, def bool) (bool, bool) {
	params := ctx.FormAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return cast.ToBool(vals), true
		}
	}
	return def, false
}

func (ctx *Context) DefaultString(key string, def string) (string, bool) {
	params := ctx.FormAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return cast.ToString(vals[0]), true
		}
	}
	return def, false
}

func (ctx *Context) DefaultStringSlice(key string, def []string) ([]string, bool) {
	params := ctx.FormAll()
	if vals, ok := params[key]; ok {
		return vals, true
	}
	return def, false
}

func (ctx *Context) DefaultForm(key string) interface{} {
	params := ctx.FormAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return vals[0]
		}
	}
	return nil
}
