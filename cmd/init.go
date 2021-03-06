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
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initializing the db before first usage.",
	Long: `Setups the db and create necessary buckets. Usage:
passvault init --username <username> --password <password>
After init, you will only allowed to access with this username and password`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("init called")
        initDB()
		// TODO Create first user
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.Flags().StringP("username", "u", "", "Set your username")
	initCmd.Flags().StringP("password", "p", "", "Set your password")
}

func initDB() (*db.DB, error) {
	db := &db.DB{}
	if err := db.Open(); err != nil {
		return nil, err
	}
    return db, nil
}

