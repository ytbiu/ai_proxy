package common

import (
	"testing"
)

func TestVerifySignature(t *testing.T) {
	type args struct {
		msg     string
		sig     string
		address string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "",
			args: args{
				msg:     "0x307835663330353238666239643637376536383866623063616266363931316262353861636561663039643461363438313534383835303231316565303631343236",
				sig:     "0xa4064506fb2af477e39baeac4cd270621aa49399f89ef63747885e207cee516b60be2da7a5ce467735b1fa99fcff9cd82a8e5066c1fef473968bf4118fa4be701b",
				address: "0xde184A6809898D81186DeF5C0823d2107c001Da2",
			},
			want:    true,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := VerifySignature(tt.args.msg, tt.args.sig, tt.args.address)
			if (err != nil) != tt.wantErr {
				t.Errorf("VerifySignature() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("VerifySignature() got = %v, want %v", got, tt.want)
			}
		})
	}
}
