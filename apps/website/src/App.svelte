<script>
	import { Router, Route } from "svelte-routing";
	import About from "./pages/About.svelte";
	import Contact from "./pages/Contact.svelte";
	import E404 from "./pages/E404.svelte";
	import Index from "./pages/Index.svelte";
	import PrivacyPolicy from "./pages/PrivacyPolicy.svelte";
	import Login from "./pages/auth/Login.svelte";
	import Register from "./pages/auth/Register.svelte";
	
	let { url } = $props();
</script>

<Router {url}>
	<Route path="/"><Index /></Route>
	<Route path="/o-nas"><About /></Route>
	<Route path="/kontakt"><Contact /></Route>
	<Route path="/polityka-prywatnosci"><PrivacyPolicy /></Route>

	<Route path="/auth/login"><Login /></Route>
	<Route path="/auth/register"><Register /></Route>
	<Route ><E404 /></Route>
</Router>