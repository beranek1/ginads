package ginads

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type successLib struct {
	version     any
	state       uint16
	deviceInfo  any
	symbolValue any
	symbolList  any
}

func (l *successLib) GetVersion() (any, error) {
	return l.version, nil
}

func (l *successLib) GetState() (any, error) {
	return l.state, nil
}

func (l *successLib) GetDeviceInfo() (any, error) {
	return l.deviceInfo, nil
}

func (l *successLib) GetSymbolInfo(name string) (any, error) {
	return name, nil
}

func (l *successLib) GetSymbolValue(_ string) (any, error) {
	return l.symbolValue, nil
}

func (l *successLib) ListSymbols() (any, error) {
	return l.symbolList, nil
}

func (l *successLib) SetSymbolValue(_ string, value any) (any, error) {
	l.symbolValue = value
	return l.symbolValue, nil
}

func (l *successLib) WriteControl(adsState uint16, deviceState uint16) (any, error) {
	l.state = adsState + deviceState
	return l.state, nil
}

type errorLib struct {
}

func (l *errorLib) GetVersion() (any, error) {
	return nil, errors.New("")
}

func (l *errorLib) GetState() (any, error) {
	return nil, errors.New("")
}

func (l *errorLib) GetDeviceInfo() (any, error) {
	return nil, errors.New("")
}

func (l *errorLib) GetSymbolInfo(_ string) (any, error) {
	return nil, errors.New("")
}

func (l *errorLib) GetSymbolValue(_ string) (any, error) {
	return nil, errors.New("")
}

func (l *errorLib) ListSymbols() (any, error) {
	return nil, errors.New("")
}

func (l *errorLib) SetSymbolValue(_ string, value any) (any, error) {
	return nil, errors.New("")
}

func (l *errorLib) WriteControl(adsState uint16, deviceState uint16) (any, error) {
	return nil, errors.New("")
}

func TestGetVersionSuccess(t *testing.T) {
	l := &successLib{1, 2, 3, 4, 5}
	b := &Backend{l}
	router := b.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/version", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "1", w.Body.String())
}

func TestGetVersionError(t *testing.T) {
	l := &errorLib{}
	b := &Backend{l}
	router := b.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/version", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 500, w.Code)
	assert.Equal(t, "{\"error\":\"\"}", w.Body.String())
}

func TestGetStateSuccess(t *testing.T) {
	l := &successLib{1, 2, 3, 4, 5}
	b := &Backend{l}
	router := b.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/state", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "2", w.Body.String())
}

func TestGetStateError(t *testing.T) {
	l := &errorLib{}
	b := &Backend{l}
	router := b.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/state", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 500, w.Code)
	assert.Equal(t, "{\"error\":\"\"}", w.Body.String())
}

func TestGetDeviceInfoSuccess(t *testing.T) {
	l := &successLib{1, 2, 3, 4, 5}
	b := &Backend{l}
	router := b.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/deviceInfo", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "3", w.Body.String())
}

func TestGetDeviceInfoError(t *testing.T) {
	l := &errorLib{}
	b := &Backend{l}
	router := b.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/deviceInfo", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 500, w.Code)
	assert.Equal(t, "{\"error\":\"\"}", w.Body.String())
}

func TestGetSymbolInfoSuccess(t *testing.T) {
	l := &successLib{1, 2, 3, 4, 5}
	b := &Backend{l}
	router := b.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/symbolInfo/test", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "\"test\"", w.Body.String())
}

func TestGetSymbolInfoError(t *testing.T) {
	l := &errorLib{}
	b := &Backend{l}
	router := b.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/symbolInfo/test", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 500, w.Code)
	assert.Equal(t, "{\"error\":\"\"}", w.Body.String())
}

func TestGetSymbolValueSuccess(t *testing.T) {
	l := &successLib{1, 2, 3, 4, 5}
	b := &Backend{l}
	router := b.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/symbolValue/test", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "4", w.Body.String())
}

func TestGetSymbolValueError(t *testing.T) {
	l := &errorLib{}
	b := &Backend{l}
	router := b.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/symbolValue/test", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 500, w.Code)
	assert.Equal(t, "{\"error\":\"\"}", w.Body.String())
}

func TestListSymbolsSuccess(t *testing.T) {
	l := &successLib{1, 2, 3, 4, 5}
	b := &Backend{l}
	router := b.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/symbolList", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "5", w.Body.String())
}

func TestListSymbolsError(t *testing.T) {
	l := &errorLib{}
	b := &Backend{l}
	router := b.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/symbolList", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 500, w.Code)
	assert.Equal(t, "{\"error\":\"\"}", w.Body.String())
}

func TestSetSymbolValueSuccess(t *testing.T) {
	l := &successLib{1, 2, 3, 4, 5}
	b := &Backend{l}
	router := b.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/symbolValue/test", strings.NewReader("{\"data\":6}"))
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "6", w.Body.String())
}

func TestSetSymbolValueError(t *testing.T) {
	l := &errorLib{}
	b := &Backend{l}
	router := b.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/symbolValue/test", strings.NewReader(""))
	router.ServeHTTP(w, req)

	assert.Equal(t, 500, w.Code)
	assert.Equal(t, "{\"error\":\"unexpected end of JSON input\"}", w.Body.String())
}

func TestWriteControlSuccess(t *testing.T) {
	l := &successLib{1, 2, 3, 4, 5}
	b := &Backend{l}
	router := b.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/state", strings.NewReader("{\"adsState\":5,\"deviceState\":7}"))
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "12", w.Body.String())

	w = httptest.NewRecorder()
	req2, _ := http.NewRequest("POST", "/state", strings.NewReader("{\"adsState\":5}"))
	router.ServeHTTP(w, req2)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "5", w.Body.String())

	w = httptest.NewRecorder()
	req3, _ := http.NewRequest("POST", "/state", strings.NewReader("{\"deviceState\":7}"))
	router.ServeHTTP(w, req3)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "7", w.Body.String())
}

func TestWriteControlError(t *testing.T) {
	l := &errorLib{}
	b := &Backend{l}
	router := b.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/state", strings.NewReader(""))
	router.ServeHTTP(w, req)

	assert.Equal(t, 500, w.Code)
	assert.Equal(t, "{\"error\":\"unexpected end of JSON input\"}", w.Body.String())
}
