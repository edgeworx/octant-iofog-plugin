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
  "k8s.io/apimachinery/pkg/runtime"
)

func (in *Application) DeepCopyInto(out *Application) {
  *out = *in
  out.TypeMeta = in.TypeMeta
  in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
  in.Spec.DeepCopyInto(&out.Spec)
  in.Status.DeepCopyInto(&out.Status)
  return
}

func (in *Application) DeepCopy() *Application {
  if in == nil {
    return nil
  }
  out := new(Application)
  in.DeepCopyInto(out)
  return out
}

func (in *Application) DeepCopyObject() runtime.Object {
  if c := in.DeepCopy(); c != nil {
    return c
  }
  return nil
}

func (in *ApplicationList) DeepCopyInto(out *ApplicationList) {
  *out = *in
  out.TypeMeta = in.TypeMeta
  out.ListMeta = in.ListMeta
  if in.Items != nil {
    in, out := &in.Items, &out.Items
    *out = make([]Application, len(*in))
    for i := range *in {
      (*in)[i].DeepCopyInto(&(*out)[i])
    }
  }
  return
}

func (in *ApplicationList) DeepCopy() *ApplicationList {
  if in == nil {
    return nil
  }
  out := new(ApplicationList)
  in.DeepCopyInto(out)
  return out
}

func (in *ApplicationList) DeepCopyObject() runtime.Object {
  if c := in.DeepCopy(); c != nil {
    return c
  }
  return nil
}

func (in *ApplicationSpec) DeepCopyInto(out *ApplicationSpec) {
  *out = *in
  if in.Microservices != nil {
    in, out := &in.Microservices, &out.Microservices
    *out = make([]apps.Microservice, len(*in))
    for i := range *in {
      (*in)[i].DeepCopyInto(&(*out)[i])
    }
  }
  if in.Routes != nil {
    in, out := &in.Routes, &out.Routes
    *out = make([]apps.Route, len(*in))
    copy(*out, *in)
  }
  return
}

func (in *ApplicationSpec) DeepCopy() *ApplicationSpec {
  if in == nil {
    return nil
  }
  out := new(ApplicationSpec)
  in.DeepCopyInto(out)
  return out
}

func (in *ApplicationStatus) DeepCopyInto(out *ApplicationStatus) {
  *out = *in
  if in.PodNames != nil {
    in, out := &in.PodNames, &out.PodNames
    *out = make([]string, len(*in))
    copy(*out, *in)
  }
  return
}

func (in *ApplicationStatus) DeepCopy() *ApplicationStatus {
  if in == nil {
    return nil
  }
  out := new(ApplicationStatus)
  in.DeepCopyInto(out)
  return out
}

func (in *Connector) DeepCopyInto(out *Connector) {
  *out = *in
  return
}

func (in *Connector) DeepCopy() *Connector {
  if in == nil {
    return nil
  }
  out := new(Connector)
  in.DeepCopyInto(out)
  return out
}

func (in *Connectors) DeepCopyInto(out *Connectors) {
  *out = *in
  if in.Instances != nil {
    in, out := &in.Instances, &out.Instances
    *out = make([]Connector, len(*in))
    copy(*out, *in)
  }
  return
}

func (in *Connectors) DeepCopy() *Connectors {
  if in == nil {
    return nil
  }
  out := new(Connectors)
  in.DeepCopyInto(out)
  return out
}

func (in *ControlPlane) DeepCopyInto(out *ControlPlane) {
  *out = *in
  out.IofogUser = in.IofogUser
  out.Database = in.Database
  return
}

func (in *ControlPlane) DeepCopy() *ControlPlane {
  if in == nil {
    return nil
  }
  out := new(ControlPlane)
  in.DeepCopyInto(out)
  return out
}

func (in *Database) DeepCopyInto(out *Database) {
  *out = *in
  return
}

func (in *Database) DeepCopy() *Database {
  if in == nil {
    return nil
  }
  out := new(Database)
  in.DeepCopyInto(out)
  return out
}

func (in *IofogUser) DeepCopyInto(out *IofogUser) {
  *out = *in
  return
}

func (in *IofogUser) DeepCopy() *IofogUser {
  if in == nil {
    return nil
  }
  out := new(IofogUser)
  in.DeepCopyInto(out)
  return out
}

func (in *Kog) DeepCopyInto(out *Kog) {
  *out = *in
  out.TypeMeta = in.TypeMeta
  in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
  in.Spec.DeepCopyInto(&out.Spec)
  in.Status.DeepCopyInto(&out.Status)
  return
}

func (in *Kog) DeepCopy() *Kog {
  if in == nil {
    return nil
  }
  out := new(Kog)
  in.DeepCopyInto(out)
  return out
}

func (in *Kog) DeepCopyObject() runtime.Object {
  if c := in.DeepCopy(); c != nil {
    return c
  }
  return nil
}

func (in *KogList) DeepCopyInto(out *KogList) {
  *out = *in
  out.TypeMeta = in.TypeMeta
  out.ListMeta = in.ListMeta
  if in.Items != nil {
    in, out := &in.Items, &out.Items
    *out = make([]Kog, len(*in))
    for i := range *in {
      (*in)[i].DeepCopyInto(&(*out)[i])
    }
  }
  return
}

func (in *KogList) DeepCopy() *KogList {
  if in == nil {
    return nil
  }
  out := new(KogList)
  in.DeepCopyInto(out)
  return out
}

func (in *KogList) DeepCopyObject() runtime.Object {
  if c := in.DeepCopy(); c != nil {
    return c
  }
  return nil
}

func (in *KogSpec) DeepCopyInto(out *KogSpec) {
  *out = *in
  out.ControlPlane = in.ControlPlane
  in.Connectors.DeepCopyInto(&out.Connectors)
  return
}

func (in *KogSpec) DeepCopy() *KogSpec {
  if in == nil {
    return nil
  }
  out := new(KogSpec)
  in.DeepCopyInto(out)
  return out
}

func (in *KogStatus) DeepCopyInto(out *KogStatus) {
  *out = *in
  if in.ControllerPods != nil {
    in, out := &in.ControllerPods, &out.ControllerPods
    *out = make([]string, len(*in))
    copy(*out, *in)
  }
  return
}

func (in *KogStatus) DeepCopy() *KogStatus {
  if in == nil {
    return nil
  }
  out := new(KogStatus)
  in.DeepCopyInto(out)
  return out
}
