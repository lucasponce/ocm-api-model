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

// MarshalClusterCredentials writes a value of the 'cluster_credentials' type to the given writer.
func MarshalClusterCredentials(object *ClusterCredentials, writer io.Writer) error {
	stream := helpers.NewStream(writer)
	WriteClusterCredentials(object, stream)
	err := stream.Flush()
	if err != nil {
		return err
	}
	return stream.Error
}

// WriteClusterCredentials writes a value of the 'cluster_credentials' type to the given stream.
func WriteClusterCredentials(object *ClusterCredentials, stream *jsoniter.Stream) {
	count := 0
	stream.WriteObjectStart()
	stream.WriteObjectField("kind")
	if len(object.fieldSet_) > 0 && object.fieldSet_[0] {
		stream.WriteString(ClusterCredentialsLinkKind)
	} else {
		stream.WriteString(ClusterCredentialsKind)
	}
	count++
	if len(object.fieldSet_) > 1 && object.fieldSet_[1] {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("id")
		stream.WriteString(object.id)
		count++
	}
	if len(object.fieldSet_) > 2 && object.fieldSet_[2] {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("href")
		stream.WriteString(object.href)
		count++
	}
	var present_ bool
	present_ = len(object.fieldSet_) > 3 && object.fieldSet_[3]
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("kubeconfig")
		stream.WriteString(object.kubeconfig)
	}
	stream.WriteObjectEnd()
}

// UnmarshalClusterCredentials reads a value of the 'cluster_credentials' type from the given
// source, which can be an slice of bytes, a string or a reader.
func UnmarshalClusterCredentials(source interface{}) (object *ClusterCredentials, err error) {
	iterator, err := helpers.NewIterator(source)
	if err != nil {
		return
	}
	object = ReadClusterCredentials(iterator)
	err = iterator.Error
	return
}

// ReadClusterCredentials reads a value of the 'cluster_credentials' type from the given iterator.
func ReadClusterCredentials(iterator *jsoniter.Iterator) *ClusterCredentials {
	object := &ClusterCredentials{
		fieldSet_: make([]bool, 4),
	}
	for {
		field := iterator.ReadObject()
		if field == "" {
			break
		}
		switch field {
		case "kind":
			value := iterator.ReadString()
			if value == ClusterCredentialsLinkKind {
				object.fieldSet_[0] = true
			}
		case "id":
			object.id = iterator.ReadString()
			object.fieldSet_[1] = true
		case "href":
			object.href = iterator.ReadString()
			object.fieldSet_[2] = true
		case "kubeconfig":
			value := iterator.ReadString()
			object.kubeconfig = value
			object.fieldSet_[3] = true
		default:
			iterator.ReadAny()
		}
	}
	return object
}
