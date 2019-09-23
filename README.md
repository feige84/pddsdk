#PDDSDK

自用SDK，看得懂的直接拿去用，如果看不懂不理解，可以联系我。

已经封装了一些常用接口。其它的可以按需自己扩展。

支持文件缓存、redis缓存。还可以支持失败重新请求，超次数重新请求。

如果有需要也可以做ClientId切换请求。

```go
	client := NewClient("xxx", "xxxx")
	//client.GetCache = FileGetCache //缓存读取方法
	//client.SetCache = FileSetCache //缓存存储方法
	//client.WriteErrLog = InsertErrLog //写错误日志的方法
	client.CacheLife = 0 //缓存生命周期（秒）
	client.AccessToken = "" //需要授权的接口需要这个。
    resp, err := client.Execute("pdd.ddk.goods.search", ApiParams{})
	if err != nil {
		return nil, err
	}
```