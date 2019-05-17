package divvycloud

import (
	"github.com/hashicorp/terraform/helper/schema"
	"testing"
)

func TestProvider(t *testing.T) {
	if err := Provider().(*schema.Provider).InternalValidate(); err != nil {
		t.Errorf("failed to validate: %s", err)
	}
}
