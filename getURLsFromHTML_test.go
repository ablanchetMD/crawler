package main

import (
	"reflect"
	"testing"
)

func TestGetURLsFromHTML(t *testing.T) {
	tests := []struct {
		name      string
		inputURL  string
		inputBody string
		expected  []string
	}{
		{
			name:     "absolute and relative URLs",
			inputURL: "https://blog.boot.dev",
			inputBody: `
		<html>
			<body>
				<a href="/path/one">
					<span>Boot.dev</span>
				</a>
				<a href="https://other.com/path/one">
					<span>Boot.dev</span>
				</a>
			</body>
		</html>
		`,
			expected: []string{"https://blog.boot.dev/path/one", "https://other.com/path/one"},
		},
		{
			name:     "absolute and relative URLs",
			inputURL: "https://blog.boot.dev",
			inputBody: `
		<html>
			<body>
				<a href="/path/one">
					<span>Boot.dev</span>
				</a>
				<div>
					<a href="/path/two">
						<span>Boot.dev</span>
					</a>
				</div>				
				<a href="https://other.com/path/one">
					<span>Boot.dev</span>
				</a>
			</body>
		</html>
		`,
			expected: []string{"https://blog.boot.dev/path/one", "https://blog.boot.dev/path/two", "https://other.com/path/one"},
		},
		{
			name:     "absolute and relative URLs",
			inputURL: "https://blog.boot.dev",
			inputBody: `
		<html>
			<body>
				<a href="/path/one">
					<span>Boot.dev</span>
				</a>
				<div>
					<div>
						<a href="http://zod.com/path/one">
						<span>Boot.dev</span>
						</a>
					</div>
				</div>
				<img src="https://blog.boot.dev/path/one">			
				
				<a href="/test/path/two">
						<span>Boot.dev</span>
				</a>
			</body>			
		</html>
		`,
			expected: []string{"https://blog.boot.dev/path/one", "http://zod.com/path/one", "https://blog.boot.dev/test/path/two"},
		},
		{
			name:     "wagslane.dev",
			inputURL: "https://wagslane.dev",
			inputBody: `
<html>			
<body class="bg-gray-100 text-gray-800">
  <nav class="z-50 text-gray-800 p-2">
  <div class="max-w-7xl mx-auto px-2 sm:px-6 lg:px-8 top-nav-bar-height">
    <div class="relative flex items-center justify-between h-full">

      <div class="absolute inset-y-0 left-0 flex items-center sm:hidden">
        <button type="button" onclick="clickMobileMenuToggle()"
          class="inline-flex items-center justify-center p-2 rounded-md text-gray-200 hover:text-white bg-gray-700 focus:outline-none focus:ring-2 focus:ring-inset focus:ring-white"
          aria-controls="mobile-menu" aria-expanded="false">
          <span class="sr-only">Open main menu</span>
          <svg class="block h-6 w-6" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24"
            stroke="currentColor" aria-hidden="true">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16M4 18h16" />
          </svg>
          <svg class="hidden h-6 w-6" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24"
            stroke="currentColor" aria-hidden="true">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
          </svg>
        </button>
      </div>

      <div class="h-10"></div>

      <div class="flex-1 flex items-center justify-center sm:items-stretch sm:justify-evenly">
        <div class="hidden sm:block sm:ml-6">
          <div class="flex space-x-4 h-full items-center">
            <a href="/" class="hover:bg-gray-700 hover:text-white px-3 py-2 rounded">
              Articles
            </a>

            <a href="/tags/" class="hover:bg-gray-700 hover:text-white px-3 py-2 rounded">
              Tags
            </a>

            <a href="/about/" class="hover:bg-gray-700 hover:text-white px-3 py-2 rounded">
              About
            </a>

            <a href="/index.xml" class="hover:bg-gray-700 hover:text-white px-3 py-2 rounded">
              RSS
            </a>
          </div>
        </div>
      </div>
    </div>
  </div>
</nav>
</body>
</html>
`, expected: []string{"https://wagslane.dev/", "https://wagslane.dev/tags/", "https://wagslane.dev/about/", "https://wagslane.dev/index.xml"},
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := getURLsFromHTML(tc.inputBody, tc.inputURL)
			if err != nil {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			}
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Test %v - %s FAIL: expected URL: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}
