package builtins_test

import (
	"errors"
	"testing"

	"github.com/wdw0058/CSCE4600/Project2/builtins"
)

func TestHistory(t *testing.T) {
	type args struct {
		args []string
	}

	tests := []struct {
		name             string
		args             args
		unsetHistoryFile bool
		function         string
		wantDir          string
		wantErr          error
	}{
		{
			name: "could not open history - write",
			args: args{
				args: []string{"cheesetestprogram"}, //does not work on windows, why its tested here
			},
			function: "write",
			wantErr:  builtins.ErrNoOpenHistory,
		},
		{
			name: "could not open history - clear",
			args: args{
				args: []string{"cheesetestprogram"}, //does not work on windows, why its tested here
			},
			function: "clear",
			wantErr:  builtins.ErrNoOpenHistory,
		},
		{
			name: "could not open history - print",
			args: args{
				args: []string{"cheesetestprogram"}, //does not work on windows, why its tested here
			},
			function: "print",
			wantErr:  builtins.ErrNoOpenHistory,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			oldVal := builtins.HistoryPath
			t.Cleanup(func() {
				builtins.HistoryPath = oldVal
			})
			builtins.HistoryPath = ""

			switch tt.function {
			case "write":
				// testing
				if err := builtins.WriteHistory(tt.args.args[0]); tt.wantErr != nil {
					if !errors.Is(err, tt.wantErr) {
						t.Fatalf("Exec() error = %v, wantErr %v", err, tt.wantErr)
					}
					return
				} else if err != nil {
					t.Fatalf("Exec() unexpected error: %v", err)
				}
			case "clear":
				// testing
				if err := builtins.ClearHistory(); tt.wantErr != nil {
					if !errors.Is(err, tt.wantErr) {
						t.Fatalf("Exec() error = %v, wantErr %v", err, tt.wantErr)
					}
					return
				} else if err != nil {
					t.Fatalf("Exec() unexpected error: %v", err)
				}
			case "print":
				// testing
				if err := builtins.PrintHistory(); tt.wantErr != nil {
					if !errors.Is(err, tt.wantErr) {
						t.Fatalf("Exec() error = %v, wantErr %v", err, tt.wantErr)
					}
					return
				} else if err != nil {
					t.Fatalf("Exec() unexpected error: %v", err)
				}
			}
		})
	}
}
