<script>
    import Basic from "../../components/widgets/Basic.svelte";
	import Breadcrumbs from "../../components/Breadcrumbs.svelte";
	import Layout from "../../components/Layout.svelte";
</script>

<Layout>
	{#snippet header()}
		<Breadcrumbs items={[{name: 'System'}]}></Breadcrumbs>
	{/snippet}
	{#snippet main()}
		<div class="grid">
			<div style="grid-area: a">
				<Basic page={{href: '/system/logs', name: 'Logi'}}>
					{#snippet main()}
						...
					{/snippet}
				</Basic>
			</div>
			
			<div style="grid-area: b">
				<Basic page={{href: '/system/sessions', name: 'Sesje'}}>
					{#snippet main()}
						...
					{/snippet}
				</Basic>
			</div>
		</div>
	{/snippet}
</Layout>
	
<style lang="scss">
    .grid {
        display: grid;
        grid-template-columns: repeat(auto-fit, minmax(100px, 1fr));
        grid-auto-rows: 1fr;
        gap: 10px;
        grid-template-areas:
            "a a b b b"
            "a a b b b";
    }
</style>