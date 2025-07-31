// Copyright (c) 2020, NVIDIA CORPORATION.  All rights reserved.
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

package nvml

import (
	"C"

	"k8s.io/klog/v2"
)

// nvml.Init()
func (l *library) Init() Return {
	if err := l.load(); err != nil {
		return ERROR_LIBRARY_NOT_FOUND
	}
	klog.Info("nvmlInit start")
	ret := nvmlInit()
	klog.Info("nvmlInit end")
	return ret
}

// nvml.InitWithFlags()
func (l *library) InitWithFlags(flags uint32) Return {
	if err := l.load(); err != nil {
		return ERROR_LIBRARY_NOT_FOUND
	}
	return nvmlInitWithFlags(flags)
}

// nvml.Shutdown()
func (l *library) Shutdown() Return {
	ret := nvmlShutdown()
	if ret != SUCCESS {
		return ret
	}

	err := l.close()
	if err != nil {
		return ERROR_UNKNOWN
	}

	return ret
}
