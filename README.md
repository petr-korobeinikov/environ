# environ

Lightweight Go library to use in 12-factor apps.

```go
s, err := environ.E("FOO").AsString()
s, err := environ.E("FOO").Default("foo_value").AsString()

// int support
i, err := environ.E("FOO").AsInt()
i, err := environ.E("FOO").Default(42).AsInt()

// hexadecimal support
os.Setenv("FOO", "0xDEADBEEF")
h, err := environ.E("FOO").AsInt()
println(h == 0xDEADBEEF)

// octal support
os.Setenv("FOO", "0755")
o, err := environ.E("FOO").AsInt()
println(o == 0755)

// float support
f, err := environ.E("FOO").AsFloat()
f, err := environ.E("FOO").Default(4.2).AsFloat()

// duration support
d, err := environ.E("FOO").AsDuration()
d, err := environ.E("FOO").Default(3 * time.Second).AsDuration()
```
