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
  metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type KogSpec struct {
  ControlPlane ControlPlane `json:"controlPlane"`
  Connectors   Connectors   `json:"connectors,omitempty"`
}

type Connectors struct {
  Image     string      `json:"image"`
  Instances []Connector `json:"instances"`
}

type Connector struct {
  Name string `json:"name"`
}

type ControlPlane struct {
  IofogUser              IofogUser `json:"iofogUser"`
  ControllerReplicaCount int32     `json:"controllerReplicaCount"`
  Database               Database  `json:"database,omitempty"`
  ControllerImage        string    `json:"controllerImage"`
  ImagePullSecret        string    `json:"imagePullSecret,omitempty"`
  KubeletImage           string    `json:"kubeletImage"`
  ServiceType            string    `json:"serviceType"`
  LoadBalancerIP         string    `json:"loadBalancerIp,omitempty"`
}

type Database struct {
  Provider     string `json:"provider"`
  Host         string `json:"host"`
  Port         int    `json:"port"`
  User         string `json:"user"`
  Password     string `json:"password"`
  DatabaseName string `json:"databaseName"`
}

type IofogUser struct {
  Name     string `json:"name"`
  Surname  string `json:"surname"`
  Email    string `json:"email"`
  Password string `json:"password"`
}

type KogStatus struct {
  ControllerPods []string `json:"controllerPods"`
}

type Kog struct {
  metav1.TypeMeta   `json:",inline"`
  metav1.ObjectMeta `json:"metadata,omitempty"`

  Spec   KogSpec   `json:"spec,omitempty"`
  Status KogStatus `json:"status,omitempty"`
}

type KogList struct {
  metav1.TypeMeta `json:",inline"`
  metav1.ListMeta `json:"metadata,omitempty"`
  Items           []Kog `json:"items"`
}
