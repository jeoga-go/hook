# hook

import "github.com/kneskgo/hook"

Package hook lets you modify the your own code or function output without actually modifying the underlying code.
There are two types of hooks exist : actions and filters.

An **Action** is a hook that is triggered at specific time when your code is running and lets you take an action. This can include things like creating send email on forgot password function or sending a Tweet when someone publishes a post.

A **Filter** allows you get and modify your returned function data before it is sent to the database or the browser. This can include things like that you have array of data and you want to modify the data before returing it actually.



## Filter Example
``` go
    f := hook.Filters{}
    fFunc1 := func(args *map[string]interface{}) {
        fmt.Println("Func 1 Running")
        (*args)["Name"] = "Kevin"
    }
    f.Add("filter1", fFunc1)

    var func2Arg = map[string]interface{}{
        "id":       "b1",
        "priority": 2,
    }

    fFunc2 := func(args *map[string]interface{}) {
        fmt.Println("Func 2 Running")
        (*args)["Name"] = "Raphel"
    }
    f.Add("filter1", fFunc2, func2Arg)

    fmt.Println(f.List)

    var fa = map[string]interface{}{
        "Name": "Lucifer",
        "Age":  35,
    }
    f.Apply("filter1", &fa)
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
```