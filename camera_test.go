package gopro

import (
  "fmt"
  "net/http"
  "net/http/httptest"
  "net/url"
  "testing"
)

const ipaddress string = "1.2.3.4"
const password string = "password"

func TestNewCamera(t *testing.T) {
  cam := NewCamera(ipaddress, password)

  if cam.ipaddress != ipaddress {
    t.Errorf("cam.ipaddress set to %s, expected %s", cam.ipaddress, ipaddress)
  }

  if cam.password != password {
    t.Errorf("cam.password set to %s, expected %s", cam.password, password)
  }
}

type recording struct {
  request *http.Request
}

func sniff(fn func(http.ResponseWriter, *http.Request)) (*recording, http.HandlerFunc) {
  var rec recording

  handler := func(writer http.ResponseWriter, request *http.Request) {
    rec.request = request
    fn(writer, request)
  }

  return &rec, http.HandlerFunc(handler)
}

func TestSend(t *testing.T) {
  rec, handler := sniff(func(writer http.ResponseWriter, req *http.Request) {
    fmt.Fprintln(writer, "request received")
  })

  srv := httptest.NewServer(handler)
  defer srv.Close()

  addr, _ := url.Parse(srv.URL)

  cam := NewCamera(addr.Host, password)
  err := cam.Send("MT", 1)

  if err != nil {
    t.Errorf("cam.Send returned an error")
  }

  request := rec.request

  if request.Method != "GET" {
    t.Errorf("cam.Send used method %s, expected %s", request.Method, "GET")
  }

  if request.URL.Path != "/camera/MT" {
    t.Errorf("cam.Send requested wrong path: %s", request.URL.Path)
  }

  query := request.URL.Query()

  if query.Get("t") != password {
    t.Errorf("cam.Send should send password")
  }

  if query.Get("p") != `%01` {
    t.Errorf("cam.Send should send parameter")
  }
}

func TestSendFail(t *testing.T) {
  // TODO
}
