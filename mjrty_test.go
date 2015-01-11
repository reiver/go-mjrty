package main


import "testing"


func TestMjrty(t *testing.T) {

	tests := []struct{
		stream   []string
		expected   string
	}{
		{
			stream: []string{"apple"},
			expected: "apple",
		},
		{
			stream: []string{"banana"},
			expected: "banana",
		},
		{
			stream: []string{"cherry"},
			expected: "cherry",
		},


		{
			stream: []string{"x", "x"},
			expected: "x",
		},


		{
			stream: []string{"apple", "apple", "banana"},
			expected: "apple",
		},
		{
			stream: []string{"apple", "banana", "apple"},
			expected: "apple",
		},
		{
			stream: []string{"banana", "apple", "apple"},
			expected: "apple",
		},


		{
			stream: []string{"apple", "apple", "apple", "banana"},
			expected: "apple",
		},
		{
			stream: []string{"apple", "apple", "banana", "apple"},
			expected: "apple",
		},
		{
			stream: []string{"apple", "banana", "apple", "apple"},
			expected: "apple",
		},
		{
			stream: []string{"banana", "apple", "apple", "apple"},
			expected: "apple",
		},


		{
			stream: []string{"apple", "apple", "apple", "apple", "banana"},
			expected: "apple",
		},
		{
			stream: []string{"apple", "apple", "apple", "banana", "apple"},
			expected: "apple",
		},
		{
			stream: []string{"apple", "apple", "banana", "apple", "apple"},
			expected: "apple",
		},
		{
			stream: []string{"apple", "banana", "apple", "apple", "apple"},
			expected: "apple",
		},
		{
			stream: []string{"banana", "apple", "apple", "apple", "apple"},
			expected: "apple",
		},


		{
			stream: []string{"apple", "apple", "apple", "banana", "banana"},
			expected: "apple",
		},
		{
			stream: []string{"apple", "apple", "banana", "apple", "banana"},
			expected: "apple",
		},
		{
			stream: []string{"apple", "banana", "apple", "apple", "banana"},
			expected: "apple",
		},
		{
			stream: []string{"banana", "apple", "apple", "apple", "banana"},
			expected: "apple",
		},
		{
			stream: []string{"apple", "apple", "banana", "banana", "apple"},
			expected: "apple",
		},
		{
			stream: []string{"apple", "banana", "apple", "banana", "apple"},
			expected: "apple",
		},
		{
			stream: []string{"banana", "apple", "apple", "banana", "apple"},
			expected: "apple",
		},
		{
			stream: []string{"apple", "banana", "banana", "apple", "apple"},
			expected: "apple",
		},
		{
			stream: []string{"banana", "apple", "banana", "apple", "apple"},
			expected: "apple",
		},
		{
			stream: []string{"banana", "banana", "apple", "apple", "apple"},
			expected: "apple",
		},


		{
			stream: []string{"apple", "apple", "apple", "banana", "cherry"},
			expected: "apple",
		},
		{
			stream: []string{"apple", "apple", "banana", "apple", "cherry"},
			expected: "apple",
		},
		{
			stream: []string{"apple", "banana", "apple", "apple", "cherry"},
			expected: "apple",
		},
		{
			stream: []string{"banana", "apple", "apple", "apple", "cherry"},
			expected: "apple",
		},
		{
			stream: []string{"apple", "apple", "banana", "cherry", "apple"},
			expected: "apple",
		},
		{
			stream: []string{"apple", "banana", "apple", "cherry", "apple"},
			expected: "apple",
		},
		{
			stream: []string{"banana", "apple", "apple", "cherry", "apple"},
			expected: "apple",
		},
		{
			stream: []string{"apple", "banana", "cherry", "apple", "apple"},
			expected: "apple",
		},
		{
			stream: []string{"banana", "apple", "cherry", "apple", "apple"},
			expected: "apple",
		},
		{
			stream: []string{"banana", "cherry", "apple", "apple", "apple"},
			expected: "apple",
		},


		{
			stream: []string{"Apple", "Apple", "Apple", "Cherry", "Cherry", "Banana", "Banana", "Cherry", "Cherry", "Cherry", "Banana", "Cherry", "Cherry"},
			expected: "Cherry",
		},
	}



	// Run tests.
		for test_number, test := range tests {

			// Initialize.
				mjrty := New()

			// Insert data.
				for _, datum := range test.stream {

					mjrty.Insert(datum)
				}

			// Test what is expected.
				actual := mjrty.Majority()

				if test.expected != actual {
					t.Errorf("For test #%d expected %q but actually received %q", test_number, test.expected, actual)
				}
		}
}
