<script>
    import { link } from "svelte5-router";
	import { user } from '../shared';
    import { logout as logoutApiCall } from "../api";

    const logout = async () => {
		await logoutApiCall()
		user.set(undefined)
    }
</script>

<nav class="nav">
    <ul class="nav__menu">
        <li><a use:link href="/"><img src="/logo2.svg" alt="logo"/></a></li>
        <li><a use:link href="/posty">Posty</a></li>
        <!-- <li><a use:link href="/o-mnie">O mnie</a></li> -->

        {#if $user !== undefined}
            <li><a href="#" on:click={logout}>Wyloguj</a></li>
        {/if}
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
                padding: 0.5rem 0;
                height: 4rem;

                &:first-child {
                    margin-right: auto;

                    img {
                        height: 3rem;
                    }
                }

                :global(a) {
                    color: inherit;
                    text-decoration: none;
                    font-size: 1rem;
                    line-height: 3rem;
                }

                :global(a img) {
                    max-height: 3rem;
                }
            }
        }
    }
</style>