package parser

import (
	"testing"
	"strings"
	"github.com/stretchr/testify/assert"
)

func TestClass(t *testing.T) {
	input := ".class public Lcom/checker/StatusChecker;"


	javaFile := &JavaFile{}
	(&ClassParser{}).Parse(javaFile, strings.Fields(input))

	expectedOutput := "public class com.checker.StatusChecker {"

	output := strings.Join(javaFile.First(), " ")

	assert.Equal(t, expectedOutput, output)
}
