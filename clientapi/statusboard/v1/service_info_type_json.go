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

package v1 // github.com/openshift-online/ocm-api-model/clientapi/statusboard/v1

import (
	"io"

	jsoniter "github.com/json-iterator/go"
	"github.com/openshift-online/ocm-api-model/clientapi/helpers"
)

// MarshalServiceInfo writes a value of the 'service_info' type to the given writer.
func MarshalServiceInfo(object *ServiceInfo, writer io.Writer) error {
	stream := helpers.NewStream(writer)
	WriteServiceInfo(object, stream)
	err := stream.Flush()
	if err != nil {
		return err
	}
	return stream.Error
}

// WriteServiceInfo writes a value of the 'service_info' type to the given stream.
func WriteServiceInfo(object *ServiceInfo, stream *jsoniter.Stream) {
	count := 0
	stream.WriteObjectStart()
	var present_ bool
	present_ = len(object.fieldSet_) > 0 && object.fieldSet_[0]
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("fullname")
		stream.WriteString(object.fullname)
		count++
	}
	present_ = len(object.fieldSet_) > 1 && object.fieldSet_[1]
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("status_type")
		stream.WriteString(object.statusType)
	}
	stream.WriteObjectEnd()
}

// UnmarshalServiceInfo reads a value of the 'service_info' type from the given
// source, which can be an slice of bytes, a string or a reader.
func UnmarshalServiceInfo(source interface{}) (object *ServiceInfo, err error) {
	iterator, err := helpers.NewIterator(source)
	if err != nil {
		return
	}
	object = ReadServiceInfo(iterator)
	err = iterator.Error
	return
}

// ReadServiceInfo reads a value of the 'service_info' type from the given iterator.
func ReadServiceInfo(iterator *jsoniter.Iterator) *ServiceInfo {
	object := &ServiceInfo{
		fieldSet_: make([]bool, 2),
	}
	for {
		field := iterator.ReadObject()
		if field == "" {
			break
		}
		switch field {
		case "fullname":
			value := iterator.ReadString()
			object.fullname = value
			object.fieldSet_[0] = true
		case "status_type":
			value := iterator.ReadString()
			object.statusType = value
			object.fieldSet_[1] = true
		default:
			iterator.ReadAny()
		}
	}
	return object
}
