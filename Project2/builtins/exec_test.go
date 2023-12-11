package builtins_test

import (
	"errors"
	"testing"

	"github.com/wdw0058/CSCE4600/Project2/builtins"
)

func TestExec(t *testing.T) {
	type args struct {
		args []string
	}

	tests := []struct {
		name         string
		args         args
		unsetHomedir bool
		wantDir      string
		wantErr      error
	}{
		{
			name: "could not start program",
			args: args{
				args: []string{"cheesetestprogram"}, //does not work on windows, why its tested here
			},
			wantErr: builtins.ErrCouldNotStartProgram,
		},
		{
			name: "too few arguments",
			args: args{
				args: []string{}, //does not work on windows, why its tested here
			},
			wantErr: builtins.ErrTooFewArguments,
		},
	}
	for _, tt := range tests {
		exit := make(chan struct{}, 2) // test exit pass

		t.Run(tt.name, func(t *testing.T) {
			// testing
			if err := builtins.Exec(exit, tt.args.args...); tt.wantErr != nil {
				if !errors.Is(err, tt.wantErr) {
					t.Fatalf("Exec() error = %v, wantErr %v", err, tt.wantErr)
				}
				return
			} else if err != nil {
				t.Fatalf("Exec() unexpected error: %v", err)
			}
		})
	}
}
