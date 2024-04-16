package todo

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

type TodosMock [2]Item

func TestAddV2_thenOk(t *testing.T) {

	i := Item{
		Task:        "",
		Description: "",
		Done:        true,
		CreatedAt:   time.Now(),
		CompletedAt: time.Now(),
	}

	it := ItemTemp{
		Title:       "titleTemp",
		Description: "titleTemp",
	}

	var x TodosMock
	x[0] = Item{
		Task:        "",
		Description: "",
		Done:        true,
		CreatedAt:   time.Now(),
		CompletedAt: time.Now(),
	}
	x[1] = Item{
		Task:        "",
		Description: "",
		Done:        true,
		CreatedAt:   time.Now(),
		CompletedAt: time.Now(),
	}

	to := Todos{i}
	to.Addv2(it)
	res := len(to)
	expected := len(x)

	if res != expected {
		t.Errorf("res %v expected %v", res, expected)
	}
}

func TestComplete_thenOk(t *testing.T) {

	i := Item{
		Task:        "",
		Description: "",
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Now(),
	}

	to := Todos{i}

	to.Complete(1)
	expected := i.Done

	if expected != false {
		t.Errorf("res %v expected %v", false, expected)
	}
}

func TestComplete_thenError(t *testing.T) {
	i := Item{
		Task:        "",
		Description: "",
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Now(),
	}

	to := Todos{i}

	res := to.Complete(-1)
	expected := errors.New("invalid index")

	if i.Done != false {
		t.Errorf("res %v expected %v", false, i.Done)
	}

	if res.Error() != expected.Error() {
		t.Errorf("res '%v' expected '%v'", res, expected)
	}
}

func TestDelete_thenOk(t *testing.T) {
	i := Item{
		Task:        "",
		Description: "",
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Now(),
	}

	to := Todos{i}

	to.Delete(1)
	expected := len(to)

	if expected != 0 {
		t.Errorf("res %v expected %v", 1, expected)
	}
}

func TestDelete_thenError(t *testing.T) {
	i := Item{
		Task:        "",
		Description: "",
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Now(),
	}

	to := Todos{i}

	res := to.Delete(-1)
	expected := errors.New("invalid index")

	if res.Error() != expected.Error() {
		t.Errorf("res '%v' expected '%v'", res, expected)
	}
}

func TestStore_thenOk(t *testing.T) {
	i := Item{
		Task:        "",
		Description: "",
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Now(),
	}

	to := Todos{i}

	res := to.Store("mock_filename")

	println(res)
	// expected := ""

	// if condition {
		
	// }
}

func TestStore_thenError(t *testing.T) {
	i := Item{
		Task:        "",
		Description: "",
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Now(),
	}

	to := Todos{i}
	// x := append(to[:0], to[0+1])

	res := x.Store("mock_error_filename")

	println(res)
}

func TestCountPending_thenPending(t *testing.T) {
	i := Item{
		Task:        "",
		Description: "",
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Now(),
	}

	to := Todos{i}

	res := to.CountPending()
	expected := 1

	if res != expected {
		t.Errorf("res %v expected %v", res, expected)
	}
}

func TestCountPending_thenDone(t *testing.T) {
	i := Item{
		Task:        "",
		Description: "",
		Done:        true,
		CreatedAt:   time.Now(),
		CompletedAt: time.Now(),
	}

	to := Todos{i}

	res := to.CountPending()
	expected := 0

	if res != expected {
		t.Errorf("res %v expected %v", res, expected)
	}
}

/***********************/

func TestPlus_thenOk(t *testing.T) {
	res := Plus(4, 6)
	expected := 10

	if res != expected {
		t.Errorf("res %v, expected %v", res, expected)
	}
}

// arg1 means argument 1 and arg2 means argument 2, and the expected stands for the 'result we expect'
type plusTest struct {
	arg1     int
	arg2     int
	expected int
}

var plusTests = []plusTest{
	plusTest{2, 3, 5},
	plusTest{8, 7, 15},
	plusTest{17, 3, 20},
	plusTest{20, 35, 55},
	plusTest{40, 34, 74},
}

func TestPlus_table(t *testing.T) {
	for _, test := range plusTests {
		if output := Plus(test.arg1, test.arg2); output != test.expected {
			t.Errorf("output %v not equal to expected value %v", output, test.expected)
		}
	}
}

func BenchmarkPlus(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Plus(4, 6)
	}
}

func ExamplePlus() {
	fmt.Println(Plus(4, 6))
	// Output: 10
}
