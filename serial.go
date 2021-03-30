package serial

import (
  "time"
  origin_serial "github.com/tarm/serial"
)

/*
usage:
import (
  "github.com/Potewo/serial"
  "log"
)
func main() {
  c := &serial.Config{
    Name: "/dev/ttyACM0",
    Baud: 9600,
  }
  s, err := serial.OpenPort(c)
  if err != nil {
    log.Fatal(err)
  }
  serial.Send(s)
}
*/

type Config struct {
  Name string
  Baud int
  ReadTimeout time.Duration
}

func OpenPort(c *Config) (*origin_serial.Port, error) {
  oc := &origin_serial.Config {Name: c.Name, Baud: c.Baud, ReadTimeout: c.ReadTimeout}
  return origin_serial.OpenPort(oc)
}

func contains(a []byte, d byte) bool {
  for _, v := range a {
    if d == v {
      return true
    }
  }
  return false
}

func send(s *origin_serial.Port, d byte) error {
  _, err := s.Write([]byte("S"))
  if err != nil {
    return err
  }

  for {
    buf := make([]byte, 128)
    n, err := s.Read(buf)
    if err != nil {
      return err
    }
    if contains(buf[:n], 'R') {
      break
    }
  }

  sendData := []byte{d}
  _, err = s.Write(sendData)
  if err != nil {
    return err
  }

  for {
    buf := make([]byte, 128)
    n, err := s.Read(buf)
    if err != nil {
      return err
    }
    if contains(buf[:n], 'O') {
      break
    }
  }
  return err
}

func Send(s *origin_serial.Port, d []byte) error {
  for _, di := range d {
    err := send(s, di)
    if err != nil {
      return err
    }
  }
  return nil
}

func Receive(s *origin_serial.Port) ([]byte, error) {
  n, err := s.Write([]byte("R"))
  if err != nil {
    return nil, err
  }

  data := make([]byte, 0)
  for {
    buf := make([]byte, 128)
    n, err = s.Read(buf)
    if err != nil {
      return nil, err
    }
    if contains(buf[:n], '\x00') {
      continue
    }
    data = append(data, buf[:n]...)
    if contains(buf, '\n') {
      break
    }
  }

  _, err = s.Write([]byte("O"))
  if err != nil {
    return nil, err
  }
  return data, nil
}
