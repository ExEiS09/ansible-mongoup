package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	cmd1, err := exec.Command("vagrant", "up").CombinedOutput()

	if err != nil {
		os.Stderr.WriteString(err.Error())
	}
	fmt.Println(string(cmd1))

	if err := cmd1.Wait(); err != nil {
		cmd2, err := exec.Command("vagrant", "destroy").CombinedOutput()
		var choice y
		fmt.Scanln(&choice)

		fmt.Println(string(cmd2))
	}

}
