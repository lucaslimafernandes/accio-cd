package handler

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
	"time"

	"github.com/lucaslimafernandes/accio-cd/logs"
	"github.com/lucaslimafernandes/accio-cd/utils"
)

func EventPush(cdFile *utils.CDRunfile) {

	var lEntry logs.LogEntry

	fmt.Println("EventPush")

	for _, v := range cdFile.Tasks {
		var cmd string
		fmt.Printf("Executando comando %v...\n", v.Command)

		cmd = buildCommand(v.Command, cdFile)
		fmt.Println("HERE")

		parts := strings.Fields(cmd)
		args := parts[1:]
		cmdExec := exec.Command(parts[0], args...)
		cmdExec.Dir = cdFile.LocalPath // Define o diretório de trabalho para o comando

		fmt.Println("HERE1")
		// Executando o comando
		output, err := executeCommand(cmd, cdFile.LocalPath)
		fmt.Println("HERE2")
		if err != nil {
			log.Fatalf("Erro ao executar o comando '%s' em '%s': %v\nSaída: %s", v.Command, cdFile.LocalPath, err, string(output))
		}
		fmt.Printf("Saída do comando %v:\n", v.Command)
		fmt.Println(string(output))
	}

	lEntry.Timestamp = time.Now()
	lEntry.Level = "INFO"
	lEntry.Message = "Execução de tarefa de push concluída."
	lEntry.Status = "OK"

	logs.InsertLog(logs.DB, &lEntry)

}

func buildCommand(command string, cdFile *utils.CDRunfile) string {

	switch command {
	case "git clone":
		return fmt.Sprintf("%v %v", command, cdFile.GitUrl)
	default:
		return command
	}
}

func executeCommand(command, dir string) ([]byte, error) {
	parts := strings.Fields(command)
	cmd := exec.Command(parts[0], parts[1:]...)
	cmd.Dir = dir
	return cmd.CombinedOutput()
}
