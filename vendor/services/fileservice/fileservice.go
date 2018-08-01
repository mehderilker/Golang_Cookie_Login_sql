package fileservice

import (
 "os"
 "fmt"
 "image"
 "image/jpeg"
    "log"
    "github.com/disintegration/imaging"
    "models"
)
func init(){
 image.RegisterFormat("jpeg","jpeg",jpeg.Decode,jpeg.DecodeConfig)
}
func PhotoAttr(picName string)(int,int ){
  imgFile,err := os.Open("public/pics/"+picName)

  if err != nil {
   fmt.Println(err)
   os.Exit(1)
  }

  defer imgFile.Close()

  imgCfg,_,err := image.DecodeConfig(imgFile)

  if err != nil {
   fmt.Println(err)
  }

  width := imgCfg.Width
  height :=imgCfg.Height


  return width,height

 }

 func NewPhotoAttrSave(opt models.PhotoAttr){

     imgFile,err := imaging.Open("public/pics/"+opt.FileName)

     if err!= nil{
         log.Fatalf("Dosya Açılırken Bir Hata Meydana Geldi : %v",err)
     }

     imgFile =imaging.Resize(imgFile,opt.Width,opt.Height,imaging.Lanczos)
     imgFile = imaging.Blur(imgFile,opt.Blur)
     imgFile = imaging.AdjustBrightness(imgFile,opt.Brightness)


     err = imaging.Save(imgFile,"public/pics/"+opt.FileName)
     if err!= nil{
         log.Fatalf("Kaydederken Bir Problem Var : %v",err)

     }

 }
