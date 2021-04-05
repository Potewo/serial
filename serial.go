package serial

import (
	origin_serial "github.com/tarm/serial"
	"time"
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
	Name        string
	Baud        int
	ReadTimeout time.Duration
}

func OpenPort(c *Config) (*origin_serial.Port, error) {
	oc := &origin_serial.Config{Name: c.Name, Baud: c.Baud, ReadTimeout: c.ReadTimeout}
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

	err = wait(s, 'R')
	if err != nil {
		return err
	}

	sendData := []byte{d}
	_, err = s.Write(sendData)
	if err != nil {
		return err
	}

	err = wait(s, 'O')
	if err != nil {
		return err
	}
	return nil
}

func wait(s *origin_serial.Port, b byte) error {
	for {
		buf := make([]byte, 128)
		n, err := s.Read(buf)
		if err != nil {
			return err
		}
		if contains(buf[:n], b) {
			break
		}
	}
	return nil
}

func Send(s *origin_serial.Port, d []byte) error {
	// send S(send) status
	_, err := s.Write([]byte("S"))
	if err != nil {
		return err
	}
	// send length of data
	_, err = s.Write([]byte {uint8(len(d))})
	if err != nil {
		return err
	}
	err = wait(s, 'O')
	if err != nil {
		return err
	}

	// send main data
	_, err = s.Write(d)

	err = wait(s, 'O')
	if err != nil {
		return err
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
