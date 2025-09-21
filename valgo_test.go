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

func ExampleInCell() {
	// Example: Validating a slice of primitive values (strings)
	tags := []string{"", "important", "urgent", ""}

	val := Is(String("Project", "name").Not().Blank())

	// Validate each tag in the slice using InCell
	for i, tag := range tags {
		val.InCell("tags", i, Is(String(tag, "tag").Not().Blank()))
	}

	if !val.Valid() {
		// NOTE: sortedErrorMarshalForDocs is an optional parameter used here for
		// documentation purposes to ensure the order of keys in the JSON output.
		out, _ := json.MarshalIndent(val.Error(sortedErrorMarshalForDocs), "", "  ")
		fmt.Println(string(out))
	}

	// output: {
	//   "tags[0]": [
	//     "Tag can't be blank"
	//   ],
	//   "tags[3]": [
	//     "Tag can't be blank"
	//   ]
	// }
}

func ExampleIf() {
	mustBeAdmin := true
	val := Is(String("", "username").Not().Blank()).
		If(mustBeAdmin, Is(String("staff", "role").EqualTo("admin")))

	if !val.Valid() {
		// NOTE: sortedErrorMarshalForDocs is an optional parameter used here for
		// documentation purposes to ensure the order of keys in the JSON output.
		out, _ := json.MarshalIndent(val.Error(sortedErrorMarshalForDocs), "", "  ")
		fmt.Println(string(out))
	}

	// output: {
	//   "role": [
	//     "Role must be equal to \"admin\""
	//   ],
	//   "username": [
	//     "Username can't be blank"
	//   ]
	// }
}

func ExampleDo() {
	mustBeAdmin := true
	val := Is(String("", "username").Not().Blank()).
		Do(func(val *Validation) {
			if mustBeAdmin {
				val.Is(String("staff", "role").EqualTo("admin"))
			}
		})

	if !val.Valid() {
		// NOTE: sortedErrorMarshalForDocs is an optional parameter used here for
		// documentation purposes to ensure the order of keys in the JSON output.
		out, _ := json.MarshalIndent(val.Error(sortedErrorMarshalForDocs), "", "  ")
		fmt.Println(string(out))
	}

	// output: {
	//   "role": [
	//     "Role must be equal to \"admin\""
	//   ],
	//   "username": [
	//     "Username can't be blank"
	//   ]
	// }
}

func ExampleWhen() {
	mustBeAdmin := true
	val := Is(String("", "username").Not().Blank()).
		When(mustBeAdmin, func(val *Validation) {
			val.Is(String("staff", "role").EqualTo("admin"))
		})

	if !val.Valid() {
		// NOTE: sortedErrorMarshalForDocs is an optional parameter used here for
		// documentation purposes to ensure the order of keys in the JSON output.
		out, _ := json.MarshalIndent(val.Error(sortedErrorMarshalForDocs), "", "  ")
		fmt.Println(string(out))
	}

	// output: {
	//   "role": [
	//     "Role must be equal to \"admin\""
	//   ],
	//   "username": [
	//     "Username can't be blank"
	//   ]
	// }
}
