package cmd

import (
	"http-test-tool/modules"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var (
	url       string
	proxy     string   //使用的代理
	AddHeader []string //添加的请求头
	bypassip  []string //添加的ip
	multiply  string //批量url头
	cfgFile   string   //配置文件
)

var rootCmd = &cobra.Command{
	Use:   "f403",
	Short: "f403 is a tool to bypass 40X response codes.",
	Long:  `f403 is a tool to bypass 40X response codes.`,
	Run: func(cmd *cobra.Command, args []string) {
		modules.Init(url, proxy, AddHeader, bypassip, multiply)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.Flags().StringVarP(&url, "url", "u", "", "the target url")
	//rootCmd.MarkFlagRequired("url")
	rootCmd.PersistentFlags().StringVarP(&proxy, "proxy", "p", "", "the proxy you will use,support http and socks5 ,for example: -p http://127.0.0.1:8080 or -p socks5://127.0.0.1:8080")
	rootCmd.PersistentFlags().StringSliceVarP(&AddHeader, "AddHeader", "a", []string{}, "the headers you will add,for explame cookie:123,Referer:https://www.baidu.com")
	rootCmd.PersistentFlags().StringSliceVarP(&bypassip, "bypassip", "b", []string{}, "the ip you will add behind some header like x-client-ip: 192.168.1.1,for example,-b 192.168.1.1,30.1.1.1 and the default value 127.0.0.1,localhost")
	rootCmd.PersistentFlags().StringVarP(&multiply, "multiply", "m", "", "the urls like http://192.168.10.2,http://192.168.11.2")

}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		//viper.AddConfigPath(os.Getenv("PWD"))
		viper.AddConfigPath("./")
		viper.SetConfigType("yaml")
		viper.SetConfigName("http")
	}
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Can't read config:", err)
		os.Exit(1)
	}
}
