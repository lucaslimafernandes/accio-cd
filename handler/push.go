package handler

import (
	"fmt"
	"log"
	"os/exec"

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

		lsCmd := exec.Command(cmd)
		lsOutput, err := lsCmd.Output()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Saída do comando %v\n:", v.Command)
		fmt.Println(string(lsOutput))

	}
}
