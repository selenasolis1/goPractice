package main

import (
	"encoding/json"

	"github.com/selenasolis1/goPractice/marshal-unmarshal-test/specx"
)

func main() {
	pdu := specx.PDU{
		FunctionCode: specx.ReadCoils,
		Data:         specx.ReadCoilsReq(0x0013, 0x0002),
	}

	blob, err := json.Marshal(pdu)
	if err != nil {
		panic(err)
	}

	pdu2 := specx.PDU{
		FunctionCode: specx.WriteMultipleCoils,
		Data:         specx.WriteMultipleCoilsReq(0x0013, 0x0002, 1, []byte{0x30}),
	}

	blob2, err := json.Marshal(pdu2)
	if err != nil {
		panic(err)
	}

	var pdu3 = specx.PDU{}
	err = json.Unmarshal(blob, &pdu3)
	if err != nil {
		panic(err)
	}

	var pdu4 = specx.PDU{}
	err = json.Unmarshal(blob2, &pdu4)
	if err != nil {
		panic(err)
	}
}
