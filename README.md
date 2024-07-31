# Gophercises - Quiz

The writeup of the exercise details: [Gophercises Quiz](https://github.com/gophercises/quiz)

## Packages in this Exercise

### `flag`

RTFM: [flag package documentation](https://pkg.go.dev/flag)

This package provides basic command-line flag parsing.

```go
// This declares an integer flag, -n, with a default value of 8080,
// stored in the pointer 'port' of type *int.
var port = flag.Int("n", 8080, "port number") // returns a pointer

// Prefer not to work with pointers?
var age int
flag.IntVar(&age, "age", 99, "your age") // uses a dereferenced pointer

// Parses the command line flags from os.Args[1:].
// Must be called after all flags are defined and before they are accessed by the program.
flag.Parse()

fmt.Println(*port) // need to dereference it
fmt.Println(age)   // no dereference needed
```

1. Define the flags needed for your program.
2. Parse the flags using `flag.Parse()`.
3. Use the flags to perform the required actions.

> The `flag` package includes a built-in `-h` flag that shows all the flags you defined in your program.

Permitted flag syntax:

```text
-flag
--flag
-flag=x
-flag x // works on non-boolean flags only
```

Permitted boolean flag values:

```text
1, 0, t, f, T, F, true, false, TRUE, FALSE, True, False
```

### `encoding/csv`

RTFM: [`encoding/csv` docs](https://pkg.go.dev/encoding/csv)

This package provides functions to read and write CSV files. This exercise mainly focus on how to read. It supports reading from and writing to CSV files with customizable delimeters, quote characters, and comment markers.

```go
// creates a new CSV reader
csvReader = csv.NewReader(fileObj)

// creates a new CSV writer
csvWriter = csv.NewWriter(fileObj)

// READ
// reads one record, returns a slice representing the fields in the record
csvReader.Read()
// reads all records from CSV file, returns a slice of records, each record is a slice of strings
csvReader.ReadAll()
// reads the NEXT line from CSV file, returns a slice representing the fields in the record
csvReader.ReadLine()

// WRITE
// write a single record to the CSV file, takes in a slice of fields
csvWriter.Write()
// write all, takes in a slice of records, each slice is a slice of strings
csvWriter.WriteAll()

// CONFIG
csvReader.Comma = ';'
csvReader.Comment = '#'
```

