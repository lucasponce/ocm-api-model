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

// MarshalNodesInfo writes a value of the 'nodes_info' type to the given writer.
func MarshalNodesInfo(object *NodesInfo, writer io.Writer) error {
	stream := helpers.NewStream(writer)
	WriteNodesInfo(object, stream)
	err := stream.Flush()
	if err != nil {
		return err
	}
	return stream.Error
}

// WriteNodesInfo writes a value of the 'nodes_info' type to the given stream.
func WriteNodesInfo(object *NodesInfo, stream *jsoniter.Stream) {
	count := 0
	stream.WriteObjectStart()
	var present_ bool
	present_ = len(object.fieldSet_) > 0 && object.fieldSet_[0] && object.nodes != nil
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("nodes")
		WriteNodeInfoList(object.nodes, stream)
	}
	stream.WriteObjectEnd()
}

// UnmarshalNodesInfo reads a value of the 'nodes_info' type from the given
// source, which can be an slice of bytes, a string or a reader.
func UnmarshalNodesInfo(source interface{}) (object *NodesInfo, err error) {
	iterator, err := helpers.NewIterator(source)
	if err != nil {
		return
	}
	object = ReadNodesInfo(iterator)
	err = iterator.Error
	return
}

// ReadNodesInfo reads a value of the 'nodes_info' type from the given iterator.
func ReadNodesInfo(iterator *jsoniter.Iterator) *NodesInfo {
	object := &NodesInfo{
		fieldSet_: make([]bool, 1),
	}
	for {
		field := iterator.ReadObject()
		if field == "" {
			break
		}
		switch field {
		case "nodes":
			value := ReadNodeInfoList(iterator)
			object.nodes = value
			object.fieldSet_[0] = true
		default:
			iterator.ReadAny()
		}
	}
	return object
}
