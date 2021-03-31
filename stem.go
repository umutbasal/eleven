package main

import (
	"fmt"
	"log"
	"os"
	"unsafe"
)

/*
#cgo LDFLAGS: -L/opt/local/lib -lstemmer
#cgo CFLAGS: -I/opt/local/include

#include <stdlib.h>
#include <libstemmer.h>
*/
import "C"

type Stemmer struct {
	stemmer *C.struct_sb_stemmer
	lang    string
}

func free(s *Stemmer) {
	if s.stemmer != nil {
		C.sb_stemmer_delete(s.stemmer)
		s.stemmer = nil
	}
}

// Create a new stemmer, ready for use with the specified language
func NewStemmer(language string) (*Stemmer, error) {
	clang := C.CString(language)
	defer C.free(unsafe.Pointer(clang))
	cchar := C.CString("UTF_8")
	defer C.free(unsafe.Pointer(cchar))

	tmp := C.sb_stemmer_new(clang, cchar)
	if tmp == nil {
		return nil, fmt.Errorf("Can't create stemmer")
	}
	stemmer := &Stemmer{
		stemmer: tmp,
		lang:    language,
	}
	return stemmer, nil
}

func (s *Stemmer) Stem(word string) string {
	ptr := unsafe.Pointer(C.CString(word))
	defer C.free(ptr)

	w := (*C.sb_symbol)(ptr)
	res := unsafe.Pointer(C.sb_stemmer_stem(s.stemmer, w, C.int(len(word))))
	size := C.sb_stemmer_length(s.stemmer)

	buf := C.GoBytes(res, size)
	return string(buf)
}

// StemEnglish reduces english words into basic form
func StemEnglish(word string) string {
	s, err := NewStemmer("english")
	if err != nil {
		log.Printf("Error creating stemmer: " + err.Error())
		os.Exit(1)
	}
	return s.Stem(word)
}
