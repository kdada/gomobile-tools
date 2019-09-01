package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"os/exec"

	apisv1 "github.com/kdada/gomobile-tools/cmdrunner/pkg/apis/v1"
)

var (
	addr = ""
	cwd  = ""
)

func init() {
	flag.StringVar(&addr, "a", addr, "addr")
	flag.StringVar(&cwd, "h", cwd, "work dir")
}

func main() {
	flag.Parse()
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	http.HandleFunc("/", run)

	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Println(err)
	}

}

func run(rw http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		write(rw, http.StatusMethodNotAllowed, &apisv1.CommandResult{Error: http.StatusText(http.StatusMethodNotAllowed)})
		return
	}

	command := &apisv1.Command{}
	if err := json.NewDecoder(req.Body).Decode(command); err != nil || len(command.Commands) <= 0 {
		write(rw, http.StatusBadRequest, &apisv1.CommandResult{Error: http.StatusText(http.StatusBadRequest)})
		return
	}

	log.Printf("Commands: %v\n", command.Commands)

	cmd := exec.Command(command.Commands[0], command.Commands[1:]...)
	if len(command.Input) > 0 {
		cmd.Stdin = bytes.NewBuffer(command.Input)
	}
	output := bytes.NewBuffer(nil)
	cmd.Stdout = output
	cmd.Stderr = output
	cmd.Dir = cwd

	result := &apisv1.CommandResult{}
	err := cmd.Run()
	if err != nil {
		result.Error = err.Error()
	}
	result.Output = output.Bytes()

	code := 200
	if err != nil {
		code = http.StatusInternalServerError
	}
	result.ExitCode = cmd.ProcessState.ExitCode()
	write(rw, code, result)
}

func write(rw http.ResponseWriter, code int, result *apisv1.CommandResult) {
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(code)
	data, _ := json.Marshal(result)
	rw.Write(data)
}
