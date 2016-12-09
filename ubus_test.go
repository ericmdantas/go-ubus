package ubus

import "testing"

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
			t.Errorf("Expected %v to equal %v", called, expectedVal)
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
			t.Errorf("Expected %v to equal %v", called, expectedVal)
		}

		if param != expectedParam {
			t.Errorf("Expected %v to equal %v", param, expectedParam)
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
			t.Errorf("Expected %v to equal %v", called, expectedVal)
		}

		if param != expectedParam {
			t.Errorf("Expected %v to equal %v", param, expectedParam)
		}

		if timesCalled != expectedCalls {
			t.Errorf("Expected number of calls to be %v, but got %v", expectedCalls, timesCalled)
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
			t.Errorf("Expected %v to equal %v", called, expectedVal)
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
			t.Errorf("Expected %v to equal %v", called, expectedVal)
		}

		if param != expectedParam {
			t.Errorf("Expected %v to equal %v", param, expectedParam)
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
			t.Errorf("Expected %v to equal %v", called, expectedVal)
		}

		if param != expectedParam {
			t.Errorf("Expected %v to equal %v", param, expectedParam)
		}

		if timesCalled != expectedCalls {
			t.Errorf("Expected number of calls to be %v, but got %v", expectedCalls, timesCalled)
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
			t.Errorf("Expected %v to equal %v", called, expectedVal)
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
			t.Errorf("Expected %v to equal %v", called, expectedVal)
		}

		if param != expectedParam {
			t.Errorf("Expected %v to equal %v", param, expectedParam)
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
			t.Errorf("Expected %v to equal %v", called, expectedVal)
		}

		if param != expectedParam {
			t.Errorf("Expected %v to equal %v", param, expectedParam)
		}

		if timesCalled != expectedCalls {
			t.Errorf("Expected number of calls to be %v, but got %v", expectedCalls, timesCalled)
		}
	})
}
