package validate

type PlayStopForm struct {
	BaseForm
	ClientId string `json:"client_id"`
}

type HlsForm struct {
	BaseForm
	File     string  `json:"file"`
	SeqNo    int64   `json:"seq_no"`
	Duration float64 `json:"duration"`
}

type BaseForm struct {
	Action string `json:"action"`
	App    string `json:"app"`
	Stream string `json:"stream"`
	Param  string `json:"param"`
	TcUrl  string `json:"tcUrl"`
}

type PublishForm struct {
	BaseForm
}

type UnPublishForm struct {
	BaseForm
}

type DvrForm struct {
	BaseForm
	Cwd  string `json:"cwd"`
	File string `json:"file"`
}

type BackendForm struct {
	BaseForm
}
