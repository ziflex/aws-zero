package awszero_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/aws/smithy-go/logging"
	"github.com/rs/zerolog"
	"github.com/ziflex/aws-zero"
)

func setup() (logging.Logger, *bytes.Buffer) {
	out := &bytes.Buffer{}
	logger := zerolog.New(out)

	return awszero.New(logger), out
}

type UseCase struct {
	Description string
	Level       string
	Format      string
	Arguments   []interface{}
	Expected    string
}

func Test_Logf(t *testing.T) {
	cases := []UseCase{
		{
			Description: "Debug",
			Level:       string(logging.Debug),
			Format:      "Foo",
			Arguments:   nil,
			Expected:    `{"level":"debug","message":"Foo"}`,
		},
		{
			Description: "Debug with format",
			Level:       string(logging.Debug),
			Format:      "Foo %s",
			Arguments: []interface{}{
				"Bar",
			},
			Expected: `{"level":"debug","message":"Foo Bar"}`,
		},
		{
			Description: "Warn",
			Level:       string(logging.Warn),
			Format:      "Foo",
			Arguments:   nil,
			Expected:    `{"level":"warn","message":"Foo"}`,
		},
		{
			Description: "Warn with format",
			Level:       string(logging.Warn),
			Format:      "Foo %s",
			Arguments: []interface{}{
				"Bar",
			},
			Expected: `{"level":"warn","message":"Foo Bar"}`,
		},
		{
			Description: "Unknown level. Fallback to trace.",
			Level:       "",
			Format:      "Foo",
			Arguments:   nil,
			Expected:    `{"level":"trace","message":"Foo"}`,
		},
		{
			Description: "Unknown level. Fallback to trace warn with format",
			Level:       "",
			Format:      "Foo %s",
			Arguments: []interface{}{
				"Bar",
			},
			Expected: `{"level":"trace","message":"Foo Bar"}`,
		},
	}

	for _, uc := range cases {
		c := uc
		t.Run(c.Description, func(t2 *testing.T) {
			log, out := setup()

			out.Reset()

			log.Logf(logging.Classification(uc.Level), uc.Format, uc.Arguments...)

			actual := strings.Trim(out.String(), "\n")

			if actual != uc.Expected {
				t2.Fatalf("%s expected to equal to %s", actual, uc.Expected)
			}
		})

	}
}
