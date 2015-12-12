package transform

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReplaceSingleRequires(t *testing.T) {

	input := "moment = require('moment');\n" +
		"Vue = require('vue');\n" +
		"\n" +
		"export default {\n" +
		"}\n"

	expected := "import moment from 'moment';\n" +
		"import Vue from 'vue';\n" +
		"\n" +
		"export default {\n" +
		"}\n\n" // XXX extra newline

	assert.Equal(t, expected, ReplaceRequiresWithEs6Imports(input))
}

func TestReplaceInlineRequires(t *testing.T) {

	input := "export default {\n" +
		"    components: {\n" +
		"        NavBar: require('./components/NavBar.vue'),\n" +
		"    }\n" +
		"}\n"

	expected := "import NavBar from './components/NavBar.vue';\n" +
		"export default {\n" +
		"    components: {\n" +
		"        NavBar,\n" +
		"    }\n" +
		"}\n\n" // XXX extra newline

	assert.Equal(t, expected, ReplaceRequiresWithEs6Imports(input))
}
