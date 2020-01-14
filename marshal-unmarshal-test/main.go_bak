package main

import (
	"encoding"
	"encoding/binary"
	"encoding/json"
	"fmt"
)

type PDUData interface {
	Size() int
	encoding.BinaryMarshaler
}

type PDUThing struct {
	Thing PDUData `json:"thing"`
}

type StartQuantData struct {
	StartingAddress uint16 `json:"starting-address,omitempty"`
	Quantity        uint16 `json:"quantity,omitempty"`
}

func (sqd StartQuantData) MarshalBinary() ([]byte, error) {
	buf := make([]byte, sqd.Size())
	binary.BigEndian.PutUint16(buf[0:], sqd.StartingAddress)
	binary.BigEndian.PutUint16(buf[2:], sqd.Quantity)

	return buf, nil
}

func (sqd StartQuantData) Size() int {
	return binary.Size(sqd)
}

func (sqd *StartQuantData) MarshalJSON() ([]byte, error) {
	m := make(map[string]interface{})
	m["type"] = "start-quant-data"
	m["starting-address"] = sqd.StartingAddress
	m["quantity"] = sqd.Quantity
	return json.Marshal(m)
}

func (pdu *PDUThing) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("objMap", objMap["thing"])
	var rm *json.RawMessage
	err = json.Unmarshal(*objMap["thing"], &rm)
	if err != nil {
		fmt.Println(err)
	}

	var m map[string]interface{}
	err = json.Unmarshal(*rm, &m)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("m", m)
	if m["type"] == "start-quant-data" {
		var sqd = StartQuantData{}
		err := json.Unmarshal(*rm, &sqd)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("sqd", sqd)
	}
	return nil
}

func main() {
	sqd1 := &StartQuantData{
		StartingAddress: 0x0001,
		Quantity:        0x0002,
	}
	fmt.Println("sqd", sqd1)
	pduthing := PDUThing{
		Thing: sqd1,
	}

	blob1, err := json.Marshal(pduthing)
	fmt.Println("blob1", string(blob1))

	newpdu := PDUThing{}

	err = json.Unmarshal(blob1, &newpdu)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("newpdu", newpdu.Thing)
	// var got1 = StartQuantData{}
	// err := json.Unmarshal(blob1, &got1)
	// if err != nil {
	//	fmt.Println(err)
	// }
	// fmt.Println(got1)
}
