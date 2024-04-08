package valgo

import (
	"encoding/json"
	"fmt"
)

func Example() {
	val := Is(String("Bob", "full_name").Not().Blank().OfLengthBetween(4, 20)).
		Is(Number(17, "age").GreaterThan(18))

	if !val.Valid() {
		// NOTE: sortedErrorMarshalForDocs is an optional parameter used here for
		// documentation purposes to ensure the order of keys in the JSON output.
		out, _ := json.MarshalIndent(val.Error(sortedErrorMarshalForDocs), "", "  ")
		fmt.Println(string(out))
	}
	// Output: {
	//   "age": [
	//     "Age must be greater than \"18\""
	//   ],
	//   "full_name": [
	//     "Full name must have a length between \"4\" and \"20\""
	//   ]
	// }
}

func ExampleIs() {

	val := Is(String("Bob", "full_name").Not().Blank().OfLengthBetween(4, 20)).
		Is(Number(17, "age").GreaterThan(18)).
		Is(String("singl", "status").InSlice([]string{"married", "single"}))

	if !val.Valid() {
		// NOTE: sortedErrorMarshalForDocs is an optional parameter used here for
		// documentation purposes to ensure the order of keys in the JSON output.
		out, _ := json.MarshalIndent(val.Error(sortedErrorMarshalForDocs), "", "  ")
		fmt.Println(string(out))
	}

	// Output: {
	//   "age": [
	//     "Age must be greater than \"18\""
	//   ],
	//   "full_name": [
	//     "Full name must have a length between \"4\" and \"20\""
	//   ],
	//   "status": [
	//     "Status is not valid"
	//   ]
	// }

}

func ExampleNew() {

	month := 5
	monthDay := 11

	val := New()

	if month == 6 {
		val.Is(Number(monthDay, "month_day").LessOrEqualTo(10))
	}

	if val.Valid() {
		fmt.Println("The validation passes")
	}

	// Output: The validation passes
}

func ExampleIn() {
	type Address struct {
		Name   string
		Street string
	}

	type Person struct {
		Name    string
		Address Address
	}

	p := Person{"Bob", Address{"", "1600 Amphitheatre Pkwy"}}

	val := Is(String(p.Name, "name").OfLengthBetween(4, 20)).
		In("address", Is(
			String(p.Address.Name, "name").Not().Blank()).Is(
			String(p.Address.Street, "street").Not().Blank()))

	if !val.Valid() {
		// NOTE: sortedErrorMarshalForDocs is an optional parameter used here for
		// documentation purposes to ensure the order of keys in the JSON output.
		out, _ := json.MarshalIndent(val.Error(sortedErrorMarshalForDocs), "", "  ")
		fmt.Println(string(out))
	}

	// output: {
	//   "address.name": [
	//     "Name can't be blank"
	//   ],
	//   "name": [
	//     "Name must have a length between \"4\" and \"20\""
	//   ]
	// }
}

func ExampleInRow() {
	type Address struct {
		Name   string
		Street string
	}

	type Person struct {
		Name      string
		Addresses []Address
	}

	p := Person{
		"Bob",
		[]Address{
			{"", "1600 Amphitheatre Pkwy"},
			{"Home", ""},
		},
	}

	val := Is(String(p.Name, "name").OfLengthBetween(4, 20))

	for i, a := range p.Addresses {
		val.InRow("addresses", i, Is(
			String(a.Name, "name").Not().Blank()).Is(
			String(a.Street, "street").Not().Blank()))
	}

	if !val.Valid() {
		// NOTE: sortedErrorMarshalForDocs is an optional parameter used here for
		// documentation purposes to ensure the order of keys in the JSON output.
		out, _ := json.MarshalIndent(val.Error(sortedErrorMarshalForDocs), "", "  ")
		fmt.Println(string(out))
	}

	// output: {
	//   "addresses[0].name": [
	//     "Name can't be blank"
	//   ],
	//   "addresses[1].street": [
	//     "Street can't be blank"
	//   ],
	//   "name": [
	//     "Name must have a length between \"4\" and \"20\""
	//   ]
	// }
}

func ExampleCheck() {
	val := Check(String("", "full_name").Not().Blank().OfLengthBetween(4, 20))

	if !val.Valid() {
		// NOTE: sortedErrorMarshalForDocs is an optional parameter used here for
		// documentation purposes to ensure the order of keys in the JSON output.
		out, _ := json.MarshalIndent(val.Error(sortedErrorMarshalForDocs), "", "  ")
		fmt.Println(string(out))
	}

	// output: {
	//   "full_name": [
	//     "Full name can't be blank",
	//     "Full name must have a length between \"4\" and \"20\""
	//   ]
	// }
}
