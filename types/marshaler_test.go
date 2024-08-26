package types_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/outofforest/proton/test/pkg1"
	"github.com/outofforest/proton/test/pkg2"
)

func TestMarshaler(t *testing.T) {
	requireT := require.New(t)

	m := pkg1.NewMarshaler(100)

	msg1 := &pkg1.MsgMapString{
		Value: map[string]string{},
	}

	for i := range 128 {
		msg1.Value[fmt.Sprintf("key-%d", i)] = fmt.Sprintf("value-%d", i)
	}

	b := make([]byte, msg1.Size())
	msgID, msgSize, err := m.Marshal(msg1, b)
	requireT.NoError(err)
	requireT.Equal(msg1.Size(), msgSize)
	requireT.Equal(pkg1.IDMsgMapString, msgID)

	msg2, msgSize2, err := m.Unmarshal(pkg1.IDMsgMapString, b)
	requireT.NoError(err)
	requireT.Equal(msg1.Size(), msgSize2)
	requireT.Equal(msg1, msg2.(*pkg1.MsgMapString))
}

func TestMarshalerUnknownType(t *testing.T) {
	requireT := require.New(t)

	m := pkg1.NewMarshaler(100)

	b := make([]byte, 100)
	msgID, msgSize, err := m.Marshal(&pkg2.SubMsg{}, b)
	requireT.Error(err)
	requireT.Equal(uint64(0), msgSize)
	requireT.Equal(uint64(0), msgID)

	msg2, msgSize2, err := m.Unmarshal(1000000000000000, b)
	requireT.Error(err)
	requireT.Nil(msg2)
	requireT.Equal(uint64(0), msgSize2)
}
