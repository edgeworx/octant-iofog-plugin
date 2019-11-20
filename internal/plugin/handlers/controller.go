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

package handlers

import (
  "github.com/eclipse-iofog/octant-plugin/internal/plugin"
  "github.com/vmware-tanzu/octant/pkg/plugin/service"
  "github.com/vmware-tanzu/octant/pkg/view/component"
)

func ControllerHandler(root bool) func(request *service.Request) (component.ContentResponse, error) {
  return func(request *service.Request) (component.ContentResponse, error) {
    if !root || plugin.ECN == nil {
      return component.ContentResponse{}, nil
    }

    contentResponse := component.NewContentResponse(component.TitleFromString("Control Plane"))
    content := plugin.SummaryHelper("Control Plane", []plugin.SectionsContent{
      {
        Label: "Name",
        Value: plugin.ECN.Kog.Spec.ControlPlane.IofogUser.Name,
      },
      {
        Label: "Surname",
        Value: plugin.ECN.Kog.Spec.ControlPlane.IofogUser.Surname,
      },
      {
        Label: "Email",
        Value: plugin.ECN.Kog.Spec.ControlPlane.IofogUser.Email,
      },
      {
        Label: "Controller IP",
        Value: plugin.ECN.Kog.Spec.ControlPlane.LoadBalancerIP,
      },
    })
    contentResponse.Add(content)
    return *contentResponse, nil
  }
}
