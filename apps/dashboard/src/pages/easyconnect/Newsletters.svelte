<script>
	import Breadcrumbs from "../../components/Breadcrumbs.svelte";
	import Layout from "../../components/Layout.svelte";
</script>

<Layout>
	{#snippet header()}
	<Breadcrumbs items={[{href: '/easyconnect', name: 'EasyConnect'}, {name: 'Newslettery'}]}></Breadcrumbs>
	{/snippet}
	{#snippet main()}
		<div class="content">

		</div>
	{/snippet}
</Layout>
	
<style lang="scss">
	.content {
		background-color: #fff;
		min-height: 400px;
		padding: 20px;
	}
</style>