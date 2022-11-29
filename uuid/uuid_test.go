package main

import "testing"

func TestSnowflake(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			"Snowflake",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("Snowflake() = %v", Snowflake())
		})
	}
}

func TestXidGenerate(t *testing.T) {
	tests := []struct {
		name          string
		containerName string
	}{
		{
			"XidGenerate",
			"Eric",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := XidGenerate(tt.containerName)
			if got != "" {
				t.Logf("XidGenerate() = %v ,Len:%d", got,len(got))
			}
		})
	}
}
