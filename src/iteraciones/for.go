package iteraciones

import (
	"fmt"
)

func Iterar() {
	// for i := 0; i < 100; i += 8 {
	// 	fmt.Println(i)
	// }
	for i := 0; i < 10; i++ {
		if i == 6 || i == 8 {
			continue
		}
		fmt.Println(i)
	}
}
