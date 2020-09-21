# API接口文档

## 环境说明

### 服务器地址
```bash
# local
http://127.0.0.1:8880
```

### 返回格式

- errCode 0 代表成功 非0表示出错
- errMsg 错误信息
- result 返回的数据

```json
{
  "errCode": 0,
  "errMsg": "",
  "result": {}
}
```


## 数据结构

user

- user_id
- email
- password
- nickname


## 注册与登录

### POST /api/emailcode/send 获取验证码

```bash
curl -s \
-X POST \
-H "Accept: application/json" \
-H "Content-type: application/json" \
-d '{"email":"dummy@dummy.com"}'  \
http://127.0.0.1:8880/api/emailcode/send
```

### POST /api/user/register 注册

```bash
curl -s \
-X POST \
-H "Accept: application/json" \
-H "Content-type: application/json" \
-d '{"email":"dummy@dummy.com", "password":"Dummy123", "code": "7582", "nickname":"dummy"}'  \
http://127.0.0.1:8880/api/user/register
```

### POST /api/user/login 登录
      
```bash
curl -i \
-X POST \
-H "Accept: application/json" \
-H "Content-type: application/json" \
-d '{"email":"dummy@dummy.com", "password":"Dummy123"}'  \
http://127.0.0.1:8880/api/user/login
```