// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package main is the main binary for the API server.
package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"reflect"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	pb "github.com/kubeflow/metadata/api"
	"github.com/kubeflow/metadata/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	host          = flag.String("host", "localhost", "Hostname to listen on.")
	rpcPort       = flag.Int("rpc_port", 9090, "RPC serving port.")
	httpPort      = flag.Int("http_port", 8080, "HTTP serving port.")
	schemaRootDir = flag.String("schema_root_dir", "/usr/local/google/home/zhenghui/go/src/github.com/kubeflow/metadata/schema/alpha", "Root directory for the predefined schemas.")
)

var typeOfBytes = reflect.TypeOf([]byte(nil))

type rawJSONPb struct {
	*runtime.JSONPb
}

// NewDecoder checks if the data is raw bytes.
func (*rawJSONPb) NewDecoder(r io.Reader) runtime.Decoder {
	return runtime.DecoderFunc(func(v interface{}) error {
		rawData, err := ioutil.ReadAll(r)
		if err != nil {
			return err
		}
		rv := reflect.ValueOf(v)

		if rv.Kind() != reflect.Ptr {
			return fmt.Errorf("%T is not a pointer", v)
		}

		rv = rv.Elem()
		if rv.Type() != typeOfBytes {
			return fmt.Errorf("Type must be []byte but got %T", v)
		}

		rv.Set(reflect.ValueOf(rawData))
		return nil
	})
}

func main() {
	flag.Parse()
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	rpcEndpoint := fmt.Sprintf("%s:%d", *host, *rpcPort)

	service, err := service.NewService(*schemaRootDir)
	if err != nil {
		glog.Fatal(err)
	}
	rpcServer := grpc.NewServer()
	pb.RegisterMetadataServiceServer(rpcServer, service)

	go func() {
		listen, err := net.Listen("tcp", rpcEndpoint)
		if err != nil {
			glog.Fatal(err)
		}
		reflection.Register(rpcServer)
		if err := rpcServer.Serve(listen); err != nil {
			glog.Fatal(err)
		}
	}()

	jsonpb := new(runtime.JSONPb)
	mux := runtime.NewServeMux(
		// Add a JSON marshaler to bind the whole HTTP body as raw proto bytes.
		runtime.WithMarshalerOption("application/raw-json", &rawJSONPb{jsonpb}),
		runtime.WithMarshalerOption(runtime.MIMEWildcard, jsonpb),
		runtime.WithProtoErrorHandler(runtime.DefaultHTTPProtoErrorHandler),
	)

	opts := []grpc.DialOption{grpc.WithInsecure()}
	err = pb.RegisterMetadataServiceHandlerFromEndpoint(ctx, mux, rpcEndpoint, opts)
	if err != nil {
		glog.Fatal(err)
	}

	httpEndpoint := fmt.Sprintf("%s:%d", *host, *httpPort)
	glog.Infof("HTTP server listening on %s", httpEndpoint)
	http.ListenAndServe(httpEndpoint, mux)
}
