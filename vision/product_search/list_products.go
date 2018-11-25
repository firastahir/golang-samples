// Copyright 2018 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package main

// [START imports]
import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	
	"context"

	"google.golang.org/api/iterator"
	vision "cloud.google.com/go/vision/apiv1"
	visionpb "google.golang.org/genproto/googleapis/cloud/vision/v1"
)
// [END imports]

// [START vision_product_search_list_products]
func listProducts(project string, location string) error {
	ctx := context.Background()
	c, err := vision.NewProductSearchClient(ctx)
	if err != nil {
		return err
	}

	req := &visionpb.ListProductsRequest{}
	req.Parent = fmt.Sprintf("projects/%s/locations/%s", project, location)

	it := c.ListProducts(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}

		fmt.Println("Product name:", resp.Name)
		fmt.Println("Product display name:", resp.DisplayName)
		fmt.Println("Product category:", resp.ProductCategory)
		fmt.Println("Product labels:", resp.ProductLabels, "\n")
	}

	return nil
}
// [END vision_product_search_list_products]

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s <project-id> <location>\n", filepath.Base(os.Args[0]))
	}
	flag.Parse()

	args := flag.Args()
	if len(args) < 2 {
		flag.Usage()
		os.Exit(1)
	}

	err := listProducts(args[0], args[1])

	if err != nil {
		fmt.Println("Error:", err)
	}
}
