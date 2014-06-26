package er101

import (
	"os"
	"testing"
)

func TestParse(t *testing.T) {
	f, err := os.Open("er-101-example-snapshots.xml")
	if err != nil {
		t.Fatal(err)
	}
	e, err := Parse(f)
	if err != nil {
		t.Fatal(err)
	}

	if got, want := len(e.Snapshots), 16; got != want {
		t.Fatalf("got %d snapshots, want %d", got, want)
	}

	snap0 := e.Snapshots[0]
	snap0WantTransform := Transform{
		CVAMode:      0,
		CVBMode:      0,
		DurationMode: 1,
		GateMode:     1,
		CVA:          12,
		CVB:          0,
		Duration:     1,
		Gate:         1,
	}
	if snap0.Transform != snap0WantTransform {
		t.Errorf("Snapshot 0 transform got %v, want %v", snap0.Transform, snap0WantTransform)
	}

	tracks0 := snap0.Tracks
	if got, want := len(tracks0), 4; got != want {
		t.Errorf("Snapshot 0 got %d tracks, want %d", got, want)
	}

	track0 := snap0.Tracks[0]
	track0WantTransform := snap0WantTransform
	if got, want := track0.Transform, track0WantTransform; got != want {
		t.Errorf("Snapshot 0 track 0 transform got %v, want %v", got, want)
	}
	if track0.Index != 0 {
		t.Errorf("Snapshot 0 track 0 index got %d, want %d", track0.Index, 0)
	}

	if len(track0.Patterns) != 2 {
		t.Errorf("Snapshot 0 track 0 got %d patterns, want %d", len(track0.Patterns), 2)
	}
	pattern1 := track0.Patterns[1]

	if pattern1.Index != 1 {
		t.Errorf("Snapshot 0 track 0 pattern 1 got index %d, want %d", pattern1.Index, 1)
	}
	if len(pattern1.Steps) != 4 {
		t.Errorf("Snapshot 0 track 0 pattern 1 got %d steps, want %d", len(pattern1.Steps), 4)
	}
	if pattern1.Options != "None" {
		t.Errorf("Snapshot 0 track 0 pattern 1 options got %s, want %s", pattern1.Options, "None")
	}

	step2 := pattern1.Steps[2]
	wantStep2 := Step{
		Index:     2,
		CVAIndex:  26,
		CVASmooth: false,
		CVBIndex:  10,
		CVBSmooth: false,
		Duration:  16,
		Gate:      8,
	}
	if step2 != wantStep2 {
		t.Errorf("Snap 0 track 0 pattern 1 step 2 got %v, want %v", step2, wantStep2)
	}

	wantLoopStart := Loop{
		Pattern: 0,
		Step:    2,
	}
	if track0.LoopStart != wantLoopStart {
		t.Errorf("Snap 0 track 0 loop start got %v, want %v", track0.LoopStart, wantLoopStart)
	}
	wantLoopEnd := Loop{
		Pattern: 1,
		Step:    3,
	}
	if track0.LoopEnd != wantLoopEnd {
		t.Errorf("Snap 0 track 0 loop end got %v, want %v", track0.LoopEnd, wantLoopEnd)
	}

	if len(track0.VoltagesA) != 100 {
		t.Errorf("track 0 VoltagesA length got %d, want %d", len(track0.VoltagesA), 100)
	}
	if len(track0.VoltagesB) != 100 {
		t.Errorf("track 0 VoltagesB length got %d, want %d", len(track0.VoltagesB), 100)
	}

	if track0.Options != "Note_Display_A" {
		t.Errorf("track 0 options got %s, want %s", track0.Options, "Note_Display_A")
	}
	if track0.PPQN != 0 {
		t.Errorf("track 0 PPQN got %d, want %d", track0.PPQN, 0)
	}
}
