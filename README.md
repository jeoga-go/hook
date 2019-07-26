
# hook  [![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)]

	import "github.com/kneskgo/hook"

Package hook lets you modify the your own code or function output without actually modifying the underlying code.
There are two types of hooks exist : actions and filters.

An **Action** is a hook that is triggered at specific time when your code is running and lets you take an action. This can include things like creating send email on forgot password function or sending a Tweet when someone publishes a post.

A **Filter** allows you get and modify your returned function data before it is sent to the database or the browser. This can include things like that you have array of data and you want to modify the data before returing it actually.

For API docs and examples, see https://godoc.org/github.com/kneskgo/hook

## Filter Example
``` go
f := hook.Filters{}
fFunc1 := func(args map[string]interface{}) {
  fmt.Println("Func 1 Running")
  args["Name"] = "Kevin"
}
f.Add("filter1", fFunc1)

var func2Arg = map[string]interface{}{
  "id":       "b1",
  "priority": 2,
}

fFunc2 := func(args map[string]interface{}) {
  fmt.Println("Func 2 Running")
  args["Name"] = "Raphel"
}
f.Add("filter1", fFunc2, func2Arg)

fmt.Println(f.List)

var fa = map[string]interface{}{
  "Name": "Lucifer",
  "Age":  35,
}
f.Apply("filter1", fa)
fmt.Println(fa)

/* 
outputs:- 
  Func 2 Running
  Func 1 Running
  map[Age:35 Name:Kevin]
*/
```
Note: func2 ran before because we did set priority to 2. Default priority is 10


### Other Filters Functions
``` go
f.RemoveById("b1") // Will remove the filter with ID b1
f.RemoveByTag("filter1") // will remove all the filter with tag name filter1
f.RemoveAll() // will remove all the filters
f.List // will list all the filters
```


## Action Example
``` go
a := hook.Actions{}
func1 := func(args map[string]interface{}) {
	fmt.Println(args)
	fmt.Println("Running Func1")
}
var actionArg1 = map[string]interface{}{
	"id":       "a1",
	"priority": 1,
}
a.Add("action1", func1, actionArg1)

func2 := func(args map[string]interface{}) {
	fmt.Println("Running Func2")
}
a.Add("action1", func2)

var args = map[string]interface{}{
		"Name": "Lucian",
		"Age":  26,
}
a.Do("action1", args)
/* 
outputs:- 
  map[Age:26 Name:Lucian]
  Func 1 Running
  Func 2 Running
*/
```
Note: func2 ran before because we did set priority to 2. Default priority is 10


### Other Actions Functions
``` go
a.RemoveByID("a1") // will remove the action with ID a1
a.RemoveByTag("action1") // will remove all the actions with tag name action1
a.RemoveAll() // will remove all the actions
a.List will list all the actions
```