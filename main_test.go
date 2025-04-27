package ntfyclient

import (
	"runtime/debug"
	"testing"
)

func TestSend(t *testing.T) {
	client := NewClient("https://ntfy.purpleglass.ru/test", nil)
	client.Send("Hello, World!")
}

func TestSendPanic(t *testing.T) {
	client := NewClient("https://ntfy.purpleglass.ru/test", nil)
	defer func() {
		if r := recover(); r != nil {
			stack := string(debug.Stack())
			client.SendError("Случилось что-то непредвиденное", stack)
		}
	}()
	panic("test")
}

func TestSendWarning(t *testing.T) {
	client := NewClient("https://ntfy.purpleglass.ru/test", nil)
	client.SendWarning("Случилось что-то непредвиденное")
}

func TestSendDebug(t *testing.T) {
	client := NewClient("https://ntfy.purpleglass.ru/test", nil)
	client.SendDebug("Просто отладочное сообщение")
}
