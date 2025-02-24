<script>
	import { Link } from "svelte-routing";
    import Navbar from "./Navbar.svelte";

    import { nav as navState } from "../store"
    let { header = undefined, main } = $props();
</script>

<input type="checkbox" id="sidebar__checkbox" checked={$navState.pinned} onchange={() => $navState.pinned = !$navState.pinned }>
<nav class="sidebar">
    <label for="sidebar__checkbox" class="sidebar__toggle">
        <svg class="unpin" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 640 512"><!--!Font Awesome Free 6.7.2 by @fontawesome - https://fontawesome.com License - https://fontawesome.com/license/free Copyright 2025 Fonticons, Inc.--><path d="M38.8 5.1C28.4-3.1 13.3-1.2 5.1 9.2S-1.2 34.7 9.2 42.9l592 464c10.4 8.2 25.5 6.3 33.7-4.1s6.3-25.5-4.1-33.7L481.4 352c9.8-.4 18.9-5.3 24.6-13.3c6-8.3 7.7-19.1 4.4-28.8l-1-3c-13.8-41.5-42.8-74.8-79.5-94.7L418.5 64 448 64c17.7 0 32-14.3 32-32s-14.3-32-32-32L192 0c-17.7 0-32 14.3-32 32s14.3 32 32 32l29.5 0-6.1 79.5L38.8 5.1zM324.9 352L177.1 235.6c-20.9 18.9-37.2 43.3-46.5 71.3l-1 3c-3.3 9.8-1.6 20.5 4.4 28.8s15.7 13.3 26 13.3l164.9 0zM288 384l0 96c0 17.7 14.3 32 32 32s32-14.3 32-32l0-96-64 0z"/></svg>
        <svg class="pin" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 384 512"><!--!Font Awesome Free 6.7.2 by @fontawesome - https://fontawesome.com License - https://fontawesome.com/license/free Copyright 2025 Fonticons, Inc.--><path d="M32 32C32 14.3 46.3 0 64 0L320 0c17.7 0 32 14.3 32 32s-14.3 32-32 32l-29.5 0 11.4 148.2c36.7 19.9 65.7 53.2 79.5 94.7l1 3c3.3 9.8 1.6 20.5-4.4 28.8s-15.7 13.3-26 13.3L32 352c-10.3 0-19.9-4.9-26-13.3s-7.7-19.1-4.4-28.8l1-3c13.8-41.5 42.8-74.8 79.5-94.7L93.5 64 64 64C46.3 64 32 49.7 32 32zM160 384l64 0 0 96c0 17.7-14.3 32-32 32s-32-14.3-32-32l0-96z"/></svg>
    </label>
    <ul class="sidebar__menu">
        <li class="sidebar__menu-item">
            <div class="sidebar__menu-item-link">
                <div class="sidebar__menu-item-link-icon">
                    <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512"><!--!Font Awesome Free 6.7.2 by @fontawesome - https://fontawesome.com License - https://fontawesome.com/license/free Copyright 2025 Fonticons, Inc.--><path d="M0 256a256 256 0 1 1 512 0A256 256 0 1 1 0 256zm320 96c0-15.9-5.8-30.4-15.3-41.6l76.6-147.4c6.1-11.8 1.5-26.3-10.2-32.4s-26.2-1.5-32.4 10.2L262.1 288.3c-2-.2-4-.3-6.1-.3c-35.3 0-64 28.7-64 64s28.7 64 64 64s64-28.7 64-64z"/></svg>
                </div>
            </div>
        </li>

        <li class="sidebar__menu-item">
            <Link to="/">
                <div class="sidebar__menu-item-link">
                    <div class="sidebar__menu-item-link-icon">
                        <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 576 512"><path d="M543.8 287.6c17 0 32-14 32-32.1c1-9-3-17-11-24L512 185l0-121c0-17.7-14.3-32-32-32l-32 0c-17.7 0-32 14.3-32 32l0 36.7L309.5 7c-6-5-14-7-21-7s-15 1-22 8L10 231.5c-7 7-10 15-10 24c0 18 14 32.1 32 32.1l32 0 0 69.7c-.1 .9-.1 1.8-.1 2.8l0 112c0 22.1 17.9 40 40 40l16 0c1.2 0 2.4-.1 3.6-.2c1.5 .1 3 .2 4.5 .2l31.9 0 24 0c22.1 0 40-17.9 40-40l0-24 0-64c0-17.7 14.3-32 32-32l64 0c17.7 0 32 14.3 32 32l0 64 0 24c0 22.1 17.9 40 40 40l24 0 32.5 0c1.4 0 2.8 0 4.2-.1c1.1 .1 2.2 .1 3.3 .1l16 0c22.1 0 40-17.9 40-40l0-16.2c.3-2.6 .5-5.3 .5-8.1l-.7-160.2 32 0z"/></svg>
                    </div>
                    <span class="sidebar__menu-item-link-label">Panel</span>
                </div>
            </Link>
        </li>

        <li class="sidebar__menu-item">
            <Link to="/easyconnect">
                <div class="sidebar__menu-item-link">
                    <div class="sidebar__menu-item-link-icon">
                        <svg xmlns="http://www.w3.org/2000/svg" fill="#2B9EB3" viewBox="0 0 640 512"><!--!Font Awesome Free 6.7.2 by @fontawesome - https://fontawesome.com License - https://fontawesome.com/license/free Copyright 2025 Fonticons, Inc.--><path d="M416 176c0 97.2-93.1 176-208 176c-38.2 0-73.9-8.7-104.7-23.9c-7.5 4-16 7.9-25.2 11.4C59.8 346.4 37.8 352 16 352c-6.9 0-13.1-4.5-15.2-11.1s.2-13.8 5.8-17.9c0 0 0 0 0 0s0 0 0 0l.2-.2c.2-.2 .6-.4 1.1-.8c1-.8 2.5-2 4.3-3.7c3.6-3.3 8.5-8.1 13.3-14.3c5.5-7 10.7-15.4 14.2-24.7C14.7 250.3 0 214.6 0 176C0 78.8 93.1 0 208 0S416 78.8 416 176zM231.5 383C348.9 372.9 448 288.3 448 176c0-5.2-.2-10.4-.6-15.5C555.1 167.1 640 243.2 640 336c0 38.6-14.7 74.3-39.6 103.4c3.5 9.4 8.7 17.7 14.2 24.7c4.8 6.2 9.7 11 13.3 14.3c1.8 1.6 3.3 2.9 4.3 3.7c.5 .4 .9 .7 1.1 .8l.2 .2s0 0 0 0s0 0 0 0c5.6 4.1 7.9 11.3 5.8 17.9c-2.1 6.6-8.3 11.1-15.2 11.1c-21.8 0-43.8-5.6-62.1-12.5c-9.2-3.5-17.8-7.4-25.2-11.4C505.9 503.3 470.2 512 432 512c-95.6 0-176.2-54.6-200.5-129zM228 72c0-11-9-20-20-20s-20 9-20 20l0 14c-7.6 1.7-15.2 4.4-22.2 8.5c-13.9 8.3-25.9 22.8-25.8 43.9c.1 20.3 12 33.1 24.7 40.7c11 6.6 24.7 10.8 35.6 14l1.7 .5c12.6 3.8 21.8 6.8 28 10.7c5.1 3.2 5.8 5.4 5.9 8.2c.1 5-1.8 8-5.9 10.5c-5 3.1-12.9 5-21.4 4.7c-11.1-.4-21.5-3.9-35.1-8.5c-2.3-.8-4.7-1.6-7.2-2.4c-10.5-3.5-21.8 2.2-25.3 12.6s2.2 21.8 12.6 25.3c1.9 .6 4 1.3 6.1 2.1c0 0 0 0 0 0s0 0 0 0c8.3 2.9 17.9 6.2 28.2 8.4l0 14.6c0 11 9 20 20 20s20-9 20-20l0-13.8c8-1.7 16-4.5 23.2-9c14.3-8.9 25.1-24.1 24.8-45c-.3-20.3-11.7-33.4-24.6-41.6c-11.5-7.2-25.9-11.6-37.1-15l-.7-.2c-12.8-3.9-21.9-6.7-28.3-10.5c-5.2-3.1-5.3-4.9-5.3-6.7c0-3.7 1.4-6.5 6.2-9.3c5.4-3.2 13.6-5.1 21.5-5c9.6 .1 20.2 2.2 31.2 5.2c10.7 2.8 21.6-3.5 24.5-14.2s-3.5-21.6-14.2-24.5c-6.5-1.7-13.7-3.4-21.1-4.7L228 72z"/></svg>
                    </div>
                    <span class="sidebar__menu-item-link-label">EasyConnect</span>
                </div>
            </Link>
            <input type="checkbox" id="menu_item_easyconnect" class="sidebar__menu-item-checkbox" checked={$navState.items["easyConnect"]} onchange={() => $navState.items["easyConnect"] = !$navState.items["easyConnect"]}>
            <label for="menu_item_easyconnect" class="sidebar__menu-item-toggle">
                <svg class="show" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 320 512"><!--!Font Awesome Free 6.7.2 by @fontawesome - https://fontawesome.com License - https://fontawesome.com/license/free Copyright 2025 Fonticons, Inc.--><path d="M310.6 233.4c12.5 12.5 12.5 32.8 0 45.3l-192 192c-12.5 12.5-32.8 12.5-45.3 0s-12.5-32.8 0-45.3L242.7 256 73.4 86.6c-12.5-12.5-12.5-32.8 0-45.3s32.8-12.5 45.3 0l192 192z"/></svg>
                <svg class="hide" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512"><!--!Font Awesome Free 6.7.2 by @fontawesome - https://fontawesome.com License - https://fontawesome.com/license/free Copyright 2025 Fonticons, Inc.--><path d="M233.4 406.6c12.5 12.5 32.8 12.5 45.3 0l192-192c12.5-12.5 12.5-32.8 0-45.3s-32.8-12.5-45.3 0L256 338.7 86.6 169.4c-12.5-12.5-32.8-12.5-45.3 0s-12.5 32.8 0 45.3l192 192z"/></svg>
            </label>
            <ul class="sidebar__menu-nested">
                <li><Link to="/easyconnect/leads">Kontakty</Link></li>
                <li><Link to="/easyconnect/newsletters">Newslettery</Link></li>
            </ul>
        </li>

        <li class="sidebar__menu-item">
            <Link to="/system">
                <div class="sidebar__menu-item-link">
                    <div class="sidebar__menu-item-link-icon">
                        <svg xmlns="http://www.w3.org/2000/svg" fill="#C44900" viewBox="0 0 512 512"><!--!Font Awesome Free 6.7.2 by @fontawesome - https://fontawesome.com License - https://fontawesome.com/license/free Copyright 2025 Fonticons, Inc.--><path d="M495.9 166.6c3.2 8.7 .5 18.4-6.4 24.6l-43.3 39.4c1.1 8.3 1.7 16.8 1.7 25.4s-.6 17.1-1.7 25.4l43.3 39.4c6.9 6.2 9.6 15.9 6.4 24.6c-4.4 11.9-9.7 23.3-15.8 34.3l-4.7 8.1c-6.6 11-14 21.4-22.1 31.2c-5.9 7.2-15.7 9.6-24.5 6.8l-55.7-17.7c-13.4 10.3-28.2 18.9-44 25.4l-12.5 57.1c-2 9.1-9 16.3-18.2 17.8c-13.8 2.3-28 3.5-42.5 3.5s-28.7-1.2-42.5-3.5c-9.2-1.5-16.2-8.7-18.2-17.8l-12.5-57.1c-15.8-6.5-30.6-15.1-44-25.4L83.1 425.9c-8.8 2.8-18.6 .3-24.5-6.8c-8.1-9.8-15.5-20.2-22.1-31.2l-4.7-8.1c-6.1-11-11.4-22.4-15.8-34.3c-3.2-8.7-.5-18.4 6.4-24.6l43.3-39.4C64.6 273.1 64 264.6 64 256s.6-17.1 1.7-25.4L22.4 191.2c-6.9-6.2-9.6-15.9-6.4-24.6c4.4-11.9 9.7-23.3 15.8-34.3l4.7-8.1c6.6-11 14-21.4 22.1-31.2c5.9-7.2 15.7-9.6 24.5-6.8l55.7 17.7c13.4-10.3 28.2-18.9 44-25.4l12.5-57.1c2-9.1 9-16.3 18.2-17.8C227.3 1.2 241.5 0 256 0s28.7 1.2 42.5 3.5c9.2 1.5 16.2 8.7 18.2 17.8l12.5 57.1c15.8 6.5 30.6 15.1 44 25.4l55.7-17.7c8.8-2.8 18.6-.3 24.5 6.8c8.1 9.8 15.5 20.2 22.1 31.2l4.7 8.1c6.1 11 11.4 22.4 15.8 34.3zM256 336a80 80 0 1 0 0-160 80 80 0 1 0 0 160z"/></svg>
                    </div>
                    <span class="sidebar__menu-item-link-label">System</span>
                </div>
            </Link>
            <input type="checkbox" id="menu_item_system" class="sidebar__menu-item-checkbox" checked={$navState.items["system"]} onchange={() => $navState.items["system"] = !$navState.items["system"] }>
            <label for="menu_item_system" class="sidebar__menu-item-toggle">
                <svg class="show" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 320 512"><!--!Font Awesome Free 6.7.2 by @fontawesome - https://fontawesome.com License - https://fontawesome.com/license/free Copyright 2025 Fonticons, Inc.--><path d="M310.6 233.4c12.5 12.5 12.5 32.8 0 45.3l-192 192c-12.5 12.5-32.8 12.5-45.3 0s-12.5-32.8 0-45.3L242.7 256 73.4 86.6c-12.5-12.5-12.5-32.8 0-45.3s32.8-12.5 45.3 0l192 192z"/></svg>
                <svg class="hide" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512"><!--!Font Awesome Free 6.7.2 by @fontawesome - https://fontawesome.com License - https://fontawesome.com/license/free Copyright 2025 Fonticons, Inc.--><path d="M233.4 406.6c12.5 12.5 32.8 12.5 45.3 0l192-192c12.5-12.5 12.5-32.8 0-45.3s-32.8-12.5-45.3 0L256 338.7 86.6 169.4c-12.5-12.5-32.8-12.5-45.3 0s-12.5 32.8 0 45.3l192 192z"/></svg>
            </label>
            <ul class="sidebar__menu-nested">
                <li><Link to="/system/sessions">Sesje</Link></li>
                <li><Link to="/system/logs">Logi</Link></li>
            </ul>
        </li>
    </ul>
</nav>
<div id="layout">
    <header>
        <Navbar/>
    </header>
    
    <div id="content">
        <header>
            {@render header?.()}
        </header>
        <main>
            {@render main()}
        </main>
    </div>

    <footer>
        Subskrybuj @PragmatycznyPiotr
    </footer>
</div>

<style lang="scss">
    @mixin sidebar-expanded {
		max-width: 250px;
        .sidebar__menu-item {
            svg.show {
                display: block;
            }
            .sidebar__menu-item-checkbox:checked {
                ~ .sidebar__menu-item-toggle {
                    .show {
                        display: none;
                    }
                    .hide {
                        display: block;
                    }
                }
                ~ .sidebar__menu-nested {
                    display: block;
                }
            }
        }
	}
    #sidebar__checkbox {
        display: none;
    }
    .sidebar {
        height: 100vh;
        position: fixed;
        top: 0;
        bottom: 0;
        background-color: #fff;
        width: 100%;
        max-width: 64px;
        transition: max-width 0.3s ease-in-out;
        padding-top: 16px;
        box-shadow: 1px 0px 2px rgba(0, 0, 0, 0.1); 
        z-index: 20;
        overflow: hidden;

        &:hover {
            @include sidebar-expanded;
        }

        &__toggle {
            display: flex;
            align-items: center;
            justify-content: center;
            position: absolute;
            right: 8px;
            top: 24px;
            width: 2rem;
            cursor: pointer;
            z-index: 50;
            svg {
                height: 1.5rem;
                fill: #707070;
                display: none;
            }
        }

        &__menu-item {
            display: block;
            position: relative;
            min-height: 48px;
            border-bottom: 1px solid #ddd;

            &-checkbox {
                display: none;
            }

            &:first-child {
                margin-bottom: 24px;
                border-bottom: none;

                .sidebar__menu-item-link-icon {
                    padding: 0 8px;
                    svg {
                        fill: var(--clr-primary);
                        max-height: 2rem !important;
                    }
                }

                .sidebar__menu-item-link:hover {
                    background-color: #fff !important;
                }
            }

            &-link {
                position: relative;
                display: flex;
                flex-direction: row;
                flex-wrap: nowrap;
                
                &:hover {
                    background-color: #f5f5f5;
                }
                
                &-icon {
                    width: 64px;
                    height: 48px;
                    display: flex;
                    align-items: center;
                    justify-content: center;
                    padding: 0 16px;
                    flex-shrink: 0;
                    left: 0;

                    svg {
                        max-height: 1rem;
                        width: auto;
                    }
                }

                &-label {
                    font-size: 16px;
                    width: calc(250px - 64px);
                    line-height: 48px;
                    transition: opacity 0.3s, transform 0.3s;
                    overflow: hidden;
                    text-overflow: ellipsis;
                    white-space: nowrap;
                }
            }
            &-toggle {
                position: absolute;
                right: 0;
                top: 0;
                padding: 14px 8px;
                cursor: pointer;

                svg {
                    max-height: 1.2rem;
                    width: 1.2rem;
                    fill: #707070;
                    display: none;
                }
            }
        }
        &__menu-nested {
			display: none;
            list-style-type: none;

            li {
                width: 100%;
				min-height: 2rem;
				line-height: 2rem;
                padding-left: 64px;
                &:hover {
                    background-color: #f5f5f5;
                }
                :global(a) {
                    font-weight: 400;
                    display: block;
                    width: 100%;
                    height: 100%;
                }
            }
        }
    }
    @media (min-width: 768px) {
        #sidebar__checkbox:checked {
            ~ .sidebar {
                @include sidebar-expanded;

                .unpin {
                    display: inline-block;
                }
            }

            ~ #layout {
                margin-left: 250px;
            }
        }
        #sidebar__checkbox:not(:checked) {
            ~ .sidebar:hover {
                .pin {
                    display: inline-block;
                }
            }
        }
    }






    #layout {
        display: flex;
        flex-direction: column;
        align-content: stretch;
        min-height: 100vh;
		margin-left: 64px;
    }

    #content {
        flex-grow: 1;

        header {
            padding-bottom: 100px;
        }

        main {
            margin: -80px 20px 0;
            min-height: 100px;
        }
    }

    header {
        background-color: var(--clr-background);
    }

    footer {
        padding: 16px 0;
        display: flex;
        justify-content: center;
        font-size: 0.7rem;
        border-top: 1px solid #e4e6ea;
        margin-top: 20px;
    }
</style>