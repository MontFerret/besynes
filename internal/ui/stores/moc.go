package stores

//#include <stdint.h>
//#include <stdlib.h>
//#include <string.h>
//#include "moc.h"
import "C"
import (
	"runtime"
	"strings"
	"unsafe"

	"github.com/therecipe/qt"
	std_core "github.com/therecipe/qt/core"
)

func cGoUnpackString(s C.struct_Moc_PackedString) string {
	if int(s.len) == -1 {
		return C.GoString(s.data)
	}
	return C.GoStringN(s.data, C.int(s.len))
}
func cGoUnpackBytes(s C.struct_Moc_PackedString) []byte {
	if int(s.len) == -1 {
		gs := C.GoString(s.data)
		return *(*[]byte)(unsafe.Pointer(&gs))
	}
	return C.GoBytes(unsafe.Pointer(s.data), C.int(s.len))
}
func unpackStringList(s string) []string {
	if len(s) == 0 {
		return make([]string, 0)
	}
	return strings.Split(s, "¡¦!")
}

type ExecutionStore_ITF interface {
	std_core.QObject_ITF
	ExecutionStore_PTR() *ExecutionStore
}

func (ptr *ExecutionStore) ExecutionStore_PTR() *ExecutionStore {
	return ptr
}

func (ptr *ExecutionStore) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *ExecutionStore) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromExecutionStore(ptr ExecutionStore_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.ExecutionStore_PTR().Pointer()
	}
	return nil
}

func NewExecutionStoreFromPointer(ptr unsafe.Pointer) (n *ExecutionStore) {
	if gPtr, ok := qt.Receive(ptr); !ok {
		n = new(ExecutionStore)
		n.SetPointer(ptr)
	} else {
		switch deduced := gPtr.(type) {
		case *ExecutionStore:
			n = deduced

		case *std_core.QObject:
			n = &ExecutionStore{QObject: *deduced}

		default:
			n = new(ExecutionStore)
			n.SetPointer(ptr)
		}
	}
	return
}

//export callbackExecutionStoref82535_Constructor
func callbackExecutionStoref82535_Constructor(ptr unsafe.Pointer) {
	this := NewExecutionStoreFromPointer(ptr)
	qt.Register(ptr, this)
}

//export callbackExecutionStoref82535_Execute
func callbackExecutionStoref82535_Execute(ptr unsafe.Pointer, addr C.struct_Moc_PackedString) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "execute"); signal != nil {
		return PointerFromResultViewModel((*(*func(string) *ResultViewModel)(signal))(cGoUnpackString(addr)))
	}

	return PointerFromResultViewModel(NewResultViewModel(nil))
}

func (ptr *ExecutionStore) ConnectExecute(f func(addr string) *ResultViewModel) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "execute"); signal != nil {
			f := func(addr string) *ResultViewModel {
				(*(*func(string) *ResultViewModel)(signal))(addr)
				return f(addr)
			}
			qt.ConnectSignal(ptr.Pointer(), "execute", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "execute", unsafe.Pointer(&f))
		}
	}
}

func (ptr *ExecutionStore) DisconnectExecute() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "execute")
	}
}

func (ptr *ExecutionStore) Execute(addr string) *ResultViewModel {
	if ptr.Pointer() != nil {
		var addrC *C.char
		if addr != "" {
			addrC = C.CString(addr)
			defer C.free(unsafe.Pointer(addrC))
		}
		tmpValue := NewResultViewModelFromPointer(C.ExecutionStoref82535_Execute(ptr.Pointer(), C.struct_Moc_PackedString{data: addrC, len: C.longlong(len(addr))}))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func ExecutionStore_QRegisterMetaType() int {
	return int(int32(C.ExecutionStoref82535_ExecutionStoref82535_QRegisterMetaType()))
}

func (ptr *ExecutionStore) QRegisterMetaType() int {
	return int(int32(C.ExecutionStoref82535_ExecutionStoref82535_QRegisterMetaType()))
}

func ExecutionStore_QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.ExecutionStoref82535_ExecutionStoref82535_QRegisterMetaType2(typeNameC)))
}

func (ptr *ExecutionStore) QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.ExecutionStoref82535_ExecutionStoref82535_QRegisterMetaType2(typeNameC)))
}

func ExecutionStore_QmlRegisterType() int {
	return int(int32(C.ExecutionStoref82535_ExecutionStoref82535_QmlRegisterType()))
}

func (ptr *ExecutionStore) QmlRegisterType() int {
	return int(int32(C.ExecutionStoref82535_ExecutionStoref82535_QmlRegisterType()))
}

func ExecutionStore_QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	return int(int32(C.ExecutionStoref82535_ExecutionStoref82535_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *ExecutionStore) QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	return int(int32(C.ExecutionStoref82535_ExecutionStoref82535_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *ExecutionStore) __children_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.ExecutionStoref82535___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *ExecutionStore) __children_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.ExecutionStoref82535___children_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *ExecutionStore) __children_newList() unsafe.Pointer {
	return C.ExecutionStoref82535___children_newList(ptr.Pointer())
}

func (ptr *ExecutionStore) __dynamicPropertyNames_atList(i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.ExecutionStoref82535___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *ExecutionStore) __dynamicPropertyNames_setList(i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.ExecutionStoref82535___dynamicPropertyNames_setList(ptr.Pointer(), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *ExecutionStore) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.ExecutionStoref82535___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *ExecutionStore) __findChildren_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.ExecutionStoref82535___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *ExecutionStore) __findChildren_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.ExecutionStoref82535___findChildren_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *ExecutionStore) __findChildren_newList() unsafe.Pointer {
	return C.ExecutionStoref82535___findChildren_newList(ptr.Pointer())
}

func (ptr *ExecutionStore) __findChildren_atList3(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.ExecutionStoref82535___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *ExecutionStore) __findChildren_setList3(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.ExecutionStoref82535___findChildren_setList3(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *ExecutionStore) __findChildren_newList3() unsafe.Pointer {
	return C.ExecutionStoref82535___findChildren_newList3(ptr.Pointer())
}

func (ptr *ExecutionStore) __qFindChildren_atList2(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.ExecutionStoref82535___qFindChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *ExecutionStore) __qFindChildren_setList2(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.ExecutionStoref82535___qFindChildren_setList2(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *ExecutionStore) __qFindChildren_newList2() unsafe.Pointer {
	return C.ExecutionStoref82535___qFindChildren_newList2(ptr.Pointer())
}

func NewExecutionStore(parent std_core.QObject_ITF) *ExecutionStore {
	tmpValue := NewExecutionStoreFromPointer(C.ExecutionStoref82535_NewExecutionStore(std_core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackExecutionStoref82535_DestroyExecutionStore
func callbackExecutionStoref82535_DestroyExecutionStore(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~ExecutionStore"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewExecutionStoreFromPointer(ptr).DestroyExecutionStoreDefault()
	}
}

func (ptr *ExecutionStore) ConnectDestroyExecutionStore(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~ExecutionStore"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~ExecutionStore", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~ExecutionStore", unsafe.Pointer(&f))
		}
	}
}

func (ptr *ExecutionStore) DisconnectDestroyExecutionStore() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~ExecutionStore")
	}
}

func (ptr *ExecutionStore) DestroyExecutionStore() {
	if ptr.Pointer() != nil {
		C.ExecutionStoref82535_DestroyExecutionStore(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *ExecutionStore) DestroyExecutionStoreDefault() {
	if ptr.Pointer() != nil {
		C.ExecutionStoref82535_DestroyExecutionStoreDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackExecutionStoref82535_ChildEvent
func callbackExecutionStoref82535_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*std_core.QChildEvent))(signal))(std_core.NewQChildEventFromPointer(event))
	} else {
		NewExecutionStoreFromPointer(ptr).ChildEventDefault(std_core.NewQChildEventFromPointer(event))
	}
}

func (ptr *ExecutionStore) ChildEventDefault(event std_core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.ExecutionStoref82535_ChildEventDefault(ptr.Pointer(), std_core.PointerFromQChildEvent(event))
	}
}

//export callbackExecutionStoref82535_ConnectNotify
func callbackExecutionStoref82535_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*std_core.QMetaMethod))(signal))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewExecutionStoreFromPointer(ptr).ConnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *ExecutionStore) ConnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.ExecutionStoref82535_ConnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackExecutionStoref82535_CustomEvent
func callbackExecutionStoref82535_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*std_core.QEvent))(signal))(std_core.NewQEventFromPointer(event))
	} else {
		NewExecutionStoreFromPointer(ptr).CustomEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *ExecutionStore) CustomEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.ExecutionStoref82535_CustomEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackExecutionStoref82535_DeleteLater
func callbackExecutionStoref82535_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewExecutionStoreFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *ExecutionStore) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.ExecutionStoref82535_DeleteLaterDefault(ptr.Pointer())
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackExecutionStoref82535_Destroyed
func callbackExecutionStoref82535_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*std_core.QObject))(signal))(std_core.NewQObjectFromPointer(obj))
	}
	qt.Unregister(ptr)

}

//export callbackExecutionStoref82535_DisconnectNotify
func callbackExecutionStoref82535_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*std_core.QMetaMethod))(signal))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewExecutionStoreFromPointer(ptr).DisconnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *ExecutionStore) DisconnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.ExecutionStoref82535_DisconnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackExecutionStoref82535_Event
func callbackExecutionStoref82535_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QEvent) bool)(signal))(std_core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewExecutionStoreFromPointer(ptr).EventDefault(std_core.NewQEventFromPointer(e)))))
}

func (ptr *ExecutionStore) EventDefault(e std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.ExecutionStoref82535_EventDefault(ptr.Pointer(), std_core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackExecutionStoref82535_EventFilter
func callbackExecutionStoref82535_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QObject, *std_core.QEvent) bool)(signal))(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewExecutionStoreFromPointer(ptr).EventFilterDefault(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
}

func (ptr *ExecutionStore) EventFilterDefault(watched std_core.QObject_ITF, event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.ExecutionStoref82535_EventFilterDefault(ptr.Pointer(), std_core.PointerFromQObject(watched), std_core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackExecutionStoref82535_ObjectNameChanged
func callbackExecutionStoref82535_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackExecutionStoref82535_TimerEvent
func callbackExecutionStoref82535_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*std_core.QTimerEvent))(signal))(std_core.NewQTimerEventFromPointer(event))
	} else {
		NewExecutionStoreFromPointer(ptr).TimerEventDefault(std_core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *ExecutionStore) TimerEventDefault(event std_core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.ExecutionStoref82535_TimerEventDefault(ptr.Pointer(), std_core.PointerFromQTimerEvent(event))
	}
}

type ResultViewModel_ITF interface {
	std_core.QObject_ITF
	ResultViewModel_PTR() *ResultViewModel
}

func (ptr *ResultViewModel) ResultViewModel_PTR() *ResultViewModel {
	return ptr
}

func (ptr *ResultViewModel) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *ResultViewModel) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromResultViewModel(ptr ResultViewModel_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.ResultViewModel_PTR().Pointer()
	}
	return nil
}

func NewResultViewModelFromPointer(ptr unsafe.Pointer) (n *ResultViewModel) {
	if gPtr, ok := qt.Receive(ptr); !ok {
		n = new(ResultViewModel)
		n.SetPointer(ptr)
	} else {
		switch deduced := gPtr.(type) {
		case *ResultViewModel:
			n = deduced

		case *std_core.QObject:
			n = &ResultViewModel{QObject: *deduced}

		default:
			n = new(ResultViewModel)
			n.SetPointer(ptr)
		}
	}
	return
}

//export callbackResultViewModelf82535_Constructor
func callbackResultViewModelf82535_Constructor(ptr unsafe.Pointer) {
	this := NewResultViewModelFromPointer(ptr)
	qt.Register(ptr, this)
}

func ResultViewModel_QRegisterMetaType() int {
	return int(int32(C.ResultViewModelf82535_ResultViewModelf82535_QRegisterMetaType()))
}

func (ptr *ResultViewModel) QRegisterMetaType() int {
	return int(int32(C.ResultViewModelf82535_ResultViewModelf82535_QRegisterMetaType()))
}

func ResultViewModel_QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.ResultViewModelf82535_ResultViewModelf82535_QRegisterMetaType2(typeNameC)))
}

func (ptr *ResultViewModel) QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.ResultViewModelf82535_ResultViewModelf82535_QRegisterMetaType2(typeNameC)))
}

func ResultViewModel_QmlRegisterType() int {
	return int(int32(C.ResultViewModelf82535_ResultViewModelf82535_QmlRegisterType()))
}

func (ptr *ResultViewModel) QmlRegisterType() int {
	return int(int32(C.ResultViewModelf82535_ResultViewModelf82535_QmlRegisterType()))
}

func ResultViewModel_QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	return int(int32(C.ResultViewModelf82535_ResultViewModelf82535_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *ResultViewModel) QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	return int(int32(C.ResultViewModelf82535_ResultViewModelf82535_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *ResultViewModel) __children_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.ResultViewModelf82535___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *ResultViewModel) __children_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.ResultViewModelf82535___children_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *ResultViewModel) __children_newList() unsafe.Pointer {
	return C.ResultViewModelf82535___children_newList(ptr.Pointer())
}

func (ptr *ResultViewModel) __dynamicPropertyNames_atList(i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.ResultViewModelf82535___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *ResultViewModel) __dynamicPropertyNames_setList(i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.ResultViewModelf82535___dynamicPropertyNames_setList(ptr.Pointer(), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *ResultViewModel) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.ResultViewModelf82535___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *ResultViewModel) __findChildren_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.ResultViewModelf82535___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *ResultViewModel) __findChildren_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.ResultViewModelf82535___findChildren_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *ResultViewModel) __findChildren_newList() unsafe.Pointer {
	return C.ResultViewModelf82535___findChildren_newList(ptr.Pointer())
}

func (ptr *ResultViewModel) __findChildren_atList3(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.ResultViewModelf82535___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *ResultViewModel) __findChildren_setList3(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.ResultViewModelf82535___findChildren_setList3(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *ResultViewModel) __findChildren_newList3() unsafe.Pointer {
	return C.ResultViewModelf82535___findChildren_newList3(ptr.Pointer())
}

func (ptr *ResultViewModel) __qFindChildren_atList2(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.ResultViewModelf82535___qFindChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *ResultViewModel) __qFindChildren_setList2(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.ResultViewModelf82535___qFindChildren_setList2(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *ResultViewModel) __qFindChildren_newList2() unsafe.Pointer {
	return C.ResultViewModelf82535___qFindChildren_newList2(ptr.Pointer())
}

func NewResultViewModel(parent std_core.QObject_ITF) *ResultViewModel {
	tmpValue := NewResultViewModelFromPointer(C.ResultViewModelf82535_NewResultViewModel(std_core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackResultViewModelf82535_DestroyResultViewModel
func callbackResultViewModelf82535_DestroyResultViewModel(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~ResultViewModel"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewResultViewModelFromPointer(ptr).DestroyResultViewModelDefault()
	}
}

func (ptr *ResultViewModel) ConnectDestroyResultViewModel(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~ResultViewModel"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~ResultViewModel", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~ResultViewModel", unsafe.Pointer(&f))
		}
	}
}

func (ptr *ResultViewModel) DisconnectDestroyResultViewModel() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~ResultViewModel")
	}
}

func (ptr *ResultViewModel) DestroyResultViewModel() {
	if ptr.Pointer() != nil {
		C.ResultViewModelf82535_DestroyResultViewModel(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *ResultViewModel) DestroyResultViewModelDefault() {
	if ptr.Pointer() != nil {
		C.ResultViewModelf82535_DestroyResultViewModelDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackResultViewModelf82535_ChildEvent
func callbackResultViewModelf82535_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*std_core.QChildEvent))(signal))(std_core.NewQChildEventFromPointer(event))
	} else {
		NewResultViewModelFromPointer(ptr).ChildEventDefault(std_core.NewQChildEventFromPointer(event))
	}
}

func (ptr *ResultViewModel) ChildEventDefault(event std_core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.ResultViewModelf82535_ChildEventDefault(ptr.Pointer(), std_core.PointerFromQChildEvent(event))
	}
}

//export callbackResultViewModelf82535_ConnectNotify
func callbackResultViewModelf82535_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*std_core.QMetaMethod))(signal))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewResultViewModelFromPointer(ptr).ConnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *ResultViewModel) ConnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.ResultViewModelf82535_ConnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackResultViewModelf82535_CustomEvent
func callbackResultViewModelf82535_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*std_core.QEvent))(signal))(std_core.NewQEventFromPointer(event))
	} else {
		NewResultViewModelFromPointer(ptr).CustomEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *ResultViewModel) CustomEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.ResultViewModelf82535_CustomEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackResultViewModelf82535_DeleteLater
func callbackResultViewModelf82535_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewResultViewModelFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *ResultViewModel) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.ResultViewModelf82535_DeleteLaterDefault(ptr.Pointer())
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackResultViewModelf82535_Destroyed
func callbackResultViewModelf82535_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*std_core.QObject))(signal))(std_core.NewQObjectFromPointer(obj))
	}
	qt.Unregister(ptr)

}

//export callbackResultViewModelf82535_DisconnectNotify
func callbackResultViewModelf82535_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*std_core.QMetaMethod))(signal))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewResultViewModelFromPointer(ptr).DisconnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *ResultViewModel) DisconnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.ResultViewModelf82535_DisconnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackResultViewModelf82535_Event
func callbackResultViewModelf82535_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QEvent) bool)(signal))(std_core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewResultViewModelFromPointer(ptr).EventDefault(std_core.NewQEventFromPointer(e)))))
}

func (ptr *ResultViewModel) EventDefault(e std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.ResultViewModelf82535_EventDefault(ptr.Pointer(), std_core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackResultViewModelf82535_EventFilter
func callbackResultViewModelf82535_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QObject, *std_core.QEvent) bool)(signal))(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewResultViewModelFromPointer(ptr).EventFilterDefault(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
}

func (ptr *ResultViewModel) EventFilterDefault(watched std_core.QObject_ITF, event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.ResultViewModelf82535_EventFilterDefault(ptr.Pointer(), std_core.PointerFromQObject(watched), std_core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackResultViewModelf82535_ObjectNameChanged
func callbackResultViewModelf82535_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackResultViewModelf82535_TimerEvent
func callbackResultViewModelf82535_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*std_core.QTimerEvent))(signal))(std_core.NewQTimerEventFromPointer(event))
	} else {
		NewResultViewModelFromPointer(ptr).TimerEventDefault(std_core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *ResultViewModel) TimerEventDefault(event std_core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.ResultViewModelf82535_TimerEventDefault(ptr.Pointer(), std_core.PointerFromQTimerEvent(event))
	}
}
