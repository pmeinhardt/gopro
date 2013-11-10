package gopro

import (
  "net/http"
  "net/http/httptest"
  "net/url"
  "testing"
)

func camcmp(t *testing.T, cam *Camera, ipaddress, password string) {
  if cam.ipaddress != ipaddress {
    t.Errorf("cam.ipaddress set to %s, expected %s", cam.ipaddress, ipaddress)
  }

  if cam.password != password {
    t.Errorf("cam.password set to %s, expected %s", cam.password, password)
  }
}

func TestNewCamera(t *testing.T) {
  cam := NewCamera("1.2.3.4", "password")
  camcmp(t, cam, "1.2.3.4", "password")
}

func TestDefaultCamera(t *testing.T) {
  cam := DefaultCamera("password")
  camcmp(t, cam, DefaultIP, "password")
}

func capture(fn func(cam *Camera) error) (*http.Request, error) {
  var req *http.Request

  handler := func(writer http.ResponseWriter, request *http.Request) {
    req = request
  }

  srv := httptest.NewServer(http.HandlerFunc(handler))
  defer srv.Close()

  addr, _ := url.Parse(srv.URL)

  cam := NewCamera(addr.Host, "testpass")
  err := fn(cam)

  return req, err
}

func TestSend(t *testing.T) {
  req, err := capture(func (cam *Camera) error {
    return cam.Send("DL")
  })

  if err != nil {
    t.Errorf("cam.Send returned an error")
  }

  if req.Method != "GET" {
    t.Errorf("cam.Send used method %s, expected %s", req.Method, "GET")
  }

  if req.URL.Path != "/camera/DL" {
    t.Errorf("cam.Send requested wrong path: %s", req.URL.Path)
  }

  query := req.URL.Query()

  if query.Get("t") != "testpass" {
    t.Errorf("cam.Send should send password")
  }
}

func TestSendFail(t *testing.T) {
  // TODO
}

func TestSendParam(t *testing.T) {
  req, err := capture(func (cam *Camera) error {
    return cam.SendParam("SH", 1)
  })

  if err != nil {
    t.Errorf("cam.SendParam returned an error")
  }

  if req.Method != "GET" {
    t.Errorf("cam.SendParam used method %s, expected %s", req.Method, "GET")
  }

  if req.URL.Path != "/camera/SH" {
    t.Errorf("cam.SendParam requested wrong path: %s", req.URL.Path)
  }

  query := req.URL.Query()

  if query.Get("t") != "testpass" {
    t.Errorf("cam.SendParam should send password")
  }

  if query.Get("p") != string(1) {
    t.Errorf("cam.SendParam should send parameter")
  }
}

func TestSendParamFail(t *testing.T) {
  // TODO
}
