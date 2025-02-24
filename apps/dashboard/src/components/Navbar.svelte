<script>
    import { Link } from "svelte-routing";
	import { user } from '../store';
    import { logout as logoutApiCall } from "../api";

    const logout = async () => {
		await logoutApiCall()
		user.set(undefined)
    }
</script>

<nav class="nav">
    <ul class="nav__menu">
        <li></li>
        <li class="nav__profile">
            <div class="nav__profile-container">
                <div class="nav__profile-img-mask">
                    <img src="/user.svg" alt=""/>
                </div>
                <div class="nav__profile-username">John Doe</div>
            </div>
            <ul class="nav__dropdown">
                <li><Link to="/">Profil</Link></li>
                <li><Link to="/">Ustawienia</Link></li>
                <li class="dropdown-divider"></li>
                <li><Link to="/">Wyloguj</Link></li>
            </ul>
        </li>
    </ul>
</nav>

<style lang="scss">
    .nav {
        &__menu {
            display: flex;
            flex-direction: row;
            gap: 2rem;
            padding: 0 32px;
            border-bottom: 1px solid #7a86a2;

            :global(a) {
                display: block;
                width: 100%;
                height: 100%;
            }

            li {
                list-style: none;
                display: block;
                position: relative;
                padding: 1rem 0;
                height: 5rem;

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

        &__profile {
            padding: 1.4rem 0 !important;
            position: relative;
            min-width: 120px;
            max-width: 200px;
            &-container {
                display: flex;
                align-items: center;
                gap: 10px; /* Space between image and text */
            }

            &-img-mask {
                min-width: 2.2rem; /* Adjust size as needed */
                height: 2.2rem;
                background-color: #fff;
                border-radius: 50%;
                display: flex;
                align-items: center;
                justify-content: center;
                overflow: hidden;
                box-shadow: 0 0 5px rgba(0, 0, 0, 0.2);

                img {
                    width: 100%;
                    height: 100%;
                    object-fit: contain;
                }
            }

            &-username {
                overflow: hidden;
                text-overflow: ellipsis;
                font-size: 1rem;
                font-weight: 600;
                white-space: nowrap;
                color: var(--clr-primary);
            }
            .nav__dropdown {
                display: none;
                position: absolute;
                top: 5rem;
                left: 0;
                background-color: #fff;
                box-shadow: 1px 1px 1px rgba(0, 0, 0, 0.1);
                min-width: 120px;
                max-width: 200px;

                li {
                    height: auto;
                    padding: 0;
                    
                    &:hover {
                        background-color: var(--clr-primary);
                    }

                    :global(a) {
                        display: block;
                        width: 100%;
                        height: 100%;
                        padding: 5px 10px;
                    }
                }

                li.dropdown-divider {
                    padding: 0;
                    margin: 0;
                    height: 1px;
                    background-color: var(--clr-primary);
                }
            }

            &:hover {
                .nav__dropdown {
                    display: block;
                }
            }
        }
    }
</style>