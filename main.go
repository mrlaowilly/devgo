package main

import (
	//"fmt"
	//"github.com/mrlaowilly/devgo/cours"
	"fmt"
	"github.com/spf13/cobra"
)

func main() {
	//fmt.Println("Hello World", cours.LeNomDeLaVariablePublic, cours.LeNomDeLaVariablePublicAvecValue)
	//fmt.Println(cours.Hello("Willy"))
	//cours.Master()
	//u := cours.MakeUser("Willy")
	//u.Hello()
	var root = &cobra.Command{
		Use:   "app [source] [destination]",
		Short: "Permet de copier des fichiers",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			source := args[0]
			destination := args[1]

			fmt.Println(source, destination)
		},
		DisableFlagsInUseLine: true,
	}

	root.Execute()
}
