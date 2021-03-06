package main

// import "C"
import (
	// "github.com/kennyzhu/go-os/dbservice/subscriber"

	"github.com/micro/go-micro"

	_ "github.com/kennyzhu/go-os/dbservice/conf"
	"github.com/kennyzhu/go-os/log"
	"github.com/kennyzhu/go-os/dbservice/handler"
	"github.com/kennyzhu/go-os/dbservice/conf"
	"github.com/kennyzhu/go-os/dbservice/models"
)

// generated by: micro new --type srv --alias dbservice github.com/kennyzhu/go-os/
func main() {
	// New Service
	service := micro.NewService(
		micro.Name(conf.AppConf.SrvName),
		micro.Version("1"),
	)

	// Initialise service
	service.Init()

	// Initialise models
	models.Init(false)

	// Register Handler,each goroutine..
	handler.Init( service.Server() )

	// Register Struct as Subscriber
	// micro.RegisterSubscriber("go.micro.srv.template", service.Server(), new(subscriber.Example))
	// micro.RegisterSubscriber("go.micro.srv.template", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
