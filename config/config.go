package config

type cfg struct {
	System struct {
		Debug bool `yaml:"debug"`
		Port  int  `yaml:"port"`
	} `yaml:"system"`
	Misskey struct {
		Instance string `yaml:"instance"`
	} `yaml:"misskey"`
}

var Config cfg
