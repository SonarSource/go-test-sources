package samples

import "fmt"

func loop() {
	for i := 1; i < 10; i++ {
		fmt.Print(i)
		break // Noncompliant
		fmt.Print(i)
	}
}
