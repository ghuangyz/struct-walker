package swalker

import (
	"fmt"
	"testing"
)

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

func TestSimpleStruct(t *testing.T) {
	value6 := map[string]string{
		"sub-key1": "value1",
		"sub-key2": "value2",
	}
	value7_v1 := "value7_v1"
	value7 := map[string]*string{
		"value7_k1": &value7_v1,
	}
	value11 := S3{
		"value8",
		[]string{"value9_a", "value9_b"},
		map[string]string{
			"value10_k1": "value10_v1",
			"value10_k2": "value10_v2",
		},
	}
	value4 := S2{
		&value6,
		[]string{"value6_a", "value6_b"},
	}

	data := &SimpleStruct{
		"value1",
		"value2",
		"value3",
		value4,
		value7,
		&value11,
	}

	key := "key4.key5.sub-key1"
	value := GetValueOf(data, key)
	fmt.Printf("%v\n", value)

	value = GetValueOf(data, "key4.key6[1]")
	fmt.Printf("%v\n", value)

	value = GetValueOf(data, "key11.key10.value10_k2")
	fmt.Printf("%v\n", value)

	value = GetValueOf(data, "key11.key10.")
	fmt.Printf("%v\n", value)

	arrayData := [3]map[string]string{
		map[string]string{
			"key01": "value01",
			"key02": "value02",
		},
		map[string]string{
			"key11": "value11",
			"key12": "value12",
		},
		map[string]string{
			"key21": "value21",
			"key22": "value22",
		},
	}
	value = GetValueOf(arrayData, "[2].key21")
	fmt.Printf("%v\n", value)

	sliceData := []map[string]string{
		map[string]string{
			"key01": "value01",
			"key02": "value02",
		},
		map[string]string{
			"key11": "value11",
			"key12": "value12",
		},
		map[string]string{
			"key21": "value21",
			"key22": "value22",
		},
	}
	value = GetValueOf(sliceData, "[1].key11")
	fmt.Printf("%v\n", value)

}
