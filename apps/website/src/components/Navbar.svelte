<script>
    import { Link } from "svelte-routing";
</script>

<nav class="nav">
    <ul class="nav__menu">
        <li><Link to="/"><img src="/website.png" alt="logo" height="48"/></Link></li>
        <li><Link to="/o-nas">O nas</Link></li>
        <li><Link to="/kontakt">Kontakt</Link></li>
    </ul>
</nav>

<style lang="scss">
    .nav {
        &__menu {
            display: flex;
            flex-direction: row;
            gap: 2rem;

            li {
                color: var(--clr-primary);
                list-style: none;
                display: block;
                position: relative;
                padding: 1rem 0;
                height: 4rem;

                &:first-child {
                    margin-right: auto;
                }

                :global(a) {
                    color: inherit;
                    text-decoration: none;
                    font-size: 1.0rem;
                    line-height: 2rem;
                }

                :global(a img) {
                    max-height: 2rem;
                }
            }
        }
    }
</style>