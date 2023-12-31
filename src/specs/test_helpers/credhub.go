package test_helpers

import (
	"crypto/tls"
	"fmt"
	"os"

	"code.cloudfoundry.org/credhub-cli/credhub"
	"code.cloudfoundry.org/credhub-cli/credhub/auth"
)

func NewCredhubClient() (*credhub.CredHub, error) {
	uaaCreds := auth.UaaClientCredentials(
		os.Getenv("CREDHUB_CLIENT"),
		os.Getenv("CREDHUB_SECRET"),
	)

	chClient, err := credhub.New(
		os.Getenv("CREDHUB_SERVER"),
		credhub.CaCerts(
			os.Getenv("CREDHUB_CA_CERT"),
		),
		credhub.SkipTLSValidation(true),
		credhub.Auth(uaaCreds),
	)

	return chClient, err
}

func credhubKey(name string) string {
	return fmt.Sprintf("%s/%s/%s", BoshCredhubPrefix, os.Getenv("BOSH_DEPLOYMENT"), name)
}

func GetDeploymentCertificateByName(name string) (tls.Certificate, error) {
	client, err := NewCredhubClient()
	if err != nil {
		return tls.Certificate{}, err
	}

	cert, err := client.GetLatestCertificate(credhubKey(name))
	if err != nil {
		return tls.Certificate{}, err
	}

	return tls.X509KeyPair([]byte(cert.Value.Certificate), []byte(cert.Value.PrivateKey))
}

func GetDeploymentCAByName(name string) (string, error) {
	client, err := NewCredhubClient()

	cert, err := client.GetLatestCertificate(credhubKey(name))
	if err != nil {
		return "", err
	}

	return cert.Value.Ca, err
}

func GetMySQLAdminPassword() (string, error) {
	client, err := NewCredhubClient()
	if err != nil {
		return "", err
	}
	pw, err := client.GetLatestPassword(credhubKey("cf_mysql_mysql_admin_password"))
	if err != nil {
		return "", err
	}

	return string(pw.Value), nil
}

func GetGaleraAgentPassword() (string, error) {
	client, err := NewCredhubClient()
	if err != nil {
		return "", err
	}
	pw, err := client.GetLatestPassword(credhubKey("cf_mysql_mysql_galera_healthcheck_endpoint_password"))
	if err != nil {
		return "", err
	}

	return string(pw.Value), nil
}

func GetProxyPassword() (string, error) {
	client, err := NewCredhubClient()
	if err != nil {
		return "", err
	}
	pw, err := client.GetLatestPassword(credhubKey("cf_mysql_proxy_api_password"))
	if err != nil {
		return "", err
	}

	return string(pw.Value), nil
}
