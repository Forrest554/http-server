package pkg

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// 上下文 选择是结合体而不是接口的原因：交给框架
// context的进一步抽象，不希望用户创建context，让web框架去创建

type Context struct {
	W http.ResponseWriter
	R *http.Request
}

func NewContext(w http.ResponseWriter, r *http.Request) *Context {
	ctx := &Context{W: w, R: r}
	return ctx
}

// 对req的json处理，并对err进行封装
// 使用obj interface{}作为参数，能接受任意类型参数
func (c *Context) ReadJson(req interface{}) error {
	r := c.R
	w := c.W
	body, err := io.ReadAll(r.Body)
	if err != nil {
		// 读失败了
		fmt.Fprintf(w, "read body failed: %v", err)
		return err
	}
	err = json.Unmarshal(body, req)
	if err != nil {
		// 失败了
		fmt.Fprintf(w, "deserialized failed: %v", err)
		return err
	}
	return nil
}

func (c *Context) WriteJson(code int, res interface{}) error {
	c.W.WriteHeader(code)
	if res != nil {
		// 正确响应
		respJson, err := json.Marshal(res)
		if err != nil {
			return err
		}
		_, err = c.W.Write(respJson)
		return err
	}
	return nil
}

func (c *Context) OKJson(res interface{}) error {
	return c.WriteJson(http.StatusOK, res)
}

func (c *Context) SystemErrorJson(res interface{}) error {
	return c.WriteJson(http.StatusInternalServerError, res)
}

func (c *Context) BadRequestJson(res interface{}) error {
	return c.WriteJson(http.StatusBadRequest, res)
}
