struct walker

Struct walker allows you to access go struct like the following with a simple chained key string, such as `"key4.key6[1]"`

```
type S2 struct {
	key5 *map[string]string
	key6 []string
}

type S3 struct {
	key8  string
	key9  []string
	key10 map[string]string
}

type SimpleStruct struct {
	key1  string
	key2  string
	key3  string
	key4  S2
	key7  map[string]*string
	key11 *S3
}
```

Usage:
```
value := swalker.GetValueOf(YourData, "somekey.subkey[0].subsubkey")
```
