// THE AUTOGENERATED LICENSE. ALL THE RIGHTS ARE RESERVED BY ROBOTS.

// WARNING: This file has automatically been generated on Tue, 22 Mar 2016 15:19:53 MSK.
// By http://git.io/cgogen. DO NOT EDIT.

package vorbis

/*
#cgo pkg-config: ogg vorbis
#include "ogg/ogg.h"
#include "vorbis/codec.h"
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"

const (
	// OvFalse as defined in vorbis/codec.h:222
	OvFalse = -1
	// OvEof as defined in vorbis/codec.h:223
	OvEof = -2
	// OvHole as defined in vorbis/codec.h:224
	OvHole = -3
	// OvEread as defined in vorbis/codec.h:226
	OvEread = -128
	// OvEfault as defined in vorbis/codec.h:227
	OvEfault = -129
	// OvEimpl as defined in vorbis/codec.h:228
	OvEimpl = -130
	// OvEinval as defined in vorbis/codec.h:229
	OvEinval = -131
	// OvEnotvorbis as defined in vorbis/codec.h:230
	OvEnotvorbis = -132
	// OvEbadheader as defined in vorbis/codec.h:231
	OvEbadheader = -133
	// OvEversion as defined in vorbis/codec.h:232
	OvEversion = -134
	// OvEnotaudio as defined in vorbis/codec.h:233
	OvEnotaudio = -135
	// OvEbadpacket as defined in vorbis/codec.h:234
	OvEbadpacket = -136
	// OvEbadlink as defined in vorbis/codec.h:235
	OvEbadlink = -137
	// OvEnoseek as defined in vorbis/codec.h:236
	OvEnoseek = -138
)
