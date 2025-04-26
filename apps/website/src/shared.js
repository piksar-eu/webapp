import { writable, get } from 'svelte/store';

const alert = writable(undefined);
const user = writable(globalThis.user);

const isLoggedIn = () => {
    return get(user) !== undefined
}

export { alert, user, isLoggedIn }