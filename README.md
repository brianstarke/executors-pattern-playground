# executors-pattern-playground
playing around with a pattern for executing stuff...

```
❯ go run cmd/main.go
2021/08/17 21:38:13 checking auth
2021/08/17 21:38:13 creating something Test Something
2021/08/17 21:38:13 with context &context.valueCtx{Context:(*context.emptyCtx)(0x1400010e000), key:"ckuid", val:"USER-AWESOME"}
2021/08/17 21:38:13 result &internal.CreateSomething{JWTAuthorized:internal.JWTAuthorized{}, ID:"mock-0001", Name:"Test Something", CreatedBy:"USER-AWESOME"}
```
