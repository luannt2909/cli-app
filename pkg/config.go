package pkg

import (
	"cli-app/pkg/common"
	"cli-app/pkg/zingmp3"
	"os"
	"path"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

const (
	configFile               = "zingmp3_downloader.toml"
	defaultMP3Quality        = 320
	defaultMP4Quality        = 1080
	defaultDownloadDirectory = "zingmp3"
)

var HomeDirectory string

func init() {
	home, err := homedir.Dir()
	if err != nil {
		panic(err)
	}

	HomeDirectory = home
}

type Config struct {
	Mp3Quality int64  `json:"mp3_quality"`
	Mp4Quality int64  `json:"mp4_quality"`
	Directory  string `json:"directory"`
}

//Check existence of the configuration file
func IsConfigFileExist() bool {
	if fi, err := os.Stat(path.Join(HomeDirectory, configFile)); err != nil || fi.IsDir() {
		return false
	}
	return true
}

//Read the configuration
func ReadConfigFile() (*Config, error) {
	viper.SetConfigFile(path.Join(HomeDirectory, configFile))
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	return &Config{
		Mp3Quality: viper.GetInt64("config.mp3_quality"),
		Mp4Quality: viper.GetInt64("config.mp4_quality"),
		Directory:  viper.GetString("config.directory"),
	}, nil
}

//Check valid data of the configuration after read successfully
func (cfg *Config) IsValidConfig() error {
	if err := IsValidMP3Quality(cfg.Mp3Quality); err != nil {
		return err
	}

	if err := IsValidMP4Quality(cfg.Mp4Quality); err != nil {
		return err
	}

	return nil
}

func IsValidMP3Quality(quality int64) error {
	if quality != zingmp3.Normal && quality != zingmp3.VIP {
		return common.InvalidMp3Quality
	}
	return nil
}

func IsValidMP4Quality(quality int64) error {
	if quality != zingmp3.SD_360 && quality != zingmp3.SD_480 && quality != zingmp3.HD_720 && quality != zingmp3.FULL_HD_1080 {
		return common.InvalidMp4Quality
	}
	return nil
}

//Get absolutely path of download folder
func (cfg *Config) GetDownloadFolder() string {
	return path.Join(HomeDirectory, cfg.Directory)
}

//Make download directory if it not exist
func (cfg *Config) CreateDownloadFolderIfNotExist() {
	downloadFolder := cfg.GetDownloadFolder()
	if _, err := os.Stat(downloadFolder); err != nil {
		err := os.MkdirAll(downloadFolder, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}
}

//Write data of config to the configuration file
func WriteConfigFile(cfg *Config) error {
	_, err := os.Create(path.Join(HomeDirectory, configFile))
	if err != nil {
		return err
	}
	viper.SetConfigName("zingmp3_downloader")
	viper.AddConfigPath(HomeDirectory)
	viper.Set("config.mp3_quality", cfg.Mp3Quality)
	viper.Set("config.mp4_quality", cfg.Mp4Quality)
	viper.Set("config.directory", cfg.Directory)
	return viper.WriteConfig()
}

//Write default config into the configuration file
func WriteDefaultConfig() error {
	return WriteConfigFile(&Config{
		Mp3Quality: defaultMP3Quality,
		Mp4Quality: defaultMP4Quality,
		Directory:  defaultDownloadDirectory,
	})
}
