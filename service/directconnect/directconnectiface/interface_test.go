// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package directconnectiface_test

import (
	"testing"

	"github.com/awslabs/aws-sdk-go/service/directconnect"
	"github.com/awslabs/aws-sdk-go/service/directconnect/directconnectiface"
	"github.com/stretchr/testify/assert"
)

func TestInterface(t *testing.T) {
	assert.Implements(t, (*directconnectiface.DirectConnectAPI)(nil), directconnect.New(nil))
}