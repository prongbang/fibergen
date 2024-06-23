package template

import "strings"

type configurationTemplate struct {
}

func (c *configurationTemplate) Text() []byte {
	lines := []string{
		"package configuration",
		"",
		"import (",
		`	"fmt"`,
		`	"github.com/fsnotify/fsnotify"`,
		`	"github.com/spf13/viper"`,
		`	"log"`,
		")",
		"",
		"type configuration struct {",
		"	Env string `mapstructure:\"env\"`",
		"	API struct {",
		"		Port int `mapstructure:\"port\"`",
		"	} `mapstructure:\"api\"`",
		"	Role struct {",
		"		Admin string `mapstructure:\"admin\"`",
		"		User  string `mapstructure:\"user\"`",
		"	} `mapstructure:\"role\"`",
		"	Jwt struct {",
		"		Secret string `mapstructure:\"secret\"`",
		"	} `mapstructure:\"jwt\"`",
		"	Casbin struct {",
		"		Model  string `mapstructure:\"model\"`",
		"		Policy string `mapstructure:\"policy\"`",
		"	} `mapstructure:\"casbin\"`",
		"	Mongodb struct {",
		"		Host     string `mapstructure:\"host\"`",
		"		Port     int    `mapstructure:\"port\"`",
		"		Database string `mapstructure:\"database\"`",
		"		User     string `mapstructure:\"user\"`",
		"		Pass     string `mapstructure:\"pass\"`",
		"	} `mapstructure:\"mongodb\"`",
		"}",
		"",
		"var Config configuration",
		"",
		"func Load(env string) {",
		"	viper.SetConfigName(env)",
		`	viper.AddConfigPath("configuration")`,
		"	if err := viper.ReadInConfig(); err != nil {",
		`		log.Fatalf("[ERROR] Reading configuration file, %s", err)`,
		"	}",
		"	viper.OnConfigChange(func(e fsnotify.Event) {",
		`		fmt.Println("[INFO] Configuration file changed:", e.Name)`,
		"	})",
		"	viper.WatchConfig()",
		"	err := viper.Unmarshal(&Config)",
		"	if err == nil {",
		`		fmt.Println("[INFO] Configuration file has been loaded.")`,
		"	}",
		"}",
	}

	return []byte(strings.Join(lines, "\n"))
}

func ConfigurationTemplate() Template {
	return &configurationTemplate{}
}
