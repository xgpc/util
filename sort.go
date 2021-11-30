package util

// =============================== uint64_int/正序 ===============================
type SortAscUint64Int struct {
	I uint32 // id
	V uint32 // value
}

type SortAscUint64IntArr []SortAscUint64Int

func (s SortAscUint64IntArr) Len() int {
	return len(s)
}

func (s SortAscUint64IntArr) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s SortAscUint64IntArr) Less(i, j int) bool {
	return s[i].V < s[j].V
}

// =============================== uint64_int/倒序 ===============================
type SortDescUint64Int struct {
	I uint64 // id
	V int    // value
}

type SortDescUint64IntArr []SortDescUint64Int

func (s SortDescUint64IntArr) Len() int {
	return len(s)
}

func (s SortDescUint64IntArr) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s SortDescUint64IntArr) Less(i, j int) bool {
	return s[i].V > s[j].V
}

// =============================== uint64/正序 ===============================
type SortAscUint []uint64

func (s SortAscUint) Len() int {
	return len(s)
}

func (s SortAscUint) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s SortAscUint) Less(i, j int) bool {
	return s[i] < s[j]
}

// =============================== uint64/倒序 ===============================
type SortDescUint []uint64

func (s SortDescUint) Len() int {
	return len(s)
}

func (s SortDescUint) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s SortDescUint) Less(i, j int) bool {
	return s[i] > s[j]
}

// =============================== int/正序 ===============================
type SortAscInt []int

func (s SortAscInt) Len() int {
	return len(s)
}

func (s SortAscInt) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s SortAscInt) Less(i, j int) bool {
	return s[i] < s[j]
}

// =============================== int/倒序 ===============================
type SortDescInt []int

func (s SortDescInt) Len() int {
	return len(s)
}

func (s SortDescInt) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s SortDescInt) Less(i, j int) bool {
	return s[i] > s[j]
}

// =============================== string_int/正序 ===============================
type SortAscStringInt struct {
	I string // id
	V int    // value
}

type SortAscStringIntArr []SortAscStringInt

func (s SortAscStringIntArr) Len() int {
	return len(s)
}

func (s SortAscStringIntArr) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s SortAscStringIntArr) Less(i, j int) bool {
	return s[i].V < s[j].V
}

// =============================== string_int/倒序 ===============================
type SortDescStringInt struct {
	I string // id
	V int    // value
}

type SortDescStringIntArr []SortDescStringInt

func (s SortDescStringIntArr) Len() int {
	return len(s)
}

func (s SortDescStringIntArr) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s SortDescStringIntArr) Less(i, j int) bool {
	return s[i].V > s[j].V
}

// =============================== uint64_string/正序 ===============================
type SortAscUint64String struct {
	I uint64 // id
	V string // value
}

type SortAscUint64StringArr []SortAscUint64String

func (s SortAscUint64StringArr) Len() int {
	return len(s)
}

func (s SortAscUint64StringArr) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s SortAscUint64StringArr) Less(i, j int) bool {
	return s[i].V < s[j].V
}

// =============================== uint64_string/倒序 ===============================
type SortDescUint64String struct {
	I uint64 // id
	V string // value
}

type SortDescUint64StringArr []SortDescUint64String

func (s SortDescUint64StringArr) Len() int {
	return len(s)
}

func (s SortDescUint64StringArr) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s SortDescUint64StringArr) Less(i, j int) bool {
	return s[i].V > s[j].V
}

// =============================== string_string/正序 ===============================
type SortAscStringString struct {
	I string // id
	V string // value
}

type SortAscStringStringArr []SortAscStringString

func (s SortAscStringStringArr) Len() int {
	return len(s)
}

func (s SortAscStringStringArr) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s SortAscStringStringArr) Less(i, j int) bool {
	return s[i].V < s[j].V
}

// =============================== string_string/倒序 ===============================
type SortDescStringString struct {
	I string // id
	V string // value
}

type SortDescStringStringArr []SortDescStringString

func (s SortDescStringStringArr) Len() int {
	return len(s)
}

func (s SortDescStringStringArr) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s SortDescStringStringArr) Less(i, j int) bool {
	return s[i].V > s[j].V
}

// =============================== uint64_int64/正序 ===============================
type SortAscUint64Int64 struct {
	I uint64 // id
	V int64  // value
}

type SortAscUint64Int64Arr []SortAscUint64Int64

func (s SortAscUint64Int64Arr) Len() int {
	return len(s)
}

func (s SortAscUint64Int64Arr) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s SortAscUint64Int64Arr) Less(i, j int) bool {
	return s[i].V < s[j].V
}
