/*
Copyright 2022 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package kind

import (
	"bytes"
	_ "embed"
	"fmt"
	"text/template"
)

//go:embed kind.yaml.tpl
var kindYamlTpl string

var kindYamlTemplate = template.Must(template.New("_").Parse(kindYamlTpl))

func BuildKind(conf BuildKindConfig) (string, error) {
	buf := bytes.NewBuffer(nil)
	err := kindYamlTemplate.Execute(buf, conf)
	if err != nil {
		return "", fmt.Errorf("failed to execute kind yaml template: %w", err)
	}
	return buf.String(), nil
}

type BuildKindConfig struct {
	KubeApiserverPort uint32
	PrometheusPort    uint32

	RuntimeConfig []string
	FeatureGates  []string
}
