package config

type Conf struct {
	File        string  `yaml:"path"`
	Length      int     `yaml:"length"`
	Head        bool    `yaml:"head"`
	TargetPath  string  `yaml:"target"`
	SourceSheet string  `yaml:"sheet"`
	Output      Output  `yaml:"output"`
	Default     Default `yaml:"default"`
}

type Output struct {
	Sheet string `yaml:"sheet"`
}

type Default struct {
	Blank string `yaml:"blank"`
}
