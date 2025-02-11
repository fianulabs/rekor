//
// Copyright 2022 The Sigstore Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cose

import (
	"context"
	"sync"
	"testing"

	fuzzUtils "github.com/sigstore/rekor/pkg/fuzz"
	"github.com/sigstore/rekor/pkg/types"
)

var initter sync.Once

func FuzzCreateProposedEntry(f *testing.F) {
	f.Fuzz(func(t *testing.T, version string) {
		initter.Do(fuzzUtils.SetFuzzLogger)
		ctx := context.Background()
		brt := New()
		props := types.ArtifactProperties{}
		_, _ = brt.CreateProposedEntry(ctx, version, props)
	})
}
