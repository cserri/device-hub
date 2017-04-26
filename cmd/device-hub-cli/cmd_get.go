// Copyright © 2017 thingful

package main

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/fiorix/protoc-gen-cobra/iocodec"
	"github.com/spf13/cobra"
	"github.com/thingful/device-hub/proto"
)

var getCommand = &cobra.Command{
	Use:   "get",
	Short: "Display one or many resources",
	Run: func(cmd *cobra.Command, args []string) {

		v := proto.GetRequest{
			Filter: strings.Join(args, ","),
		}

		err := roundTrip(v, func(cli proto.HubClient, in iocodec.Decoder, out iocodec.Encoder) error {

			resp, err := cli.Get(context.Background(), &v)

			if err != nil {
				return err
			}

			return out.Encode(resp)

		})
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(args)
	},
}
