package gin

import (
	"encoding/xml"
	"fmt"
	"html/template"
	"net/http"
	"net/url"
	"project/framework/gin/internal/json"
)

type IResponse interface {
	// IJson Json输出
	IJson(obj interface{}) IResponse

	// IJsonP Jsonp输出
	IJsonP(obj interface{}) IResponse

	// IXml xml输出
	IXml(obj interface{}) IResponse

	// IHtml html输出
	IHtml(template string, obj interface{}) IResponse

	// IText string
	IText(format string, values ...interface{}) IResponse

	// IRedirect 重定向
	IRedirect(path string) IResponse

	// ISetHeader header
	ISetHeader(key string, val string) IResponse

	// ISetCookie Cookie
	ISetCookie(key string, val string, maxAge int, path, domain string, secure, httpOnly bool) IResponse

	// ISetStatus 设置状态码
	ISetStatus(code int) IResponse

	// ISetOkStatus 设置200状态
	ISetOkStatus() IResponse
}

func (ctx *Context) IJson(obj interface{}) IResponse {
	byt, err := json.Marshal(obj)
	if err != nil {
		return ctx.ISetStatus(http.StatusInternalServerError)
	}
	ctx.ISetHeader("Content-Type", "application/json")
	ctx.Writer.Write(byt)
	return ctx
}

func (ctx *Context) IJsonP(obj interface{}) IResponse {
	callbackFunc := ctx.Query("callback")
	ctx.ISetHeader("Content-Type", "application/json")
	//输出到前端页面的时候需要注意进行字符过滤，否则有可能造成xss攻击
	callback := template.JSEscapeString(callbackFunc)

	//输出函数名
	_, err := ctx.Writer.Write([]byte(callback))
	if err != nil {
		return ctx
	}

	//输出左括号
	_, err = ctx.Writer.Write([]byte("("))
	if err != nil {
		return ctx
	}
	//数据函数参数
	ret, err := json.Marshal(obj)
	if err != nil {
		return ctx
	}

	_, err = ctx.Writer.Write(ret)
	if err != nil {
		return ctx
	}

	//输出左括号
	_, err = ctx.Writer.Write([]byte(")"))
	if err != nil {
		return ctx
	}

	return ctx

}

func (ctx *Context) IXml(obj interface{}) IResponse {
	byt, err := xml.Marshal(obj)
	if err != nil {
		return ctx.ISetStatus(http.StatusInternalServerError)
	}
	ctx.ISetHeader("Content-Type","application/xml")
	ctx.Writer.Write(byt)
	return ctx
}

func (ctx *Context) IHtml(file string, obj interface{}) IResponse {
	t, err := template.New("output").ParseFiles(file)
	if err != nil {
		return ctx
	}
	if err = t.Execute(ctx.Writer, obj);err != nil{
		return ctx
	}
	ctx.ISetHeader("Content-Type","application/html")
	return ctx
}

func (ctx *Context) IText(format string, values ...interface{}) IResponse {
	out := fmt.Sprintf(format, values...)
	ctx.ISetHeader("Content-Type","application/text")
	ctx.Writer.Write([]byte(out))
	return ctx
}

func (ctx *Context) IRedirect(path string) IResponse {
	http.Redirect(ctx.Writer,ctx.Request,path,http.StatusMovedPermanently)
	return ctx
}

func (ctx *Context) ISetHeader(key string, val string) IResponse {
	ctx.Writer.Header().Add(key,val)
	return ctx
}

func (ctx *Context) ISetCookie(key string, val string, maxAge int, path, domain string, secure, httpOnly bool) IResponse {
	if path == "" {
		path = "/"
	}
	http.SetCookie(ctx.Writer,&http.Cookie{
		Name: key,
		Value: url.QueryEscape(val),
		MaxAge: maxAge,
		Path: path,
		Domain: domain,
		SameSite: 1,
		Secure:secure,
		HttpOnly: httpOnly,
	})
	return ctx
}

func (ctx *Context) ISetStatus(code int) IResponse {
	ctx.Writer.WriteHeader(code)
	return ctx
}

func (ctx *Context) ISetOkStatus() IResponse {
	ctx.Writer.WriteHeader(http.StatusOK)
	return ctx
}
