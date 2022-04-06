# jws-signature-util
检查 k8s jws-token 是否正确  

使用方法
```bash
./bin/jws -h                     

jws sign or verify: jws <sign or verify> --content <content> --token-id <token id> --token-secret <token secret>

content: the content(base64) that wanted to be sign, such as kube-public/configmap/cluster-info.data.kubeconfig
token id: the token id, such as kube-system/secret/bootstrap-token-<token id>.data.token-id
token secret: the token secret, such as kube-system/secret/bootstrap-token-<token id>.data.token-secret
jws token: the jws token, such as kube-public/configmap/cluster-info.data.jws-token-<token id>

Usage:
  jws [flags]
  jws [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  sign        jws sign : jws sign --content <content> --token-id <token id> --token-secret <token secret>
  verify      jws verify : jws verify --content <content> --token-id <token id> --token-secret <token secret> --jws-token <jws token>

Flags:
  -h, --help   help for jws

Use "jws [command] --help" for more information about a command.

```