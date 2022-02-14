# Numeral Systems

## Task 1: Long Addition Base 10 (100%)

### Definition

Write a program (with tests) that implements "long" arithmetics addition. The program may contain
a `LongAdd` function, which accepts two decimal numbers encoded as strings. The function should
return the sum of `a` and `b`.

```go
func LongAdd(a, b string) string { ... }
```

### Test examples

```go
{"123", "210", "333"},
{"55555", "55", "55610"},
{"98765432109876543210987654321098765432109876543210",
  "22222222222222222222222222222222222222222222222222",
  "120987654332098765433209876543320987654332098765432"},
```

## Extra Task 1: Long Fibonacci (+25%)

### Definition

Write a program that outputs N-th Fibonacci sequence number, where N can be large (500 is large here, since
the sequence grows exponentially). The program may contain the following function:

```go
func LongFibonacci(n int) string { ... }
```

### Test examples

```go
{0, "0"},
{10, "55"},
{500, "139423224561697880139724382870407283950070256587697307264108962948325571622863290691557658876222521294125"},
```

## Extra Task 2: Long Addition Base N (+25%)

### Definition

Write a program that adds two long integers in an arbitrary numeral system. The program may contain
the following function:

```go
func LongAdd(a, b string, base int) string { ... }
```

### Test examples

```go
{"55555", "55", 10, "55610"},
{"123", "2", 5, "130"},
```
