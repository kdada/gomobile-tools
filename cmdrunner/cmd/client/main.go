package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	apisv1 "github.com/kdada/gomobile-tools/cmdrunner/pkg/apis/v1"
)

var (
	addr = ""
)

func init() {
	env := os.Getenv("COMMAND_RUNNER_ADDR")
	if env != "" {
		addr = env
	}
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	args := os.Args
	args = args[1:]

	os.Stdin.Close()
	input, _ := ioutil.ReadAll(os.Stdin)

	cmd := apisv1.Command{}
	cmd.Input = input
	cmd.Commands = args

	data, _ := json.Marshal(cmd)
	resp, err := http.Post(fmt.Sprintf("http://%s/", addr), "application/json", bytes.NewBuffer(data))
	if err != nil {
		log.Fatalln(err)
	}

	result := &apisv1.CommandResult{}
	err = json.NewDecoder(resp.Body).Decode(result)
	if err != nil {
		log.Fatalln(err)
	}

	if result.Error != "" {
		log.Printf("Run command with error: %s\n", result.Error)
	}

	fmt.Printf("%s", result.Output)

	os.Exit(result.ExitCode)
}
