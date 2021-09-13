package ssm

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
)

func Run(cfg *aws.Config) {
	svc := ssm.New(session.New(), cfg)
	getAll(svc)
}

func getAll(svc *ssm.SSM) {
	path := "/st/qa/"
	out, err := svc.GetParametersByPath(&ssm.GetParametersByPathInput{
		Path:           aws.String(path),
		WithDecryption: aws.Bool(false),
	})
	if err != nil {
		panic(out)
	}

	parameters := parametersToMap(path, out.Parameters)
	j, err := json.Marshal(parameters)
	if err != nil {
		panic(out)
	}

	fmt.Printf("Parameters: %v", out.Parameters)
	fmt.Printf("JSON Parameters: %s", j)
}

func parametersToMap(path string, parameters []*ssm.Parameter) map[string]string {
	res := make(map[string]string, len(parameters))
	for _, v := range parameters {
		res[strings.Replace(*v.Name, path, "", 1)] = *v.Value
	}
	return res
}
