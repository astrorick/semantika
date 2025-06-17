package semantika_test

import (
	"testing"

	"github.com/astrorick/semantika"
)

// TestNew calls New() with different values, and checks for valid returns.
func TestNew(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name            string
		versionString   string
		expectedVersion *semantika.Version
		errorIsExpected bool
	}{
		// successfull tests
		{
			name:          "TestNew_0.0.0",
			versionString: "0.0.0",
			expectedVersion: &semantika.Version{
				Major: 0,
				Minor: 0,
				Patch: 0,
			},
			errorIsExpected: false,
		},
		{
			name:          "TestNew_1.0.0",
			versionString: "1.0.0",
			expectedVersion: &semantika.Version{
				Major: 1,
				Minor: 0,
				Patch: 0,
			},
			errorIsExpected: false,
		},
		{
			name:          "TestNew_0.2.0",
			versionString: "0.2.0",
			expectedVersion: &semantika.Version{
				Major: 0,
				Minor: 2,
				Patch: 0,
			},
			errorIsExpected: false,
		},
		{
			name:          "TestNew_0.0.3",
			versionString: "0.0.3",
			expectedVersion: &semantika.Version{
				Major: 0,
				Minor: 0,
				Patch: 3,
			},
			errorIsExpected: false,
		},
		{
			name:          "TestNew_1.2.3",
			versionString: "1.2.3",
			expectedVersion: &semantika.Version{
				Major: 1,
				Minor: 2,
				Patch: 3,
			},
			errorIsExpected: false,
		},

		// characters in input string
		{
			name:            "TestNew_a.2.3",
			versionString:   "a.2.3",
			expectedVersion: nil,
			errorIsExpected: true,
		},
		{
			name:            "TestNew_1.b.3",
			versionString:   "1.b.3",
			expectedVersion: nil,
			errorIsExpected: true,
		},
		{
			name:            "TestNew_1.2.c",
			versionString:   "1.2.c",
			expectedVersion: nil,
			errorIsExpected: true,
		},
		{
			name:            "TestNew_a.b.c",
			versionString:   "a.b.c",
			expectedVersion: nil,
			errorIsExpected: true,
		},

		// negative numbers
		{
			name:            "TestNew_-1.0.0",
			versionString:   "-1.0.0",
			expectedVersion: nil,
			errorIsExpected: true,
		},
		{
			name:            "TestNew_0.-1.0",
			versionString:   "0.-1.0",
			expectedVersion: nil,
			errorIsExpected: true,
		},
		{
			name:            "TestNew_0.0.-1",
			versionString:   "0.0.-1",
			expectedVersion: nil,
			errorIsExpected: true,
		},

		// invalid strings
		{
			name:            "TestNew_...",
			versionString:   "...",
			expectedVersion: nil,
			errorIsExpected: true,
		},
		{
			name:            "TestNew_1a.2b.3c",
			versionString:   "1a.2b.3c",
			expectedVersion: nil,
			errorIsExpected: true,
		},
		{
			name:            "TestNew_hello",
			versionString:   "hello",
			expectedVersion: nil,
			errorIsExpected: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			res, err := semantika.New(test.versionString)

			// test for presence of error
			if err != nil {
				if !test.errorIsExpected {
					t.Errorf("expected: %v, result: %v", test.errorIsExpected, err != nil)
				}
				return
			}

			// test for version string consistency
			if *res != *test.expectedVersion {
				t.Errorf("expected: %v, result: %v", test.expectedVersion, res)
			}
		})
	}
}

// TestString calls Version.String() with different values, and checks for valid returns.
func TestString(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		version        *semantika.Version
		expectedString string
	}{
		{
			name: "TestString_0.0.0",
			version: &semantika.Version{
				Major: 0,
				Minor: 0,
				Patch: 0,
			},
			expectedString: "0.0.0",
		},
		{
			name: "TestString_1.0.0",
			version: &semantika.Version{
				Major: 1,
				Minor: 0,
				Patch: 0,
			},
			expectedString: "1.0.0",
		},
		{
			name: "TestString_0.2.0",
			version: &semantika.Version{
				Major: 0,
				Minor: 2,
				Patch: 0,
			},
			expectedString: "0.2.0",
		},
		{
			name: "TestString_0.0.3",
			version: &semantika.Version{
				Major: 0,
				Minor: 0,
				Patch: 3,
			},
			expectedString: "0.0.3",
		},
		{
			name: "TestString_1.2.3",
			version: &semantika.Version{
				Major: 1,
				Minor: 2,
				Patch: 3,
			},
			expectedString: "1.2.3",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if res := test.version.String(); res != test.expectedString {
				t.Errorf("expected: %v, result: %v", test.expectedString, res)
			}
		})
	}
}

// TestCompare calls Version.Compare() with different reference and comparison versions, and checks for valid returns.
func TestCompare(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name              string
		referenceVersion  *semantika.Version
		comparisonVersion *semantika.Version
		expectedResult    int8
	}{
		// tests with referenceVersion < comparisonVersion
		{
			name: "TestCompare_1.1.1_2.2.2",
			referenceVersion: &semantika.Version{
				Major: 1,
				Minor: 1,
				Patch: 1,
			},
			comparisonVersion: &semantika.Version{
				Major: 2,
				Minor: 2,
				Patch: 2,
			},
			expectedResult: -1,
		},
		{
			name: "TestCompare_1.1.1_2.1.1",
			referenceVersion: &semantika.Version{
				Major: 1,
				Minor: 1,
				Patch: 1,
			},
			comparisonVersion: &semantika.Version{
				Major: 2,
				Minor: 1,
				Patch: 1,
			},
			expectedResult: -1,
		},
		{
			name: "TestCompare_1.1.1_1.2.1",
			referenceVersion: &semantika.Version{
				Major: 1,
				Minor: 1,
				Patch: 1,
			},
			comparisonVersion: &semantika.Version{
				Major: 1,
				Minor: 2,
				Patch: 1,
			},
			expectedResult: -1,
		},
		{
			name: "TestCompare_1.1.1_1.1.2",
			referenceVersion: &semantika.Version{
				Major: 1,
				Minor: 1,
				Patch: 1,
			},
			comparisonVersion: &semantika.Version{
				Major: 1,
				Minor: 1,
				Patch: 2,
			},
			expectedResult: -1,
		},

		// tests with referenceVersion = comparisonVersion
		{
			name: "TestCompare_1.2.3",
			referenceVersion: &semantika.Version{
				Major: 1,
				Minor: 2,
				Patch: 3,
			},
			comparisonVersion: &semantika.Version{
				Major: 1,
				Minor: 2,
				Patch: 3,
			},
			expectedResult: 0,
		},
		{
			name: "TestCompare_1.1.1",
			referenceVersion: &semantika.Version{
				Major: 1,
				Minor: 1,
				Patch: 1,
			},
			comparisonVersion: &semantika.Version{
				Major: 1,
				Minor: 1,
				Patch: 1,
			},
			expectedResult: 0,
		},
		{
			name: "TestCompare_1.11.111",
			referenceVersion: &semantika.Version{
				Major: 1,
				Minor: 11,
				Patch: 111,
			},
			comparisonVersion: &semantika.Version{
				Major: 1,
				Minor: 11,
				Patch: 111,
			},
			expectedResult: 0,
		},

		// tests with referenceVersion > comparisonVersion
		{
			name: "TestCompare_2.2.2_1.1.1",
			referenceVersion: &semantika.Version{
				Major: 2,
				Minor: 2,
				Patch: 2,
			},
			comparisonVersion: &semantika.Version{
				Major: 1,
				Minor: 1,
				Patch: 1,
			},
			expectedResult: 1,
		},
		{
			name: "TestCompare_2.2.2_1.2.2",
			referenceVersion: &semantika.Version{
				Major: 2,
				Minor: 2,
				Patch: 2,
			},
			comparisonVersion: &semantika.Version{
				Major: 1,
				Minor: 2,
				Patch: 2,
			},
			expectedResult: 1,
		},
		{
			name: "TestCompare_2.2.2_2.1.2",
			referenceVersion: &semantika.Version{
				Major: 2,
				Minor: 2,
				Patch: 2,
			},
			comparisonVersion: &semantika.Version{
				Major: 2,
				Minor: 1,
				Patch: 2,
			},
			expectedResult: 1,
		},
		{
			name: "TestCompare_2.2.2_2.2.1",
			referenceVersion: &semantika.Version{
				Major: 2,
				Minor: 2,
				Patch: 2,
			},
			comparisonVersion: &semantika.Version{
				Major: 2,
				Minor: 2,
				Patch: 1,
			},
			expectedResult: 1,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if res := test.referenceVersion.Compare(test.comparisonVersion); res != test.expectedResult {
				t.Errorf("expected: %v, result: %v", test.expectedResult, res)
			}
		})
	}
}

// TestOlderThan calls Version.OlderThan() with different reference and comparison versions, and checks for valid returns.
func TestOlderThan(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name              string
		referenceVersion  *semantika.Version
		comparisonVersion *semantika.Version
		expectedResult    bool
	}{
		// tests with referenceVersion < comparisonVersion
		{
			name: "TestOlderThan_1.5.5_2.3.4",
			referenceVersion: &semantika.Version{
				Major: 1,
				Minor: 5,
				Patch: 5,
			},
			comparisonVersion: &semantika.Version{
				Major: 2,
				Minor: 3,
				Patch: 4,
			},
			expectedResult: true,
		},

		// tests with referenceVersion = comparisonVersion
		{
			name: "TestOlderThan_8.9.4_8.9.4",
			referenceVersion: &semantika.Version{
				Major: 8,
				Minor: 9,
				Patch: 4,
			},
			comparisonVersion: &semantika.Version{
				Major: 8,
				Minor: 9,
				Patch: 4,
			},
			expectedResult: false,
		},

		// tests with referenceVersion > comparisonVersion
		{
			name: "TestOlderThan_3.3.4_3.2.3",
			referenceVersion: &semantika.Version{
				Major: 3,
				Minor: 3,
				Patch: 4,
			},
			comparisonVersion: &semantika.Version{
				Major: 3,
				Minor: 2,
				Patch: 3,
			},
			expectedResult: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if res := test.referenceVersion.OlderThan(test.comparisonVersion); res != test.expectedResult {
				t.Errorf("expected: %v, result: %v", test.expectedResult, res)
			}
		})
	}
}

// TestOlderThanOrEquals calls Version.OlderThanOrEquals() with different reference and comparison versions, and checks for valid returns.
func TestOlderThanOrEquals(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name              string
		referenceVersion  *semantika.Version
		comparisonVersion *semantika.Version
		expectedResult    bool
	}{
		// tests with referenceVersion < comparisonVersion
		{
			name: "TestOlderThan_1.5.5_2.3.4",
			referenceVersion: &semantika.Version{
				Major: 1,
				Minor: 5,
				Patch: 5,
			},
			comparisonVersion: &semantika.Version{
				Major: 2,
				Minor: 3,
				Patch: 4,
			},
			expectedResult: true,
		},

		// tests with referenceVersion = comparisonVersion
		{
			name: "TestOlderThan_8.9.4_8.9.4",
			referenceVersion: &semantika.Version{
				Major: 8,
				Minor: 9,
				Patch: 4,
			},
			comparisonVersion: &semantika.Version{
				Major: 8,
				Minor: 9,
				Patch: 4,
			},
			expectedResult: true,
		},

		// tests with referenceVersion > comparisonVersion
		{
			name: "TestOlderThan_3.3.4_3.2.3",
			referenceVersion: &semantika.Version{
				Major: 3,
				Minor: 3,
				Patch: 4,
			},
			comparisonVersion: &semantika.Version{
				Major: 3,
				Minor: 2,
				Patch: 3,
			},
			expectedResult: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if res := test.referenceVersion.OlderThanOrEquals(test.comparisonVersion); res != test.expectedResult {
				t.Errorf("expected: %v, result: %v", test.expectedResult, res)
			}
		})
	}
}

// TestEquals calls Version.Equals() with different reference and comparison versions, and checks for valid returns.
func TestEquals(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name              string
		referenceVersion  *semantika.Version
		comparisonVersion *semantika.Version
		expectedResult    bool
	}{
		// tests with referenceVersion < comparisonVersion
		{
			name: "TestOlderThan_1.5.5_2.3.4",
			referenceVersion: &semantika.Version{
				Major: 1,
				Minor: 5,
				Patch: 5,
			},
			comparisonVersion: &semantika.Version{
				Major: 2,
				Minor: 3,
				Patch: 4,
			},
			expectedResult: false,
		},

		// tests with referenceVersion = comparisonVersion
		{
			name: "TestOlderThan_8.9.4_8.9.4",
			referenceVersion: &semantika.Version{
				Major: 8,
				Minor: 9,
				Patch: 4,
			},
			comparisonVersion: &semantika.Version{
				Major: 8,
				Minor: 9,
				Patch: 4,
			},
			expectedResult: true,
		},

		// tests with referenceVersion > comparisonVersion
		{
			name: "TestOlderThan_3.3.4_3.2.3",
			referenceVersion: &semantika.Version{
				Major: 3,
				Minor: 3,
				Patch: 4,
			},
			comparisonVersion: &semantika.Version{
				Major: 3,
				Minor: 2,
				Patch: 3,
			},
			expectedResult: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if res := test.referenceVersion.Equals(test.comparisonVersion); res != test.expectedResult {
				t.Errorf("expected: %v, result: %v", test.expectedResult, res)
			}
		})
	}
}

// TestNewerThanOrEquals calls Version.NewerThanOrEquals() with different reference and comparison versions, and checks for valid returns.
func TestNewerThanOrEquals(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name              string
		referenceVersion  *semantika.Version
		comparisonVersion *semantika.Version
		expectedResult    bool
	}{
		// tests with referenceVersion < comparisonVersion
		{
			name: "TestOlderThan_1.5.5_2.3.4",
			referenceVersion: &semantika.Version{
				Major: 1,
				Minor: 5,
				Patch: 5,
			},
			comparisonVersion: &semantika.Version{
				Major: 2,
				Minor: 3,
				Patch: 4,
			},
			expectedResult: false,
		},

		// tests with referenceVersion = comparisonVersion
		{
			name: "TestOlderThan_8.9.4_8.9.4",
			referenceVersion: &semantika.Version{
				Major: 8,
				Minor: 9,
				Patch: 4,
			},
			comparisonVersion: &semantika.Version{
				Major: 8,
				Minor: 9,
				Patch: 4,
			},
			expectedResult: true,
		},

		// tests with referenceVersion > comparisonVersion
		{
			name: "TestOlderThan_3.3.4_3.2.3",
			referenceVersion: &semantika.Version{
				Major: 3,
				Minor: 3,
				Patch: 4,
			},
			comparisonVersion: &semantika.Version{
				Major: 3,
				Minor: 2,
				Patch: 3,
			},
			expectedResult: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if res := test.referenceVersion.NewerThanOrEquals(test.comparisonVersion); res != test.expectedResult {
				t.Errorf("expected: %v, result: %v", test.expectedResult, res)
			}
		})
	}
}

// TestNewerThan calls Version.NewerThan() with different reference and comparison versions, and checks for valid returns.
func TestNewerThan(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name              string
		referenceVersion  *semantika.Version
		comparisonVersion *semantika.Version
		expectedResult    bool
	}{
		// tests with referenceVersion < comparisonVersion
		{
			name: "TestOlderThan_1.5.5_2.3.4",
			referenceVersion: &semantika.Version{
				Major: 1,
				Minor: 5,
				Patch: 5,
			},
			comparisonVersion: &semantika.Version{
				Major: 2,
				Minor: 3,
				Patch: 4,
			},
			expectedResult: false,
		},

		// tests with referenceVersion = comparisonVersion
		{
			name: "TestOlderThan_8.9.4_8.9.4",
			referenceVersion: &semantika.Version{
				Major: 8,
				Minor: 9,
				Patch: 4,
			},
			comparisonVersion: &semantika.Version{
				Major: 8,
				Minor: 9,
				Patch: 4,
			},
			expectedResult: false,
		},

		// tests with referenceVersion > comparisonVersion
		{
			name: "TestOlderThan_3.3.4_3.2.3",
			referenceVersion: &semantika.Version{
				Major: 3,
				Minor: 3,
				Patch: 4,
			},
			comparisonVersion: &semantika.Version{
				Major: 3,
				Minor: 2,
				Patch: 3,
			},
			expectedResult: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if res := test.referenceVersion.NewerThan(test.comparisonVersion); res != test.expectedResult {
				t.Errorf("expected: %v, result: %v", test.expectedResult, res)
			}
		})
	}
}
