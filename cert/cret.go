package cert

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"log"
	"math/big"
	"os"
	"time"
)

var (
	rootCsr = &x509.Certificate{
		Version:      3,
		SerialNumber: big.NewInt(time.Now().Unix()),
		Subject: pkix.Name{
			Organization:       []string{"MiaoMint"},
			OrganizationalUnit: []string{"Sphinx"},
			CommonName:         "Sphinx Root CA",
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(10, 0, 0),
		BasicConstraintsValid: true,
		IsCA:                  true,
		MaxPathLen:            1,
		MaxPathLenZero:        false,
		KeyUsage:              x509.KeyUsageCertSign | x509.KeyUsageCRLSign,
	}

	rootKey   *ecdsa.PrivateKey
	rootCert  *x509.Certificate
	TlsConfig = &tls.Config{}
)

func init() {
	err := os.MkdirAll("certs", 0755)
	if err != nil {
		log.Fatalln(err)
	}

	if _, err := os.Stat("certs/root.crt"); err == nil {
		rootCert, rootKey, _, err = LoadCertFromPem("root")
		if err != nil {
			log.Fatalln(err)
		}
	} else {
		rootKey = GeneratePrivateKey()
		rootDer, err := x509.CreateCertificate(rand.Reader, rootCsr, rootCsr, rootKey.Public(), rootKey)
		if err != nil {
			log.Fatalln(err)
		}
		rootCert, err = x509.ParseCertificate(rootDer)
		if err != nil {
			log.Fatalln(err)
		}
	}

	os.WriteFile("certs/root.crt", []byte(ParseCertPEM(rootCert)), 0644)
	os.WriteFile("certs/root.key", []byte(ParseKeyPEM(rootKey)), 0644)
}

func GeneratePrivateKey() (key *ecdsa.PrivateKey) {
	key, _ = ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	return
}

func GenerateClientCert(domain string, ips ...string) error {
	csr := &x509.Certificate{
		Version:      3,
		SerialNumber: big.NewInt(time.Now().Unix()),
		Subject: pkix.Name{
			Organization:       []string{"MiaoMint"},
			OrganizationalUnit: []string{"Sphinx"},
			CommonName:         domain,
		},
		DNSNames:              append(ips, domain),
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(1, 0, 0),
		BasicConstraintsValid: true,
		IsCA:                  false,
		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageKeyEncipherment,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
	}
	key := GeneratePrivateKey()
	der, err := x509.CreateCertificate(rand.Reader, csr, rootCert, key.Public(), rootKey)
	if err != nil {
		return err
	}
	cert, err := x509.ParseCertificate(der)
	if err != nil {
		return err
	}

	os.WriteFile("certs/"+domain+".crt", []byte(ParseCertPEM(cert)), 0644)
	os.WriteFile("certs/"+domain+".key", []byte(ParseKeyPEM(key)), 0644)

	return nil
}

func RemoveCert(domain string) error {
	err := os.Remove(GetCertPath(domain))
	if err != nil {
		return err
	}
	err = os.Remove(GetKeyPath(domain))
	if err != nil {
		return err
	}
	return nil
}

func ParseCertPEM(cert *x509.Certificate) string {
	pemData := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: cert.Raw})
	return string(pemData)
}

func ParseKeyPEM(key *ecdsa.PrivateKey) string {
	pemData, err := x509.MarshalECPrivateKey(key)
	if err != nil {
		log.Fatalln(err)
	}
	return string(pem.EncodeToMemory(&pem.Block{Type: "EC PRIVATE KEY", Bytes: pemData}))
}

func GetCertPath(domain string) string {
	return "certs/" + domain + ".crt"
}

func GetKeyPath(domain string) string {
	return "certs/" + domain + ".key"
}

func LoadCertFromPem(domain string) (cert *x509.Certificate, key *ecdsa.PrivateKey, tlsCert tls.Certificate, err error) {
	tlsCert, err = tls.LoadX509KeyPair(GetCertPath(domain), GetKeyPath(domain))
	if err != nil {
		return
	}
	cert, err = x509.ParseCertificate(tlsCert.Certificate[0])
	if err != nil {
		return
	}
	key = tlsCert.PrivateKey.(*ecdsa.PrivateKey)
	return
}

func LoadAllCerts() (certs []tls.Certificate, err error) {
	files, err := os.ReadDir("certs")
	if err != nil {
		return
	}
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		if file.Name() == "root.crt" || file.Name() == "root.key" {
			continue
		}

		var tlsCert tls.Certificate
		_, _, tlsCert, err = LoadCertFromPem(file.Name()[:len(file.Name())-4])
		if err != nil {
			return
		}
		certs = append(certs, tlsCert)
	}
	return
}

func RefreshCerts() error {
	certs, err := LoadAllCerts()
	if err != nil {
		return err
	}
	TlsConfig.Certificates = certs
	return nil
}
