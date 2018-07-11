/*******************************************************************************
 * Copyright © 2017-2018 VMware, Inc. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the License
 * is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
 * or implied. See the License for the specific language governing permissions and limitations under
 * the License.
 *******************************************************************************/

package main

import (
	"github.com/edgexfoundry-holding/edgex-ui-go/initial"
	"github.com/edgexfoundry-holding/edgex-ui-go/web/app"
	"github.com/edgexfoundry-holding/edgex-ui-go/web/app/mongo"
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"
	"github.com/edgexfoundry-holding/edgex-ui-go/web/app/common"
)

func main() {

	ok := mongo.DBConnect()
	if !ok {
		return
	}

	initial.Initialize()

	r := app.InitRestRoutes()

	server := &http.Server{
		Handler:      common.GeneralFilter(r),
		Addr:         ":4000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("EdgeX ui server listen at " + server.Addr)

	log.Fatal(server.ListenAndServe())
}