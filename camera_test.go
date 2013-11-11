package gopro

import (
	"errors"
	"net/http"
	"testing"
)

func camcmp(t *testing.T, actual, expected *Camera) {
	if actual.client != expected.client {
		t.Errorf("cam.client not set correctly")
	}

	if actual.ipaddress != expected.ipaddress {
		t.Errorf("cam.ipaddress set to %v, expected %v", actual.ipaddress, expected.ipaddress)
	}

	if actual.password != expected.password {
		t.Errorf("cam.password set to %v, expected %v", actual.password, expected.password)
	}
}

func TestNewCamera(t *testing.T) {
	exp := &Camera{http.DefaultClient, "1.2.3.4", "password"}
	cam := NewCamera("1.2.3.4", "password")
	camcmp(t, cam, exp)
}

func TestDefaultCamera(t *testing.T) {
	exp := &Camera{http.DefaultClient, DefaultIP, "password"}
	cam := DefaultCamera("password")
	camcmp(t, cam, exp)
}

type MockRequestHandler func(req *http.Request) (*http.Response, error)

type MockTransport struct {
	handler MockRequestHandler
	request *http.Request
}

func (t *MockTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	res, err := t.handler(req)
	t.request = req
	return res, err
}

func TestSend(t *testing.T) {
	cam := DefaultCamera("testpass")

	handler := func(req *http.Request) (*http.Response, error) {
		return &http.Response{StatusCode: 200}, nil
	}

	transport := &MockTransport{handler: handler}
	cam.client = &http.Client{Transport: transport}

	err := cam.Send("DL")
	req := transport.request

	if err != nil {
		t.Errorf("cam.Send returned an error: %v", err)
	}

	if req.Method != "GET" {
		t.Errorf("cam.Send used method %v, expected %v", req.Method, "GET")
	}

	if req.URL.Host != cam.ipaddress {
		t.Errorf("cam.Send requested host %v, expected %v", req.URL.Host, cam.ipaddress)
	}

	if req.URL.Path != "/camera/DL" {
		t.Errorf("cam.Send requested wrong path: %v", req.URL.Path)
	}

	query := req.URL.Query()

	if query.Get("t") != "testpass" {
		t.Errorf("cam.Send should send password")
	}
}

func TestSendFail(t *testing.T) {
	cam := DefaultCamera("testpass")

	handler := func(req *http.Request) (*http.Response, error) {
		return nil, errors.New("this is bad")
	}

	transport := &MockTransport{handler: handler}
	cam.client = &http.Client{Transport: transport}

	res := cam.Send("DL")

	if res == nil {
		t.Errorf("cam.Send should return HTTP error")
	}
}

func TestSendParam(t *testing.T) {
	cam := DefaultCamera("testpass")

	handler := func(req *http.Request) (*http.Response, error) {
		return &http.Response{StatusCode: 200}, nil
	}

	transport := &MockTransport{handler: handler}
	cam.client = &http.Client{Transport: transport}

	err := cam.SendParam("SH", 1)
	req := transport.request

	if err != nil {
		t.Errorf("cam.SendParam returned an error: %v", err)
	}

	if req.Method != "GET" {
		t.Errorf("cam.SendParam used method %v, expected %v", req.Method, "GET")
	}

	if req.URL.Host != cam.ipaddress {
		t.Errorf("cam.SendParam requested host %v, expected %v", req.URL.Host, cam.ipaddress)
	}

	if req.URL.Path != "/camera/SH" {
		t.Errorf("cam.SendParam requested wrong path: %v", req.URL.Path)
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
	cam := DefaultCamera("testpass")

	handler := func(req *http.Request) (*http.Response, error) {
		return nil, errors.New("this is bad")
	}

	transport := &MockTransport{handler: handler}
	cam.client = &http.Client{Transport: transport}

	res := cam.SendParam("SH", 1)

	if res == nil {
		t.Errorf("cam.SendParam should return HTTP error")
	}
}
