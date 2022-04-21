快递100 相关接口

目前实现：实时快递查询、快递查询地图轨迹

`下载`
```shell
    go get github.com/ymboom0042/kd100
```

`使用`
````go
    express, err := kd100.NewExpress("customer","key")
    if err != nil {
    panic("快递业务失败")
    }

    p := make(kd100.Param)
    p.set("num","num").set("com","yuantong")
    // 实时快递查询
    //resp,err := express.SyncQuery(p)
    // 快递查询地图轨迹
    resp,err := express.mapTrack(p)
````