/*
Copyright 2019 The Knative Authors.

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

package v1beta1

import (
	"context"

	"knative.dev/pkg/apis"
	v1 "knative.dev/serving/pkg/apis/serving/v1"
)

// ConvertUp implements apis.Convertible
func (source *Configuration) ConvertUp(ctx context.Context, obj apis.Convertible) error {
	switch sink := obj.(type) {
	case *v1.Configuration:
		sink.ObjectMeta = source.ObjectMeta
		sink.Spec = source.Spec
		sink.Status = source.Status
		return nil
	default:
		return apis.ConvertUpViaProxy(ctx, source, &v1.Configuration{}, sink)
	}
}

// ConvertDown implements apis.Convertible
func (sink *Configuration) ConvertDown(ctx context.Context, obj apis.Convertible) error {
	switch source := obj.(type) {
	case *v1.Configuration:
		sink.ObjectMeta = source.ObjectMeta
		sink.Spec = source.Spec
		sink.Status = source.Status
		return nil
	default:
		return apis.ConvertDownViaProxy(ctx, source, &v1.Configuration{}, sink)
	}
}
