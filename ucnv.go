package icu4go

type UConverterType int

const (
	UCNV_UNSUPPORTED_CONVERTER UConverterType = iota
	UCNV_SBCS
	UCNV_DBCS
	UCNV_MBCS
	UCNV_LATIN_1
	UCNV_UTF8
	UCNV_UTF16_BigEndian
	UCNV_UTF16_LittleEndian
	UCNV_UTF32_BigEndian
	UCNV_UTF32_LittleEndian
	UCNV_EBCDIC_STATEFUL
	UCNV_ISO_2022
	UCNV_LMBCS_1
	UCNV_LMBCS_2
	UCNV_LMBCS_3
	UCNV_LMBCS_4
	UCNV_LMBCS_5
	UCNV_LMBCS_6
	UCNV_LMBCS_8
	UCNV_LMBCS_11
	UCNV_LMBCS_16
	UCNV_LMBCS_17
	UCNV_LMBCS_18
	UCNV_LMBCS_19
	UCNV_LMBCS_LAST
	UCNV_HZ
	UCNV_SCSU
	UCNV_ISCII
	UCNV_US_ASCII
	UCNV_UTF7
	UCNV_BOCU1
	UCNV_UTF16
	UCNV_UTF32
	UCNV_CESU8
	UCNV_IMAP_MAILBOX
	UCNV_COMPOUND_TEXT
	UCNV_NUMBER_OF_SUPPORTED_CONVERTER_TYPES
)

type UConverterPlatform int

const (
	UCNV_UNKNOWN UConverterPlatform = iota
	UCNV_IBM
)

type UConverterCallbackReason int

const (
	UCNV_UNASSIGNED = iota
	UCNV_ILLEGAL
	UCNV_IRREGULAR
	UCNV_RESET
	UCNV_CLOSE
	UCNV_CLONE
)

const (
	UCNV_OPTION_SEP_CHAR         = ','
	UCNV_OPTION_SEP_STRING       = ","
	UCNV_VALUE_SEP_CHAR          = '='
	UCNV_VALUE_SEP_STRING        = "="
	UCNV_LOCALE_OPTION_STRING    = ",locale="
	UCNV_VERSION_OPTION_STRING   = ",version="
	UCNV_SWAP_LFNL_OPTION_STRING = ",swaplfnl"
)

type UConverterUnicodeSet int

const (
	UCNV_ROUNDTRIP_SET              UConverterUnicodeSet = 0
	UCNV_ROUNDTRIP_AND_FALLBACK_SET                      = 1
	UCNV_SET_COUNT                                       = 2
)

type UConverter struct{}

// Clone the converter (thread-safe).
func (arg1 *UConverter) SafeClone(arg2 *UErrorCode) UConverter

// Closes the converter.  All the internally allocated memory will be released.  This must be called after you finished using the converter.
func (arg1 *UConverter) Close()

// Retrieves the substition byte sequence inserted in place of an Unicode character that is inconvertible from Unicode to the codepage.
func (arg1 *UConverter) GetSubstChars(arg2 *UErrorCode) []byte

// Sets the substition byte sequence inserted in place of an Unicode character that is inconvertible from Unicode to the codepage.  The argument must denote a single character in the codepage.  Note that the converter may apply some restrictions as for stateful encodings, for example only a single byte substitution is allowed for some ISO-2022 variants
func (arg1 *UConverter) SetSubstChars(arg2 []byte, arg3 *UErrorCode)

// Sets the substition characters inserted in place of an Unicode character that is inconvertible from Unicode to the codepage.
func (arg1 *UConverter) SetSubstString(arg2 string, arg3 *UErrorCode)

// Retrieves the portion of the input bytes resulting from the last failing codepage-to-Unicode conversion.
func (arg1 *UConverter) GetInvalidChars(arg2 *UErrorCode) []byte

// Retrieves the portion of the input string resulting from the last failing Unicode-to-codepage conversion.
func (arg1 *UConverter) GetInvalidUChars(arg2 *UErrorCode) string

// Resets the converter's internal state altogether.
func (arg1 *UConverter) Reset()

// Resets the internal state of codepage-to-Unicode part of the converter.
func (arg1 *UConverter) ResetToUnicode()

// Resets the internal state of Unicode-to-codepage part of the converter.
func (arg1 *UConverter) ResetFromUnicode()

// Retrieves the maximum number of bytes that can be yielded from a Unicode character.
func (arg1 *UConverter) GetMaxCharSize() int

// Retrieves the minimum number of bytes that can be yielded from a Unicode character.
func (arg1 *UConverter) GetMinCharSize() int

// Retrieves the name of the converter
func (arg1 *UConverter) GetName(arg2 *UErrorCode) string

// Retrieves the IBM codepage number of the converter.  IBM codepage numbers differ from other vendors' (such as Microsoft's) codepage numbers in some ways.
func (arg1 *UConverter) GetCCSID(arg2 *UErrorCode) int

// Retrieves the type of the converter.
func (arg1 *UConverter) GetType() UConverterType

// Retrieves the starter bytes of the multibyte character set.  For example, an UTF-8's multibyte sequence starts with a byte in range of 0xc0-0xfd.
func (arg1 *UConverter) GetStarters(arg2 *UErrorCode) []int

// Retrieves the whole set of the Unicode characters convertible from / to the codepage.
func (arg1 *UConverter) GetUnicodeSet(arg2 UConverterUnicodeSet, arg3 *UErrorCode) USet

// Sets the callback function that will be called on the various events during the codepage-to-Unicode conversion.
func (arg1 *UConverter) SetToUCallBack(arg2 func(bool, UConverter, []byte, []uint16, []int32, UConverterCallbackReason, *UErrorCode), arg3 *UErrorCode)

// Sets the callback function that will be called on the various events during the Unicode-to-codepage conversion.
func (arg1 *UConverter) SetFromUCallBack(arg2 func(bool, UConverter, []uint16, []byte, []int32, UConverterCallbackReason, *UErrorCode), arg3 *UErrorCode)

// Performs incremental conversion from Unicode to the codepage.  Given slices are modified reflecting the progress.  Offsets needs to be at least the length of the resulting UTF-16 string.  It can also be nil.
func (arg1 *UConverter) FromUnicode(in *[]byte, out *[]uint16, offsets *[]int32, arg5 bool, arg6 *UErrorCode)

// Performs incremental conversion from Unicode to the codepage.  Given slices are modified reflecting the progress.  Offsets needs to be at least the length of the resulting UTF-16 string.  It can also be nil.
func (arg1 *UConverter) ToUnicode(in *[]uint16, out *[]byte, offsets *[]int32, arg5 bool, arg6 *UErrorCode)

// Performs codepage-to-Unicode conversion. Returns a UTF-16 uint16 slice.
func (arg1 *UConverter) ToUChars(arg2 []byte, arg3 *UErrorCode) []uint16

// Performs codepage-to-Unicode conversion. Returns a UTF-8 string.
func (arg1 *UConverter) ToUTF8(arg2 []byte, arg3 *UErrorCode) string

// Replaces the occurrences of the character that turns into \x5c in the Unicode-to-codepage conversion by backslashes (U+005C).  For example, yen signs (U+00A5) will be converted to backslashes for Shift_JIS to Unicode converter.
func (arg1 *UConverter) FixFileSeparator(arg2 string) string

// Returns true if the convertor uses ambiguous mappings of the same character.
func (arg1 *UConverter) IsAmbiguous() (_swig_ret UBool)

// Giving true to the argument instructs the converter to use the fallback mapping.
func (arg1 *UConverter) SetFallback(arg2 UBool)

// Returns true if the converter uses the fallback mapping.
func (arg1 *UConverter) UsesFallback() (_swig_ret UBool)

// Returns the number of UChars held in the converter's internal buffer to await more input for completing the conversion.
func (arg1 *UConverter) FromUCountPending(arg2 *UErrorCode) (_swig_ret int)

// Returns the number of bytes held in the converter's internal buffer to await more input for completing the conversion.
func (arg1 *UConverter) ToUCountPending(arg2 *UErrorCode) (_swig_ret int)

// Returns true if the target codepage is a fixed-width encoding.
func (arg1 *UConverter) IsFixedWidth(arg2 *UErrorCode) (_swig_ret UBool)

// Fluses the internal cache that is shared amongst the converters.
func UConverterFlushCache() (_swig_ret int)

// Detects the Unicode encoding according to the BOM sequence.
func UConverterDetectUnicodeSignature(arg1 []byte, arg2 *int, arg3 *UErrorCode) (_swig_ret string)

// Returns all the available converters.
func UConverterGetAvailableNames() (_swig_ret []string)

// Returns an UEnumeration instance that enumerates all the available converters.
func UConverterOpenAllNames(arg1 *UErrorCode) (_swig_ret X_UEnumeration)

// Returns the aliases for the specified encoding.
func UConverterGetAliases(arg1 string, arg2 *UErrorCode) (_swig_ret []string)

// Returns the UEnumeration instance that enumerates the aliases of the encoding that conform to the specified standard.
func UConverterOpenStandardNames(arg1 string, arg2 string, arg3 *UErrorCode) (_swig_ret X_UEnumeration)

// Returns the name of the encoding that is canonical to the specified standard.
func UConverterGetCanonicalName(arg2 string, arg3 string, arg4 *UErrorCode) (_swig_ret string)
