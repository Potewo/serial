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
    ReadTimeout: time.Second * 1,
  }

  s, err := OpenPort(c)
  if err != nil {
    t.Fatal(err)
  }
  if err := Send(s, byte(0x3e)); err != nil {
    t.Fatal(err)
  }
}
