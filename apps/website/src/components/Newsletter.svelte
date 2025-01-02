<script>
    import { Link } from "svelte-routing";
</script>

<div class="newsletter">
    <div class="newsletter__head">Zapisz się do newslettera</div>
    <div class="newsletter__input">
        <input type="email" value="" placeholder="Twój email">
        <button class="button">Zapisz się</button>
    </div>
    <div class="newsletter__terms">Klikając „Zapisz się”, wyrażasz zgodę na przetwarzanie Twojego adresu e-mail w celu otrzymywania naszego newslettera zgodnie z&nbsp;<Link to="/polityka-prywatnosci">Polityką&nbsp;Prywatności</Link>.</div>
</div>

<style lang="scss">
    .newsletter {
        max-width: 480px;
        display: flex;
        flex-direction: column;

        &__head {
            padding: 1rem 0;
            font-size: 1.6rem;
            color: var(--clr-primary);
        }

        &__input {
            display: flex;
            
            input {
                padding: 0.6em 1.2em;
                font-size: 1em;
                flex-grow: 1;
            }

            button {
                margin-left: 8px;
            }
        }

        &__terms {
            font-size: 0.6rem;
            padding: 0.5rem 0;
            color: #767373;
        }
    }
</style>