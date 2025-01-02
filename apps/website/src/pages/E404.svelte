<script>
    import { Link } from "svelte-routing";
</script>
<div class="e404">
    <div>
        Strona nie została znaleziona.<br/>
        <Link to="/">Wróć do strony głównej</Link>
    </div>
</div>

<style lang="scss">
    .e404 {
        width: 100vw;
        height: 100vh;
        display: flex;
        align-items: center;
        justify-content: center;
        text-align: center;
        background-color: var(--clr-secondary);
        color: #EBEBEB;

        :global(a) {
            color: #42a5e0
        }
    }
</style>