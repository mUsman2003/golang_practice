# Go Backend Developer Learning Notes
*Comprehensive guide for developers with C++/Java experience*

## Table of Contents
1. [Key Differences from C++/Java](#key-differences)
2. [Go Fundamentals](#fundamentals)
3. [Go's Type System](#type-system)
4. [Memory Management](#memory-management)
5. [Concurrency (Go's Superpower)](#concurrency)
6. [Error Handling](#error-handling)
7. [Interfaces and Composition](#interfaces)
8. [Package Management](#packages)
9. [Backend-Specific Topics](#backend-topics)
10. [Best Practices](#best-practices)
11. [Common Pitfalls](#pitfalls)
12. [Performance Optimization](#performance)
13. [Testing in Go](#testing)
14. [Deployment and DevOps](#deployment)

---

## 1. Key Differences from C++/Java {#key-differences}

### Philosophy and Design Principles

**Simplicity Over Features**
```go
// Go deliberately excludes many features
// No generics (until Go 1.18), no inheritance, no operator overloading
// No exceptions, no function overloading, no default parameters

// C++ might have:
// template<typename T> class Container<T> { /* complex template code */ }
// operator+(const MyClass& other) { /* operator overloading */ }

// Go approach - simple and explicit
type Container struct {
    items []interface{}  // Before Go 1.18
    // items []T         // With Go 1.18+ generics
}

func (c *Container) Add(item interface{}) {
    c.items = append(c.items, item)
}
```

**Compilation Speed**
```go
// Go compiles extremely fast compared to C++
// - No header files to parse
// - Dependency graph is explicit
// - Unused imports cause compilation errors
// - Simple syntax requires less parsing

// Example: Large Go projects compile in seconds
// vs C++ projects that might take minutes/hours
```

**Composition Over Inheritance**
```go
// Java/C++ inheritance
// class Dog extends Animal implements Runnable { }

// Go composition - embed structs
type Animal struct {
    Name string
    Age  int
}

func (a Animal) Speak() string {
    return a.Name + " makes a sound"
}

type Dog struct {
    Animal              // Embedded - gets all Animal methods
    Breed string
}

func (d Dog) Bark() string {
    return d.Name + " barks"  // Can access Animal fields directly
}

// Usage
dog := Dog{
    Animal: Animal{Name: "Buddy", Age: 3},
    Breed:  "Labrador",
}
fmt.Println(dog.Speak())  // From Animal
fmt.Println(dog.Bark())   // From Dog
```

**No Hidden Magic**
```go
// C++: Hidden constructors, destructors, copy constructors
// Java: Hidden boxing/unboxing, autoboxing

// Go: Everything is explicit
type Person struct {
    name string
}

// No constructor - use factory function
func NewPerson(name string) *Person {
    return &Person{name: name}
}

// No destructor - garbage collector handles cleanup
// No hidden copying - you see exactly what's copied
```

### Syntax Comparison Table

| Feature | C++/Java | Go |
|---------|----------|-----|
| Variable Declaration | `int x = 5;` | `var x int = 5` or `x := 5` |
| Function Declaration | `int add(int a, int b)` | `func add(a, b int) int` |
| Class/Struct | `class Person { public: ... }` | `type Person struct { ... }` |
| Method | `void Person::getName()` | `func (p Person) getName() string` |
| Error Handling | `try/catch` | `if err != nil { ... }` |
| Null Check | `if (ptr != null)` | `if ptr != nil` |
| Array/Collection | `vector<int> v` | `[]int` (slice) |
| Interface | `interface IWriter` | `type Writer interface` |

---

## 2. Go Fundamentals {#fundamentals}

### Variables and Constants

**Variable Declaration Methods**
```go
package main

import "fmt"

func main() {
    // Method 1: var keyword with type
    var name string = "John"
    var age int = 30
    var salary float64 = 50000.50
    
    // Method 2: var with type inference
    var city = "New York"        // string inferred
    var isEmployed = true        // bool inferred
    var weight = 70.5           // float64 inferred
    
    // Method 3: Short declaration (only inside functions)
    country := "USA"            // Equivalent to var country string = "USA"
    zipCode := 10001           // Equivalent to var zipCode int = 10001
    
    // Multiple variable declaration
    var (
        firstName string = "John"
        lastName  string = "Doe"
        email     string = "john@example.com"
    )
    
    // Multiple assignment
    x, y := 10, 20
    a, b, c := 1, 2.5, "hello"
    
    // Zero values (important concept in Go)
    var defaultInt int        // 0
    var defaultString string  // ""
    var defaultBool bool      // false
    var defaultSlice []int    // nil
    var defaultMap map[string]int // nil
    var defaultPointer *int   // nil
    
    fmt.Printf("Zero values: %d, '%s', %t, %v, %v, %v\n", 
        defaultInt, defaultString, defaultBool, defaultSlice, defaultMap, defaultPointer)
}
```

**Constants**
```go
// Constants are computed at compile time
const (
    // Basic constants
    Pi = 3.14159
    E  = 2.71828
    
    // String constants
    AppName = "MyBackendApp"
    Version = "1.0.0"
    
    // Numeric constants with iota (auto-incrementing)
    StatusPending = iota    // 0
    StatusRunning           // 1
    StatusCompleted         // 2
    StatusFailed           // 3
)

// Typed constants
const MaxUsers int = 1000
const DefaultTimeout time.Duration = 30 * time.Second

// Iota examples
const (
    // File permissions
    Read = 1 << iota    // 1 (binary: 001)
    Write               // 2 (binary: 010)  
    Execute             // 4 (binary: 100)
)

const (
    // Skip values with blank identifier
    _ = iota           // 0 (ignored)
    KB = 1 << (10 * iota)  // 1024
    MB                     // 1048576
    GB                     // 1073741824
)
```

### Data Types Deep Dive

**Numeric Types**
```go
// Signed integers
var i8 int8 = 127              // -128 to 127
var i16 int16 = 32767          // -32,768 to 32,767
var i32 int32 = 2147483647     // -2^31 to 2^31-1
var i64 int64 = 9223372036854775807  // -2^63 to 2^63-1

// Unsigned integers
var ui8 uint8 = 255            // 0 to 255 (same as byte)
var ui16 uint16 = 65535        // 0 to 65,535
var ui32 uint32 = 4294967295   // 0 to 2^32-1
var ui64 uint64 = 18446744073709551615 // 0 to 2^64-1

// Platform-specific integers
var i int = 42       // int32 or int64 depending on platform
var ui uint = 42     // uint32 or uint64 depending on platform

// Floating point
var f32 float32 = 3.14159      // IEEE-754 32-bit
var f64 float64 = 3.14159265359 // IEEE-754 64-bit (default for literals)

// Complex numbers
var c64 complex64 = 1 + 2i
var c128 complex128 = complex(1.5, 2.5) // real=1.5, imag=2.5

// Type aliases
type byte = uint8    // byte is alias for uint8
type rune = int32    // rune is alias for int32 (Unicode code point)

// Conversion between types (explicit only)
var a int = 42
var b float64 = float64(a)    // Must explicitly convert
var c int32 = int32(a)        // Must explicitly convert
```

**Strings and Runes**
```go
import (
    "fmt"
    "unicode/utf8"
)

func stringExamples() {
    // String literals
    str1 := "Hello, World!"                    // Interpreted string
    str2 := `Raw string
    with newlines
    and "quotes"`                              // Raw string literal
    
    // Unicode and UTF-8
    unicode := "Hello, 世界! 🌍"
    fmt.Printf("String: %s\n", unicode)
    fmt.Printf("Length in bytes: %d\n", len(unicode))        // 17 bytes
    fmt.Printf("Length in runes: %d\n", utf8.RuneCountInString(unicode)) // 11 runes
    
    // Rune operations
    for i, r := range unicode {
        fmt.Printf("Index %d: rune %c (Unicode %U)\n", i, r, r)
    }
    
    // String manipulation
    name := "John"
    greeting := "Hello, " + name + "!"         // Concatenation
    
    // String formatting (like printf in C++)
    age := 30
    formatted := fmt.Sprintf("Name: %s, Age: %d", name, age)
    
    // String conversion
    num := 42
    numStr := fmt.Sprintf("%d", num)           // int to string
    // strconv.Itoa(num) is more efficient for simple int conversion
    
    // Strings are immutable
    original := "Hello"
    // original[0] = 'h'  // This would cause compile error
    modified := strings.ToLower(original)      // Creates new string
}
```

**Arrays vs Slices (Critical Distinction)**
```go
import "fmt"

func arrayVsSlice() {
    // ARRAYS - Fixed size, value type
    var arr1 [5]int                    // Array of 5 integers, zero-initialized
    arr2 := [5]int{1, 2, 3, 4, 5}    // Array literal
    arr3 := [...]int{1, 2, 3}        // Compiler counts elements
    
    // Arrays are values - copying creates new array
    arr4 := arr2                      // arr4 is completely separate from arr2
    arr4[0] = 999                     // arr2[0] is still 1
    
    fmt.Printf("arr2: %v, arr4: %v\n", arr2, arr4)
    
    // SLICES - Dynamic arrays, reference type
    var slice1 []int                  // nil slice
    slice2 := []int{1, 2, 3, 4, 5}   // Slice literal
    slice3 := make([]int, 5)          // Slice of length 5, zero-initialized
    slice4 := make([]int, 3, 10)      // Length 3, capacity 10
    
    // Slice operations
    slice2 = append(slice2, 6, 7, 8)  // Append elements
    subSlice := slice2[1:4]           // Elements at index 1, 2, 3
    
    // Slices share underlying array
    slice2[1] = 999
    fmt.Printf("slice2: %v, subSlice: %v\n", slice2, subSlice) // Both show 999
    
    // Slice internals
    fmt.Printf("Length: %d, Capacity: %d\n", len(slice2), cap(slice2))
    
    // Converting array to slice
    arrSlice := arr2[:]               // Slice backed by array
    
    // 2D slices
    matrix := [][]int{
        {1, 2, 3},
        {4, 5, 6},
        {7, 8, 9},
    }
    
    // Efficient slice operations
    // Copy slice
    source := []int{1, 2, 3, 4, 5}
    dest := make([]int, len(source))
    copy(dest, source)
    
    // Remove element at index i
    i := 2
    slice2 = append(slice2[:i], slice2[i+1:]...)
}
```

**Maps (Hash Tables)**
```go
import "fmt"

func mapExamples() {
    // Map creation methods
    var m1 map[string]int              // nil map (cannot be written to)
    m1 = make(map[string]int)          // Initialize empty map
    
    m2 := make(map[string]int)         // Create and initialize
    m3 := map[string]int{              // Map literal
        "apple":  5,
        "banana": 3,
        "orange": 8,
    }
    
    // Map operations
    m2["key1"] = 10                    // Set value
    m2["key2"] = 20
    
    value := m2["key1"]                // Get value
    fmt.Printf("Value: %d\n", value)
    
    // Check if key exists (important pattern)
    if value, exists := m2["key3"]; exists {
        fmt.Printf("key3 exists with value: %d\n", value)
    } else {
        fmt.Println("key3 does not exist")
    }
    
    // Delete key
    delete(m2, "key1")
    
    // Iterate over map
    for key, value := range m3 {
        fmt.Printf("%s: %d\n", key, value)
    }
    
    // Maps are reference types
    m4 := m3                          // m4 points to same underlying map
    m4["apple"] = 999
    fmt.Printf("m3 apple: %d\n", m3["apple"]) // Also 999
    
    // Complex map types
    userPrefs := map[string]map[string]interface{}{
        "user1": {
            "theme": "dark",
            "notifications": true,
            "language": "en",
        },
        "user2": {
            "theme": "light",
            "notifications": false,
            "language": "es",
        },
    }
    
    // Map of slices
    groups := map[string][]string{
        "admins": {"alice", "bob"},
        "users":  {"charlie", "david", "eve"},
    }
    
    // Concurrent map access requires synchronization
    // Use sync.Map for concurrent access or protect with mutex
}
```

### Control Structures

**For Loops (The Only Loop in Go)**
```go
import (
    "fmt"
    "time"
)

func forLoopExamples() {
    // Traditional for loop
    for i := 0; i < 10; i++ {
        fmt.Printf("Iteration: %d\n", i)
    }
    
    // While-style loop
    count := 0
    for count < 5 {
        fmt.Printf("Count: %d\n", count)
        count++
    }
    
    // Infinite loop
    for {
        fmt.Println("This runs forever")
        time.Sleep(1 * time.Second)
        break // Don't actually run forever in this example
    }
    
    // Range over slice
    numbers := []int{10, 20, 30, 40, 50}
    for index, value := range numbers {
        fmt.Printf("Index: %d, Value: %d\n", index, value)
    }
    
    // Range over slice (index only)
    for index := range numbers {
        fmt.Printf("Index: %d\n", index)
    }
    
    // Range over slice (value only)
    for _, value := range numbers {
        fmt.Printf("Value: %d\n", value)
    }
    
    // Range over map
    scores := map[string]int{
        "Alice": 95,
        "Bob":   87,
        "Carol": 92,
    }
    for name, score := range scores {
        fmt.Printf("%s scored %d\n", name, score)
    }
    
    // Range over string (iterates over runes)
    text := "Hello, 世界"
    for index, char := range text {
        fmt.Printf("Index: %d, Char: %c\n", index, char)
    }
    
    // Range over channel
    ch := make(chan int, 3)
    ch <- 1
    ch <- 2
    ch <- 3
    close(ch)
    
    for value := range ch {
        fmt.Printf("Received: %d\n", value)
    }
    
    // Nested loops with labels
    outer:
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            if i == 1 && j == 1 {
                break outer  // Break out of outer loop
            }
            fmt.Printf("i=%d, j=%d\n", i, j)
        }
    }
}
```

**Conditional Statements**
```go
import (
    "fmt"
    "os"
    "strconv"
)

func conditionalExamples() {
    // Basic if statement
    age := 25
    if age >= 18 {
        fmt.Println("Adult")
    }
    
    // If-else
    if age < 13 {
        fmt.Println("Child")
    } else if age < 20 {
        fmt.Println("Teenager") 
    } else {
        fmt.Println("Adult")
    }
    
    // If with initialization (very common pattern)
    if err := doSomething(); err != nil {
        fmt.Printf("Error occurred: %v\n", err)
        return
    }
    
    // Another common pattern
    if value, exists := myMap["key"]; exists {
        fmt.Printf("Found value: %v\n", value)
    }
    
    // Switch statement (no break needed!)
    day := "Monday"
    switch day {
    case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
        fmt.Println("Weekday")
    case "Saturday", "Sunday":
        fmt.Println("Weekend")
    default:
        fmt.Println("Unknown day")
    }
    
    // Switch with expression
    switch hour := time.Now().Hour(); {
    case hour < 12:
        fmt.Println("Good morning")
    case hour < 18:
        fmt.Println("Good afternoon")
    default:
        fmt.Println("Good evening")
    }
    
    // Type switch (for interfaces)
    var value interface{} = 42
    switch v := value.(type) {
    case int:
        fmt.Printf("Integer: %d\n", v)
    case string:
        fmt.Printf("String: %s\n", v)
    case bool:
        fmt.Printf("Boolean: %t\n", v)
    default:
        fmt.Printf("Unknown type: %T\n", v)
    }
    
    // Fallthrough (rarely used)
    switch num := 2; num {
    case 1:
        fmt.Println("One")
        fallthrough
    case 2:
        fmt.Println("Two or after One")
        fallthrough
    case 3:
        fmt.Println("Three or continuation")
    }
}

func doSomething() error {
    // Simulate an operation that might fail
    if _, err := strconv.Atoi("not-a-number"); err != nil {
        return err
    }
    return nil
}
```

---

## 3. Go's Type System {#type-system}

### Structs (Go's "Classes")

**Basic Struct Definition and Usage**
```go
package main

import (
    "fmt"
    "time"
)

// Basic struct definition
type Person struct {
    // Exported fields (start with capital letter)
    FirstName string
    LastName  string  
    Age       int
    Email     string
    
    // Unexported fields (start with lowercase)
    ssn       string
    birthDate time.Time
}

// Constructor pattern (Go doesn't have constructors)
func NewPerson(firstName, lastName, email string, age int) *Person {
    return &Person{
        FirstName: firstName,
        LastName:  lastName,
        Email:     email,
        Age:       age,
        birthDate: time.Now().AddDate(-age, 0, 0), // Approximate birth date
    }
}

// Factory function with validation
func NewPersonWithValidation(firstName, lastName, email string, age int) (*Person, error) {
    if firstName == "" || lastName == "" {
        return nil, fmt.Errorf("first name and last name are required")
    }
    if age < 0 || age > 150 {
        return nil, fmt.Errorf("invalid age: %d", age)
    }
    
    return &Person{
        FirstName: firstName,
        LastName:  lastName,
        Email:     email,
        Age:       age,
    }, nil
}

// Methods on structs
func (p Person) FullName() string {
    return p.FirstName + " " + p.LastName
}

func (p Person) IsAdult() bool {
    return p.Age >= 18
}

// Pointer receiver methods (can modify the struct)
func (p *Person) SetEmail(email string) {
    p.Email = email
}

func (p *Person) HaveBirthday() {
    p.Age++
}

// Value receiver vs pointer receiver
func (p Person) GetAgeValue() int {     // Receives copy of struct
    p.Age = 999  // This won't affect the original struct
    return p.Age
}

func (p *Person) GetAgePointer() int {  // Receives pointer to struct
    // p.Age = 999  // This WOULD affect the original struct
    return p.Age
}

func structExamples() {
    // Creating structs
    var p1 Person                       // Zero value initialization
    fmt.Printf("Zero value: %+v\n", p1)

    p2 := Person{                       // Struct literal
        FirstName: "John",
        LastName:  "Doe",
        Age:       30,
        Email:     "john@example.com",
    }
    
    p3 := Person{"Jane", "Smith", 25, "jane@example.com", "", time.Now()}  // Positional
    
    // Using constructor
    p4 := NewPerson("Alice", "Johnson", "alice@example.com", 28)
    
    // Method calls
    fmt.Printf("Full name: %s\n", p2.FullName())
    fmt.Printf("Is adult: %t\n", p2.IsAdult())
    
    // Pointer receiver methods
    p2.SetEmail("newemail@example.com")
    p2.HaveBirthday()
    fmt.Printf("Updated person: %+v\n", p2)
    
    // Anonymous structs (useful for temporary data structures)
    config := struct {
        Host     string
        Port     int
        UseHTTPS bool
    }{
        Host:     "localhost",
        Port:     8080,
        UseHTTPS: false,
    }
    fmt.Printf("Config: %+v\n", config)
}
```

**Struct Embedding (Composition)**
```go
import (
    "fmt"
    "time"
)

// Base struct
type Animal struct {
    Name     string
    Species  string
    Age      int
    Weight   float64
}

func (a Animal) Eat() string {
    return fmt.Sprintf("%s is eating", a.Name)
}

func (a Animal) Sleep() string {
    return fmt.Sprintf("%s is sleeping", a.Name)
}

func (a *Animal) GrowOlder() {
    a.Age++
}

// Embedded struct (composition)
type Dog struct {
    Animal              // Embedded struct - Dog "has" Animal fields/methods
    Breed   string
    IsGoodBoy bool
}

func (d Dog) Bark() string {
    return fmt.Sprintf("%s says Woof!", d.Name)  // Can access Animal.Name directly
}

func (d Dog) Fetch() string {
    return fmt.Sprintf("%s is fetching the ball", d.Name)
}

// Method with same name as embedded struct method
func (d Dog) Eat() string {
    return fmt.Sprintf("%s is eating dog food", d.Name)  // Overrides Animal.Eat()
}

// Multiple embedding
type Address struct {
    Street  string
    City    string
    Country string
}

func (a Address) FullAddress() string {
    return fmt.Sprintf("%s, %s, %s", a.Street, a.City, a.Country)
}

type Employee struct {
    Person              // Embedded
    Address             // Embedded  
    EmployeeID   string
    Department   string
    Salary       float64
    HireDate     time.Time
}

func (e Employee) GetDetails() string {
    return fmt.Sprintf("Employee: %s, ID: %s, Department: %s", 
        e.FullName(), e.EmployeeID, e.Department)  // Uses Person.FullName()
}

func embeddingExamples() {
    // Create Dog with embedded Animal
    dog := Dog{
        Animal: Animal{
            Name:    "Buddy",
            Species: "Canine",
            Age:     3,
            Weight:  25.5,
        },
        Breed:     "Golden Retriever", 
        IsGoodBoy: true,
    }
    
    // Access embedded fields directly
    fmt.Printf("Dog name: %s\n", dog.Name)        // From Animal
    fmt.Printf("Dog breed: %s\n", dog.Breed)      // From Dog
    
    // Call embedded methods
    fmt.Println(dog.Sleep())      // From Animal
    fmt.Println(dog.Bark())       // From Dog
    fmt.Println(dog.Eat())        // Dog's version, not Animal's
    
    // Call Animal's Eat method explicitly
    fmt.Println(dog.Animal.Eat()) // Explicitly call Animal's method
    
    // Modify through embedded struct
    dog.GrowOlder()               // From Animal
    fmt.Printf("Dog age: %d\n", dog.Age)
    
    // Multiple embedding
    emp := Employee{
        Person: Person{
            FirstName: "Alice",
            LastName:  "Johnson",
            Age:       30,
            Email:     "alice@company.com",
        },
        Address: Address{
            Street:  "123 Main St",
            City:    "New York",
            Country: "USA",
        },
        EmployeeID: "EMP001",
        Department: "Engineering",
        Salary:     75000,
        HireDate:   time.Now(),
    }
    
    // Access methods from both embedded structs
    fmt.Println(emp.GetDetails())     // Uses Person.FullName()
    fmt.Println(emp.FullAddress())    // From Address
    fmt.Printf("Employee age: %d\n", emp.Age)  // From Person
}
```

### Pointers in Go

**Pointer Basics (Simpler than C++)**
```go
import "fmt"

func pointerExamples() {
    // Basic pointer operations
    var x int = 42
    var p *int = &x        // p points to x
    
    fmt.Printf("Value of x: %d\n", x)
    fmt.Printf("Address of x: %p\n", &x)
    fmt.Printf("Value of p: %p\n", p)
    fmt.Printf("Value pointed by p: %d\n", *p)
    
    // Modify through pointer
    *p = 100
    fmt.Printf("New value of x: %d\n", x)  // x is now 100
    
    // Pointer to struct
    person := Person{FirstName: "John", LastName: "Doe", Age: 30}
    personPtr := &person
    
    // Access struct fields through pointer (automatic dereferencing)
    fmt.Printf("Name: %s\n", personPtr.FirstName)  // Go automatically dereferences
    // Equivalent to: (*personPtr).FirstName
    
    // Modify struct through pointer
    personPtr.Age = 31
    fmt.Printf("Updated age: %d\n", person.Age)    // person.Age is now 31
    
    // new() function - allocates memory and returns pointer
    ptrToInt := new(int)               // Allocates int, returns *int
    fmt.Printf("Value: %d\n", *ptrToInt)  // 0 (zero value)
    
    ptrToPerson := new(Person)         // Allocates Person, returns *Person
    ptrToPerson.FirstName = "Jane"
    
    // Nil pointers
    var nilPtr *int
    fmt.Printf("Nil pointer: %v\n", nilPtr)  // <nil>
    
    // Check for nil before dereferencing
    if nilPtr != nil {
        fmt.Printf("Value: %d\n", *nilPtr)
    } else {
        fmt.Println("Pointer is nil")
    }
    
    // Pointers to pointers (less common in Go)
    var pp **int = &p      // Pointer to pointer
    fmt.Printf("Value through double pointer: %d\n", **pp)
}

// Function parameters: value vs pointer
func modifyValue(x int) {
    x = 999  // This won't affect the original variable
}

func modifyPointer(x *int) {
    *x = 999  // This will affect the original variable
}

func modifyStruct(p Person) {
    p.Age = 999  // Won't affect original
}

func modifyStructPointer(p *Person) {
    p.Age = 999  // Will affect original
}

func parameterExamples() {
    num := 42
    fmt.Printf("Before modifyValue: %d\n", num)
    modifyValue(num)
    fmt.Printf("After modifyValue: %d\n", num)  // Still 42
    
    modifyPointer(&num)
    fmt.Printf("After modifyPointer: %d\n", num)  // Now 999
    
    person := Person{FirstName: "John", Age: 30}
    fmt.Printf("Before modifyStruct: %d\n", person.Age)
    modifyStruct(person)
    fmt.Printf("After modifyStruct: %d\n", person.Age)  // Still 30
    
    modifyStructPointer(&person)  
    fmt.Printf("After modifyStructPointer: %d\n", person.Age)  // Now 999
}
```

### Custom Types and Type Aliases

**Type Definitions**
```go
import (
    "fmt"
    "strconv"
    "strings"
)

// Define new types based on existing types
type UserID int
type Email string
type Temperature float64
type StatusCode int

// Constants for custom types
const (
    StatusOK StatusCode = 200
    StatusNotFound StatusCode = 404
    StatusInternalError StatusCode = 500
)

// Methods on custom types
func (uid UserID) String() string {
    return fmt.Sprintf("User-%d", int(uid))
}

func (e Email) IsValid() bool {
    return strings.Contains(string(e), "@") && len(string(e)) > 3
}

func (e Email) Domain() string {
    parts := strings.Split(string(e), "@")
    if len(parts) == 2 {
        return parts[1]
    }
    return ""
}

func (t Temperature) Celsius() float64 {
    return float64(t)
}

func (t Temperature) Fahrenheit() float64 {
    return float64(t)*9/5 + 32
}

func (t Temperature) String() string {
    return fmt.Sprintf("%.2f°C", float64(t))
}

func (sc StatusCode) IsSuccess() bool {
    return sc >= 200 && sc < 300
}

func (sc StatusCode) IsError() bool {
    return sc >= 400
}

// Type aliases (different from type definitions)
type MyString = string  // Alias - same type as string
type MyInt int          // New type - different from int

func customTypeExamples() {
    // Using custom types
    var userID UserID = 12345
    var email Email = "user@example.com"
    var temp Temperature = 25.5
    var status StatusCode = StatusOK
    
    fmt.Printf("User ID: %s\n", userID.String())
    fmt.Printf("Email valid: %t\n", email.IsValid())
    fmt.Printf("Email domain: %s\n", email.Domain())
    fmt.Printf("Temperature: %s (%.2f°F)\n", temp.String(), temp.Fahrenheit())
    fmt.Printf("Status success: %t\n", status.IsSuccess())
    
    // Type conversion required for different types
    var regularInt int = 42
    var myInt MyInt = 42
    
    // regularInt = myInt        // Error: cannot assign MyInt to int
    regularInt = int(myInt)      // OK: explicit conversion
    myInt = MyInt(regularInt)    // OK: explicit conversion
    
    // Type alias works like the original type
    var myStr MyString = "hello"
    var regularStr string = myStr  // OK: no conversion needed
    
    // Slice of custom types
    userIDs := []UserID{1001, 1002, 1003}
    for _, id := range userIDs {
        fmt.Printf("Processing %s\n", id.String())
    }
}

// Struct with custom types
type User struct {
    ID       UserID
    Email    Email
    Name     string
    IsActive bool
}

func (u User) Validate() error {
    if u.ID <= 0 {
        return fmt.Errorf("invalid user ID: %d", u.ID)
    }
    if !u.Email.IsValid() {
        return fmt.Errorf("invalid email: %s", u.Email)
    }
    if u.Name == "" {
        return fmt.Errorf("name is required")
    }
    return nil
}

// Function types
type ProcessFunc func(string) (string, error)
type ValidateFunc func(interface{}) bool
type HandlerFunc func(http.ResponseWriter, *http.Request)

// Using function types
func processData(data string, processor ProcessFunc) (string, error) {
    return processor(data)
}

func stringProcessor(s string) (string, error) {
    if s == "" {
        return "", fmt.Errorf("empty string")
    }
    return strings.ToUpper(s), nil
}

func functionTypeExamples() {
    result, err := processData("hello world", stringProcessor)
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return
    }
    fmt.Printf("Processed: %s\n", result)
    
    // Anonymous function
    result2, _ := processData("go lang", func(s string) (string, error) {
        return strings.ReplaceAll(s, " ", "_"), nil
    })
    fmt.Printf("Processed: %s\n", result2)
}
```

---

## 4. Memory Management {#memory-management}

### Garbage Collection in Go

**Understanding Go's GC**
```go
import (
    "fmt"
    "runtime"
    "time"
)

func memoryManagementExamples() {
    // Go uses automatic garbage collection
    // No manual memory management like C++ (no malloc/free, new/delete)
    // No reference counting like some languages
    
    // Print memory statistics
    var m runtime.MemStats
    runtime.ReadMemStats(&m)
    fmt.Printf("Memory allocated: %d KB\n", m.Alloc/1024)
    fmt.Printf("Total allocations: %d\n", m.TotalAlloc/1024)
    fmt.Printf("System memory: %d KB\n", m.Sys/1024)
    fmt.Printf("Number of GC cycles: %d\n", m.NumGC)
    
    // Force garbage collection (rarely needed in production)
    runtime.GC()
    
    // Create many objects to see GC in action
    for i := 0; i < 100000; i++ {
        _ = make([]byte, 1024)  // Allocate 1KB
    }
    
    runtime.ReadMemStats(&m)
    fmt.Printf("After allocations - Memory: %d KB, GC cycles: %d\n", 
        m.Alloc/1024, m.NumGC)
}

// Stack vs Heap allocation examples
func stackVsHeap() {
    // Go compiler decides where to allocate based on escape analysis
    
    // This will likely be allocated on stack (doesn't escape function)
    localVar := 42
    fmt.Printf("Local var: %d\n", localVar)
    
    // This will be allocated on heap (escapes function via return)
    heapVar := createOnHeap()
    fmt.Printf("Heap var: %d\n", *heapVar)
    
    // Large objects typically go to heap
    largeSlice := make([]int, 1000000)
    fmt.Printf("Large slice length: %d\n", len(largeSlice))
}

func createOnHeap() *int {
    x := 42
    return &x  // x escapes to heap because we return its address
}

// Memory-efficient patterns
func efficientPatterns() {
    // 1. Reuse slices instead of creating new ones
    var buffer []byte
    for i := 0; i < 10; i++ {
        buffer = buffer[:0]  // Reset slice but keep capacity
        buffer = append(buffer, fmt.Sprintf("Item %d", i)...)
        // Process buffer
    }
    
    // 2. Use object pools for frequently allocated objects
    pool := &sync.Pool{
        New: func() interface{} {
            return make([]byte, 1024)
        },
    }
    
    // Get from pool
    buf := pool.Get().([]byte)
    // Use buf...
    // Return to pool
    pool.Put(buf)
    
    // 3. Pre-allocate slices with known capacity
    items := make([]string, 0, 100)  // Length 0, capacity 100
    for i := 0; i < 50; i++ {
        items = append(items, fmt.Sprintf("item%d", i))
    }
    
    // 4. Use strings.Builder for string concatenation
    var sb strings.Builder
    sb.Grow(1000)  // Pre-allocate capacity
    for i := 0; i < 100; i++ {
        sb.WriteString(fmt.Sprintf("part%d ", i))
    }
    result := sb.String()
}
```

### Memory Allocation Patterns

**make() vs new() vs Composite Literals**
```go
import "fmt"

func allocationPatterns() {
    // make() - creates slices, maps, channels (reference types)
    slice1 := make([]int, 5)           // Length 5, capacity 5
    slice2 := make([]int, 5, 10)       // Length 5, capacity 10
    map1 := make(map[string]int)       // Empty map
    map2 := make(map[string]int, 10)   // Empty map with hint capacity
    chan1 := make(chan int)            // Unbuffered channel
    chan2 := make(chan int, 5)         // Buffered channel
    
    // new() - allocates memory, returns pointer (rarely used)
    ptr1 := new(int)                   // *int pointing to zero value
    ptr2 := new(Person)                // *Person pointing to zero value
    
    // Composite literals (preferred for structs)
    person1 := Person{FirstName: "John", Age: 30}      // Value
    person2 := &Person{FirstName: "Jane", Age: 25}     // Pointer
    
    // Which to use when:
    fmt.Printf("Slice: %v\n", slice1)
    fmt.Printf("Map: %v\n", map1)
    fmt.Printf("Channel: %v\n", chan1)
    fmt.Printf("Pointer from new: %v\n", ptr1)
    fmt.Printf("Person value: %v\n", person1)
    fmt.Printf("Person pointer: %v\n", person2)
}

// Memory leaks to avoid
func memoryLeaks() {
    // 1. Goroutine leaks
    // BAD - goroutine never terminates
    go func() {
        for {
            // Do work without exit condition
            time.Sleep(time.Second)
        }
    }()
    
    // GOOD - goroutine with exit condition
    ctx, cancel := context.WithCancel(context.Background())
    go func() {
        for {
            select {
            case <-ctx.Done():
                return
            default:
                // Do work
                time.Sleep(time.Second)
            }
        }
    }()
    // Later: cancel() to stop goroutine
    
    // 2. Slice leaks - keeping reference to large underlying array
    // BAD
    largeSlice := make([]byte, 1000000)
    smallSlice := largeSlice[:5]  // Still references large array
    
    // GOOD - copy when you need small part of large slice
    smallSliceCopy := make([]byte, 5)
    copy(smallSliceCopy, largeSlice[:5])
    largeSlice = nil  // Can be garbage collected
    
    // 3. Map leaks - maps don't shrink automatically
    m := make(map[int][]byte)
    for i := 0; i < 100000; i++ {
        m[i] = make([]byte, 1024)
    }
    // Delete all entries but map still holds memory
    for k := range m {
        delete(m, k)
    }
    // Better: create new map
    m = make(map[int][]byte)
}
```

---

## 5. Concurrency (Go's Superpower) {#concurrency}

### Goroutines Deep Dive

**Goroutine Basics and Lifecycle**
```go
import (
    "fmt"
    "runtime"
    "sync"
    "time"
)

func goroutineBasics() {
    // Get number of CPU cores
    fmt.Printf("Number of CPUs: %d\n", runtime.NumCPU())
    fmt.Printf("Number of goroutines: %d\n", runtime.NumGoroutine())
    
    // Simple goroutine
    go func() {
        fmt.Println("Hello from goroutine")
    }()
    
    // Goroutine with parameters
    message := "Hello World"
    go func(msg string) {
        fmt.Printf("Message: %s\n", msg)
    }(message)
    
    // Multiple goroutines
    for i := 0; i < 5; i++ {
        go func(id int) {
            fmt.Printf("Goroutine %d is running\n", id)
            time.Sleep(time.Millisecond * 100)
            fmt.Printf("Goroutine %d finished\n", id)
        }(i)  // Important: pass i as parameter to avoid closure variable capture
    }
    
    // Wait for goroutines (not ideal - use sync.WaitGroup instead)
    time.Sleep(time.Second)
    
    fmt.Printf("Final number of goroutines: %d\n", runtime.NumGoroutine())
}

// Proper goroutine synchronization with WaitGroup
func goroutineWithWaitGroup() {
    var wg sync.WaitGroup
    numWorkers := 5
    
    for i := 0; i < numWorkers; i++ {
        wg.Add(1)  // Add one goroutine to wait group
        go func(workerID int) {
            defer wg.Done()  // Signal completion when function exits
            
            fmt.Printf("Worker %d starting\n", workerID)
            // Simulate work
            time.Sleep(time.Millisecond * time.Duration(workerID*100))
            fmt.Printf("Worker %d completed\n", workerID)
        }(i)
    }
    
    wg.Wait()  // Wait for all goroutines to complete
    fmt.Println("All workers completed")
}

// Worker pool pattern
func workerPool() {
    const numWorkers = 3
    const numJobs = 10
    
    jobs := make(chan int, numJobs)
    results := make(chan int, numJobs)
    
    // Start workers
    var wg sync.WaitGroup
    for w := 1; w <= numWorkers; w++ {
        wg.Add(1)
        go func(id int, jobs <-chan int, results chan<- int) {
            defer wg.Done()
            for job := range jobs {
                fmt.Printf("Worker %d processing job %d\n", id, job)
                time.Sleep(time.Millisecond * 100)  // Simulate work
         
---