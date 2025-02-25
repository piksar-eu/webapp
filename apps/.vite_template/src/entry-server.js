import { render as _render } from 'svelte/server'
import App from './App.svelte'

/**
 * @param {string} _url
 */
function render(_url) {
  return _render(App, {props: {url: _url} })
}

globalThis.render = render