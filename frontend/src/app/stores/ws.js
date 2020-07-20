'use strict';

export function makeWebsocketStore(url, socketOptions, initialValue={}) {
    let socket;
    let reopenCount;
    const subscriptions = new Set();
    let reopenTimeoutHandler;
    const reopenTimeouts = [2000, 5000, 10000, 30000, 60000];

    function reopenTimeout() {
        const n = reopenCount;
        reopenCount++;
        return reopenTimeout[
            n >= reopenTimeouts.length -1 ? reopenTimeouts.length -1 : n
        ];
    }

    function close() {
        if (reopenTimeoutHandler) {
            clearTimeout(reopenTimeoutHandler);
        }

        if (socket) {
            socket.close();
            socket = undefined;
        }
    }

    function reopen() {
        close();
        if(subscriptions.size > 0) {
            reopenTimeoutHandler = setTimeout(open, reopenTimeout());
        }
    }

    async function open() {
        if(reopenTimeoutHandler) {
            clearTimeout(reopenTimeoutHandler);
            reopenTimeoutHandler = undefined;
        }

        if (socket) return;
        socket = new WebSocket(url, socketOptions);

        socket.onmessage = event =>{
            initialValue = JSON.parse(event.data);
            subscriptions.forEach(subscription=>subscription(initialValue));
        };

        socket.onclose = event => reopen();

        return new Promise((resolve, reject)=>{
            socket.onopen = event =>{
                reopenCount = 0;
                resolve();
            };
        });
    }

    return {
        set(value) {
            open().then(_=>socket.send(JSON.stringify(value)));
        },
        subscribe(subscription) {
            open();
            subscription(initialValue);
            subscriptions.add(subscription);
            return _=>{
                subscriptions.delete(subscription);
                if(subscriptions.size === 0) close();
            };
        }
    }
}

export default makeWebsocketStore;