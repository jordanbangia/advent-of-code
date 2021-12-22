package main

import (
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"

	"github.com/jordanbangia/advent-of-code/goutil"
)

func main() {
	input, err := goutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	// i := "9C0141080250320F1802104A08"
	i := input[0]

	data, err := hex.DecodeString(i)
	if err != nil {
		return
	}

	packet, _, _ := parseHexTransmission(byteArrToString(data))
	// fmt.Printf("%+v\n", packet)
	// for _, sp := range packet.SubPackets {
	// 	fmt.Printf("%+v\n", sp)
	// }
	fmt.Println(packet.VersionSum())
	fmt.Println(packet.Evaluate())

}

type Packet struct {
	Version int
	TypeID  int

	Literal int

	LengthTypeID int
	SubPackets   []*Packet
}

func (p *Packet) VersionSum() int {
	subPacketSums := 0
	for _, sp := range p.SubPackets {
		subPacketSums += sp.VersionSum()
	}

	return subPacketSums + p.Version
}

func (p *Packet) Evaluate() int {
	switch p.TypeID {
	case 4:
		return p.Literal
	case 0:
		total := 0
		for _, sb := range p.SubPackets {
			total += sb.Evaluate()
		}
		return total
	case 1:
		product := 1
		for _, sb := range p.SubPackets {
			product *= sb.Evaluate()
		}
		return product
	case 2:
		min := -1
		for _, sb := range p.SubPackets {
			if min == -1 {
				min = sb.Evaluate()
			} else {
				min = goutil.Min(min, sb.Evaluate())
			}
		}
		return min
	case 3:
		max := -1
		for _, sb := range p.SubPackets {
			max = goutil.Max(max, sb.Evaluate())
		}
		return max
	case 5:
		if p.SubPackets[0].Evaluate() > p.SubPackets[1].Evaluate() {
			return 1
		}
		return 0
	case 6:
		if p.SubPackets[0].Evaluate() < p.SubPackets[1].Evaluate() {
			return 1
		}
		return 0
	case 7:
		if p.SubPackets[0].Evaluate() == p.SubPackets[1].Evaluate() {
			return 1
		}
		return 0
	default:
		fmt.Println("SOMETHINGS NOT RIGHT")
		return -1
	}
}

func parseHexTransmission(data string) (*Packet, int, error) {
	if len(data) < 6 {
		return nil, 0, fmt.Errorf("packet isn't long enough")
	}

	// fmt.Println(data)
	packet := &Packet{
		Version: stringBinToInt(data[:3]),
		TypeID:  stringBinToInt(data[3:6]),
	}

	offSet := 6
	if packet.TypeID == 4 {
		literal, parsedOffset := parseLiteral(data[6:])
		packet.Literal = literal
		offSet += parsedOffset
	} else {
		offSet = 7
		// this is an operator packet, so parse out the length type ID
		packet.LengthTypeID = stringBinToInt(data[6:7])

		packet.SubPackets = []*Packet{}

		if packet.LengthTypeID == 0 {
			subPacketLength := stringBinToInt(data[7 : 7+15])
			l := subPacketLength
			start := 7 + 15
			offSet += 15
			for l > 0 {
				p, processed, err := parseHexTransmission(data[start:])
				if err != nil {
					break
				}
				packet.SubPackets = append(packet.SubPackets, p)
				l -= processed
				start += processed
				offSet += processed
			}
		} else {
			subPacketCount := stringBinToInt(data[7 : 7+11])
			offSet += 11
			start := 7 + 11
			for subPacketCount > 0 {
				p, processed, err := parseHexTransmission(data[start:])
				if err != nil {
					break
				}
				packet.SubPackets = append(packet.SubPackets, p)
				start += processed
				offSet += processed
				subPacketCount--
			}
		}
	}

	// fmt.Printf("%+v\n", packet)

	return packet, offSet, nil
}

func parseLiteral(data string) (int, int) {
	i := 0
	lastGroupFound := false

	var literalBits strings.Builder
	for !lastGroupFound {
		if data[i] == '0' {
			lastGroupFound = true
		}
		i += 1
		write := data[i : i+4]
		literalBits.WriteString(write)
		i += 4
	}

	literalVal, _ := strconv.ParseInt(literalBits.String(), 2, 64)
	return int(literalVal), i
}

func byteArrToString(data []byte) string {
	var str strings.Builder
	for _, d := range data {
		str.WriteString(fmt.Sprintf("%08b", d))
	}
	return str.String()
}

func stringBinToInt(data string) int {
	subPacketLenght, _ := strconv.ParseInt(data, 2, 64)
	return int(subPacketLenght)
}
