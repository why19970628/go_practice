package observer

import "testing"

func ExampleObserver() {
	//subject := NewSubject()
	//reader1 := NewReader("reader1")
	//reader2 := NewReader("reader2")
	//reader3 := NewReader("reader3")
	//subject.Attach(reader1)
	//subject.Attach(reader2)
	//subject.Attach(reader3)
	//
	//subject.UpdateContext("observer mode")
	// Output:
	// reader1 receive observer mode
	// reader2 receive observer mode
	// reader3 receive observer mode
}

func TestExampleObserver(t *testing.T) {
	subject := NewNewSubject()
	reader1 := NewObj("reader1")
	reader2 := NewObj("reader2")
	reader3 := NewObj("reader3")

	subject.add(reader1)
	subject.add(reader2)
	subject.add(reader3)

	subject.update("aaa")
}
