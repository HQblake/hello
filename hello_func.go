package hello

import "fmt"

func (h Hello) String() string {
	return fmt.Sprintf("%v, %v%v", h.Greeting, h.Name, h.Pun)
}
