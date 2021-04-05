package serial

import (
  "testing"
  "time"
)

func TestReceiveSuccess(t *testing.T) {
  c := &Config {
    Name: "/dev/ttyACM0",
    Baud: 9600,
	ReadTimeout: time.Second,
  }

  s, err := OpenPort(c)
  if err != nil {
    t.Fatal(err)
  }
  d, err := Receive(s)
  if err != nil {
    t.Fatal(err)
  } else {
    t.Logf("%q", d);
  }
}

func TestSendSuccess(t *testing.T) {
  c := &Config {
    Name: "/dev/ttyACM0",
    Baud: 9600,
    ReadTimeout: time.Second,
  }

  s, err := OpenPort(c)
  if err != nil {
    t.Fatal(err)
  }
  err = Send(s, []byte {0x3e, 0x64})
  if err != nil {
    t.Fatal(err)
  }
}
