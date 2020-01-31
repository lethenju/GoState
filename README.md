# State Machine Framework
A simple framework for designing state machines in Go

## What is it ?

You like to design state machines but you dont want to waste time with the implementation ? 
This project aim to answer your needs and provide an easy way to design your state machines based programs very easily.

## Cool ! How can I use it ?

You need to declare your state machine first before playing with it, of course

### Declaring State Machines

There is 2 ways of declaring your state machine.
* Statically in the code
* Via CSV files

#### Declaring my State Machine Statically

You can simply declare your state machine statically. In this perspective, you need to create and populate State structures.

```go
/** Represents a State.
  */
type State struct {
	Name string               // Name of the State
	Core_function func()      // user-defined Function that gets executed when you enter that state. 
	Connected    []Connection // List of connections from that state
}
```

You can do it like so : 

```go
var my_state sm.State
my_state.Name = "My State"
my_state.Core_function = func() {
	fmt.Println("Hey !")
}
my_state.Connected   = append(my_state.Connected, 
    sm.Connection{ Connection_state : &another_state,
        Reason_to_move : func () bool { 
                return true;
            },
            Transition : func () {
                fmt.Println("[Mine] -> [Another]")
            }},
```
It is as simple as that !

The `core function` will be called when entering to the state.
The `Reason to move` function will be tested for the transition to happen. If it returns true, the transition will occur, firing the `transition function` too.

#### Declaring my State Machine via CSV 

The last solution is a bit heavy for fat state machines, and needs to actually modelise directly in code.
It can be very interesting in many ways to put the model elsewhere.

For example, in CSV files ?

*At first I wanted to use plantuml syntax, but I need to have a very specific format to have the callback names needed. Maybe I will try to implement it, later.*

Here is the format of a line of the `states` CSV
```csv
Name of the state,core_function_callback
```

Here is the format of a line of the `transitions` CSV
```csv
Name of a first state,Name of a second state,reason_callback,transition_callback
```

After writing them down, in the code, you need several things :
```go
    // Declaring our void callbacks
	map_functions := map[string] func ()  {
		"core_function_callback": core_function_callback,
		"transition_callback": transition_callback,
		
	}
	// Declaring our boolean-returning callbacks
	map_reasons := map[string] func () bool {
		"reason_callback": reason_callback,
	}
	
	// Launch the parsing, install the model and get a reference to the first state
	first_state := sm.Parse_and_install(map_functions, map_reasons)
```

I know, it is not ideal, but we need to do it that way for now, for the parser to understand the function names..


### Launching State Machine

Structure of the `runtime` Function

The runtime function is the heart of the State Machine. You need this minimal code to make the framework run
```go
func runtime(s *sm.State) { // The parameter is the entry state of the SM.
  for {
	    // Launch the user-defined core function of the state 
		  s.Core_function()
		  // Transition to another state if possible (the state function will take care of everything)
		  s = (*s).State_function()
	}
}
```

And now, in your main, you should have at least a reference to the first state (The entry one).
So you can simply call `runtime(&my_state)`, and thats it !

Any questions ?
Send to julien.letheno@gmail.com, I will be happy to answer them !
