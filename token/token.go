package token

type TokenType string

type Token struct {
    Type    TokenType
    Literal string
}

const (
    Illegal      = "Illegal"
    EOF          = "EOF"

    Identifier   = "Identifier"
    Int          = "Int"
    String       = "String"

    Assign       = "="
    Plus         = "+"
    Minus        = "-"
    Bang         = "!"
    Asterisk     = "*"
    Slash        = "/"
    Equal        = "=="
    NotEqual     = "!="

    LessThan     = "<"
    GreaterThan  = ">"

    Comma        = ","
    Semicolon    = ";"
    Colon        = ":"

    LeftParen    = "("
    RightParen   = ")"
    LeftBrace    = "{"
    RightBrace   = "}"
    LeftBracket  = "["
    RightBracket = "]"

    Function     = "Function"
    Let          = "Let"
    True         = "True"
    False        = "False"
    If           = "If"
    Else         = "Else"
    Return       = "Return"
)

var keywords = map[string]TokenType{
    "fn":     Function,
    "let":    Let,
    "true":   True,
    "false":  False,
    "if":     If,
    "else":   Else,
    "return": Return,
}

func LookupIdentifierType(identifier string) TokenType {
    if tok, ok := keywords[identifier]; ok {
        return tok
    }

    return Identifier
}
