package debugmode

import (
	"bytes"
	"log"
	"os/exec"
)

var debug = false

func ActivateDebugMode() {
	debug = true
}

func DeactivateDebugMode() {
	debug = false
}

func IsDebugMode() bool {
	return debug
}

func debugExec(cmd *exec.Cmd) (err error) {
	args := cmd.Args
	newargs := make([]interface{}, len(args))
	for i, v := range args {
		newargs[i] = v
	}
	log.Print("Exec command: ")
	log.Println(newargs...)

	output, err := cmd.CombinedOutput()
	if len(output) > 0 {
		log.Println("Command output: ")
		log.Print(bytes.NewBuffer(output).String())
	}

	return err
}

func DebugExec(name string, args ...string) (err error) {
	cmd := exec.Command(name, args...)
	if IsDebugMode() {
		err = debugExec(cmd)
	} else {
		err = cmd.Run()
	}

	return err
}
