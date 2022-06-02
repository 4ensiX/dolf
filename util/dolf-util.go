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

type layer_sum struct {
    Manifest []string
    Img_id []string
}

func DLtar(tarfile io.ReadCloser) ([]layers,layer_sum){
        tarReader := tar.NewReader(tarfile)

        var l []layers
	    var l_sum layer_sum

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
                        l_sum.Manifest = readManifest(tarReader)
                        continue
                } else if strings.HasSuffix(name, ".json") && !strings.HasPrefix(name, "manifest") {
                        l_sum.Img_id = readImgConfig(tarReader)
                }
        }
	return l,l_sum
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

        var l []manifest // caution! array!

        if err := json.Unmarshal([]byte(jsonfile), &l); err != nil {
                panic(err)
        }

        manf := l[0] // array[0]

        var layers_sum []string
        var tmp string

        for _,layer := range manf.Layers {
                tmp = strings.Replace(layer, "/layer.tar", "", -1)
                layers_sum = append(layers_sum,tmp)
        }
        return layers_sum

}




type imgfs struct {
        Diff_ids []string `json: "diff_ids"`
}

type imgconfig struct {
        Rootfs  imgfs `json: "rootfs"`
}//大文字じゃないと読み込まない


func readImgConfig(tarReader *tar.Reader) ([]string){
        jsonfile, err := ioutil.ReadAll(tarReader)
        if err != nil {
                fmt.Println(err)
                os.Exit(1)
        }

        var c imgconfig

        if err := json.Unmarshal([]byte(jsonfile), &c); err != nil {
                panic(err)
        }

//        for _,layer := range c.Rootfs.Diff_ids {
//                fmt.Println(layer)
//        }
        return c.Rootfs.Diff_ids
}

