package config

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Config the asd.json file should be set at the root level
type Config struct {
	SERVER struct {
		Host string `mapstructure:"host"`
		Port int    `mapstructure:"port"`
	}

	DB struct {
		Driver      string `mapstructure:"driver"`
		URL         string `mapstructure:"url"`
		Automigrate bool   `mapstructure:"automigrate"`
	}
}

var cfgFile string

func LoadTestConfig(path string) (*Config, error) {
	cmd := &cobra.Command{}
	cmd.Flags().StringVar(&cfgFile, "config", path, "Config file")

	return LoadConfig(cmd)
}

// LoadConfig should load and unmarshal the config file
func LoadConfig(cmd *cobra.Command) (*Config, error) {
	err := viper.BindPFlags(cmd.Flags())
	if err != nil {
		return nil, err
	}

	viper.SetEnvPrefix("GOKIT")

	if configFile, _ := cmd.Flags().GetString("config"); configFile != "" {
		viper.SetConfigFile(configFile)
	} else {
		viper.AddConfigPath("./config")
		viper.SetConfigName("config")
	}

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	viper.AutomaticEnv() // read in environment variables that match

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}

	config := new(Config)

	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	fmt.Println(config)

	return config, nil
}
