package main

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
    cloudflare "github.com/pulumi/pulumi-cloudflare/sdk/v3/go/cloudflare"
    "github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
        c := config.New(ctx, "")
        zoneId := c.Require("ZONE_ID")
        _, err := cloudflare.NewRecord(ctx, "wildcard", &cloudflare.RecordArgs{
            Name: pulumi.String("*"),
            ZoneId: pulumi.String(zoneId),
            Type: pulumi.String("A"),
            Value: pulumi.String("192.168.1.80"),
            Ttl: pulumi.Int(3600),
        })
        if err != nil {
            return nil
        }

        return nil
	})
}
