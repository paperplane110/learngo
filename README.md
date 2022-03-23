# Learn Go

## Week 1

### 1-10 Pointer

- go only has pass-by-value
- could use *Pointer* to achieve pass-by-reference

```go
var a *int  // Pointer to int variable
var b int = 10
a = &b
*a = 5
fmt.Println(b)
```

### 1-11 Array

```go
arr1 := [...]int{1,2,3,4,5}
```

- `[10]int` and `[20]int` are different types
- `func f(arr [5]int)` is pass-by-value
- We won't use Array directly, but using `slice` frequently

### 1-12 Slice

Slice is a view of array

```go
arr := [...]int{1,2,3,4,5,6}
s := arr[1:3]
// s = [2,3]
s[0] = 100
// arr = [1,100,3,4,5,6]
```

Reslice

```go
arr := [...]int{1,2,3,4,5,6}
s1 := arr[:4]
s1 = s1[2:]
```

**Extending slice**

```go
arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
s1 := arr[2:6]
s2 := s1[3:5]
fmt.Println("s1: ", s1)
fmt.Println("s2: ", s2)

// s1:  [2 3 4 5]
// s2:  [5 6]
```

### 1-13 Slice Ops

Pop
```go
s := []int{1,2,3,4,5,6}
// pop from front
s = s[1:]
// pop from back
s = s[:len(s)-1]
// pop from middle e.g. 2
s = append(s[:2], s[3:]...)
```

Append
```go
s := []int{1,2,3,4,5,6}
s1 := append(s, 10)
```

### 1-14 Map

Create map
```go
m := map[string]int {
	"jack": 90,
	"marry": 91,
}

// get value
jackScore := m["jack"]
tomScore, ok := m["tom"]    // tomScore = "", ok = false

// delete value
delet(m, "jack")

// traverse
for k, v := range m {
	fmt.Println(k, v)
}
```