# TinyCMS

cms system

## 基本运行

```bash
tinycms -version             # 显示服务程序的当前版本
tinycms                      # 运行服务程序
```

## http

```bash
# 身份认证模式分为Red与Blue两个等级，都是通过request的head中 Authorization 进行传递
"Authorization": "Red dXNlcjE6cGFzc3dvcmQ="
# 或
"Authorization": "Blue dXNlcjE6cGFzc3dvcmQ="
```

无论`Red`或`Blue`，后面内容格式被base64解码后，都应为 user:token

**Red**表示后面内容中 user应该存在有意义用户id

**Blue**则表明后面格式中用户名部分无意义或无人登录。该部分主要用于核定客户端发回的token是否有效，用于简易判断是否爬虫程序还是正常的浏览器访问
