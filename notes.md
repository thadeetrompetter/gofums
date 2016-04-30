# Go fundamentals

## Introduction

### `iota`

assign the value `iota` to a variable to have it increment automatically. You create enums this way.

```go
const (
	foo = iota
	bar
	baz
)

func main() {
	fmt.Printf("%d,%d,%d\n", foo, bar, baz)
}
```

### variables

`Printf` noteworthy:

```go
fmt.Printf("2 decimal places: %f.2\n", 3.4567);
fmt.Printf("print a boolean value: %t\n", true);
fmt.Printf("print a hex value: %x\n", byte(65));

```

### strings

Take a substring:

```go
str := "lorem ipsum golang"
fmt.Printf("Print lorem: %s\n", str[0:5]) // provide start and end index
fmt.Printf("%s\n", str[:5]) // default to beginning of string
fmt.Printf("%s\n", str[6:]) // default to end of string
```

Characters in a string are called **runes** in go. You can loop over the runes
in a string like this

```go
str := "lorem ipsum golang"
for _, r := range str {
	fmt.Printf("%c\n", r) // print each rune on a new line
}
```

Delimiting a string with backticks means that its contents will be taken literally. Special characters will lose their meaning.

### control structures

It's possible to declare a variable in an if statement. The variable will then
be scoped to the if statement's block.

```go
if bytes, err := fmt.Printf("%s\n", "foobar"); err == nil {
	fmt.Printf("%d\n", bytes) // bytes is available inside if, else, else if
}
// but not outside
```

The **switch** statement differs from javascript in that it is not necessary
 to add a `break` after each case that you want not to fall through.
If you want a case to fall through, you need to add `fallthrough` at the end
to get fallthrough behavior.

You can use `switch` on a particular variable, or not. Usage without a
variable following the `switch` keyword is not allowed in javascript.
Inside a case. You can also separate conditions with a comma. You can't rely
on consecutive case fallthrough, hence the comma separation.

```go
foo := "bananas"
switch {
case foo[:1] == "b", foo[:2] == "ba":
	fmt.Println("great success")
default:
	fmt.Println("that's not it")
}
```

### loops

looping in go is done with `for`

```go
// an infinite loop
for {
	fmt.Println("i could do this all day")
}
// a while loop
i := 0
for i < 10 {
	fmt.Println(i);
	i++
}
// a for loop
for i := 0; i < 10; i++ {
	fmt.Println(i)
	// i is scoped to this block
}
```

### function definitions

Create a function with `func`. Group arguments of same type.

You can return multiple values from a `func`

```go
func foo(first, second string, third int) (string, string) {
	fmt.Println(first, second, third)
	return "foo", "bar"
}
```

`defer` can be used to queue code that needs to execute before a function
returns. This way it's easy to keep related code grouped together.

```go
func foo(filename string) string {
	file, err := os.Open(filename)
	if err != nil {
		return "not good"
	}
	defer file.Close()
	// do many things that could possibly make you forget to close the file
	// ...
	// ...
	return "success value"
}
```

`rest` parameter like in es6, but prepended to the type declaration instead
of the argument itself.

```go
func foo(things ...string) {
	for _, thing := range things {
		fmt.Println(thing)
	}
}
foo("one", "two", "three", "four", "five")
```

### arrays and slices

For **arrays**, the length is part of the type. Unlike in javascript, arrays
are always passed as value, unless specifically stated otherwise. Arrays use
dots syntax to specify the array length by passing its values, or a number
specifying the array length.

```go
// length is determined on the fly
thingsArray := [...]string{"foo", "bar", "baz"}
// or beforehand
thingsArray := [3]string{"foo", "bar", "baz"}
```

Slices are used more often for reasons of speed and flexibility.
A slice is like a window into an array. Slices are passed by reference,
unlike arrays. Slices, being pointers, perform better because they're but
pointers to an underlying array or another slice.

```go
thingsSlice := []string{"this", "is", "passed", "by", "reference"}
```

Similar to extracting a substring, you can get a slice of a slice using
bracket notation with the start and end offset, separated by a colon. The
 second index means ***up to but not including***.

```go
numbers := []int{1,2,3,4,5}
fmt.Println(numbers[3:4]) // prints [4]
```

If you want to create a slice without but don't have items to put in it yet,
you use `make` and pass the amount of items that are going to be in it. The
 slice will be created with given the amount of elements set to empty values.

You can also pass an integer for capacity to make as a third argument, which
will create an empty array expecting a certain amount of slots. Capacity is
not a hard limit. You can keep appending items to the slice (with `append`)
and the slice will grow to accommodate (capacity doubles).

```go
numbers := make([]int, 5)
numbers[0] = 1 // [1 0 0 0 0]
// or
numbers := make([]int, 0, 5)
numbers = append(numbers, 1) // [1]
```

`len` and `cap` give you the actual length or the capacity for a slice.

Trying array access notation for an index out of bounds will throw a runtime
error.

`copy` will make a fresh copy of a slice. The copied slice will no longer
reference the memory of the original slice.

### maps

Similar to a plain object in javascript (even more to Map in es6). To create
a new map, you also use make.

```go
myMap := make(map[string]int)
myMap["MeaningOfLife"] = 42
```

If you use a key that doesn't exist, you'll get the zero value for the value
type that is set.

Actually, trying to get the value for a non existent key will return both a
zero value and a falsy boolean, which specifies that the lookup failed
because you used a bad key.

```go
if val, ok := myMap["MeaningOfLife"]; ok {
	fmt.Println(val)
} else {
	fmt.Println("key does not exist")
}
```

Use `delete` to remove a key:value pair from a map

```go
delete(myMap, "MeaningOfLife")
```

Any type, except a slice, can be the key in a key:value pair on a map.

### byte slices

a byteslice is a slice of bytes. Like a buffer in nodejs.
When you do any sort of **io**, you will likely use byteslices. In this
example, the contents of some-file are read into a byteslice.

```go
const file string = "/home/user/dir/some-file"
f, err := os.Open(file)
if err != nil {
	fmt.Println("could not open file")
	os.Exit(1)
}
defer f.Close()

b := make([]byte, 100)

n, _ := f.Read(b)

// convert byteslice to a string.
fromBytes := string(b)
fmt.Printf("%d, % x: %s\n", n, b, fromBytes)
```

Similar to calling `toString` on a buffer in node.

If you want to write some string to a file (os.Write), you need to
convert the string to a byteslice.

```go
someString := "hi there"
fromString := []byte(someString) // ready for io!
```

### errors

Errors in go are generally returned by functions in `result, error` style.
Which reminds of nodejs error handling in callback functions.

When you have a simple error to return that you don't need to refer to
elsewhere, use `fmt.Errorf`.

```go
func errFunc(str string) error {
	if str == "foobar" {
		return fmt.Errorf("foobar is not allowed")
	}
	_, err := fmt.Printf("%s\n", str)
	return err
}

func main() {
	if err := errFunc("foobar"); err != nil {
		fmt.Printf("could not print it, because: %s\n", err)
		os.Exit(1)
	}
}
```

When you need to provide more elaborate errors, for example, when you're
writing packages, you'll have to make use of the **errors** package.

Instead of returning `fmt.Errorf` from errorFunc, make a new error and
return that instead. The error's var name is idiomatic go.

```go
var errFoobar = errors.New("foobar is not allowed")

// in errorFunc
if str == "foobar" {
	return errFoobar
}
// you can use the returned error in comparison
if(err == errFoobar){
	// do something specific
} else {
	// do something else
}
```

### goroutines and channels

A **goroutine** is a separately running process or thread, executing
concurrently with the rest of the program.

A **channel** is go's fundamental communication mechanism for goroutines. It
works much like a generator function in javascript.

```go
func emit(c chan string) {
	words := []string{"beware", "of", "gophers"}

	// push into the channel
	for _, word := range words {
		c <- word
	}
	close(c)
}

func main() {
	wordChannel := make(chan string)

	// start a goroutine. func emit will run in a separate thread
	go emit(wordChannel)

	// get all words as they come in over the channel
	for _, word := range wordChannel {
		fmt.Printf("%s\n", word)
	}
	// alternatively (get one word)
	word := <-wordChannel
	fmt.Printf("%s\n", word)

	// or even
	fmt.Printf("%s\n", <-wordChannel)
}
```

reading from a channel with `<-` returns `value, done` like a javascript
iterator. this you can use to determine if you want to keep reading from the
channel or not.

You can run as many goroutines as you want in parallel, but you need to make
sure to close the channel at some point to avoid a **deadlock** situation
where the code that reads from the channel will never receive word that it
should stop reading.

A use case for channels would be to use a channel to receive unique id's
from a generator function, which keeps its state to itself.

```go
func getID(idChannel chan int) {
	id := 0
	for {
		id++
		idChannel <- id
	}
}
func main() {
	idChan := make(chan int)

	go getID(idChan)

	fmt.Printf("id: %d\n", <-idChan)
	fmt.Printf("id: %d\n", <-idChan)
	fmt.Printf("id: %d\n", <-idChan)
	// on and on!
}
```

### the `select` keyword

Coordination between goroutines. Using `select`, it is possible to transmit
or to listen on multiple channels at once. `select` has `case`s, like a
`switch`. The following program will loop over a slice of words, and send a
boolean into a channel that will be picked up by the generator function as a
termination signal.

```go
func emit(wordChannel chan string, doneChannel chan bool) {
	i := 0
	words := []string{"feed", "the", "monkey"}
	for {
		select {
		case wordChannel <- words[i]:
			i++
			if i == len(words) {
				i = 0
			}
		case <-doneChannel:
			close(doneChannel)
			doneChannel <- true
		}
	}
}
func main() {
	wordChannel := make(chan string)
	doneChannel := make(chan bool)

	go emit(wordChannel, doneChannel)

	for i := 0; i <= 100; i++ {
		fmt.Println(<-wordChannel)
	}
	doneChannel <- true
	<- doneChannel // will block until generator sends value
}
```
