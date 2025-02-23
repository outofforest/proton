package types_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/outofforest/proton/test/pkg1"
	"github.com/outofforest/proton/test/pkg1/spkg"
)

func TestMarshaller(t *testing.T) {
	requireT := require.New(t)

	m := pkg1.NewMarshaller(100)

	msg1 := &pkg1.MsgMapString{
		Value: map[string]string{},
	}

	for i := range 128 {
		msg1.Value[fmt.Sprintf("key-%d", i)] = fmt.Sprintf("value-%d", i)
	}

	size, err := m.Size(msg1)
	requireT.NoError(err)
	b := make([]byte, size)
	msgID, msgSize, err := m.Marshal(msg1, b)
	requireT.NoError(err)
	requireT.Equal(size, msgSize)

	msg2, msgSize2, err := m.Unmarshal(msgID, b)
	requireT.NoError(err)
	requireT.Equal(size, msgSize2)
	requireT.Equal(msg1, msg2.(*pkg1.MsgMapString))
}

func TestMarshallerUnknownType(t *testing.T) {
	requireT := require.New(t)

	m := pkg1.NewMarshaller(100)

	b := make([]byte, 100)
	msgID, msgSize, err := m.Marshal(&spkg.SubMsg{}, b)
	requireT.Error(err)
	requireT.Equal(uint64(0), msgSize)
	requireT.Equal(uint64(0), msgID)

	msg2, msgSize2, err := m.Unmarshal(1000000000000000, b)
	requireT.Error(err)
	requireT.Nil(msg2)
	requireT.Equal(uint64(0), msgSize2)
}
