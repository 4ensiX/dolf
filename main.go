package main

import (
	"fmt"
	"io"
	"os"

	"dolf/util"

	"github.com/docker/docker/client"
	"golang.org/x/net/context"
)

func saveImage(id string) (io.ReadCloser, error) {
        var err error
        var cli *client.Client

        ctx := context.Background()

        cli, err = client.NewClientWithOpts(client.FromEnv)
        if err != nil {
                return nil, err
        }

        readCloser, err := cli.ImageSave(ctx,[]string{id})
        if err != nil {
                return nil, err
        }

        return readCloser, nil
}


func main() {

    if len(os.Args) < 2 {
        fmt.Println("how to use: dolf [image name]")
        return 
    }else if len(os.Args) > 2 {
        fmt.Println("invalid args")
        return 
    }
    var id string = os.Args[1]

    reader, err := saveImage(id)
	if err != nil {
		fmt.Println(err)
        os.Exit(1)
	}
	defer reader.Close()

	layer, _ := util.DLtar(reader)

    dir := "temp"
    err = os.Mkdir(dir, 0755)

    if err != nil {
        fmt.Println(err)
        return
    }

    for i, _ := range layer {
        wf, err := os.Create("temp/" + layer[i].Layer + ".txt")
        if err != nil {
            fmt.Println(err)
            return
        }
        defer wf.Close()
        for _, temp := range layer[i].LayerFiles {
            wf.WriteString(temp + "\n")
        }
    }
}
