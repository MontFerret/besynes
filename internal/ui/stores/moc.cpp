

#define protected public
#define private public

#include "moc.h"
#include "_cgo_export.h"

#include <QByteArray>
#include <QChildEvent>
#include <QEvent>
#include <QMetaMethod>
#include <QMetaObject>
#include <QObject>
#include <QOffscreenSurface>
#include <QPaintDeviceWindow>
#include <QPdfWriter>
#include <QString>
#include <QTimerEvent>
#include <QWindow>

#ifdef QT_QML_LIB
	#include <QQmlEngine>
#endif


class ExecutionStoref82535: public QObject
{
Q_OBJECT
public:
	ExecutionStoref82535(QObject *parent = Q_NULLPTR) : QObject(parent) {qRegisterMetaType<quintptr>("quintptr");ExecutionStoref82535_ExecutionStoref82535_QRegisterMetaType();ExecutionStoref82535_ExecutionStoref82535_QRegisterMetaTypes();callbackExecutionStoref82535_Constructor(this);};
	 ~ExecutionStoref82535() { callbackExecutionStoref82535_DestroyExecutionStore(this); };
	void childEvent(QChildEvent * event) { callbackExecutionStoref82535_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackExecutionStoref82535_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackExecutionStoref82535_CustomEvent(this, event); };
	void deleteLater() { callbackExecutionStoref82535_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackExecutionStoref82535_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackExecutionStoref82535_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackExecutionStoref82535_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackExecutionStoref82535_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackExecutionStoref82535_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackExecutionStoref82535_TimerEvent(this, event); };
signals:
public slots:
	ResultViewModelf82535* execute(QString addr) { QByteArray t00926c = addr.toUtf8(); Moc_PackedString addrPacked = { const_cast<char*>(t00926c.prepend("WHITESPACE").constData()+10), t00926c.size()-10 };return static_cast<ResultViewModelf82535*>(callbackExecutionStoref82535_Execute(this, addrPacked)); };
private:
};

Q_DECLARE_METATYPE(ExecutionStoref82535*)


void ExecutionStoref82535_ExecutionStoref82535_QRegisterMetaTypes() {
}

class ResultViewModelf82535: public QObject
{
Q_OBJECT
public:
	ResultViewModelf82535(QObject *parent = Q_NULLPTR) : QObject(parent) {qRegisterMetaType<quintptr>("quintptr");ResultViewModelf82535_ResultViewModelf82535_QRegisterMetaType();ResultViewModelf82535_ResultViewModelf82535_QRegisterMetaTypes();callbackResultViewModelf82535_Constructor(this);};
	 ~ResultViewModelf82535() { callbackResultViewModelf82535_DestroyResultViewModel(this); };
	void childEvent(QChildEvent * event) { callbackResultViewModelf82535_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackResultViewModelf82535_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackResultViewModelf82535_CustomEvent(this, event); };
	void deleteLater() { callbackResultViewModelf82535_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackResultViewModelf82535_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackResultViewModelf82535_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackResultViewModelf82535_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackResultViewModelf82535_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackResultViewModelf82535_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackResultViewModelf82535_TimerEvent(this, event); };
signals:
public slots:
private:
};

Q_DECLARE_METATYPE(ResultViewModelf82535*)


void ResultViewModelf82535_ResultViewModelf82535_QRegisterMetaTypes() {
}

void* ExecutionStoref82535_Execute(void* ptr, struct Moc_PackedString addr)
{
	ResultViewModelf82535* returnArg;
	QMetaObject::invokeMethod(static_cast<ExecutionStoref82535*>(ptr), "execute", Q_RETURN_ARG(ResultViewModelf82535*, returnArg), Q_ARG(QString, QString::fromUtf8(addr.data, addr.len)));
	return returnArg;
}

int ExecutionStoref82535_ExecutionStoref82535_QRegisterMetaType()
{
	return qRegisterMetaType<ExecutionStoref82535*>();
}

int ExecutionStoref82535_ExecutionStoref82535_QRegisterMetaType2(char* typeName)
{
	return qRegisterMetaType<ExecutionStoref82535*>(const_cast<const char*>(typeName));
}

int ExecutionStoref82535_ExecutionStoref82535_QmlRegisterType()
{
#ifdef QT_QML_LIB
	return qmlRegisterType<ExecutionStoref82535>();
#else
	return 0;
#endif
}

int ExecutionStoref82535_ExecutionStoref82535_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName)
{
#ifdef QT_QML_LIB
	return qmlRegisterType<ExecutionStoref82535>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName));
#else
	return 0;
#endif
}

void* ExecutionStoref82535___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void ExecutionStoref82535___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* ExecutionStoref82535___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* ExecutionStoref82535___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void ExecutionStoref82535___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* ExecutionStoref82535___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* ExecutionStoref82535___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void ExecutionStoref82535___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* ExecutionStoref82535___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* ExecutionStoref82535___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void ExecutionStoref82535___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* ExecutionStoref82535___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* ExecutionStoref82535___qFindChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void ExecutionStoref82535___qFindChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* ExecutionStoref82535___qFindChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* ExecutionStoref82535_NewExecutionStore(void* parent)
{
	if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new ExecutionStoref82535(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new ExecutionStoref82535(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new ExecutionStoref82535(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new ExecutionStoref82535(static_cast<QWindow*>(parent));
	} else {
		return new ExecutionStoref82535(static_cast<QObject*>(parent));
	}
}

void ExecutionStoref82535_DestroyExecutionStore(void* ptr)
{
	static_cast<ExecutionStoref82535*>(ptr)->~ExecutionStoref82535();
}

void ExecutionStoref82535_DestroyExecutionStoreDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void ExecutionStoref82535_ChildEventDefault(void* ptr, void* event)
{
	static_cast<ExecutionStoref82535*>(ptr)->QObject::childEvent(static_cast<QChildEvent*>(event));
}

void ExecutionStoref82535_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<ExecutionStoref82535*>(ptr)->QObject::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void ExecutionStoref82535_CustomEventDefault(void* ptr, void* event)
{
	static_cast<ExecutionStoref82535*>(ptr)->QObject::customEvent(static_cast<QEvent*>(event));
}

void ExecutionStoref82535_DeleteLaterDefault(void* ptr)
{
	static_cast<ExecutionStoref82535*>(ptr)->QObject::deleteLater();
}

void ExecutionStoref82535_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<ExecutionStoref82535*>(ptr)->QObject::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char ExecutionStoref82535_EventDefault(void* ptr, void* e)
{
	return static_cast<ExecutionStoref82535*>(ptr)->QObject::event(static_cast<QEvent*>(e));
}

char ExecutionStoref82535_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<ExecutionStoref82535*>(ptr)->QObject::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}



void ExecutionStoref82535_TimerEventDefault(void* ptr, void* event)
{
	static_cast<ExecutionStoref82535*>(ptr)->QObject::timerEvent(static_cast<QTimerEvent*>(event));
}

int ResultViewModelf82535_ResultViewModelf82535_QRegisterMetaType()
{
	return qRegisterMetaType<ResultViewModelf82535*>();
}

int ResultViewModelf82535_ResultViewModelf82535_QRegisterMetaType2(char* typeName)
{
	return qRegisterMetaType<ResultViewModelf82535*>(const_cast<const char*>(typeName));
}

int ResultViewModelf82535_ResultViewModelf82535_QmlRegisterType()
{
#ifdef QT_QML_LIB
	return qmlRegisterType<ResultViewModelf82535>();
#else
	return 0;
#endif
}

int ResultViewModelf82535_ResultViewModelf82535_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName)
{
#ifdef QT_QML_LIB
	return qmlRegisterType<ResultViewModelf82535>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName));
#else
	return 0;
#endif
}

void* ResultViewModelf82535___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void ResultViewModelf82535___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* ResultViewModelf82535___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* ResultViewModelf82535___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void ResultViewModelf82535___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* ResultViewModelf82535___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* ResultViewModelf82535___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void ResultViewModelf82535___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* ResultViewModelf82535___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* ResultViewModelf82535___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void ResultViewModelf82535___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* ResultViewModelf82535___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* ResultViewModelf82535___qFindChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void ResultViewModelf82535___qFindChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* ResultViewModelf82535___qFindChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* ResultViewModelf82535_NewResultViewModel(void* parent)
{
	if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new ResultViewModelf82535(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new ResultViewModelf82535(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new ResultViewModelf82535(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new ResultViewModelf82535(static_cast<QWindow*>(parent));
	} else {
		return new ResultViewModelf82535(static_cast<QObject*>(parent));
	}
}

void ResultViewModelf82535_DestroyResultViewModel(void* ptr)
{
	static_cast<ResultViewModelf82535*>(ptr)->~ResultViewModelf82535();
}

void ResultViewModelf82535_DestroyResultViewModelDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void ResultViewModelf82535_ChildEventDefault(void* ptr, void* event)
{
	static_cast<ResultViewModelf82535*>(ptr)->QObject::childEvent(static_cast<QChildEvent*>(event));
}

void ResultViewModelf82535_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<ResultViewModelf82535*>(ptr)->QObject::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void ResultViewModelf82535_CustomEventDefault(void* ptr, void* event)
{
	static_cast<ResultViewModelf82535*>(ptr)->QObject::customEvent(static_cast<QEvent*>(event));
}

void ResultViewModelf82535_DeleteLaterDefault(void* ptr)
{
	static_cast<ResultViewModelf82535*>(ptr)->QObject::deleteLater();
}

void ResultViewModelf82535_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<ResultViewModelf82535*>(ptr)->QObject::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char ResultViewModelf82535_EventDefault(void* ptr, void* e)
{
	return static_cast<ResultViewModelf82535*>(ptr)->QObject::event(static_cast<QEvent*>(e));
}

char ResultViewModelf82535_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<ResultViewModelf82535*>(ptr)->QObject::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}



void ResultViewModelf82535_TimerEventDefault(void* ptr, void* event)
{
	static_cast<ResultViewModelf82535*>(ptr)->QObject::timerEvent(static_cast<QTimerEvent*>(event));
}

#include "moc_moc.h"
