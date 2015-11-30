# weixin
## cacher
cacher := weixin.NewCacher(wx)
redisStorage := weixin.NewRediStorage(redisAddress)
memcacheStorage := weixin.NewMemcacheStorage(memcacheAddress)
cacher.AddStorage(redisStorage)
cacher.AddStorage(memcacheStorage)
cacher.SetAccessTokenKey(accessTokenKey)
cacher.SetJSSDKTicketKey(jssdkTicketKey)
cacher.Run()
