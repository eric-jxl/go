package redis

import "testing"

func TestInitRedis(t *testing.T) {
	r := InitRedis()
	pong,err := r.Ping(ctx).Result()
	if err != nil {
		t.Fatalf("%v",err)
	}
	t.Log(pong)
}

func TestLock(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Lock",
			args{
				"get",
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Lock(tt.args.key); got != tt.want {
				t.Errorf("Lock() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnLock(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			"Unlock",
			args{"get"},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UnLock(tt.args.key); got != tt.want {
				t.Errorf("UnLock() = %v, want %v", got, tt.want)
			}
		})
	}
}