##### srs接口回调服务

- 简单的api回调接口,所有接口皆为POST方式
    - `/api/publish`,开始推流api回调接口
    - `/api/un_publish`,停止推流api回调接口
    - `/api/dvr`,录制视频api回调接口
    - `/api/backend`,forward的api回调接口
    - `/api/play`,客户端开始播放api回调接口
    - `/api/stop`,客户端停止播放api回调接口
    - `/api/hls`,ts视频api回调接口
    - `/api/test`,api回调接口信息打印
- 以下接口暂未提供，因为客户端连接和关闭会很频繁，可以在客户端播放和停止播放来处理权限
    - `/api/connect`,客户端连接api回调接口，暂未提供
    - `/api/disconnect`,客户端断开连接api回调接口，暂未提供
