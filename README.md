# Package bytesconv

	import "https://github.com/bytesconv/bytesconv"

## Overview

Package bytesconv implements conversions to and from bytes representations of basic data types.

## Numeric Conversions

The most common numeric conversions are Bytestoi ([]byte to int) and ItoBytes (int to []byte).

	i, err := bytesconv.Bytestoi([]byte("-42"))
	bytes := bytesconv.Itobytes(-42)
	