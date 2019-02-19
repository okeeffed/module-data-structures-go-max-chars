package maxchars

type testCase struct {
	description string
	input       string
	expected    string
}

var testCases = []testCase{
	{
		description: "should have c as most frequest char in abcccccf",
		input:       "abcccccf",
		expected:    "c",
	},
	{
		description: "should have d as most frequest char in abcccdddddf",
		input:       "abcccdddddf",
		expected:    "d",
	},
	{
		description: "should have d as most frequest char in abcccdddddf",
		input:       "abcccdddddf!!!!!!!!!!",
		expected:    "d",
	},
}
