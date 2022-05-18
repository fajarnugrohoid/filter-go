package models

type PpdbOptionList struct {
	Options []*PpdbOption
}

func (optionList *PpdbOptionList) AddOpt(item PpdbOption) {
	optionList.Options = append(optionList.Options, &item)
}
