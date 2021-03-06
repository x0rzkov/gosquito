package core

import (
	log "github.com/livelace/logrus"
	"github.com/spf13/viper"
	"os"
	"os/user"
	"path/filepath"
	"runtime"
)

var (
	FLOW_CONF_DIR    = filepath.Join("flow", "conf")
	PLUGIN_DATA_DIR  = filepath.Join("plugin", "data")
	PLUGIN_STATE_DIR = filepath.Join("plugin", "state")
	PLUGIN_TEMP_DIR  = filepath.Join("plugin", "temp")
)

func initConfig() (string, error) {
	// Get current user info.
	userAccount, err := user.Current()
	if err != nil {
		return "", err
	}

	// Possible config paths.
	userDir := filepath.Join(userAccount.HomeDir, ".gosquito")
	configFile := "config.toml"

	// Return existing config path.
	if IsFile(DEFAULT_ETC_PATH, configFile) {
		return DEFAULT_ETC_PATH, nil

	} else if IsFile(userDir, configFile) {
		return userDir, nil

	} else if IsFile(DEFAULT_CURRENT_PATH, configFile) {
		return DEFAULT_CURRENT_PATH, nil
	}

	// Write config sample if config not found.
	if err := CreateDirIfNotExist(userDir); err != nil {
		return userDir, err
	}

	if err := WriteStringToFile(userDir, configFile, CONFIG_SAMPLE); err != nil {
		return userDir, err
	}

	return userDir, nil
}

func GetConfig() *viper.Viper {
	configPath, err := initConfig()

	if err != nil {
		log.WithFields(log.Fields{
			"path":  configPath,
			"error": err,
		}).Error(LOG_CONFIG_ERROR)
		os.Exit(1)
	}

	// Show user config path.
	if configPath != "" {
		log.WithFields(log.Fields{
			"path": configPath,
		}).Info(LOG_CONFIG_INIT)
	}

	// Read generated/existed configuration.
	v := viper.New()
	v.SetConfigName("config.toml")
	v.SetConfigType("toml")
	v.AddConfigPath(configPath)

	if err := v.ReadInConfig(); err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error(LOG_CONFIG_ERROR)
		os.Exit(1)
	}

	// Set defaults.
	v.SetDefault(VIPER_DEFAULT_EXPIRE_ACTION, make([]string, 0))
	v.SetDefault(VIPER_DEFAULT_EXPIRE_ACTION_DELAY, DEFAULT_EXPIRE_ACTION_DELAY)
	v.SetDefault(VIPER_DEFAULT_EXPIRE_ACTION_TIMEOUT, DEFAULT_EXPIRE_ACTION_TIMEOUT)
	v.SetDefault(VIPER_DEFAULT_EXPIRE_INTERVAL, DEFAULT_EXPIRE_INTERVAL)
	v.SetDefault(VIPER_DEFAULT_EXPORTER_LISTEN, DEFAULT_EXPORTER_LISTEN)
	v.SetDefault(VIPER_DEFAULT_FLOW_CONF, filepath.Join(configPath, FLOW_CONF_DIR))
	v.SetDefault(VIPER_DEFAULT_FLOW_ENABLE, make([]string, 0))
	v.SetDefault(VIPER_DEFAULT_FLOW_INTERVAL, DEFAULT_FLOW_INTERVAL)
	v.SetDefault(VIPER_DEFAULT_FLOW_NUMBER, DEFAULT_FLOW_NUMBER)
	v.SetDefault(VIPER_DEFAULT_LOG_LEVEL, DEFAULT_LOG_LEVEL)
	v.SetDefault(VIPER_DEFAULT_PLUGIN_DATA, filepath.Join(configPath, PLUGIN_DATA_DIR))
	v.SetDefault(VIPER_DEFAULT_PLUGIN_INCLUDE, DEFAULT_PLUGIN_INCLUDE)
	v.SetDefault(VIPER_DEFAULT_PLUGIN_STATE, filepath.Join(configPath, PLUGIN_STATE_DIR))
	v.SetDefault(VIPER_DEFAULT_PLUGIN_TEMP, filepath.Join(configPath, PLUGIN_TEMP_DIR))
	v.SetDefault(VIPER_DEFAULT_PLUGIN_TIMEOUT, DEFAULT_PLUGIN_TIMEOUT)
	v.SetDefault(VIPER_DEFAULT_PROC_NUM, runtime.GOMAXPROCS(0))
	v.SetDefault(VIPER_DEFAULT_TIME_FORMAT, DEFAULT_TIME_FORMAT)
	v.SetDefault(VIPER_DEFAULT_TIME_ZONE, DEFAULT_TIME_ZONE)
	v.SetDefault(VIPER_DEFAULT_USER_AGENT, DEFAULT_USER_AGENT)

	// Directories must exist for proper work.
	workDirs := []string{
		v.GetString(VIPER_DEFAULT_FLOW_CONF),
		v.GetString(VIPER_DEFAULT_PLUGIN_DATA),
		v.GetString(VIPER_DEFAULT_PLUGIN_STATE),
		v.GetString(VIPER_DEFAULT_PLUGIN_TEMP),
	}

	for _, workDir := range workDirs {
		if err := CreateDirIfNotExist(workDir); err != nil {
			log.WithFields(log.Fields{
				"path":  workDir,
				"error": err,
			}).Error(LOG_CONFIG_ERROR)
			os.Exit(1)
		}
	}

	return v
}
