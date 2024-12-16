package utils

import (
	"bufio"
	"fmt"
	"os"
)

// waitForKeyPress waits until the user presses Enter or Spacebar
func WaitForKeyPress() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Press Enter or Spacebar to continue...")
	for {
		input, _ := reader.ReadByte()
		if input == '\n' || input == ' ' { // Check for Enter or Space
			break
		}
	}
	fmt.Print("\033[F\033[K") // Clear the "Press Enter" message
}
