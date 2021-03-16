package terraformcli

import (
	"fmt"
	"reflect"
	"testing"
)

const testRunnerStaterListSuccessWriteStdout = `
null_resource.a
null_resource.b
null_resource.c
`

const testRunnerStaterPullSuccessWriteStdout = `
{
  "version": 4,
  "terraform_version": "0.14.0",
  "serial": 1,
  "lineage": "9454d3ba-725d-93a0-4ec2-b600ce1ff3b0",
  "outputs": {},
  "resources": [
    {
      "mode": "managed",
      "type": "null_resource",
      "name": "a",
      "provider": "provider[\"registry.terraform.io/hashicorp/null\"]",
      "instances": [
        {
          "schema_version": 0,
          "attributes": {
            "id": "8101096104556242691",
            "triggers": null
          },
          "sensitive_attributes": [],
          "private": "bnVsbA=="
        }
      ]
    },
    {
      "mode": "managed",
      "type": "null_resource",
      "name": "b",
      "provider": "provider[\"registry.terraform.io/hashicorp/null\"]",
      "instances": [
        {
          "schema_version": 0,
          "attributes": {
            "id": "4993183039109330160",
            "triggers": null
          },
          "sensitive_attributes": [],
          "private": "bnVsbA=="
        }
      ]
    },
    {
      "mode": "managed",
      "type": "null_resource",
      "name": "c",
      "provider": "provider[\"registry.terraform.io/hashicorp/null\"]",
      "instances": [
        {
          "schema_version": 0,
          "attributes": {
            "id": "4613527663746728648",
            "triggers": null
          },
          "sensitive_attributes": [],
          "private": "bnVsbA=="
        }
      ]
    }
  ]
}
`

func TestRunnerStater_List(t *testing.T) {
	type fields struct {
		runner Runner
	}
	type args struct {
		options ListOptions
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				runner: &mockRunner{
					writeStdout: []byte(testRunnerStaterListSuccessWriteStdout),
				},
			},
			want: []string{
				"null_resource.a",
				"null_resource.b",
				"null_resource.c",
			},
		},
		{
			name: "error",
			fields: fields{
				runner: &mockRunner{
					returnErr: fmt.Errorf("error"),
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &RunnerStater{
				runner: tt.fields.runner,
			}
			got, err := o.List(tt.args.options)
			if (err != nil) != tt.wantErr {
				t.Errorf("RunnerStater.List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RunnerStater.List() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRunnerStater_Move(t *testing.T) {
	type fields struct {
		runner Runner
	}
	type args struct {
		src     string
		dst     string
		options MoveOptions
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				runner: &mockRunner{},
			},
			args: args{
				src: "null_resource.src",
				dst: "null_resource.dst",
			},
		},
		{
			name: "success with options",
			fields: fields{
				runner: &mockRunner{},
			},
			args: args{
				src: "null_resource.src",
				dst: "null_resource.dst",
				options: MoveOptions{
					State:    "./state.tfstate",
					StateOut: "./stateout.tfstate",
				},
			},
		},
		{
			name: "error",
			fields: fields{
				runner: &mockRunner{
					returnErr: fmt.Errorf("error"),
				},
			},
			args: args{
				src: "null_resource.src",
				dst: "null_resource.dst",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &RunnerStater{
				runner: tt.fields.runner,
			}
			if err := o.Move(tt.args.src, tt.args.dst, tt.args.options); (err != nil) != tt.wantErr {
				t.Errorf("RunnerStater.Move() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRunnerStater_Pull(t *testing.T) {
	type fields struct {
		runner Runner
	}
	type args struct {
		options PullOptions
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				runner: &mockRunner{
					writeStdout: []byte(testRunnerStaterPullSuccessWriteStdout),
				},
			},
			want: []byte(testRunnerStaterPullSuccessWriteStdout),
		},
		{
			name: "error",
			fields: fields{
				runner: &mockRunner{
					returnErr: fmt.Errorf("error"),
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &RunnerStater{
				runner: tt.fields.runner,
			}
			got, err := o.Pull(tt.args.options)
			if (err != nil) != tt.wantErr {
				t.Errorf("RunnerStater.Pull() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RunnerStater.Pull() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRunnerStater_Push(t *testing.T) {
	type fields struct {
		runner Runner
	}
	type args struct {
		path    string
		options PushOptions
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				runner: &mockRunner{},
			},
		},
		{
			name: "error",
			fields: fields{
				runner: &mockRunner{
					returnErr: fmt.Errorf("error"),
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &RunnerStater{
				runner: tt.fields.runner,
			}
			if err := o.Push(tt.args.path, tt.args.options); (err != nil) != tt.wantErr {
				t.Errorf("RunnerStater.Push() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewRunnerStater(t *testing.T) {
	type args struct {
		runner Runner
	}
	tests := []struct {
		name string
		args args
		want *RunnerStater
	}{
		{
			name: "success",
			args: args{
				runner: &mockRunner{},
			},
			want: &RunnerStater{
				runner: &mockRunner{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRunnerStater(tt.args.runner); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRunnerStater() = %v, want %v", got, tt.want)
			}
		})
	}
}
