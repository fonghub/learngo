package m1


import (
    "testing"
    "regexp"
)

// TestHello3Name calls greetings.Hello3 with a name, checking
// for a valid return value.
func TestHello3Name(t *testing.T) {
    name := "Gladys"
    want := regexp.MustCompile(`\b`+name+`\b`)
    msg, err := Hello3("Gladys")
    if !want.MatchString(msg) || err != nil {
        t.Fatalf(`Hello3("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)
    }
}

// TestHello3Empty calls greetings.Hello3 with an empty string,
// checking for an error.
func TestHello3Empty(t *testing.T) {
    msg, err := Hello3("")
    if msg != "" || err == nil {
        t.Fatalf(`Hello3("") = %q, %v, want "", error`, msg, err)
    }
}
