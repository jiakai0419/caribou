package arithmetic

import "testing"

func TestMaxI64(t *testing.T) {
	if MaxI64() != LOWER_BOUND_I64 {
		t.FailNow()
	}

	if MaxI64(1024) != 1024 {
		t.FailNow()
	}

	if MaxI64(3, 5) != 5 {
		t.FailNow()
	}

	if MaxI64(1, 3, 5) != 5 {
		t.FailNow()
	}
}

func TestMinI64(t *testing.T) {
	if MinI64() != UPPER_BOUND_I64 {
		t.FailNow()
	}

	if MinI64(1024) != 1024 {
		t.FailNow()
	}

	if MinI64(3, 5) != 3 {
		t.FailNow()
	}

	if MinI64(1, 3, 5) != 1 {
		t.FailNow()
	}
}

func TestMaxI(t *testing.T) {
	if MaxI() != LOWER_BOUND_I {
		t.FailNow()
	}

	if MaxI(1024) != 1024 {
		t.FailNow()
	}

	if MaxI(3, 5) != 5 {
		t.FailNow()
	}

	if MaxI(1, 3, 5) != 5 {
		t.FailNow()
	}
}

func TestMinI(t *testing.T) {
	if MinI() != UPPER_BOUND_I {
		t.FailNow()
	}

	if MinI(1024) != 1024 {
		t.FailNow()
	}

	if MinI(3, 5) != 3 {
		t.FailNow()
	}

	if MinI(1, 3, 5) != 1 {
		t.FailNow()
	}
}
