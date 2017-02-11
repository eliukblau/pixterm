// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sfnt

/*
This file contains opt-in tests for popular, high quality, proprietary fonts,
made by companies such as Adobe and Microsoft. These fonts are generally
available, but copies are not explicitly included in this repository due to
licensing differences or file size concerns. To opt-in, run:

go test golang.org/x/image/font/sfnt -args -proprietary

Not all tests pass out-of-the-box on all systems. For example, the Microsoft
Times New Roman font is downloadable gratis even on non-Windows systems, but as
per the ttf-mscorefonts-installer Debian package, this requires accepting an
End User License Agreement (EULA) and a CAB format decoder. These tests assume
that such fonts have already been installed. You may need to specify the
directories for these fonts:

go test golang.org/x/image/font/sfnt -args -proprietary -adobeDir=/foo/bar/aFonts -microsoftDir=/foo/bar/mFonts

To only run those tests for the Microsoft fonts:

go test golang.org/x/image/font/sfnt -test.run=ProprietaryMicrosoft -args -proprietary
*/

// TODO: add Apple system fonts? Google fonts (Droid? Noto?)? Emoji fonts?

// TODO: enable Apple/Microsoft tests by default on Darwin/Windows?

import (
	"flag"
	"io/ioutil"
	"path/filepath"
	"testing"
)

var (
	proprietary = flag.Bool("proprietary", false, "test proprietary fonts not included in this repository")

	adobeDir = flag.String(
		"adobeDir",
		// This needs to be set explicitly. There is no default dir on Debian:
		// https://bugs.debian.org/cgi-bin/bugreport.cgi?bug=736680
		//
		// Get the fonts from https://github.com/adobe-fonts, e.g.:
		//	- https://github.com/adobe-fonts/source-code-pro/releases/latest
		//	- https://github.com/adobe-fonts/source-han-sans/releases/latest
		//	- https://github.com/adobe-fonts/source-sans-pro/releases/latest
		//
		// Copy all of the TTF and OTF files to the one directory, such as
		// $HOME/adobe-fonts, and pass that as the -adobeDir flag here.
		"",
		"directory name for the Adobe proprietary fonts",
	)

	microsoftDir = flag.String(
		"microsoftDir",
		"/usr/share/fonts/truetype/msttcorefonts",
		"directory name for the Microsoft proprietary fonts",
	)
)

type proprietor int

const (
	adobe proprietor = iota
	microsoft
)

func TestProprietaryAdobeSourceCodeProOTF(t *testing.T) {
	testProprietary(t, adobe, "SourceCodePro-Regular.otf", 1500, 2)
}

func TestProprietaryAdobeSourceCodeProTTF(t *testing.T) {
	testProprietary(t, adobe, "SourceCodePro-Regular.ttf", 1500, 36)
}

func TestProprietaryAdobeSourceHanSansSC(t *testing.T) {
	testProprietary(t, adobe, "SourceHanSansSC-Regular.otf", 65535, 2)
}

func TestProprietaryAdobeSourceSansProOTF(t *testing.T) {
	testProprietary(t, adobe, "SourceSansPro-Regular.otf", 1800, 2)
}

func TestProprietaryAdobeSourceSansProTTF(t *testing.T) {
	// The 1000 here is smaller than the 1800 above. For some reason, the TTF
	// version of the file has fewer glyphs than the (presumably canonical) OTF
	// version. The number of glyphs in the .otf and .ttf files can be verified
	// with the ttx tool.
	testProprietary(t, adobe, "SourceSansPro-Regular.ttf", 1000, 56)
}

func TestProprietaryMicrosoftArial(t *testing.T) {
	testProprietary(t, microsoft, "Arial.ttf", 1200, 98)
}

func TestProprietaryMicrosoftComicSansMS(t *testing.T) {
	testProprietary(t, microsoft, "Comic_Sans_MS.ttf", 550, 98)
}

func TestProprietaryMicrosoftTimesNewRoman(t *testing.T) {
	testProprietary(t, microsoft, "Times_New_Roman.ttf", 1200, 98)
}

func TestProprietaryMicrosoftWebdings(t *testing.T) {
	testProprietary(t, microsoft, "Webdings.ttf", 200, -1)
}

// testProprietary tests that we can load every glyph in the named font.
//
// The exact number of glyphs in the font can differ across its various
// versions, but as a sanity check, there should be at least minNumGlyphs.
//
// While this package is a work-in-progress, not every glyph can be loaded. The
// firstUnsupportedGlyph argument, if non-negative, is the index of the first
// unsupported glyph in the font. This number should increase over time (or set
// negative), as the TODO's in this package are done.
func testProprietary(t *testing.T, p proprietor, filename string, minNumGlyphs, firstUnsupportedGlyph int) {
	if !*proprietary {
		t.Skip("skipping proprietary font test")
	}

	file, err := []byte(nil), error(nil)
	switch p {
	case adobe:
		file, err = ioutil.ReadFile(filepath.Join(*adobeDir, filename))
		if err != nil {
			t.Fatalf("%v\nPerhaps you need to set the -adobeDir=%v flag?", err, *adobeDir)
		}
	case microsoft:
		file, err = ioutil.ReadFile(filepath.Join(*microsoftDir, filename))
		if err != nil {
			t.Fatalf("%v\nPerhaps you need to set the -microsoftDir=%v flag?", err, *microsoftDir)
		}
	}
	f, err := Parse(file)
	if err != nil {
		t.Fatalf("Parse: %v", err)
	}

	numGlyphs := f.NumGlyphs()
	if numGlyphs < minNumGlyphs {
		t.Fatalf("NumGlyphs: got %d, want at least %d", numGlyphs, minNumGlyphs)
	}

	var buf Buffer
	iMax := numGlyphs
	if firstUnsupportedGlyph >= 0 {
		iMax = firstUnsupportedGlyph
	}
	for i, numErrors := 0, 0; i < iMax; i++ {
		if _, err := f.LoadGlyph(&buf, GlyphIndex(i), nil); err != nil {
			t.Errorf("LoadGlyph(%d): %v", i, err)
			numErrors++
		}
		if numErrors == 10 {
			t.Fatal("LoadGlyph: too many errors")
		}
	}
}
