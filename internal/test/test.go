package test

import (
	"strings"
	"testing"
)

func AssertThatInformed(t *testing.T, outs strings.Builder, expect string) {
	t.Helper()
	assertStringBuilderContains(t, outs, "info", expect)
}

func AssertThatAlerted(t *testing.T, errs strings.Builder, expect string) {
	t.Helper()
	assertStringBuilderContains(t, errs, "alert", expect)
}

func assertStringBuilderContains(t *testing.T, sb strings.Builder, name, expect string) {
	t.Helper()
	if !strings.Contains(sb.String(), expect) {
		t.Errorf("did not find:\n\t%s\nin %s stream:\n\t%s\n", indent(expect), name, indent(sb.String()))
	} else if len(sb.String()) > 0 && len(expect) == 0 {
		t.Errorf("expected %s stream to be empty but got:\n\t%s", name, indent(sb.String()))
	}
}

func indent(str string) string {
	return strings.Replace(str, "\n", "\n\t", -1)
}
