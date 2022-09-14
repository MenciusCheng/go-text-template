# go-text-template

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

## 参考文档

- [template package](https://pkg.go.dev/text/template)
- [[译]Golang template 小抄](https://colobu.com/2019/11/05/Golang-Templates-Cheatsheet/)
