package helper

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/micro/cli"
	"github.com/micro/go-micro/metadata"
)

func ACMEHosts(ctx *cli.Context) []string {
	var hosts []string
	for _, host := range strings.Split(ctx.String("acme_hosts"), ",") {
		if len(host) > 0 {
			hosts = append(hosts, host)
		}
	}
	return hosts
}

func RequestToContext(r *http.Request) context.Context {
	ctx := context.Background()
	md := make(metadata.Metadata)
	for k, v := range r.Header {
		md[k] = strings.Join(v, ",")
	}
	return metadata.NewContext(ctx, md)
}

func TLSConfig(ctx *cli.Context) (*tls.Config, error) {
	cert := ctx.GlobalString("tls_cert_file")
	key := ctx.GlobalString("tls_key_file")
	ca := ctx.GlobalString("tls_client_ca_file")

	if len(cert) > 0 && len(key) > 0 {
		certs, err := tls.LoadX509KeyPair(cert, key)
		if err != nil {
			return nil, err
		}

		if len(ca) > 0 {
			caCert, err := ioutil.ReadFile(ca)
			if err != nil {
				return nil, err
			}

			caCertPool := x509.NewCertPool()
			caCertPool.AppendCertsFromPEM(caCert)

			return &tls.Config{
				Certificates: []tls.Certificate{certs},
				ClientCAs:    caCertPool,
				ClientAuth:   tls.RequireAndVerifyClientCert,
			}, nil
		}

		return &tls.Config{
			Certificates: []tls.Certificate{certs},
		}, nil
	}

	return nil, errors.New("TLS certificate and key files not specified")
}
