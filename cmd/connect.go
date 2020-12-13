/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/Ogguz/passvault/cryption"
	"os"
)

// connectCmd represents the connect command
var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		username, _:= cmd.Flags().GetString("username")
		password, _:= cmd.Flags().GetString("password")
		// TODO env ile al https://ordina-jworks.github.io/development/2018/10/20/make-your-own-cli-with-golang-and-cobra.html
		u := user{
			username: username,
			password: password,
		}
		if checkCredentials(u) {
			fmt.Println("Login succeed")
		} else {
			fmt.Println("Login failed")
			os.Exit(2)
		}
	},
}

type user struct {
	username string
	password string
}

func checkCredentials(u user) bool {

    if u.username == "emrah" && u.password == "adamindibi" {
    	return true
	}
	return false

}

func getUserPassFromFile (filename string) user {
    cryption.DecryptFromFile(filename,"a")
    return user{
		username: "",
		password: "",
	}
}

func init() {
	rootCmd.AddCommand(connectCmd)

	// Here you will define your flags and configuration settings.

    connectCmd.Flags().StringP("username", "u", "", "Username for auth")
	connectCmd.Flags().StringP("password", "p", "", "Password for auth")
}
