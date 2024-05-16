// Package clickhouseinfra provides a set of functions to interact with ClickHouse containers.
package clickhouseinfra

import (
	"bytes"
	"context"
	"crypto/tls"
	"crypto/x509"
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/docker/go-connections/nat"
	"github.com/mdelapenya/tlscert"
	"github.com/testcontainers/testcontainers-go"
	chmodule "github.com/testcontainers/testcontainers-go/modules/clickhouse"
)

const (
	defaultUser      = "default"
	defaultDB        = "dimo"
	SecureNativePort = nat.Port("9440/tcp")
)

var secureConfigXML = []byte(`
<clickhouse>
	<tcp_port_secure>9440</tcp_port_secure>
	<openSSL>
		<server>
			<certificateFile>/etc/clickhouse-server/certs/client.crt</certificateFile>
			<privateKeyFile>/etc/clickhouse-server/certs/client.key</privateKeyFile>
			<verificationMode>relaxed</verificationMode>
			<caConfig>/etc/clickhouse-server/certs/ca.crt</caConfig>
		</server>
	</openSSL>
</clickhouse>
`)

// Container is a struct that holds the clickhouse and zookeeper containers.
type Container struct {
	*chmodule.ClickHouseContainer
	RootCAs *x509.CertPool
}

// CreateClickHouseContainer function starts and testcontainer for clickhouse.
// The caller is responsible for terminating the container.
func CreateClickHouseContainer(ctx context.Context, userName, password string) (*Container, error) {
	if userName == "" {
		userName = defaultUser
	}
	caCert, clientCerts, err := createCert()
	if err != nil {
		return nil, fmt.Errorf("failed to create certs: %w", err)
	}
	clickHouseContainer, err := chmodule.RunContainer(ctx,
		testcontainers.WithImage("clickhouse/clickhouse-server:23.3.8.21-alpine"),
		chmodule.WithDatabase(defaultDB),
		chmodule.WithUsername(userName),
		chmodule.WithPassword(password),
		WithCerts(caCert, clientCerts),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to start clickhouse container: %w", err)
	}
	rootCAs, _ := x509.SystemCertPool()
	if rootCAs == nil {
		rootCAs = x509.NewCertPool()
	}

	// add our cert to the system pool
	rootCAs.AppendCertsFromPEM(caCert.Bytes)
	rootCAs.AppendCertsFromPEM(clientCerts.Bytes)
	rootCAs.AppendCertsFromPEM(clientCerts.KeyBytes)
	rootCAs.AppendCertsFromPEM(caCert.Cert.AuthorityKeyId)
	return &Container{
		ClickHouseContainer: clickHouseContainer,
		RootCAs:             rootCAs,
	}, nil
}

// GetClickHouseAsConn function returns a clickhouse.Conn connection which uses native ClickHouse protocol.
func (c *Container) GetClickHouseAsConn() (clickhouse.Conn, error) {
	host, err := c.ConnectionHost(context.TODO())
	if err != nil {
		return nil, fmt.Errorf("failed to get clickhouse host: %w", err)
	}
	conn, err := clickhouse.Open(&clickhouse.Options{
		Protocol: clickhouse.Native,
		Addr:     []string{host},
		Auth: clickhouse.Auth{
			Username: c.User,
			Password: c.Password,
			Database: c.DbName,
		},
		TLS: &tls.Config{
			RootCAs: c.RootCAs,
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to open clickhouse connection: %w", err)
	}
	return conn, nil
}

// GetClickhouseAsDB function returns a sql.DB connection which allows interfaceing with the stdlib database/sql package.
func (c *Container) GetClickhouseAsDB(ctx context.Context) (*sql.DB, error) {
	host, err := c.ConnectionHost(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get clickhouse host: %w", err)
	}
	dbConn := clickhouse.OpenDB(&clickhouse.Options{
		Protocol: clickhouse.Native,
		Addr:     []string{host},
		Auth: clickhouse.Auth{
			Username: c.User,
			Password: c.Password,
			Database: c.DbName,
		},
		TLS: &tls.Config{
			RootCAs: c.RootCAs,
		},
	})
	const retries = 3
	for i := 0; i < retries; i++ {
		err = dbConn.Ping()
		if err != nil {
			time.Sleep(500 * time.Millisecond)
			continue
		}
		return dbConn, nil
	}

	return nil, fmt.Errorf("failed to ping clickhouse after %d retries: %w", retries, err)
}

// ConnectionHost returns the host and port of the clickhouse container, using the default, native 9000 port, and
// obtaining the host and exposed port from the container
func (c *Container) ConnectionHost(ctx context.Context) (string, error) {
	host, err := c.Host(ctx)
	if err != nil {
		return "", err
	}

	port, err := c.MappedPort(ctx, SecureNativePort)
	if err != nil {
		return "", err
	}

	return host + ":" + port.Port(), nil
}

// Terminate function terminates the clickhouse containers.
// If an error occurs, it will be printed to stderr.
func (c *Container) Terminate(ctx context.Context) {
	if err := c.ClickHouseContainer.Terminate(ctx); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed to terminate clickhouse container: %v", err)
	}
}

func createCert() (*tlscert.Certificate, *tlscert.Certificate, error) {
	// Generate a certificate for localhost and save it to disk.
	caCert := tlscert.SelfSignedFromRequest(tlscert.Request{
		Name:              "ca-cert",
		Host:              "localhost",
		SubjectCommonName: "localhost",
		IsCA:              true,
		ValidFor:          time.Hour * 24 * 365 * 10,
	})
	if caCert == nil {
		return nil, nil, fmt.Errorf("failed to generate CA certificate")
	}

	cert := tlscert.SelfSignedFromRequest(tlscert.Request{
		Name:              "test-cert",
		SubjectCommonName: "test-cert",
		Host:              "localhost",
		Parent:            caCert,
		ValidFor:          time.Hour * 24 * 365,
	})
	if cert == nil {
		return nil, nil, fmt.Errorf("failed to generate client certificate")
	}

	return caCert, cert, nil
}

// WithCerts is a customize request option that adds the certificates to the clickhouse container.
func WithCerts(caCert, clientCerts *tlscert.Certificate) testcontainers.CustomizeRequestOption {
	return func(req *testcontainers.GenericContainerRequest) {
		req.ExposedPorts = append(req.ExposedPorts, SecureNativePort.Port())
		ca := testcontainers.ContainerFile{
			Reader:            bytes.NewReader(caCert.Bytes),
			ContainerFilePath: "/etc/clickhouse-server/certs/ca.crt",
			FileMode:          0o755,
		}
		cert := testcontainers.ContainerFile{
			Reader:            bytes.NewReader(clientCerts.Bytes),
			ContainerFilePath: "/etc/clickhouse-server/certs/client.crt",
			FileMode:          0o755,
		}
		key := testcontainers.ContainerFile{
			Reader:            bytes.NewReader(clientCerts.KeyBytes),
			ContainerFilePath: "/etc/clickhouse-server/certs/client.key",
			FileMode:          0o755,
		}
		config := testcontainers.ContainerFile{
			Reader:            bytes.NewReader(secureConfigXML),
			ContainerFilePath: "/etc/clickhouse-server/config.d/aconfig.xml",
			FileMode:          0o755,
		}
		req.Files = append(req.Files, ca, cert, key, config)
	}
}
