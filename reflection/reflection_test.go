package reflection

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"Struct with one string field",
			struct {
				Name string
			}{"Niels"},
			[]string{"Niels"},
		},
		{
			"Struct with two string fields",
			struct {
				Name string
				City string
			}{"Niels", "Berlin"},
			[]string{"Niels", "Berlin"},
		},
		{
			"Struct with non string field",
			struct {
				Name string
				Age  int
			}{"Niels", 41},
			[]string{"Niels"},
		},
		{
			"Nested Fields",
			Person{
				"Niels",
				Profile{41, "Berlin"},
			},
			[]string{"Niels", "Berlin"},
		},
		{
			"Pointers to things",
			&Person{
				"Niels",
				Profile{41, "Berlin"},
			},
			[]string{"Niels", "Berlin"},
		},
		{
			"Slices",
			[]Profile{
				{33, "London"},
				{41, "Berlin"},
			},
			[]string{"London", "Berlin"},
		},
		{
			"Arrays",
			[2]Profile{
				{33, "London"},
				{41, "Berlin"},
			},
			[]string{"London", "Berlin"},
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

	t.Run("with Maps", func(t *testing.T) {
		aMap := map[string]string{
			"Foo": "Bar",
			"Baz": "Boz",
		}

		var got []string
		walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Bar")
		assertContains(t, got, "Boz")
	})
}

func assertContains(t *testing.T, haystack []string, needle string) {
	t.Helper()
	contains := false

	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}
	if !contains {
		t.Errorf("expected %v to contain %q but it didn't", haystack, needle)
	}
}
