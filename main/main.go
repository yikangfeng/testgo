package main

import (
	"fmt"
	//"os"
	_ "api"
	"goc"
	"service/commons"
	"utils"
)

//import "github.com/tidwall/evio"

const datafile = "D:/data.json"

func init() {
	fmt.Println("1")
}

type mp map[string]string

func (m mp) set(k, v string) {
	m[k] = v
}

type test struct {
	name string
	age  int
}

func (p *test) setValue() {

	p.age = 32
	p.name = "kf"
}
func main() {

	utils.ToJSONString()

	goc.GoCallC()

	goc.GoCallCTest()

	fmt.Println("abcdef" +
		"test")

	t := &test{name: "", age: 0}
	t.setValue()
	fmt.Printf("name=%s,age=%d\n", t.name, t.age)

	commons.Int2string(2)

	fmt.Println("2")
	fmt.Printf("call func return string=%s \n", commons.Int2string(2))
	var r, _ = commons.String2int("2")
	fmt.Printf("call func return int=%d \n", r)

	m := make(mp)
	m.set("k", "v")

	var s string = "abc"
	var a *string = &s
	*a = "def"
	fmt.Printf("a value is=%s\n", *a)

	var i int = 1
	var ip *int = &i
	var iip **int = &ip
	fmt.Printf("iip value is=%p", iip)

	var iiip ***int = &iip
	fmt.Printf("iiip value is=%p", iiip)
	var iiiip ****int = &iiip
	fmt.Printf("iiiip value is=%d", ****iiiip)

	/*var events evio.Events
		events.Data = func(c evio.Conn, in []byte) (out []byte, action evio.Action) {
			out = in
			return
		}
		if err := evio.Serve(events, "tcp://localhost:4000"); err != nil {
			panic(err.Error())
		}


		fi, err := os.Open(datafile)
	    if err != nil {
	        panic(err)
	    }
	     fi.Close()
	     fmt.Println("abc")
	*/

}
