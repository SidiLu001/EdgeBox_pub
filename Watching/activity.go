package sample

import (
	"github.com/project-flogo/core/activity"
//	"github.com/project-flogo/core/data/metadata"
)

func init() {
	_ = activity.Register(&Activity{}) //activity.Register(&Activity{}, New) to create instances using factory method 'New'
}

var activityMd = activity.ToMetadata(&Output{})

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

	//call neural network here
    ctx.Logger().Debugf("result of picking out a person: %s", "found") //log is also dummy here
	err = nil //set if neural network go wrong
	if err != nil {
		return true, err
	}

	

	output := &Output{Serial: "imgPath"}//should be serial of the record in the database
	err = ctx.SetOutputObject(output)
	if err != nil {
		return true, err
	}

	return true, nil
}
