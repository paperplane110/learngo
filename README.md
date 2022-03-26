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

### 1-17 Struct and method

1. Declare a `struct`

Name is using CamelCase. 
The first letter is top-case means the var is `Public`, which would be exposed to user.
The first letter is lowercase means that the var is `private`.

```go
type Node struct {
	Value       int
	Left, Right *Node
}
```

2. Add some methods

Use `func` to add some methods to a `struct`.

```go
// pass by reference: 
// 传入的是指针，内部操作会影响传入的指针所指向的变量/对象
func (node *Node) SetValue(value int) {
	&node.value = value
}

// pass by value: 
// 传入的是值，相当于把对象/变量拷贝一份，然后传入方法，内部改变不会影响外部的对象/变量
func (node Node) Print() {
    fmt.Println(node.value)
}
```

### 1-18 Package

1. One folder, one package
2. The methods could be defined in different go files (in the same folder)
3. `Public` var/method starts with `TopCase`
4. `Private` var/method starts with `lowerCase`

### 1-19 Extent the package

1. 组合的方式
2. 别名的方式

### 1-20 Use embedding to extend the package

...

### 1-21 Package management

#### History

1. GOPATH
2. GOVENDOR
3. go mod (recommended)

- Use `import` directly in the code to add a new package.
- or `go get -u <package>`: install a package (the newest version)
- or `go get -u <package>@<version>`: specify the version
- `go mod tidy`: clean go.sum
- `go build ./...`: build all go file and pull relative packages


## Week 2

### 2-1 Interface

When to use interface: decouple the code, only remain the logic. 
Especially in Strong type language.

```go
// Indicates the "retriever" has Get method.
// Users don't need to worry about what the specific type of the retriever is.
type retriever interface {
	Get(string) string
	Delete(string) string
}
```

### 2-2 Interface & Duck typing

> ***Duck typing***: 描述外部行为，而非内部结构