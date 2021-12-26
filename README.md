# environ

Lightweight Go library to use in 12-factor apps.

```go
s, err := environ.E("FOO").AsString()
s, err := environ.E("FOO").Default("foo_value").AsString()

i, err := environ.E("FOO").AsInt()
i, err := environ.E("FOO").Default(42).AsInt()

f, err := environ.E("FOO").AsFloat()
f, err := environ.E("FOO").Default(4.2).AsFloat()

d, err := environ.E("FOO").AsDuration()
d, err := environ.E("FOO").Default(3 * time.Second).AsDuration()
```
