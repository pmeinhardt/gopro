package gopro

import (
  "net/http"
  "os"
  "testing"
)

func TestListFiles(t *testing.T) {
  cam := DefaultCamera("testpass")

  handler := func(req *http.Request) (*http.Response, error) {
    file, err := os.Open("test/100GOPRO.html")

    if err != nil {
      t.Errorf("Failed to read test file 100GOPRO.html: %v", err)
    }

    return &http.Response{StatusCode: 200, Body: file}, nil
  }

  transport := &MockTransport{handler: handler}
  cam.client = &http.Client{Transport: transport}

  files, err := cam.ListFiles(DefaultMediaPath)

  if err != nil {
    t.Errorf("cam.ListFiles returned an error: %v", err)
  }

  if length := len(files); length != 3 {
    t.Errorf("cam.ListFiles returned %d, expected %d", length, 3)
  }

  names := []string{
    "http://10.5.5.9:8080/DCIM/100GOPRO/GOPR0017.LRV",
    "http://10.5.5.9:8080/DCIM/100GOPRO/GOPR0017.MP4",
    "http://10.5.5.9:8080/DCIM/100GOPRO/GOPR0017.THM",
  }

  for i, expected := range names {
    if actual := files[i].Name; actual != expected {
      t.Errorf("cam.ListFiles returned [%d] %v, expected %v", i, actual, expected)
    }
  }
}

func TestDownload(t *testing.T) {
  // TODO
}
