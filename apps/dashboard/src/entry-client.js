import './assets/app.scss'
import { hydrate } from 'svelte'
import App from './App.svelte'

hydrate(App, {
  target: document.getElementById('app'),
  props: {
    url: window.location.pathname
  }
})
