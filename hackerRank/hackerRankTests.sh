#!/usr/bin/env bash

# Terminal colors for output formatting
GREEN='\033[0;32m'
RED='\033[0;31m'
NC='\033[0m' # No Color

run_test() {
    local test_name="$1"
    local expected_code="$2"
    local expected_output="$3"
    shift 3 # Remove the first 3 arguments, leaving the command to execute

    # Execute command and capture combined stdout/stderr along with exit code
    local actual_output
    actual_output=$("$@" 2>&1)
    local actual_code=$?

    # Verify results
    if [ "$actual_code" -eq "$expected_code" ] && [ "$actual_output" = "$expected_output" ]; then
        echo "[ ${GREEN}PASS${NC} ] $test_name"
    else
        echo "[ ${RED}FAIL${NC} ] $test_name"
        echo "   Expected exit code: $expected_code, got: $actual_code"
        echo "   Expected output:   '$expected_output'"
        echo "   Actual output:     '$actual_output'"
    fi
}

echo "=== Running Tests ==="
echo ""

### TEST CASES

run_test "mars-exploration - 0 errors" 0 "0" go run mars-exploration-go/main.go SOSSOS

run_test "mars-exploration - 1 error" 0 "1" go run mars-exploration-go/main.go SOSSIS

run_test "mars-exploration - 3 errors" 0 "3" go run mars-exploration-go/main.go SOSSOS999

