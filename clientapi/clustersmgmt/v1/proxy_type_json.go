/*
Copyright (c) 2020 Red Hat, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// IMPORTANT: This file has been generated automatically, refrain from modifying it manually as all
// your changes will be lost when the file is generated again.

package v1 // github.com/openshift-online/ocm-api-model/clientapi/clustersmgmt/v1

import (
	"io"

	jsoniter "github.com/json-iterator/go"
	"github.com/openshift-online/ocm-api-model/clientapi/helpers"
)

// MarshalProxy writes a value of the 'proxy' type to the given writer.
func MarshalProxy(object *Proxy, writer io.Writer) error {
	stream := helpers.NewStream(writer)
	WriteProxy(object, stream)
	err := stream.Flush()
	if err != nil {
		return err
	}
	return stream.Error
}

// WriteProxy writes a value of the 'proxy' type to the given stream.
func WriteProxy(object *Proxy, stream *jsoniter.Stream) {
	count := 0
	stream.WriteObjectStart()
	var present_ bool
	present_ = len(object.fieldSet_) > 0 && object.fieldSet_[0]
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("http_proxy")
		stream.WriteString(object.httpProxy)
		count++
	}
	present_ = len(object.fieldSet_) > 1 && object.fieldSet_[1]
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("https_proxy")
		stream.WriteString(object.httpsProxy)
		count++
	}
	present_ = len(object.fieldSet_) > 2 && object.fieldSet_[2]
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("no_proxy")
		stream.WriteString(object.noProxy)
	}
	stream.WriteObjectEnd()
}

// UnmarshalProxy reads a value of the 'proxy' type from the given
// source, which can be an slice of bytes, a string or a reader.
func UnmarshalProxy(source interface{}) (object *Proxy, err error) {
	iterator, err := helpers.NewIterator(source)
	if err != nil {
		return
	}
	object = ReadProxy(iterator)
	err = iterator.Error
	return
}

// ReadProxy reads a value of the 'proxy' type from the given iterator.
func ReadProxy(iterator *jsoniter.Iterator) *Proxy {
	object := &Proxy{
		fieldSet_: make([]bool, 3),
	}
	for {
		field := iterator.ReadObject()
		if field == "" {
			break
		}
		switch field {
		case "http_proxy":
			value := iterator.ReadString()
			object.httpProxy = value
			object.fieldSet_[0] = true
		case "https_proxy":
			value := iterator.ReadString()
			object.httpsProxy = value
			object.fieldSet_[1] = true
		case "no_proxy":
			value := iterator.ReadString()
			object.noProxy = value
			object.fieldSet_[2] = true
		default:
			iterator.ReadAny()
		}
	}
	return object
}
