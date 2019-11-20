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

package k8s

import (
  "errors"
  "fmt"
  "github.com/eclipse-iofog/iofog-go-sdk/pkg/client"
  kogv1 "github.com/eclipse-iofog/octant-plugin/internal/apis/iofog/v1"
  extsclientset "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
  k8serrors "k8s.io/apimachinery/pkg/api/errors"
  metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
  "k8s.io/client-go/kubernetes"
  restclient "k8s.io/client-go/rest"
  "k8s.io/client-go/tools/clientcmd"
)

type Kubernetes struct {
  config          *restclient.Config
  kogClient       kogv1.KogV1Client
  clientset       *kubernetes.Clientset
  extsClientset   *extsclientset.Clientset
  ns              string
  kogInstanceName string
}

type Deployment struct {
  Agents               map[string]client.AgentInfo
  Connectors           map[string]client.ConnectorInfo
  Flows                map[string]client.FlowInfo
  Kog                  kogv1.Kog
  MicroservicesPerFlow map[string]map[string]client.MicroserviceInfo
}

func NewKubernetes(configFilename, namespace string) (*Kubernetes, error) {
  config, err := clientcmd.BuildConfigFromFlags("", configFilename)
  if err != nil {
    return nil, err
  }

  clientset, err := kubernetes.NewForConfig(config)
  if err != nil {
    return nil, err
  }
  extsClientset, err := extsclientset.NewForConfig(config)
  if err != nil {
    return nil, err
  }
  kogClient, err := kogv1.NewForConfig(config)
  if err != nil {
    return nil, err
  }

  return &Kubernetes{
    config:          config,
    kogClient:       *kogClient,
    clientset:       clientset,
    extsClientset:   extsClientset,
    ns:              namespace,
    kogInstanceName: "iokog",
  }, nil
}

func (k8s *Kubernetes) NamespaceExists(namespace string) error {
  if _, err := k8s.clientset.CoreV1().Namespaces().Get(namespace, metav1.GetOptions{}); err != nil {
    if k8serrors.IsNotFound(err) {
      return errors.New("Could not find Namespace " + namespace + " on Kubernetes cluster")
    }
    return err
  }

  return nil
}

func (k8s *Kubernetes) GetKogDeployment() (*kogv1.Kog, error) {
  kogsList, err := k8s.kogClient.Kogs(k8s.ns).List(metav1.ListOptions{})
  if err != nil {
    return nil, err
  }

  if len(kogsList.Items) > 0 {
    return &kogsList.Items[0], nil
  }

  return nil, nil
}

func (k8s *Kubernetes) GetControllerEndpoint(controllerIp string) (endpoint string, err error) {
  controllerPort := 51121
  endpoint = fmt.Sprintf("%s:%d", controllerIp, controllerPort)

  return
}

func (k8s *Kubernetes) GetControllerIp() (controllerIp string, err error) {
  services, err := k8s.clientset.CoreV1().Services(k8s.ns).List(metav1.ListOptions{LabelSelector: "name=controller"})
  if err != nil {
    return "", err
  }
  if len(services.Items) != 1 {
    return "", errors.New("Could not find Controller Service in Kubernetes namespace " + k8s.ns)
  }

  return services.Items[0].Status.LoadBalancer.Ingress[0].IP, nil
}