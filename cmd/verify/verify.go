package verify

import (
	"fmt"
	"github.com/spf13/cobra"
	"k8s.io/cluster-bootstrap/token/jws"
)

// Options is a struct to support verify command
type Options struct {
	Content     string
	TokenId     string
	TokenSecret string
	JWSToken    string
}

// NewOptions return initialized Options
func NewOptions() *Options {
	return &Options{}
}

func NewCmdVerify() *cobra.Command {
	o := NewOptions()
	cmd := &cobra.Command{
		Use:   "verify",
		Short: "jws verify : jws verify --content <content> --token-id <token id> --token-secret <token secret> --jws-token <jws token>",
		Long: `
jws verify : jws verify --content <content> --token-id <token id> --token-secret <token secret>

content: the content(base64) that wanted to be verify, such as kube-public/configmap/cluster-info.data.kubeconfig
token id: the token id, such as kube-system/secret/bootstrap-token-<token id>.data.token-id
token secret:  the token secret, such as kube-system/secret/bootstrap-token-<token id>.data.token-secret
jws token: the jws token, such as kube-public/configmap/cluster-info.data.jws-token-<token id>`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("jws verify --content %s --token-id %s --token-secret %s --jws-token %s\n", o.Content, o.TokenId, o.TokenSecret, o.JWSToken)
			Verify(o)
		},
	}

	cmd.Flags().StringVarP(&o.Content, "content", "c", o.Content, "the content(base64)")
	cmd.Flags().StringVarP(&o.TokenId, "token-id", "i", o.TokenId, "the token id")
	cmd.Flags().StringVarP(&o.TokenSecret, "token-secret", "s", o.TokenSecret, "the token secret")
	cmd.Flags().StringVarP(&o.JWSToken, "jws-token", "j", o.JWSToken, "the jws token that wanted to be verify")
	return cmd
}

func Verify(o *Options) {
	if !bootstrap.DetachedTokenIsValid(o.JWSToken, o.Content, o.TokenId, o.TokenSecret) {
		fmt.Println("failed to verify JWS signature of received content")
		return
	}
	fmt.Println("success to verify JWS signature of received content")
}
