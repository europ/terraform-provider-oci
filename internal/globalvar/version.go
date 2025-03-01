// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package globalvar

import (
	"log"
)

const Version = "4.107.0"
const ReleaseDate = "2023-02-09"

func PrintVersion() {
	log.Printf("[INFO] terraform-provider-oci %s\n", Version)
}
