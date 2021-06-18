package main

import (
	"testing"
	"time"
)

func Test_shouldBeKicked(t *testing.T) {
	tests := []struct {
		name     string
		roles    []string
		joinedAt time.Time
		want     bool
	}{
		{
			name:     "has_role_should_not_kick",
			roles:    []string{"hi"},
			joinedAt: time.Now().Add(-time.Hour),
			want:     false,
		},
		{
			name:     "no_role_recently_joined_should_not_kick",
			roles:    []string{},
			joinedAt: time.Now().Add(-time.Second),
			want:     false,
		},
		{
			name:     "no_role_joined_long_time_ago_should_kick",
			roles:    []string{},
			joinedAt: time.Now().Add(-time.Hour),
			want:     true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := shouldBeKicked(tt.roles, tt.joinedAt)
			if got != tt.want {
				t.Errorf("shouldBeKicked() got = %v, want %v", got, tt.want)
			}
		})
	}
}
