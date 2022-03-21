package v1

import (
	"log"

	"github.com/akazwz/fhub/model/response"
	"github.com/anacrolix/torrent"
	"github.com/gin-gonic/gin"
)

// DownloadByMagnet 根据magnet 下载
func DownloadByMagnet(c *gin.Context) {
	tClient, err := torrent.NewClient(nil)
	defer tClient.Close()
	if err != nil {
		log.Println(err)
		response.BadRequest(4000, "new torrent client error", c)
		return
	}
	magnet, err := tClient.AddMagnet("magnet:?xt=urn:btih:ZOCMZQIPFFW7OLLMIC5HUB6BPCSDEOQU")
	if err != nil {
		log.Println(err)
		response.BadRequest(4000, "add magnet error", c)
		return
	}
	<-magnet.GotInfo()
	info := magnet.Info()
	response.Ok(2000, info, "get meta info success", c)
	return
}
