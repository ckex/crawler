package browser

import (
	"testing"
	"fmt"
)

func Test_ua(t *testing.T) {
	for key, val := range UserAgents {
		fmt.Println(key," -------------------------------------------------------------> ")
		for index,value := range  val{
			fmt.Println(index,value)
		}
	}
}
