package main

import "fmt"

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.
func (i IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", i[0], i[2], i[2], i[3])
}
func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}

	// slice story
	var holderslice [4]string
	holderslice[3] = "array it dont make use of append"
	slice := []string{"literal"}
	slice2 := make([]int, 5)
	fmt.Println((holderslice), slice, slice2)
	fmt.Println()

	// user define type slice story
	type slice3 []string
	h := slice3{"how", "you", "do"}
	h = append(h, "new")
	fmt.Println(h, len(h), cap(h))

	fmt.Println()

	// user define array story
	type array1 [4]string
	a := array1{"aow", "you", "do"}
	a[len(a)-1] = "new"
	fmt.Println(a, len(a), cap(a))

	fmt.Println(byte(3))
	fmt.Println([]byte("4"))
	fmt.Println(string(52))

	type bit []byte
	fmt.Println(bit{2, 3})
	fmt.Printf("%T--- %T %T %T\n", byte(0), 4, rune(-4), 'A')

	// what started this is how to use assertion in go
	var i interface{} = "3"

	s ,_ := i.(int)
	fmt.Printf("%T -- %v\n", s, s)

	b := []byte{2, '3'}
	fmt.Println(b)
	fmt.Printf("%T %T\n", b, '2')

	e:= []byte("thing")
	fmt.Printf("%T %v", e, e)
	fmt.Println('2')
}

