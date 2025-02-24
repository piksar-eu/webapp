<script>
    import { Link } from "svelte-routing";
    let { page, main, style = 'light' } = $props();

</script>
<div class="widget" class:widget--dark={style && style=='dark'}>
    <header>
        <h3>{page.name}</h3>
        {#if page?.href}
            <Link class="button button--sm" to={page.href}>więcej</Link>
        {/if}
    </header>

    {@render main?.()}
</div>

<style lang="scss">
    .widget {
        background-color: #fff;
        box-shadow: 1px 0px 2px rgba(0, 0, 0, 0.1); 
        padding: 20px;
        min-height: 120px;
        height: 100%;

        header {
            display: flex;
            justify-content: space-between;
        }
    }
    .widget--dark {
        background-color: #2c3249;
        color: #fff;
    }
</style>