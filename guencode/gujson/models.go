package gujson

import (
	"fmt"
)

type Cmd struct {
	Action   string
	Status   string
	Priority int
	//	StartDate string
	// // // //
	AAAA string
	BBBB string
	CCCC string
	DDDD string
	// // // //
	Options []byte //
}

func (c *Cmd) Print() {

	fmt.Println(" :: Action     : ", c.Action)
	fmt.Println(" :: Status     : ", c.Status)
	fmt.Println(" :: Priority   : ", c.Priority)
	//	fmt.Println(" :: Start date : ", c.StartDate)

	fmt.Println(" :: URL      : ", c.AAAA)
	fmt.Println(" :: IP       :", c.BBBB)
	fmt.Println(" :: Domain   : ", c.CCCC)
	fmt.Println(" :: FQDN     : ", c.DDDD)
	fmt.Println(" :: Options  : ", string(c.Options))

}

func (s *Cmd) SetDefaultValues() {

	s.Action = "N/A"
	s.Status = "N/A"
	s.Priority = -1
	//	s.StartDate = "N/A"
	// // // //
	s.AAAA = "N/A"
	s.BBBB = "N/A"
	s.CCCC = "N/A"
	s.DDDD = "N/A"

}

// // // //// // // //// // // //// // // //// // // //// // // //// // // //// // //

type ArrCmd struct {
	Items []Cmd
}

func (c *ArrCmd) AddItem(item Cmd) []Cmd {
	c.Items = append(c.Items, item)
	return c.Items
}

func (t *ArrCmd) Print() {

	for _, i := range t.Items {
		i.Print()

	}
}
