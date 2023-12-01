package main

import (
	"log"

	"github.com/sesaquecruz/go-payment-processor/config"
	"github.com/sesaquecruz/go-payment-processor/di"
	"github.com/sesaquecruz/go-payment-processor/internal/acquirer"
	"github.com/sesaquecruz/go-payment-processor/internal/infra/connection"
	"github.com/sesaquecruz/go-payment-processor/internal/infra/service"
)

//	@title			Payment Processor
//	@version		1.0.0
//	@description	A Rest API for Payment Processing.
//	@termsOfService	https://github.com/sesaquecruz/go-payment-processor

//	@contact.name	Support
//	@contact.url	https://github.com/sesaquecruz/go-payment-processor

//	@license.name	MIT
//	@license.url	https://github.com/sesaquecruz/go-payment-processor

//	@BasePath	/api/v1

func main() {
	cfg := config.GetConfig()

	db, err := connection.DBConnection(cfg.DbDsn)
	if err != nil {
		log.Fatal(err)
	}

	app := di.NewApp(
		db,
		service.PaymentWithAcquirer(acquirer.NewCielo(cfg.CieloUrl, cfg.CieloKey)),
		service.PaymentWithAcquirer(acquirer.NewRede(cfg.RedeUrl, cfg.RedeKey)),
		service.PaymentWithAcquirer(acquirer.NewStone(cfg.StoneUrl, cfg.StoneKey)),
	)

	app.Listen(":8080")
}
