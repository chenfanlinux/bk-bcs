/*
 * Tencent is pleased to support the open source community by making Blueking Container Service available.
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package types

type BcsDeployment struct {
	TypeMeta   `json:",inline"`
	ObjectMeta `json:"metadata"`

	Spec BcsDeploymentSpec `json:"spec"`

	UpPolicy      UpdatePolicy  `json:"updatePolicy,omitempty"`
	RestartPolicy RestartPolicy `json:"restartPolicy,omitempty"`
	KillPolicy    KillPolicy    `json:"killPolicy,omitempty"`
	Constraints   *Constraint   `json:"constraint,omitempty"`
}

type BcsDeploymentSpec struct {
	Instance int `json:"instance"`
	//RestartPolicy RestartPolicy     `json:"restartPolicy,omitempty"`
	Selector map[string]string `json:"selector,omitempty"`
	Template *PodTemplateSpec  `json:"template"`
	Strategy UpgradeStrategy   `json:"strategy"`

	// pause allows you to pause the rolling update process when you are in it.
	// default value is false.
	PauseDeployment bool `json:"pauseDeployment"`
}

type UpgradeStrategy struct {
	Type          UpgradeStrategyType `json:"type"`
	RollingUpdate *RollingUpdate      `json:"rollingupdate"`
}

type UpgradeStrategyType string
type RollingOrderType string

const (
	RecreateUpgradeStrategyType      UpgradeStrategyType = "Recreate"
	RollingUpdateUpgradeStrategyType UpgradeStrategyType = "RollingUpdate"

	// ForceUpdate means that all the old pods will be deleted. and then recreate
	// pods with new deployment.
	// ForceUpdate is "valid" only when you call the k8s update restful api.
	ForceUpdateStrategyType UpgradeStrategyType = "ForceUpdate"

	// CreateFirstOrder means that the new pod will be created and then delete the old pod
	// during the whole rolling update operation process. DeleteFirstOrder is quite the opposite.
	CreateFirstOrder RollingOrderType = "CreateFirst"
	DeleteFirstOrder RollingOrderType = "DeleteFirst"
)

type RollingUpdate struct {
	// The maximum number of pods that can be unavailable during the update.
	// This can not be 0 if MaxSurge is 0.
	// By default, a fixed value of 1 is used.
	MaxUnavailable int `json:"maxUnavilable"`

	// The maximum number of pods that can be scheduled above the original number of pods.
	// By default, a value of 1 is used.
	MaxSurge int `json:"maxSurge"`

	// the time duration between the rolling update operation.
	// in second unit. By default, a value of 10s is used.
	UpgradeDuration uint32 `json:"upgradeDuration"`

	// rolling update will do step by step manully or automatically, if manully, it will pause after every step.
	// by default is false
	RollingManually bool `json:"rollingManually"`

	// rolling update order, create first or delete first.
	// by default, a value of CreateFirst is used.
	RollingOrder RollingOrderType `json:"rollingOrder"`
}
