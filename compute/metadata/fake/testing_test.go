package fakemetadata_test

import (
	"testing"

	fakemetadata "github.com/zchee/gce-metadata-server/compute/metadata/fake"
)

func TestOnTest(t *testing.T) {
	if ok := fakemetadata.OnTest(); !ok {
		t.Fatal("expected true")
	}
}
