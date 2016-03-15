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
`lister`, `queuer`, `treer` and `stacker` can store any kind of object. However, it's up to the user to
perform the type assertion whenever an element retrieval happens. This is because objects
are stored as `interface{}`'s in them. As I mentioned above, I strived to provide
flexibility, yet it's up to the user to remember what types of objects were stored before.

In terms of `seter` and `ultiseter`, objects need to be ordered (so that they can be sorted).
That's why `XLesser` interfaces with `XLess` methods needs to be implemented. `X` is the letter for the corresponding
data structure -- `S` for `seter` and `M` for `multiseter`. Thanks to that, it's
possible to use all of these data structures in a single file (the names are differentiated).

`XLesser` for `int`s can be implemented like that:
```go
type Int int
func (i Int) XLess(other XDataStructure.XLesser) bool {
	return i < other.(Int)
}
```
`XLesser`s consist only of the `Less()` function. It takes another `XLesser` object as an argument
and returns a bool. Remember about performing type assertion of the provided object.

Practical example:
```go
type Int int
func (i Int) XLess(other multiseter.MLesser) bool {
	return i < other.(Int)
}
```

There's a special case for the treer:
```go
type Int int
func (i Int) TLess(other interface{}) bool {
	return i < other.(Int)
}
```

`treer` can take any kind of value, because the comparing anything there bases on the keys. The key
is represented by the `TLesser` type (see example above). Insertion should consist of providing both
the key and the object to the `Insert()` method. The key can be, for example, an alias for `int`, just
like in the case of `MLesser` and `SLesser`.

It just takes an `interface{}` as an argument.

Yes, that means that Go primitive types need to be aliased and given the `XLess()` function.

(The project is in development and many changes can happen.)

Using the sakeven's rbtree source file for the treer (https://github.com/sakeven/RbTree).

There's documentation of every package available in every package's directory (generated with godoc).

# TO DO:
- nothing as for now... most likely.