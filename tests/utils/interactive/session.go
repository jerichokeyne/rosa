package interactive

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/acarl005/stripansi"
	"github.com/creack/pty"
)

// InteractiveSession manages a PTY session for testing interactive commands
type InteractiveSession struct {
	cmd         *exec.Cmd
	ptmx        *os.File
	reader      *bufio.Reader
	output      strings.Builder
	outputChan  chan string
	stopReading chan bool
}

// NewInteractiveSession creates a new interactive testing session
func NewInteractiveSession(command string, args ...string) (*InteractiveSession, error) {
	cmd := exec.Command(command, args...)

	// Start the command with a PTY
	ptmx, err := pty.Start(cmd)
	if err != nil {
		return nil, fmt.Errorf("failed to start command with PTY: %w", err)
	}

	// Set a reasonable terminal size
	_ = pty.Setsize(ptmx, &pty.Winsize{
		Rows: 40,
		Cols: 120,
	})

	session := &InteractiveSession{
		cmd:         cmd,
		ptmx:        ptmx,
		reader:      bufio.NewReader(ptmx),
		outputChan:  make(chan string, 1000),
		stopReading: make(chan bool),
	}

	// Start a single goroutine to handle all reading and terminal control sequences
	go session.handleIO()

	// Give it a moment to start
	time.Sleep(50 * time.Millisecond)

	return session, nil
}

// handleIO is the single goroutine that reads from PTY and responds to control sequences
func (s *InteractiveSession) handleIO() {
	buf := make([]byte, 1024)
	for {
		select {
		case <-s.stopReading:
			return
		default:
			n, err := s.ptmx.Read(buf)
			if n > 0 {
				data := string(buf[:n])
				s.output.WriteString(data)
				s.outputChan <- data

				// Check for Device Status Report (DSR) query: ESC[6n
				// Survey uses this to detect cursor position
				if strings.Contains(data, "\x1b[6n") {
					// Respond with cursor position report: ESC[{row};{col}R
					// Report cursor at row 1, column 1
					s.ptmx.Write([]byte("\x1b[1;1R"))
				}
			}
			if err != nil {
				if err != io.EOF {
					s.outputChan <- fmt.Sprintf("[IO Error: %v]", err)
				}
				return
			}
		}
	}
}

// ExpectPrompt waits for expected text to appear in the output
// Returns all output received up to and including the expected text
func (s *InteractiveSession) ExpectPrompt(expected string, timeout time.Duration) (string, error) {
	deadline := time.Now().Add(timeout)
	var collected strings.Builder

	for time.Now().Before(deadline) {
		select {
		case data := <-s.outputChan:
			collected.WriteString(data)

			// Check if we found the expected text
			if expected == "" || strings.Contains(collected.String(), expected) {
				return collected.String(), nil
			}

		case <-time.After(100 * time.Millisecond):
			// Check timeout
			if time.Now().After(deadline) {
				break
			}
		}
	}

	return collected.String(), fmt.Errorf("timeout waiting for %q, got: %s", expected, collected.String())
}

// SendLine sends input followed by a carriage return (for raw terminal mode)
func (s *InteractiveSession) SendLine(input string) error {
	// Survey uses raw terminal mode, which expects \r instead of \n
	_, err := s.ptmx.Write([]byte(input + "\r"))
	if err != nil {
		return fmt.Errorf("failed to send input: %w", err)
	}
	// Small delay to ensure input is processed
	time.Sleep(50 * time.Millisecond)
	return nil
}

// Send sends input without a newline
func (s *InteractiveSession) Send(input string) error {
	_, err := s.ptmx.Write([]byte(input))
	if err != nil {
		return fmt.Errorf("failed to send input: %w", err)
	}
	return nil
}

// SendEnter sends just a carriage return (accept default)
func (s *InteractiveSession) SendEnter() error {
	_, err := s.ptmx.Write([]byte("\r"))
	if err != nil {
		return fmt.Errorf("failed to send enter: %w", err)
	}
	// Small delay to ensure input is processed
	time.Sleep(50 * time.Millisecond)
	return nil
}

// Wait waits for the command to complete with a timeout
func (s *InteractiveSession) Wait(timeout time.Duration) error {
	done := make(chan error, 1)
	go func() {
		done <- s.cmd.Wait()
	}()

	select {
	case err := <-done:
		return err
	case <-time.After(timeout):
		return fmt.Errorf("command timed out after %v", timeout)
	}
}

// Close closes the PTY and kills the process if still running
func (s *InteractiveSession) Close() error {
	// Signal the reader goroutine to stop
	select {
	case s.stopReading <- true:
	default:
	}

	if s.cmd.Process != nil {
		_ = s.cmd.Process.Kill()
	}
	if s.ptmx != nil {
		return s.ptmx.Close()
	}
	return nil
}

// GetFullOutput returns all output captured so far
func (s *InteractiveSession) GetFullOutput() string {
	return s.output.String()
}

// GetFullOutput returns all output captured so far, and strip ANSI characters
func (s *InteractiveSession) GetFullCleanOutput() string {
	return stripansi.Strip(s.output.String())
}

// ReadAll reads all available output from the PTY with a timeout
// Uses idle timeout - stops when no new data arrives for 300ms
func (s *InteractiveSession) ReadAll(timeout time.Duration) string {
	deadline := time.Now().Add(timeout)
	lastReadTime := time.Now()
	idleTimeout := 300 * time.Millisecond // If no data for 300ms, consider done

	for time.Now().Before(deadline) {
		select {
		case <-s.outputChan:
			// Drain the channel
			lastReadTime = time.Now()

		case <-time.After(50 * time.Millisecond):
			// Check if we've been idle too long
			if time.Since(lastReadTime) > idleTimeout {
				return s.output.String()
			}
		}
	}

	return s.output.String()
}

// ExpectAndSend waits for a prompt and sends a response
func (s *InteractiveSession) ExpectAndSend(prompt, response string, timeout time.Duration) error {
	_, err := s.ExpectPrompt(prompt, timeout)
	if err != nil {
		return err
	}
	// Give the prompt a moment to fully render before sending
	time.Sleep(100 * time.Millisecond)
	return s.SendLine(response)
}

// DumpOutput logs all output received so far (useful for debugging)
func (s *InteractiveSession) DumpOutput() {
	output := s.output.String()
	fmt.Printf("=== RAW OUTPUT ===\n%q\n", output)
	fmt.Printf("=== CLEANED OUTPUT ===\n%s\n", stripansi.Strip(output))
	fmt.Printf("=== END OUTPUT ===\n")
}

// ContinuousRead starts reading output in the background and printing it
// This is useful for debugging to see what's happening in real-time
func (s *InteractiveSession) ContinuousRead(done chan bool) {
	go func() {
		buf := make([]byte, 1024)
		for {
			select {
			case <-done:
				return
			default:
				s.ptmx.SetReadDeadline(time.Now().Add(100 * time.Millisecond))
				n, err := s.ptmx.Read(buf)
				if n > 0 {
					data := string(buf[:n])
					s.output.WriteString(data)
					fmt.Printf("[PTY OUTPUT] %q\n", data)
				}
				if err != nil && err != io.EOF && !os.IsTimeout(err) {
					return
				}
			}
		}
	}()
}
