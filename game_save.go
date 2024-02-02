package main

import (
	"os"
	"path/filepath"
	"time"

	"github.com/fxamacker/cbor/v2"
)

func (g *Game) SavePeriodically() {
	for {
		time.Sleep(time.Second / 2)
		g.Save()
	}
}

func (g *Game) SavePath() (string, error) {
	bin, err := os.Executable()
	if err != nil {
		return "", err
	}

	return filepath.Join(filepath.Dir(bin), ".save"), nil
}

func (g *Game) Save() error {
	b, err := cbor.Marshal(g)
	if err != nil {
		return err
	}

	path, err := g.SavePath()
	if err != nil {
		return err
	}

	err = os.WriteFile(path, b, 0600)
	if err != nil {
		return err
	}

	return nil
}

func (g *Game) Load() error {
	path, err := g.SavePath()
	if err != nil {
		return err
	}

	b, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	err = cbor.Unmarshal(b, g)
	if err != nil {
		return err
	}

	return nil
}
