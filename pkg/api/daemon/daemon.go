//
// Last.Backend LLC CONFIDENTIAL
// __________________
//
// [2014] - [2017] Last.Backend LLC
// All Rights Reserved.
//
// NOTICE:  All information contained herein is, and remains
// the property of Last.Backend LLC and its suppliers,
// if any.  The intellectual and technical concepts contained
// herein are proprietary to Last.Backend LLC
// and its suppliers and may be covered by Russian Federation and Foreign Patents,
// patents in process, and are protected by trade secret or copyright law.
// Dissemination of this information or reproduction of this material
// is strictly forbidden unless prior written permission is obtained
// from Last.Backend LLC.
//

package daemon

import (
	_cfg "github.com/lastbackend/lastbackend/pkg/common/config"

	"github.com/lastbackend/lastbackend/pkg/api/config"
	"github.com/lastbackend/lastbackend/pkg/api/context"
	"github.com/lastbackend/lastbackend/pkg/api/events"
	"github.com/lastbackend/lastbackend/pkg/api/http"
	"github.com/lastbackend/lastbackend/pkg/log"
	"github.com/lastbackend/lastbackend/pkg/sockets"
	"github.com/lastbackend/lastbackend/pkg/storage"
	"os"
	"os/signal"
	"syscall"
)

const app = "api"

func Daemon(_cfg *_cfg.Config) {

	var (
		ctx  = context.Get()
		cfg  = config.Set(_cfg)
		sigs = make(chan os.Signal)
		done = make(chan bool, 1)
	)

	log.New(app, *cfg.LogLevel)
	log.Info("Start API server")

	ctx.SetConfig(cfg)

	stg, err := storage.Get(cfg.GetEtcdDB())
	if err != nil {
		panic(err)
	}
	ctx.SetStorage(stg)
	ctx.SetWssHub(sockets.NewHub())

	go func() {
		events.NewEventListener().Listen()
	}()

	go func() {
		if err := http.Listen(*cfg.APIServer.Host, *cfg.APIServer.Port); err != nil {
			log.Warnf("Http server start error: %s", err)
		}
	}()

	// Handle SIGINT and SIGTERM.
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		for {
			select {
			case <-sigs:
				done <- true
				return
			}
		}
	}()

	<-done
}
