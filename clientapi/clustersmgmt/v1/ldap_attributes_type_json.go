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

// MarshalLDAPAttributes writes a value of the 'LDAP_attributes' type to the given writer.
func MarshalLDAPAttributes(object *LDAPAttributes, writer io.Writer) error {
	stream := helpers.NewStream(writer)
	WriteLDAPAttributes(object, stream)
	err := stream.Flush()
	if err != nil {
		return err
	}
	return stream.Error
}

// WriteLDAPAttributes writes a value of the 'LDAP_attributes' type to the given stream.
func WriteLDAPAttributes(object *LDAPAttributes, stream *jsoniter.Stream) {
	count := 0
	stream.WriteObjectStart()
	var present_ bool
	present_ = len(object.fieldSet_) > 0 && object.fieldSet_[0] && object.id != nil
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("id")
		WriteStringList(object.id, stream)
		count++
	}
	present_ = len(object.fieldSet_) > 1 && object.fieldSet_[1] && object.email != nil
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("email")
		WriteStringList(object.email, stream)
		count++
	}
	present_ = len(object.fieldSet_) > 2 && object.fieldSet_[2] && object.name != nil
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("name")
		WriteStringList(object.name, stream)
		count++
	}
	present_ = len(object.fieldSet_) > 3 && object.fieldSet_[3] && object.preferredUsername != nil
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("preferred_username")
		WriteStringList(object.preferredUsername, stream)
	}
	stream.WriteObjectEnd()
}

// UnmarshalLDAPAttributes reads a value of the 'LDAP_attributes' type from the given
// source, which can be an slice of bytes, a string or a reader.
func UnmarshalLDAPAttributes(source interface{}) (object *LDAPAttributes, err error) {
	iterator, err := helpers.NewIterator(source)
	if err != nil {
		return
	}
	object = ReadLDAPAttributes(iterator)
	err = iterator.Error
	return
}

// ReadLDAPAttributes reads a value of the 'LDAP_attributes' type from the given iterator.
func ReadLDAPAttributes(iterator *jsoniter.Iterator) *LDAPAttributes {
	object := &LDAPAttributes{
		fieldSet_: make([]bool, 4),
	}
	for {
		field := iterator.ReadObject()
		if field == "" {
			break
		}
		switch field {
		case "id":
			value := ReadStringList(iterator)
			object.id = value
			object.fieldSet_[0] = true
		case "email":
			value := ReadStringList(iterator)
			object.email = value
			object.fieldSet_[1] = true
		case "name":
			value := ReadStringList(iterator)
			object.name = value
			object.fieldSet_[2] = true
		case "preferred_username":
			value := ReadStringList(iterator)
			object.preferredUsername = value
			object.fieldSet_[3] = true
		default:
			iterator.ReadAny()
		}
	}
	return object
}
