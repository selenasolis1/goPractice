package specx

import (
	"encoding"
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"fmt"
)

const (
	// ReadCoils is function code 1
	ReadCoils byte = 0x01

	// WriteMultipleCoils is function code 15
	WriteMultipleCoils = 0x0F
)

// PDU represents the modbus protocol data unit.
type PDU struct {
	FunctionCode byte    `json:"function-code,omitempty"`
	Data         PDUData `json:"data,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for PDUs.
func (pdu PDU) MarshalJSON() ([]byte, error) {
	m := make(map[string]interface{})
	m["function-code"] = pdu.FunctionCode

	// See https://tour.golang.org/methods/16 for type switching
	switch v := pdu.Data.(type) {
	case StartQuantData:
		m["type"] = "start-quant-data"
		m["starting-address"] = v.StartingAddress
		m["quantity"] = v.Quantity
	case StartWriteMultData:
		m["type"] = "start-write-mult-data"
		m["starting-address"] = v.StartingAddress
		m["quantity"] = v.Quantity
		m["byte-count"] = v.ByteCount
		m["write-data"] = v.WriteData
	default:
		m["type"] = "unknown"
	}

	return json.Marshal(m)
}

// UnmarshalJSON implements the json.Unmarshaler interface for
// PDUs.
func (pdu *PDU) UnmarshalJSON(blob []byte) error {
	// After looking at the blob we can unmarshal this into a map
	// because we see it is just a set of key-value pairs.
	m := make(map[string]interface{})
	err := json.Unmarshal(blob, &m)
	if err != nil {
		return err
	}

	// Now that we have a map we can start making assignments to the
	// PDU members, but we have to make type assertions on the types
	// since the map just holds them as empty interfaces.
	i, ok := m["function-code"]
	if !ok {
		return fmt.Errorf("pdu json unmarshal: no function-code key")
	}
	f, ok := i.(float64)
	if !ok {
		return fmt.Errorf("pdu json unmarshal: unable to assert function code as float64")
	}
	b := byte(f)
	pdu.FunctionCode = b

	t, ok := m["type"]
	if !ok {
		return fmt.Errorf("pdu unmarshal json: no pdu-type information")
	}
	switch t.(string) {
	case "start-quant-data":
		pdu.Data = StartQuantData{
			StartingAddress: uint16(m["starting-address"].(float64)),
			Quantity:        uint16(m["quantity"].(float64)),
		}
	case "start-write-mult-data":
		str := fmt.Sprint(m["write-data"])
		b, err := base64.StdEncoding.DecodeString(str)
		if err != nil {
			return err
		}
		pdu.Data = StartWriteMultData{
			StartingAddress: uint16(m["starting-address"].(float64)),
			Quantity:        uint16(m["quantity"].(float64)),
			ByteCount:       byte(m["byte-count"].(float64)),
			WriteData:       b,
		}

	default:
		pdu.Data = nil
	}
	return nil
}

// PDUData is the interface for types out of the specification
// that are structured modbus data packets.
type PDUData interface {
	Size() int
	encoding.BinaryMarshaler
}

// StartQuantData is the PDUData type associated with function codes that provide
// as StartingAddress and a Quantity of registers.  For example, this type is
// used by the function codes for ReadCoils, ReadDiscreteInputs, etc.
type StartQuantData struct {
	StartingAddress uint16 `json:"starting-address,omitempty"`
	Quantity        uint16 `json:"quantity,omitempty"`
}

// Size implements the PDUData interface for StartQuantData.
func (sqd StartQuantData) Size() int {
	return binary.Size(sqd)
}

// MarshalBinary implements the encoding.BinaryMarshaler interface for
// StartQuantData.
func (sqd StartQuantData) MarshalBinary() ([]byte, error) {
	buf := make([]byte, sqd.Size())
	binary.BigEndian.PutUint16(buf[0:], sqd.StartingAddress)
	binary.BigEndian.PutUint16(buf[2:], sqd.Quantity)

	return buf, nil
}

// StartWriteMultData is the PDUData type associated with function codes that
// provide a starting address and a quantity to range over many addresses, but also
// provide data that should be written to those addresses.  This type is
// used by the function code requests for WriteMultipleCoils and WriteMultipleRegisters.
type StartWriteMultData struct {
	StartingAddress uint16 `json:"starting-address,omitempty"`
	Quantity        uint16 `json:"quantity,omitempty"`
	ByteCount       byte   `json:"byte-count,omitempty"`
	WriteData       []byte `json:"write-data,omitempty"`
}

// Size implements the PDUData interface for StartWriteMultData.
func (swmd StartWriteMultData) Size() int {
	saSize := binary.Size(swmd.StartingAddress)
	qSize := binary.Size(swmd.Quantity)
	bcSize := binary.Size(swmd.ByteCount)
	wdSize := binary.Size(swmd.WriteData)
	return saSize + qSize + bcSize + wdSize
}

// MarshalBinary implements the encoding.BinaryMarshaler interface for
// StartWriteMultData.
func (swmd StartWriteMultData) MarshalBinary() ([]byte, error) {
	buf := make([]byte, swmd.Size())
	binary.BigEndian.PutUint16(buf[0:], swmd.StartingAddress)
	binary.BigEndian.PutUint16(buf[2:], swmd.Quantity)
	buf[4] = swmd.ByteCount
	for i, wd := range swmd.WriteData {
		start := 5
		idx := start + i
		buf[idx] = wd
	}
	return buf, nil
}

// ReadCoilsReq provides a specification compliant name
// for a PDUData implementation that supports the FunctionCode ReadCoils.
func ReadCoilsReq(startingAddress uint16, quantity uint16) StartQuantData {
	return StartQuantData{startingAddress, quantity}
}

// WriteMultipleCoilsReq provides a specification compliant name
// for a PDUData implementation that supports the FunctionCode WriteMultipleCoilsReq
func WriteMultipleCoilsReq(startingAddress uint16, quantity uint16, byteCount byte, writeData []byte) StartWriteMultData {
	return StartWriteMultData{startingAddress, quantity, byteCount, writeData}
}
