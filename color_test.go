package color

import "testing"

// Clearly not a real test - just a visual reference
func TestOutput(t *testing.T) {
	println(Red + "Hello" + Reset)
	println(Green + "Hello" + Reset)
	println(Yellow + "Hello" + Reset)
	println(Blue + "Hello" + Reset)
	println(Purple + "Hello" + Reset)
	println(Cyan + "Hello" + Reset)
	println(Gray + "Hello" + Reset)
	println(White + "Hello" + Reset)
	println(Red + "H" + Green + "el" + Purple + "lo" + Reset)
}

func TestColorize(t *testing.T) {
	tests := []struct {
		color string
		want  string
	}{
		{
			color: Red,
			want:  "\033[31mtest\033[0m",
		},
		{
			color: Green,
			want:  "\033[32mtest\033[0m",
		},
		{
			color: Yellow,
			want:  "\033[33mtest\033[0m",
		},
		{
			color: Blue,
			want:  "\033[34mtest\033[0m",
		},
		{
			color: Purple,
			want:  "\033[35mtest\033[0m",
		},
		{
			color: Cyan,
			want:  "\033[36mtest\033[0m",
		},
		{
			color: Gray,
			want:  "\033[37mtest\033[0m",
		},
		{
			color: White,
			want:  "\033[97mtest\033[0m",
		},
	}

	for _, tc := range tests {
		t.Run(tc.color, func(t *testing.T) {
			output := Colorize(tc.color, "test")
			if output != tc.want {
				t.Errorf("Expected %s, got %s", tc.want, output)
			}
		})
	}
}

func TestIn(t *testing.T) {
	tests := []struct {
		fn    func(string) string
		color string
		want  string
	}{
		{
			fn:    InRed,
			color: Red,
			want:  "\033[31mtest\033[0m",
		},
		{
			fn:    InGreen,
			color: Green,
			want:  "\033[32mtest\033[0m",
		},
		{
			fn:    InYellow,
			color: Yellow,
			want:  "\033[33mtest\033[0m",
		},
		{
			fn:    InBlue,
			color: Blue,
			want:  "\033[34mtest\033[0m",
		},
		{
			fn:    InPurple,
			color: Purple,
			want:  "\033[35mtest\033[0m",
		},
		{
			fn:    InCyan,
			color: Cyan,
			want:  "\033[36mtest\033[0m",
		},
		{
			fn:    InGray,
			color: Gray,
			want:  "\033[37mtest\033[0m",
		},
		{
			fn:    InWhite,
			color: White,
			want:  "\033[97mtest\033[0m",
		},
	}

	for _, tc := range tests {
		t.Run(tc.color, func(t *testing.T) {
			output := tc.fn("test")
			if output != tc.want {
				t.Errorf("Expected %s, got %s", tc.want, output)
			}
		})
	}
}
