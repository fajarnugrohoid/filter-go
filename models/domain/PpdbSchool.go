package domain

type PpdbSchool struct {
	Type  string `bson:"type,omitempty"`
	Level string `bson:"level,omitempty"`
	code  int    `bson:"code,omitempty"`
}
