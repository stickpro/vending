package command

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/stickpro/vending/pkg/database/pgsql"
	"github.com/stickpro/vending/pkg/logger"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
)

var migrateUpCmd *cobra.Command

type DBConfig struct {
	Host     string `mapstructure:"DB_HOSTNAME"`
	Port     string `mapstructure:"DB_PORT"`
	Username string `mapstructure:"DB_USERNAME"`
	Password string `mapstructure:"DB_PASSWORD"`
	DBName   string `mapstructure:"DB_NAME"`
}

func init() {
	migrateUpCmd = &cobra.Command{
		Use:   "up",
		Short: "migrate to v1 command",
		Long:  `Command to install version 1 of our application`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Running migrate up command")
			const configDir = "configs"
			cfg, err := LoadConfig()

			if err != nil {
				logger.Error(err)
				return
			}

			fmt.Println(cfg)
			db, _ := pgsql.ConnectionDataBase(cfg.Host, cfg.Username, cfg.Password, cfg.DBName, cfg.Port)

			file, err := os.Open("./internal/domain")
			if err != nil {
				logger.Error("failed opening directory: %s", err)
			}
			defer file.Close()

			var models = []interface{}{}

			list, _ := file.Readdirnames(0) // 0 to read all files and folders
			for _, name := range list {
				fmt.Println(name)
				models = append(models, name)
			}

			fset := token.NewFileSet()
			f, err := parser.ParseFile(fset, "./internal/domain/user.go", nil, 0)
			if err != nil {
				logger.Error(err)
			}

			// Collect the struct types in this slice.
			var structTypes []*ast.StructType

			// Use the Inspect function to walk AST looking for struct
			// type nodes.
			ast.Inspect(f, func(n ast.Node) bool {
				if n, ok := n.(*ast.StructType); ok {
					structTypes = append(structTypes, n)
				}
				return true
			})
			fmt.Println(structTypes)
			//db.AutoMigrate(structTypes)
			fmt.Println("pkgs", db)
			//db.AutoMigrate(&domain)
		},
	}

	migrateCmd.AddCommand(migrateUpCmd)
}

func LoadConfig() (config DBConfig, err error) {
	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	config.Host = viper.GetString("DB_HOSTNAME_DOCKER")
	config.Port = viper.GetString("DB_PORT_DOCKER")
	config.Username = viper.GetString("DB_NAME")
	config.Password = viper.GetString("DB_USERNAME")
	config.DBName = viper.GetString("DB_PASSWORD")
	return
}
