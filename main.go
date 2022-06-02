package main

import (
	"fmt"
	"io"
	"os"
        "strconv"

//	"dolf/util"
	"github.com/4ensiX/dolf/util"

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

    layer, layers_sum := util.DLtar(reader)

    dir := "temp"
    err = os.Mkdir(dir, 0755)

    if err != nil {
        fmt.Println(err)
        //return
    }

    for i, _ := range layer {
        var wf *os.File
        var err error
        for j, _ :=  range layers_sum.Manifest {
            if layer[i].Layer == layers_sum.Manifest[j] {
                wf, err = os.Create("temp/" + layers_sum.Img_id[j] + ".txt")
                break
            }
        }
        if err != nil {
            fmt.Println(err)
            return
        }
        defer wf.Close()
        for _, temp := range layer[i].LayerFiles {
            wf.WriteString(temp + "\n")
        }
    }
    lay, err := os.Create("temp/layers.txt")
    for i, _ := range layers_sum.Manifest {
        lay.WriteString(strconv.Itoa(i+1) + " " + layers_sum.Manifest[i] + " " + layers_sum.Img_id[i] + "\n")
    }
}
