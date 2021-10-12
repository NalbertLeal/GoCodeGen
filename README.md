# GoCodeGen

Golang is an awesome language, the STD have a lot of modules writen to help us to develop fast and reliable code. But of course go isn't perfect. Some times golang can be repetitive or/and verbose because of the use of some practices. For example:

Image a situation that you want to use a linked list instead of an array, fine, golang have **container/list**, that implementation of linked list is fine enoth to most of the situations, but it uses **interface{}** as the list type, so the list can hold any value we need. That's the go default aproach to write code that may need to support multiple types.

That cause the follow code to show quite often into the code:

```go
list := list.New()
// list insertions of integers
fisrtNode, err := l.Front().Value
if err != nil {
  // ERROR, VALUE DOESN'T EXIST
}
firstValue := fisrtNode.(int)
```

So its nescessary to make a conversion every time we receive a value from the list because the use of interface{}.

To avoid this we can write ower own struct IntList that wrap a `container/list` and write methods that return the values as integers:

```go
type IntList struct {
  l list.List
}

func (il *IntList) Front() (int, error) {
  value, err := il.l,Front()
  if err != nil {
    return 0, err
  }
  return value.(int), nil
}
```

But if you need to use a list to integers, another to floats, another to string... then you need to rewrite this wrapper to each type, that's waste of time.

GoCodeGen only need to receive the path of the file with the code and the types. For example the last IntList became:

```go
type {{.list_name}}List struct {
  l list.List
}

func (il *{{.list_name}}List) Front() ({{.list_type}}, error) {
  value, err := il.l,Front()
  if err != nil {
    return 0, err
  }
  return value.({{.list_type}}), nil
}
```

Now just write the comand on terminal:

```shell
$ GoCodeGen ./gen/list.gen.go int --list_name Int --list_type int
$ GoCodeGen ./gen/list.gen.go float --list_name Float --list_type float
$ GoCodeGen ./gen/list.gen.go string --list_name String --list_type string
```

Now you have three files with the correct names and types. The result is writen into files with the same name but with the **NAME** passed as parameter in front of the current file name. So we have the files `int_list.gen.go`, `float_list.gen.go` e `string_list.gen.go`.

The so the command to run GoCodeGen is:

```
$ GoCodeGen ./<relative-path-to-file> <PREFIX_NAME> --var1 value1 --var2 value2
```