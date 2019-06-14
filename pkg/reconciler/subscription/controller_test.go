/*
Copyright 2019 The Knative Authors

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

package subscription

import (
	"testing"

	"github.com/knative/pkg/configmap"
	logtesting "github.com/knative/pkg/logging/testing"
	. "github.com/knative/pkg/reconciler/testing"

	_ "github.com/knative/pkg/injection/informers/apiextinformers/apiextensionsv1beta1/crd/fake"

	// Fake injection informers
	_ "github.com/knative/eventing/pkg/client/injection/informers/eventing/v1alpha1/subscription/fake"
)

func TestNew(t *testing.T) {
	defer logtesting.ClearAll()
	ctx, _ := SetupFakeContext(t)

	c := NewController(ctx, configmap.NewFixedWatcher())

	if c == nil {
		t.Fatal("Expected NewController to return a non-nil value")
	}
}