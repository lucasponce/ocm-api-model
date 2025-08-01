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

package v1 // github.com/openshift-online/ocm-api-model/clientapi/servicemgmt/v1

import (
	"io"
	"sort"

	jsoniter "github.com/json-iterator/go"
	"github.com/openshift-online/ocm-api-model/clientapi/helpers"
)

// MarshalAWS writes a value of the 'AWS' type to the given writer.
func MarshalAWS(object *AWS, writer io.Writer) error {
	stream := helpers.NewStream(writer)
	WriteAWS(object, stream)
	err := stream.Flush()
	if err != nil {
		return err
	}
	return stream.Error
}

// WriteAWS writes a value of the 'AWS' type to the given stream.
func WriteAWS(object *AWS, stream *jsoniter.Stream) {
	count := 0
	stream.WriteObjectStart()
	var present_ bool
	present_ = len(object.fieldSet_) > 0 && object.fieldSet_[0] && object.sts != nil
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("sts")
		WriteSTS(object.sts, stream)
		count++
	}
	present_ = len(object.fieldSet_) > 1 && object.fieldSet_[1]
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("access_key_id")
		stream.WriteString(object.accessKeyID)
		count++
	}
	present_ = len(object.fieldSet_) > 2 && object.fieldSet_[2]
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("account_id")
		stream.WriteString(object.accountID)
		count++
	}
	present_ = len(object.fieldSet_) > 3 && object.fieldSet_[3]
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("private_link")
		stream.WriteBool(object.privateLink)
		count++
	}
	present_ = len(object.fieldSet_) > 4 && object.fieldSet_[4]
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("secret_access_key")
		stream.WriteString(object.secretAccessKey)
		count++
	}
	present_ = len(object.fieldSet_) > 5 && object.fieldSet_[5] && object.subnetIDs != nil
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("subnet_ids")
		WriteStringList(object.subnetIDs, stream)
		count++
	}
	present_ = len(object.fieldSet_) > 6 && object.fieldSet_[6] && object.tags != nil
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("tags")
		if object.tags != nil {
			stream.WriteObjectStart()
			keys := make([]string, len(object.tags))
			i := 0
			for key := range object.tags {
				keys[i] = key
				i++
			}
			sort.Strings(keys)
			for i, key := range keys {
				if i > 0 {
					stream.WriteMore()
				}
				item := object.tags[key]
				stream.WriteObjectField(key)
				stream.WriteString(item)
			}
			stream.WriteObjectEnd()
		} else {
			stream.WriteNil()
		}
	}
	stream.WriteObjectEnd()
}

// UnmarshalAWS reads a value of the 'AWS' type from the given
// source, which can be an slice of bytes, a string or a reader.
func UnmarshalAWS(source interface{}) (object *AWS, err error) {
	iterator, err := helpers.NewIterator(source)
	if err != nil {
		return
	}
	object = ReadAWS(iterator)
	err = iterator.Error
	return
}

// ReadAWS reads a value of the 'AWS' type from the given iterator.
func ReadAWS(iterator *jsoniter.Iterator) *AWS {
	object := &AWS{
		fieldSet_: make([]bool, 7),
	}
	for {
		field := iterator.ReadObject()
		if field == "" {
			break
		}
		switch field {
		case "sts":
			value := ReadSTS(iterator)
			object.sts = value
			object.fieldSet_[0] = true
		case "access_key_id":
			value := iterator.ReadString()
			object.accessKeyID = value
			object.fieldSet_[1] = true
		case "account_id":
			value := iterator.ReadString()
			object.accountID = value
			object.fieldSet_[2] = true
		case "private_link":
			value := iterator.ReadBool()
			object.privateLink = value
			object.fieldSet_[3] = true
		case "secret_access_key":
			value := iterator.ReadString()
			object.secretAccessKey = value
			object.fieldSet_[4] = true
		case "subnet_ids":
			value := ReadStringList(iterator)
			object.subnetIDs = value
			object.fieldSet_[5] = true
		case "tags":
			value := map[string]string{}
			for {
				key := iterator.ReadObject()
				if key == "" {
					break
				}
				item := iterator.ReadString()
				value[key] = item
			}
			object.tags = value
			object.fieldSet_[6] = true
		default:
			iterator.ReadAny()
		}
	}
	return object
}
