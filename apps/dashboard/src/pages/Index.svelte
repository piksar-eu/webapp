<script>
	import Layout from "../components/Layout.svelte";
    import Basic from "../components/widgets/Basic.svelte";
</script>

<Layout>
	{#snippet main()}
        <div class="grid">
            <div style="grid-area: a">
                <Basic page={{href: '/easyconnect', name: 'EasyConnect'}}>
                    {#snippet main()}
                        ...
                    {/snippet}
                </Basic>
            </div>
            <div style="grid-area: b">
                <Basic page={{href: '/system/logs', name: 'Logi'}} style="dark">
                    {#snippet main()}
                        ...
                    {/snippet}
                </Basic>
            </div>
            <div style="grid-area: c">
                <Basic page={{href: '/system/sessions', name: 'Sesje'}}>
                    {#snippet main()}
                        ...
                    {/snippet}
                </Basic>
            </div>
            <div style="grid-area: d">
                <Basic page={{href: '/', name: 'D'}} style="dark">
                    {#snippet main()}
                        ...
                    {/snippet}
                </Basic>
            </div>
            <div style="grid-area: e">
                <Basic page={{href: '/', name: 'E'}}>
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
            "a a a b b b"
            "a a a c c c"
            "d d e e e e";
    }

    @media (max-width: 768px) {
        .grid {
            grid-template-areas:
                "a a b b"
                "a a b b"
                "d c c e";
        }
    }

    @media (max-width: 576px) {
        .grid {
            grid-template-areas:
                "a a"
                "b b"
                "c d"
                "e e";
        }
    }
</style>