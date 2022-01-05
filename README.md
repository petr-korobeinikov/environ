# environ

Lightweight Go library to use in 12-factor apps.

```go
// export FOO="Hello, World!"
s, err := environ.E("FOO").AsString()
s, err := environ.E("FOO").Default("Hello!").AsString()

// int support
// export FOO=42
i, err := environ.E("FOO").AsInt()
i, err := environ.E("FOO").Default(42).AsInt()

// hexadecimal support
// export FOO=0xDEADBEEF
h, err := environ.E("FOO").AsInt()
println(h == 0xDEADBEEF)

// octal support
// export FOO=0755
o, err := environ.E("FOO").AsInt()
println(o == 0755)

// float support
// export FOO=4.2
f, err := environ.E("FOO").AsFloat()
f, err := environ.E("FOO").Default(4.2).AsFloat()

// duration support
// export FOO=42s
d, err := environ.E("FOO").AsDuration()
d, err := environ.E("FOO").Default(3 * time.Second).AsDuration()
```
