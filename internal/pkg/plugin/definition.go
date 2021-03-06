// Licensed to Apache Software Foundation (ASF) under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Apache Software Foundation (ASF) licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package plugin

// Plugin defines the plugin model in Satellite.
type Plugin interface {
	// Name returns the name of the specific plugin.
	Name() string
	// Description returns the description of the specific plugin.
	Description() string
	// Config returns the default config, that is a YAML pattern.
	DefaultConfig() string
}

// SharingPlugin the plugins cloud be sharing with different modules in different namespaces.
type SharingPlugin interface {
	Plugin

	// Prepare the sharing plugins, such as build the connection with the external services.
	Prepare() error
	// Close the sharing plugin.
	Close() error
}

// Config is used to initialize the DefaultInitializingPlugin.
type Config map[string]interface{}

// NameField is a required field in Config.
const NameField = "plugin_name"
