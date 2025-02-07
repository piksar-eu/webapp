import { writable } from 'svelte/store';

const alert = writable(undefined);
const user = writable(globalThis.user);

export { alert, user }