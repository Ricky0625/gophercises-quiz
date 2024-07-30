# Gophercises - Quiz

The writeup of the exercise details: https://github.com/gophercises/quiz

## Packages in this exercise

### `flag`

RTFM: https://pkg.go.dev/flag

This package provides basic command-line flag parsing.

```go
// this declares an integer flag, -n with default value of 8080, stored in the pointer port, with type *int
var port = flag.Int("n", 8080, "port number") // returns a pointer

// don't prefer working with pointer?
var age int
flag.IntVar(&age, "age", 99, "your age") // returns dereferenced pointer

// parses the command line flags from os.Args[1:]
// must be called after all flags are defined and before acessed by the program
flag.Parse()

fmt.Println(*port) // need to dereference it
fmt.Println(age)   // no dereference needed
```

1. Define the flags needed for your program
2. Parse flag using `flag.Parse()`
3. Use the flags to do whatever you want

> Has builtin -h flag which shows all the flag you defined in your program.

Permitted flag syntax:

```text
-flag
--flag
-flag=x
-flag x // works on non-boolean flags only
```

Permitted boolean flags value:

```text
1, 0, t, f, T, F, true, false, TRUE, FALSE, True, False
```

TODO: notes of `encoding/csv`
