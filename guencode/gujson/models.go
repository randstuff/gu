package gujson

import (
	"fmt"
)

type MyStructA struct {
	AAAA string
	BBBB string
	CCCC int
	DDDD string
	EEEE string
	FFFF string
	GGGG string
	HHHH string
	// // // //
	ZZZZ []byte //
}

func (c *MyStructA) Print() {

	fmt.Println(" :: AAAA : ", c.AAAA)
	fmt.Println(" :: BBBB : ", c.BBBB)
	fmt.Println(" :: CCCC : ", c.CCCC)
	fmt.Println(" :: DDDD : ", c.DDDD)

	fmt.Println(" :: EEEE : ", c.EEEE)
	fmt.Println(" :: FFFF :", c.FFFF)
	fmt.Println(" :: GGGG : ", c.GGGG)
	fmt.Println(" :: HHHH : ", c.HHHH)
	fmt.Println(" :: ZZZZ : ", string(c.ZZZZ))

}

func (s *MyStructA) SetDefaultValues() {

	s.AAAA = "N/A"
	s.BBBB = "N/A"
	s.CCCC = -1
	s.DDDD = "N/A"
	// // // //
	s.EEEE = "N/A"
	s.FFFF = "N/A"
	s.GGGG = "N/A"
	s.HHHH = "N/A"

}

// // // //// // // //// // // //// // // //// // // //// // // //// // // //// // //

type ArrMyStructA struct {
	Items []MyStructA
}

func (c *ArrMyStructA) AddItem(item MyStructA) []MyStructA {
	c.Items = append(c.Items, item)
	return c.Items
}

func (t *ArrMyStructA) Print() {

	for _, i := range t.Items {
		i.Print()
	}
}
