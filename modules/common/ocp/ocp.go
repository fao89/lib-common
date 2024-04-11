/*
Copyright 2024 Red Hat

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

package ocp

import (
	"context"

	"github.com/fao89/lib-common/modules/common/helper"

	"gopkg.in/yaml.v3"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
)

// IsFipsCluster - Check if OCP has fips enabled which is a day 1 operation
func IsFipsCluster(ctx context.Context, h *helper.Helper) (bool, error) {
	configMap := &corev1.ConfigMap{}
	err := h.GetClient().Get(ctx, types.NamespacedName{Name: "cluster-config-v1", Namespace: "kube-system"}, configMap)
	if err != nil {
		return false, err
	}

	var installConfig map[string]interface{}
	installConfigYAML := configMap.Data["install-config"]
	err = yaml.Unmarshal([]byte(installConfigYAML), &installConfig)
	if err != nil {
		return false, err
	}

	fipsEnabled, ok := installConfig["fips"].(bool)
	if !ok {
		return false, nil
	}
	return fipsEnabled, nil
}
