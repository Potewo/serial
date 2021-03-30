package serial

import (
  "testing"
  "time"
)

func TestReceiveSuccess(t *testing.T) {
  c := &Config {
    Name: "/dev/ttyACM0",
    Baud: 9600,
    ReadTimeout: time.Second * 1,
  }

  s, err := OpenPort(c)
  if err != nil {
    t.Fatal(err)
  }
  if err := Receive(s); err != nil {
    t.Fatal(err)
  }
}

func TestSendSuccess(t *testing.T) {
  c := &Config {
    Name: "/dev/ttyACM0",
    Baud: 9600,
    ReadTimeout: time.Second * 1,
  }

  s, err := OpenPort(c)
  if err != nil {
    t.Fatal(err)
  }
  if err := Send(s); err != nil {
    t.Fatal(err)
  }
}
