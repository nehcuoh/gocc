package main

type Token string

const TK_BEGIN Token = "BEGIN"
const TK_UNKNOWN = "UNKNOWN"
const TK_ID Token = "ID"
const TK_INT_CONST Token = "INT"
const TK_STRING Token = "STRING"
const TK_FLOAT_CONST = "FLOAT"
const TK_INC Token = "++"
const TK_ADD_ASSIGN Token = "+="
const TK_ADD Token = "+"
const TK_DEC Token = "--"
const TK_SUB_ASSIGN Token = "-="
const TK_POINTER Token = "->"
const TK_SUB Token = "-"
const TK_MUL_ASSIGN Token = "*="
const TK_MUL Token = "*"
const TK_DIV_ASSIGN = "/="
const TK_DIV = "/"
const TK_MOD_ASSIGN = "%="
const TK_MOD = "%"

const TK_LSHIFT_ASSIGN = "<<="
const TK_LSHIFT = "<<"
const TK_LESS_EQ = "<="
const TK_LESS = "<"

const TK_RSHIFT_ASSIGN = ">>="
const TK_RSHIFT = ">>"
const TK_GREAT_EQ = ">="
const TK_GREAT = "<"

const TK_UNEQUAL = "!="
const TK_NOT = "!"

const TK_EQUAL = "=="
const TK_ASSIGN = "="

const TK_OR = "||"
const TK_BITOR = "|"
const TK_BITOR_ASSIGN = "|="

const TK_AND = "&&"
const TK_BITAND = "&"
const TK_BITAND_ASSIGN = "&="

const TK_ELLIPSE = ".."
const TK_DOT = "."

const TK_LEFT_BRACE = "{"
const TK_RIGHT_BRACE = "}"

const TK_LEFT_BRACKET = "["
const TK_RIGHT_BRACKET = "]"

const TK_LEFT_PAREN = "("
const TK_RIGHT_PAREN = ")"

const TK_COMMA = ","
const TK_SEMICOLON = ";"
const TK_COMP = "~"
const TK_QUESTION = "?"
const TK_COLON = ":"

const TK_AUTO = "auto"
const TK_EXTERN = "extern"
const TK_REGISTER = "register"
const TK_STATIC = "static"
const TK_TYPEDEF = "typedef"
const TK_CONST = "const"
const TK_VOLATILE = "volatile"
const TK_CHAR = "char"

// ignore signed, unsigned, short, long
const TK_INT = "int"
const TK_INT64 = "int64"
const TK_FLOAT = "float"

// ignore double
const TK_ENUM = "enum"
const TK_STRUCT = "struct"
const TK_UNION = "union"
const TK_VOID = "void"
const TK_BREAK = "break"
const TK_CASE = "case"
const TK_CONTINUE = "continue"
const TK_DEFAULT = "default"
const TK_DO = "do"
const TK_ELSE = "else"
const TK_FOR = "for"
const TK_GOTO = "goto"
const TK_IF = "if"
const TK_RETURN = "return"
const TK_SWITCH = "switch"
const TK_WHILE = "while"
const TK_SIZEOF = "sizeof"

const TK_END Token = "EOF"
