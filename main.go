package serial

import (
  "log"
  "time"
  origin_serial "github.com/tarm/serial"
)

// usage:
// func main() {
//   c := &serial.Config{
//     Name: "/dev/ttyACM0",
//     Baud: 9600,
//   }
//   s, err := serial.OpenPort(c)
//   if err != nil {
//     log.Fatal(err)
//   }
//   send(s)
// }

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

func Send(s *origin_serial.Port) error {
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

  sendData := []byte{0xe3}
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

func Receive(s *origin_serial.Port) error {
  n, err := s.Write([]byte("R"))
  if err != nil {
    return err
  }

  data := make([]byte, 0)
  for {
    buf := make([]byte, 128)
    n, err = s.Read(buf)
    if err != nil {
      return err
    }
    if contains(buf[:n], '\x00') {
      continue
    }
    data = append(data, buf[:n]...)
    if contains(buf, '\n') {
      break
    }
  }
  log.Printf("%q", data)

  _, err = s.Write([]byte("O"))
  if err != nil {
    return err
  }
  return nil
}
