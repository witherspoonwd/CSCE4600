package builtins_test

import (
	"errors"
	"testing"

	"github.com/wdw0058/CSCE4600/Project2/builtins"
)

func TestListDirectory(t *testing.T) {
	type args struct {
		args []string
	}

	tests := []struct {
		name             string
		args             args
		unsetOsDirectory bool
		wantDir          string
		wantErr          error
	}{
		{
			name: "couldnt open directory - OS fail",
			args: args{
				args: []string{}, //does not work on windows, why its tested here
			},
			wantErr:          builtins.ErrCouldNotOpenDirectory,
			unsetOsDirectory: true,
		},
		{
			name: "couldnt open directory - bad pass",
			args: args{
				args: []string{"fakefakefakefaaaakeFAKEFAKE"}, //does not work on windows, why its tested here
			},
			wantErr:          builtins.ErrCouldNotOpenDirectory,
			unsetOsDirectory: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.unsetOsDirectory {
				oldVal := builtins.OsDirectory
				t.Cleanup(func() {
					builtins.OsDirectory = oldVal
				})
				builtins.OsDirectory = ""
			}

			// testing
			if err := builtins.ListDirectory(tt.args.args...); tt.wantErr != nil {
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
