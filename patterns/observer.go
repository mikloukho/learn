package patterns

// Observer.
// Subject хранит []Observer и вызывает Notify.
// Альтернатива по-гошному: каналы вместо колбэков.
//
// TODO: интерфейс Observer с Update(event string), Subject с Subscribe/Unsubscribe/Notify.
