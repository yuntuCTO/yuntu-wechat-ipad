package config

import (
	"flag"
	"fmt"
	"os"
	"time"
)

const (
	EnvDev  = "dev"
	EnvProd = "prod"
)

type Config struct {
	Addr            string
	Env             string
	ShutdownTimeout time.Duration
}

func Load() (Config, error) {
	addr := flag.String("addr", env("WXIPAD_ADDR", ":8080"), "监听地址")
	envName := flag.String("env", env("WXIPAD_ENV", EnvProd), "运行环境: dev | prod")
	shutdownTimeout := flag.Duration("shutdown-timeout", 10*time.Second, "优雅退出超时")
	flag.Parse()

	cfg := Config{
		Addr:            *addr,
		Env:             *envName,
		ShutdownTimeout: *shutdownTimeout,
	}
	if err := cfg.validate(); err != nil {
		return Config{}, err
	}
	return cfg, nil
}

func (c Config) validate() error {
	if c.Addr == "" {
		return fmt.Errorf("监听地址不能为空")
	}
	if c.Env != EnvDev && c.Env != EnvProd {
		return fmt.Errorf("非法运行环境 %q，可选 dev | prod", c.Env)
	}
	if c.ShutdownTimeout <= 0 {
		return fmt.Errorf("优雅退出超时必须大于 0")
	}
	return nil
}

func env(key, fallback string) string {
	if v, ok := os.LookupEnv(key); ok && v != "" {
		return v
	}
	return fallback
}
