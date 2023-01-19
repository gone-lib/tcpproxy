package tcpproxy

import (
	"io"
	"testing"
)

func TestListenerAccept(t *testing.T) {
	tl := new(TargetListener)
	ch := make(chan interface{}, 1)
	go func() {
		for {
			conn, err := tl.Accept()
			if err != nil {
				ch <- err
				return
			}
			ch <- conn
		}
	}()

	for i := 0; i < 3; i++ {
		conn := new(Conn)
		tl.HandleConn(conn)
		got := <-ch
		if got != conn {
			t.Errorf("Accept conn = %v; want %v", got, conn)
		}
	}
	tl.Close()
	got := <-ch
	if got != io.EOF {
		t.Errorf("Accept error post-Close = %v; want io.EOF", got)
	}
}
