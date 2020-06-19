Tutorial: Writing an Interpreter in Go
======================================

## About
Following a tutorial that implements an interpreter for the
Monkey programming language (Thorsten Ball's _Writing an
Interpreter in Go_).

## Requirements
* Go 1.13+

## Running Locally
```sh
$ make repl
```

## Testing
```sh
$ make test
```

## Examples
```
>> puts("Hello, world!")
Hello, world!
null
>> 1 + 1
2
>> let items = {"foo": "bar", "baz": "quux"}
>> items
{foo: bar, baz: quux}
>> items["foo"]
bar
>> items["oh no"]
null
>> let numbers = [1, 2, 3, 4, 5]
>> numbers
[1, 2, 3, 4, 5]
>> numbers[0]
1
>> numbers[9]
null
>> let square = fn(x) { x * x; }
>> square(2)
4
>> "o" + "m" + "g"
omg
```
