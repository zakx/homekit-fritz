package cmd

import (
	"github.com/chris-rock/homekit-fritz/homekit"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Starts the HomeKit Bridge for Fritz!Box",
	Long:  `Connects to Fritz!Box and publishes all devices in HomeKit`,
	Run: func(cmd *cobra.Command, args []string) {
		setLogging()

		hk := &homekit.HKConfig{
			Pin:     viper.GetString("homekit.pin"),
			SetupID: viper.GetString("homekit.setupid"),
		}
		fb := &homekit.FritzBoxConfig{
			Username: viper.GetString("fritzbox.username"),
			Password: viper.GetString("fritzbox.password"),
			URL:      viper.GetString("fritzbox.url"),
		}

		config := &homekit.Config{
			HomeKit:  hk,
			FritzBox: fb,
		}

		// print qr code
		homekit.Qrcode(hk)

		// start service
		homekit.Start(config)
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
