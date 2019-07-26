package hook_test

import (
	"fmt"
	"hook"
	"testing"
)

func TestAction_WithoutArgs(t *testing.T) {
	f := hook.Actions{}
	// f.Hello()
	fFunc1 := func(args map[string]interface{}) {
		// t.Log("Running ok without args")
	}
	f.Add("filter1", fFunc1)
	f.Do("filter1")
}

func Demo() {

}

func TestAction_MultipleWithSort(t *testing.T) {
	testVar := "Lucifer"
	a := hook.Actions{}
	func1 := func(args map[string]interface{}) {
		// fmt.Println(args)
		// args["Name"] = "a"
		testVar = "Kevin"
	}
	var actionArg1 = map[string]interface{}{
		"id":       "a1",
		"priority": 1,
	}
	a.Add("action1", func1, actionArg1)

	var args = map[string]interface{}{
		"Name": "Lucifer",
		"Age":  35,
	}
	a.Do("action1", args)

	fmt.Println(args)

	if testVar != "Kevin" {
		t.Error(
			"expected", "Kevin",
			"got", testVar,
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
