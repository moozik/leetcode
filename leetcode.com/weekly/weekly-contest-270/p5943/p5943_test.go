package p5943

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_deleteMiddle(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			args: args{head: &ListNode{Val: 2, Next: &ListNode{Val: 1, Next: nil}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteMiddle(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				for {
					if got == nil {
						break
					}
					fmt.Printf("%+v", got)
					got = got.Next
				}
				t.Errorf("deleteMiddle() = %v, want %v", got, tt.want)
			}
		})
	}
}
