/*
 * This file is part of the KubeVirt project
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
 *
 * Copyright 2017, 2018 Red Hat, Inc.
 *
 */

package watch

var HotMigrateMaps = map[string][]string{
	"host-model-cpu.node.kubevirt.io/Cascadelake-Server":  {"Cascadelake-Server", "Icelake-Server"},
	"host-model-cpu.node.kubevirt.io/Icelake-Server":      {"Cascadelake-Server", "Icelake-Server"},
	"host-model-cpu.node.kubevirt.io/Haswell-noTSX":       {"Haswell-noTSX", "Haswell-noTSX-IBRS"},
	"host-model-cpu.node.kubevirt.io/Haswell-noTSX-IBRS":  {"Haswell-noTSX", "Haswell-noTSX-IBRS"},
	"host-model-cpu.node.kubevirt.io/Broadwell-IBRS":      {"Broadwell-IBRS"},
	"host-model-cpu.node.kubevirt.io/IvyBridge":           {"IvyBridge", "IvyBridge-IBRS"},
	"host-model-cpu.node.kubevirt.io/IvyBridge-IBRS":      {"IvyBridge", "IvyBridge-IBRS"},
	"host-model-cpu.node.kubevirt.io/EPYC-Milan":          {"EPYC-Milan"},
	"host-model-cpu.node.kubevirt.io/Cooperlake":          {"Cooperlake", "EPYC"},
	"host-model-cpu.node.kubevirt.io/EPYC":                {"Cooperlake", "EPYC"},
	"host-model-cpu.node.kubevirt.io/Skylake-Server-IBRS": {"Skylake-Server-IBRS"},
}
