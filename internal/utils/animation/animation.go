package animation

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/afonsocraposo/advent-of-code-2024/internal/utils/matrix"
)

func printLine(n int) {
	for j := 0; j < n; j++ {
		fmt.Print("-")
	}
	fmt.Print("\n")
}

// waitForKeyPress waits until the user presses Enter or Spacebar
func waitForKeyPress() {
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

func handleAnimation(framerate int) {
	if framerate == 0 {
		waitForKeyPress()
	} else {
		time.Sleep(time.Duration(1000.0 / float64(framerate) * float64(time.Millisecond)))
	}
}

func PrintMatrix(mat matrix.Matrix, redraw bool, framerate int) {
	m, n := mat.Size()
	if redraw {
		fmt.Printf("\033[%dA", m+2)
	}
	for i, vector := range mat.Rows {
		if i == 0 {
			printLine(n + 2)
		}
		fmt.Printf("|%s|\n", vector.ToValuesString())
		if i == n-1 {
			printLine(n + 2)
		}
	}
	handleAnimation(framerate)
}

func PrintRuneMatrix(mat matrix.Matrix, title string, redraw bool, framerate int) {
	m, n := mat.Size()
	if redraw {
		if title != "" {
			fmt.Printf("\033[%dA", m+3)
		} else {
			fmt.Printf("\033[%dA", m+2)
		}
	}
    if title != "" {
        fmt.Println(title)
    }
	for i, vector := range mat.Rows {
		if i == 0 {
			printLine(n + 2)
		}
		fmt.Printf("|%s|\n", vector.ToTextString())
		if i == m-1 {
			printLine(n + 2)
		}
	}
	handleAnimation(framerate)
}
