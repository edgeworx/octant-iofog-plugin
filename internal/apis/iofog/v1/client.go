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
  "k8s.io/client-go/kubernetes/scheme"
  "k8s.io/client-go/rest"
)

type KogV1Interface interface {
  Kogs(namespace string) KogInterface
}

type KogV1Client struct {
  restClient rest.Interface
}

func NewForConfig(c *rest.Config) (*KogV1Client, error) {
  config := *c
  config.ContentConfig.GroupVersion = &SchemeGroupVersion
  config.APIPath = "/apis"
  config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()
  config.UserAgent = rest.DefaultKubernetesUserAgent()

  client, err := rest.RESTClientFor(&config)
  if err != nil {
    return nil, err
  }

  return &KogV1Client{restClient: client}, nil
}

func (c *KogV1Client) Kogs(namespace string) KogInterface {
  return &kogClient{
    restClient: c.restClient,
    ns:         namespace,
  }
}
