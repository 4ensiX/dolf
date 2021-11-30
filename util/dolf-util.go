package util

import (
    "archive/tar"
    "fmt"
    "io"
    "os"
    "strings"
    "io/ioutil"
    "encoding/json"
)

type layers struct {
	Layer string
	LayerFiles []string 
}

func DLtar(tarfile io.ReadCloser) ([]layers,[]string){
        tarReader := tar.NewReader(tarfile)

        var l []layers
	var layers_sum []string

        for {
                tarHeader, err := tarReader.Next()
                if err == io.EOF {
                        break
                }

                if err != nil {
                        fmt.Println(err)
                        os.Exit(1)
                }

                name := tarHeader.Name

                if tarHeader.Typeflag == tar.TypeDir {// layer Dir
                        l = append(l,layers{strings.Trim(name,"/"),layerTar(tarReader)})
                } else if strings.HasSuffix(name, "manifest.json") {
                        layers_sum = readManifest(tarReader)
                } else {// [sha256].json,repositories
                        continue
                }
        }
	return l,layers_sum
}


func layerTar(tarReader *tar.Reader) ([]string){

        var filelist []string
        for {
                tarHeader, err := tarReader.Next()
                if err == io.EOF {
                        break
                }

                if err != nil {
                        fmt.Println(err)
                        os.Exit(1)
                }
                name := tarHeader.Name

                if tarHeader.Typeflag == tar.TypeDir {// Dir
                        continue
                }else if strings.HasSuffix(name, "layer.tar") {//layer.tar
			layerReader := tar.NewReader(tarReader)
			filelist = layerTar(layerReader)
			return filelist
		}else if strings.HasSuffix(name, "VERSION") || strings.HasSuffix(name, "json") {
			continue
		}else {
	                filelist = append(filelist, "/" + name)
		}
        }
	return filelist
}

type manifest struct {
        Config  string `json: "config"`
        RepoTags []string `json: "repotag"`
        Layers []string `json: "layers"`
}

func readManifest(tarReader *tar.Reader) ([]string){

        jsonfile, err := ioutil.ReadAll(tarReader)
        if err != nil {
                fmt.Println(err)
                os.Exit(1)
        }

        var l []manifest // why array?

        if err := json.Unmarshal([]byte(jsonfile), &l); err != nil {
                panic(err)
        }

        manf := l[0] // WANT:TO modify

        var layers_sum []string
        var tmp string

        for _,layer := range manf.Layers {
                tmp = strings.Trim(layer, "/layers.tar")
                layers_sum = append(layers_sum,tmp)
        }
        return layers_sum

}
