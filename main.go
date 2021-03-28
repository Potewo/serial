package main

import (
  "log"
  "github.com/tarm/serial"
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
  receive(s)
}

func contains(a []byte, d byte) bool {
  for _, v := range a {
    if d == v {
      return true
    }
  }
  return false
}

func send(s *serial.Port) {
  _, err := s.Write([]byte("S"))
  if err != nil {
    log.Fatal(err)
  }

  for {
    buf := make([]byte, 128)
    n, err := s.Read(buf)
    if err != nil {
      log.Fatal(err)
    }
    if contains(buf[:n], 'R') {
      break
    }
  }

  sendData := []byte("3")
  _, err = s.Write(sendData)
  if err != nil {
    log.Fatal(err)
  }

  for {
    buf := make([]byte, 128)
    n, err := s.Read(buf)
    if err != nil {
      log.Fatal(err)
    }
    if contains(buf[:n], 'O') {
      break
    }
  }
}

func receive(s *serial.Port) {
  n, err := s.Write([]byte("R"))
  if err != nil {
    log.Fatal(err)
  }

  data := make([]byte, 0)
  for {
    buf := make([]byte, 128)
    n, err = s.Read(buf)
    if err != nil {
      log.Fatal(err)
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
    log.Fatal(err)
  }
}
