package client

import (
	"context"
	"log"
	"log/slog"
	"net"
	"net/http"
	"os"

	"github.com/cushydigit/go-freelancer-sdk/freelancer"
	"github.com/joho/godotenv"
	"golang.org/x/net/proxy"
)

const (
	PROXY_ADDR   = "PROXY_ADDR"
	ACCESS_TOKEN = "FREELANCER_ACCESS_TOKEN"
)

func Init(l *slog.Logger, useSandBox bool) *freelancer.Client {

	if err := godotenv.Load(); err != nil {
		log.Fatalf("failed to load .env: %v", err)
	}

	accessToken := os.Getenv(ACCESS_TOKEN)
	if accessToken == "" {
		log.Fatalf("environment variable is not set: %s", ACCESS_TOKEN)
	}

	opts := []freelancer.ClientOption{
		freelancer.WithLogger(l),
	}

	if useSandBox {
		opts = append(opts, freelancer.WithSandBox())
	}

	return freelancer.NewClient(accessToken, opts...)
}

func InitWithProxy(l *slog.Logger, useSandBox bool) *freelancer.Client {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("failed to load .env: %v", err)
	}
	proxyAddr := os.Getenv(PROXY_ADDR)
	if proxyAddr == "" {
		log.Fatalf("environment variable is not set: %s", PROXY_ADDR)
	}

	accessToken := os.Getenv(ACCESS_TOKEN)
	if accessToken == "" {
		log.Fatalf("environment variable is not set: %s", ACCESS_TOKEN)
	}

	dialer, err := proxy.SOCKS5("tcp", proxyAddr, nil, proxy.Direct)
	if err != nil {
		log.Fatalf("failed to create SOCKS5 dialer: %v", err)
	}

	transport := &http.Transport{
		DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
			return dialer.Dial(network, addr)
		},
	}

	httpClient := &http.Client{
		Transport: transport,
	}

	opts := []freelancer.ClientOption{
		freelancer.WithHttpClient(httpClient),
		freelancer.WithLogger(l),
	}

	if useSandBox {
		opts = append(opts, freelancer.WithSandBox())
	}

	return freelancer.NewClient(accessToken, opts...)
}
