package main

import (
	"github.com/remko/go-mkvparse"
	"time"
)

type handler struct {
}

func (h *handler) HandleMasterBegin(id mkvparse.ElementID, info mkvparse.ElementInfo) (bool, error) {
	println("HandleMasterBegin: ", mkvparse.NameForElementID(id))
	switch id {
	case mkvparse.TagsElement,
		//mkvparse.ClusterElement,
		mkvparse.BlockGroupElement,
		mkvparse.EBMLElement,
		mkvparse.TracksElement,
		mkvparse.InfoElement,
		mkvparse.SeekHeadElement,
		mkvparse.CuesElement:
		return false, nil
	default:
		return true, nil
	}
	return true, nil
}

func (h *handler) HandleMasterEnd(id mkvparse.ElementID, info mkvparse.ElementInfo) (err error) {
	println("HandleMasterEnd: ", mkvparse.NameForElementID(id))
	return
}

func (h *handler) HandleString(id mkvparse.ElementID, data string, info mkvparse.ElementInfo) (err error) {
	println("HandleString")
	println(mkvparse.NameForElementID(id))
	println(data)
	return
}

func (h *handler) HandleInteger(id mkvparse.ElementID, data int64, info mkvparse.ElementInfo) (err error) {
	println("HandleInteger")
	println(mkvparse.NameForElementID(id))
	println(data)
	return
}

func (h *handler) HandleFloat(id mkvparse.ElementID, data float64, info mkvparse.ElementInfo) (err error) {
	println("HandleFloat")
	println(mkvparse.NameForElementID(id))
	println(data)
	return
}

func (h *handler) HandleDate(id mkvparse.ElementID, data time.Time, info mkvparse.ElementInfo) (err error) {
	println("HandleDate")
	println(mkvparse.NameForElementID(id))
	println(data.String())
	return
}

func (h *handler) HandleBinary(id mkvparse.ElementID, data []byte, info mkvparse.ElementInfo) (err error) {
	println("HandleBinary")
	println(mkvparse.NameForElementID(id))
	return
}

func main() {
	err := mkvparse.ParsePath("video.mkv", &handler{})
	if err != nil {
		panic(err)
	}

}
