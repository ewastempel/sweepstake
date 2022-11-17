package main

import (
	"testing"
)

func TestAssignCountriesToParticipants(t *testing.T) {
	t.Run("Testing the length of the map and the length of its elements", func(t *testing.T) {
		countries := map[int]string{
			0: "Canada",
			1: "England",
			2: "USA",
			3: "Argentina",
			4: "Uruguay",
			5: "Wales",
			6: "Spain",
			7: "Netherlands",
			8: "Morocco",
			9: "Japan",
		}
		participants := map[string]int{
			"Siobhan": 2,
			"Aidan":   4,
			"Paul":    1,
			"Ruby":    3,
		}
		want := map[string][]string{
			"Siobhan": {"USA", "Argentina"},
			"Aidan":   {"Spain", "Wales", "Netherlands", "Canada"},
			"Paul":    {"England"},
			"Ruby":    {"Uruguay", "Morocco", "Japan"},
		}
		got := AssignCountriesToParticipants(countries, participants)
		if len(got) != len(want) {
			t.Errorf("got len: %q want len: %q", len(got), len(want))
		}
		if len(got["Siobhan"]) != len(want["Siobhan"]) {
			t.Errorf("got: %q want: %q", len(got["Siobhan"]), len(want["Siobhan"]))
		}
		if len(got["Aidan"]) != len(want["Aidan"]) {
			t.Errorf("got: %q want: %q", len(got["Aidan"]), len(want["Aidan"]))
		}
		if len(got["Paul"]) != len(want["Paul"]) {
			t.Errorf("got: %q want: %q", len(got["Paul"]), len(want["Paul"]))
		}
		if len(got["Ruby"]) != len(want["Ruby"]) {
			t.Errorf("got: %q want: %q", len(got["Ruby"]), len(want["Ruby"]))
		}
		if len(got["Siobhan"]) != participants["Siobhan"] {
			t.Errorf("got: %q want: %q", len(got["Siobhan"]), participants["Siobhan"])
		}
		if len(got["Aidan"]) != participants["Aidan"] {
			t.Errorf("got: %q want: %q", len(got["Aidan"]), participants["Aidan"])
		}
		if len(got["Paul"]) != participants["Paul"] {
			t.Errorf("got: %q want: %q", len(got["Paul"]), participants["Paul"])
		}
		if len(got["Ruby"]) != participants["Ruby"] {
			t.Errorf("got: %q want: %q", len(got["Ruby"]), participants["Ruby"])
		}
	})
}
