package serializer

import (
	"io/ioutil"

	// "github.com/golang/protobuf/jsonpb"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	// "google.golang.org/protobuf/proto"
	// "google.golang.org/protobuf/proto"
)

func WriteProtobufFileToJson(message proto.Message, fileName string) error {
	data, err := ProtobufToJson(message)

	if err != nil {
		return err
	}
	err = ioutil.WriteFile(fileName, []byte(data), 0644)
	if err != nil {
		return err
	}
	return nil
}
func ProtobufToJson(message proto.Message) (string, error) {
	// marshaler := jsonpb.Marshaler{
	// 	EnumsAsInts:  false,
	// 	EmitDefaults: true,
	// 	Indent:       " ",
	// 	OrigName:     true,
	// }
	// marshaler := jsonpb.Marshaler{
	// 	EnumsAsInts:  false,
	// 	EmitDefaults: true,
	// 	Indent:       " ",
	// 	OrigName:     true,
	// }
	b, err := protojson.Marshal(message)
	return string(b), err
}
func WriteProtoBufToBinaryFile(message proto.Message, fileName string) error {
	data, err := proto.Marshal(message)
	if err != nil {
		return err
	}
	ioutil.WriteFile(fileName, data, 0644)
	if err != nil {
		return err
	}
	return nil
}
func ReadProtoBufFromBinaryFile(fileName string, message proto.Message) error {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}
	err = proto.Unmarshal(data, message)
	if err != nil {
		return err
	}
	return nil
}
