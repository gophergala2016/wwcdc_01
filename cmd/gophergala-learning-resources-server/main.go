package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/go-swagger/go-swagger/spec"
	flags "github.com/jessevdk/go-flags"
	graceful "github.com/tylerb/graceful"

	"github.com/gophergala2016/wwcdc_01/restapi/operations"
)

// This file was generated by the swagger tool.
// Make sure not to overwrite this file after you generated it because all your edits would be lost!
// It would only be overwritten if you explicitly specify --include-main for the generate all or support commands
//go:generate swagger generate server -t ../.. -A GophergalaLearningResources

var opts struct {
	TLSHost           string         `long:"tls-host" description:"the IP to listen on for tls, when not specified it's the same as --host" env:"TLS_HOST"`
	TLSPort           int            `long:"tls-port" description:"the port to listen on for secure connections, defaults to a random value" env:"TLS_PORT"`
	TLSCertificate    flags.Filename `long:"tls-certificate" description:"the certificate to use for secure connections" required:"true" env:"TLS_CERTIFICATE"`
	TLSCertificateKey flags.Filename `long:"tls-key" description:"the private key to use for secure conections" required:"true" env:"TLS_PRIVATE_KEY"`
}

var chttp = http.NewServeMux()
var handler http.Handler

func ApiHandler(w http.ResponseWriter, r *http.Request) {
	if strings.Contains(r.URL.Path, "/api/") {
		//fmt.Println("api", r.URL.Path)
		handler.ServeHTTP(w, r)
	} else {
		//fmt.Println("static", r.URL.Path)
		chttp.ServeHTTP(w, r)
	}
}

func main() {
	swaggerSpec, err := spec.New(SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}

	parser := flags.NewParser(&opts, flags.Default)
	parser.ShortDescription = swaggerSpec.Spec().Info.Title
	parser.LongDescription = swaggerSpec.Spec().Info.Description

	api := operations.NewGophergalaLearningResourcesAPI(swaggerSpec)
	handler = configureAPI(api)

	for _, optsGroup := range api.CommandLineOptionsGroups {
		parser.AddGroup(optsGroup.ShortDescription, optsGroup.LongDescription, optsGroup.Options)
	}

	if _, err := parser.Parse(); err != nil {
		os.Exit(1)
	}

	httpsServer := &graceful.Server{Server: new(http.Server)}
	// Normal resources
	chttp.Handle("/", http.FileServer(http.Dir("public")))

	mux := http.NewServeMux()
	mux.HandleFunc("/", ApiHandler)
	httpsServer.Handler = mux
	httpsServer.TLSConfig = new(tls.Config)
	httpsServer.TLSConfig.NextProtos = []string{"http/1.1"}
	// https://www.owasp.org/index.php/Transport_Layer_Protection_Cheat_Sheet#Rule_-_Only_Support_Strong_Protocols
	httpsServer.TLSConfig.MinVersion = tls.VersionTLS11
	httpsServer.TLSConfig.Certificates = make([]tls.Certificate, 1)
	httpsServer.TLSConfig.Certificates[0], err = tls.LoadX509KeyPair(string(opts.TLSCertificate), string(opts.TLSCertificateKey))
	if err != nil {
		api.ServerShutdown()
		log.Fatal(err)
	}

	tlsListener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", opts.TLSHost, opts.TLSPort))
	if err != nil {
		api.ServerShutdown()
		log.Fatalln(err)
	}

	fmt.Printf("serving gophergala learning resources at https://%s:%d\n", opts.TLSHost, opts.TLSPort)

	wrapped := tls.NewListener(tcpKeepAliveListener{tlsListener.(*net.TCPListener)}, httpsServer.TLSConfig)
	if err := httpsServer.Serve(wrapped); err != nil {
		api.ServerShutdown()
		log.Fatalln(err)
	}
	go func() {

		<-httpsServer.StopChan()
		api.ServerShutdown()
	}()
}

// tcpKeepAliveListener is copied from the stdlib net/http package

// tcpKeepAliveListener sets TCP keep-alive timeouts on accepted
// connections. It's used by ListenAndServe and ListenAndServeTLS so
// dead TCP connections (e.g. closing laptop mid-download) eventually
// go away.
type tcpKeepAliveListener struct {
	*net.TCPListener
}

func (ln tcpKeepAliveListener) Accept() (c net.Conn, err error) {
	tc, err := ln.AcceptTCP()
	if err != nil {
		return
	}
	tc.SetKeepAlive(true)
	tc.SetKeepAlivePeriod(3 * time.Minute)
	return tc, nil
}
