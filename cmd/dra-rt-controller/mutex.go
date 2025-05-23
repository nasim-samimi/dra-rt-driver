/*
 * Copyright 2023 The Kubernetes Authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"sync"
)

type PerNodeMutex struct {
	sync.Mutex
	submutex map[string]*sync.Mutex
}

func NewPerNodeMutex() *PerNodeMutex {
	return &PerNodeMutex{
		submutex: make(map[string]*sync.Mutex),
	}
}

func (pnm *PerNodeMutex) Get(node string) *sync.Mutex {
	pnm.Mutex.Lock()
	defer pnm.Mutex.Unlock()
	if pnm.submutex[node] == nil {
		pnm.submutex[node] = &sync.Mutex{}
	}
	return pnm.submutex[node]
}

type PerCoreMutex struct {
	sync.Mutex
	submutex map[string]*sync.Mutex
}

func NewPerCoreMutex() *PerCoreMutex {
	return &PerCoreMutex{
		submutex: make(map[string]*sync.Mutex),
	}
}

func (pcm *PerCoreMutex) Get(coreID string) *sync.Mutex {
	pcm.Lock()
	defer pcm.Unlock()
	if pcm.submutex[coreID] == nil {
		pcm.submutex[coreID] = &sync.Mutex{}
	}
	return pcm.submutex[coreID]
}
