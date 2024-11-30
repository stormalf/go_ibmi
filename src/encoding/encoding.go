// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package encoding defines interfaces shared by other packages that
// convert data to and from byte-level and textual representations.
// Packages that check for these interfaces include encoding/gob,
// encoding/json, and encoding/xml. As a result, implementing an
// interface once can make a type useful in multiple encodings.
// Standard types that implement these interfaces include time.Time and net.IP.
// The interfaces come in pairs that produce and consume encoded data.
//
// Adding encoding/decoding methods to existing types may constitute a breaking change,
// as they can be used for serialization in communicating with programs
// written with different library versions.
// The policy for packages maintained by the Go project is to only allow
// the addition of marshaling functions if no existing, reasonable marshaling exists.
package encoding

// BinaryMarshaler is the interface implemented by an object that can
// marshal itself into a binary form.
//
// MarshalBinary encodes the receiver into a binary form and returns the result.
type BinaryMarshaler interface {
	MarshalBinary() (data []byte, err error)
}

// BinaryUnmarshaler is the interface implemented by an object that can
// unmarshal a binary representation of itself.
//
// UnmarshalBinary must be able to decode the form generated by MarshalBinary.
// UnmarshalBinary must copy the data if it wishes to retain the data
// after returning.
type BinaryUnmarshaler interface {
	UnmarshalBinary(data []byte) error
}

// BinaryAppender is the interface implemented by an object
// that can append the binary representation of itself.
// If a type implements both [BinaryAppender] and [BinaryMarshaler],
// then v.MarshalBinary() must be semantically identical to v.AppendBinary(nil).
type BinaryAppender interface {
	// AppendBinary appends the binary representation of itself to the end of b
	// (allocating a larger slice if necessary) and returns the updated slice.
	//
	// Implementations must not retain b, nor mutate any bytes within b[:len(b)].
	AppendBinary(b []byte) ([]byte, error)
}

// TextMarshaler is the interface implemented by an object that can
// marshal itself into a textual form.
//
// MarshalText encodes the receiver into UTF-8-encoded text and returns the result.
type TextMarshaler interface {
	MarshalText() (text []byte, err error)
}

// TextUnmarshaler is the interface implemented by an object that can
// unmarshal a textual representation of itself.
//
// UnmarshalText must be able to decode the form generated by MarshalText.
// UnmarshalText must copy the text if it wishes to retain the text
// after returning.
type TextUnmarshaler interface {
	UnmarshalText(text []byte) error
}

// TextAppender is the interface implemented by an object
// that can append the textual representation of itself.
// If a type implements both [TextAppender] and [TextMarshaler],
// then v.MarshalText() must be semantically identical to v.AppendText(nil).
type TextAppender interface {
	// AppendText appends the textual representation of itself to the end of b
	// (allocating a larger slice if necessary) and returns the updated slice.
	//
	// Implementations must not retain b, nor mutate any bytes within b[:len(b)].
	AppendText(b []byte) ([]byte, error)
}