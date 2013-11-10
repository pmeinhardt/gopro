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

func (cam *Camera) queryURL(action string, param int) string {
  u := url.URL{}

  u.Scheme = "http"
  u.Host = cam.ipaddress
  u.Path = "/camera/" + action

  params := url.Values{}
  params.Set("t", cam.password)
  params.Set("p", `%0` + strconv.Itoa(param))

  u.RawQuery = params.Encode()

  return u.String()
}

// func (cam *Camera) Send(action string) error {
// }

func (cam *Camera) SendParam(action string, param int) error {
  _, err := http.Get(cam.queryURL(action, param))
  return err
}

func (cam *Camera) StartCapture() error {
  return cam.SendParam("SH", 1)
}

func (cam *Camera) StopCapture() error {
  return cam.SendParam("SH", 0)
}

func (cam *Camera) StartBeeping() error {
  return cam.SendParam("LL", 1)
}

func (cam *Camera) StopBeeping() error {
  return cam.SendParam("LL", 0)
}

// func (cam *Camera) DeleteLast() error {
//   return cam.Send("DL")
// }

// func (cam *Camera) DeleteAll() error {
//   return cam.Send("DA")
// }
