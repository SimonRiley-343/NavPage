# navpage

## 简介

### 前端

前端使用 `VUE`

- 使用 `Element` 框架
- 通过 `axios` 库处理 HTTP 请求

由于后端项目直接读取自身目录下 `webdist` 目录文件作为前端资源，出于方便调试的目的，在 `package.json` 中 `scripts build` 项添加了 `--dest backend/webdist` 配置项，用于将 `npm run build` 命令的目标路径修改为 `backend/webdist`

### 后端

后端使用 `Golang`

- 使用 `go-gin`，`go-router` 提供 HTTP API
- 使用 `cors` 处理跨域请求问题
- 使用 `sqlite3` 存储配置
- 使用 `Go Modules` 进行包管理

### 项目结构

前端项目位于根目录 `/`，其中 `/src` 为前端项目源码，`/src/assets` 目录用于存储前端静态资源

后端项目位于 `/backend`

## 部署

### 使用构建好的可执行文件

```shell
./navpage -p 8080
# -p 参数指定业务端口
```

### 自行构建

```shell
git clone https://git.leviatan.cn/WindSpirit/navpage.git
cd navpage/
npm install
npm run build
cd backend/
go build
```

### Nginx 代理

```conf
server {
    listen      443 ssl;
    server_name domain.com;
    root        path/to/gitbook;

    location = / {
        proxy_pass          http://localhost:8081/;
        proxy_redirect      off;
        proxy_set_header    Host $host;
        proxy_set_header    X-Real-IP $remote_addr;
        proxy_set_header    X-Forwarded-For $proxy_add_x_forwarded_for;
    }

    location ~ ^/(css|js|img) {
        proxy_pass          http://localhost:8081$request_uri;
        proxy_redirect      off;
        proxy_set_header    Host $host;
        proxy_set_header    X-Real-IP $remote_addr;
        proxy_set_header    X-Forwarded-For $proxy_add_x_forwarded_for;
    }

    location /api/ {
        proxy_pass          http://localhost:8081/;
        proxy_redirect      off;
        proxy_set_header    Host $host;
        proxy_set_header    X-Real-IP $remote_addr;
        proxy_set_header    X-Forwarded-For $proxy_add_x_forwarded_for;
    }
}
```

## TODO List

### 前端

- [x] 前台基础框架
- [ ] 登录页
- [ ] 后台基础框架
- [ ] 后台配置页
- [ ] 前台排序选择

### 后端

- [x] 获取所有页面配置
- [x] 处理前端页面路由
- [x] 登陆认证
- [ ] 配置接口
