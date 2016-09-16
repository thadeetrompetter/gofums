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

#### re-declaring variables

Go complains if you assign no new values to a variable that has been declared
before, using `:=` syntax.

```go
// will complain: no new variables on left side of :=
foo, bar := true, 2
foo, bar := true, 2

// it's ok if you redeclare the first variable
foo, bar := true, 2
foo2, bar := true, 2
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

Delimiting a string with backticks means that its contents will be taken
literally. Special characters will lose their meaning.

A single rune (character) can be enclosed by single quotes: `'r'`.


### control structures

#### if
It's possible to declare a variable in an if statement. The variable will then
be scoped to the if statement's block.

```go
if bytes, err := fmt.Printf("%s\n", "foobar"); err == nil {
	fmt.Printf("%d\n", bytes) // bytes is available inside if, else, else if
}
// but not outside
```

Take care not to inadvertently shadow a named return value:
Go vet will warn you about this.

```go
func empty(i []int) (n int) {
	if n := len(i); n > 0 {
		n = 1 // this is likely not the variable you want to return
	}
	return n
}

func main() {
	intSlice := []int{1, 2, 3, 4, 5}
	result := empty(intSlice)
	fmt.Printf("foo: %d\n", result)
}
```
#### switch

The **switch** statement differs from javascript in that it is not necessary
 to add a `break` after each case that you want not to fall through.
You can put as much comma separated logic as you like in each case statement,
making the **go** switch much more flexible than in other c-like languages.
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

#### loop labels

Labeling loops allows you to break/continue out of one or more levels of loop
nesting.

```go
func main() {
	var (
		numbers = []int{1, 2, 3, 5, 0, 5, 8, 6}
		odd     int
		even    int
	)

	// Label the first loop level
First:
	for i := 0; i < 10; i++ {
		for _, number := range numbers {
			if number == 0 {
				break First // break loop when 0 is encountered
			}
			if number%2 == 0 {
				even++
			} else {
				odd++
			}
		}
	}
	// will print 1 2 3 5 and then stop
	fmt.Printf("odd: %d, even: %d\n", odd, even)
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
A deadlock error will occur if `main` doesn't receive the stop signal over
`doneChannel`.

Here is an example of using a timer from the `time` package to close a channel
after three seconds have passed.

```go
func emit(wordChannel chan string) {
	defer close(wordChannel)
	i := 0
	t := time.NewTimer(3 * time.Second)

	words := []string{"feed", "the", "monkey"}
	for {
		select {
		case wordChannel <- words[i]:
			i++
			if i == len(words) {
				i = 0
			}
		case <-t.C:
			return // will trigger the deferred closing of the channel
		}
	}
}
func main() {
	wordChannel := make(chan string)

	go emit(wordChannel)

	for word := range wordChannel {
		fmt.Printf("%s\n", word)
	}
}
```

### channels of channels

The following code will make the emit function responsible for creating the
channel(s) on which the main func will receive words. Now the main function
only needs to create one channel to receive on. `emit` will push words into the
word channel until the timeout strikes.

```go
func emit(chanCh chan chan string) {
	wordCh := make(chan string)
	i := 0
	t := time.NewTimer(3 * time.Second)
	words := []string{"feed", "the", "monkey"}

	defer close(wordCh) // defer close (will happen after timeout)

	chanCh <- wordCh // send the word channel back

	for {
		select {
		case wordCh <- words[i]:
			i++
			if i == len(words) {
				i = 0
			}
		case <-t.C:
			return
		}
	}
}
func main() {
	chanCh := make(chan chan string)

	go emit(chanCh)

	wordCh := <-chanCh // use whatever comes from chanCh as wordChan

	for word := range wordCh {
		fmt.Printf("%s\n", word)
	}
}
```

### multiple readers, multiple writers

A program that retrieves webpages and returns their content length over a
channel. Makes use of the `net/http` package.

```go
func getPage(url string) (int, error) {
	res, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	body, err := ioutil.ReadAll(res.Body)
	// body needs to be closed
	defer res.Body.Close()
	// if error: return 0 length and the error.
	if err != nil {
		return 0, err
	}
	return len(body), nil // otherwise length of the byteslice
}
// run this func as a goroutine
func worker(urlCh chan string, sizeCh chan string) {
	for {
		url := <-urlCh
		length, err := getPage(url)
		if err == nil {
			sizeCh <- fmt.Sprintf("%s: %d", url, length)
		} else {
			sizeCh <- fmt.Sprintf("Error getting %s: %s", url, err)
		}
	}
}
// run this func as a goroutine
func generator(url string, urlCh chan string) {
	urlCh <- url
}

func main() {
	urls := []string{
		"http://voorhoede.nl",
		"http://trabalhar.voorhoede.nl",
		"http://nu.nl",
	}
	sizeCh := make(chan string)
	urlCh := make(chan string)

	// create some work for the workers
	for _, url := range urls {
		go generator(url, urlCh)
	}
	// determine the number of workers (as many as there are urls here)
	for i := 0; i < len(urls); i++ {
		go worker(urlCh, sizeCh)
	}
	//
	for i := 0; i < len(urls); i++ {
		size := <-sizeCh
		fmt.Printf("%s\n", size)
	}
}
```

### Closing channels

It's not possible to 'broadcast' a message on a channel that will be received
by all 'listeners'. However, when you `close` a channel, you achieve a similar
effect. When the channel is closed, `<- goCh` will stop blocking the thread and
the messages will be printed.

The alternative is to use a `select` in `printer` with two cases. This way you
can repeat an action (printing messages) until `goCh` is closed which causes
the function to return.

```go
func printer(msg string, goCh chan bool) {
	<-goCh // this will block until channel close
	fmt.Printf("%s\n", msg)
}
// OR
func printer(msg string, goCh chan bool) {
	for {
		select {
		case <-goCh:
			return
		default:
			fmt.Printf("%s\n", msg) // will print until goCh terminates
		}
	}
}
func main() {
	goCh := make(chan bool)
	for i := 0; i < 10; i++ {
		go printer(fmt.Sprintf("Printer: %d\n", i), goCh)
	}
	time.Sleep(2 * time.Second)
	close(goCh)
	time.Sleep(3 * time.Second)
}
```

### nil channel

A nil channel is non-blocking. It's useful to selectively disable part of a
select statement. Setting a channel to `nil` inside a `select` means that the
case that listens to that specific channel, will be ignored.

You can also set the emitting side of the channel to nil to stop transmitting.

If you save a copy of the channel before setting the channel to nil, you can
resume reception or transmission later.

```go
func reader(intCh chan int) {
	// will stop receiving after 5 seconds
	t := time.NewTimer(5 * time.Second)
	for {
		select {
		case i := <-intCh:
			fmt.Printf("%d\n", i)
		case <-t.C:
			intCh = nil
		}
	}
}
func writer(intCh chan int) {
	stopper := time.NewTimer(1 * time.Second)
	restarter := time.NewTimer(3 * time.Second)
	// save a copy of the channel to restore it later
	intChBackup := intCh

	for {
		select {
		// will stop transmitting after 1 second and restart after 3 seconds
		case intCh <- rand.Intn(42):
		case <-stopper.C:
			intCh = nil
		case <-restarter.C:
			intCh = intChBackup
		}
	}
}
func main() {
	intCh := make(chan int)
	go reader(intCh)
	go writer(intCh)
	// prevent the program from exiting
	time.Sleep(10 * time.Second)
}
```

### buffered channels

Buffered channels allow you to transmit on a channel up to a certain threshold,
without a receiver having registered. It can be used as a semaphore to control
the amount of concurrent tasks running.

You buffer a channel by passing a numerical value after the channel declaration
in `make`:

```go
bufferedChannel := make(chan string, 100) // will contain up to 100 strings
```

Here is an example having a lot of workers of which 10 can run at a time.
`atomic.AddInt64` will count the number of routines running in parallel.

```go
// conversion because int * int64 is not possible
func work() {
	atomic.AddInt64(&running, 1)
	fmt.Printf("%d ", running)
	time.Sleep(time.Second * time.Duration(rand.Intn(10)))
	atomic.AddInt64(&running, -1)
}
func worker(sema chan bool) {
	<-sema       // block until value comes in
	work()       // synchronously do some work
	sema <- true // re-insert value into channel for the next worker to start
}
func main() {
	sema := make(chan bool, 10) // buffered channel
	// 1000 worker, but only 10 can run at a time
	for i := 0; i < 1000; i++ {
		go worker(sema)
	}
	// kick off the queue by inserting 10 boolean into buffered channel
	for i := 0; i < cap(sema); i++ {
		sema <- true
	}
	time.Sleep(time.Second * 30)
}
```

### types and the `type` keyword

User defined types are go's take on object oriented programming. This is how
you define a custom type:

```go
type foo struct {
	name string
	id int
	attributes []string
}
```

You 'instantiate' a type with the `new` keyword

```go
foo := new(foo)
// or
foo2 := &foo {
	// properties
}
// or
foo3 := foo {
	// properties
}
```

In general, you'll use pointers to the struct instead of its value (foo3).

You can create an anonymous struct.

```go
foo4 := struct {
	// properties
}
```
Useful for one-off structs, like the ones you would use to specify the template
for (un)marshalling a json response.

You can define a method on the custom type:

```go
func (f *foo) myMethod(arg string) string {
	return fmt.Sprintf("Foo: %s")
}
```

An example of the use of a custom type in context:

```go
type webPage struct {
	url  string
	body []byte
	err  error
}

func (w *webPage) get() {
	resp, err := http.Get(w.url)
	if err != nil {
		w.err = err
		return
	}
	defer resp.Body.Close()
	w.body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		w.err = err
	}
}

func (w *webPage) isOk() bool {
	return w.err == nil
}

func main() {
	w := &webPage{url: "http://www.voorhoede.nl"}
	w.get()
	if w.isOk() {
		fmt.Printf("URL: %s, Error: %s, Length: %d\n", w.url, w.err, len(w.body))
	} else {
		fmt.Println("Something went wrong")
	}
}
```
Types can be structures, but you can also redefine existing types:

```go
// SummableSlice ...
type SummableSlice []int

// can't feed a pointer here..
func (s SummableSlice) sum() int {
	sum := 0
	for _, n := range s {
		sum += n
	}
	return sum
}
func main() {
	numbers := SummableSlice{4, 7, 83, 29, 2, 1, 8, 5}
	fmt.Println(numbers.sum())
}
```

### interfaces

Interfaces define behavior instead of a specific type. Types will satisfy an
interface as soon as they have methods that the interface requires.

```go
type shuffler interface {
	Len() int
	Swap(i, j int)
}

func shuffle(s shuffler) {
	for i := 0; i < s.Len(); i++ {
		j := rand.Intn(s.Len() - i)
		s.Swap(i, j)
	}
}

// will satisfies shuffle interface because it has Len and Shuffle methods
type intSlice []int

func (is intSlice) Len() int {
	return len(is)
}
func (is intSlice) Swap(i, j int) {
	is[i], is[j] = is[j], is[i]
}

// stringSlice will also satisfy the interface, with different data
type stringSlice []string

func (ss stringSlice) Len() int {
	return len(ss)
}
func (ss stringSlice) Swap(i, j int) {
	ss[i], ss[j] = ss[j], ss[i]
}

func main() {
	is := intSlice{1, 2, 3, 4, 5, 6, 7, 8, 9}
	ss := stringSlice{"foo", "bar", "baz", "one", "two", "three"}
	shuffle(is)
	shuffle(ss)
	fmt.Printf("is: %d. ss: %s\n", is, ss)
}
```

### empty interfaces

Empty interfaces are used like a wildcard type. Anything satisfies an empty
interface. This way you can provide an unknown type and determine what you're
dealing with at runtime.

Asserting the type will cause a runtime panic if the program can not determine
the type.

```go
func whatIsThis(i interface{}) {
	// one way to find out a type is to use %T in Printf
	fmt.Printf("%T\n", i)
	// another way (e.g. in a switch statement)
	switch v := i.(type) {
	case int:
		// using the type switch
		fmt.Printf("it's an int: %d\n", v)
		// or a type assertion
		fmt.Printf("it's an int: %d\n", i.(int))
	case string:
		// using the type switch
		fmt.Printf("it's a string: %s\n", v)
		// or a type assertion
		fmt.Printf("it's a string: %s\n", i.(string))
	default:
		// useful in Printf: %v prints any value
		fmt.Printf("who knows? %v\n", v)
	}
}
func main() {
	whatIsThis(123)
}
```

### user defined packages

packages are not bound by directory structure, as long as they are in your
`$GOPATH`. In the `import` statement, packages from the standard library go on
top, followed by user defined packages, in alphabetical order.

Go will create a `pkg` folder in `$GOPATH` where it stores objects related to
the packages you create or import.

To allow variables to be exported from a package, you should write them with a
capital letter.

Let's put the shuffle functionality in its own package.

```go
// in $GOPATH/src/shuffler/shuffler.go

package shuffler

import "math/rand"

// Shuffleable ...
type Shuffleable interface {
	Len() int
	Swap(i, j int)
}

type WeightedShuffleable interface {
	Shuffleable
	Weight(i int) int
}

// Shuffle ...
func Shuffle(s Shuffleable) {
	for i := 0; i < s.Len(); i++ {
		j := rand.Intn(s.Len() - i)
		s.Swap(i, j)
	}
}

func WeightedShuffle(w WeightedShuffleable) {
	total := 0
	for i := 0; i < w.Len(); i++ {
		total += w.Weight(i)
	}
	for i := 0; i < w.Len(); i++ {
		pos := rand.Intn(total)
		cum := 0
		for j := i; j < w.Len(); j++ {
			cum += w.Weight(j)
			if pos >= cum {
				total -= w.Weight(j)
				w.Swap(i, j)
				break
			}
		}
	}
}

// in main.go

// import the package
import (
	"fmt"

	"shuffler"
)

// use
shuffler.Shuffle( /* type which satisfies the interface */ )
```

-- Example that doesn't work

All packages can define an `init` function which executes to do setup work when
the module loads. First all the standard packages fire init, afterward the user
defined packages.

You can reference an imported package by a shorter name:

```go
import shortname "myunreasonablylongpackagename"
```

If you want to keep a package in your imports list, but don't need it right
away, prefix it with an underscore to have go ignore it.

```go
import _ "idonotneedthispackagerightnow"

```
### go executable

`go build` will compile and check for errors
`go install` will compile and create an executable in `bin`
`go run` takes a filename for a quick run

### unit testing

Unit tests are easily created for your package by adding a file postfixed with
`_test`, containing the tests for the functions you export.

You run the tests with `go test {package}`.

Here's the package **poem**, which needs some unit tests.

```go

package poem

// Poem ...
type Poem []Stanza

// Stanza ...
type Stanza []Line

// Line ...
type Line string

// NewPoem you can use this kind of 'constructor' notation to return an instance
// but you could also use Poem directly, because it's exported
func NewPoem() Poem {
	return Poem{}
}

// NumStanzas ...
func (p Poem) NumStanzas() int {
	return len(p)
}

// NumLines ...
func (s Stanza) NumLines() int {
	return len(s)
}

// NumLines ...
func (p Poem) NumLines() (count int) {
	for _, s := range p {
		count += s.NumLines()
	}
	return
}

// Stats ...
func (p Poem) Stats() (numVowels, numConsonants int) {
	for _, s := range p {
		for _, l := range s {
			for _, r := range l {
				switch string(r) {
				case "a", "e", "i", "o", "u":
					numVowels++
				default:
					numConsonants++
				}
			}
		}
	}
	return
}

// poem_test.go

package poem

import "testing"

// make a test for every function
func TestNumStanzas(t *testing.T) {
	// because you're inside the package you're testing, you don't need imports
	poem := Poem{}
	if poem.NumStanzas() != 0 {
		t.Fatalf("Should be empty")
	}
	poem = Poem{{
		"flap jop nleuk",
		"bobo wammakk jenk",
	}}
	if poem.NumStanzas() != 1 {
		// Fatalf will make the test process end. If you want other tests to
		// run regardless of failure, use Errorf instead.
		t.Fatalf("Unexpected stanza count: %d\n", poem.NumStanzas())
	}
}
```

To easily create a collection of test input and expected output, you can use a
slice of anonymous structs. Loop over this slice of test io to keep it DRY

```go
var tests = []struct {
	x int
	y int
	out int
}{
	{1, 2, 3},
	{3, 2, 5},
	{10, 5, 15},
}

func TestAdd(t *testing.T) {
	for _, test := range tests {
		result := Add(test.x, test.x)
		if result != test.out {
			t.Errorf("expected %d + %d to be %d, got %d", test.x, test.y, test.out, result)
		}
	}
}
```
#### summary

* Create a test for a go file by creating a test file named
`{filename}_test.go`.
* Declare the same package name as the file you are testing.
* Import the `testing` package.
* Test a function from your file by creating a function in the test file:
`Test{Exported function name}`.
* The test function receives a pointer to the `T` method of the `testing`
package: `TestFn(t *testing.T)`.
* use methods `t.Fail()` to fail without reporting.
* `t.Errorf` will fail and allow you to print an printf style message.

**// TODO: verify this**
When testing a package, you can declare the package name in the same fashion
that groups go files with their associated tests, by appending a **_test**
postfix. This will make the test more strict.

### fmt

**Verbs** in c-like languages are the interpolation characters in Printf (among
others) such as `%s`, `%d` or `%v`. There are a some verbs that are specific to
go, such as:

* `%v` and `%#v` will print a complex type.
* `%T` gives you the type of a value.
* `%t` prints the value of a boolean
* `%q` prints a safely quoted string

read `godoc fmt` for more interesting verbs.

`fmt` exposes the `Stringer` interface, which allows you to determine how a
custom type should be printed. It's like overriding the default `toString`
method for an object in javascript.

```go
type banaan string

// banaan will satisfy the Stringer interface
func (b banaan) String() string {
	message := string(b)
	return fmt.Sprintf("%s, %s", message, "ik ben een banaan")
}
func main() {
	var b banaan = "hoi"
	fmt.Printf("%s\n", b)
}
```

### os, io, bufio

Go's standard library for interaction with the system and for doing io.

Low-level io is done with the **io** package, while **ioutil** and **bufio**
add some higher level abstractions.

This func, added to the `poem` package, reads the lines from a file into a
poem made up of stanzas separated by an empty line.

```go
func LoadPoem(name string) (Poem, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	p := Poem{}
	var s Stanza

	scan := bufio.NewScanner(f)

	for scan.Scan() { // loop over lines
		l := scan.Text() // read the current line
		if l == "" { // empty line means a new stanza
			p = append(p, s)
			s = Stanza{}
			continue
		}
		s = append(s, Line(l)) // explicitly convert string to a Line
	}

	if scan.Err() != nil {
		return nil, scan.Err()
	}

	p = append(p, s)

	return p, nil
}
```

### the standard `http` package

A built-in web client and server.

```go
// in main.go

func poemHandler(writer http.ResponseWriter, req *http.Request) {
	// tell go to collect form parameters
	req.ParseForm()
	// this will cause panic if you don't set the name parameter in the url
	poem := req.Form["name"][0]
	// Fprintf is printf for an io destination
	fmt.Fprintf(writer, "poem %s on its way", poem)
}
func main() {
	// define a route/handler mapping
	http.HandleFunc("/poem", poemHandler)
	// make the web server listen for incoming requests
	http.ListenAndServe(":8088", nil)
}
```

### json

`encoding/json` convert between json and go types.

You can transform responses into json by feeding a http.ResponseWriter into
json.NewEncoder.

The normal go rules with regard to variables being public or private also apply
here. lowercase variables will not be in the exported json.

```go
// will be written to ouput stream.
enc := json.NewEncoder(writer)
enc.Encode(poem)
```

You can also read json. e.g. a configuration file. The following function reads
from json, translating the **info** property into **Information** on the
`Config` struct.

```go

type Config struct {
	Route       string
	// this hints makes the decoder look for `info` instead of `Information`
	Information string `json:"info"`
}

func config(filename string) (*Config, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// use a pointer to store the read json.
	config := &Config{}
	decoder := json.NewDecoder(f)
	// Decode will only give an error back. The decoded json will go to config
	err = decoder.Decode(config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
```

### strings package

Useful methods for working with strings.

Using `strings.Split` to split a string on a separator to count the words in a
line of the poem.

```go
func (p Poem) NumWords() int {
	var count int
	for _, s := range p {
		for _, l := range s {
			sl := string(l)
			words := strings.Split(sl, " ")
			count += len(words)
		}
	}
	return count
}

// and a test
func TestNumWords(t *testing.T) {
	poem := Poem{{"hi there"}}
	if poem.NumWords() != 2 {
		t.Fatalf("expected 2 words")
	}
}
```

If you want to convert to the string type from another type that go can't just
convert, use the `strconv` package.

`strconv.ParseInt` is like javascript's parseInt. `strconv.Atoi` is a shorthand
for the most common application of `ParseInt`.

The `sort` package can sort collection of many types. If you have a custom type
that satisfies the `sort.Interface` interface, you can sort it with the `Sort`
method. This requires the type to have the methods `Len`, `Less` and `Swap`.

### the `sync` package

The sync package contains functions that help protecting data in concurrent
processing. Maps are not safe from overwriting each other. If you run a bunch
of goroutines in a loop which all write to the same map, you'll need the sync
package to protect its values.

Put a mutex around a write operation to a map that happens in a goroutine.

```go
cache := make(map[string]poem.Poem)
var mutex sync.Mutex
// inside a goroutine:
mutex.Lock()
cache["foo"] = "some value"
mutex.Unlock()
```

You can use `sync.WaitGroup` to make go wait for a collection of goroutines to
close. It works in three steps.

```go
var wg sync.WaitGroup

func main(){
	for _, foo := range something {
		wg.Add(1) // increment no of goroutines to wait for by 1
		go execute(){
			// do stuff
			wg.Done() // decrement by 1.
		}
	}
	wg.Wait() // blocks until all routines have finished
}
```

Prescription rule of thumb for using mutexes is that you should only use
mutexes if you're writing something with very little logic. If much logic is
needed you will end up copying it all over the place and you should then
consider using goroutines and channels instead.

### the standard `log` package

Does logging.

`log.Printf` does roughly the same as `fmt.Printf`.
`log.Fatalf` logs and then exits the program non-zero.

You can configure the log output formatting with `log.SetFlags`

The `log/syslog` package adds compatibility for logging to syslog files.

### the `flag` package

For supplying command line parameters to your program.

```go
func main() {
	bla := flag.String("foo", "bar", "set some value")
	// provide a flag calling the binary: -foo=somevalue
	flag.Parse()
	// a pointer is returned, so get the value like this:
	fmt.Println(*bla)
}
```

An alternative way to do this is to provide the pointer yourself instead of
receiving it from `flag`. Use `flag.StringVar`

```go
func main(){
	var bla string
	// provide a pointer here:
	flag.StringVar(&bla, "foo", "bar", "set some value")
	flag.Parse()
	// and write a regular variable where you need its value.
	fmt.Println(bla)
}
```

To access the raw arguments to the executable, use `flag.Arg(n)`. It's like
node's `process.argv[n]`

### the `time` package

Coming soon...

### dependency management with `godep`

* Install godep `go get godep`
* When installing a dependency, use `godep get {dep}`
* As soon as you have all your dependencies installed, run `godep save` to
write a json (npm shrinkwrap like) file describing your dependency tree, as
well as a `Godeps` directory containing all dependencies. These generated files
should go into versioning.
* A developer cloning the project will run `godep restore`. this will use the
files checked in under Godeps and put them in the `GOPATH` to make them usable.

Godep can also run your app sandboxed, with the contents of the `Godeps` dir
appended to your `GOPATH`: `godep go run {file}`.

### running shell commands with `exec`

Run a shell command from go. First you can check for the program you want to
run actually exists.
The command is: `const cmd string = "pwd"`

```go
path, err := exec.LookPath(cmd)
if err != nil {
	log.Fatalf("%s not found", cmd)
}
```
Now that you have found out if the command exists, you can proceed with
building an actual command. This will return you a pointer to `exec.Cmd`.

```go
command := exec.Command(path)
```

Next, you need to connect the command's stdout somewhere. Attach a
reference to a `bytes.Buffer`.

```go
var out bytes.Buffer
command.Stdout = &out
```

Ready to run the command. Another opportunity to check for an error.

```go
err = command.Run()
if err != nil {
	log.Fatalf("%s gave an error: %s", cmd, err)
}
fmt.Println(out.String()) // buffer toString
```

Using `command.Start`, you can run the command in the background. The thread
would block otherwise. When you execute `command.Wait()` later on, the thread
will block there until the result of running the command comes through.
