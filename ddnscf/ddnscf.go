package main

import (
	"flag"
	"log"
	"os"

	"nontajid.com/cloudflareclient"
	"nontajid.com/ddns"
)

func main() {
	cloudFlareEmailEnvVar := "CLOUDFLARE_EMAIL"
	cloudFlareApiKeyEnvVar := "CLOUDFLARE_API_KEY"

	targetIp := flag.String("ip", "", "Ip address that you want to update")
	tld := flag.String("tld", "", "Top level domain name that you want to update")
	domain := flag.String("domain", "", "targetDomain that you want to update")
	email := flag.String("email", "", "email registered with cloudflare")
	apiKey := flag.String("apiKey", "", "cloudflare api key")

	flag.Parse()

	// Fallback to Environment variable if not provide by cli
	if *email == "" {
		*email = os.Getenv(cloudFlareEmailEnvVar)
	}
	if *apiKey == "" {
		*apiKey = os.Getenv(cloudFlareApiKeyEnvVar)
	}

	if *targetIp == "" {
		log.Fatalln("Ip Address is required field")
	}
	if *tld == "" {
		log.Fatalln("tld is required field")
	}
	if *domain == "" {
		log.Fatalln("domain is required field")
	}
	if *email == "" {
		log.Fatalln("email is required field")
	}
	if *apiKey == "" {
		log.Fatalln("apiKey is required field")
	}

	recordInfo := cloudflareclient.NewRecordInfo{
		Ip:     *targetIp,
		Tld:    *tld,
		Domain: *domain,
	}

	cloudflareConnectionInfo := cloudflareclient.CloudFlareConnectionInfo{
		Email:  *email,
		ApiKey: *apiKey,
	}

	connection := ddns.CreateCloudFlareUpdater(cloudflareConnectionInfo)
	ddns.UpdateIp(recordInfo, connection)
}
