package codecs

import (
	"fmt"
	"reflect"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/bson/bsonrw"
)

// func UserEncodeValue(_ bsoncodec.EncodeContext, vw bsonrw.ValueWriter, val reflect.Value) error {
// 	fmt.Println("passou pelo Encoder")

// 	switch val.Kind() {
// 	case reflect.String:
// 		fmt.Println("É uma string")
// 	case reflect.Struct:
// 		fmt.Println("É uma struct")
// 	}
// 	return nil
// }

func UserDecodeValue(_ bsoncodec.DecodeContext, vr bsonrw.ValueReader, val reflect.Value) error {
	if !val.CanSet() || val.Kind() != reflect.Struct {
		return ValueDecoderError{Name: "codecs.UserDecodeValue", Kinds: []reflect.Kind{reflect.Struct}, Received: val}
	}

	val.Elem()

	fmt.Println("passou pelo Decoder")
	return nil
}

// func StringEncodeValue(_ bsoncodec.EncodeContext, vw bsonrw.ValueWriter, val reflect.Value) error {
// 	switch val.Kind() {
// 	case reflect.String:
// 		fmt.Printf("É uma string: %v", val.Interface().(string))
// 		vw.WriteString(fmt.Sprintf("É uma string: %v", val.Interface()))
// 	case reflect.Struct:
// 		fmt.Println("É uma struct")
// 	}
// 	return nil
// }

func CustomRegistry() *bsoncodec.Registry {
	// Create a registry that decodes nothing.
	registry := bson.NewRegistry()
	// registry.RegisterTypeEncoder(reflect.TypeOf(model.User{}), bsoncodec.ValueEncoderFunc(UserEncodeValue))
	// registry.RegisterTypeDecoder(reflect.TypeOf(model.User{}), bsoncodec.ValueDecoderFunc(UserDecodeValue))
	// registry.RegisterTypeEncoder(reflect.TypeOf(string("")), bsoncodec.ValueEncoderFunc(StringEncodeValue))
	return registry
}
