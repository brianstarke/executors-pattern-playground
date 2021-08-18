# executors-pattern-playground

playing around with a pattern for executing stuff...

```
❯ go run cmd/main.go
2021-08-17T22:16:10.283-0400    INFO    internal/executable.go:25       checking auth
2021-08-17T22:16:10.284-0400    INFO    internal/executable.go:53       publishing event        {"topic": "actions.createFoo.mock-exec-id-13029381.start", "payload": {"Bar":"Apple","Baz":"Christopher","Result":null}}
2021-08-17T22:16:10.284-0400    INFO    internal/executable.go:71       publishing event        {"topic": "actions.createFoo.mock-exec-id-13029381.complete", "payload": {"Bar":"Apple","Baz":"Christopher","Result":{"ID":"mock-001","Bar":"Apple","Baz":"Christopher","CreatedBy":"USER-AWESOME"}}}
```
