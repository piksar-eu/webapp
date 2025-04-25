<script>
	import { Router, Route } from "svelte5-router";
	import E404 from "./pages/E404.svelte";
	import Index from "./pages/Index.svelte";
    import EasyConnect from "./pages/easyconnect/EasyConnect.svelte";
    import Leads from "./pages/easyconnect/Leads.svelte";
    import Newsletters from "./pages/easyconnect/Newsletters.svelte";
    import System from "./pages/system/System.svelte";
    import Roles from "./pages/system/Roles.svelte";
	import Role from "./pages/system/Role.svelte";
    import Users from "./pages/system/Users.svelte";
	import User from "./pages/system/User.svelte";
	
	let { url } = $props();
</script>

<Router {url}>
	<Route path="/"><Index /></Route>
	<Route path="/easyconnect"><EasyConnect /></Route>
	<Route path="/easyconnect/leads"><Leads /></Route>
	<Route path="/easyconnect/newsletters"><Newsletters /></Route>
	<Route path="/system"><System /></Route>
	<Route path="/system/roles"><Roles /></Route>
	<Route path="/system/role/add"><Role id={undefined}/></Route>
	<Route path="/system/role/:id/edit" let:params><Role id={params.id}/></Route>
	<Route path="/system/users"><Users /></Route>
	<Route path="/system/users/:id/edit" let:params><User id={params.id}/></Route>

	<Route ><E404 /></Route>
</Router>