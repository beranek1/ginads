package ginads

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
)

type ADSLib[T any] interface {
	GetVersion() (T, error)
	GetState() (T, error)
	GetDeviceInfo() (T, error)
	GetSymbolInfo(name string) (T, error)
	GetSymbolValue(name string) (T, error)
	ListSymbols() (any, error)
	SetSymbolValue(name string, value any) (T, error)
	WriteControl(adsState uint16, deviceState uint16) (T, error)
}

type Backend[T any] struct {
	lib ADSLib[T]
}

func Create[T any](adsLib ADSLib[T]) *Backend[T] {
	return &Backend[T]{adsLib}
}

func returnADSResult(c *gin.Context, dat any, err error) {
	if err != nil {
		c.String(500, "{\"error\":\""+err.Error()+"\"}")
	} else {
		byt, err := json.Marshal(dat)
		if err != nil {
			c.String(500, "{\"error\":\""+err.Error()+"\"}")
		} else {
			c.String(200, string(byt))
		}
	}
}

func (b *Backend[T]) GetVersion(c *gin.Context) {
	dat, err := b.lib.GetVersion()
	returnADSResult(c, dat, err)
}

func (b *Backend[T]) GetState(c *gin.Context) {
	dat, err := b.lib.GetState()
	returnADSResult(c, dat, err)
}

func (b *Backend[T]) GetDeviceInfo(c *gin.Context) {
	dat, err := b.lib.GetDeviceInfo()
	returnADSResult(c, dat, err)
}

func (b *Backend[T]) GetSymbolInfo(c *gin.Context) {
	name := c.Param("name")
	dat, err := b.lib.GetSymbolInfo(name)
	returnADSResult(c, dat, err)
}

func (b *Backend[T]) GetSymbolValue(c *gin.Context) {
	name := c.Param("name")
	dat, err := b.lib.GetSymbolValue(name)
	returnADSResult(c, dat, err)
}

func (b *Backend[T]) ListSymbols(c *gin.Context) {
	dat, err := b.lib.ListSymbols()
	returnADSResult(c, dat, err)
}

func (b *Backend[T]) SetSymbolValue(c *gin.Context) {
	name := c.Param("name")
	rawData, err := c.GetRawData()
	if err == nil {
		var data map[string]interface{}
		err := json.Unmarshal(rawData, &data)
		if err == nil {
			if value, exists := data["data"]; exists {
				dat, err := b.lib.SetSymbolValue(name, value)
				returnADSResult(c, dat, err)
			} else {
				dat, err := b.lib.GetSymbolValue(name)
				returnADSResult(c, dat, err)
			}
		} else {
			c.String(500, "{\"error\":\""+err.Error()+"\"}")
		}
	} else {
		c.String(500, "{\"error\":\""+err.Error()+"\"}")
	}
}

func (b *Backend[T]) WriteControl(c *gin.Context) {
	rawData, err := c.GetRawData()
	if err == nil {
		var data map[string]uint16
		err := json.Unmarshal(rawData, &data)
		if err == nil {
			adsState := data["adsState"]
			deviceState := data["deviceState"]
			dat, err := b.lib.WriteControl(adsState, deviceState)
			returnADSResult(c, dat, err)
		} else {
			c.String(500, "{\"error\":\""+err.Error()+"\"}")
		}
	} else {
		c.String(500, "{\"error\":\""+err.Error()+"\"}")
	}
}

func (b *Backend[T]) SetupRouter() *gin.Engine {
	r := gin.Default()
	r.SetTrustedProxies(nil)
	r.GET("/version", b.GetVersion)

	r.GET("/state", b.GetState)

	r.POST("/state", b.WriteControl)

	r.GET("/deviceInfo", b.GetDeviceInfo)

	r.GET("/symbolInfo/:name", b.GetSymbolInfo)

	r.GET("/symbolValue/:name", b.GetSymbolValue)

	r.POST("/symbolValue/:name", b.SetSymbolValue)

	r.GET("/symbolList", b.ListSymbols)

	return r
}
