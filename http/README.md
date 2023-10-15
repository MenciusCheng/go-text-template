# http

提供 gen 方法入口的 http 服务。

## 用法

```bash
curl --location --request POST '127.0.0.1:8080/ping'

curl --location --request POST '127.0.0.1:8080/gen' \
--header 'Content-Type: application/json' \
--data-raw '{
    "template": "your name is {{ .name }}, age is {{ .age }}",
    "data": "{\"name\":\"coco\",\"age\":18}"
}'
```
