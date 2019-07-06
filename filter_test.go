package hook_test

import (
	"hook"
	"testing"
)

func TestFilter_WithoutArgs(t *testing.T) {
	f := hook.Filters{}
	// f.Hello()
	fFunc1 := func(args *map[string]interface{}) {
		// t.Log("Running ok without orgs")
	}
	f.Add("filter1", fFunc1)
	f.Apply("filter1")
}

func TestFilter_MultipleWithSort(t *testing.T) {
	f := hook.Filters{}

	fFunc1 := func(args *map[string]interface{}) {
		// fmt.Println("Filter 1")
		// fmt.Println(args)
		(*args)["Name"] = "a"
	}
	f.Add("filter1", fFunc1)

	var b = map[string]interface{}{
		"id":       "b1",
		"priority": 2,
	}

	fFunc2 := func(args *map[string]interface{}) {
		// fmt.Println("Filter 2")
		// fmt.Println(args)

		(*args)["Name"] = "b"
	}
	f.Add("filter1", fFunc2, b)

	// f.RemoveById("abc1")
	// fmt.Println(f.List)

	var fa = map[string]interface{}{
		"Name": "Lucifer",
		"Age":  35,
	}
	f.Apply("filter1", &fa)

	if fa["Name"] != "a" {
		t.Error(
			"For", fa,
			"expected", "a",
			"got", fa["Name"],
		)
	}
}

// This is a package-level example:
// func Example() {
// 	fmt.Printf("Hi")
// }

// // Use Real like a general floating type:
// func ExampleFilter() {
// 	fmt.Printf("Real")
// }
