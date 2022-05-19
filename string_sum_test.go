package string_sum

import "testing"

func TestStringSum(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name       string
		args       args
		wantOutput string
		wantErr    bool
	}{
		{name: "empty", args: struct{ input string }{input: ""}, wantOutput: "", wantErr: true},
		{name: "simple", args: struct{ input string }{input: "3+5"}, wantOutput: "8", wantErr: false},
		{name: "simple", args: struct{ input string }{input: "-3+5"}, wantOutput: "2", wantErr: false},
		{name: "simple", args: struct{ input string }{input: "-3-5"}, wantOutput: "-8", wantErr: false},
		{name: "simple", args: struct{ input string }{input: "-3-5-8"}, wantOutput: "", wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOutput, err := StringSum(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("StringSum() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotOutput != tt.wantOutput {
				t.Errorf("StringSum() gotOutput = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}
