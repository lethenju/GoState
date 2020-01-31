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

TODO (for now see fizzbuzz example)

#### Declaring my State Machine via CSV 

TODO (for now see game example)

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

TODO
