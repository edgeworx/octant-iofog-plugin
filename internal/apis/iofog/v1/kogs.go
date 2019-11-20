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
  "k8s.io/apimachinery/pkg/watch"
  "k8s.io/client-go/kubernetes/scheme"
  "k8s.io/client-go/rest"
  "time"
)

type KogInterface interface {
  List(opts metav1.ListOptions) (*KogList, error)
  Get(name string, options metav1.GetOptions) (*Kog, error)
  Create(*Kog) (*Kog, error)
  Watch(opts metav1.ListOptions) (watch.Interface, error)
  Delete(name string, options *metav1.DeleteOptions) error
  Update(*Kog) (*Kog, error)
}

type kogClient struct {
  restClient rest.Interface
  ns         string
}

func (c *kogClient) List(opts metav1.ListOptions) (*KogList, error) {
  var timeout time.Duration
  if opts.TimeoutSeconds != nil {
    timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
  }

  result := &KogList{}
  err := c.restClient.
    Get().
    Namespace(c.ns).
    Resource("kogs").
    VersionedParams(&opts, scheme.ParameterCodec).
    Timeout(timeout).
    Do().
    Into(result)

  return result, err
}

func (c *kogClient) Get(name string, opts metav1.GetOptions) (*Kog, error) {
  result := &Kog{}
  err := c.restClient.
    Get().
    Namespace(c.ns).
    Resource("kogs").
    Name(name).
    VersionedParams(&opts, scheme.ParameterCodec).
    Do().
    Into(result)

  return result, err
}

func (c *kogClient) Create(kog *Kog) (*Kog, error) {
  result := &Kog{}
  err := c.restClient.
    Post().
    Namespace(c.ns).
    Resource("kogs").
    Body(kog).
    Do().
    Into(result)

  return result, err
}

func (c *kogClient) Watch(opts metav1.ListOptions) (watch.Interface, error) {
  var timeout time.Duration
  if opts.TimeoutSeconds != nil {
    timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
  }

  opts.Watch = true
  return c.restClient.
    Get().
    Namespace(c.ns).
    Resource("kogs").
    VersionedParams(&opts, scheme.ParameterCodec).
    Timeout(timeout).
    Watch()
}

func (c *kogClient) Delete(name string, opts *metav1.DeleteOptions) error {
  return c.restClient.
    Delete().
    Namespace(c.ns).
    Resource("kogs").
    Name(name).
    Body(opts).
    Do().
    Error()
}

func (c *kogClient) Update(kog *Kog) (*Kog, error) {
  result := &Kog{}
  err := c.restClient.
    Put().
    Namespace(c.ns).
    Resource("kogs").
    Name(kog.Name).
    Body(kog).
    Do().
    Into(result)

  return result, err
}

func init() {
  AddToScheme(scheme.Scheme)
}
