// +build ignore

/*
 * Minio Cloud Storage, (C) 2017 Minio, Inc.
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
 */

package main

import (
	"log"

	"github.com/minio/minio/pkg/madmin"
)

func main() {
	// Note: YOUR-ACCESSKEYID, YOUR-SECRETACCESSKEY are
	// dummy values, please replace them with original values.

	// Note: YOUR-ACCESSKEYID, YOUR-SECRETACCESSKEY are
	// dummy values, please replace them with original values.

	// API requests are secure (HTTPS) if secure=true and insecure (HTTPS) otherwise.
	// New returns an Minio Admin client object.
	madmClnt, err := madmin.New("your-minio.example.com:9000", "YOUR-ACCESSKEYID", "YOUR-SECRETACCESSKEY", true)
	if err != nil {
		log.Fatalln(err)
	}

	// Heal upload mybucket/myobject/uploadID - dry run.
	isDryRun := true
	_, err = madmClnt.HealUpload("mybucket", "myobject", "myUploadID", isDryRun)
	if err != nil {
		log.Fatalln(err)
	}

	// Heal upload mybucket/myobject/uploadID - this time for real.
	isDryRun = false
	healResult, err := madmClnt.HealUpload("mybucket", "myobject", "myUploadID", isDryRun)
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Heal result for mybucket/myobject/myUploadID: %#v\n", healResult)
}
