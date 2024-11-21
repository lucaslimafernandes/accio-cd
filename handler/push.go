package handler

import (
	"fmt"
	"log"
	"os/exec"
	"strings"

	"github.com/lucaslimafernandes/accio-cd/utils"
)

func EventPush(cdFile *utils.CDRunfile) {
	fmt.Println("EventPush")

	for _, v := range cdFile.Tasks {
		var cmd string
		fmt.Printf("Executando comando %v...\n", v.Command)

		switch v.Command {
		case "git clone":
			cmd = fmt.Sprintf("%v %v", v.Command, cdFile.GitUrl)
		default:
			cmd = v.Command
		}

		parts := strings.Fields(cmd)
		args := parts[1:]
		cmdExec := exec.Command(parts[0], args...)
		cmdExec.Dir = cdFile.LocalPath // Define o diretório de trabalho para o comando

		// Executando o comando
		cmdOutput, err := cmdExec.CombinedOutput() // Usando CombinedOutput para capturar stdout e stderr
		if err != nil {
			log.Fatalf("Erro ao executar o comando '%s' em '%s': %v\nSaída: %s", v.Command, cdFile.LocalPath, err, string(cmdOutput))
		}
		fmt.Printf("Saída do comando %v:\n", v.Command)
		fmt.Println(string(cmdOutput))
	}
}
