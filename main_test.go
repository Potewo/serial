package serial

import (
  "testing"
)

func SendTest(t *testing.T) {
  c := &Config {
    Name: "/dev/ttyACM0",
    Baud: 9600,
  }

  s, err := OpenPort(c)
  if err != nil {
    t.Fatal(err)
  }
  Send(s)
}
