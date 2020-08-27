#### mod_init

##### Concurrently initializes global structs in designated modules.

-----------------------------------------------

*Motivation:* There are cases where use of a module depends on a globally defined and initialized struct within that module. A pattern to ensure such a global struct would be be available in such a module would be to define the global struct as being assigned the output of an initialization function. When a project pulls in numerous modules following this pattern, there can be performance lag in startup of the program since each assignment is blocking. The `mod_init` module alleviates this issue by initializing each global struct within it's own respective thread.

*Example use:*

```
package example

import (
	mi "github.com/georgercarder/mod_init"
)

func G_ExampleGlobal() *MyExampleStruct {
	e, err := modInitialzer.Get()
	if err != nil {
		// handle err 
	}
	return e.(*MyExampleStruct)	
}

var modInitializer = mi.NewModInit(
		newMyExampleStruct, 
		2 * time.Second, 
		fmt.Errorf("error I want displayed " +
			"if timeout exceeded."))

func newMyExampleStruct() interface{} {
	s := new(MyExampleStruct)
	// do things to initialize
	// eg .. network calls
	//    .. file io
	//    .. etc.
	return s
}
```

```
package calling_in_example

import (
	"example_path_to/example"
)

func myFuncUsingExampleGlobal() (o Output) {
	eg := example.G_ExampleGlobal().MemberFnOfExample()
	// ...
	// ...
}
```

##### Please let me know in the git issues if you have questions or comments.
