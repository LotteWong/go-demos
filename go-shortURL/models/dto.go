package models

type ShortLinkReq struct {
	Url          string `json:"url" validate:"required"`
	ExpireWithin int64  `json:"expire_within" validate:"min=0"`
}

type ShortLinkRes struct {
	ShortLink string `json:"shortlink"`
}

type HealthCheckDto struct {
	Msg string `json:"msg"`
}
