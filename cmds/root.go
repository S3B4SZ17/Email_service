package cmds

import (
	"fmt"
	"os"

	"github.com/S3B4SZ17/Email_service/app"
	config "github.com/S3B4SZ17/Email_service/config"
	"github.com/S3B4SZ17/Email_service/management"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var rootCmd = &cobra.Command{
	Use:   "Email_service client",
	Short: "A simple email service client that just works",
	Long: `A simple email service client that just works.
				  Complete documentation is available at http://github.com/S3B4SZ17/Email_service`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		cfgFile, err := cmd.Flags().GetString("config")

		if err != nil {
			cmd.Usage()
			management.Log.Error(err.Error())
		} else {
			cfgObj, _ := config.LoadConfig(cfgFile)
			management.InitializeZapCustomLogger()
			management.Log.Info("Loaded config initial configuration", zap.String("configFile", cfgFile))
			management.Log.Info("Email_service configuration", zap.String("email", cfgObj.Smtp_server.Email_from), zap.Int("port", cfgObj.Smtp_server.Port), zap.String("host_url", cfgObj.Smtp_server.Host_url))
			app.StartApp(&cfgObj)
		}

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	var cfgFile string
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is ./app.env)")
	// rootCmd.PersistentFlags().StringVarP(&projectBase, "projectbase", "b", "", "base project directory eg. github.com/spf13/")
	// rootCmd.PersistentFlags().StringP("author", "a", "YOUR NAME", "Author name for copyright attribution")
	// rootCmd.PersistentFlags().StringVarP(&userLicense, "license", "l", "", "Name of license for the project (can provide `licensetext` in config)")
	// rootCmd.PersistentFlags().Bool("viper", true, "Use Viper for configuration")
	// viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
	// viper.BindPFlag("projectbase", rootCmd.PersistentFlags().Lookup("projectbase"))
	// viper.BindPFlag("useViper", rootCmd.PersistentFlags().Lookup("viper"))
	// viper.SetDefault("author", "NAME HERE <EMAIL ADDRESS>")
	// viper.SetDefault("license", "apache")
}
