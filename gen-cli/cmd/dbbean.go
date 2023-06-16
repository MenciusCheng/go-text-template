package cmd

import (
	"fmt"
	"github.com/MenciusCheng/go-text-template/gen-cli/config"
	"github.com/MenciusCheng/go-text-template/genxy/generate_method"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"strings"
)

var table string
var db string
var patterns string

func init() {
	dbbeanCmd.Flags().StringVarP(&table, "table", "t", "", "table name")
	dbbeanCmd.MarkFlagRequired("table")

	dbbeanCmd.Flags().StringVarP(&db, "db", "d", "", "db name")
	dbbeanCmd.MarkFlagRequired("db")

	dbbeanCmd.Flags().StringVarP(&patterns, "patterns", "p", "", "more patterns")

	rootCmd.AddCommand(dbbeanCmd)
}

var dbbeanCmd = &cobra.Command{
	Use:   "dbbean",
	Short: "",
	Run: func(cmd *cobra.Command, args []string) {
		beanDir := ""
		if len(args) > 0 {
			beanDir = args[0]
			if beanDir[len(beanDir)-1] != '/' || beanDir[len(beanDir)-1] != '\\' {
				beanDir = fmt.Sprintf("%s%c", beanDir, os.PathSeparator)
			}
		}

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

		generate_method.WriteBean(beans, &generate_method.Config{
			BeanDir: beanDir,
		})
		beanFile := fmt.Sprintf("%s%s_generate.go", beanDir, table)
		if err := exec.Command("go", "fmt", beanFile).Run(); err != nil {
			fmt.Println("go fmt err", err)
			os.Exit(1)
		}

		ps := strings.Split(patterns, ",")
		for _, item := range ps {
			switch item {
			case "dao":
				daoDir := fmt.Sprintf("%s..%cdbdao%c", beanDir, os.PathSeparator, os.PathSeparator)
				generate_method.WriteDaoOneTableProjectId(beans, &generate_method.Config{
					DaoDir: daoDir,
				})
				daoFile := fmt.Sprintf("%s%s_dao_generate.go", daoDir, table)
				if err := exec.Command("go", "fmt", daoFile).Run(); err != nil {
					fmt.Println("go fmt err", err)
					os.Exit(1)
				}
			}
		}
	},
}
