package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var (
	cfgFile string
	rootCmd = &cobra.Command{
		Use:     "workit",
		Short:   "小航书-golang快速开发模板CLI工具",
		Version: Version,
	}
)

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "配置文件 (默认: $HOME/.workit.yaml)")

	// 只保留必要的命令
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(newCmd)
}

func initConfig() {
	if cfgFile != "" {
		return
	}

	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	cfgFile = filepath.Join(home, ".workit.yaml")
}

func Execute() error {
	return rootCmd.Execute()
}
