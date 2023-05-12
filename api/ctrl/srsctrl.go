package ctrl

import (
	"apiForSrs/api/response"
	"apiForSrs/api/validate"
	"apiForSrs/global"
	"apiForSrs/poolservice"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
)

var (
	NewSrsCtrl = &srsCtrl{}
)

type srsCtrl struct {
}

// on_stop
func (s *srsCtrl) Stop(ctx *gin.Context) {
	var form validate.PlayStopForm
	if err := ctx.ShouldBind(&form); err != nil {
		response.SrsErrJson(ctx)
		return
	}

	// 可以启动携程
	poolservice.AddFuncWork(func() {
		by, _ := json.Marshal(form)
		fmt.Println("Stop:", string(by))
	})

	// 开始解析数据
	response.SrsOkJson(ctx)
	return

}

// Play on_play
// 可以在当前接口中统计有多少客户端播放
// 可以处理是否允许客户端播放
func (s *srsCtrl) Play(ctx *gin.Context) {
	var form validate.PlayStopForm
	if err := ctx.ShouldBind(&form); err != nil {
		response.SrsErrJson(ctx)
		return
	}

	// 可以启动携程
	poolservice.AddFuncWork(func() {
		by, _ := json.Marshal(form)
		fmt.Println("play获取到数据信息:", string(by))
	})

	// 开始解析数据
	response.SrsOkJson(ctx)
	return

}

// Backend forward到录制服务
// 处理是否允许某个app下的stream流或者是某一个单独的stream流是否需要录制
func (s *srsCtrl) Backend(ctx *gin.Context) {
	var form validate.BackendForm
	if err := ctx.ShouldBind(&form); err != nil {
		response.SrsErrForwardJson(ctx)
		return
	}

	// 需要及时获取数据信息,判断是否需要转推
	// 此处直接返回空
	response.SrsErrForwardJson(ctx)
	return
}

// Hls m3u8文件接收处理
// hls生成的ts文件应该自行处理，待视频停止播流时生成单独的m3u8文件，供回放使用
func (s *srsCtrl) Hls(ctx *gin.Context) {
	var form validate.HlsForm
	if err := ctx.ShouldBind(&form); err != nil {
		response.SrsErrJson(ctx)
		return
	}

	// 启动协程处理，有排序和时间，即便不是顺序发送也可以正常排序的
	poolservice.AddFuncWork(func() {
		by, _ := json.Marshal(form)
		fmt.Println("Hls获取到数据信息:", string(by))
	})

	// 开始解析数据
	response.SrsOkJson(ctx)
	return

}

// Dvr 录制
// 录制成mp4，可以自行处理视频呢是否需要剪辑等
func (s *srsCtrl) Dvr(ctx *gin.Context) {
	var form validate.DvrForm
	if err := ctx.ShouldBind(&form); err != nil {
		response.SrsErrJson(ctx)
		return
	}

	// 录制的无需通过队列顺序发送，可以直接有协程池中获取worker进行处理
	poolservice.AddFuncWork(func() {
		by, _ := json.Marshal(form)
		fmt.Println("Dvr获取到数据信息:", string(by))
	})

	// 开始解析数据
	response.SrsOkJson(ctx)
	return

}

// UnPublish 客户端停止推流，
func (s *srsCtrl) UnPublish(ctx *gin.Context) {
	var form validate.UnPublishForm
	if err := ctx.ShouldBind(&form); err != nil {
		response.SrsErrJson(ctx)
		return
	}

	// 写入队列，顺序发送,后台执行消息推送
	global.RunnerBox.AddTask(func() {
		by, _ := json.Marshal(form)
		fmt.Println("UnPublish获取到数据信息:", string(by))
	})

	// 开始解析数据
	response.SrsOkJson(ctx)
	return

}

// Publish 推流
// 可以在客户端推流中增加对推流地址的鉴权，限制等功能
func (s *srsCtrl) Publish(ctx *gin.Context) {
	var form validate.PublishForm
	if err := ctx.ShouldBind(&form); err != nil {
		response.SrsErrJson(ctx)
		return
	}
	if form.Stream == "" {
		fmt.Println("publish的流为空啊:", form.App)
		response.SrsErrJson(ctx)
		return
	}

	// 如果一个视频流是否多次重复推流

	// 判断是否需要判断app推流时符合规则

	// 写入队列，顺序发送,后台执行消息推送
	global.RunnerBox.AddTask(func() {
		by, _ := json.Marshal(form)
		fmt.Println("Publish获取到数据信息:", string(by))
	})

	// 开始解析数据
	response.SrsOkJson(ctx)
	return

}

// 打印信息接口
func (s *srsCtrl) Test(ctx *gin.Context) {
	var data interface{}
	if err := ctx.ShouldBind(&data); err != nil {
		response.SrsErrJson(ctx)
		return
	}
	//// 开始解析数据
	by, _ := json.Marshal(data)
	fmt.Println("Test获取到数据信息：", string(by))

	response.SrsOkJson(ctx)
	return

}
