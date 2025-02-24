<script>
	import { Router, Route } from "svelte-routing";
	import E404 from "./pages/E404.svelte";
	import Index from "./pages/Index.svelte";
    import EasyConnect from "./pages/easyconnect/EasyConnect.svelte";
    import Leads from "./pages/easyconnect/Leads.svelte";
    import Newsletters from "./pages/easyconnect/Newsletters.svelte";
    import System from "./pages/system/System.svelte";
    import Logs from "./pages/system/Logs.svelte";
    import Sessions from "./pages/system/Sessions.svelte";
	
	let { url } = $props();
</script>

<Router {url}>
	<Route path="/"><Index /></Route>
	<Route path="/easyconnect"><EasyConnect /></Route>
	<Route path="/easyconnect/leads"><Leads /></Route>
	<Route path="/easyconnect/newsletters"><Newsletters /></Route>
	<Route path="/system"><System /></Route>
	<Route path="/system/logs"><Logs /></Route>
	<Route path="/system/sessions"><Sessions /></Route>

	<Route ><E404 /></Route>
</Router>