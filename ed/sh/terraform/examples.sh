# tf xamples

# aws.st
cd ed/sh/terraform/examples/aws.st
# aws lambda
export GOPATH=$PWD/ed/l/go/examples/aws
go get -u "github.com/aws/aws-sdk-go/aws"
# build
GOOS=linux go build -o /tmp/awsLambdaOne $GOPATH/src/app/lambda
cd /tmp && zip awsLambdaOne.zip awsLambdaOne && mv awsLambdaOne.zip /Users/kovpakvolodymyr/Downloads && cd -
# test
open https://realtimelog.herokuapp.com:443/64kfym341kp2
go run $GOPATH/src/app/main.go k31 val200 200
for i in $(seq 2000 2999); do go run $GOPATH/src/app/main.go "k$i" val200 $i; done

# aws.ec2
cd ed/sh/terraform/examples/aws.ec2
c=plan
c=apply
c=refresh
c=destroy
terraform $c -var-file=ec2.tfvars -lock=false

# aws.x
# @see: docket image build in k8s
cd ed/sh/terraform/examples/aws.x/environments/dev
export AWS_PROFILE=x
terraform init
terraform plan
terraform apply
# run
h=''
key=/Users/kovpakvolodymyr/web/kovpak/gh/ed/sh/ssh/examples/nopwd/id_rsa
scp -i $key /tmp/xgoapp ec2-user@$h:/tmp
# ssh and run bin file
ssh -i $key ec2-user@$h

#
cd ed/sh/terraform/examples/aws.k8s
terraform init
terraform plan
terraform apply
terraform destroy
cfg=kubeconfig_my-cluster
d=~/web/kovpak/gh/
kubectl --kubeconfig=$cfg get pods
kubectl --kubeconfig=$cfg apply --force=true -f $d/ed/sh/kubernetes/examples/log/pod.yaml
kubectl --kubeconfig=$cfg apply --force=true -f $d/ed/sh/kubernetes/examples/sh/pod.yaml
kubectl --kubeconfig=$cfg delete pod log-pod
kubectl --kubeconfig=$cfg delete pod ksh-pod
