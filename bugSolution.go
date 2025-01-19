```go
func main() {
    var i int = 0
    fmt.Println(i) // Output: 0
    fmt.Println(i == 0) // Output: true
    fmt.Println(i == 0.0) // Output: true
    fmt.Println(i == float64(0)) // Output: true

    var f float64 = 0.0
    fmt.Println(i == int(f)) // explicit conversion to int

    var j int64 = 0
    fmt.Println(i == int(j)) // explicit conversion to int
}
```