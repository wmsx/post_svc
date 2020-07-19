package main

import (
	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/wmsx/post_svc/handler"
	"github.com/wmsx/post_svc/models"
	"github.com/wmsx/post_svc/setting"

	categoryProto "github.com/wmsx/post_svc/proto/category"
	postProto "github.com/wmsx/post_svc/proto/post"
)

const name = "wm.sx.svc.post"

func main() {
	service := micro.NewService(
		micro.Name(name),
		micro.Flags(
			&cli.StringFlag{
				Name:    "env",
				Usage:   "指定运行环境",
				Value:   "dev",
				EnvVars: []string{"RUN_ENV"},
			},
		),
	)

	var env string
	service.Init(
		micro.Action(func(c *cli.Context) error {
			env = c.String("env")
			return nil
		}),

		micro.BeforeStart(func() (err error) {
			err = setting.SetUp(name, env)
			models.SetUp()
			return err
		}),
		micro.AfterStop(func() error {
			models.CloseDB()
			return nil
		}),
	)

	_ = postProto.RegisterPostHandler(service.Server(), new(handler.PostHandler))
	_ = categoryProto.RegisterCategoryHandler(service.Server(), new(handler.CategoryHandler))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
