package chapter3

//func TestKadai3(t *testing.T) {
//	type args struct {
//		x interface{}
//	}
//	tests := []struct {
//		name    string
//		args    args
//		want    string
//		wantErr bool
//	}{
//		{
//			name: "dog",
//			args:args{x:Dog{}},
//			want: "わんわん",
//			wantErr: false,
//		},
//		{
//			name: "cat",
//			args:args{x:Cat{}},
//			want: "にゃーにゃ",
//			wantErr: false,
//		},
//		{
//			name: "else",
//			args:args{x:"hoge"},
//			wantErr: true,
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			got, err := Kadai3(tt.args.x)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("Kadai3() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if got != tt.want {
//				t.Errorf("Kadai3() got = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
