package gorange

import (
	"reflect"
	"testing"
)

func TestExtractRange(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []int
		wantErr  bool
	}{
		{
			name:     "正常范围",
			input:    "1-3, 5, 7-9",
			expected: []int{1, 2, 3, 5, 7, 8, 9},
			wantErr:  false,
		},
		{
			name:     "中文逗号",
			input:    "1，3-5，7",
			expected: []int{1, 3, 4, 5, 7},
			wantErr:  false,
		},
		{
			name:     "带噪声数据",
			input:    "xxx1x---x3x, uie4kjdf, ---88---",
			expected: []int{1, 2, 3, 4, 88},
			wantErr:  false,
		},
		{
			name:     "空输入",
			input:    "",
			expected: []int{},
			wantErr:  false,
		},
		{
			name:     "单个数字",
			input:    "5",
			expected: []int{5},
			wantErr:  false,
		},
		{
			name:     "多个范围",
			input:    "1-3,5-7,9-10",
			expected: []int{1, 2, 3, 5, 6, 7, 9, 10},
			wantErr:  false,
		},
		{
			name:     "重复数字",
			input:    "1,1-3,3",
			expected: []int{1, 1, 2, 3, 3},
			wantErr:  false,
		},
		{
			name:     "多余的分隔符",
			input:    "1,,2,,,3",
			expected: []int{1, 2, 3},
			wantErr:  false,
		},
		{
			name:     "多余的横杠",
			input:    "1----3, ----5----",
			expected: []int{1, 2, 3, 5},
			wantErr:  false,
		},
		{
			name:     "非法输入被忽略",
			input:    "abc,1-3,def,5,ghi",
			expected: []int{1, 2, 3, 5},
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ExtractRange(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExtractRange() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("ExtractRange() = %v, want %v", got, tt.expected)
			}
		})
	}
}

// 测试边界情况
func TestExtractRangeEdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []int
		wantErr  bool
	}{
		{
			name:     "范围颠倒",
			input:    "3-1",
			expected: []int{1, 2, 3},
			wantErr:  false,
		},
		{
			name:     "特殊字符",
			input:    "!@#$%^&*()1-3",
			expected: []int{1, 2, 3},
			wantErr:  false,
		},
		{
			name:     "大量空格",
			input:    "   1   -   3   ,   5   ",
			expected: []int{1, 2, 3, 5},
			wantErr:  false,
		},
		{
			name:     "噪声数据",
			input:    "xxx3-1xxx, xjlkjfd13slsv-sdf",
			expected: []int{1, 2, 3, 13},
			wantErr:  false,
		},
		{
			name:     "多余的横杠",
			input:    " 1----   3  ，5 ",
			expected: []int{1, 2, 3, 5},
			wantErr:  false,
		},
		{
			name:     "只有分隔符",
			input:    ",,,,",
			expected: []int{},
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ExtractRange(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExtractRange() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("ExtractRange() = %v, want %v", got, tt.expected)
			}
		})
	}
}
