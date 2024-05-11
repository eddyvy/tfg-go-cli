package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tfg [project name]",
	Short: "Application for creating a REST API from a Postgresql Database",
	Long: `This applications is a tool created as a final degree project
with the aim of creating a new project with the use of a CLI.`,
	Run: func(cmd *cobra.Command, args []string) {
		pConf, err := readProjectConfig()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println(pConf)

		dConf, err := readDatabaseConfig()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println(dConf)

		db, err := connectDatabase(dConf)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer db.Close()

		tables, err := chooseTables(db)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println("Chosen tables:", tables)
		os.Stdout.Sync()
		fmt.Println("Creating project...")

		err = createProject(pConf, dConf, tables)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
	Args: cobra.MatchAll(cobra.ArbitraryArgs),
}

func init() {
	initProjectFlags(rootCmd)
	initDatabaseFlags(rootCmd)
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
