# mergejson

merge two JSON strings into one

### Example

```go

 // merge map: same key name but different value
 a, err := mergejson.MergeJson([]byte(`{"user":{ "name": "hello" }}`), []byte(`{"user":{ "name": "world" }}`))
 if err != nil {
  t.Fatal(err)
  return
 }
 fmt.Printf("same key name but different value: %s\n", a)

 // merge map: same map but a different number of keys(src < dst)
 a1, err := mergejson.MergeJson([]byte(`{"user":{ "name": "hello" }}`), []byte(`{"user":{ "name": "world", "age": 12 }}`))
 if err != nil {
  t.Fatal(err)
  return
 }
 fmt.Printf("same key name but a different number of keys(src < dst): %s\n", a1)

 // merge map: same key name but a different number of keys(src > dst)
 a2, err := mergejson.MergeJson([]byte(`{"user":{ "name": "hello", "age": 12}}`), []byte(`{"user":{ "name": "world" }}`))
 if err != nil {
  t.Fatal(err)
  return
 }
 fmt.Printf("same key name but a different number of keys(src > dst): %s\n", a2)

 // merge slice
 a3, err := mergejson.MergeJson([]byte(`{"user":{ "name": "hello", "lucky_num": [1,2,3,4]}}`), []byte(`{"user":{ "name": "world", "lucky_num": [5,6,7]}}`))
 if err != nil {
  t.Fatal(err)
  return
 }
 fmt.Printf("same key name but a different number of keys(src > dst): %s\n", a3)
```
