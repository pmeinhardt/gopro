package gopro

import (
	"bufio"
	"io"
	"io/ioutil"
	"net/url"
	"os"
	"regexp"
)

const DefaultMediaPath string = "/DCIM/100GOPRO/"

func (cam *Camera) weburl(path string) string {
	u := url.URL{}

	u.Scheme = "http"
	u.Host = cam.ipaddress + ":8080"
	u.Path = path

	return u.String()
}

type FileInfo struct {
	Name  string
	Size  int64
	IsDir bool
}

func (cam *Camera) ListFiles(dirpath string) (fi []FileInfo, err error) {
	if dirpath[len(dirpath)-1] != '/' {
		dirpath += "/"
	}

	u := cam.weburl(dirpath)
	res, err := cam.get(u)

	if err != nil {
		return
	}

	defer res.Body.Close()
	html, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return
	}

	re, err := regexp.Compile(u + "[^?#\" ]+")

	if err != nil {
		return
	}

	names := re.FindAllString(string(html), -1)

	fi = make([]FileInfo, len(names))

	for i, name := range names {
		fi[i].Name = name
	}

	return
}

func (cam *Camera) Download(src, dest string) (err error) {
	res, err := cam.get(src)

	if err != nil {
		return
	}

	file, err := os.OpenFile(dest, os.O_CREATE|os.O_EXCL, 0644)

	if err != nil {
		return
	}

	defer file.Close()

	_, err = io.Copy(bufio.NewWriter(file), res.Body)

	return
}
