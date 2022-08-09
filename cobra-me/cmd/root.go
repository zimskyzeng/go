package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// 创建根命令，后续所有的命令都基于此
var (
	rootCmd = cobra.Command{
		Use:   "hugo",
		Short: "Hugo hugo hugo",
		Long:  "Hugo hugo hugo Hugo hugo hugo",
		SilenceUsage: true,
		Args: func(cmd *cobra.Command, args []string) error {
			fmt.Println("args", args)
			return nil
		},
		// 在运行前执行的钩子函数
		PreRun: func(cmd *cobra.Command, args []string) {
			fmt.Println("prerun hugo")
		},
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("hugo", Pname, Name)
		},
		// 在运行后执行的钩子函数
		PostRun: func(cmd *cobra.Command, args []string) {
			fmt.Println("postrun hugo")
		},
	}

	Name string
	Pname string
)

func init(){
	// 使用Flags仅对当前的该子命令有效
	sayHello.Flags().StringVarP(&Name, "name", "n", "zimskyzeng", "default zimskyzeng")
	// 使用PersistentFlags对下级的子命令也可以生效使用
	rootCmd.PersistentFlags().StringVarP(&Pname, "pname", "p", "pp", "ppp")
	// args的设置方式
	rootCmd.SetArgs([]string{"a", "b", "c"})
	rootCmd.AddCommand(sayHello)
}

var sayHello = &cobra.Command{
	Use: 		"hello",
	Short:  	"hello",
	Long: 		"hello",
	Args: func(cmd *cobra.Command, args []string) error {
		fmt.Println("args", args)
		return nil
	},
	PreRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("prerun hello")
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("hello", Name, Pname)
	},
}


func Execute() error {
	return rootCmd.Execute()
}
