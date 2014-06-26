// Package er101 is a library for parsing snapshots exported from the
// Orthogonal Devices ER-101: Indexed Quad Sequencer
//
// See http://www.orthogonaldevices.com/er-101 for more details.
package er101

import (
	"encoding/xml"
	"io"
)

type ER101 struct {
	Snapshots []Snapshot `xml:"snapshots>Snapshot"`
}

type Snapshot struct {
	Transform Transform `xml:"transform"`
	Tracks    []Track   `xml:"tracks>Track"`
}

type Transform struct {
	CVAMode      int `xml:"cvA_mode"`
	CVBMode      int `xml:"cvB_mode"`
	DurationMode int `xml:"duration_mode"`
	GateMode     int `xml:"gate_mode"`
	CVA          int `xml:"cvA"`
	CVB          int `xml:"cvB"`
	Duration     int `xml:"duration"`
	Gate         int `xml:"gate"`
}

type Track struct {
	Transform Transform `xml:"transform"`
	Index     int       `xml:"index"`
	Patterns  []Pattern `xml:"patterns>Pattern"`
	LoopStart Loop      `xml:"loop_start"`
	LoopEnd   Loop      `xml:"loop_end"`
	VoltagesA []int     `xml:"voltagesA>int"`
	VoltagesB []int     `xml:"voltagesB>int"`
	Options   string    `xml:"options"`
	PPQN      int       `xml:"ppqn"`
}

type Pattern struct {
	Index   int    `xml:"index"`
	Steps   []Step `xml:"steps>Step"`
	Options string `xml:"options"`
}

type Step struct {
	Index     int  `xml:"index"`
	CVAIndex  int  `xml:"cvA_index"`
	CVASmooth bool `xml:"cvA_smooth"`
	CVBIndex  int  `xml:"cvB_index"`
	CVBSmooth bool `xml:"cvB_smooth"`
	Duration  int  `xml:"duration"`
	Gate      int  `xml:"gate"`
}

type Loop struct {
	Pattern int `xml:"pattern"`
	Step    int `xml:"step"`
}

func Parse(r io.Reader) (ER101, error) {
	dec := xml.NewDecoder(r)
	e := ER101{}
	err := dec.Decode(&e)
	return e, err
}
