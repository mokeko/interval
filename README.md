# interval

[![Go Reference](https://pkg.go.dev/badge/github.com/mokeko/interval.svg)](https://pkg.go.dev/github.com/mokeko/interval)
[![codecov](https://codecov.io/gh/mokeko/interval/branch/main/graph/badge.svg?token=G8EH1PG5V4)](https://codecov.io/gh/mokeko/interval)
[![Go Report Card](https://goreportcard.com/badge/github.com/mokeko/interval)](https://goreportcard.com/report/github.com/mokeko/interval)

Generic interval types and operations for Go.

# Installation

`go get github.com/mokeko/interval`

# Usage

- Int

```go
type Int = interval.Int // wrapper type for int

// [1, 3]
i := interval.New(
  interval.ClosedEp(Int(1)),
  interval.ClosedEp(Int(3)),
)

fmt.Println(i.Contains(Int(2))) // true

// (3, +inf)
i2 := interval.New(
  interval.OpenEp(Int(3)),
  interval.UnboundedEp[Int](),
)

fmt.Println(i.Overlaps(i2)) // false
```
- Time
```go
type Time = interval.Time // wrapper type for time.Time

// [2020-01-01, 2020-01-04]
i := interval.New(
  interval.ClosedEp(Time(time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC))),
  interval.ClosedEp(Time(time.Date(2020, 1, 4, 0, 0, 0, 0, time.UTC))),
)

fmt.Println(i.Contains(Time(time.Date(2020, 1, 5, 0, 0, 0, 0, time.UTC)))) // false

// [2020-01-02, 2020-01-03]
i2 := interval.New(
  interval.ClosedEp(Time(time.Date(2020, 1, 2, 0, 0, 0, 0, time.UTC))),
  interval.ClosedEp(Time(time.Date(2020, 1, 3, 0, 0, 0, 0, time.UTC))),
)

fmt.Println(i.Overlaps(i2)) // true
```
- Other types
```go
// Need to implement Ordered interface.
type ver struct {
  major int
  minor int
}

func (v ver) Equal(v2 ver) bool {
  return v == v2
}

func (v ver) LessThan(v2 ver) bool {
  return v.major < v2.major || (v.major == v2.major && v.minor < v2.minor)
}
```
```go
// [1.0, 2.0)
i := interval.New(
  interval.ClosedEp(ver{1, 0}),
  interval.OpenEp(ver{2, 0}),
)

fmt.Println(i.Contains(ver{2, 0})) // false

// [1.5, 2.5)
i2 := interval.New(
  interval.ClosedEp(ver{1, 5}),
  interval.OpenEp(ver{2, 5}),
)

fmt.Println(i.Overlaps(i2)) // true
```
