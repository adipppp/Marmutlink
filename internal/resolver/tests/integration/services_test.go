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

	link := "https://youtube.com/watch?v=RWFPC17YUhw"
	stream, err := resolver.Resolve(link)
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

	query := "23.exe tyfw"
	_, err := resolver.Resolve(query)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}

func TestResolveInvalidLink(t *testing.T) {
	resolver := setUp()

	invalidLink := "https://youtube.com/watch?v=INVALID_LINK"
	_, err := resolver.Resolve(invalidLink)
	if err == nil {
		t.Fatalf("Expected error for invalid link, got nil")
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

func TestResolveInvalidIdentifier(t *testing.T) {
	resolver := setUp()

	invalidIdentifier := "not_a_link_or_search"
	_, err := resolver.Resolve(invalidIdentifier)
	if err == nil {
		t.Fatalf("Expected error for invalid identifier, got nil")
	}
}
