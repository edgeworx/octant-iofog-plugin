/*
 *  *******************************************************************************
 *  * Copyright (c) 2019 Edgeworx, Inc.
 *  *
 *  * This program and the accompanying materials are made available under the
 *  * terms of the Eclipse Public License v. 2.0 which is available at
 *  * http://www.eclipse.org/legal/epl-2.0
 *  *
 *  * SPDX-License-Identifier: EPL-2.0
 *  *******************************************************************************
 *
 */

package v1

import (
  "github.com/eclipse-iofog/iofog-go-sdk/pkg/apps"
  metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ApplicationSpec struct {
  Microservices []apps.Microservice `json:"microservices"`
  Routes        []apps.Route        `json:"routes"`
  Replicas      int32               `json:"replicas"`
}

type ApplicationStatus struct {
  Replicas      int32    `json:"replicas"`
  LabelSelector string   `json:"labelSelector"`
  PodNames      []string `json:"podNames"`
}

type Application struct {
  metav1.TypeMeta   `json:",inline"`
  metav1.ObjectMeta `json:"metadata,omitempty"`

  Spec   ApplicationSpec   `json:"spec,omitempty"`
  Status ApplicationStatus `json:"status,omitempty"`
}

type ApplicationList struct {
  metav1.TypeMeta `json:",inline"`
  metav1.ListMeta `json:"metadata,omitempty"`
  Items           []Application `json:"items"`
}
