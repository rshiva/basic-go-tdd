package main
import (
				"testing" 
				"reflect"
			)

func TestWalk(t *testing.T) {

	type Profile struct {
    Age  int
    City string
}

	type Person struct {
    Name    string
    Profile Profile
}

	cases := []struct{
			Name string
			Input interface{}
			ExpectedCalls []string
	} {
		{
			"Pointers to things",
			&Person{
					"Chris",
					Profile{33, "London"},
			},
			[]string{"Chris", "London"},
	},
	}

	for _, test := range cases {
			t.Run(test.Name, func(t *testing.T) {
					var got []string
					walk(test.Input, func(input string) {
							got = append(got, input)
					})

					if !reflect.DeepEqual(got, test.ExpectedCalls) {
							t.Errorf("got %v, want %v", got, test.ExpectedCalls)
					}
			})
	}
}