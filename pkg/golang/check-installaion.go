package golang

import (
	"fmt"
	"os/exec"
)

func CheckInstallation() (bool, error) {
	cmd := exec.Command("go", "version")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("cmd.Run() failed with %s\n", err)
		return false, err
	}

	fmt.Println("System has go installed\n" + string(out))
	return true, nil
}
