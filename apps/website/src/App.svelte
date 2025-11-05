<script>
	import { Router, Route } from "svelte5-router";
	import About from "./pages/About.svelte";
	import Contact from "./pages/Contact.svelte";
	import E404 from "./pages/E404.svelte";
	import Index from "./pages/Index.svelte";
	import PrivacyPolicy from "./pages/PrivacyPolicy.svelte";
	import Posts from "./pages/Posts.svelte";
	import Post from "./pages/Post.svelte";
	import Login from "./pages/auth/Login.svelte";
	import Register from "./pages/auth/Register.svelte";

	
	let { url } = $props();
</script>

<Router {url}>
	<Route path="/"><Index /></Route>
	<Route path="/o-mnie"><About /></Route>
	<Route path="/posty"><Posts /></Route>
	<Route path="/post/:slug">{#snippet children(params)}<Post slug={params.slug}/>{/snippet}</Route>
	<Route path="/polityka-prywatnosci"><PrivacyPolicy /></Route>

	<Route path="/auth/login"><Login /></Route>
	<Route path="/auth/register"><Register /></Route>
	<Route ><E404 /></Route>
</Router>