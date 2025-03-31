import { writable } from 'svelte/store';

const createLocalStorageStore = (key, initialValue) => {
    if (!import.meta.env.SSR) {
        const storedValue = localStorage.getItem(key);
        initialValue = storedValue ? JSON.parse(storedValue) : initialValue;
    }
    
    const store = writable(initialValue);

    if (!import.meta.env.SSR) {
        store.subscribe(val => {
            if (val !== undefined) {
                localStorage.setItem(key, JSON.stringify(val));
            }
        });
    }

    return store;
}

const nav = createLocalStorageStore('nav', {
    pinned: false,
    items: {}
});

const alert = writable(undefined);
const user = writable(globalThis.user);

const can = (...permissions) => {
    const userPermissions = globalThis.user?.permissions ?? [];
    return permissions.some(permission => 
        userPermissions.includes(permission)
    );
}

export { alert, user, nav, can }