package systemcpu

import "testing"

func testfake(t *testing.T) {
	var srun SysRun
	var fake bool

	srun.Update()
	if !fake {
		t.Errorf("No test.")
	}
}
