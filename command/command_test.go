package command

import (
	"bytes"
	"io"
	"testing"

	ui "github.com/chroju/alfrednize/UI"
)

var (
	jsonMultipleItems = "{\"items\":[{\"uid\":\"foo\",\"title\":\"foo\",\"subtitle\":\"\",\"arg\":\"foo\",\"match\":\"foo\",\"autocomplete\":\"foo\"},{\"uid\":\"bar\",\"title\":\"bar\",\"subtitle\":\"\",\"arg\":\"bar\",\"match\":\"bar\",\"autocomplete\":\"bar\"}]}\n"
)

type mockUI struct {
	in     []string
	out    io.Writer
	errOut io.Writer
}

func newMockUI(in []string, out, errOut io.Writer) ui.UI {
	return &mockUI{
		in:     in,
		out:    out,
		errOut: errOut,
	}
}

func (mockUi *mockUI) In() []string {
	return mockUi.in
}

func (mockUi *mockUI) Out() io.Writer {
	return mockUi.out
}

func (mockUi *mockUI) ErrOut() io.Writer {
	return mockUi.errOut
}

func Test_execCommand(t *testing.T) {
	tests := []struct {
		name          string
		args          []string
		stdin         []string
		want          int
		wantOutWriter string
		wantErrWriter string
	}{
		{
			name:          "help",
			args:          []string{"--help"},
			want:          0,
			wantOutWriter: helpMessage,
		},
		{
			name:          "help short hand",
			args:          []string{"-h"},
			want:          0,
			wantOutWriter: helpMessage,
		},
		{
			name:          "version",
			args:          []string{"--version"},
			want:          0,
			wantOutWriter: versionCommand(),
		},
		{
			name:          "version short hand",
			args:          []string{"-v"},
			want:          0,
			wantOutWriter: versionCommand(),
		},
		{
			name:          "hyphen without stdin",
			args:          []string{},
			want:          1,
			wantErrWriter: helpCommand(),
		},
		{
			name:          "hyphen with stdin",
			args:          []string{},
			stdin:         []string{"foo", "bar"},
			want:          0,
			wantOutWriter: jsonMultipleItems,
		},
		{
			name:          "no args without stdin",
			args:          []string{},
			want:          1,
			wantErrWriter: helpCommand(),
		},
		{
			name:          "no args with stdin",
			args:          []string{},
			want:          0,
			stdin:         []string{"foo", "bar"},
			wantOutWriter: jsonMultipleItems,
		},
		{
			name:          "unexpected args without stdin",
			args:          []string{"test"},
			want:          1,
			wantErrWriter: helpCommand(),
		},
		{
			name:          "unexpected args with stdin",
			args:          []string{"test"},
			want:          1,
			stdin:         []string{"foo", "bar"},
			wantErrWriter: helpCommand(),
		},
		{
			name:          "multiple args",
			args:          []string{"test", "test"},
			want:          1,
			wantErrWriter: helpCommand(),
		},
		{
			name:          "empty stdin",
			want:          1,
			stdin:         []string{},
			wantErrWriter: helpCommand(),
		},
	}
	for _, tt := range tests {
		outWriter := &bytes.Buffer{}
		errWriter := &bytes.Buffer{}
		ui := newMockUI(tt.stdin, outWriter, errWriter)
		t.Run(tt.name, func(t *testing.T) {
			if got := ExecCommand(ui, tt.args); got != tt.want {
				t.Errorf("execCommand() = %v, want %v", got, tt.want)
			}
			if gotOutWriter := outWriter.String(); gotOutWriter != tt.wantOutWriter {
				t.Errorf("get() = %v, want %v", gotOutWriter, tt.wantOutWriter)
			}
			if gotErrWriter := errWriter.String(); gotErrWriter != tt.wantErrWriter {
				t.Errorf("get() = %v, want %v", gotErrWriter, tt.wantErrWriter)
			}
		})
	}
}
