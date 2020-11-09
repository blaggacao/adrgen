package adr

import (
	"fmt"
	"github.com/spf13/viper"
	"path/filepath"
)

const CONFIG_FILENAME = "adrgen.config"
const CONFIG_FORMAT = "yaml"

type Config struct {
	TargetDirectory string
	TemplateFilename string
	MetaParams []string
	Statuses []string
	DefaultStatus string
	IdDigitNumber int
}

func CreateConfigFile(config Config) error {
	viper.SetConfigName(CONFIG_FILENAME)
	viper.SetConfigType(CONFIG_FORMAT)
	viper.Set("directory", config.TargetDirectory)
	viper.Set("template_file", filepath.Join(config.TargetDirectory, config.TemplateFilename))
	viper.Set("default_meta", config.MetaParams)
	viper.Set("supported_statuses", config.Statuses)
	viper.Set("default_status", config.DefaultStatus)
	viper.Set("id_digit_number", config.IdDigitNumber)
	return viper.WriteConfigAs(CONFIG_FILENAME + ".yml")
}

func GetConfig(directory string) (Config, error) {
	viper.SetConfigName(CONFIG_FILENAME)
	viper.SetConfigType(CONFIG_FORMAT)
	viper.AddConfigPath(directory)
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil { // Handle errors reading the config file
		return DefaultConfig(), fmt.Errorf("Fatal error config file: %s \n", err)
	}
	return Config{
		TargetDirectory: viper.GetString("directory"),
		TemplateFilename: viper.GetString("template_file"),
		MetaParams: viper.GetStringSlice("default_meta"),
		Statuses: viper.GetStringSlice("supported_statuses"),
		DefaultStatus: viper.GetString("default_status"),
		IdDigitNumber: viper.GetInt("id_digit_number"),
	}, nil
}

func DefaultConfig() Config  {
	return Config{
		TemplateFilename: "",
		TargetDirectory: ".",
		Statuses: []string{"proposed", "accepted", "rejected", "superseeded", "amended", "deprecated"},
		DefaultStatus: "proposed",
		MetaParams: []string{},
		IdDigitNumber: 4,
	}
}