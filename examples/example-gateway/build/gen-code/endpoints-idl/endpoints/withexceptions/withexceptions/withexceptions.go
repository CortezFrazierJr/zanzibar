// Code generated by thriftrw v1.29.2. DO NOT EDIT.
// @generated

package withexceptions

import (
	errors "errors"
	fmt "fmt"
	strings "strings"

	multierr "go.uber.org/multierr"
	stream "go.uber.org/thriftrw/protocol/stream"
	wire "go.uber.org/thriftrw/wire"
	zapcore "go.uber.org/zap/zapcore"
)

type EndpointExceptionType1 struct {
	Message1 string `json:"message1,required"`
}

// ToWire translates a EndpointExceptionType1 struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//   x, err := v.ToWire()
//   if err != nil {
//     return err
//   }
//
//   if err := binaryProtocol.Encode(x, writer); err != nil {
//     return err
//   }
func (v *EndpointExceptionType1) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	w, err = wire.NewValueString(v.Message1), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 1, Value: w}
	i++

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

// FromWire deserializes a EndpointExceptionType1 struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a EndpointExceptionType1 struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v EndpointExceptionType1
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *EndpointExceptionType1) FromWire(w wire.Value) error {
	var err error

	message1IsSet := false

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TBinary {
				v.Message1, err = field.Value.GetString(), error(nil)
				if err != nil {
					return err
				}
				message1IsSet = true
			}
		}
	}

	if !message1IsSet {
		return errors.New("field Message1 of EndpointExceptionType1 is required")
	}

	return nil
}

// Encode serializes a EndpointExceptionType1 struct directly into bytes, without going
// through an intermediary type.
//
// An error is returned if a EndpointExceptionType1 struct could not be encoded.
func (v *EndpointExceptionType1) Encode(sw stream.Writer) error {
	if err := sw.WriteStructBegin(); err != nil {
		return err
	}

	if err := sw.WriteFieldBegin(stream.FieldHeader{ID: 1, Type: wire.TBinary}); err != nil {
		return err
	}
	if err := sw.WriteString(v.Message1); err != nil {
		return err
	}
	if err := sw.WriteFieldEnd(); err != nil {
		return err
	}

	return sw.WriteStructEnd()
}

// Decode deserializes a EndpointExceptionType1 struct directly from its Thrift-level
// representation, without going through an intemediary type.
//
// An error is returned if a EndpointExceptionType1 struct could not be generated from the wire
// representation.
func (v *EndpointExceptionType1) Decode(sr stream.Reader) error {

	message1IsSet := false

	if err := sr.ReadStructBegin(); err != nil {
		return err
	}

	fh, ok, err := sr.ReadFieldBegin()
	if err != nil {
		return err
	}

	for ok {
		switch {
		case fh.ID == 1 && fh.Type == wire.TBinary:
			v.Message1, err = sr.ReadString()
			if err != nil {
				return err
			}
			message1IsSet = true
		default:
			if err := sr.Skip(fh.Type); err != nil {
				return err
			}
		}

		if err := sr.ReadFieldEnd(); err != nil {
			return err
		}

		if fh, ok, err = sr.ReadFieldBegin(); err != nil {
			return err
		}
	}

	if err := sr.ReadStructEnd(); err != nil {
		return err
	}

	if !message1IsSet {
		return errors.New("field Message1 of EndpointExceptionType1 is required")
	}

	return nil
}

// String returns a readable string representation of a EndpointExceptionType1
// struct.
func (v *EndpointExceptionType1) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [1]string
	i := 0
	fields[i] = fmt.Sprintf("Message1: %v", v.Message1)
	i++

	return fmt.Sprintf("EndpointExceptionType1{%v}", strings.Join(fields[:i], ", "))
}

// ErrorName is the name of this type as defined in the Thrift
// file.
func (*EndpointExceptionType1) ErrorName() string {
	return "EndpointExceptionType1"
}

// Equals returns true if all the fields of this EndpointExceptionType1 match the
// provided EndpointExceptionType1.
//
// This function performs a deep comparison.
func (v *EndpointExceptionType1) Equals(rhs *EndpointExceptionType1) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}
	if !(v.Message1 == rhs.Message1) {
		return false
	}

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of EndpointExceptionType1.
func (v *EndpointExceptionType1) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v == nil {
		return nil
	}
	enc.AddString("message1", v.Message1)
	return err
}

// GetMessage1 returns the value of Message1 if it is set or its
// zero value if it is unset.
func (v *EndpointExceptionType1) GetMessage1() (o string) {
	if v != nil {
		o = v.Message1
	}
	return
}

func (v *EndpointExceptionType1) Error() string {
	return v.String()
}

type EndpointExceptionType2 struct {
	Message2 string `json:"message2,required"`
}

// ToWire translates a EndpointExceptionType2 struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//   x, err := v.ToWire()
//   if err != nil {
//     return err
//   }
//
//   if err := binaryProtocol.Encode(x, writer); err != nil {
//     return err
//   }
func (v *EndpointExceptionType2) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	w, err = wire.NewValueString(v.Message2), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 1, Value: w}
	i++

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

// FromWire deserializes a EndpointExceptionType2 struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a EndpointExceptionType2 struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v EndpointExceptionType2
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *EndpointExceptionType2) FromWire(w wire.Value) error {
	var err error

	message2IsSet := false

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TBinary {
				v.Message2, err = field.Value.GetString(), error(nil)
				if err != nil {
					return err
				}
				message2IsSet = true
			}
		}
	}

	if !message2IsSet {
		return errors.New("field Message2 of EndpointExceptionType2 is required")
	}

	return nil
}

// Encode serializes a EndpointExceptionType2 struct directly into bytes, without going
// through an intermediary type.
//
// An error is returned if a EndpointExceptionType2 struct could not be encoded.
func (v *EndpointExceptionType2) Encode(sw stream.Writer) error {
	if err := sw.WriteStructBegin(); err != nil {
		return err
	}

	if err := sw.WriteFieldBegin(stream.FieldHeader{ID: 1, Type: wire.TBinary}); err != nil {
		return err
	}
	if err := sw.WriteString(v.Message2); err != nil {
		return err
	}
	if err := sw.WriteFieldEnd(); err != nil {
		return err
	}

	return sw.WriteStructEnd()
}

// Decode deserializes a EndpointExceptionType2 struct directly from its Thrift-level
// representation, without going through an intemediary type.
//
// An error is returned if a EndpointExceptionType2 struct could not be generated from the wire
// representation.
func (v *EndpointExceptionType2) Decode(sr stream.Reader) error {

	message2IsSet := false

	if err := sr.ReadStructBegin(); err != nil {
		return err
	}

	fh, ok, err := sr.ReadFieldBegin()
	if err != nil {
		return err
	}

	for ok {
		switch {
		case fh.ID == 1 && fh.Type == wire.TBinary:
			v.Message2, err = sr.ReadString()
			if err != nil {
				return err
			}
			message2IsSet = true
		default:
			if err := sr.Skip(fh.Type); err != nil {
				return err
			}
		}

		if err := sr.ReadFieldEnd(); err != nil {
			return err
		}

		if fh, ok, err = sr.ReadFieldBegin(); err != nil {
			return err
		}
	}

	if err := sr.ReadStructEnd(); err != nil {
		return err
	}

	if !message2IsSet {
		return errors.New("field Message2 of EndpointExceptionType2 is required")
	}

	return nil
}

// String returns a readable string representation of a EndpointExceptionType2
// struct.
func (v *EndpointExceptionType2) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [1]string
	i := 0
	fields[i] = fmt.Sprintf("Message2: %v", v.Message2)
	i++

	return fmt.Sprintf("EndpointExceptionType2{%v}", strings.Join(fields[:i], ", "))
}

// ErrorName is the name of this type as defined in the Thrift
// file.
func (*EndpointExceptionType2) ErrorName() string {
	return "EndpointExceptionType2"
}

// Equals returns true if all the fields of this EndpointExceptionType2 match the
// provided EndpointExceptionType2.
//
// This function performs a deep comparison.
func (v *EndpointExceptionType2) Equals(rhs *EndpointExceptionType2) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}
	if !(v.Message2 == rhs.Message2) {
		return false
	}

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of EndpointExceptionType2.
func (v *EndpointExceptionType2) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v == nil {
		return nil
	}
	enc.AddString("message2", v.Message2)
	return err
}

// GetMessage2 returns the value of Message2 if it is set or its
// zero value if it is unset.
func (v *EndpointExceptionType2) GetMessage2() (o string) {
	if v != nil {
		o = v.Message2
	}
	return
}

func (v *EndpointExceptionType2) Error() string {
	return v.String()
}

type Response struct {
}

// ToWire translates a Response struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//   x, err := v.ToWire()
//   if err != nil {
//     return err
//   }
//
//   if err := binaryProtocol.Encode(x, writer); err != nil {
//     return err
//   }
func (v *Response) ToWire() (wire.Value, error) {
	var (
		fields [0]wire.Field
		i      int = 0
	)

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

// FromWire deserializes a Response struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a Response struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v Response
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *Response) FromWire(w wire.Value) error {

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		}
	}

	return nil
}

// Encode serializes a Response struct directly into bytes, without going
// through an intermediary type.
//
// An error is returned if a Response struct could not be encoded.
func (v *Response) Encode(sw stream.Writer) error {
	if err := sw.WriteStructBegin(); err != nil {
		return err
	}

	return sw.WriteStructEnd()
}

// Decode deserializes a Response struct directly from its Thrift-level
// representation, without going through an intemediary type.
//
// An error is returned if a Response struct could not be generated from the wire
// representation.
func (v *Response) Decode(sr stream.Reader) error {

	if err := sr.ReadStructBegin(); err != nil {
		return err
	}

	fh, ok, err := sr.ReadFieldBegin()
	if err != nil {
		return err
	}

	for ok {
		switch {
		default:
			if err := sr.Skip(fh.Type); err != nil {
				return err
			}
		}

		if err := sr.ReadFieldEnd(); err != nil {
			return err
		}

		if fh, ok, err = sr.ReadFieldBegin(); err != nil {
			return err
		}
	}

	if err := sr.ReadStructEnd(); err != nil {
		return err
	}

	return nil
}

// String returns a readable string representation of a Response
// struct.
func (v *Response) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [0]string
	i := 0

	return fmt.Sprintf("Response{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this Response match the
// provided Response.
//
// This function performs a deep comparison.
func (v *Response) Equals(rhs *Response) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of Response.
func (v *Response) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v == nil {
		return nil
	}
	return err
}

// WithExceptions_Func1_Args represents the arguments for the WithExceptions.Func1 function.
//
// The arguments for Func1 are sent and received over the wire as this struct.
type WithExceptions_Func1_Args struct {
}

// ToWire translates a WithExceptions_Func1_Args struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//   x, err := v.ToWire()
//   if err != nil {
//     return err
//   }
//
//   if err := binaryProtocol.Encode(x, writer); err != nil {
//     return err
//   }
func (v *WithExceptions_Func1_Args) ToWire() (wire.Value, error) {
	var (
		fields [0]wire.Field
		i      int = 0
	)

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

// FromWire deserializes a WithExceptions_Func1_Args struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a WithExceptions_Func1_Args struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v WithExceptions_Func1_Args
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *WithExceptions_Func1_Args) FromWire(w wire.Value) error {

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		}
	}

	return nil
}

// Encode serializes a WithExceptions_Func1_Args struct directly into bytes, without going
// through an intermediary type.
//
// An error is returned if a WithExceptions_Func1_Args struct could not be encoded.
func (v *WithExceptions_Func1_Args) Encode(sw stream.Writer) error {
	if err := sw.WriteStructBegin(); err != nil {
		return err
	}

	return sw.WriteStructEnd()
}

// Decode deserializes a WithExceptions_Func1_Args struct directly from its Thrift-level
// representation, without going through an intemediary type.
//
// An error is returned if a WithExceptions_Func1_Args struct could not be generated from the wire
// representation.
func (v *WithExceptions_Func1_Args) Decode(sr stream.Reader) error {

	if err := sr.ReadStructBegin(); err != nil {
		return err
	}

	fh, ok, err := sr.ReadFieldBegin()
	if err != nil {
		return err
	}

	for ok {
		switch {
		default:
			if err := sr.Skip(fh.Type); err != nil {
				return err
			}
		}

		if err := sr.ReadFieldEnd(); err != nil {
			return err
		}

		if fh, ok, err = sr.ReadFieldBegin(); err != nil {
			return err
		}
	}

	if err := sr.ReadStructEnd(); err != nil {
		return err
	}

	return nil
}

// String returns a readable string representation of a WithExceptions_Func1_Args
// struct.
func (v *WithExceptions_Func1_Args) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [0]string
	i := 0

	return fmt.Sprintf("WithExceptions_Func1_Args{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this WithExceptions_Func1_Args match the
// provided WithExceptions_Func1_Args.
//
// This function performs a deep comparison.
func (v *WithExceptions_Func1_Args) Equals(rhs *WithExceptions_Func1_Args) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of WithExceptions_Func1_Args.
func (v *WithExceptions_Func1_Args) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v == nil {
		return nil
	}
	return err
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the arguments.
//
// This will always be "Func1" for this struct.
func (v *WithExceptions_Func1_Args) MethodName() string {
	return "Func1"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Call for this struct.
func (v *WithExceptions_Func1_Args) EnvelopeType() wire.EnvelopeType {
	return wire.Call
}

// WithExceptions_Func1_Helper provides functions that aid in handling the
// parameters and return values of the WithExceptions.Func1
// function.
var WithExceptions_Func1_Helper = struct {
	// Args accepts the parameters of Func1 in-order and returns
	// the arguments struct for the function.
	Args func() *WithExceptions_Func1_Args

	// IsException returns true if the given error can be thrown
	// by Func1.
	//
	// An error can be thrown by Func1 only if the
	// corresponding exception type was mentioned in the 'throws'
	// section for it in the Thrift file.
	IsException func(error) bool

	// WrapResponse returns the result struct for Func1
	// given its return value and error.
	//
	// This allows mapping values and errors returned by
	// Func1 into a serializable result struct.
	// WrapResponse returns a non-nil error if the provided
	// error cannot be thrown by Func1
	//
	//   value, err := Func1(args)
	//   result, err := WithExceptions_Func1_Helper.WrapResponse(value, err)
	//   if err != nil {
	//     return fmt.Errorf("unexpected error from Func1: %v", err)
	//   }
	//   serialize(result)
	WrapResponse func(*Response, error) (*WithExceptions_Func1_Result, error)

	// UnwrapResponse takes the result struct for Func1
	// and returns the value or error returned by it.
	//
	// The error is non-nil only if Func1 threw an
	// exception.
	//
	//   result := deserialize(bytes)
	//   value, err := WithExceptions_Func1_Helper.UnwrapResponse(result)
	UnwrapResponse func(*WithExceptions_Func1_Result) (*Response, error)
}{}

func init() {
	WithExceptions_Func1_Helper.Args = func() *WithExceptions_Func1_Args {
		return &WithExceptions_Func1_Args{}
	}

	WithExceptions_Func1_Helper.IsException = func(err error) bool {
		switch err.(type) {
		case *EndpointExceptionType1:
			return true
		case *EndpointExceptionType2:
			return true
		default:
			return false
		}
	}

	WithExceptions_Func1_Helper.WrapResponse = func(success *Response, err error) (*WithExceptions_Func1_Result, error) {
		if err == nil {
			return &WithExceptions_Func1_Result{Success: success}, nil
		}

		switch e := err.(type) {
		case *EndpointExceptionType1:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for WithExceptions_Func1_Result.E1")
			}
			return &WithExceptions_Func1_Result{E1: e}, nil
		case *EndpointExceptionType2:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for WithExceptions_Func1_Result.E2")
			}
			return &WithExceptions_Func1_Result{E2: e}, nil
		}

		return nil, err
	}
	WithExceptions_Func1_Helper.UnwrapResponse = func(result *WithExceptions_Func1_Result) (success *Response, err error) {
		if result.E1 != nil {
			err = result.E1
			return
		}
		if result.E2 != nil {
			err = result.E2
			return
		}

		if result.Success != nil {
			success = result.Success
			return
		}

		err = errors.New("expected a non-void result")
		return
	}

}

// WithExceptions_Func1_Result represents the result of a WithExceptions.Func1 function call.
//
// The result of a Func1 execution is sent and received over the wire as this struct.
//
// Success is set only if the function did not throw an exception.
type WithExceptions_Func1_Result struct {
	// Value returned by Func1 after a successful execution.
	Success *Response               `json:"success,omitempty"`
	E1      *EndpointExceptionType1 `json:"e1,omitempty"`
	E2      *EndpointExceptionType2 `json:"e2,omitempty"`
}

// ToWire translates a WithExceptions_Func1_Result struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//   x, err := v.ToWire()
//   if err != nil {
//     return err
//   }
//
//   if err := binaryProtocol.Encode(x, writer); err != nil {
//     return err
//   }
func (v *WithExceptions_Func1_Result) ToWire() (wire.Value, error) {
	var (
		fields [3]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	if v.Success != nil {
		w, err = v.Success.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 0, Value: w}
		i++
	}
	if v.E1 != nil {
		w, err = v.E1.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}
	if v.E2 != nil {
		w, err = v.E2.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 2, Value: w}
		i++
	}

	if i != 1 {
		return wire.Value{}, fmt.Errorf("WithExceptions_Func1_Result should have exactly one field: got %v fields", i)
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _Response_Read(w wire.Value) (*Response, error) {
	var v Response
	err := v.FromWire(w)
	return &v, err
}

func _EndpointExceptionType1_Read(w wire.Value) (*EndpointExceptionType1, error) {
	var v EndpointExceptionType1
	err := v.FromWire(w)
	return &v, err
}

func _EndpointExceptionType2_Read(w wire.Value) (*EndpointExceptionType2, error) {
	var v EndpointExceptionType2
	err := v.FromWire(w)
	return &v, err
}

// FromWire deserializes a WithExceptions_Func1_Result struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a WithExceptions_Func1_Result struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v WithExceptions_Func1_Result
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *WithExceptions_Func1_Result) FromWire(w wire.Value) error {
	var err error

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 0:
			if field.Value.Type() == wire.TStruct {
				v.Success, err = _Response_Read(field.Value)
				if err != nil {
					return err
				}

			}
		case 1:
			if field.Value.Type() == wire.TStruct {
				v.E1, err = _EndpointExceptionType1_Read(field.Value)
				if err != nil {
					return err
				}

			}
		case 2:
			if field.Value.Type() == wire.TStruct {
				v.E2, err = _EndpointExceptionType2_Read(field.Value)
				if err != nil {
					return err
				}

			}
		}
	}

	count := 0
	if v.Success != nil {
		count++
	}
	if v.E1 != nil {
		count++
	}
	if v.E2 != nil {
		count++
	}
	if count != 1 {
		return fmt.Errorf("WithExceptions_Func1_Result should have exactly one field: got %v fields", count)
	}

	return nil
}

// Encode serializes a WithExceptions_Func1_Result struct directly into bytes, without going
// through an intermediary type.
//
// An error is returned if a WithExceptions_Func1_Result struct could not be encoded.
func (v *WithExceptions_Func1_Result) Encode(sw stream.Writer) error {
	if err := sw.WriteStructBegin(); err != nil {
		return err
	}

	if v.Success != nil {
		if err := sw.WriteFieldBegin(stream.FieldHeader{ID: 0, Type: wire.TStruct}); err != nil {
			return err
		}
		if err := v.Success.Encode(sw); err != nil {
			return err
		}
		if err := sw.WriteFieldEnd(); err != nil {
			return err
		}
	}

	if v.E1 != nil {
		if err := sw.WriteFieldBegin(stream.FieldHeader{ID: 1, Type: wire.TStruct}); err != nil {
			return err
		}
		if err := v.E1.Encode(sw); err != nil {
			return err
		}
		if err := sw.WriteFieldEnd(); err != nil {
			return err
		}
	}

	if v.E2 != nil {
		if err := sw.WriteFieldBegin(stream.FieldHeader{ID: 2, Type: wire.TStruct}); err != nil {
			return err
		}
		if err := v.E2.Encode(sw); err != nil {
			return err
		}
		if err := sw.WriteFieldEnd(); err != nil {
			return err
		}
	}

	count := 0
	if v.Success != nil {
		count++
	}
	if v.E1 != nil {
		count++
	}
	if v.E2 != nil {
		count++
	}

	if count != 1 {
		return fmt.Errorf("WithExceptions_Func1_Result should have exactly one field: got %v fields", count)
	}

	return sw.WriteStructEnd()
}

func _Response_Decode(sr stream.Reader) (*Response, error) {
	var v Response
	err := v.Decode(sr)
	return &v, err
}

func _EndpointExceptionType1_Decode(sr stream.Reader) (*EndpointExceptionType1, error) {
	var v EndpointExceptionType1
	err := v.Decode(sr)
	return &v, err
}

func _EndpointExceptionType2_Decode(sr stream.Reader) (*EndpointExceptionType2, error) {
	var v EndpointExceptionType2
	err := v.Decode(sr)
	return &v, err
}

// Decode deserializes a WithExceptions_Func1_Result struct directly from its Thrift-level
// representation, without going through an intemediary type.
//
// An error is returned if a WithExceptions_Func1_Result struct could not be generated from the wire
// representation.
func (v *WithExceptions_Func1_Result) Decode(sr stream.Reader) error {

	if err := sr.ReadStructBegin(); err != nil {
		return err
	}

	fh, ok, err := sr.ReadFieldBegin()
	if err != nil {
		return err
	}

	for ok {
		switch {
		case fh.ID == 0 && fh.Type == wire.TStruct:
			v.Success, err = _Response_Decode(sr)
			if err != nil {
				return err
			}

		case fh.ID == 1 && fh.Type == wire.TStruct:
			v.E1, err = _EndpointExceptionType1_Decode(sr)
			if err != nil {
				return err
			}

		case fh.ID == 2 && fh.Type == wire.TStruct:
			v.E2, err = _EndpointExceptionType2_Decode(sr)
			if err != nil {
				return err
			}

		default:
			if err := sr.Skip(fh.Type); err != nil {
				return err
			}
		}

		if err := sr.ReadFieldEnd(); err != nil {
			return err
		}

		if fh, ok, err = sr.ReadFieldBegin(); err != nil {
			return err
		}
	}

	if err := sr.ReadStructEnd(); err != nil {
		return err
	}

	count := 0
	if v.Success != nil {
		count++
	}
	if v.E1 != nil {
		count++
	}
	if v.E2 != nil {
		count++
	}
	if count != 1 {
		return fmt.Errorf("WithExceptions_Func1_Result should have exactly one field: got %v fields", count)
	}

	return nil
}

// String returns a readable string representation of a WithExceptions_Func1_Result
// struct.
func (v *WithExceptions_Func1_Result) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [3]string
	i := 0
	if v.Success != nil {
		fields[i] = fmt.Sprintf("Success: %v", v.Success)
		i++
	}
	if v.E1 != nil {
		fields[i] = fmt.Sprintf("E1: %v", v.E1)
		i++
	}
	if v.E2 != nil {
		fields[i] = fmt.Sprintf("E2: %v", v.E2)
		i++
	}

	return fmt.Sprintf("WithExceptions_Func1_Result{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this WithExceptions_Func1_Result match the
// provided WithExceptions_Func1_Result.
//
// This function performs a deep comparison.
func (v *WithExceptions_Func1_Result) Equals(rhs *WithExceptions_Func1_Result) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}
	if !((v.Success == nil && rhs.Success == nil) || (v.Success != nil && rhs.Success != nil && v.Success.Equals(rhs.Success))) {
		return false
	}
	if !((v.E1 == nil && rhs.E1 == nil) || (v.E1 != nil && rhs.E1 != nil && v.E1.Equals(rhs.E1))) {
		return false
	}
	if !((v.E2 == nil && rhs.E2 == nil) || (v.E2 != nil && rhs.E2 != nil && v.E2.Equals(rhs.E2))) {
		return false
	}

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of WithExceptions_Func1_Result.
func (v *WithExceptions_Func1_Result) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v == nil {
		return nil
	}
	if v.Success != nil {
		err = multierr.Append(err, enc.AddObject("success", v.Success))
	}
	if v.E1 != nil {
		err = multierr.Append(err, enc.AddObject("e1", v.E1))
	}
	if v.E2 != nil {
		err = multierr.Append(err, enc.AddObject("e2", v.E2))
	}
	return err
}

// GetSuccess returns the value of Success if it is set or its
// zero value if it is unset.
func (v *WithExceptions_Func1_Result) GetSuccess() (o *Response) {
	if v != nil && v.Success != nil {
		return v.Success
	}

	return
}

// IsSetSuccess returns true if Success is not nil.
func (v *WithExceptions_Func1_Result) IsSetSuccess() bool {
	return v != nil && v.Success != nil
}

// GetE1 returns the value of E1 if it is set or its
// zero value if it is unset.
func (v *WithExceptions_Func1_Result) GetE1() (o *EndpointExceptionType1) {
	if v != nil && v.E1 != nil {
		return v.E1
	}

	return
}

// IsSetE1 returns true if E1 is not nil.
func (v *WithExceptions_Func1_Result) IsSetE1() bool {
	return v != nil && v.E1 != nil
}

// GetE2 returns the value of E2 if it is set or its
// zero value if it is unset.
func (v *WithExceptions_Func1_Result) GetE2() (o *EndpointExceptionType2) {
	if v != nil && v.E2 != nil {
		return v.E2
	}

	return
}

// IsSetE2 returns true if E2 is not nil.
func (v *WithExceptions_Func1_Result) IsSetE2() bool {
	return v != nil && v.E2 != nil
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the result.
//
// This will always be "Func1" for this struct.
func (v *WithExceptions_Func1_Result) MethodName() string {
	return "Func1"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Reply for this struct.
func (v *WithExceptions_Func1_Result) EnvelopeType() wire.EnvelopeType {
	return wire.Reply
}
