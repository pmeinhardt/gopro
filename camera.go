package gopro

import (
  "net/http"
  "net/url"
  "strconv"
)

type Camera struct {
  ipaddress string
  password string
}

func NewCamera(ipaddress, password string) *Camera {
  return &Camera{ipaddress, password}
}

func (cam *Camera) queryString(param ...int) string {
  q := url.Values{}

  q.Set("t", cam.password)

  if len(param) > 0 {
    q.Set("p", `%0` + strconv.Itoa(param[0]))
  }
  return q.Encode()
}

func (cam *Camera) queryURL(action string, param ...int) string {
  u := url.URL{}

  u.Scheme = "http"
  u.Host = cam.ipaddress
  u.Path = "/camera/" + action
  u.RawQuery = cam.queryString(param...)

  return u.String()
}

func (cam *Camera) Send(action string, params ...int) error {
  _, err := http.Get(cam.queryURL(action, params...))
  return err
}

func (cam *Camera) StartCapture() error {
  return cam.Send("SH", 1)
}

func (cam *Camera) StopCapture() error {
  return cam.Send("SH", 0)
}

func (cam *Camera) StartBeeping() error {
  return cam.Send("LL", 1)
}

func (cam *Camera) StopBeeping() error {
  return cam.Send("LL", 0)
}

func (cam *Camera) DeleteLast() error {
  return cam.Send("DL")
}

func (cam *Camera) DeleteAll() error {
  return cam.Send("DA")
}
