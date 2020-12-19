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
)
// TODO add description and get flags as parameters
// TODO add control if vault already exist
// addVaultCmd represents the addVault command
var addVaultCmd = &cobra.Command{
	Use:   "addVault",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("addVault called")
		save("a","a","d")
	},
}

func init() {
	rootCmd.AddCommand(addVaultCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addVaultCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addVaultCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func save(name,username,password string) {
	dbBolt, err := newDB()
	if err != nil {
		log.Fatal(err)
	}
	defer dbBolt.Close()

	dbBolt.Update(func(tx *db.Tx) error {

		v := db.Vault{
			Tx:       tx,
			Name:     []byte(name),
			Username: []byte(username),
			Password: []byte(password),
		}

		return v.Save()
	})

}
