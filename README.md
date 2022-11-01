# ptrutil

> Package ptrutil provides helpers to convert values to pointers and pointers values

## Install

```sh
go get github.com/verify-lab/ptrutil
```

## Ptr function example

```go
values := []interface{}{
    "test", 1, true,
}

for _, val := range values {
    fmt.Printf("val: %+v; ptr: %+v\n", val, Ptr(val))
}

// Output:
// val: test; ptr: 0xc000014250
// val: 1;    ptr: 0xc000014260
// val: true; ptr: 0xc000014270
```

## FromPtr function example

```go
val := "test"

fmt.Printf("val: %+v; ptr: %+v\n", val, FromPtr(&val))
// Output:
// val: test; ptr: 0xc000014250
```
