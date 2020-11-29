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

package api

//   Init()     Initiating stage: Init plugin by config
//    ||
//    \/
//   Prepare()   Preparing stage: Prepare the Forwarder, such as get remote client.
//    ||
//    \/
//   Forward()  Running stage: Forward the batch events
//    ||
//    \/
//   Close()    Closing stage: Close the Collector, such as close connection with SkyWalking javaagent.

// Forwarder is a plugin interface, that defines new forwarders.
type Forwarder interface {
	Initializer
	Preparer
	Closer

	// Forward the batch events to the external services, such as Kafka MQ and SkyWalking OAP cluster.
	Forward(batch BatchEvents)
}
