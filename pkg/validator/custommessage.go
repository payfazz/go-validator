package validator

type CustomFieldMessage map[string]string

type ValidatorWithFieldMessage struct {
	val      *Validator
	messages CustomFieldMessage
}

func (val *Validator) WithCustomFieldMessage(messages CustomFieldMessage) *ValidatorWithFieldMessage {
	return &ValidatorWithFieldMessage{
		val:      val,
		messages: messages,
	}
}

func (val *ValidatorWithFieldMessage) ValidateStruct(obj interface{}) error {
	return nil
}
