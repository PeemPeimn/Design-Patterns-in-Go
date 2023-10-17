package factory_method

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFactoryMethod(t *testing.T) {
	tests := map[string]struct {
		dialogType       string
		expectedRendered string
		expectedError    bool
	}{
		"wrong type": {
			dialogType:    "unknown",
			expectedError: true,
		},
		"windows": {
			dialogType:       "Windows",
			expectedRendered: "Windows button",
			expectedError:    false,
		},
		"web": {
			dialogType:       "Web",
			expectedRendered: "HTML button",
			expectedError:    false,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			dialog, err := getDialog(test.dialogType)

			if test.expectedError {
				assert.NotNil(t, err)
				return
			}

			assert.Equal(t, test.expectedRendered, dialog.Render())
		})
	}
}
