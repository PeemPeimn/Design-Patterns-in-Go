package abstract_factory

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAbstractFactory(t *testing.T) {
	tests := map[string]struct {
		osType        string
		expectedPaint string
		expectedCheck string
		expectedError bool
	}{
		"wrong type": {
			osType:        "unknown",
			expectedError: true,
		},
		"windows": {
			osType:        "Windows",
			expectedPaint: "WindowsButton",
			expectedCheck: "WindowsCheckBox",
			expectedError: false,
		},
		"mac": {
			osType:        "Mac",
			expectedPaint: "MacButton",
			expectedCheck: "MacCheckBox",
			expectedError: false,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			factory, err := NewFactory(test.osType)

			if test.expectedError {
				assert.NotNil(t, err)
				return
			}

			button := factory.CreateButton()
			checkBox := factory.CreateCheckBox()

			assert.Equal(t, test.expectedPaint, button.Paint())
			assert.Equal(t, test.expectedCheck, checkBox.Check())
		})
	}
}
