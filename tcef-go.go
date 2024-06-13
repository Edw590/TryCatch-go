/*******************************************************************************
 * Copyright 2023-2023 Edw590
 *
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 ******************************************************************************/

package Tcef

// Main credits: https://dzone.com/articles/try-and-catch-in-golang

type Exception interface{}

func Throw(up Exception) {
	panic(up)
}

func (tcef Tcef) Do() {
	var panicked bool = true

	if tcef.Finally != nil {
		defer tcef.Finally()
	}
	if tcef.Else != nil {
		defer func() {
			if !panicked {
				tcef.Else()
			}
		}()
	}
	defer func() {
		if panicked {
			var recovered any = recover()
			if tcef.Catch != nil {
				tcef.Catch(recovered)
			}
		}
	}()

	tcef.Try()
	panicked = false
}

type Tcef struct {
	Try     func()
	Catch   func(Exception)
	Else    func()
	Finally func()
}
