# GoCodeGen

Golang is an awesome language, the std library have a lot of modules writen to help us to develop fast and reliable code. But of course go isn't perfect. Some times golang can be repetitive or/and verbose because of some practices. For example:

Imagine a situation that you want to use a linked list instead of an array, fine, golang have `container/list`, that implement a linked list, this is fine enough to most of the situations but it uses `interface{}` as the list type, so the list can hold any value we need. That's the go default aproach to write code that may need to support multiple types.

That cause the follow happen quite often into the code:

```go
list := list.New()
// list insertions of integers
firstNode, err := l.Front()
if err != nil {
  // ERROR, VALUE DOESN'T EXIST
}
firstValue, ok := fisrtNode.Value.(int)
if !ok {
  // THIS ISN'T AN INTEGER!!!! (╯°□°）╯︵ ┻━┻
}
```

So its nescessary to make a conversion and error verification every time we receive a value from the list because the use of `interface{}`.

To avoid this we can write our own struct IntList that wrap a `container/list` and write methods that return the values as integers:

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

To solve this i have creted this code generation based on templates. GoCodeGen only need to receive the path of the template file and the temaplate variables to substitute. That's how we call this tool:

```
$ GoCodeGen ./<relative-path-to-file> <FILE_PREFIX_NAME> --var1 value1 --var2 value2
```

To undestand better what GoCodeGen does let's see an example, the last IntList became the follow template:

```go
package main

type {{.list_name}}List struct {
  l list.List
}

func (il *{{.list_name}}List) Front() ({{.list_type}}, error) {
  value, err := il.l.Front()
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

Here in the first command `./gen/list.gen.go` (the `./` is obrigatory to indicate that that's a relative path) is the template file path, `int` is the file result prefix and the arguments with `--` is the name of the template variable and the next argument the value of this template variable.

Now we have three files with the correct names and types. The result is writen into files with the same name but with the prefix passed as parameter in front of the current file name. So we have the files `int_list.gen.go`, `float_list.gen.go` and `string_list.gen.go`.

Remember that the files are writen in the same path that tou called GoCodeGen.
