package valgo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidatorComparableNot(t *testing.T) {

	v := Is(Comparable(10).Not().EqualTo(11))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorComparableEqualToValid(t *testing.T) {

	var v *Validation

	v = Is(Comparable(10).EqualTo(10))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Test with string
	v = Is(Comparable("hello").EqualTo("hello"))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Test with struct (comparable)
	type Point struct {
		X, Y int
	}
	p1 := Point{X: 1, Y: 2}
	p2 := Point{X: 1, Y: 2}
	v = Is(Comparable(p1).EqualTo(p2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	a := 10
	v = Is(Comparable(&a).EqualTo(&a))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorComparableEqualToInvalid(t *testing.T) {

	var v *Validation

	v = Is(Comparable(11).EqualTo(10))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be equal to \"10\"",
		v.Errors()["value_0"].Messages()[0])

	// Test with string
	v = Is(Comparable("hello").EqualTo("world"))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be equal to \"world\"",
		v.Errors()["value_0"].Messages()[0])

	// Test with struct (comparable)
	type Point struct {
		X, Y int
	}
	p1 := Point{X: 1, Y: 2}
	p2 := Point{X: 1, Y: 3}
	v = Is(Comparable(p1).EqualTo(p2))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be equal to \"{1 3}\"",
		v.Errors()["value_0"].Messages()[0])

	a := 10
	b := 10

	v = Is(Comparable(&a).EqualTo(&b))
	assert.False(t, v.Valid())
	assert.Contains(t,
		v.Errors()["value_0"].Messages()[0],
		"Value 0 must be equal to")
}

func TestValidatorComparablePassingValid(t *testing.T) {

	var v *Validation

	valTen := 10

	v = Is(Comparable(valTen).Passing(func(val int) bool {
		return val == 10
	}))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorComparablePassingInvalid(t *testing.T) {

	var v *Validation

	valTen := 10

	v = Is(Comparable(valTen).Passing(func(val int) bool {
		return val == 9
	}))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 is not valid",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorComparablePassingWithCustomType(t *testing.T) {

	var v *Validation

	type Status string
	status := Status("running")

	// Test valid status
	v = Is(Comparable(status).Passing(func(s Status) bool {
		return s == "running" || s == "paused"
	}))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Test invalid status
	status = Status("stopped")
	v = Is(Comparable(status).Passing(func(s Status) bool {
		return s == "running" || s == "paused"
	}))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 is not valid",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorComparablePassingWithStruct(t *testing.T) {

	var v *Validation

	type User struct {
		Name  string
		Age   int
		Email string
	}

	user := User{Name: "John", Age: 25, Email: "john@example.com"}

	// Test valid user
	v = Is(Comparable(user).Passing(func(u User) bool {
		return u.Age >= 18 && u.Email != ""
	}))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Test invalid user
	user = User{Name: "John", Age: 17, Email: ""}
	v = Is(Comparable(user).Passing(func(u User) bool {
		return u.Age >= 18 && u.Email != ""
	}))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 is not valid",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorComparableOrOperatorWithIs(t *testing.T) {
	var v *Validation

	var _true = true
	var _false = false

	// Testing Or operation with two valid conditions
	v = Is(Comparable(true).EqualTo(true).Or().EqualTo(true))
	assert.True(t, v.Valid())
	assert.Equal(t, _true || true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing Or operation with left invalid and right valid conditions
	v = Is(Comparable(true).EqualTo(false).Or().EqualTo(true))
	assert.True(t, v.Valid())
	assert.Equal(t, false || true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing Or operation with left valid and right invalid conditions
	v = Is(Comparable(true).EqualTo(true).Or().EqualTo(false))
	assert.True(t, v.Valid())
	assert.Equal(t, true || false, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing Or operation with two invalid conditions
	v = Is(Comparable(true).EqualTo(false).Or().EqualTo(false))
	assert.False(t, v.Valid())
	assert.Equal(t, _false || false, v.Valid())
	assert.NotEmpty(t, v.Errors())

	// // Testing And operation (default when no Or() function is used) with left valid and right invalid conditions
	// v = Is(Comparable(true).EqualTo(true).EqualTo(false))
	// assert.False(t, v.Valid())
	// assert.Equal(t, true && false, v.Valid())
	// assert.NotEmpty(t, v.Errors())

	// // Testing combination of Not and Or operators with left valid and right invalid conditions
	// v = Is(Comparable(true).Not().EqualTo(false).Or().EqualTo(false))
	// assert.True(t, v.Valid())
	// assert.Equal(t, !false || false, v.Valid())
	// assert.Empty(t, v.Errors())

	// // Testing combination of Not and Or operators with left invalid and right valid conditions
	// v = Is(Comparable(true).Not().EqualTo(true).Or().EqualTo(true))
	// assert.True(t, v.Valid())
	// assert.Equal(t, !true || true, v.Valid())
	// assert.Empty(t, v.Errors())

	// // Testing multiple Or operations in sequence with the first condition being valid
	// v = Is(Comparable(true).EqualTo(true).Or().EqualTo(false).Or().EqualTo(false))
	// assert.True(t, v.Valid())
	// assert.Equal(t, true || _false || false, v.Valid())
	// assert.Empty(t, v.Errors())

	// // Testing multiple Or operations in sequence with the last condition being valid
	// v = Is(Comparable(true).EqualTo(false).Or().EqualTo(false).Or().EqualTo(true))
	// assert.True(t, v.Valid())
	// assert.Equal(t, _false || false || true, v.Valid())
	// assert.Empty(t, v.Errors())

	// // Testing invalid Or operation then valid And operation
	// v = Is(Comparable(true).EqualTo(false).Or().EqualTo(true).EqualTo(true))
	// assert.True(t, v.Valid())
	// assert.Equal(t, false || _true && true, v.Valid())
	// assert.Empty(t, v.Errors())

	// // Testing valid Or operation then invalid And operation
	// v = Is(Comparable(true).EqualTo(false).Or().EqualTo(true).EqualTo(false))
	// assert.False(t, v.Valid())
	// assert.Equal(t, false || true && false, v.Valid())
	// assert.NotEmpty(t, v.Errors())

	// // Testing valid And operation then invalid Or operation
	// v = Is(Comparable(true).EqualTo(true).EqualTo(true).Or().EqualTo(false))
	// assert.True(t, v.Valid())
	// assert.Equal(t, _true && true || false, v.Valid())
	// assert.Empty(t, v.Errors())

	// // Testing invalid And operation then valid Or operation
	// v = Is(Comparable(true).EqualTo(true).EqualTo(false).Or().EqualTo(true))
	// assert.True(t, v.Valid())
	// assert.Equal(t, true && false || true, v.Valid())
	// assert.Empty(t, v.Errors())
}

func TestValidatorComparableOrOperatorWithCheck(t *testing.T) {
	var v *Validation

	// Check are Non-Short-circuited operations

	var _true = true
	var _false = false

	// Testing Or operation with two valid conditions
	v = Check(Comparable(true).EqualTo(true).Or().EqualTo(true))
	assert.True(t, v.Valid())
	assert.Equal(t, _true || true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing Or operation with left invalid and right valid conditions
	v = Check(Comparable(true).EqualTo(false).Or().EqualTo(true))
	assert.True(t, v.Valid())
	assert.Equal(t, false || true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing Or operation with left valid and right invalid conditions
	v = Check(Comparable(true).EqualTo(true).Or().EqualTo(false))
	assert.True(t, v.Valid())
	assert.Equal(t, true || false, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing Or operation with two invalid conditions
	v = Check(Comparable(true).EqualTo(false).Or().EqualTo(false))
	assert.False(t, v.Valid())
	assert.Equal(t, _false || false, v.Valid())
	assert.NotEmpty(t, v.Errors())

	// Testing And operation (default when no Or() function is used) with left valid and right invalid conditions
	v = Check(Comparable(true).EqualTo(true).EqualTo(false))
	assert.False(t, v.Valid())
	assert.Equal(t, true && false, v.Valid())
	assert.NotEmpty(t, v.Errors())

	// Testing combination of Not and Or operators with left valid and right invalid conditions
	v = Check(Comparable(true).Not().EqualTo(false).Or().EqualTo(false))
	assert.True(t, v.Valid())
	assert.Equal(t, !false || false, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing combination of Not and Or operators with left invalid and right valid conditions
	v = Check(Comparable(true).Not().EqualTo(true).Or().EqualTo(true))
	assert.True(t, v.Valid())
	assert.Equal(t, !true || true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing multiple Or operations in sequence with the first condition being valid
	v = Check(Comparable(true).EqualTo(true).Or().EqualTo(false).Or().EqualTo(false))
	assert.True(t, v.Valid())
	assert.Equal(t, true || _false || false, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing multiple Or operations in sequence with the last condition being valid
	v = Check(Comparable(true).EqualTo(false).Or().EqualTo(false).Or().EqualTo(true))
	assert.True(t, v.Valid())
	assert.Equal(t, _false || false || true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing invalid Or operation then valid And operation
	v = Check(Comparable(true).EqualTo(false).Or().EqualTo(true).EqualTo(true))
	assert.True(t, v.Valid())
	assert.Equal(t, false || _true && true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing valid Or operation then invalid And operation
	v = Check(Comparable(true).EqualTo(false).Or().EqualTo(true).EqualTo(false))
	assert.False(t, v.Valid())
	assert.Equal(t, false || true && false, v.Valid())
	assert.NotEmpty(t, v.Errors())

	// Testing valid And operation then invalid Or operation
	v = Check(Comparable(true).EqualTo(true).EqualTo(true).Or().EqualTo(false))
	assert.True(t, v.Valid())
	assert.Equal(t, _true && true || false, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing invalid And operation then valid Or operation
	v = Check(Comparable(true).EqualTo(true).EqualTo(false).Or().EqualTo(true))
	assert.True(t, v.Valid())
	assert.Equal(t, true && false || true, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorComparableWithString(t *testing.T) {
	var v *Validation

	// Test with string
	text := "hello world"
	v = Is(Comparable(text).Passing(func(s string) bool {
		return len(s) > 0 && s[0] == 'h'
	}))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Test with empty string
	emptyText := ""
	v = Is(Comparable(emptyText).Passing(func(s string) bool {
		return len(s) == 0
	}))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Test EqualTo with string
	v = Is(Comparable("test").EqualTo("test"))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	v = Is(Comparable("test").EqualTo("different"))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be equal to \"different\"",
		v.Errors()["value_0"].Messages()[0])

	// Test with another string comparison
	v = Is(Comparable("hello").EqualTo("hello"))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorComparableWithNumericTypes(t *testing.T) {
	var v *Validation

	// Test with int
	v = Is(Comparable(42).EqualTo(42))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	v = Is(Comparable(42).EqualTo(43))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be equal to \"43\"",
		v.Errors()["value_0"].Messages()[0])

	// Test with float64
	v = Is(Comparable(3.14).EqualTo(3.14))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	v = Is(Comparable(3.14).EqualTo(3.15))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be equal to \"3.15\"",
		v.Errors()["value_0"].Messages()[0])

	// Test with int32
	v = Is(Comparable(int32(100)).EqualTo(int32(100)))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	v = Is(Comparable(int32(100)).EqualTo(int32(101)))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be equal to \"101\"",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorComparableWithCustomStruct(t *testing.T) {
	var v *Validation

	type Person struct {
		Name string
		Age  int
	}

	person := Person{Name: "Alice", Age: 30}

	// Test EqualTo with struct
	person2 := Person{Name: "Alice", Age: 30}
	v = Is(Comparable(person).EqualTo(person2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Test EqualTo with different struct
	person3 := Person{Name: "Bob", Age: 30}
	v = Is(Comparable(person).EqualTo(person3))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be equal to \"{Bob 30}\"",
		v.Errors()["value_0"].Messages()[0])

	// Test Passing with struct
	v = Is(Comparable(person).Passing(func(p Person) bool {
		return p.Age >= 18 && len(p.Name) > 0
	}))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorComparableWithArray(t *testing.T) {
	var v *Validation

	// Test with array (comparable)
	arr1 := [3]int{1, 2, 3}
	arr2 := [3]int{1, 2, 3}
	v = Is(Comparable(arr1).EqualTo(arr2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Test with different array
	arr3 := [3]int{1, 2, 4}
	v = Is(Comparable(arr1).EqualTo(arr3))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be equal to \"[1 2 4]\"",
		v.Errors()["value_0"].Messages()[0])

	// Test Passing with array
	v = Is(Comparable(arr1).Passing(func(a [3]int) bool {
		return a[0] == 1 && a[1] == 2
	}))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorComparableWithComplexType(t *testing.T) {
	var v *Validation

	// Test with complex number
	c1 := complex(1, 2)
	c2 := complex(1, 2)
	v = Is(Comparable(c1).EqualTo(c2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Test with different complex number
	c3 := complex(1, 3)
	v = Is(Comparable(c1).EqualTo(c3))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be equal to \"(1+3i)\"",
		v.Errors()["value_0"].Messages()[0])

	// Test Passing with complex
	v = Is(Comparable(c1).Passing(func(c complex128) bool {
		return real(c) == 1 && imag(c) == 2
	}))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorComparableWithChannel(t *testing.T) {
	var v *Validation

	// Test with channel
	ch1 := make(chan int)
	ch2 := make(chan int)

	// Channels are comparable but different instances are not equal
	v = Is(Comparable(ch1).EqualTo(ch2))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be equal to \"0x",
		v.Errors()["value_0"].Messages()[0][:len("Value 0 must be equal to \"0x")])

	// Test with same channel
	v = Is(Comparable(ch1).EqualTo(ch1))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Test Passing with channel
	v = Is(Comparable(ch1).Passing(func(ch chan int) bool {
		return ch != nil
	}))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorComparableInSliceValid(t *testing.T) {

	var v *Validation

	v = Is(Comparable(10).InSlice([]int{5, 10, 15}))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Test with string
	v = Is(Comparable("hello").InSlice([]string{"world", "hello", "test"}))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Test with float
	v = Is(Comparable(3.14).InSlice([]float64{3.13, 3.14, 3.15}))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Test with struct
	type Point struct {
		X, Y int
	}
	p1 := Point{X: 1, Y: 2}
	p2 := Point{X: 1, Y: 2}
	p3 := Point{X: 2, Y: 3}
	v = Is(Comparable(p1).InSlice([]Point{p2, p3}))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Test with array
	arr1 := [3]int{1, 2, 3}
	arr2 := [3]int{1, 2, 3}
	arr3 := [3]int{2, 3, 4}
	v = Is(Comparable(arr1).InSlice([][3]int{arr2, arr3}))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Test with complex
	c1 := complex(1, 2)
	c2 := complex(1, 2)
	c3 := complex(2, 3)
	v = Is(Comparable(c1).InSlice([]complex128{c2, c3}))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyInt int
	var myInt1 MyInt = 10

	v = Is(Comparable(myInt1).InSlice([]MyInt{5, 10, 15}))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorComparableInSliceInvalid(t *testing.T) {

	var v *Validation

	v = Is(Comparable(10).InSlice([]int{5, 15, 20}))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 is not valid",
		v.Errors()["value_0"].Messages()[0])

	// Test with string
	v = Is(Comparable("hello").InSlice([]string{"world", "test", "example"}))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 is not valid",
		v.Errors()["value_0"].Messages()[0])

	// Test with float
	v = Is(Comparable(3.14).InSlice([]float64{3.13, 3.15, 3.16}))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 is not valid",
		v.Errors()["value_0"].Messages()[0])

	// Test with struct
	type Point struct {
		X, Y int
	}
	p1 := Point{X: 1, Y: 2}
	p2 := Point{X: 2, Y: 3}
	p3 := Point{X: 3, Y: 4}
	v = Is(Comparable(p1).InSlice([]Point{p2, p3}))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 is not valid",
		v.Errors()["value_0"].Messages()[0])

	// Test with array
	arr1 := [3]int{1, 2, 3}
	arr2 := [3]int{2, 3, 4}
	arr3 := [3]int{3, 4, 5}
	v = Is(Comparable(arr1).InSlice([][3]int{arr2, arr3}))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 is not valid",
		v.Errors()["value_0"].Messages()[0])

	// Test with complex
	c1 := complex(1, 2)
	c2 := complex(2, 3)
	c3 := complex(3, 4)
	v = Is(Comparable(c1).InSlice([]complex128{c2, c3}))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 is not valid",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type MyInt int
	var myInt1 MyInt = 10

	v = Is(Comparable(myInt1).InSlice([]MyInt{5, 15, 20}))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 is not valid",
		v.Errors()["value_0"].Messages()[0])
}
