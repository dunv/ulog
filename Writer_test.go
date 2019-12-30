package ulog

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/dunv/uhelpers"
)

func TestWriterMultipleBuffers(t *testing.T) {

	defaultBuffer := bytes.NewBuffer([]byte{})
	addedBuffer := bytes.NewBuffer([]byte{})

	SetWriter(defaultBuffer, nil)
	SetWriter(addedBuffer, uhelpers.PtrToString("added"))

	Info("halloWelt")
	checkDisregardPackage(defaultBuffer, _LEVEL_INFO, "halloWelt", t)
	checkDisregardPackage(addedBuffer, _LEVEL_INFO, "halloWelt", t)
}

func TestWriterRemoveBuffer(t *testing.T) {

	defaultBuffer := bytes.NewBuffer([]byte{})
	removedBuffer := bytes.NewBuffer([]byte{})

	SetWriter(defaultBuffer, nil)
	SetWriter(removedBuffer, uhelpers.PtrToString("removed"))
	err := RemoveWriter("removed")
	if err != nil {
		t.Error(err)
	}

	Info("halloWelt")
	checkDisregardPackage(defaultBuffer, _LEVEL_INFO, "halloWelt", t)
	res, err := ioutil.ReadAll(removedBuffer)
	if err != nil {
		t.Error(err)
	}
	if string(res) != "" {
		t.Errorf("did write to removed buffer")
	}
}
