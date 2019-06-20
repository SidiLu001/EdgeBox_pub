package sample

import (
	"log"

	"github.com/project-flogo/core/activity"
	//	"github.com/project-flogo/core/data/metadata"

	"fmt"
	"os"

	// "image/color"
	// "image"
	// "log"
	"strconv"

	"github.com/Kagami/go-face"
	"gocv.io/x/gocv"
)

// const dataDir = "testdata"

// var rec, _ = face.NewRecognizer("testdata")
var window = gocv.NewWindow("EdgeWare")
var img gocv.Mat
var webcam, _ = gocv.OpenVideoCapture("resource/the_car_lab.mp4")
var rec, _ = face.NewRecognizer("resource/faceModel")

// var boxcolor color.RGBA
var frameIndex = 0
var deviceID string
var filename string
var activityMd = activity.ToMetadata(&Output{})
var err error
var imgDir = os.Getenv("HOME") + "/flogo"

func init() {
	_ = activity.Register(&Activity{}) //activity.Register(&Activity{}, New) to create instances using factory method 'New'
	// window = gocv.NewWindow("Flogo")
	// defer window.Close()
	// frameIndex = 0
	// img = gocv.NewMat()
	// defer img.Close()

	// // Init the recognizer.
	// rec, err = face.NewRecognizer(dataDir)
	// if err != nil {
	// 	log.Fatalf("Can't init face recognizer: %v", err)
	// }
	// // Free the resources when you're finished.
	// defer rec.Close()

	//*****************************************
	// deviceID = "the_car_lab.mp4"
	// // open capture device
	// webcam, err = gocv.OpenVideoCapture(deviceID)
	// if err != nil {
	// 	fmt.Printf("Error opening video capture device: %v\n", deviceID)
	// 	return
	// }

	// defer webcam.Close()

	// boxcolor = color.RGBA{0, 255, 0, 0}
}

//New optional factory method, should be used if one activity instance per configuration is desired
func New(ctx activity.InitContext) (activity.Activity, error) {

	//	s := &Settings{}
	//	err := metadata.MapToStruct(ctx.Settings(), s, true)
	//	if err != nil {
	//		return nil, err
	//	}

	//	ctx.Logger().Debugf("Setting: %s", s.ASetting)

	act := &Activity{} //add aSetting to instance//nothing to add now

	return act, nil
}

// Activity is an sample Activity that can be used as a base to create a custom activity
type Activity struct {
}

// Metadata returns the activity's metadata
func (a *Activity) Metadata() *activity.Metadata {
	return activityMd
}

// Eval implements api.Activity.Eval - Logs the Message
func (a *Activity) Eval(ctx activity.Context) (done bool, err error) {
	img = gocv.NewMat()
	defer img.Close()

	// deviceID = "the_car_lab.mp4"
	// open capture device
	// webcam, err = gocv.OpenVideoCapture(deviceID)
	// if err != nil {
	// 	fmt.Printf("Error opening video capture device: %v\n", deviceID)
	// 	return
	// }
	// defer webcam.Close()

	//call neural network here
	ctx.Logger().Debugf("result of picking out a person: %s", "found") //log is also dummy here
	err = nil                                                          //set if neural network go wrong
	if err != nil {
		return true, err
	}

	// *************************************
	// add by Yongtao
	if ok := webcam.Read(&img); !ok {
		fmt.Printf("Device closed: %v\n", deviceID)
		return
	}
	fmt.Println(img.Size())
	frameIndex++
	filename = imgDir + "/flogo" + strconv.Itoa(frameIndex) + ".jpg"

	window.IMShow(img)
	window.WaitKey(1)

	if frameIndex%30 == 0 {
		return true, nil
	}

	testImagePristin := "resource/temp/tmp.jpg"
	gocv.IMWrite(testImagePristin, img)

	// Recognize faces on that image.
	faces, err := rec.RecognizeFile(testImagePristin)
	if err != nil {
		log.Fatalf("Can't recognize: %v", err)
	}

	fmt.Printf("\n %c[%d;%d;%dm# of faces: %d%c[0m\n", 0x1B, 0, 0, 33, len(faces), 0x1B)
	// imgFace := gocv.IMRead(testImagePristin, gocv.IMReadColor)

	save := false
	// if save is true, indicating that the face is detected

	sendString := filename
	for _, f := range faces {
		mRect := f.Rectangle
		fmt.Println(mRect)
		// 	mRect.Min.X -= 20
		// 	mRect.Min.Y -= 60
		// 	mRect.Max.X += 20
		// 	mRect.Max.Y += 20
		// 	fmt.Println(mRect.Min.X, mRect.Min.Y, mRect.Max.X, mRect.Max.Y)
		// 	// gocv.Rectangle(&img, mRect, color.RGBA{0, 255, 0, 0}, 2)
		save = true
		// 	// rect := image.Rect(mRect.Min.X, mRect.Min.Y, mRect.Max.X, mRect.Max.Y)
		// 	// imgFace := img.Region(rect)

		// // 	frameIndex++
		// // 	filename = "/home/yyt/flogo/flogo" + strconv.Itoa(frameIndex) + ".jpg"
		// // 	gocv.IMWrite(filename, imgFace)
		sendString += ";" + mRect.String()
	}
	// *************************

	// if !save {
	// 	return false, nil
	// }
	// ***********************
	// filename = testImagePristin
	//todo:
	// A frame of pictures may contain multiple faces, which will be stored as multiple files.
	// These file paths should be merged and transmitted in strings.
	// Now each picture only transmitted a face's path for testing

	//
	if save {
		gocv.IMWrite(filename, img)
		output := &Output{Serial: sendString} //should be serial of the record in the database
		err = ctx.SetOutputObject(output)
		if err != nil {
			return true, err
		}
	}

	return true, nil
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
