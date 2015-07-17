package args

import "testing"

func TestNew(t *testing.T) {
	argv := "test"
	testNew1 := New()
	testNew2 := New()
	func(vals ...V) {
		v := []interface{}{testNew1.Get(vals), testNew2.Get(vals)}
		for i := range v {
			if v[i] != argv {
				t.Errorf("invalid default value: want=%v gets=%v", argv, v[i])
			}
		}
	}(testNew1(argv), testNew2(argv))
}

func TestNewDefaults(t *testing.T) {
	testNewNil := New()
	testNewOne := New(Default(1))
	testNewFlag := NewInt(Flag("testnewflag", 1, "usage"))
	func(vals ...V) {
		if v := testNewNil.Get(vals); v != nil {
			t.Errorf("invalid default value: want=nil gets=%v", v)
		}

		if v := testNewOne.Get(vals); v != 1 {
			t.Errorf("invalid default value: want=1 gets=%v", v)
		}

		if i := testNewFlag.Get(vals); i != 1 {
			t.Errorf("invalid default value: want=1 gets=%v", i)
		}
	}()
}

func TestNewInt(t *testing.T) {
	argv := 1
	testNew1 := NewInt()
	testNew2 := NewInt()
	func(vals ...V) {
		v := []int{testNew1.Get(vals), testNew2.Get(vals)}
		for i := range v {
			if v[i] != argv {
				t.Errorf("invalid value: want=%d gets=%d", argv, v[i])
			}
		}
	}(testNew1(argv), testNew2(argv))
}

func TestNewUint(t *testing.T) {
	argv := uint(1)
	testNew1 := NewUint()
	testNew2 := NewUint()
	func(vals ...V) {
		v := []uint{testNew1.Get(vals), testNew2.Get(vals)}
		for i := range v {
			if v[i] != argv {
				t.Errorf("invalid value: want=%d gets=%d", argv, v[i])
			}
		}
	}(testNew1(argv), testNew2(argv))
}
