package jsonio

import (
	"os"
	"strings"
	"testing"
)

type TmpStruct struct {
	A int
	B string
	C bool
}

type Example struct {
	Text string
	MD5  string
}

var TheStruct []TmpStruct
var Ejson Example

func TestBoth(t *testing.T) {

	t1 := TmpStruct{A: 10, B: "One", C: false}
	TheStruct = append(TheStruct, t1)
	t2 := TmpStruct{A: 20, B: "Two", C: true}
	TheStruct = append(TheStruct, t2)
	t3 := TmpStruct{A: 30, B: "Three", C: true}
	TheStruct = append(TheStruct, t3)

	err := SaveJSon("tmp.json", TheStruct)
	if err != nil {
		t.Errorf("SaveJSon(tmp.json,TheStruct) = io error")
	}

	l := len(TheStruct)
	if l != 3 {
		t.Errorf("TheStruct want %d got %d", 3, l)
	}
	TheStruct = nil

	err = LoadJSon("tmp.json", &TheStruct)
	if err != nil {
		t.Errorf("LoadJSon(tmp.json,TheStruct) = io error")
	}
	err = os.Remove("tmp.json")
	if err != nil {
		t.Errorf("LoadJSon(tmp.json,TheStruct) = io error")
	}
	if TheStruct[0].A != 10 || TheStruct[0].B != "One" || TheStruct[0].C != false {
		t.Errorf("LoadJSon want %d %s %t got %d %s %t", 10, "One", false, TheStruct[0].A, TheStruct[0].B, TheStruct[0].C)
	}
	if TheStruct[1].A != 20 || TheStruct[1].B != "Two" || TheStruct[1].C != true {
		t.Errorf("LoadJSon want %d %s %t got %d %s %t", 20, "One", true, TheStruct[1].A, TheStruct[1].B, TheStruct[1].C)
	}
	if TheStruct[2].A != 30 || TheStruct[2].B != "Three" || TheStruct[2].C != true {
		t.Errorf("LoadJSon want %d %s %t got %d %s %t", 30, "Three", true, TheStruct[2].A, TheStruct[2].B, TheStruct[2].C)
	}

	es := "{ \"text\": \"example_text\", \"md5\": \"fa4c6baa0812e5b5c80ed8885e55a8a6\" }"
	LoadJSonFromString(es, &Ejson)
	if strings.Compare(Ejson.Text, "example_text") != 0 || strings.Compare(Ejson.MD5, "fa4c6baa0812e5b5c80ed8885e55a8a6") != 0 {
		t.Errorf("LoadJSonFromString wanted text:%s md5:%s but got text:%s md5:%s", "example_text", "fa4c6baa0812e5b5c80ed8885e55a8a6", Ejson.Text, Ejson.MD5)
	}

}
