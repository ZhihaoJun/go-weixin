package cacher

import (
)

type Storage struct {

}

type Cacher struct {
    wx Weixin
    AccessTokenKey string
    JSSDKTicketKey string
    storages []*Storage
}

func New(wx Weixin) *Cacher {
    return Cacher{
        wx: wx,
    }
}

func (this *Cacher) SetAccessTokenKey(key string) {
    this.AccessTokenKey = key
}

func (this *Cacher) SetJSSDKTicketKey(key string) {
    this.JSSDKTicketKey = key
}

func (this *Cacher) AddStorage(storage *Storage) {
    this.storages = append(this.storages, storage)
}

func (this *Cacher) Run() {

}
