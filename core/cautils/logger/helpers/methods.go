package helpers

type IDetails interface {
	Key() string
	Value() interface{}
}

// ======================================================================================
// ============================== String ================================================
// ======================================================================================

// Key
func (s *StringObj) Key() string {
	return s.key
}

// Value
func (s *StringObj) Value() interface{} {
	return s.value
}

// ======================================================================================
// =============================== Error ================================================
// ======================================================================================

// Key
func (s *ErrorObj) Key() string {
	return s.key
}

// Value
func (s *ErrorObj) Value() interface{} {
	return s.value
}

// ======================================================================================
// ================================= Int ================================================
// ======================================================================================

// Key
func (s *IntObj) Key() string {
	return s.key
}

// Value
func (s *IntObj) Value() interface{} {
	return s.value
}

// ======================================================================================
// =========================== Interface ================================================
// ======================================================================================

// Key
func (s *InterfaceObj) Key() string {
	return s.key
}

// Value
func (s *InterfaceObj) Value() interface{} {
	return s.value
}
