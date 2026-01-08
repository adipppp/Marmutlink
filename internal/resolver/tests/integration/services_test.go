package integration_test

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"testing"

	resolver "adipppp/Marmutlink/internal/resolver/src"
)

func setUp() resolver.Resolver {
	return resolver.NewIDResolver()
}

func TestResolveYTLinkSuccess(t *testing.T) {
	resolver := setUp()

	identifier := "RWFPC17YUhw"
	stream, err := resolver.Resolve(identifier)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	data, err := io.ReadAll(stream)
	if err != nil {
		t.Fatalf("Expected no error reading stream, got %v", err)
	}

	sha256Sum := sha256.Sum256(data)
	sha256String := hex.EncodeToString(sha256Sum[:])

	expected_hash := "436d6ac6fd2fb420b4870e94c9cf1317b58725e0f6a011cdc365b8cbf49ac76b"

	if sha256String != expected_hash {
		t.Fatalf("Expected hash %s, got %s", expected_hash, sha256String)
	}
}

func TestResolveYTSearchSuccess(t *testing.T) {
	resolver := setUp()

	query := "ytsearch:23.exe tyfw"
	_, err := resolver.Resolve(query)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}

func TestResolveInvalidIdentifier(t *testing.T) {
	resolver := setUp()

	invalidIdentifier := "INVALID_LINK"
	_, err := resolver.Resolve(invalidIdentifier)
	if err == nil {
		t.Fatalf("Expected error for invalid identifier, got nil")
	}
}

func TestResolveInvalidSearch(t *testing.T) {
	resolver := setUp()

	invalidQuery := "ytsearch:asdlkfjasldkfjalksdjflkasjdf"
	_, err := resolver.Resolve(invalidQuery)
	if err == nil {
		t.Fatalf("Expected error for invalid search, got nil")
	}
}
