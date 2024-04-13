package wc_test

import (
	"basic-wc/wc"
	"testing"
)

const ART_OF_WAR_WORD_COUNT = 58164
const TEST_WORD_DATA_COUNT = 9
const TEST_BYTE_DATA_COUNT = 10
const TEST_LINE_DATA_COUNT = 20
const TEST_CHAR_DATA_COUNT = 38
const TEST_CHAR_DATA_WORD_COUNT = 7
const TEST_CHAR_DATA_LINE_COUNT = 1
const TEST_CHAR_DATA_BYTE_COUNT = 60

func TestWCWordCount(t *testing.T) {
	// Get the word count only
	statTypes := wc.StatTypes{WordRetrieve: true}

	statCounts, _ := wc.GetStats("testdata/test_word_data.txt", statTypes)

	if statCounts.WordCount != TEST_WORD_DATA_COUNT {
		t.Fatalf("Word Count failed, expected=%v, got=%v", TEST_WORD_DATA_COUNT, statCounts.WordCount)
	}
}

func TestWCLineCount(t *testing.T) {
	statTypes := wc.StatTypes{LineRetrieve: true}

	statCounts, _ := wc.GetStats("testdata/test_line_data.txt", statTypes)

	if statCounts.LineCount != TEST_LINE_DATA_COUNT {
		t.Fatalf("Line Count failed, expected=%v, got=%v", TEST_LINE_DATA_COUNT, statCounts.LineCount)
	}
}

func TestWCByteCount(t *testing.T) {
	statTypes := wc.StatTypes{ByteRetrieve: true}

	statCounts, _ := wc.GetStats("testdata/test_byte_data.txt", statTypes)

	if statCounts.ByteCount != TEST_BYTE_DATA_COUNT {
		t.Fatalf("Byte Count failed, expected=%v, got=%v", TEST_BYTE_DATA_COUNT, statCounts.ByteCount)
	}
}

func TestWCCharCount(t *testing.T) {
	statTypes := wc.StatTypes{CharRetrieve: true}

	statCounts, _ := wc.GetStats("testdata/test_char_data.txt", statTypes)

	if statCounts.CharCount != TEST_CHAR_DATA_COUNT {
		t.Fatalf("Char Count failed, expected=%v, got=%v", TEST_CHAR_DATA_COUNT, statCounts.CharCount)
	}
}

func TestWCAllCount(t *testing.T) {
	statTypes := wc.StatTypes{CharRetrieve: true, LineRetrieve: true, ByteRetrieve: true, WordRetrieve: true}

	statCounts, _ := wc.GetStats("testdata/test_char_data.txt", statTypes)

	if statCounts.CharCount != TEST_CHAR_DATA_COUNT {
		t.Fatalf("Char Count failed, expected=%v, got=%v", TEST_CHAR_DATA_COUNT, statCounts.CharCount)
	}

	if statCounts.ByteCount != TEST_CHAR_DATA_BYTE_COUNT {
		t.Fatalf("Byte Count failed, expected=%v, got=%v", TEST_CHAR_DATA_BYTE_COUNT, statCounts.ByteCount)
	}

	if statCounts.LineCount != TEST_CHAR_DATA_LINE_COUNT {
		t.Fatalf("Line Count failed, expected=%v, got=%v", TEST_CHAR_DATA_LINE_COUNT, statCounts.LineCount)
	}

	if statCounts.WordCount != TEST_CHAR_DATA_WORD_COUNT {
		t.Fatalf("Word Count failed, expected=%v, got=%v", TEST_CHAR_DATA_WORD_COUNT, statCounts.WordCount)
	}
}
