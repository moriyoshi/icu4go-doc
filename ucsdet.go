package icu4go

type UCharsetDetector struct {}
type UCharsetMatch struct {}

// Creates the new UCharsetDetector instance.
func UCharsetDetectorOpen(arg1 *UErrorCode) (_swig_ret UCharsetDetector)

// Destroys the UCharsetDetector instance.
func (arg1 *UCharsetDetector) Close()

// Sets the bytes to detect the encoding from.
func (arg1 *UCharsetDetector) SetText(arg2 []byte, arg3 *UErrorCode)

// Sets the "hint" encoding that helps it to detect the encoding, e.g. the value of charset parameter that appears in a html's <meta> tag.
func (arg1 *UCharsetDetector) SetDeclaredEncoding(arg2 []byte, arg3 *UErrorCode)

// Performs detection.
func (arg1 *UCharsetDetector) Detect(arg2 *UErrorCode) (_swig_ret *UCharsetMatch)

// Performs detection and returns all the candidates.
func (arg1 *UCharsetDetector) DetectAll(arg2 *UErrorCode) (_swig_ret []*UCharsetMatch)

// Returns all the encodings that can be told by the detector.
func (arg1 *UCharsetDetector) GetAllDetectableCharsets(arg2 *UErrorCode) (_swig_ret X_UEnumeration)

func (arg1 *UCharsetDetector) IsInputFilterEnabled() (_swig_ret bool)
func (arg1 *UCharsetDetector) EnableInputFilter(arg2 bool)

func (arg1 *UCharsetMatch) GetName(arg2 *UErrorCode) (_swig_ret string)
func (arg1 *UCharsetMatch) GetConfidence(arg2 *UErrorCode) (_swig_ret int)
func (arg1 *UCharsetMatch) GetLanguage(arg2 *UErrorCode) (_swig_ret string)
func DeleteUCharsetMatch(arg1 *UCharsetMatch)
