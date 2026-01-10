package integrationtest

import (
	"bytes"
	"io"
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v2"
)

type FiberTest struct {
	app *fiber.App
	t   *testing.T
}

func (f *FiberTest) Handlers(m ...fiber.Handler) *FiberTest {
	for _, h := range m {
		f.app.Use(h)
	}
	return f
}

func (f *FiberTest) TestJSON(j string) (*http.Response, error) {
	return f.TestRequest(NewHttpRequestJSON(f.t, j))
}

func (f *FiberTest) TestRequest(r *http.Request) (*http.Response, error) {
	return f.app.Test(r, -1)
}

func NewHttpRequestJSON(t *testing.T, b string) *http.Request {
	r := NewHttpRequestBody(t, b)
	r.Header.Set("content-type", "application/json")
	return r
}

func NewHttpRequestBody(t *testing.T, b string) *http.Request {
	if b == "" {
		t.Fatal("cannot create http request, the body is empty")
	}
	return NewHttpRequest(t, "http://test.com", []byte(b))
}

func NewHttpRequest(t *testing.T, u string, b []byte) *http.Request {
	var bf io.Reader
	if len(b) > 0 {
		bf = bytes.NewReader(b)
	}
	r, err := http.NewRequest(http.MethodPost, u, bf)

	if err != nil {
		t.Fatal(err.Error())
	}

	return r
}

func NewFiberTest(t *testing.T, m ...fiber.Handler) *FiberTest {
	a := fiber.New()

	return (&FiberTest{
		app: a,
		t:   t,
	}).Handlers(m...)
}
