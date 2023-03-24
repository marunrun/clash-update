# clash-update

> 引用项目 https://github.com/XIU2/CloudflareSpeedTest
> 
> 因为`go get github.com/XIU2/CloudflareSpeedTest` 报错所以使用 git submodule 的方式引入这个包。
> 
> 参考项目 https://github.com/Loyalsoldier/clash-rules
> 


----

## 使用

注意将`example.yaml`中的`proxies`替换成你自己的配置,运行软件的时候将代理关闭。
```bash
git clone --recurse-submodules https://github.com/marunrun/clash-update.git

cd clash-update

cp example.yaml clash.yaml

go run main.go
```



替换好的文件， 你可以在本地跑个nginx，然后托管配置`http://127.0.0.1:7880/clash.yaml`

```conf
server {
    listen   7880;
    server_name  _;

    charset utf-8;
    root ~/code/clash-update;
    
    access_log off;
}
```
