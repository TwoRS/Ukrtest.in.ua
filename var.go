package main
 
import(

)

var UsersList = make(map[int ] UserType)   // Список пользователей
var CounterConnections = NumberConnectionType{counters: 0} //Количество сокет подключений
var ActiveClients = make(map[uint64 ] ClientConnType) //Активный соккет клиента

