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
  // TODO
}

func TestSendFail(t *testing.T) {
  // TODO
}

func TestSendParam(t *testing.T) {
  rec, handler := sniff(func(writer http.ResponseWriter, req *http.Request) {
    fmt.Fprintln(writer, "request received")
  })

  srv := httptest.NewServer(handler)
  defer srv.Close()

  addr, _ := url.Parse(srv.URL)

  cam := NewCamera(addr.Host, password)
  err := cam.SendParam("MT", 1)

  if err != nil {
    t.Errorf("cam.SendParam returned an error")
  }

  request := rec.request

  if request.Method != "GET" {
    t.Errorf("cam.SendParam used method %s, expected %s", request.Method, "GET")
  }

  if request.URL.Path != "/camera/MT" {
    t.Errorf("cam.SendParam requested wrong path: %s", request.URL.Path)
  }

  query := request.URL.Query()

  if query.Get("t") != password {
    t.Errorf("cam.SendParam should send password")
  }

  fmt.Println(request.URL)

  if query.Get("p") != string(1) {
    t.Errorf("cam.SendParam should send parameter")
  }
}

func TestSendParamFail(t *testing.T) {
  // TODO
}
