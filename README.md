# Test Encoding Go Empty Map

```sh
$ go run main.go
nil
encoding/json {}
mailru/easyjson {}

nil map
encoding/json {}
mailru/easyjson {}

map of nil
encoding/json {"kv":null}
mailru/easyjson {"kv":null}

empty map
encoding/json {"kv":{}}
mailru/easyjson {"kv":{}}

not nil map
encoding/json {"kv":{"2":1}}
mailru/easyjson {"kv":{"2":1}}
```
