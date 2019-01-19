### How to use

websocket url: 

`http://adelytics.com:5000`

subscribe to a room: 

`io.emit('join', '<ROOM_NAME>')`

if the room does not exist, the `join` event will create a new room and subscribe to it by default.

Emit and event to a room: 

`io.emit('event', JSON.stringify({event: '<EVENT_NAME>', room: '<ROOM_NAME>', data: '<PAYLOAD>'}))`

Listen to an event:

First, join the room: `io.emit('join', '<ROOM_NAME>')`

Next, start listening to any event in this room: `io.on('<EVENT_NAME>', '<PAYLOAD>')`