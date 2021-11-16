# TinyCMS

cms system

## 基本运行

```bash
tinycms -version             # 显示服务程序的当前版本
tinycms                      # 运行服务程序
```

## 身份验证（中间件）

### 基本意义

```bash
# 身份认证模式分为Red与Blue两个等级，都是通过request的head中 Authorization 进行传递
"Authorization": "Red dXNlcjE6cGFzc3dvcmQ="
# 或
"Authorization": "Blue dXNlcjE6cGFzc3dvcmQ="
```

无论`Red`或`Blue`，后面内容格式被base64解码后，都应为 user:token

**Red**表示后面内容中 user应该存在有意义用户id

**Blue**则表明后面格式中用户名部分无意义或无人登录。该部分主要用于核定客户端发回的token是否有效，用于简易判断是否爬虫程序还是正常的浏览器访问

### 相关配置

第30行决定是否进行token校验（审计）。同时需要redis服务配合，redis服务配置位于第8行 - 11行

```yaml
publishing:
  server: "127.0.0.1:9000"
redis:
  local:
    server: "127.0.0.1:6379"
    pwd: 
    db: 0
  audit_auth_token:   
    server: "127.0.0.1:6379" 
    pwd:
    db: 0
mariadb:
  tcms:
    server: "172.16.1.250"
    port: 3306
    uid: root
    pwd: "123"
    db: tinycms
    timeout: 10
  logclientip:
    server: "172.16.1.250"
    port: 3306
    uid: root
    pwd: "123"
    db: clientiplog
    table: tcms
    timeout: 10    
audit:
  auth_token:  # 审计auth_token策略
    allow: true # 是否允许审计auth_token
    expire_ms: 60000  # token保存有效时间（毫秒）
    jam_ms: 10000  # 请求被审计后，每次请求都会被卡住jam_ms毫秒数，然后才能继续
    repeat: 1 # token被重复n次后会被触发审计
  log_clientip:
    allow: true 
```

## ip记录(中间件)

### 相关配置

第35行决定是否进行来访ip记录。同时需要MariaDB服务配合，配置位于第20行 - 27行。**需要注意**由于请求由Nginx转发过来，需要注意配置Nginx的转发内容，以确保客户端ip的真实有效。

```yaml
publishing:
  server: "127.0.0.1:9000"
redis:
  local:
    server: "127.0.0.1:6379"
    pwd: 
    db: 0
  audit_auth_token:   
    server: "127.0.0.1:6379" 
    pwd:
    db: 0
mariadb:
  tcms:
    server: "172.16.1.250"
    port: 3306
    uid: root
    pwd: "123"
    db: tinycms
    timeout: 10
  logclientip:
    server: "172.16.1.250"
    port: 3306
    uid: root
    pwd: "123"
    db: clientiplog
    table: tcms
    timeout: 10    
audit:
  auth_token:  # 审计auth_token策略
    allow: true # 是否允许审计auth_token
    expire_ms: 60000  # token保存有效时间（毫秒）
    jam_ms: 10000  # 请求被审计后，每次请求都会被卡住jam_ms毫秒数，然后才能继续
    repeat: 1 # token被重复n次后会被触发审计
  log_clientip:
    allow: true 
```

## 
