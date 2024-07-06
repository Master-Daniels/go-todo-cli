package cmd

import (
	"fmt"
	"os"

	"github.com/MasterDaniels/todo/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var datafile, cfgFile, userLicense string

func init() {
	cobra.OnInitialize(initConfig)

	home, err := utils.Dirname()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	rootCmd.PersistentFlags().StringVarP(&datafile, "datafile", "f", home+"/"+"todos.json", "datafile to store todos(must end in `.json` extension)")
	rootCmd.Flags().StringP("author", "n", "MASTER DANIELS", "Author name for copyright attribution")
	rootCmd.PersistentFlags().StringVarP(&userLicense, "license", "l", "", "Name of license for the project (can provide `licensetext` in config)")
	rootCmd.PersistentFlags().Bool("viper", true, "Use Viper for configuration")
	viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
	viper.BindPFlag("useViper", rootCmd.PersistentFlags().Lookup("viper"))
	viper.SetDefault("author", "MASTER DANIELS adebayooluwasegun011@gmail.com")
	viper.SetDefault("license", "apache")
}

var rootCmd = &cobra.Command{
	Use:   "todo",
	Short: "Todo is a very fast todo-cli application built using Go",
	Long: `Todo is a very fast todo-cli application built using Goloang, 
        	the language of tomorrow.
          The todo application allows users to manage and organize tasks
          directly from the terminal. Users can input and manage tasks,
          set`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to the Todo cli app.\nUse --help to see a summary on how to use this app.")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func GetdataFile() string {
	if viper.GetString("datafile") == "" {
		return datafile
	}
	return viper.GetString("datafile")
}

func initConfig() {
	// Don't forget to read config either from cfgFile or from home directory!
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := utils.Dirname()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".todo" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".todo")
		viper.AutomaticEnv()
		viper.SetEnvPrefix("todo")
	}

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
