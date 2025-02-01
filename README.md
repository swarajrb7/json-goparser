# JSON Parser

This is simple implementation of JSON parser in Go.
Example:
```
{
  "key": "value",
  "key-n": 101,
  "key-o": {
    "inner key": "inner value"
  },
  "key-l": ['list value']
}# 
```

It would return
```
Error:  2025/02/01 23:36:29 lexer error: char ' at line 7 col 14 is invalid
```


## Build and run

``` 
go run . <file_path>
```

OR

```
go build
 
./json-parser <file_path>

```

## Resources

[Blog Post](https://codingchallenges.fyi/challenges/challenge-json-parser/)

