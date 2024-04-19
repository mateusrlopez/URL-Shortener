package settings

import "time"

type Settings struct {
	Router struct {
		Context string `default:"/api/v1"`
	}

	Server struct {
		Port         uint          `default:"8080"`
		ReadTimeout  time.Duration `default:"10s"`
		WriteTimeout time.Duration `default:"10s"`
	}

	Mongo struct {
		Uri               string        `required:"true"`
		ConnectionTimeout time.Duration `default:"10s"`
		Database          string        `required:"true"`
	}
}
