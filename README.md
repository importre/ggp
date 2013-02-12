ggp (Go Get Packages)
===
- - -

Search for Go packages.

## Installation

`$ go get github.com/importre/ggp`

## Usage

`$ ggp -q <query>`

## Example

If you want to find packages related to "xml", just type the following.

`$ ggp -q xml`

Then **ggp** will show you something.

    package  description
    =============================
    * encoding/xml
        - Package xml implements a simple XML 1.0 parser that understands XML name spaces.

    * github.com/mkrautz/plist/xmlplist
        - Package xmlplist decodes XML plist files

    * github.com/mewmew/tmx
        - Package tmx implements access to Tiled's tmx (Tile Map XML) files.

    * github.com/mewkiz84/tmx
        - Package tmx implements access to Tiled's tmx (Tile Map XML) files.
      
      ...
      omitted
      ...
  
