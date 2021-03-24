package learn_go

import (
	"context"
	"errors"
	"fmt"
	"net"
	"runtime"
	"testing"
	"time"

	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttputil"
	"golang.org/x/sync/errgroup"
)

var defaultClientsCount = runtime.NumCPU()

func Test(t *testing.T) {
	run := func(t *testing.T, requestsCount int) {
		t.Run("sync", func(t *testing.T) {
			start := time.Now()
			benchmarkSync(t, defaultClientsCount, requestsCount)
			fmt.Println("-", t.Name(), ": ", time.Since(start))
		})

		t.Run("async", func(t *testing.T) {
			start := time.Now()
			benchmarkAsync(t, defaultClientsCount, requestsCount)
			fmt.Println("-", t.Name(), ": ", time.Since(start))
		})
	}

	t.Run("100", func(t *testing.T) { run(t, 100) })
	t.Run("500", func(t *testing.T) { run(t, 500) })
	t.Run("1k", func(t *testing.T) { run(t, 1000) })
	t.Run("10k", func(t *testing.T) { run(t, 10000) })
}

func benchmark(
	t testing.TB,
	ctx context.Context,
	client *fasthttp.HostClient,
	clientsCount int,
	requestsCount int,
) {
	defer ctx.Done()

	for request := 0; request < requestsCount; request++ {
		eg := errgroup.Group{}

		for sessionId := 0; sessionId < clientsCount; sessionId++ {
			eg.Go(
				func() error {
					statusCode, _, err := client.Get(nil, "http://localhost/ping")

					if statusCode != fasthttp.StatusOK {
						if nil == err {
							return errors.New("bad response")
						}
					}

					return err
				},
			)
		}

		if err := eg.Wait(); nil != err {
			panic(err)
		}
	}
}

func benchmarkSync(tb testing.TB, clientsCount int, requestsCount int) {
	ctx := context.Background()

	apiListener := fasthttputil.NewInmemoryListener()
	apiClient := &fasthttp.HostClient{
		Addr: "localhost:8181",
		Dial: func(addr string) (net.Conn, error) {
			return apiListener.Dial()
		},
	}

	soaListener := fasthttputil.NewInmemoryListener()
	soaClient := &fasthttp.HostClient{
		Addr: "localhost:8282",
		Dial: func(addr string) (net.Conn, error) {
			return soaListener.Dial()
		},
	}

	defer func() {
		_ = apiListener.Close()
		_ = soaListener.Close()
	}()

	// start SOA service
	soaServer := &fasthttp.Server{
		Handler: func(ctx *fasthttp.RequestCtx) {
			ctx.SetBody([]byte(`pong`))
		},
	}

	go func() { _ = soaServer.Serve(soaListener) }()

	// start API service
	apiServer := &fasthttp.Server{
		Handler: func(ctx *fasthttp.RequestCtx) {
			statusCode, body, err := soaClient.Get(nil, "http://localhost/ping")
			if nil != err {
				ctx.SetStatusCode(fasthttp.StatusInternalServerError)
			} else {
				ctx.SetStatusCode(statusCode)
				ctx.SetBody(body)
			}
		},
	}

	go func() { _ = apiServer.Serve(apiListener) }()

	benchmark(tb, ctx, apiClient, clientsCount, requestsCount)
}

func benchmarkAsync(t testing.TB, clientsCount int, requestsCount int) {
	ctx := context.Background()
	apiListener := fasthttputil.NewInmemoryListener()

	type (
		requestStruct struct {
			body     []byte
			response chan []byte
		}
	)

	bus := make(chan requestStruct, 1024)

	defer func() {
		_ = apiListener.Close()
	}()

	// start SOA service
	go func(ctx context.Context) {
		for {
			select {
			case req := <-bus:
				req.response <- []byte("pong")
				continue

			case <-ctx.Done():
				return
			}
		}
	}(ctx)

	// start API service
	apiServer := &fasthttp.Server{
		Handler: func(ctx *fasthttp.RequestCtx) {
			req := requestStruct{
				body:     nil,
				response: make(chan []byte),
			}

			bus <- req

			select {
			case out := <-req.response:
				ctx.SetStatusCode(fasthttp.StatusOK)
				ctx.SetBody(out)

			case <-time.After(100 * time.Millisecond):
				ctx.SetStatusCode(fasthttp.StatusInternalServerError)
			}
		},
	}

	go func() { _ = apiServer.Serve(apiListener) }()

	benchmark(
		t,
		ctx,
		&fasthttp.HostClient{
			Addr: "localhost:8181",
			Dial: func(addr string) (net.Conn, error) {
				return apiListener.Dial()
			},
		},
		clientsCount, // number of sessions
		requestsCount,
	)
}
