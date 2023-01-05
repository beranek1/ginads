package ginads

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/beranek1/goadsinterface"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type successLib struct {
	version     goadsinterface.AdsVersion
	state       goadsinterface.AdsState
	deviceInfo  goadsinterface.AdsDeviceInfo
	symbolValue goadsinterface.AdsData
	symbolList  goadsinterface.AdsSymbolList
	symbolInfo  goadsinterface.AdsSymbolInfo
	symbol      goadsinterface.AdsSymbol
}

func (l *successLib) GetVersion() (goadsinterface.AdsVersion, error) {
	return l.version, nil
}

func (l *successLib) GetState() (goadsinterface.AdsState, error) {
	return l.state, nil
}

func (l *successLib) GetDeviceInfo() (goadsinterface.AdsDeviceInfo, error) {
	return l.deviceInfo, nil
}

func (l *successLib) GetSymbol(name string) (goadsinterface.AdsSymbol, error) {
	return l.symbol, nil
}

func (l *successLib) GetSymbolValue(_ string) (goadsinterface.AdsData, error) {
	return l.symbolValue, nil
}

func (l *successLib) GetSymbolInfo() (goadsinterface.AdsSymbolInfo, error) {
	return l.symbolInfo, nil
}

func (l *successLib) GetSymbolList() (goadsinterface.AdsSymbolList, error) {
	return l.symbolList, nil
}

func (l *successLib) SetSymbolValue(_ string, value goadsinterface.AdsData) (goadsinterface.AdsData, error) {
	l.symbolValue = value
	return l.symbolValue, nil
}

func (l *successLib) SetState(state goadsinterface.AdsState) (goadsinterface.AdsState, error) {
	l.state = state
	return l.state, nil
}

type errorLib struct {
}

func (l *errorLib) GetVersion() (goadsinterface.AdsVersion, error) {
	return goadsinterface.AdsVersion{}, errors.New("")
}

func (l *errorLib) GetState() (goadsinterface.AdsState, error) {
	return goadsinterface.AdsState{}, errors.New("")
}

func (l *errorLib) GetDeviceInfo() (goadsinterface.AdsDeviceInfo, error) {
	return goadsinterface.AdsDeviceInfo{}, errors.New("")
}

func (l *errorLib) GetSymbol(_ string) (goadsinterface.AdsSymbol, error) {
	return goadsinterface.AdsSymbol{}, errors.New("")
}

func (l *errorLib) GetSymbolValue(_ string) (goadsinterface.AdsData, error) {
	return goadsinterface.AdsData{}, errors.New("")
}

func (l *errorLib) GetSymbolList() (goadsinterface.AdsSymbolList, error) {
	return goadsinterface.AdsSymbolList{}, errors.New("")
}

func (l *errorLib) GetSymbolInfo() (goadsinterface.AdsSymbolInfo, error) {
	return goadsinterface.AdsSymbolInfo{}, errors.New("")
}

func (l *errorLib) SetSymbolValue(_ string, value goadsinterface.AdsData) (goadsinterface.AdsData, error) {
	return goadsinterface.AdsData{}, errors.New("")
}

func (l *errorLib) SetState(state goadsinterface.AdsState) (goadsinterface.AdsState, error) {
	return goadsinterface.AdsState{}, errors.New("")
}

func createTestBackendRouterSuccess() *gin.Engine {
	v := goadsinterface.AdsVersion{Version: 1, Revision: 2, Build: 3}
	st := goadsinterface.AdsState{Ads: 4, Device: 5}
	di := goadsinterface.AdsDeviceInfo{Name: "test", Version: v}
	dat := goadsinterface.AdsData{Data: "data"}
	sl := []string{"symbol"}
	sy := goadsinterface.AdsSymbol{Name: "symbol", IndexGroup: 0, IndexOffset: 0, Size: 0, Type: "string", Comment: "nocomment"}
	si := goadsinterface.AdsSymbolInfo{}
	si["sumbol"] = sy
	l := &successLib{v, st, di, dat, sl, si, sy}
	b := Create(l)
	return createTestRouter(b)
}

func createTestBackendRouterError() *gin.Engine {
	l := &errorLib{}
	b := Create(l)
	return createTestRouter(b)
}

func createTestRouter(b *Backend) *gin.Engine {
	router := b.SetupRouter()
	return router
}

func TestGetVersionSuccess(t *testing.T) {
	router := createTestBackendRouterSuccess()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/version", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	jsonStr, _ := json.Marshal(goadsinterface.AdsVersion{Version: 1, Revision: 2, Build: 3})
	assert.Equal(t, string(jsonStr), w.Body.String())
}

func TestGetVersionError(t *testing.T) {
	router := createTestBackendRouterError()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/version", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 500, w.Code)
	assert.Equal(t, "{\"error\":\"\"}", w.Body.String())
}

func TestGetStateSuccess(t *testing.T) {
	router := createTestBackendRouterSuccess()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/state", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	jsonStr, _ := json.Marshal(goadsinterface.AdsState{Ads: 4, Device: 5})
	assert.Equal(t, string(jsonStr), w.Body.String())
}

func TestGetStateError(t *testing.T) {
	router := createTestBackendRouterError()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/state", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 500, w.Code)
	assert.Equal(t, "{\"error\":\"\"}", w.Body.String())
}

func TestGetDeviceInfoSuccess(t *testing.T) {
	router := createTestBackendRouterSuccess()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/device/info", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	jsonStr, _ := json.Marshal(goadsinterface.AdsDeviceInfo{Name: "test", Version: goadsinterface.AdsVersion{Version: 1, Revision: 2, Build: 3}})
	assert.Equal(t, string(jsonStr), w.Body.String())
}

func TestGetDeviceInfoError(t *testing.T) {
	router := createTestBackendRouterError()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/device/info", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 500, w.Code)
	assert.Equal(t, "{\"error\":\"\"}", w.Body.String())
}

func TestGetSymbolSuccess(t *testing.T) {
	router := createTestBackendRouterSuccess()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/symbol/test", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	jsonStr, _ := json.Marshal(goadsinterface.AdsSymbol{Name: "symbol", IndexGroup: 0, IndexOffset: 0, Size: 0, Type: "string", Comment: "nocomment"})
	assert.Equal(t, string(jsonStr), w.Body.String())
}

func TestGetSymbolError(t *testing.T) {
	router := createTestBackendRouterError()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/symbol/test", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 500, w.Code)
	assert.Equal(t, "{\"error\":\"\"}", w.Body.String())
}

func TestGetSymbolValueSuccess(t *testing.T) {
	router := createTestBackendRouterSuccess()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/symbol/test/value", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	jsonStr, _ := json.Marshal(goadsinterface.AdsData{Data: "data"})
	assert.Equal(t, string(jsonStr), w.Body.String())
}

func TestGetSymbolValueError(t *testing.T) {
	router := createTestBackendRouterError()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/symbol/test/value", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 500, w.Code)
	assert.Equal(t, "{\"error\":\"\"}", w.Body.String())
}

func TestGetSymbolInfoSuccess(t *testing.T) {
	router := createTestBackendRouterSuccess()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/symbol", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	si := goadsinterface.AdsSymbolInfo{}
	si["sumbol"] = goadsinterface.AdsSymbol{Name: "symbol", IndexGroup: 0, IndexOffset: 0, Size: 0, Type: "string", Comment: "nocomment"}
	jsonStr, _ := json.Marshal(si)
	assert.Equal(t, string(jsonStr), w.Body.String())
}

func TestGetSymbolInfoError(t *testing.T) {
	router := createTestBackendRouterError()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/symbol", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 500, w.Code)
	assert.Equal(t, "{\"error\":\"\"}", w.Body.String())
}

func TestSetSymbolValueSuccess(t *testing.T) {
	router := createTestBackendRouterSuccess()

	w := httptest.NewRecorder()
	jsonStr, _ := json.Marshal(goadsinterface.AdsData{Data: "data"})
	req, _ := http.NewRequest("POST", "/symbol/test/value", strings.NewReader(string(jsonStr)))
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, string(jsonStr), w.Body.String())
}

func TestSetSymbolValueError(t *testing.T) {
	router := createTestBackendRouterError()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/symbol/test/value", strings.NewReader(""))
	router.ServeHTTP(w, req)

	assert.Equal(t, 500, w.Code)
	assert.Equal(t, "{\"error\":\"unexpected end of JSON input\"}", w.Body.String())
}

func TestSetStateSuccess(t *testing.T) {
	router := createTestBackendRouterSuccess()

	w := httptest.NewRecorder()
	jsonStr, _ := json.Marshal(goadsinterface.AdsState{Ads: 4, Device: 5})
	req, _ := http.NewRequest("POST", "/state", strings.NewReader(string(jsonStr)))
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, string(jsonStr), w.Body.String())

	// w = httptest.NewRecorder()
	// req2, _ := http.NewRequest("POST", "/state", strings.NewReader("{\"adsState\":5}"))
	// router.ServeHTTP(w, req2)

	// assert.Equal(t, 200, w.Code)
	// assert.Equal(t, "5", w.Body.String())

	// w = httptest.NewRecorder()
	// req3, _ := http.NewRequest("POST", "/state", strings.NewReader("{\"deviceState\":7}"))
	// router.ServeHTTP(w, req3)

	// assert.Equal(t, 200, w.Code)
	// assert.Equal(t, "7", w.Body.String())
}

func TestSetStateError(t *testing.T) {
	router := createTestBackendRouterError()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/state", strings.NewReader(""))
	router.ServeHTTP(w, req)

	assert.Equal(t, 500, w.Code)
	assert.Equal(t, "{\"error\":\"unexpected end of JSON input\"}", w.Body.String())
}
