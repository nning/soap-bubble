package main

import (
	"log"
	"os"
	"strconv"
	"strings"

	"gopkg.in/yaml.v3"
)

type Config struct {
	WindowResolution string `yaml:"windowResolution"`
	WindowWidth      int
	WindowHeight     int
	PixelResolution  string `yaml:"pixelResolution"`
	PixelWidth       int
	PixelHeight      int
	Fullscreen       bool `yaml:"fullscreen"`
	Debug            bool `yaml:"debug"`
	MaxBubbles       int  `yaml:"maxBubbles"`
	BubbleCollision  bool `yaml:"bubbleCollision"`
	EdgeCollision    bool `yaml:"edgeCollision"`
	ShowCursor       bool `yaml:"showCursor"`
	AutoSave         bool `yaml:"autoSave"`

	path string
}

func (c *Config) Load() *Config {
	data, err := os.ReadFile(c.path)
	if err != nil {
		return c
	}

	err = yaml.Unmarshal(data, c)
	if err != nil {
		log.Fatal(err)
	}

	res := strings.Split(c.WindowResolution, "x")
	if len(res) != 2 {
		log.Fatalf("invalid window resolution: %s", c.WindowResolution)
	}

	c.WindowWidth, err = strconv.Atoi(res[0])
	if err != nil {
		log.Fatalf("invalid window : %s", res[0])
	}

	c.WindowHeight, err = strconv.Atoi(res[1])
	if err != nil {
		log.Fatalf("invalid window height: %s", res[1])
	}

	res = strings.Split(c.PixelResolution, "x")
	if len(res) != 2 {
		log.Fatalf("invalid pixel resolution: %s", c.PixelResolution)
	}

	c.PixelWidth, err = strconv.Atoi(res[0])
	if err != nil {
		log.Fatalf("invalid pixel width: %s", res[0])
	}

	c.PixelHeight, err = strconv.Atoi(res[1])
	if err != nil {
		log.Fatalf("invalid pixel height: %s", res[1])
	}

	return c
}

func NewConfig() *Config {
	return &Config{
		WindowResolution: "1920x1080",
		WindowWidth:      1920,
		WindowHeight:     1080,
		PixelResolution:  "960x540",
		PixelWidth:       960,
		PixelHeight:      540,
		Fullscreen:       false,
		Debug:            false,
		MaxBubbles:       1,
		BubbleCollision:  true,
		EdgeCollision:    true,
		ShowCursor:       true,
		AutoSave:         true,

		path: "config.yml",
	}
}
