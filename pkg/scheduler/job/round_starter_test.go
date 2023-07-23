package job

import "testing"

func Test_isServiceScorableInRound(t *testing.T) {
	type args struct {
		roundNumber int
		roundFreq   int
		roundDelay  int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{name: "Basic", args: args{roundNumber: 1, roundFreq: 1, roundDelay: 0}, want: true},
		{name: "Basic", args: args{roundNumber: 1, roundFreq: 1, roundDelay: 1}, want: false},
		{name: "Basic", args: args{roundNumber: 1, roundFreq: 1, roundDelay: 2}, want: false},
		{name: "3rd round, 3 F, 2 D", args: args{roundNumber: 3, roundFreq: 3, roundDelay: 2}, want: false},
		{name: "5th round, 3 F, 1 D", args: args{roundNumber: 5, roundFreq: 3, roundDelay: 1}, want: false},
		{name: "10th round, 4 F, 2 D", args: args{roundNumber: 10, roundFreq: 4, roundDelay: 2}, want: true},
		{name: "20th round, 2 F, 2 D", args: args{roundNumber: 20, roundFreq: 2, roundDelay: 2}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isServiceScorableInRound(tt.args.roundNumber, tt.args.roundFreq, tt.args.roundDelay); got != tt.want {
				t.Errorf("isServiceScorableInRound() = %v, want %v", got, tt.want)
			}
		})
	}
}
