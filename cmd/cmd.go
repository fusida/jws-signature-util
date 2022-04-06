package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"jws-signature-util/cmd/sign"
	"jws-signature-util/cmd/verify"
)

func NewJwsCommand() *cobra.Command {
	cmds := &cobra.Command{
		Use:   "jws",
		Short: "jws sign or verify",
		Long: `
jws sign or verify: jws <sign or verify> --content <content> --token-id <token id> --token-secret <token secret>

content: the content(base64) that wanted to be sign, such as kube-public/configmap/cluster-info.data.kubeconfig
token id: the token id, such as kube-system/secret/bootstrap-token-<token id>.data.token-id
token secret: the token secret, such as kube-system/secret/bootstrap-token-<token id>.data.token-secret
jws token: the jws token, such as kube-public/configmap/cluster-info.data.jws-token-<token id>`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("jws sign or verify")
		},
	}
	cmds.AddCommand(sign.NewCmdSign())
	cmds.AddCommand(verify.NewCmdVerify())
	return cmds
}
