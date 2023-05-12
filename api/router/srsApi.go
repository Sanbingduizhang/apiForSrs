package router

import (
	"github.com/gin-gonic/gin"

	"apiForSrs/api/ctrl"
)

func SrsApi(r *gin.RouterGroup) {
	r.POST("test", ctrl.NewSrsCtrl.Test)
	// on_publish
	r.POST("publish", ctrl.NewSrsCtrl.Publish)
	// on_unpublish
	r.POST("un_publish", ctrl.NewSrsCtrl.UnPublish)
	// on_dvr
	r.POST("dvr", ctrl.NewSrsCtrl.Dvr)
	// on_forward
	r.POST("backend", ctrl.NewSrsCtrl.Backend)
	// on_play
	r.POST("play", ctrl.NewSrsCtrl.Play)
	// on_stop
	r.POST("stop", ctrl.NewSrsCtrl.Stop)
	// on_hls
	r.POST("hls", ctrl.NewSrsCtrl.Hls)

}
