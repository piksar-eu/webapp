<script>
	import { Router, Route } from "svelte-routing";
	import E404 from "./pages/E404.svelte";
	import Index from "./pages/Index.svelte";
	
	let { url } = $props();
</script>

<Router {url}>
	<Route path="/"><Index /></Route>
	<Route ><E404 /></Route>
</Router>