# expecterlint

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/d0ubletr0uble/expecterlint)
[![Go Report Card](https://goreportcard.com/badge/github.com/d0ubletr0uble/expecterlint)](https://goreportcard.com/report/github.com/d0ubletr0uble/expecterlint)
[![Go Doc](https://godoc.org/github.com/d0ubletr0uble/expecterlint?status.svg)](https://godoc.org/github.com/d0ubletr0uble/expecterlint)
![Pipeline](https://github.com/d0ubletr0uble/expecterlint/actions/workflows/test.yml/badge.svg)

Since `v2.10.0` [mockery](https://github.com/vektra/mockery) introduced [Expecter Structs](https://vektra.github.io/mockery/latest/features/#expecter-structs). 

expecterlint checks if calls to `.On("Method")` could be replaced with syntax `.EXPECT().Method()`.

For example tests that register mock calls with `.On` 

```go
func Test_CreateUser(t *testing.T) {
    u := NewMockUserIFace(t)
    u.On("CreateUser", mock.Anything, User{}).Return(nil)

    err := u.CreateUser(t.Context(), User{})
    if err != nil {
        t.Error(err)
    }
}
```

could be replaced with:

```go
func Test_CreateUser(t *testing.T) {
    u := NewMockUserIFace(t)
    u.EXPECT().CreateUser(mock.Anything, User{}).Return(nil)

    err := u.CreateUser(t.Context(), User{})
    if err != nil {
        t.Error(err)
    }
}
```

which benefits from argument hints, type safety and better IDE support.

## Usage
expecterlint checks only `_test.go` files.

* To automatically replace findings
```shell
go run github.com/d0ubletr0uble/expecterlint@latest -fix ./...
```
* To only list findings
```shell
go run github.com/d0ubletr0uble/expecterlint@latest ./...
```
