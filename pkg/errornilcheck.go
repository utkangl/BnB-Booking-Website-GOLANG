package pkg

import "log"

func ErrorNilCheckPrint(err error) {
	if err != nil {
		log.Println(err)
	}
}

func ErrorNilCheckFatal(err error) {
	if err != nil {
		log.Println(err)
	}
}

func ErrorNilCheckReturn(err error) error {
	return err
}
