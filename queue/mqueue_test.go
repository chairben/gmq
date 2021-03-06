package gmq

import "testing"

var message = []byte("TEST Message: you know, for testing...")
var queue = &Queue{QName: "test_queue"}

func TestQueueSimplePush(t *testing.T) {
	queue.Init(DEFAULT_QUEUE_CAP)
	err := queue.Push(message)
	if err != nil {
		t.Errorf("Error %T %s", err, err.Error())
	}
	if len(queue.QObj) != 1 {
		t.Error("Push failed!")
	}
}

func TestQueueSimplePop(t *testing.T) {
	ret, err := queue.Pop()
	if err != nil {
		t.Errorf("Error %T %s", err, err.Error())
	}
	if len(ret) != len(message) {
		t.Errorf("Message pop'd from queue incomplete! \n"+
			"message: %d \n"+
			"returned: %d", len(message), len(ret))
	}
}

func TestQueueSequentialPush(t *testing.T) {
	queue.Init(DEFAULT_QUEUE_CAP)
	for i := 0; i < 10; i++ {
		err := queue.Push(message)
		if err != nil {
			t.Errorf("Error %T %s", err, err.Error())
		}
	}
	if len(queue.QObj) < 10 {
		t.Errorf("Push failed for some object! Length of queue %d", len(queue.QObj))
	}
}

func TestQueueSequentialPop(t *testing.T) {
	for i := 0; i < 10; i++ {
		ret, err := queue.Pop()
		if err != nil {
			t.Errorf("Error %T %s", err, err.Error())
		}
		if len(ret) != len(message) {
			t.Errorf("Message pop'd from queue incomplete! \n"+
				"message: %d \n"+
				"returned: %d", len(message), len(ret))
		}
	}
}

func TestQueueConcurrentPush(t *testing.T) {
	queue.Init(DEFAULT_QUEUE_CAP)
	for i := 0; i < 10; i++ {
		go func() {
			err := queue.Push(message)
			if err != nil {
				t.Errorf("Error %T %s", err, err.Error())
			}
		}()
	}
}

func TestQueueConcurrentPop(t *testing.T) {
	for i := 0; i < 10; i++ {
		go func() {
			ret, err := queue.Pop()
			if err != nil {
				t.Errorf("Error %T %s", err, err)
			}
			if len(ret) != len(message) {
				t.Errorf("Message pop'd from queue incomplete! \n"+
					"message: %d \n"+
					"returned: %d", len(message), len(ret))
			}
		}()
	}
}
