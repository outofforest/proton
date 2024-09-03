package factory

import (
	"reflect"

	"github.com/pkg/errors"

	"github.com/outofforest/proton/types"
	"github.com/outofforest/proton/types/tarray"
	"github.com/outofforest/proton/types/tarrayfloat32"
	"github.com/outofforest/proton/types/tarrayfloat64"
	"github.com/outofforest/proton/types/tarrayint8"
	"github.com/outofforest/proton/types/tarrayuint8"
	"github.com/outofforest/proton/types/tbool"
	"github.com/outofforest/proton/types/tfloat32"
	"github.com/outofforest/proton/types/tfloat64"
	"github.com/outofforest/proton/types/tint16"
	"github.com/outofforest/proton/types/tint32"
	"github.com/outofforest/proton/types/tint64"
	"github.com/outofforest/proton/types/tint8"
	"github.com/outofforest/proton/types/tmap"
	"github.com/outofforest/proton/types/tslice"
	"github.com/outofforest/proton/types/tslicefloat32"
	"github.com/outofforest/proton/types/tslicefloat64"
	"github.com/outofforest/proton/types/tsliceint8"
	"github.com/outofforest/proton/types/tsliceuint8"
	"github.com/outofforest/proton/types/tstring"
	"github.com/outofforest/proton/types/tstruct"
	"github.com/outofforest/proton/types/tuint16"
	"github.com/outofforest/proton/types/tuint32"
	"github.com/outofforest/proton/types/tuint64"
	"github.com/outofforest/proton/types/tuint8"
)

// Get returns builder for particular type.
//
//nolint:gocyclo
func Get(msgType, fieldType reflect.Type, tm types.TypeMap) (types.BuilderFactory, error) {
	switch fieldType.Kind() {
	case reflect.Bool:
		return newConstantAdapter(tbool.New()), nil
	case reflect.Int8:
		return newConstantAdapter(tint8.New(msgType, fieldType, tm)), nil
	case reflect.Int16:
		return tint16.New(msgType, fieldType, tm), nil
	case reflect.Int32:
		return tint32.New(msgType, fieldType, tm), nil
	case reflect.Int64:
		return tint64.New(msgType, fieldType, tm), nil
	case reflect.Uint8:
		return newConstantAdapter(tuint8.New(msgType, fieldType, tm)), nil
	case reflect.Uint16:
		return tuint16.New(msgType, fieldType, tm), nil
	case reflect.Uint32:
		return tuint32.New(msgType, fieldType, tm), nil
	case reflect.Uint64:
		return tuint64.New(msgType, fieldType, tm), nil
	case reflect.Float32:
		return newConstantAdapter(tfloat32.New(msgType, fieldType, tm)), nil
	case reflect.Float64:
		return newConstantAdapter(tfloat64.New(msgType, fieldType, tm)), nil
	case reflect.Array:
		switch fieldType.Elem().Kind() {
		case reflect.Uint8:
			return newConstantAdapter(tarrayuint8.New(fieldType, tm)), nil
		case reflect.Int8:
			return newConstantAdapter(tarrayint8.New(fieldType, tm)), nil
		case reflect.Float32:
			return newConstantAdapter(tarrayfloat32.New(fieldType, tm)), nil
		case reflect.Float64:
			return newConstantAdapter(tarrayfloat64.New(fieldType, tm)), nil
		default:
			elementBuilder, err := Get(msgType, fieldType.Elem(), tm)
			if err != nil {
				return nil, err
			}
			return tarray.New(msgType, fieldType, elementBuilder), nil
		}
	case reflect.Map:
		keyBuilder, err := Get(msgType, fieldType.Key(), tm)
		if err != nil {
			return nil, err
		}

		elementBuilder, err := Get(msgType, fieldType.Elem(), tm)
		if err != nil {
			return nil, err
		}

		return tmap.New(msgType, fieldType, keyBuilder, elementBuilder, tm), nil
	case reflect.Slice:
		switch fieldType.Elem().Kind() {
		case reflect.Uint8:
			return tsliceuint8.New(msgType, fieldType, tm), nil
		case reflect.Int8:
			return tsliceint8.New(msgType, fieldType, tm), nil
		case reflect.Float32:
			return tslicefloat32.New(msgType, fieldType, tm), nil
		case reflect.Float64:
			return tslicefloat64.New(msgType, fieldType, tm), nil
		default:
			elementBuilder, err := Get(msgType, fieldType.Elem(), tm)
			if err != nil {
				return nil, err
			}

			return tslice.New(msgType, fieldType, elementBuilder, tm), nil
		}
	case reflect.String:
		return tstring.New(msgType, fieldType, tm), nil
	case reflect.Struct:
		fieldCount := fieldType.NumField()
		fieldBuilders := make([]types.BuilderFactory, 0, fieldCount)

		for i := range fieldCount {
			fieldType := fieldType.Field(i).Type
			fieldBuilder, err := Get(msgType, fieldType, tm)
			if err != nil {
				return nil, err
			}
			fieldBuilders = append(fieldBuilders, fieldBuilder)
		}

		return newNonConstantAdapter(tstruct.New(msgType, fieldType, fieldBuilders, tm)), nil
	default:
		return nil, errors.Errorf("unsupported type %s", fieldType.Name())
	}
}

func newConstantAdapter(builder types.BuilderFactoryConstant) constantAdapter {
	return constantAdapter{
		builder: builder,
	}
}

type constantAdapter struct {
	builder types.BuilderFactoryConstant
}

func (b constantAdapter) Dependencies() []reflect.Type {
	return b.builder.Dependencies()
}

func (b constantAdapter) Allocators() []reflect.Type {
	return b.builder.Allocators()
}

func (b constantAdapter) ConstantSize() uint64 {
	return b.builder.ConstantSize()
}

func (b constantAdapter) SizeCodeTemplate(_ *uint64) (string, bool) {
	return "", false
}

func (b constantAdapter) MarshalCodeTemplate(varIndex *uint64) string {
	return b.builder.MarshalCodeTemplate(varIndex)
}

func (b constantAdapter) UnmarshalCodeTemplate(varIndex *uint64) string {
	return b.builder.UnmarshalCodeTemplate(varIndex)
}

func newNonConstantAdapter(builder types.BuilderFactoryNonConstant) nonConstantAdapter {
	return nonConstantAdapter{
		builder: builder,
	}
}

type nonConstantAdapter struct {
	builder types.BuilderFactoryNonConstant
}

func (b nonConstantAdapter) Dependencies() []reflect.Type {
	return b.builder.Dependencies()
}

func (b nonConstantAdapter) Allocators() []reflect.Type {
	return b.builder.Allocators()
}

func (b nonConstantAdapter) ConstantSize() uint64 {
	return 0
}

func (b nonConstantAdapter) SizeCodeTemplate(varIndex *uint64) (string, bool) {
	return b.builder.SizeCodeTemplate(varIndex), true
}

func (b nonConstantAdapter) MarshalCodeTemplate(varIndex *uint64) string {
	return b.builder.MarshalCodeTemplate(varIndex)
}

func (b nonConstantAdapter) UnmarshalCodeTemplate(varIndex *uint64) string {
	return b.builder.UnmarshalCodeTemplate(varIndex)
}
