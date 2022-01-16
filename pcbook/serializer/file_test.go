package serializer

import (
	"testing"

	"github.com/rohit-tambe/go-grpc/pcbook/pb"
	"github.com/rohit-tambe/go-grpc/pcbook/sample"
)

func TestWriteProtoBufToBinaryFile(t *testing.T) {
	t.Parallel()
	binaryFile := "laptop.bin"
	jsonFile := "laptop.json"
	laptop1 := sample.NewLapTop()
	err := WriteProtoBufToBinaryFile(laptop1, binaryFile)
	if err != nil {
		t.Errorf("%s error happend", err)
	}
	laptop2 := &pb.Laptop{}
	err = ReadProtoBufFromBinaryFile(binaryFile, laptop2)
	if err != nil {
		t.Errorf("%s error happend", err)
	}
	err = WriteProtobufFileToJson(laptop1, jsonFile)
	if err != nil {
		t.Errorf("%s error happend", err)
	}
}
