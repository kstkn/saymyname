package saymyname

import "fmt"

func (n Name) String() string {
	return fmt.Sprintf("%s %s", n.First, n.Last)
}

func (n Name) Reverse() string {
	return fmt.Sprintf("%s %s", n.Last, n.First)
}
