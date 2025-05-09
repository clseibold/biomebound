package main

import (
	"bytes"
	"fmt"
	"os"
	"time"

	"gitlab.com/clseibold/biomebound"
	sis "gitlab.com/sis-suite/smallnetinformationservices"
)

func main() {
	context, _ := sis.InitConfiglessMode()

	// Check if "dev.pem" exists
	certPath := "dev.pem"
	if _, err := os.Stat(certPath); err == nil {

	} else if os.IsNotExist(err) {
		fmt.Printf("Generating new self-signed certificate.")
		// Generate new certificate and save it as "dev.pem"
		var pemBuffer bytes.Buffer
		sis.GenerateSelfSignedCertificateAndPrivateKey(sis.KeyType_ED25519, "localhost", []string{"localhost"}, "", time.Now().UTC().Add(time.Hour*24*365*200), &pemBuffer)
		os.WriteFile(certPath, pemBuffer.Bytes(), 0644)
	} else {
	}

	hosts := []sis.HostConfig{
		{BindAddress: "localhost", BindPort: "7000", Hostname: "localhost", Port: "7000", Upload: false, CertPath: certPath, SCGI: false},
		{BindAddress: "localhost", BindPort: "7000", Hostname: "localhost", Port: "7000", Upload: true, CertPath: certPath, SCGI: false},
	}
	geminiServer, _ := context.CreateServer(sis.ServerType_Gemini, "biomebound-dev", "en", hosts...)

	gameContext := biomebound.NewContext()
	gameContext.Start()
	gameContext.Attach(geminiServer)

	context.Start()
}
