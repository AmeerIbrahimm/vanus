// Copyright 2022 Linkall Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/linkall-labs/vanus/internal/primitive"
	"github.com/linkall-labs/vanus/internal/trigger"
	"github.com/linkall-labs/vanus/internal/util/signal"
	"github.com/linkall-labs/vanus/observability/log"
	pbtrigger "github.com/linkall-labs/vanus/proto/pkg/trigger"
	"google.golang.org/grpc"
	"net"
	"os"
	"sync"
)

func main() {
	ctx := signal.SetupSignalContext()
	f := flag.String("config", "./config/trigger.yaml", "trigger worker config file path")
	flag.Parse()
	c, err := trigger.InitConfig(*f)
	if err != nil {
		log.Error(nil, "init config error", map[string]interface{}{log.KeyError: err})
		os.Exit(-1)
	}
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", c.Port))
	if err != nil {
		log.Error(context.Background(), "failed to listen", map[string]interface{}{
			"error": err,
		})
		os.Exit(-1)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	srv := trigger.NewTriggerServer(*c)
	pbtrigger.RegisterTriggerWorkerServer(grpcServer, srv)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		log.Info(ctx, "the grpc server ready to work", nil)
		err = grpcServer.Serve(listen)
		if err != nil {
			log.Error(ctx, "grpc server occurred an error", map[string]interface{}{
				log.KeyError: err,
			})
		}
	}()
	init := srv.(primitive.Initializer)
	if err = init.Initialize(ctx); err != nil {
		log.Error(ctx, "the trigger worker has initialized failed", map[string]interface{}{
			log.KeyError: err,
		})
		os.Exit(1)
	}
	<-ctx.Done()
	closer := srv.(primitive.Closer)
	closer.Close(ctx)
	grpcServer.GracefulStop()
	wg.Wait()
	log.Info(ctx, "trigger worker stopped", nil)
}
