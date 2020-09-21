## HelloService
### HelloService.Say 说你好

- 请求示例

```bash
curl -s \
-X POST \
-H "Accept: application/json" \
-H "Content-type: application/json" \
-H "X-TraceID: hello" \
-d '{"jsonrpc":"2.0", "id": 123, "method": "HelloService.Say", "params":[{"who":"changnian"}]}'  \
http://127.0.0.1:8880/goweb/jsonrpc
```

- 输出示例

```json
{"result":{"Message":"Hello, changnian!"},"error":null,"id":123}
```

