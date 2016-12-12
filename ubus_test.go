package ubus

import "testing"

func BenchmarkOn(b *testing.B) {
	b.Run("simple", func(b *testing.B) {
		u := NewBus()

		for i := 0; i < b.N; i++ {
			u.On("a", func(_ interface{}) {})
		}
	})

	b.Run("simple_with_emit", func(b *testing.B) {
		u := NewBus()

		for i := 0; i < b.N; i++ {
			u.On("a", func(_ interface{}) {})
			u.Emit("a", "x")
		}
	})
}

func BenchmarkOnce(b *testing.B) {
	b.Run("simple", func(b *testing.B) {
		u := NewBus()

		for i := 0; i < b.N; i++ {
			u.Once("a", func(_ interface{}) {})
		}
	})

	b.Run("simple_with_emit", func(b *testing.B) {
		u := NewBus()

		for i := 0; i < b.N; i++ {
			u.Once("a", func(_ interface{}) {})
			u.Emit("a", "x")
		}
	})
}

func BenchmarkEmit(b *testing.B) {
	b.Run("simple", func(b *testing.B) {
		u := NewBus()

		u.On("a", func(_ interface{}) {})

		for i := 0; i < b.N; i++ {
			u.Emit("a", "x")
		}
	})
}

func BenchmarkOff(b *testing.B) {
	b.Run("simple", func(b *testing.B) {
		u := NewBus()

		u.On("a", func(_ interface{}) {})

		for i := 0; i < b.N; i++ {
			u.Off([]string{"a"})
		}
	})

	b.Run("multiple", func(b *testing.B) {
		u := NewBus()

		for i := 0; i < b.N; i++ {
			u.On("a", func(_ interface{}) {})
			u.Off([]string{"a"})
		}
	})
}

func BenchmarkDestroyFn(b *testing.B) {
	b.Run("simple", func(b *testing.B) {
		u := NewBus()

		for i := 0; i < b.N; i++ {
			destroy := u.On("a", func(_ interface{}) {})
			destroy()
		}
	})
}

func TestOn(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		u := NewBus()

		expectedVal := true
		called := false

		u.On("a", func(_ interface{}) {
			called = true
		})

		u.Emit("a", nil)

		if called != expectedVal {
			t.Errorf("Expected %t to equal %t", called, expectedVal)
		}
	})

	t.Run("not_so_simple", func(t *testing.T) {
		u := NewBus()

		expectedVal := true
		expectedParam := "1"
		called := false
		param := ""

		u.On("a", func(x interface{}) {
			called = true
			param = x.(string)
		})

		u.Emit("a", "1")

		if called != expectedVal {
			t.Errorf("Expected %t to equal %t", called, expectedVal)
		}

		if param != expectedParam {
			t.Errorf("Expected %s to equal %s", param, expectedParam)
		}
	})

	t.Run("times_called", func(t *testing.T) {
		u := NewBus()

		expectedVal := true
		expectedParam := "1"
		expectedCalls := 3
		called := false
		param := ""
		timesCalled := 0

		u.On("a", func(x interface{}) {
			called = true
			param = x.(string)
			timesCalled += 1
		})

		u.Emit("a", "1")
		u.Emit("a", "1")
		u.Emit("a", "1")

		if called != expectedVal {
			t.Errorf("Expected %t to equal %t", called, expectedVal)
		}

		if param != expectedParam {
			t.Errorf("Expected %s to equal %s", param, expectedParam)
		}

		if timesCalled != expectedCalls {
			t.Errorf("Expected number of calls to be %d, but got %d", expectedCalls, timesCalled)
		}
	})
}

func TestOnce(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		u := NewBus()

		expectedVal := true
		called := false

		u.Once("a", func(_ interface{}) {
			called = true
		})

		u.Emit("a", nil)

		if called != expectedVal {
			t.Errorf("Expected %t to equal %t", called, expectedVal)
		}
	})

	t.Run("not_so_simple", func(t *testing.T) {
		u := NewBus()

		expectedVal := true
		expectedParam := "1"
		called := false
		param := ""

		u.Once("a", func(x interface{}) {
			called = true
			param = x.(string)
		})

		u.Emit("a", "1")

		if called != expectedVal {
			t.Errorf("Expected %t to equal %t", called, expectedVal)
		}

		if param != expectedParam {
			t.Errorf("Expected %s to equal %s", param, expectedParam)
		}
	})

	t.Run("times_called", func(t *testing.T) {
		u := NewBus()

		expectedVal := true
		expectedParam := "1"
		expectedCalls := 1
		called := false
		param := ""
		timesCalled := 0

		u.Once("a", func(x interface{}) {
			called = true
			param = x.(string)
			timesCalled += 1
		})

		u.Emit("a", "1")
		u.Emit("a", "1")
		u.Emit("a", "1")

		if called != expectedVal {
			t.Errorf("Expected %t to equal %t", called, expectedVal)
		}

		if param != expectedParam {
			t.Errorf("Expected %s to equal %s", param, expectedParam)
		}

		if timesCalled != expectedCalls {
			t.Errorf("Expected number of calls to be %d, but got %d", expectedCalls, timesCalled)
		}
	})
}

func TestEmit(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		u := NewBus()

		expectedVal := true
		called := false

		u.On("a", func(_ interface{}) {
			called = true
		})

		u.Emit("a", nil)

		if called != expectedVal {
			t.Errorf("Expected %t to equal %t", called, expectedVal)
		}
	})

	t.Run("not_so_simple", func(t *testing.T) {
		u := NewBus()

		expectedVal := true
		expectedParam := "1"
		called := false
		param := ""

		u.On("a", func(x interface{}) {
			called = true
			param = x.(string)
		})

		u.Emit("a", "1")

		if called != expectedVal {
			t.Errorf("Expected %t to equal %t", called, expectedVal)
		}

		if param != expectedParam {
			t.Errorf("Expected %s to equal %s", param, expectedParam)
		}
	})

	t.Run("times_called", func(t *testing.T) {
		u := NewBus()

		expectedVal := true
		expectedParam := "1"
		expectedCalls := 3
		called := false
		param := ""
		timesCalled := 0

		u.On("a", func(x interface{}) {
			called = true
			param = x.(string)
			timesCalled += 1
		})

		u.Emit("a", "1")
		u.Emit("a", "1")
		u.Emit("a", "1")

		if called != expectedVal {
			t.Errorf("Expected %t to equal %t", called, expectedVal)
		}

		if param != expectedParam {
			t.Errorf("Expected %s to equal %s", param, expectedParam)
		}

		if timesCalled != expectedCalls {
			t.Errorf("Expected number of calls to be %d, but got %d", expectedCalls, timesCalled)
		}
	})

	t.Run("multiple_emits", func(t *testing.T) {
		u := NewBus()

		aCalled := false
		aCalledTimes := 0
		expectedACalled := true
		expectedACalledTimes := 1

		bCalled := false
		bCalledTimes := 0
		expectedBCalled := true
		expectedBCalledTimes := 1

		cCalled := false
		cCalledTimes := 0
		expectedCCalled := true
		expectedCCalledTimes := 1

		u.On("a", func(_ interface{}) {
			aCalled = true
			aCalledTimes += 1
		})

		u.On("b", func(_ interface{}) {
			bCalled = true
			bCalledTimes += 1
		})

		u.On("c", func(_ interface{}) {
			cCalled = true
			cCalledTimes += 1
		})

		u.Emit("a", nil)
		u.Emit("b", nil)
		u.Emit("c", nil)

		if aCalled != expectedACalled {
			t.Errorf("Expected a to have been called, but it wasn't. %t != %t", aCalled, expectedACalled)
		}

		if aCalledTimes != expectedACalledTimes {
			t.Errorf("Expected a to have been called %d times, but it was called %d.", expectedACalledTimes, aCalledTimes)
		}

		if bCalled != expectedBCalled {
			t.Errorf("Expected b to have been called, but it wasn't. %t != %t", bCalled, expectedBCalled)
		}

		if bCalledTimes != expectedBCalledTimes {
			t.Errorf("Expected b to have been called %d times, but it was called %d.", expectedBCalledTimes, bCalledTimes)
		}

		if cCalled != expectedCCalled {
			t.Errorf("Expected c to have been called, but it wasn't. %t != %t", bCalled, expectedCCalled)
		}

		if cCalledTimes != expectedCCalledTimes {
			t.Errorf("Expected c to have been called %d times, but it was called %d.", expectedCCalledTimes, cCalledTimes)
		}
	})
}

func TestOff(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		u := NewBus()

		u.Off([]string{"xyz"})
	})

	t.Run("simple", func(t *testing.T) {
		u := NewBus()

		aCalled := false
		aCalledTimes := 0
		expectedACalled := true
		expectedACalledTimes := 1

		u.On("a", func(_ interface{}) {
			aCalled = true
			aCalledTimes += 1
		})

		u.Emit("a", nil)
		u.Off([]string{"a"})
		u.Emit("a", nil)

		if aCalled != expectedACalled {
			t.Errorf("Expected a to have been called, but it wasn't. %t != %t", aCalled, expectedACalled)
		}

		if aCalledTimes != expectedACalledTimes {
			t.Errorf("Expected a to have been called %d times, but it was called %d.", expectedACalledTimes, aCalledTimes)
		}
	})

	t.Run("multiple_off", func(t *testing.T) {
		u := NewBus()

		aCalled := false
		aCalledTimes := 0
		expectedACalled := true
		expectedACalledTimes := 1

		bCalled := false
		bCalledTimes := 0
		expectedBCalled := true
		expectedBCalledTimes := 1

		cCalled := false
		cCalledTimes := 0
		expectedCCalled := true
		expectedCCalledTimes := 1

		u.On("a", func(_ interface{}) {
			aCalled = true
			aCalledTimes += 1
		})

		u.On("b", func(_ interface{}) {
			bCalled = true
			bCalledTimes += 1
		})

		u.On("c", func(_ interface{}) {
			cCalled = true
			cCalledTimes += 1
		})

		u.Emit("a", nil)
		u.Off([]string{"a"})
		u.Emit("a", nil)
		u.Emit("a", nil)

		u.Emit("b", nil)
		u.Off([]string{"b"})
		u.Emit("b", nil)
		u.Emit("b", nil)

		u.Emit("c", nil)
		u.Off([]string{"c"})

		if aCalled != expectedACalled {
			t.Errorf("Expected a to have been called, but it wasn't. %t != %t", aCalled, expectedACalled)
		}

		if aCalledTimes != expectedACalledTimes {
			t.Errorf("Expected a to have been called %d times, but it was called %d.", expectedACalledTimes, aCalledTimes)
		}

		if bCalled != expectedBCalled {
			t.Errorf("Expected b to have been called, but it wasn't. %t != %t", bCalled, expectedBCalled)
		}

		if bCalledTimes != expectedBCalledTimes {
			t.Errorf("Expected b to have been called %d times, but it was called %d.", expectedBCalledTimes, bCalledTimes)
		}

		if cCalled != expectedCCalled {
			t.Errorf("Expected c to have been called, but it wasn't. %t != %t", bCalled, expectedCCalled)
		}

		if cCalledTimes != expectedCCalledTimes {
			t.Errorf("Expected c to have been called %d times, but it was called %d.", expectedCCalledTimes, cCalledTimes)
		}
	})
}

func TestDestroyFn(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		u := NewBus()

		aCalled := false
		aCalledTimes := 0
		expectedACalled := true
		expectedACalledTimes := 1

		destroyA := u.On("a", func(_ interface{}) {
			aCalled = true
			aCalledTimes += 1
		})

		u.Emit("a", nil)
		destroyA()
		u.Emit("a", nil)

		if aCalled != expectedACalled {
			t.Errorf("Expected a to have been called, but it wasn't. %t != %t", aCalled, expectedACalled)
		}

		if aCalledTimes != expectedACalledTimes {
			t.Errorf("Expected a to have been called %d times, but it was called %d.", expectedACalledTimes, aCalledTimes)
		}
	})

	t.Run("multiple_off", func(t *testing.T) {
		u := NewBus()

		aCalled := false
		aCalledTimes := 0
		expectedACalled := true
		expectedACalledTimes := 1

		bCalled := false
		bCalledTimes := 0
		expectedBCalled := true
		expectedBCalledTimes := 1

		cCalled := false
		cCalledTimes := 0
		expectedCCalled := true
		expectedCCalledTimes := 1

		destroyA := u.On("a", func(_ interface{}) {
			aCalled = true
			aCalledTimes += 1
		})

		destroyB := u.On("b", func(_ interface{}) {
			bCalled = true
			bCalledTimes += 1
		})

		destroyC := u.On("c", func(_ interface{}) {
			cCalled = true
			cCalledTimes += 1
		})

		u.Emit("a", nil)
		destroyA()
		u.Emit("a", nil)
		u.Emit("a", nil)

		u.Emit("b", nil)
		destroyB()
		u.Emit("b", nil)
		u.Emit("b", nil)

		u.Emit("c", nil)
		destroyC()

		if aCalled != expectedACalled {
			t.Errorf("Expected a to have been called, but it wasn't. %t != %t", aCalled, expectedACalled)
		}

		if aCalledTimes != expectedACalledTimes {
			t.Errorf("Expected a to have been called %d times, but it was called %d.", expectedACalledTimes, aCalledTimes)
		}

		if bCalled != expectedBCalled {
			t.Errorf("Expected b to have been called, but it wasn't. %t != %t", bCalled, expectedBCalled)
		}

		if bCalledTimes != expectedBCalledTimes {
			t.Errorf("Expected b to have been called %d times, but it was called %d.", expectedBCalledTimes, bCalledTimes)
		}

		if cCalled != expectedCCalled {
			t.Errorf("Expected c to have been called, but it wasn't. %t != %t", bCalled, expectedCCalled)
		}

		if cCalledTimes != expectedCCalledTimes {
			t.Errorf("Expected c to have been called %d times, but it was called %d.", expectedCCalledTimes, cCalledTimes)
		}
	})
}
