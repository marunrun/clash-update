# clash-update

引用项目 https://github.com/XIU2/CloudflareSpeedTest

> 因为`go get github.com/XIU2/CloudflareSpeedTest` 报错所以使用 git submodule 的方式引入这个包。

利用cloudflare优先ip更新clash配置文件

复制`example.yaml`为`clash.yaml`，修改`clash.yaml`中的`proxies`为你的代理列表。

` cp example.yaml clash.yaml`

测试代理，并替换文件

`go run main.go`

注意: 在代码中，我将下载测速关闭了，如果你需要的话，可以自行打开。

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
