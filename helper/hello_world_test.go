package helper

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"runtime"
	"testing"
)

func BenchmarkTable(b *testing.B) {
	benchmarks := []struct {
		name    string
		request string
	}{
		{
			name:    "Kahfi",
			request: "Kahfi",
		},
		{
			name:    "Kurnia",
			request: "Kurnia",
		},
		{
			name:    "KahfiKurniaAji",
			request: "Kahfi Kurnia Aji",
		},
		{
			name:    "Fauzan",
			request: "Fauzan Nugraha",
		},
	}
	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.request)
			}
		})
	}
}

func BenchmarkSub(b *testing.B) {
	b.Run("Kahfi", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Kahfi")
		}
	})
	b.Run("Kurnia", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Kurnia")
		}
	})
}

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Kahfi")
	}
}

func BenchmarkHelloWorldKurnia(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Kurnia")
	}
}

func TestTableHelloWorld(t *testing.T) {
	var tests = []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Kahfi",
			request:  "Kahfi",
			expected: "Hello Kahfi",
		},
		{
			name:     "Kurnia",
			request:  "Kurnia",
			expected: "Hello Kurnia",
		},
		{
			name:     "Aji",
			request:  "Aji",
			expected: "Hello Aji",
		},
		{
			name:     "Fauzan",
			request:  "Fauzan",
			expected: "Hello Fauzan",
		},
		{
			name:     "Widi",
			request:  "Widi",
			expected: "Hello Widi",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

func TestSubTest(t *testing.T) {
	t.Run("Kahfi", func(t *testing.T) {
		var result string = HelloWorld("Kahfi")
		require.Equal(t, "Hello Kahfi", result, "Result must be 'Hello Kahfi'")
	})

	t.Run("Kurnia", func(t *testing.T) {
		var result string = HelloWorld("Kurnia")
		require.Equal(t, "Hello Kurnia", result, "Result must be 'Hello Kurnia'")
	})

	t.Run("Aji", func(t *testing.T) {
		var result string = HelloWorld("Aji")
		require.Equal(t, "Hello Aji", result, "Result must be 'Hello Aji'")
	})
}

func TestMain(m *testing.M) {
	// before
	fmt.Println("BEFORE UNIT TEST")

	m.Run()

	// after
	fmt.Println("AFTER UNIT TEST")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "linux" {
		t.Skip("Can't run on Linux")
	}

	var result string = HelloWorld("Kahfi")
	require.Equal(t, "Hello Kahfi", result, "Result must be 'Hello Kahfi'")
}

func TestHelloWorldRequire(t *testing.T) {
	var result string = HelloWorld("Kahfi")
	require.Equal(t, "Hello Kahfi", result, "Result must be 'Hello Kahfi'")
	fmt.Println("TestHelloWorld with Require done !!!!!!!!!!!!!")
}

func TestHelloWorldAssert(t *testing.T) {
	var result string = HelloWorld("Kahfi")
	assert.Equal(t, "Hello Kahfi", result, "Result must be 'Hello Kahfi'")
	fmt.Println("TestHelloWorld with Assert done !!!!!!!!!!!!!")
}

func TestHelloWorldKahfi(t *testing.T) {
	var result string = HelloWorld("Kahfi")

	if result != "Hello Kahfi" {
		// error
		t.Error("Result must be 'Hello Kahfi'")
	}

	fmt.Println("===== TestHelloWorldKahfi Done (Selesai) =====")
}

func TestHelloWorldAji(t *testing.T) {
	var result string = HelloWorld("Aji")

	if result != "Hello Aji" {
		// error
		t.Fatal("Result must be 'Hello Aji'")
	}

	fmt.Println("===== TestHelloWorldAji Done (Selesai) =====")
}

func TestHelloWorldKurnia(t *testing.T) {
	var result string = HelloWorld("Kurnia")

	if result != "Hello Kurnia" {
		// error
		t.Fatal("Result must be 'Hello Kurnia'")
	}

	fmt.Println("===== TestHelloWorldKurnia Done (Selesai) =====")
}
