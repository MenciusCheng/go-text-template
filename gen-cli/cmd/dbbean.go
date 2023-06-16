package cmd

import (
	"fmt"
	"github.com/MenciusCheng/go-text-template/gen-cli/config"
	"github.com/MenciusCheng/go-text-template/genxy/generate_method"
	"github.com/spf13/cobra"
	"os"
)

var table string
var db string

func init() {
	dbbeanCmd.Flags().StringVarP(&table, "table", "t", "", "table name")
	dbbeanCmd.MarkFlagRequired("table")

	dbbeanCmd.Flags().StringVar(&db, "db", "", "db name")
	dbbeanCmd.MarkFlagRequired("db")

	rootCmd.AddCommand(dbbeanCmd)
}

var dbbeanCmd = &cobra.Command{
	Use:   "dbbean",
	Short: "",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fileName := args[0]

		database := config.GetDatabaseByDb(db)
		if database == nil {
			fmt.Println("db config not found")
			os.Exit(1)
		}

		beans := generate_method.GetBeans(generate_method.DbParam{
			User:     database.User,
			Password: database.Password,
			Host:     database.Host,
			Port:     database.Port,
			TableNameLike: map[string]generate_method.TableConfig{
				table: {},
			},
			TableSchema: []string{db},
		})

		fmt.Printf("fileName: %s", fileName)
		fmt.Printf("beans: %+v", beans)
	},
}
