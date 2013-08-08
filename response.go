package alpha

import (
  "mime"
  "strings"
  "net/http"
)

type Response struct {
  Out       http.ResponseWriter
  Headers   http.Header
}

func (res *Response) Status(code int) *Response {
  res.Out.WriteHeader(code)
  return res
}

func (res *Response) SendString(body string) *Response {
  res.Out.Write([]byte(body))
  return res
}

func (res *Response) Type(t string) *Response {
  if t == "" {
    return res
  }

  field := "Content-Type"
  t = strings.ToLower(t)

  if ^strings.Index(t, "/") == 0 {
    if '.' != t[0] {
      t = "." + t
    }
    t = mime.TypeByExtension(t)
  }

  return res.SetHeader(field, t)
}

func (res *Response) ContentType(t string) *Response {
  return res.Type(t)
}

func (res *Response) Send() *Response {
  return res
}

func (res *Response) SetHeader(field string, val string) *Response {
  res.Headers.Set(field, val)
  return res
}

func (res *Response) Set() *Response {
  return res
}
