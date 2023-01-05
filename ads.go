package ginads

import (
	"encoding/json"

	"github.com/beranek1/goadsinterface"
	"github.com/gin-gonic/gin"
)

type Backend struct {
	lib goadsinterface.AdsLibrary
}

func Create(adsLib goadsinterface.AdsLibrary) *Backend {
	return &Backend{adsLib}
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

func (b *Backend) GetVersion(c *gin.Context) {
	dat, err := b.lib.GetVersion()
	returnADSResult(c, dat, err)
}

func (b *Backend) GetState(c *gin.Context) {
	dat, err := b.lib.GetState()
	returnADSResult(c, dat, err)
}

func (b *Backend) GetDeviceInfo(c *gin.Context) {
	dat, err := b.lib.GetDeviceInfo()
	returnADSResult(c, dat, err)
}

func (b *Backend) GetSymbol(c *gin.Context) {
	name := c.Param("name")
	dat, err := b.lib.GetSymbol(name)
	returnADSResult(c, dat, err)
}

func (b *Backend) GetSymbolInfo(c *gin.Context) {
	dat, err := b.lib.GetSymbolInfo()
	returnADSResult(c, dat, err)
}

func (b *Backend) GetSymbolValue(c *gin.Context) {
	name := c.Param("name")
	dat, err := b.lib.GetSymbolValue(name)
	returnADSResult(c, dat, err)
}

func (b *Backend) GetSymbolList(c *gin.Context) {
	dat, err := b.lib.GetSymbolList()
	returnADSResult(c, dat, err)
}

func (b *Backend) SetState(c *gin.Context) {
	rawData, err := c.GetRawData()
	if err == nil {
		var data goadsinterface.AdsState
		err := json.Unmarshal(rawData, &data)
		if err == nil {
			dat, err := b.lib.SetState(data)
			returnADSResult(c, dat, err)
		} else {
			c.String(500, "{\"error\":\""+err.Error()+"\"}")
		}
	} else {
		c.String(500, "{\"error\":\""+err.Error()+"\"}")
	}
}

func (b *Backend) SetSymbolValue(c *gin.Context) {
	name := c.Param("name")
	rawData, err := c.GetRawData()
	if err == nil {
		var data goadsinterface.AdsData
		err := json.Unmarshal(rawData, &data)
		if err == nil {
			dat, err := b.lib.SetSymbolValue(name, data)
			returnADSResult(c, dat, err)
		} else {
			c.String(500, "{\"error\":\""+err.Error()+"\"}")
		}
	} else {
		c.String(500, "{\"error\":\""+err.Error()+"\"}")
	}
}

func (b *Backend) AttachToRouter(path string, r *gin.Engine) {
	r.GET(path+"/version", b.GetVersion)
	r.GET(path+"/state", b.GetState)
	r.POST(path+"/state", b.SetState)
	r.GET(path+"/device/info", b.GetDeviceInfo)
	r.GET(path+"/symbol/:name", b.GetSymbol)
	r.GET(path+"/symbol", b.GetSymbolInfo)
	r.GET(path+"/symbol/:name/value", b.GetSymbolValue)
	r.POST(path+"/symbol/:name/value", b.SetSymbolValue)
}

func (b *Backend) SetupRouter() *gin.Engine {
	r := gin.Default()
	r.SetTrustedProxies(nil)
	b.AttachToRouter("", r)
	return r
}
