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
	"github.com/Ogguz/passvault/db"
	"github.com/spf13/cobra"
	"log"
	"os"
)

// getVaultCmd represents the getVault command
var getVaultCmd = &cobra.Command{
	Use:   "getVault",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		getVault("abc")
	},
}

func init() {
	rootCmd.AddCommand(getVaultCmd)

}

func getVault(name string)  {
	dbBolt, err := newDB()
	if err != nil {
		log.Fatal(err)
	}
	defer dbBolt.Close()
	
	dbBolt.View(func(tx *db.Tx) error {
		v, err := tx.Vault([]byte(name))
		if err != nil || len(v.Name) == 0 || len(v.Credentials) == 0 {
			fmt.Printf("%s %s\n", name, err)
			os.Exit(2)
		}

		fmt.Printf("Credentials for vault %s;\n", v.Name)
		fmt.Printf("Username: %s\n", v.Name)
		fmt.Printf("Password: %s\n", v.Credentials)

		return nil
	})
	
}