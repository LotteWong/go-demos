package models

import "time"

type UrlDetail struct {
	Url          string        `json:"url"`
	CreateAt     string        `json:"create_at"`
	ExpireWithin time.Duration `json:"expire_within"`
}
