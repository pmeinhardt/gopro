package gopro

import (
  "net/http"
  "net/url"
)

type Camera struct {
  ipaddress string
  password string
}

const DefaultIP string = "10.5.5.9"

func NewCamera(ipaddress, password string) *Camera {
  return &Camera{ipaddress, password}
}

func (cam *Camera) url(action string, params *map[string]string) string {
  u := url.URL{}
  q := url.Values{}

  u.Scheme = "http"
  u.Host = cam.ipaddress
  u.Path = "/camera/" + action

  if params != nil {
    for key, value := range *params {
      q.Set(key, value)
    }
  }

  q.Set("t", cam.password)

  u.RawQuery = q.Encode()

  return u.String()
}

func (cam *Camera) Send(action string) error {
  _, err := http.Get(cam.url(action, nil))
  return err
}

func (cam *Camera) SendParam(action string, param int) error {
  params := map[string]string{"p": string(param)}
  _, err := http.Get(cam.url(action, &params))
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

func (cam *Camera) DeleteLast() error {
  return cam.Send("DL")
}

func (cam *Camera) DeleteAll() error {
  return cam.Send("DA")
}
