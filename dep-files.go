package main

import (
	"fmt"
	"os"
    "strconv"
    "bufio"
    "path/filepath"
    "strings"
)

const layers_file string = "layers.txt"
var splitf []string

func main() {

    if len(os.Args) < 2 {
        fmt.Println("how to use?")
        return
    }else if len(os.Args) > 2 {
        fmt.Println("invalid args")
        return
    }
    var search_dir string = os.Args[1]

    layersf ,err := os.Open(search_dir + "/" + layers_file)
    if err != nil {
        fmt.Println(err)
        return
    }
    defer layersf.Close()

    scanner := bufio.NewScanner(layersf)
    var layers []string
    for scanner.Scan() {
        split := strings.Split(scanner.Text()," ")
        layers =  append(layers,split[2])
    }

    var lflist string
    for i, l := range layers {
        err :=  filepath.Walk(search_dir, func(path string, info os.FileInfo, err error) error {
                    if err != nil {
                        return err
                    }
                    lflist = search_dir + "/" + l + ".txt"
                    if path == lflist {
                        create_graph(lflist,i)
                    }
                    return nil
                })
        if err != nil {
            fmt.Println(err)
            return
        }
    }
}


func create_graph(layer_files_list string, layer_number int) {
    var lfl *os.File
    var err error

    lfl, err = os.Open(layer_files_list)
    if err != nil {
        fmt.Println(err)
        return
    }
    defer lfl.Close()

    var line_tmp []string
    var tmp string
    scanner := bufio.NewScanner(lfl)
    for scanner.Scan() {
        var x int = 0
        tmp = scanner.Text()
        for i,f := range splitf {
            line_tmp = strings.Split(f," ")
            if strings.Compare(tmp,line_tmp[0]) == 0 {
                line_tmp[1] = line_tmp[1] + "," + strconv.Itoa(layer_number)
                splitf[i] = line_tmp[0] + " " + line_tmp[1]
                if err != nil {
                    fmt.Println(err)
                    return
                }
                x = 1
                break
            }
        }
        if x == 1 {continue}
        ltmp := tmp + " " + strconv.Itoa(layer_number)
        splitf = append(splitf, ltmp)
    }
    var ldf *os.File
    ldf, err = os.OpenFile("layer_deps.txt",os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
    if err != nil {
        fmt.Println(err)
        return
    }
    defer ldf.Close()
    for _,f := range splitf {
        ldf.WriteString(f + "\n")
    }
}
