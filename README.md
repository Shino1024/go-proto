# go-proto
## The collection of data structures and some algorithms in Go.

A collection of ready-to-import Go packages, which are friendly in usage. They provide
a quite wide functionality, ensuring a somewhat good performance.

This collection takes advantage of the reflect, fmt and errors packages mostly. 

The structures are capable of holding any type of data, though not for certain structures,
such as set or tree.

Lots of useful methods are provided with these structures in order to ensure flexibility.

To install the structures, use the `go get` command.

## Up-to-date list of the data structures:
- `lister` (linked list),
- `queuer` (queue),
- `stacker` (stack),
- `multiseter` (multiset),
- `seter` (set),
- `treer` (red-black tree).

## Information and warnings.
lister, queuer and stacker can store any kind of object. However, it's up to the user to
perform the type assertion whenever an element retrieval happens. This is because objects
are stored as `interface{}`'s in them. As I mentioned above, I strived to provide
flexibility, yet it's up to the user to remember what types of objects were stored before.

In terms of stacker, multiseter and treer, objects need to be ordered (so that they can be sorted).
That's why `XLesser` interfaces needs to be implemented. `X` is the letter for the corresponding
data structure -- `S` for `seter`, `M` for `multiseter` and `T` for `treer`. Thanks to that, it's
possible to use all of these data structures in a single file (the names are differentiated).

`XLesser` for `int`s can be implemented like that:
```go
type Int int
func (i Int) Less(other XDataStructure.XLesser) bool {
	return i < other.(Int)
}
```
`XLesser`s consist only of the `Less()` function. It takes another `XLesser` object as an argument
and returns a bool. Remember about performing type assertion of the provided object.

Practical example:
```go
type Int int
func (i Int) Less(other multiseter.MLesser) bool {
	return i < other.(Int)
}
```

Yes, that means that Go primitive types need to be aliased and given the `Less()` function.

(The project is in development and many changes can happen.)

Using the sakeven's rbtree source file for the treer (https://github.com/sakeven/RbTree).

There's documentation of every package available in every package's directory (generated with godoc).

# TO DO: