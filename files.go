package gopro

import (
  "io/ioutil"
  "net/url"
  "regexp"
)

const DefaultMediaPath string = "/DCIM/100GOPRO"

func (cam *Camera) weburl(path string) string {
  u := url.URL{}

  u.Scheme = "http"
  u.Host = cam.ipaddress + ":8080"
  u.Path = path

  return u.String()
}

type FileInfo struct {
  Name string
  Size int64
  IsDir bool
}

func (cam *Camera) ListFiles(dirpath string) (fi []FileInfo, err error) {
  u := cam.weburl(dirpath)
  res, err := cam.get(u)

  if err != nil {
    return fi, err
  }

  defer res.Body.Close()
  html, err := ioutil.ReadAll(res.Body)

  re, err := regexp.Compile(u + "/[^?#\" ]+")

  if err != nil {
    return fi, err
  }

  names := re.FindAllString(string(html), -1)

  fi = make([]FileInfo, len(names))

  for i, name := range names {
    fi[i].Name = name
  }

  return fi, err
}
