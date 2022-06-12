package fakemetadata_test

import (
	"testing"

	fakemetadata "github.com/zchee/gce-metadata-server/metadata/fake"
)

func TestIsTest(t *testing.T) {
	if ok := fakemetadata.IsTest(); !ok {
		t.Fatal("expected true")
	}
}
