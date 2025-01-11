// Copyright 2022 Camunda Services GmbH
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package camunda

import (
	"path/filepath"
	"strings"
	"testing"

	"github.com/gruntwork-io/terratest/modules/helm"
	"github.com/gruntwork-io/terratest/modules/k8s"
	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	corev1 "k8s.io/api/core/v1"
)

type configMapTemplateTest struct {
	suite.Suite
	chartPath string
	release   string
	namespace string
	templates []string
}

func TestConfigMapTemplate(t *testing.T) {
	t.Parallel()

	chartPath, err := filepath.Abs("../../../")
	require.NoError(t, err)

	suite.Run(t, &configMapTemplateTest{
		chartPath: chartPath,
		release:   "camunda-platform-test",
		namespace: "camunda-platform-" + strings.ToLower(random.UniqueId()),
		templates: []string{"templates/camunda/configmap-identity-auth.yaml"},
	})
}

func (s *configMapTemplateTest) TestConfigMapIdentityIssuerURL() {
	// given
	options := &helm.Options{
		SetValues: map[string]string{
			"global.identity.auth.issuerBackendUrl": "http://keycloak:80/auth/realms/camunda-platform",
		},
		KubectlOptions: k8s.NewKubectlOptions("", "", s.namespace),
	}

	// when
	output := helm.RenderTemplate(s.T(), options, s.chartPath, s.release, s.templates)
	var configmap corev1.ConfigMap
	helm.UnmarshalK8SYaml(s.T(), output, &configmap)

	// then
	s.Require().Equal("http://keycloak:80/auth/realms/camunda-platform", configmap.Data["CAMUNDA_IDENTITY_ISSUER_BACKEND_URL"])
}

func (s *configMapTemplateTest) TestConfigMapIdentityIssuerURLWithKeycloakURLSyntax() {
	// given
	options := &helm.Options{
		SetValues: map[string]string{
			"global.identity.keycloak.url.protocol": "http",
			"global.identity.keycloak.url.host":     "keycloak",
			"global.identity.keycloak.url.port":     "80",
			"global.identity.keycloak.contextPath":  "/auth/realms/",
			"global.identity.keycloak.realm":        "camunda-platform",
		},
		KubectlOptions: k8s.NewKubectlOptions("", "", s.namespace),
	}

	// when
	output := helm.RenderTemplate(s.T(), options, s.chartPath, s.release, s.templates)
	var configmap corev1.ConfigMap
	helm.UnmarshalK8SYaml(s.T(), output, &configmap)

	// then
	s.Require().Equal("http://keycloak:80/auth/realms/camunda-platform", configmap.Data["CAMUNDA_IDENTITY_ISSUER_BACKEND_URL"])
}
