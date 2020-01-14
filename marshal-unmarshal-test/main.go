package main

import (
	"encoding/json"
	"fmt"

	"github.com/selenasolis1/goPractice/marshal-unmarshal-test/specx"
)

func main() {
	pdu := specx.PDU{
		FunctionCode: specx.ReadCoils,
		Data:         specx.ReadCoilsReq(0x0013, 0x0002),
	}
	fmt.Println("pdu is", pdu)

	blob, err := json.Marshal(pdu)
	if err != nil {
		panic(err)
	}
	fmt.Printf("pdu after marshal is: %s\n", string(blob))

	pdu2 := specx.PDU{
		FunctionCode: specx.WriteMultipleCoils,
		Data:         specx.WriteMultipleCoilsReq(0x0013, 0x0002, 1, []byte{0x30}),
	}
	fmt.Println("pdu2 is", pdu2)

	blob2, err := json.Marshal(pdu2)
	if err != nil {
		panic(err)
	}
	fmt.Printf("pdu2 after marshal is: %s\n", string(blob2))

	var pdu3 = specx.PDU{}
	err = json.Unmarshal(blob, &pdu3)
	if err != nil {
		panic(err)
	}
	fmt.Printf("pdu3 unmarshaled from blob is: %+v\n", pdu3)

	var pdu4 = specx.PDU{}
	err = json.Unmarshal(blob2, &pdu4)
	if err != nil {
		panic(err)
	}
	fmt.Printf("pdu4 unmarshaled from blob2 is: %+v\n", pdu4)
}
