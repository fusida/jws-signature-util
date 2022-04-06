package sign

import (
	"fmt"
	"github.com/spf13/cobra"
	bootstrap "k8s.io/cluster-bootstrap/token/jws"
)

// Options is a struct to support sign command
type Options struct {
	Content     string
	TokenId     string
	TokenSecret string
}

// NewOptions return initialized Options
func NewOptions() *Options {
	return &Options{}
}

func NewCmdSign() *cobra.Command {
	o := NewOptions()
	cmd := &cobra.Command{
		Use:   "sign",
		Short: "jws sign : jws sign --content <content> --token-id <token id> --token-secret <token secret>",
		Long: `
jws sign : jws sign --content <content> --token-id <token id> --token-secret <token secret>

content: the content(base64) that wanted to be sign, such as kube-public/configmap/cluster-info.data.kubeconfig
token id: the token id, such as kube-system/secret/bootstrap-token-<token id>.data.token-id
token secret:  the token secret, such as kube-system/secret/bootstrap-token-<token id>.data.token-secret`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("jws sign --content %s --token-id %s --token-secret %s\n", o.Content, o.TokenId, o.TokenSecret)
			Sign(o)
		},
	}

	cmd.Flags().StringVarP(&o.Content, "content", "c", o.Content, "the content(base64) that wanted to be sign")
	cmd.Flags().StringVarP(&o.TokenId, "token-id", "i", o.TokenId, "the token id")
	cmd.Flags().StringVarP(&o.TokenSecret, "token-secret", "s", o.TokenSecret, "the token secret")
	return cmd
}

func Sign(o *Options) {
	jwsToken, err := bootstrap.ComputeDetachedSignature(o.Content, o.TokenId, o.TokenSecret)
	if err != nil {
		fmt.Printf("fail to sign: %v\n", err)
		return
	}
	fmt.Printf("success sign: %s\n", jwsToken)
}
