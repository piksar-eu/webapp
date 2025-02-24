<script>
    import Layout from "./Layout.svelte";

    let { header, children } = $props();
</script>

<Layout>
    <div class="container subpage">
        <h1>{header}</h1>
        {@render children?.()}
    </div>
</Layout>

<style lang="scss">
    .subpage {
        padding-bottom: 32px;

        h1 {
            padding: 32px 0;
        }
    }
</style>