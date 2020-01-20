package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"

	"github.com/audibleblink/bothan/modules"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	hostsfile string
	masscan   bool
	verbose   bool

	rootCmd = &cobra.Command{
		Use:   "bothan host:ip",
		Short: "A tool to identify malicious-looking C2 servers online",
		Run:   execute,
	}
)

func init() {
	rootCmd.PersistentFlags().StringVarP(&hostsfile, "hostsfile", "f", "",
		"Newline-serperated list of host:port entries to check")
	rootCmd.PersistentFlags().BoolVarP(&masscan, "masscan", "", false,
		"Indicate whether input is from masscan (-oD output only)")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false,
		"Sets verbose logging")
}

func main() {
	rootCmd.Execute()
}

func execute(cmd *cobra.Command, args []string) {

	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})

	if verbose {
		log.SetLevel(log.DebugLevel)
	}

	var scanner *bufio.Scanner
	if hostsfile == "" {
		host := bytes.NewBufferString(args[0])
		scanner = bufio.NewScanner(host)
	}

	if hostsfile == "-" {
		scanner = bufio.NewScanner(os.Stdin)
	}

	if hostsfile != "" && hostsfile != "-" {
		file, err := os.Open(hostsfile)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		scanner = bufio.NewScanner(file)
	}

	for scanner.Scan() {
		host := scanner.Text()

		hostLogger := log.WithFields(log.Fields{
			"host": host,
		})

		if masscan {
			entry := &modules.MasscanRecord{}
			err := json.Unmarshal([]byte(host), entry)
			if err != nil {
				hostLogger.Error(err)
				continue
			}

			host = fmt.Sprintf("%s:%d", entry.IP, entry.Port)
		}

		req := newRequest("http", host, "GET")

		hostLogger.Debug("Requesting...")
		resp, err := doRequest(req)
		if err != nil {
			hostLogger.Error(err)
			continue
		}

		raw, err := httputil.DumpResponse(resp, false)
		if err != nil {
			hostLogger.Error(err)
			continue
		}

		rawReq := &modules.Query{
			Request:     req,
			RawResponse: raw,
		}

		successLogger := log.New()
		successLogger.SetFormatter(&log.TextFormatter{
			FullTimestamp: true,
		})
		successLogger.SetOutput(os.Stdout)

		for name, ost := range modules.Registry {
			if ost.IsBad(rawReq) {
				successLogger.WithFields(log.Fields{
					"tool": name,
					"host": rawReq.Request.URL.String(),
				}).Info("SUCCESS")
			}
		}
	}
}

func newRequest(scheme, hostAndPort, method string) *http.Request {
	uri := &url.URL{
		Scheme: scheme,
		Host:   hostAndPort,
	}

	return &http.Request{
		URL:    uri,
		Method: method,
	}
}

func doRequest(req *http.Request) (resp *http.Response, err error) {
	client := &http.Client{}
	resp, err = client.Do(req)
	if err != nil {
		if req.URL.Scheme == "https" {
			return
		}
		if terr := err.(*url.Error); terr != nil {
			log.Debug("Possible EOF. Re-trying with HTTPS: ", terr)
			req = newRequest("https", req.URL.Host, req.Method)
			return doRequest(req)
		}
		log.Error(err)
	}
	return
}
