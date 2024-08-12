package model

import "time"

type RedisReq struct {
	Key        string      `json:"key" validate:"required"`
	Value      interface{} `json:"value"`
	ExpireTime string      `json:"expireTime"`
}

type SetValueReq struct {
	Key        string
	Value      interface{}
	ExpireTime time.Duration
}
